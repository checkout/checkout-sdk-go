package payments

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestIncrementAuthorization(t *testing.T) {

	pk := os.Getenv("CHECKOUT_FOUR_PUBLIC_KEY")
	sk := os.Getenv("CHECKOUT_FOUR_SECRET_KEY")

	config, err := checkout.SdkConfig(&sk, &pk, checkout.Sandbox)
	if err != nil {
		require.Fail(t, "failed creating the creating the configuration!")
	}
	var paymentsClient = NewClient(*config)

	paymentResponse := requestPayment(t, paymentsClient)
	assert.NotNil(t, paymentResponse)

	var authorizationRequest = AuthorizationRequest{
		Amount:    10,
		Reference: "Auth Reference",
	}
	authorizationResponse, err := paymentsClient.IncrementAuthorization(paymentResponse.Processed.ID, &authorizationRequest, nil)

	assert.Nil(t, err)
	assert.NotNil(t, authorizationResponse)
	assert.NotNil(t, authorizationResponse.Amount)
	assert.NotNil(t, authorizationResponse.ActionID)
	assert.NotNil(t, authorizationResponse.Currency)
	assert.NotNil(t, authorizationResponse.Approved)
	assert.NotNil(t, authorizationResponse.ResponseCode)
	assert.NotNil(t, authorizationResponse.ResponseSummary)
	assert.NotNil(t, authorizationResponse.ExpiresOn)
	assert.NotNil(t, authorizationResponse.ProcessedOn)
	assert.NotNil(t, authorizationResponse.Balances)
	assert.NotNil(t, authorizationResponse.Links)
	assert.NotNil(t, authorizationResponse.Risk)

}

func requestPayment(t *testing.T, client *Client) *Response {

	var source = CardSource{
		Type:        common.Card.String(),
		Number:      "4556447238607884",
		ExpiryMonth: 6,
		ExpiryYear:  2025,
		CVV:         "100",
		Stored:      checkout.Bool(false),
	}
	var request = &Request{
		Capture:           checkout.Bool(false),
		Reference:         "Payment Reference",
		Amount:            10,
		Currency:          "EUR",
		AuthorizationType: "Estimated",
		Source:            source,
		Customer: &Customer{
			Email: "example@email.com",
			Name:  "First Name Last Name",
		},
		Metadata: map[string]string{
			"test1": "test_metadata_1",
		},
	}
	response, err := client.Request(request, nil)
	if err != nil || response.Processed == nil {
		require.Fail(t, "payment request failed!")
	}
	return response
}
