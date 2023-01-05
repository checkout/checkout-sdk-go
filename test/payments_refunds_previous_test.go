package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
)

func TestRefundCardPaymentPrevious(t *testing.T) {

	paymentResponse := makeCardPaymentPrevious(t, true, 10)

	Wait(time.Duration(3))

	refundRequest := payments.RefundRequest{
		Reference: uuid.New().String(),
	}
	response, err := PreviousApi().Payments.RefundPayment(paymentResponse.Id, &refundRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Reference)
	assert.NotEmpty(t, response.ActionId)
	assert.NotEmpty(t, response.Links)
	assert.NotEmpty(t, response.Links["payment"])

}

func TestRefundCardPaymentIdempotentlyPrevious(t *testing.T) {

	paymentResponse := makeCardPaymentPrevious(t, true, 10)

	Wait(time.Duration(3))

	refundRequest := payments.RefundRequest{
		Amount:    5,
		Reference: uuid.New().String(),
		Metadata:  nil,
	}

	idempotencyKey := uuid.New().String()

	response1, err := PreviousApi().Payments.RefundPayment(paymentResponse.Id, &refundRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response1)

	response2, err := PreviousApi().Payments.RefundPayment(paymentResponse.Id, &refundRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response2)

	assert.Equal(t, response1.ActionId, response2.ActionId)

}
