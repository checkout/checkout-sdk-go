package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/paymentmethods"
)

// tests

func TestGetAvailablePaymentMethods(t *testing.T) {
	t.Skip("use on demand")
	cases := []struct {
		name    string
		query   paymentmethods.GetPaymentMethodsQuery
		checker func(*paymentmethods.GetPaymentMethodsResponse, error)
	}{
		{
			name:  "when processing channel id is valid then return payment methods",
			query: paymentmethods.GetPaymentMethodsQuery{ProcessingChannelId: "pc_5jp2az55l3cuths25t5p3xhwru"},
			checker: func(response *paymentmethods.GetPaymentMethodsResponse, err error) {
				assert.Nil(t, err)
				assertPaymentMethodsResponse(t, response)
			},
		},
		{
			name:  "when processing channel id is invalid then return error",
			query: paymentmethods.GetPaymentMethodsQuery{ProcessingChannelId: "pc_invalid"},
			checker: func(response *paymentmethods.GetPaymentMethodsResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
			},
		},
	}

	client := OAuthApi().PaymentMethods

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetAvailablePaymentMethods(tc.query))
		})
	}
}

// common methods

func assertPaymentMethodsResponse(t *testing.T, response *paymentmethods.GetPaymentMethodsResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Methods)
	for _, method := range response.Methods {
		assert.NotEmpty(t, method.Type)
	}
}
