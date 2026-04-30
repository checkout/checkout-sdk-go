package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/payments"
)

// tests

func TestCancelAScheduledPaymentRetry(t *testing.T) {
	t.Skip("Use on demand: requires a payment with an active scheduled retry")

	paymentResponse := makeCardPayment(t, false, 10)

	cases := []struct {
		name      string
		paymentId string
		request   payments.CancellationRequest
		checker   func(*payments.CancellationResponse, error)
	}{
		{
			name:      "when request is valid then return a cancellation response",
			paymentId: paymentResponse.Id,
			request:   buildCancellationIntegrationRequest(),
			checker: func(response *payments.CancellationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.ActionId)
				assert.NotEmpty(t, response.Reference)
				assert.NotEmpty(t, response.Links)
				assert.NotEmpty(t, response.Links["payment"])
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checker(client.CancelAScheduledRetry(tc.paymentId, &tc.request, nil))
		})
	}
}

func TestCancelAScheduledPaymentRetryIdempotently(t *testing.T) {
	t.Skip("Use on demand: requires a payment with an active scheduled retry")

	paymentResponse := makeCardPayment(t, false, 10)

	idempotencyKey := uuid.New().String()

	cases := []struct {
		name           string
		paymentId      string
		idempotencyKey string
		checker        func(*payments.CancellationResponse, error)
	}{
		{
			name:           "when request is valid with idempotencyKey then return a cancellation response",
			paymentId:      paymentResponse.Id,
			idempotencyKey: idempotencyKey,
			checker: func(response *payments.CancellationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)

				response2, err := DefaultApi().Payments.CancelAScheduledRetry(paymentResponse.Id, nil, &idempotencyKey)
				assert.Nil(t, err)
				assert.NotNil(t, response2)

				assert.Equal(t, response.ActionId, response2.ActionId)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checker(client.CancelAScheduledRetry(tc.paymentId, nil, &tc.idempotencyKey))
		})
	}
}

// common methods

func buildCancellationIntegrationRequest() payments.CancellationRequest {
	return payments.CancellationRequest{
		Reference: uuid.New().String(),
	}
}
