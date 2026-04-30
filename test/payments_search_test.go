package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/payments"
	"github.com/checkout/checkout-sdk-go/v2/payments/nas"
)

// tests

func TestSearchPayments(t *testing.T) {
	t.Skip("Avoid because can create timeout in the pipeline, activate when needed")

	authorizedPayment := makeCardPayment(t, false, 10)
	capturedPayment := makeCardPayment(t, true, 10)

	cases := []struct {
		name    string
		request nas.SearchPaymentsRequest
		checker func(interface{}, error)
	}{
		{
			name:    "when searching by id returns authorized payment",
			request: buildSearchPaymentsRequest(authorizedPayment.Id),
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assertSearchPaymentsResponse(t, response.(*nas.SearchPaymentsResponse), authorizedPayment.Id)
			},
		},
		{
			name:    "when searching by id returns captured payment",
			request: buildSearchPaymentsRequest(capturedPayment.Id),
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assertSearchPaymentsResponse(t, response.(*nas.SearchPaymentsResponse), capturedPayment.Id)
			},
		},
		{
			name:    "when searching by reference returns payments with matching reference",
			request: buildSearchPaymentsByQuery(fmt.Sprintf("reference:'%s'", Reference)),
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				resp := response.(*nas.SearchPaymentsResponse)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data)
				for _, payment := range resp.Data {
					assert.NotEmpty(t, payment.Id)
					assert.Equal(t, Reference, payment.Reference)
					assertSearchPaymentSource(t, payment)
				}
			},
		},
		{
			name:    "when searching by status Authorized returns only authorized payments",
			request: buildSearchPaymentsByQuery("status:'Authorized'"),
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				resp := response.(*nas.SearchPaymentsResponse)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data)
				for _, payment := range resp.Data {
					assert.Equal(t, payments.Authorized, payment.Status)
					assertSearchPaymentSource(t, payment)
				}
			},
		},
		{
			name:    "when searching by status Captured returns only captured payments",
			request: buildSearchPaymentsByQuery("status:'Captured'"),
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				resp := response.(*nas.SearchPaymentsResponse)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data)
				for _, payment := range resp.Data {
					assert.Equal(t, payments.Captured, payment.Status)
					assertSearchPaymentSource(t, payment)
				}
			},
		},
		{
			name:    "when searching by currency GBP returns only GBP payments",
			request: buildSearchPaymentsByQuery("currency:'GBP'"),
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				resp := response.(*nas.SearchPaymentsResponse)
				assert.NotNil(t, resp)
				assert.NotEmpty(t, resp.Data)
				for _, payment := range resp.Data {
					assert.Equal(t, common.GBP, payment.Currency)
					assertSearchPaymentSource(t, payment)
				}
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
				return len(response.Data) > 0
			}

			tc.checker(retriable(process, predicate, 2))
		})
	}
}

// common methods

func buildSearchPaymentsRequest(paymentId string) nas.SearchPaymentsRequest {
	return buildSearchPaymentsByQuery(fmt.Sprintf("id:'%s'", paymentId))
}

func buildSearchPaymentsByQuery(query string) nas.SearchPaymentsRequest {
	from := time.Now().Add(-5 * time.Minute)
	to := time.Now().Add(5 * time.Minute)
	return nas.SearchPaymentsRequest{
		Query: query,
		Limit: 10,
		From:  &from,
		To:    &to,
	}
}

func assertSearchPaymentsResponse(t *testing.T, response *nas.SearchPaymentsResponse, expectedPaymentId string) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Data)
	assert.Equal(t, expectedPaymentId, response.Data[0].Id)
	assertSearchPaymentSource(t, response.Data[0])
}

func assertSearchPaymentSource(t *testing.T, payment nas.GetPaymentResponse) {
	assert.NotNil(t, payment.Source)
	if payment.Source.ResponseCardSource != nil {
		assert.NotZero(t, payment.Source.ResponseCardSource.ExpiryMonth)
		assert.NotZero(t, payment.Source.ResponseCardSource.ExpiryYear)
	}
}
