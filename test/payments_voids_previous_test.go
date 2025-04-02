package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
)

func TestVoidCardPaymentPrevious(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")
	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	voidRequest := payments.VoidRequest{
		Reference: uuid.New().String(),
	}

	cases := []struct {
		name        string
		paymentId   string
		voidRequest payments.VoidRequest
		checker     func(interface{}, error)
	}{
		{
			name:        "when request valid then return a void response",
			paymentId:   paymentResponse.Id,
			voidRequest: voidRequest,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.(*payments.VoidResponse).Reference)
				assert.NotEmpty(t, response.(*payments.VoidResponse).ActionId)
				assert.NotEmpty(t, response.(*payments.VoidResponse).Links)
				assert.NotEmpty(t, response.(*payments.VoidResponse).Links["payment"])
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			process := func() (interface{}, error) {
				return client.VoidPayment(tc.paymentId, &tc.voidRequest, nil)
			}
			predicate := func(data interface{}) bool {
				response := data.(*payments.VoidResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			tc.checker(retriable(process, predicate, 2))
		})
	}
}

func TestVoidCardPaymentIdempotentlyPrevious(t *testing.T) {
	t.Skip("unavailable")
	paymentResponse := makeCardPaymentPrevious(t, false, 10)

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

				response2, err := PreviousApi().Payments.VoidPayment(paymentResponse.Id, nil, &idempotencyKey)
				assert.Nil(t, err)
				assert.NotNil(t, response2)

				assert.Equal(t, response.ActionId, response2.ActionId)
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checker(client.VoidPayment(tc.paymentId, nil, &tc.idempotencyKey))
		})
	}
}
