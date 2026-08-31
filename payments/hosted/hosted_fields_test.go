package hosted

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

// Verifies the fields recently aligned with the Checkout.com swagger spec
// (HostedPaymentsRequest) serialize to their snake_case wire names.
//   - authorization_type
//   - payment_plan
func TestPaymentHostedRequest_AuthorizationTypeAndPaymentPlan(t *testing.T) {
	request := PaymentHostedRequest{
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
func TestPaymentHostedRequest_OmitsNewFieldsWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(PaymentHostedRequest{})

	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "authorization_type")
	assert.NotContains(t, string(marshalled), "payment_plan")
}
