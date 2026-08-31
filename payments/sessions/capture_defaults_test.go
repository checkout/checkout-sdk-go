package payment_sessions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

// Verifies that capture is omitted from every payment session request unless the caller sets it.
//
// Capture was previously declared as a non-pointer bool without omitempty, so a zero-value struct
// serialized "capture":false. On session creation that silently disabled auto-capture, contradicting
// the API default of true. On submit it overwrote the value provided when the session was created.

func TestPaymentSessionsRequest_OmitsCaptureWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(PaymentSessionsRequest{})

	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "capture")
}

func TestPaymentSessionsWithPaymentRequest_OmitsCaptureWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(PaymentSessionsWithPaymentRequest{SessionData: "SD"})

	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "capture")
}

func TestSubmitPaymentSessionRequest_OmitsCaptureWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(SubmitPaymentSessionRequest{SessionData: "SD"})

	assert.NoError(t, err)
	assert.Equal(t, `{"session_data":"SD"}`, string(marshalled))
}

// Verifies capture is still fully expressible in both directions.

func TestSubmitPaymentSessionRequest_SendsCaptureFalseWhenExplicitlySet(t *testing.T) {
	capture := false

	marshalled, err := json.Marshal(SubmitPaymentSessionRequest{
		SessionData: "SD",
		Capture:     &capture,
	})

	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"capture":false`)
}

func TestSubmitPaymentSessionRequest_SendsCaptureTrueWhenExplicitlySet(t *testing.T) {
	capture := true

	marshalled, err := json.Marshal(SubmitPaymentSessionRequest{
		SessionData: "SD",
		Capture:     &capture,
	})

	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"capture":true`)
}

func TestPaymentSessionsRequest_SendsCaptureFalseWhenExplicitlySet(t *testing.T) {
	capture := false

	marshalled, err := json.Marshal(PaymentSessionsRequest{
		Amount:   1000,
		Currency: common.GBP,
		Capture:  &capture,
	})

	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"capture":false`)
}

// Verifies the fields added to close the gaps found while auditing these three request bodies
// against the swagger.

func TestSubmitPaymentSessionRequest_SerializesAmountAllocations(t *testing.T) {
	marshalled, err := json.Marshal(SubmitPaymentSessionRequest{
		SessionData: "SD",
		AmountAllocations: []common.AmountAllocations{
			{
				Id:     "ent_abcdefghijklmnopqrstuvwxyz",
				Amount: 1000,
			},
		},
	})

	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"amount_allocations":`)
	assert.Contains(t, string(marshalled), `"id":"ent_abcdefghijklmnopqrstuvwxyz"`)
}

func TestPaymentSessionsWithPaymentRequest_SerializesAuthorizationTypeAndPaymentPlan(t *testing.T) {
	marshalled, err := json.Marshal(PaymentSessionsWithPaymentRequest{
		SessionData:       "SD",
		AuthorizationType: nas.EstimatedAuthorizationType,
		PaymentPlan: &payments.PaymentPlan{
			AmountVariability:   payments.FixedAVT,
			DaysBetweenPayments: 28,
		},
	})

	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"authorization_type":"Estimated"`)
	assert.Contains(t, string(marshalled), `"payment_plan":`)
	assert.Contains(t, string(marshalled), `"days_between_payments":28`)
}

func TestPaymentSessionsWithPaymentRequest_OmitsNewFieldsWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(PaymentSessionsWithPaymentRequest{SessionData: "SD"})

	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "authorization_type")
	assert.NotContains(t, string(marshalled), "payment_plan")
}

func TestSubmitPaymentSessionRequest_UnmarshalsCapture(t *testing.T) {
	var request SubmitPaymentSessionRequest

	err := json.Unmarshal([]byte(`{"session_data":"SD","capture":false}`), &request)

	assert.NoError(t, err)
	assert.Equal(t, "SD", request.SessionData)
	assert.NotNil(t, request.Capture)
	assert.False(t, *request.Capture)
}

func TestSubmitPaymentSessionRequest_UnmarshalsDocumentedExample(t *testing.T) {
	var request SubmitPaymentSessionRequest

	err := json.Unmarshal([]byte(`{"session_data":"{SESSION_DATA_FROM_FLOW}","3ds":{"enabled":true}}`), &request)

	assert.NoError(t, err)
	assert.Equal(t, "{SESSION_DATA_FROM_FLOW}", request.SessionData)
	assert.Nil(t, request.Capture, "the documented example does not carry capture")
}
