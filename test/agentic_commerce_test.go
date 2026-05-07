package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/agenticcommerce"
	"github.com/checkout/checkout-sdk-go/v2/common"
)

// tests

func TestCreateDelegatedPaymentToken(t *testing.T) {
	t.Skip("Requires a valid HMAC signing key and merchant enabled for agentic commerce")
	cases := []struct {
		name    string
		request agenticcommerce.CreateDelegatedPaymentTokenRequest
		checker func(*agenticcommerce.CreateDelegatedPaymentTokenResponse, error)
	}{
		{
			name:    "when request is valid then create delegated payment token",
			request: buildCreateDelegatedPaymentTokenIntegrationRequest(),
			checker: func(response *agenticcommerce.CreateDelegatedPaymentTokenResponse, err error) {
				assert.Nil(t, err)
				assertDelegatedPaymentTokenResponse(t, response)
			},
		},
	}

	client := DefaultApi().AgenticCommerce

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateDelegatedPaymentToken(tc.request, nil))
		})
	}
}

// common methods

func buildCreateDelegatedPaymentTokenIntegrationRequest() agenticcommerce.CreateDelegatedPaymentTokenRequest {
	expiresAt := time.Now().Add(time.Hour)
	card := agenticcommerce.NewDelegatedPaymentMethodCard()
	card.CardNumberType = agenticcommerce.Fpan
	card.Number = "4242424242424242"
	card.ExpMonth = "11"
	card.ExpYear = "2026"
	card.Metadata = map[string]string{"issuing_bank": "test"}

	return agenticcommerce.CreateDelegatedPaymentTokenRequest{
		PaymentMethod: *card,
		Allowance: agenticcommerce.DelegatedPaymentAllowance{
			Reason:            agenticcommerce.OneTime,
			MaxAmount:         10000,
			Currency:          common.USD,
			MerchantId:        "cli_vkuhvk4vjn2edkps7dfsq6emqm",
			CheckoutSessionId: "1PQrsT",
			ExpiresAt:         &expiresAt,
		},
		RiskSignals: []agenticcommerce.DelegatedPaymentRiskSignal{
			{Type: "card_testing", Score: 10, Action: "blocked"},
		},
		Metadata: map[string]string{"campaign": "q4"},
		// Signature must be computed as: Base64(HMAC-SHA256(signingKey, Timestamp + RequestBody))
		Headers: agenticcommerce.DelegatedPaymentHeaders{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Signature: "computed-hmac-sha256-base64-signature",
		},
	}
}

func assertDelegatedPaymentTokenResponse(t *testing.T, response *agenticcommerce.CreateDelegatedPaymentTokenResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotNil(t, response.Created)
	assert.NotNil(t, response.Metadata)
}
