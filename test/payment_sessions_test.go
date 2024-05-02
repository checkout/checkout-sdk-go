package test

import (
	"github.com/checkout/checkout-sdk-go/payments"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments/sessions"
)

var (
	paymentSessionsRequest = payment_sessions.PaymentSessionsRequest{
		Amount:     int64(2000),
		Currency:   common.GBP,
		Reference:  "ORD-123A",
		Billing:    &payments.BillingInformation{Address: Address()},
		SuccessUrl: "https://example.com/payments/success",
		FailureUrl: "https://example.com/payments/fail",
	}
)

func TestRequestPaymentSessions(t *testing.T) {
	cases := []struct {
		name    string
		request payment_sessions.PaymentSessionsRequest
		checker func(response *payment_sessions.PaymentSessionsResponse, err error)
	}{
		{
			name:    "when payment context is valid the return a response",
			request: paymentSessionsRequest,
			checker: func(response *payment_sessions.PaymentSessionsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				if response.Links != nil {
					assert.NotNil(t, response.Links)
				}
			},
		},
	}

	client := DefaultApi().PaymentSessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestPaymentSessions(tc.request))
		})
	}

}
