package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/payments"
	"github.com/checkout/checkout-sdk-go/v2/payments/nas"
)

// tests

func TestReverseCardPayment(t *testing.T) {
	paymentResponse := makeCardPayment(t, false, 10)

	cases := []struct {
		name        string
		paymentId   string
		request     payments.PaymentReversalRequest
		checkerOne  func(interface{}, error)
		checkerTwo  func(interface{}, error)
	}{
		{
			name:      "when request is valid then return a reversal response",
			paymentId: paymentResponse.Id,
			request:   buildReversalIntegrationRequest(),
			checkerOne: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.(*payments.PaymentReversalResponse).ActionId)
				assert.NotEmpty(t, response.(*payments.PaymentReversalResponse).Links)
				assert.NotEmpty(t, response.(*payments.PaymentReversalResponse).Links["payment"])
			},
			checkerTwo: func(response interface{}, err error) {
				assert.NotEmpty(t, response.(*nas.GetPaymentResponse).Balances)
				assert.Equal(t, int64(10), response.(*nas.GetPaymentResponse).Balances.TotalAuthorized)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.TotalCaptured)
				assert.Equal(t, int64(10), response.(*nas.GetPaymentResponse).Balances.TotalVoided)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.AvailableToCapture)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.AvailableToVoid)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			processOne := func() (interface{}, error) {
				return client.ReversePayment(tc.paymentId, &tc.request, nil)
			}
			predicateOne := func(data interface{}) bool {
				response := data.(*payments.PaymentReversalResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			tc.checkerOne(retriable(processOne, predicateOne, 2))

			Wait(3)

			processTwo := func() (interface{}, error) {
				return client.GetPaymentDetails(tc.paymentId)
			}
			predicateTwo := func(data interface{}) bool {
				response := data.(*nas.GetPaymentResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			tc.checkerTwo(retriable(processTwo, predicateTwo, 2))
		})
	}
}

func TestReverseCardPaymentIdempotently(t *testing.T) {
	paymentResponse := makeCardPayment(t, false, 10)

	idempotencyKey := uuid.New().String()

	cases := []struct {
		name           string
		paymentId      string
		idempotencyKey string
		checker        func(*payments.PaymentReversalResponse, error)
	}{
		{
			name:           "when request is valid with idempotencyKey then return a reversal response",
			paymentId:      paymentResponse.Id,
			idempotencyKey: idempotencyKey,
			checker: func(response *payments.PaymentReversalResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)

				response2, err := DefaultApi().Payments.ReversePayment(paymentResponse.Id, nil, &idempotencyKey)
				assert.Nil(t, err)
				assert.NotNil(t, response2)

				assert.Equal(t, response.ActionId, response2.ActionId)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checker(client.ReversePayment(tc.paymentId, nil, &tc.idempotencyKey))
		})
	}
}

// common methods

func buildReversalIntegrationRequest() payments.PaymentReversalRequest {
	return payments.PaymentReversalRequest{
		Reference: uuid.New().String(),
		Metadata: map[string]interface{}{
			"coupon_code": "NY2018",
			"partner_id":  123989,
		},
	}
}

