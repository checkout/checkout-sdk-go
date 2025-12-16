package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas"
)

func TestCaptureCardPayment(t *testing.T) {
	paymentResponse := makeCardPayment(t, false, 10)

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := nas.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
		Amount:    5,
	}

	cases := []struct {
		name           string
		paymentId      string
		captureRequest nas.CaptureRequest
		checkerOne     func(*payments.CaptureResponse, error)
		checkerTwo     func(*nas.GetPaymentResponse, error)
	}{
		{
			name:           "when get a capture payment request then return a response",
			paymentId:      paymentResponse.Id,
			captureRequest: captureRequest,
			checkerOne: func(response *payments.CaptureResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.Reference)
				assert.NotEmpty(t, response.ActionId)
				assert.NotEmpty(t, response.Links)
				assert.NotEmpty(t, response.Links["payment"])
			},
			checkerTwo: func(response *nas.GetPaymentResponse, err error) {
				assert.NotEmpty(t, response.Balances)
				assert.Equal(t, int64(10), response.Balances.TotalAuthorized)
				assert.Equal(t, int64(5), response.Balances.TotalCaptured)
				assert.Equal(t, int64(0), response.Balances.TotalRefunded)
				assert.Equal(t, int64(5), response.Balances.TotalVoided)
				assert.Equal(t, int64(0), response.Balances.AvailableToCapture)
				assert.Equal(t, int64(5), response.Balances.AvailableToRefund)
				assert.Equal(t, int64(0), response.Balances.AvailableToVoid)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checkerOne(client.CapturePayment(tc.paymentId, tc.captureRequest, nil))
			Wait(time.Duration(3))
			tc.checkerTwo(client.GetPaymentDetails(tc.paymentId))
		})
	}
}

func TestCaptureCardPaymentWithoutRequest(t *testing.T) {
	paymentResponse := makeCardPayment(t, false, 10)

	cases := []struct {
		name       string
		paymentId  string
		checkerOne func(*payments.CaptureResponse, error)
	}{
		{
			name:      "when get a capture payment request without request then return a response",
			paymentId: paymentResponse.Id,
			checkerOne: func(response *payments.CaptureResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.ActionId)
				assert.NotEmpty(t, response.Links)
				assert.NotEmpty(t, response.Links["payment"])
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checkerOne(client.CapturePaymentWithoutRequest(tc.paymentId, nil))
		})
	}
}

func TestCaptureCardPaymentIdempotently(t *testing.T) {
	t.Skip("unavailable")
	paymentResponse := makeCardPayment(t, false, 10)

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := nas.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
	}

	idempotencyKeyRandom1 := uuid.New().String()

	idempotencyKeyRandom2 := uuid.New().String()

	cases := []struct {
		name                  string
		paymentId             string
		captureRequest        nas.CaptureRequest
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
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.NotEqual(t, response1.(*payments.CaptureResponse).ActionId, response2.(*payments.CaptureResponse).ActionId)
			},
		},
	}

	client := DefaultApi().Payments

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

func TestCaptureCardPaymentWithoutrequestIdempotently(t *testing.T) {
	t.Skip("unavailable")
	paymentResponse := makeCardPayment(t, false, 10)

	idempotencyKeyRandom1 := uuid.New().String()

	idempotencyKeyRandom2 := uuid.New().String()

	cases := []struct {
		name                  string
		paymentId             string
		idempotencyKeyRandom1 string
		idempotencyKeyRandom2 string
		checker               func(interface{}, error, interface{}, error)
	}{
		{
			name:                  "when request is valid then capture payment without request idempotently",
			paymentId:             paymentResponse.Id,
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
			name:                  "when request is valid then capture without request payment idempotently error",
			paymentId:             paymentResponse.Id,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom2,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.NotEqual(t, response1.(*payments.CaptureResponse).ActionId, response2.(*payments.CaptureResponse).ActionId)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			processOne := func() (interface{}, error) {
				return client.CapturePaymentWithoutRequest(tc.paymentId, &tc.idempotencyKeyRandom1)
			}
			predicateOne := func(data interface{}) bool {
				response := data.(*payments.CaptureResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			processTwo := func() (interface{}, error) {
				return client.CapturePaymentWithoutRequest(tc.paymentId, &tc.idempotencyKeyRandom2)
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
