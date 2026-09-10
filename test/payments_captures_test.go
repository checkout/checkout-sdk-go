package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

func TestCaptureCardPayment(t *testing.T) {
	// A partial capture: half of the authorized amount is captured.
	const authorizedAmount = int64(10)
	const capturedAmount = int64(5)

	paymentResponse := makeCardPayment(t, false, authorizedAmount)

	metadata := make(map[string]interface{})
	metadata["TestCaptureCardPayment"] = "metadata"

	captureRequest := nas.CaptureRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
		Amount:    capturedAmount,
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
				assert.Equal(t, authorizedAmount, response.Balances.TotalAuthorized)
				assert.Equal(t, capturedAmount, response.Balances.TotalCaptured)
				assert.Equal(t, int64(0), response.Balances.TotalRefunded)
				// The uncaptured remainder is released, not voided: it is neither available
				// to capture nor to void, and total_voided stays 0. Observed stable from
				// t+2s to t+24s against the sandbox on 2026-09-10.
				assert.Equal(t, int64(0), response.Balances.TotalVoided)
				assert.Equal(t, int64(0), response.Balances.AvailableToCapture)
				assert.Equal(t, capturedAmount, response.Balances.AvailableToRefund)
				assert.Equal(t, int64(0), response.Balances.AvailableToVoid)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checkerOne(client.CapturePayment(tc.paymentId, tc.captureRequest, nil))

			// Poll until the capture is reflected, bounded to
			// MaxRetryAttemps (10) attempts * 1s wait = ~10s total. This mirrors the
			// .NET suite, which polls this endpoint on TotalCaptured for this scenario.
			process := func() (interface{}, error) {
				return client.GetPaymentDetails(tc.paymentId)
			}
			predicate := func(data interface{}) bool {
				response, ok := data.(*nas.GetPaymentResponse)
				return ok && response != nil && response.Balances != nil &&
					response.Balances.TotalCaptured == capturedAmount
			}

			response, err := retriable(process, predicate, 1)
			if err != nil {
				t.Fatalf("GetPaymentDetails failed while waiting for the capture to settle: %v", err)
			}
			// retriable returns a nil response with a nil error when it exhausts its
			// attempts. Balances is a pointer, so handing that straight to the checker
			// would panic instead of reporting why the test failed.
			typed, ok := response.(*nas.GetPaymentResponse)
			if !ok || typed == nil || typed.Balances == nil {
				t.Fatalf(
					"payment balances did not reflect the capture of %d after %d attempts",
					capturedAmount, MaxRetryAttemps)
			}

			tc.checkerTwo(typed, err)
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
