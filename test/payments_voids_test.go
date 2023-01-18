package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
)

func TestVoidCardPayment(t *testing.T) {

	paymentResponse := makeCardPayment(t, false, 10)

	Wait(time.Duration(3))

	voidRequest := payments.VoidRequest{
		Reference: uuid.New().String(),
	}

	response, err := DefaultApi().Payments.VoidPayment(paymentResponse.Id, &voidRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Reference)
	assert.NotEmpty(t, response.ActionId)
	assert.NotEmpty(t, response.Links)
	assert.NotEmpty(t, response.Links["payment"])

	Wait(time.Duration(3))

	payment, err := DefaultApi().Payments.GetPaymentDetails(paymentResponse.Id)

	assert.NotEmpty(t, payment.Balances)
	assert.Equal(t, int64(10), payment.Balances.TotalAuthorized)
	assert.Equal(t, int64(0), payment.Balances.TotalCaptured)
	assert.Equal(t, int64(0), payment.Balances.TotalRefunded)
	assert.Equal(t, int64(10), payment.Balances.TotalVoided)
	assert.Equal(t, int64(0), payment.Balances.AvailableToCapture)
	assert.Equal(t, int64(0), payment.Balances.AvailableToRefund)
	assert.Equal(t, int64(0), payment.Balances.AvailableToVoid)
}

func TestVoidCardPaymentIdempotently(t *testing.T) {

	paymentResponse := makeCardPayment(t, false, 10)

	Wait(time.Duration(3))

	idempotencyKey := uuid.New().String()

	response1, err := DefaultApi().Payments.VoidPayment(paymentResponse.Id, nil, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response1)

	response2, err := DefaultApi().Payments.VoidPayment(paymentResponse.Id, nil, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response2)

	assert.Equal(t, response1.ActionId, response2.ActionId)
}
