package nas

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/payments"
)

// Verifies PaymentInstructionResponse aligned with the 2026-08-05 Checkout.com
// swagger delta: funds_transfer_type on the card payout response instruction.

func TestPaymentInstructionResponse_FundsTransferTypePresentWhenSet(t *testing.T) {
	payload := `{"value_date":"2026-08-05T00:00:00Z","funds_transfer_type":"AA"}`

	var instruction PaymentInstructionResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &instruction))
	assert.Equal(t, "AA", instruction.FundsTransferType)

	marshalled, err := json.Marshal(instruction)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"funds_transfer_type":"AA"`)
}

func TestPaymentInstructionResponse_FundsTransferTypeAbsentWhenUnset(t *testing.T) {
	instruction := PaymentInstructionResponse{}

	marshalled, err := json.Marshal(instruction)
	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "funds_transfer_type")
}

func TestGetPaymentResponse_DeserializesPaymentIp(t *testing.T) {
	var response GetPaymentResponse
	assert.NoError(t, json.Unmarshal([]byte(`{"id":"pay_123","payment_ip":"90.197.169.245"}`), &response))
	assert.Equal(t, "90.197.169.245", response.PaymentIp)
}

func TestPaymentRequest_RetryRoundtripsAndSendsEnabledFalse(t *testing.T) {
	request := PaymentRequest{
		Amount:    6540,
		Reference: "ORD-5023-4E89",
		PaymentIp: "90.197.169.245",
		Retry: &payments.PaymentRetryRequest{
			Enabled:  true,
			Dunning:  &payments.Dunning{Enabled: true, MaxAttempts: 6, EndAfterDays: 30},
			Downtime: &payments.Downtime{Enabled: false},
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"retry":{"dunning":{"enabled":true,"max_attempts":6,"end_after_days":30},"downtime":{"enabled":false},"enabled":true}`)
	assert.Contains(t, string(marshalled), `"payment_ip":"90.197.169.245"`)

	var roundtrip PaymentRequest
	assert.NoError(t, json.Unmarshal(marshalled, &roundtrip))
	assert.Equal(t, request.Retry, roundtrip.Retry)
	assert.Equal(t, request.PaymentIp, roundtrip.PaymentIp)

	marshalled, err = json.Marshal(PaymentRequest{Retry: &payments.PaymentRetryRequest{Enabled: false}})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"retry":{"enabled":false}`)
}

func TestGetPaymentResponse_DeserializesRetry(t *testing.T) {
	payload := `{"id":"pay_mbabizu24mvu3mela5njyhpit4","retry":{"enabled":false,"attempts_made":0,"max_attempts":6}}`

	var response GetPaymentResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.False(t, response.Retry.Enabled)
	assert.Equal(t, 6, response.Retry.MaxAttempts)
	assert.Equal(t, "pay_mbabizu24mvu3mela5njyhpit4", response.Id)
}
