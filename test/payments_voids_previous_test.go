package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go-beta/payments"
)

func TestVoidCardPaymentPrevious(t *testing.T) {

	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	Wait(time.Duration(3))

	voidRequest := payments.VoidRequest{
		Reference: uuid.New().String(),
	}

	response, err := PreviousApi().Payments.VoidPayment(paymentResponse.Id, &voidRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Reference)
	assert.NotEmpty(t, response.ActionId)
	assert.NotEmpty(t, response.Links)
	assert.NotEmpty(t, response.Links["payment"])
}

func TestVoidCardPaymentIdempotentlyPrevious(t *testing.T) {

	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	Wait(time.Duration(3))

	idempotencyKey := uuid.New().String()

	response1, err := PreviousApi().Payments.VoidPayment(paymentResponse.Id, nil, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response1)

	response2, err := PreviousApi().Payments.VoidPayment(paymentResponse.Id, nil, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response2)

	assert.Equal(t, response1.ActionId, response2.ActionId)
}
