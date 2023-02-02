package test

import (
	"github.com/checkout/checkout-sdk-go/payments/nas"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
)

func TestVoidCardPayment(t *testing.T) {
	paymentResponse := makeCardPayment(t, false, 10)

	voidRequest := payments.VoidRequest{
		Reference: uuid.New().String(),
	}

	cases := []struct {
		name        string
		paymentId   string
		voidRequest payments.VoidRequest
		checkerOne  func(interface{}, error)
		checkerTwo  func(interface{}, error)
	}{
		{
			name:        "when request valid then return a void response",
			paymentId:   paymentResponse.Id,
			voidRequest: voidRequest,
			checkerOne: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.(*payments.VoidResponse).Reference)
				assert.NotEmpty(t, response.(*payments.VoidResponse).ActionId)
				assert.NotEmpty(t, response.(*payments.VoidResponse).Links)
				assert.NotEmpty(t, response.(*payments.VoidResponse).Links["payment"])
			},
			checkerTwo: func(response interface{}, err error) {
				assert.NotEmpty(t, response.(*nas.GetPaymentResponse).Balances)
				assert.Equal(t, int64(10), response.(*nas.GetPaymentResponse).Balances.TotalAuthorized)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.TotalCaptured)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.TotalRefunded)
				assert.Equal(t, int64(10), response.(*nas.GetPaymentResponse).Balances.TotalVoided)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.AvailableToCapture)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.AvailableToRefund)
				assert.Equal(t, int64(0), response.(*nas.GetPaymentResponse).Balances.AvailableToVoid)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			processOne := func() (interface{}, error) {
				return client.VoidPayment(tc.paymentId, &tc.voidRequest, nil)
			}
			predicateOne := func(data interface{}) bool {
				response := data.(*payments.VoidResponse)
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

func TestVoidCardPaymentIdempotently(t *testing.T) {
	paymentResponse := makeCardPayment(t, false, 10)

	idempotencyKey := uuid.New().String()

	cases := []struct {
		name           string
		paymentId      string
		idempotencyKey string
		checker        func(*payments.VoidResponse, error)
	}{
		{
			name:           "when request valid with idempotencyKey then return a void response",
			paymentId:      paymentResponse.Id,
			idempotencyKey: idempotencyKey,
			checker: func(response *payments.VoidResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)

				response2, err := DefaultApi().Payments.VoidPayment(paymentResponse.Id, nil, &idempotencyKey)
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
			tc.checker(client.VoidPayment(tc.paymentId, nil, &tc.idempotencyKey))
		})
	}
}
