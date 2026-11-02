package links

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

// Verifies the fields recently aligned with the Checkout.com swagger spec
// (PaymentLinksRequest) serialize to their snake_case wire names.
//   - authorization_type
//   - payment_plan
func TestPaymentLinkRequest_AuthorizationTypeAndPaymentPlan(t *testing.T) {
	request := PaymentLinkRequest{
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
func TestPaymentLinkRequest_OmitsNewFieldsWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(PaymentLinkRequest{})

	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "authorization_type")
	assert.NotContains(t, string(marshalled), "payment_plan")
}

func TestPaymentLinkRequest_CustomerRetryOnlySendsMaxAttempts(t *testing.T) {
	maxAttempts := 2
	request := PaymentLinkRequest{CustomerRetry: &payments.CustomerRetry{MaxAttempts: &maxAttempts}}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)

	var raw map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &raw))
	assert.Equal(t, map[string]interface{}{"max_attempts": float64(2)}, raw["customer_retry"])
}

// Roundtrip of the documented scalar properties using the swagger field examples.
func TestPaymentLinkRequest_RoundtripsSwaggerExampleValues(t *testing.T) {
	maxAttempts := 2
	request := PaymentLinkRequest{
		Amount:              1000,
		Currency:            common.GBP,
		PaymentType:         payments.Regular,
		PaymentIp:           "90.197.169.245",
		Reference:           "ORD-123A",
		Description:         "Payment for Gold Necklace",
		DisplayName:         "The Jewelry Shop",
		ProcessingChannelId: "pc_q4dbxom5jbgudnjzjpz7j2z6uq",
		ExpiresIn:           604800,
		ReturnUrl:           "https://example.com/success",
		Locale:              "en-GB",
		Capture:             true,
		CustomerRetry:       &payments.CustomerRetry{MaxAttempts: &maxAttempts},
		AuthorizationType:   nas.EstimatedAuthorizationType,
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"payment_type":"Regular"`)
	assert.Contains(t, string(marshalled), `"payment_ip":"90.197.169.245"`)
	assert.Contains(t, string(marshalled), `"customer_retry":{"max_attempts":2}`)

	var roundtrip PaymentLinkRequest
	assert.NoError(t, json.Unmarshal(marshalled, &roundtrip))
	assert.Equal(t, request, roundtrip)
}

func TestPaymentLinkRequest_CustomerRetrySendsZeroMaxAttempts(t *testing.T) {
	zero := 0
	marshalled, err := json.Marshal(PaymentLinkRequest{CustomerRetry: &payments.CustomerRetry{MaxAttempts: &zero}})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"customer_retry":{"max_attempts":0}`)
}
