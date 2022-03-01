package payments

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestCapturePayment(t *testing.T) {

	pk := os.Getenv("CHECKOUT_FOUR_PUBLIC_KEY")
	sk := os.Getenv("CHECKOUT_FOUR_SECRET_KEY")

	config, err := checkout.SdkConfig(&sk, &pk, checkout.Sandbox)
	if err != nil {
		require.Fail(t, "failed creating the creating the configuration!")
	}
	var paymentsClient = NewClient(*config)

	paymentResponse := requestCardPayment(t, paymentsClient)
	assert.NotNil(t, paymentResponse)

	var captureRequest = CapturesRequest{
		Amount:      5,
		Reference:   "Capture Reference",
		CaptureType: Final,
	}
	captureResponse, err := paymentsClient.Captures(paymentResponse.Processed.ID, &captureRequest, nil)

	assert.Nil(t, err)
	assert.NotNil(t, captureResponse)
	assert.NotNil(t, captureResponse.Accepted)
	assert.NotNil(t, captureResponse.Accepted.ActionID)
	assert.NotNil(t, captureResponse.Accepted.Reference)
	assert.NotNil(t, captureResponse.Accepted.Links)

}

func requestCardPayment(t *testing.T, client *Client) *Response {

	var source = CardSource{
		Type:        common.Card.String(),
		Number:      "4556447238607884",
		ExpiryMonth: 6,
		ExpiryYear:  2025,
		CVV:         "100",
		Stored:      checkout.Bool(false),
	}
	var request = &Request{
		Capture:   checkout.Bool(false),
		Reference: "Payment Reference",
		Amount:    10,
		Currency:  "EUR",
		Source:    source,
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
