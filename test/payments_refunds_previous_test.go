package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
)

func TestRefundCardPaymentPrevious(t *testing.T) {
	t.Skip("unavailable")

	paymentResponse := makeCardPaymentPrevious(t, true, 10)

	refundRequest := payments.RefundRequest{
		Reference: uuid.New().String(),
	}

	cases := []struct {
		name          string
		paymentId     string
		refundRequest payments.RefundRequest
		checker       func(interface{}, error)
	}{
		{
			name:          "when request is valid then return a refund response",
			paymentId:     paymentResponse.Id,
			refundRequest: refundRequest,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.(*payments.RefundResponse).Reference)
				assert.NotEmpty(t, response.(*payments.RefundResponse).ActionId)
				assert.NotEmpty(t, response.(*payments.RefundResponse).Links)
				assert.NotEmpty(t, response.(*payments.RefundResponse).Links["payment"])
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(2)
			process := func() (interface{}, error) {
				return client.RefundPayment(tc.paymentId, &tc.refundRequest, nil)
			}
			predicate := func(data interface{}) bool {
				response := data.(*payments.RefundResponse)
				return response.Links != nil && len(response.Links) >= 0
			}
			Wait(2)
			tc.checker(retriable(process, predicate, 2))
		})
	}
}

func TestRefundCardPaymentIdempotentlyPrevious(t *testing.T) {
	t.Skip("unavailable")

	paymentResponse := makeCardPaymentPrevious(t, true, 10)

	refundRequest := payments.RefundRequest{
		Amount:    5,
		Reference: uuid.New().String(),
		Metadata:  nil,
	}

	idempotencyKeyRandom1 := uuid.New().String()

	idempotencyKeyRandom2 := uuid.New().String()

	cases := []struct {
		name                  string
		paymentId             string
		refundRequest         payments.RefundRequest
		idempotencyKeyRandom1 string
		idempotencyKeyRandom2 string
		checker               func(interface{}, error, interface{}, error)
	}{
		{
			name:                  "when get a refund payment with idempotencyKey request then return a response",
			paymentId:             paymentResponse.Id,
			refundRequest:         refundRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom1,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.Equal(t, response1.(*payments.RefundResponse).ActionId, response2.(*payments.RefundResponse).ActionId)
			},
		},
		{
			name:                  "when request is valid then capture payment idempotently error",
			paymentId:             paymentResponse.Id,
			refundRequest:         refundRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom2,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.NotEqual(t, response1.(*payments.RefundResponse).ActionId, response2.(*payments.RefundResponse).ActionId)
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			processOne := func() (interface{}, error) {
				return client.RefundPayment(tc.paymentId, &tc.refundRequest, &tc.idempotencyKeyRandom1)
			}
			predicateOne := func(data interface{}) bool {
				response := data.(*payments.RefundResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			processTwo := func() (interface{}, error) {
				return client.RefundPayment(tc.paymentId, &tc.refundRequest, &tc.idempotencyKeyRandom2)
			}
			predicateTwo := func(data interface{}) bool {
				response := data.(*payments.RefundResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			Wait(2)
			retriableOne, errOne := retriable(processOne, predicateOne, 2)
			Wait(2)
			retriableTwo, errTwo := retriable(processTwo, predicateTwo, 2)
			Wait(2)
			tc.checker(retriableOne, errOne, retriableTwo, errTwo)
		})
	}
}
