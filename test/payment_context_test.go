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
	paymentContextPayPalSource = sources.NewPaymentContextsPayPalSource()

	paymentContextKlarnaSource = getKlarnaPaymentContextsSource()

	customer = contexts.PaymentContextCustomerRequest{
		Email:         Email,
		EmailVerified: true,
		Name:          Name,
	}

	item = contexts.PaymentContextsItems{
		Name:        "mask",
		Quantity:    1,
		UnitPrice:   1000,
		TotalAmount: 1000,
	}

	paymentContextsPayPalRequest = contexts.PaymentContextsRequest{
		Source:              paymentContextPayPalSource,
		Amount:              1000,
		Currency:            common.EUR,
		PaymentType:         payments.Regular,
		Customer:            &customer,
		Capture:             true,
		ProcessingChannelId: os.Getenv("CHECKOUT_PROCESSING_CHANNEL_ID"),
		SuccessUrl:          "https://example.com/payments/success",
		FailureUrl:          "https://example.com/payments/fail",
		Items: []contexts.PaymentContextsItems{
			item,
		},
	}

	processing = contexts.PaymentContextsProcessing{
		Locale: "en-GB",
	}

	paymentContextsKlarnaRequest = contexts.PaymentContextsRequest{
		Source:              paymentContextKlarnaSource,
		Amount:              1000,
		Currency:            common.EUR,
		PaymentType:         payments.Regular,
		ProcessingChannelId: os.Getenv("CHECKOUT_PROCESSING_CHANNEL_ID"),
		Items: []contexts.PaymentContextsItems{
			item,
		},
		Processing: &processing,
	}
)

func TestRequestPaymentContextPayPal(t *testing.T) {
	t.Skip("unavailable")
	cases := []struct {
		name    string
		request contexts.PaymentContextsRequest
		checker func(response *contexts.PaymentContextsRequestResponse, err error)
	}{
		{
			name:    "when payment context is valid the return a response",
			request: paymentContextsPayPalRequest,
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

func TestRequestPaymentContextKlarna(t *testing.T) {
	cases := []struct {
		name    string
		request contexts.PaymentContextsRequest
		checker func(response *contexts.PaymentContextsRequestResponse, err error)
	}{
		{
			name:    "test Klarna source for request payment contexts",
			request: paymentContextsKlarnaRequest,
			checker: func(response *contexts.PaymentContextsRequestResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.PartnerMetadata.ClientToken)
				assert.NotNil(t, response.PartnerMetadata.SessionId)
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

func getKlarnaPaymentContextsSource() payments.PaymentSource {
	source := sources.NewPaymentContextsKlarnaSource()
	source.AccountHolder = &common.AccountHolder{
		BillingAddress: &common.Address{
			Country: common.DE,
		},
	}

	return source
}

func makePaymentContextRequest(t *testing.T) *contexts.PaymentContextsRequestResponse {

	response, err := DefaultApi().Contexts.RequestPaymentContexts(paymentContextsPayPalRequest)

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
				assert.Equal(t, int64(1000), response.PaymentRequest.Amount)
				assert.Equal(t, common.EUR, response.PaymentRequest.Currency)
				assert.Equal(t, payments.Regular, response.PaymentRequest.PaymentType)
				assert.Equal(t, true, response.PaymentRequest.Capture)
				assert.Equal(t, "mask", response.PaymentRequest.Items[0].Name)
				assert.Equal(t, 1, response.PaymentRequest.Items[0].Quantity)
				assert.Equal(t, 1000, response.PaymentRequest.Items[0].UnitPrice)
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
