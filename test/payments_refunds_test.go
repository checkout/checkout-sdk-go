package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
)

func TestRefundCardPayment(t *testing.T) {

	paymentResponse := makeCardPayment(t, true, 10)

	Wait(time.Duration(3))

	refundRequest := payments.RefundRequest{
		Reference: uuid.New().String(),
	}
	response, err := DefaultApi().Payments.RefundPayment(paymentResponse.Id, &refundRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Reference)
	assert.NotEmpty(t, response.ActionId)
	assert.NotEmpty(t, response.Links)
	assert.NotEmpty(t, response.Links["payment"])

	payment, err := DefaultApi().Payments.GetPaymentDetails(paymentResponse.Id)

	assert.NotEmpty(t, payment.Balances)
	assert.Equal(t, int64(10), payment.Balances.TotalAuthorized)
	assert.Equal(t, int64(10), payment.Balances.TotalCaptured)
	assert.Equal(t, int64(10), payment.Balances.TotalRefunded)
	assert.Equal(t, int64(0), payment.Balances.TotalVoided)
	assert.Equal(t, int64(0), payment.Balances.AvailableToCapture)
	assert.Equal(t, int64(0), payment.Balances.AvailableToRefund)
	assert.Equal(t, int64(0), payment.Balances.AvailableToVoid)

}

func TestRefundCardPaymentIdempotently(t *testing.T) {

	paymentResponse := makeCardPayment(t, true, 10)

	Wait(time.Duration(3))

	refundRequest := payments.RefundRequest{
		Amount:    5,
		Reference: uuid.New().String(),
		Metadata:  nil,
	}

	idempotencyKey := uuid.New().String()

	response1, err := DefaultApi().Payments.RefundPayment(paymentResponse.Id, &refundRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response1)

	response2, err := DefaultApi().Payments.RefundPayment(paymentResponse.Id, &refundRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response2)

	assert.Equal(t, response1.ActionId, response2.ActionId)

}
