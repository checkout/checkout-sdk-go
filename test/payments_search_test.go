package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/payments/nas"
)

// tests

func TestSearchPayments(t *testing.T) {
	t.Skip("Avoid because can create timeout in the pipeline, activate when needed")

	paymentResponse := makeCardPayment(t, false, 10)

	cases := []struct {
		name    string
		request nas.SearchPaymentsRequest
		checker func(interface{}, error)
	}{
		{
			name:    "when search is valid then return payments",
			request: buildSearchPaymentsRequest(paymentResponse.Id),
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assertSearchPaymentsResponse(t, response.(*nas.SearchPaymentsResponse), paymentResponse.Id)
			},
		},
	}

	client := OAuthApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			process := func() (interface{}, error) {
				return client.SearchPayments(tc.request)
			}
			predicate := func(data interface{}) bool {
				response := data.(*nas.SearchPaymentsResponse)
				return response.Data != nil && len(response.Data) > 0
			}

			tc.checker(retriable(process, predicate, 2))
		})
	}
}

// common methods

func buildSearchPaymentsRequest(paymentId string) nas.SearchPaymentsRequest {
	from := time.Now().Add(-5 * time.Minute)
	to := time.Now().Add(5 * time.Minute)
	return nas.SearchPaymentsRequest{
		Query: fmt.Sprintf("id:'%s'", paymentId),
		Limit: 10,
		From:  &from,
		To:    &to,
	}
}

func assertSearchPaymentsResponse(t *testing.T, response *nas.SearchPaymentsResponse, expectedPaymentId string) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Data)
	assert.Equal(t, expectedPaymentId, response.Data[0].Id)
}
