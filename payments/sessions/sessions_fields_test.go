package payment_sessions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

// Verifies the fields recently aligned with the Checkout.com swagger spec
// (CreatePaymentSessionsBaseRequest) serialize to their snake_case wire names.
//   - authorization_type
//   - payment_plan
func TestPaymentSessionsRequest_AuthorizationTypeAndPaymentPlan(t *testing.T) {
	request := PaymentSessionsRequest{
		AuthorizationType: nas.EstimatedAuthorizationType,
		PaymentPlan: &payments.PaymentPlan{
			AmountVariability:   payments.FixedAVT,
			DaysBetweenPayments: 28,
		},
	}

	marshalled, err := json.Marshal(request)

	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"authorization_type":"Estimated"`)
	assert.Contains(t, string(marshalled), `"payment_plan":`)
	assert.Contains(t, string(marshalled), `"days_between_payments":28`)
}

// Verifies the new fields are omitted when unset.
func TestPaymentSessionsRequest_OmitsNewFieldsWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(PaymentSessionsRequest{})

	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "authorization_type")
	assert.NotContains(t, string(marshalled), "payment_plan")
}

func TestPaymentSessionsRequest_CustomerRetryOnlySendsMaxAttempts(t *testing.T) {
	maxAttempts := 2
	request := PaymentSessionsRequest{CustomerRetry: &payments.CustomerRetry{MaxAttempts: &maxAttempts}}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)

	var raw map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &raw))
	assert.Equal(t, map[string]interface{}{"max_attempts": float64(2)}, raw["customer_retry"])
}

// Roundtrip of the documented scalar properties using the swagger field examples.
func TestPaymentSessionsRequest_RoundtripsSwaggerExampleValues(t *testing.T) {
	maxAttempts := 2
	capture := true
	request := PaymentSessionsRequest{
		Amount:              1000,
		Currency:            common.USD,
		PaymentType:         payments.Regular,
		Reference:           "ORD-123A",
		Description:         "Payment for gold necklace",
		ProcessingChannelId: "pc_q4dbxom5jbgudnjzjpz7j2z6uq",
		SuccessUrl:          "https://example.com/payments/success",
		FailureUrl:          "https://example.com/payments/failure",
		Locale:              payments.EnGBLT,
		Capture:             &capture,
		CustomerRetry:       &payments.CustomerRetry{MaxAttempts: &maxAttempts},
		AuthorizationType:   nas.EstimatedAuthorizationType,
		IpAddress:           "90.197.169.245",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"customer_retry":{"max_attempts":2}`)
	assert.Contains(t, string(marshalled), `"ip_address":"90.197.169.245"`)

	var roundtrip PaymentSessionsRequest
	assert.NoError(t, json.Unmarshal(marshalled, &roundtrip))
	assert.Equal(t, request, roundtrip)
}

func TestPaymentSessionsRequest_CustomerRetrySendsZeroMaxAttempts(t *testing.T) {
	zero := 0
	marshalled, err := json.Marshal(PaymentSessionsRequest{CustomerRetry: &payments.CustomerRetry{MaxAttempts: &zero}})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"customer_retry":{"max_attempts":0}`)
}
