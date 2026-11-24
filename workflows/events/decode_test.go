package events

import (
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
)

func TestUnmarshalPaymentEventKeepsDataMap(t *testing.T) {
	body := []byte(`{
		"id": "evt_123",
		"source": "gateway",
		"type": "payment_approved",
		"timestamp": "2026-01-02T03:04:05Z",
		"version": "1.0.0",
		"data": {
			"id": "pay_123",
			"action_id": "act_123",
			"amount": 1500,
			"currency": "USD",
			"approved": true,
			"status": "Authorized"
		},
		"action_invocations": [{
			"workflow_id": "wf_1",
			"workflow_action_id": "wfa_1",
			"status": "successful"
		}]
	}`)

	var event EventResponse
	require.NoError(t, json.Unmarshal(body, &event))

	assert.Equal(t, "evt_123", event.Id)
	assert.Equal(t, "gateway", event.Source)
	assert.Equal(t, EventTypePaymentApproved, event.Type)
	assert.Equal(t, "pay_123", event.Data["id"])
	assert.Equal(t, float64(1500), event.Data["amount"])
	require.Len(t, event.ActionInvocations, 1)
	assert.Equal(t, "wf_1", event.ActionInvocations[0].WorkflowId)
	assert.Equal(t, Successful, event.ActionInvocations[0].Status)
	assert.JSONEq(t, `{
		"id": "pay_123",
		"action_id": "act_123",
		"amount": 1500,
		"currency": "USD",
		"approved": true,
		"status": "Authorized"
	}`, string(event.RawData()))

	payload, err := event.DecodePayment()
	require.NoError(t, err)
	assert.Equal(t, "pay_123", payload.Id)
	assert.Equal(t, "act_123", payload.ActionId)
	assert.Equal(t, int64(1500), payload.Amount)
	assert.Equal(t, common.Currency("USD"), payload.Currency)
	assert.True(t, payload.Approved)
	assert.Equal(t, payments.Authorized, payload.Status)
	assert.Empty(t, payload.Reference)
}

func TestDecodePaymentRejectsUnknownEventType(t *testing.T) {
	body := []byte(`{
		"id": "evt_dispute",
		"type": "dispute_received",
		"data": {"id": "dsp_1", "amount": 2500}
	}`)

	var event EventResponse
	require.NoError(t, json.Unmarshal(body, &event))

	payload, err := event.DecodePayment()
	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrUnexpectedEventType)
	assert.Equal(t, "dsp_1", event.Data["id"])
	assert.JSONEq(t, `{"id":"dsp_1","amount":2500}`, string(event.RawData()))

	var custom struct {
		Id string `json:"id"`
	}
	require.NoError(t, event.DecodeData(&custom))
	assert.Equal(t, "dsp_1", custom.Id)
}

func TestDecodePaymentMissingData(t *testing.T) {
	t.Run("data field absent", func(t *testing.T) {
		var event EventResponse
		require.NoError(t, json.Unmarshal([]byte(`{"id":"evt_1","type":"payment_approved"}`), &event))

		payload, err := event.DecodePayment()
		assert.Nil(t, payload)
		assert.ErrorIs(t, err, ErrMissingEventData)

		var dest PaymentEventData
		assert.ErrorIs(t, event.DecodeData(&dest), ErrMissingEventData)
	})

	t.Run("data is null", func(t *testing.T) {
		var event EventResponse
		require.NoError(t, json.Unmarshal([]byte(`{"id":"evt_1","type":"payment_captured","data":null}`), &event))
		assert.Nil(t, event.Data)

		_, err := event.DecodePayment()
		assert.ErrorIs(t, err, ErrMissingEventData)
	})

	t.Run("payment id absent", func(t *testing.T) {
		var event EventResponse
		require.NoError(t, json.Unmarshal([]byte(`{
			"type": "payment_declined",
			"data": {"amount": 10, "currency": "GBP"}
		}`), &event))

		payload, err := event.DecodePayment()
		assert.Nil(t, payload)
		assert.ErrorIs(t, err, ErrMissingEventField)
		assert.True(t, strings.Contains(err.Error(), `"id"`))
	})

	t.Run("optional field absent", func(t *testing.T) {
		var event EventResponse
		require.NoError(t, json.Unmarshal([]byte(`{
			"type": "payment_refunded",
			"data": {"id": "pay_optional"}
		}`), &event))

		payload, err := event.DecodePayment()
		require.NoError(t, err)
		assert.Equal(t, "pay_optional", payload.Id)
		assert.Zero(t, payload.Amount)
		assert.Empty(t, payload.Reference)
		assert.Nil(t, payload.Customer)
		assert.Nil(t, payload.Source)
	})
}

func TestDecodePaymentMalformedData(t *testing.T) {
	t.Run("data is an array", func(t *testing.T) {
		var event EventResponse
		err := json.Unmarshal([]byte(`{"type":"payment_approved","data":[1,2]}`), &event)
		assert.Error(t, err)
		assert.True(t, strings.Contains(err.Error(), "event data"))
	})

	t.Run("data is a string", func(t *testing.T) {
		var event EventResponse
		err := json.Unmarshal([]byte(`{"type":"payment_approved","data":"pay_123"}`), &event)
		assert.Error(t, err)
	})

	t.Run("amount has the wrong type", func(t *testing.T) {
		var event EventResponse
		require.NoError(t, json.Unmarshal([]byte(`{
			"type": "payment_voided",
			"data": {"id": "pay_123", "amount": "1500"}
		}`), &event))
		assert.Equal(t, "1500", event.Data["amount"])

		payload, err := event.DecodePayment()
		assert.Nil(t, payload)
		var typeErr *json.UnmarshalTypeError
		assert.ErrorAs(t, err, &typeErr)
		assert.Equal(t, "int64", typeErr.Type.String())
	})

	t.Run("body is not json", func(t *testing.T) {
		var event EventResponse
		err := json.Unmarshal([]byte(`{"data":`), &event)
		assert.Error(t, err)
	})
}

func TestDecodeDataFromConstructedMap(t *testing.T) {
	event := EventResponse{
		Type: EventTypePaymentPending,
		Data: map[string]interface{}{
			"id":       "pay_memory",
			"amount":   float64(80),
			"currency": "EUR",
		},
	}
	assert.Nil(t, event.RawData())

	payload, err := event.DecodePayment()
	require.NoError(t, err)
	assert.Equal(t, "pay_memory", payload.Id)
	assert.Equal(t, int64(80), payload.Amount)
	assert.Equal(t, common.Currency("EUR"), payload.Currency)
}

func TestDecodeDataRejectsNilDestination(t *testing.T) {
	event := EventResponse{
		Type: EventTypePaymentExpired,
		Data: map[string]interface{}{"id": "pay_1"},
	}

	err := event.DecodeData(nil)
	assert.ErrorIs(t, err, errNilDecodeDestination)

	var payload *PaymentEventData
	err = event.DecodeData(payload)
	assert.ErrorIs(t, err, errNilDecodeDestination)
}

func TestDecodePaymentNilReceiver(t *testing.T) {
	var event *EventResponse
	payload, err := event.DecodePayment()
	assert.Nil(t, payload)
	assert.ErrorIs(t, err, ErrMissingEventData)
	assert.Nil(t, event.RawData())
}

func TestPaymentEventTypesSharePayload(t *testing.T) {
	types := []string{
		EventTypePaymentApproved,
		EventTypePaymentCaptured,
		EventTypePaymentDeclined,
		EventTypePaymentRefunded,
		EventTypePaymentVoided,
		EventTypePaymentPending,
		EventTypePaymentExpired,
		EventTypePaymentCanceled,
	}
	for _, eventType := range types {
		t.Run(eventType, func(t *testing.T) {
			body := []byte(`{"type":"` + eventType + `","data":{"id":"pay_1","amount":1}}`)
			var event EventResponse
			require.NoError(t, json.Unmarshal(body, &event))
			payload, err := event.DecodePayment()
			require.NoError(t, err)
			assert.Equal(t, "pay_1", payload.Id)
		})
	}
}

func TestDecodeDataDoesNotReportMissingWhenMapIsSet(t *testing.T) {
	event := EventResponse{Data: map[string]interface{}{"id": "pay_map"}}
	var dest struct {
		Id string `json:"id"`
	}
	require.NoError(t, event.DecodeData(&dest))
	assert.Equal(t, "pay_map", dest.Id)
	assert.False(t, errors.Is(event.DecodeData(&dest), ErrMissingEventData))
}
