package test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/abc"
)

func TestCaptureCardPaymentPrevious(t *testing.T) {
	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := abc.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
		Amount:    5,
	}

	cases := []struct {
		name           string
		paymentId      string
		captureRequest abc.CaptureRequest
		checker        func(interface{}, error)
	}{
		{
			name:           "when request is valid then capture payment",
			paymentId:      paymentResponse.Id,
			captureRequest: captureRequest,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.(*payments.CaptureResponse).Reference)
				assert.NotEmpty(t, response.(*payments.CaptureResponse).ActionId)
				assert.NotEmpty(t, response.(*payments.CaptureResponse).Links)
				assert.NotEmpty(t, response.(*payments.CaptureResponse).Links["payment"])
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			process := func() (interface{}, error) { return client.CapturePayment(tc.paymentId, tc.captureRequest, nil) }
			predicate := func(data interface{}) bool {
				response := data.(*payments.CaptureResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			tc.checker(retriable(process, predicate, 2))
		})
	}
}

func TestCaptureCardPaymentIdempotentlyPrevious(t *testing.T) {
	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := abc.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
	}

	idempotencyKeyRandom1 := uuid.New().String()

	idempotencyKeyRandom2 := uuid.New().String()

	cases := []struct {
		name                  string
		paymentId             string
		captureRequest        abc.CaptureRequest
		idempotencyKeyRandom1 string
		idempotencyKeyRandom2 string
		checker               func(interface{}, error, interface{}, error)
	}{
		{
			name:                  "when request is valid then capture payment idempotently",
			paymentId:             paymentResponse.Id,
			captureRequest:        captureRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom1,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.Equal(t, response1.(*payments.CaptureResponse).ActionId, response2.(*payments.CaptureResponse).ActionId)
			},
		},
		{
			name:                  "when request is valid then capture payment idempotently error",
			paymentId:             paymentResponse.Id,
			captureRequest:        captureRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom2,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.NotNil(t, err2)
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			processOne := func() (interface{}, error) {
				return client.CapturePayment(tc.paymentId, tc.captureRequest, &tc.idempotencyKeyRandom1)
			}
			predicateOne := func(data interface{}) bool {
				response := data.(*payments.CaptureResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			processTwo := func() (interface{}, error) {
				return client.CapturePayment(tc.paymentId, tc.captureRequest, &tc.idempotencyKeyRandom2)
			}
			predicateTwo := func(data interface{}) bool {
				response := data.(*payments.CaptureResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			retriableOne, errOne := retriable(processOne, predicateOne, 2)
			retriableTwo, errTwo := retriable(processTwo, predicateTwo, 2)
			tc.checker(retriableOne, errOne, retriableTwo, errTwo)
		})
	}
}
