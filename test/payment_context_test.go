package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/contexts"
	sources "github.com/checkout/checkout-sdk-go/payments/nas/sources/contexts"
)

var (
	paymentContextSource = sources.NewPaymentContextsPaypalSource()

	item = contexts.PaymentContextsItems{
		Name:      "mask",
		Quantity:  1,
		UnitPrice: 2000,
	}

	paymentContextsRequest = contexts.PaymentContextsRequest{
		Source:              paymentContextSource,
		Amount:              2000,
		Currency:            common.EUR,
		PaymentType:         payments.Regular,
		Capture:             true,
		ProcessingChannelId: os.Getenv("CHECKOUT_PROCESSING_CHANNEL_ID"),
		SuccessUrl:          "https://example.com/payments/success",
		FailureUrl:          "https://example.com/payments/fail",
		Items: []contexts.PaymentContextsItems{
			item,
		},
	}
)

func TestRequestPaymentContext(t *testing.T) {
	cases := []struct {
		name    string
		request contexts.PaymentContextsRequest
		checker func(response *contexts.PaymentContextsRequestResponse, err error)
	}{
		{
			name:    "when payment context is valid the return a response",
			request: paymentContextsRequest,
			checker: func(response *contexts.PaymentContextsRequestResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.PartnerMetadata.OrderId)
				if response.Links != nil {
					assert.NotNil(t, response.Links)
				}
			},
		},
	}

	client := DefaultApi().Contexts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestPaymentContexts(tc.request))
		})
	}

}

func makePaymentContextRequest(t *testing.T) *contexts.PaymentContextsRequestResponse {

	response, err := DefaultApi().Contexts.RequestPaymentContexts(paymentContextsRequest)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.PartnerMetadata.OrderId)
	if response.Links != nil {
		assert.NotNil(t, response.Links)
	}

	return response
}

func TestGetPaymentContext(t *testing.T) {

	request := makePaymentContextRequest(t)

	cases := []struct {
		name             string
		paymentContextId string
		checker          func(*contexts.PaymentContextDetailsResponse, error)
	}{
		{
			name:             "when get a payment context and the response return correct data",
			paymentContextId: request.Id,
			checker: func(response *contexts.PaymentContextDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, int64(2000), response.PaymentRequest.Amount)
				assert.Equal(t, common.EUR, response.PaymentRequest.Currency)
				assert.Equal(t, payments.Regular, response.PaymentRequest.PaymentType)
				assert.Equal(t, true, response.PaymentRequest.Capture)
				assert.Equal(t, "mask", response.PaymentRequest.Items[0].Name)
				assert.Equal(t, 1, response.PaymentRequest.Items[0].Quantity)
				assert.Equal(t, 2000, response.PaymentRequest.Items[0].UnitPrice)
				assert.Equal(t, "https://example.com/payments/success", response.PaymentRequest.SuccessUrl)
				assert.Equal(t, "https://example.com/payments/fail", response.PaymentRequest.FailureUrl)
				assert.NotNil(t, response.PartnerMetadata)
				assert.NotNil(t, response.PartnerMetadata.OrderId)
			},
		},
	}

	client := DefaultApi().Contexts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetPaymentContextDetails(tc.paymentContextId))
		})
	}
}
