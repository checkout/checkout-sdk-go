package payments

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func intPtr(v int) *int { return &v }

func TestPaymentRetryRequest_SerializesAllPropertiesAndRoundtrips(t *testing.T) {
	request := PaymentRetryRequest{
		Enabled:  true,
		Dunning:  &Dunning{Enabled: true, MaxAttempts: 5, EndAfterDays: 30},
		Downtime: &Downtime{Enabled: true},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"enabled":true,"dunning":{"enabled":true,"max_attempts":5,"end_after_days":30},"downtime":{"enabled":true}}`, string(marshalled))

	var roundtrip PaymentRetryRequest
	assert.NoError(t, json.Unmarshal(marshalled, &roundtrip))
	assert.Equal(t, request, roundtrip)
}

func TestPaymentRetryRequest_SerializesEnabledFalse(t *testing.T) {
	request := PaymentRetryRequest{
		Enabled:  false,
		Dunning:  &Dunning{Enabled: false},
		Downtime: &Downtime{Enabled: false},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"enabled":false,"dunning":{"enabled":false},"downtime":{"enabled":false}}`, string(marshalled))
}

func TestPaymentRetryRequest_DeserializesDocumentedDefaults(t *testing.T) {
	payload := `{"enabled":true,"dunning":{"enabled":true,"max_attempts":6,"end_after_days":30},"downtime":{"enabled":false}}`

	var request PaymentRetryRequest
	assert.NoError(t, json.Unmarshal([]byte(payload), &request))
	assert.True(t, request.Enabled)
	assert.Equal(t, 6, request.Dunning.MaxAttempts)
	assert.Equal(t, 30, request.Dunning.EndAfterDays)
	assert.False(t, request.Downtime.Enabled)
}

func TestCustomerRetry_SerializesMaxAttemptsAndRoundtrips(t *testing.T) {
	// swagger example: max_attempts 2
	request := CustomerRetry{MaxAttempts: intPtr(2)}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"max_attempts":2}`, string(marshalled))

	var roundtrip CustomerRetry
	assert.NoError(t, json.Unmarshal(marshalled, &roundtrip))
	assert.Equal(t, request, roundtrip)
}

func TestCustomerRetry_SerializesZeroMaxAttempts(t *testing.T) {
	marshalled, err := json.Marshal(CustomerRetry{MaxAttempts: intPtr(0)})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"max_attempts":0}`, string(marshalled))

	marshalled, err = json.Marshal(CustomerRetry{})
	assert.NoError(t, err)
	assert.JSONEq(t, `{}`, string(marshalled))
}

func TestPaymentRetryRequest_SerializesDeprecatedTopLevelFields(t *testing.T) {
	request := PaymentRetryRequest{Enabled: true, MaxAttempts: 3, EndAfterDays: 10}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"enabled":true,"max_attempts":3,"end_after_days":10}`, string(marshalled))
}

func TestPaymentRetryResponse_DeserializesAllPropertiesAndRoundtrips(t *testing.T) {
	payload := `{"enabled":true,"attempts_made":2,"max_attempts":6,"ends_on":"2026-11-25T10:00:00Z","next_attempt_on":"2026-10-27T10:00:00Z"}`

	var response PaymentRetryResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.True(t, response.Enabled)
	assert.Equal(t, 2, response.AttemptsMade)
	assert.Equal(t, 6, response.MaxAttempts)
	assert.Equal(t, time.Date(2026, 11, 25, 10, 0, 0, 0, time.UTC), response.EndsOn.UTC())
	assert.Equal(t, time.Date(2026, 10, 27, 10, 0, 0, 0, time.UTC), response.NextAttemptOn.UTC())

	marshalled, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.JSONEq(t, payload, string(marshalled))
}

func TestPaymentRetryResponse_DeserializesEnabledFalse(t *testing.T) {
	var response PaymentRetryResponse
	assert.NoError(t, json.Unmarshal([]byte(`{"enabled":false}`), &response))
	assert.False(t, response.Enabled)

	marshalled, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"enabled":false}`, string(marshalled))
}
