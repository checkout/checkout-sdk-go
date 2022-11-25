package test

import (
	"github.com/checkout/checkout-sdk-go-beta/payments/abc"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCaptureCardPaymentPrevious(t *testing.T) {

	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	Wait(time.Duration(3))

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := abc.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
		Amount:    5,
	}

	response, err := PreviousApi().Payments.CapturePayment(paymentResponse.Id, captureRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Reference)
	assert.NotEmpty(t, response.ActionId)
	assert.NotEmpty(t, response.Links)
	assert.NotEmpty(t, response.Links["payment"])

}

func TestCaptureCardPaymentIdempotentlyPrevious(t *testing.T) {

	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	Wait(time.Duration(3))

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := abc.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
	}

	idempotencyKey := uuid.New().String()

	response1, err := PreviousApi().Payments.CapturePayment(paymentResponse.Id, captureRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response1)

	response2, err := PreviousApi().Payments.CapturePayment(paymentResponse.Id, captureRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, response2)

	assert.Equal(t, response1.ActionId, response2.ActionId)

}
