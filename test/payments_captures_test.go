package test

import (
	"github.com/checkout/checkout-sdk-go-beta/payments/nas"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCaptureCardPayment(t *testing.T) {

	paymentResponse := makeCardPayment(t, false, 10)

	Wait(time.Duration(3))

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := nas.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
		Amount:    5,
	}

	response, err := DefaultApi().Payments.CapturePayment(paymentResponse.Id, captureRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Reference)
	assert.NotEmpty(t, response.ActionId)
	assert.NotEmpty(t, response.Links)
	assert.NotEmpty(t, response.Links["payment"])

	Wait(time.Duration(3))

	payment, err := DefaultApi().Payments.GetPaymentDetails(paymentResponse.Id)

	assert.NotEmpty(t, payment.Balances)
	assert.Equal(t, 10, payment.Balances.TotalAuthorized)
	assert.Equal(t, 5, payment.Balances.TotalCaptured)
	assert.Equal(t, 0, payment.Balances.TotalRefunded)
	assert.Equal(t, 0, payment.Balances.TotalVoided)
	assert.Equal(t, 0, payment.Balances.AvailableToCapture)
	assert.Equal(t, 5, payment.Balances.AvailableToRefund)
	assert.Equal(t, 0, payment.Balances.AvailableToVoid)

}

func TestCaptureCardPaymentIdempotently(t *testing.T) {

	paymentResponse := makeCardPayment(t, false, 10)

	Wait(time.Duration(3))

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := nas.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
	}

	idempotencyKey := uuid.New().String()

	response1, err := DefaultApi().Payments.CapturePayment(paymentResponse.Id, captureRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response1)

	response2, err := DefaultApi().Payments.CapturePayment(paymentResponse.Id, captureRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response2)

	assert.Equal(t, response1.ActionId, response2.ActionId)

}
