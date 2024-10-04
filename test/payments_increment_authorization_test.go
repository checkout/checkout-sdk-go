package test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources"
)

func TestIncrementAuthorization(t *testing.T) {
	paymentResponse := makeCardPaymentPartialAuthorization(t, false, 10)
	assert.NotNil(t, paymentResponse, "Expected paymentResponse not to be nil")

	metadata := make(map[string]interface{})
	metadata["TestIncrementAuthorization"] = "metadata"

	incrementAuthorizationRequest := nas.IncrementAuthorizationRequest{
		Amount:    5,
		Reference: uuid.New().String(),
		Metadata:  metadata,
	}

	cases := []struct {
		name       string
		paymentId  string
		request    nas.IncrementAuthorizationRequest
		checkerOne func(*nas.IncrementAuthorizationResponse, error)
		checkerTwo func(*nas.GetPaymentResponse, error)
	}{
		{
			name:      "when request an increment authorization then return a response",
			paymentId: paymentResponse.Id,
			request:   incrementAuthorizationRequest,
			checkerOne: func(response *nas.IncrementAuthorizationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, int64(5), response.Amount)
				assert.NotEmpty(t, response.ActionId)
				assert.NotEmpty(t, response.Currency)
				assert.False(t, response.Approved)
				assert.NotEmpty(t, response.ResponseCode)
				assert.NotEmpty(t, response.ResponseSummary)
				assert.NotEmpty(t, response.ExpiresOn)
				assert.NotEmpty(t, response.ProcessedOn)
				assert.NotEmpty(t, response.Balances)
				assert.NotEmpty(t, response.Links)
			},
			checkerTwo: func(response *nas.GetPaymentResponse, err error) {
				assert.NotEmpty(t, response.Balances)
				assert.Equal(t, int64(10), response.Balances.TotalAuthorized)
				assert.Equal(t, int64(0), response.Balances.TotalCaptured)
				assert.Equal(t, int64(0), response.Balances.TotalRefunded)
				assert.Equal(t, int64(0), response.Balances.TotalVoided)
				assert.Equal(t, int64(10), response.Balances.AvailableToCapture)
				assert.Equal(t, int64(0), response.Balances.AvailableToRefund)
				assert.Equal(t, int64(10), response.Balances.AvailableToVoid)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Wait(time.Duration(3))
			tc.checkerOne(client.IncrementAuthorization(tc.paymentId, tc.request, nil))
			Wait(time.Duration(3))
			tc.checkerTwo(client.GetPaymentDetails(tc.paymentId))
		})
	}
}

func TestIncrementAuthorizationIdempotently(t *testing.T) {
	paymentResponse := makeCardPaymentPartialAuthorization(t, false, 10)
	assert.NotNil(t, paymentResponse, "Expected paymentResponse not to be nil")

	metadata := make(map[string]interface{})
	metadata["TestIncrementAuthorization"] = "metadata"

	incrementAuthorizationRequest := nas.IncrementAuthorizationRequest{
		Reference: uuid.New().String(),
		Metadata:  metadata,
	}

	idempotencyKeyRandom1 := uuid.New().String()

	idempotencyKeyRandom2 := uuid.New().String()

	cases := []struct {
		name                  string
		paymentId             string
		request               nas.IncrementAuthorizationRequest
		idempotencyKeyRandom1 string
		idempotencyKeyRandom2 string
		checker               func(interface{}, error, interface{}, error)
	}{
		{
			name:                  "when request is valid then increment authorization idempotently",
			paymentId:             paymentResponse.Id,
			request:               incrementAuthorizationRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom1,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.Equal(t, response1.(*nas.IncrementAuthorizationResponse).ActionId, response2.(*nas.IncrementAuthorizationResponse).ActionId)
			},
		},
		{
			name:                  "when request is valid then capture payment idempotently error",
			paymentId:             paymentResponse.Id,
			request:               incrementAuthorizationRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom2,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.NotEqual(t, response1.(*nas.IncrementAuthorizationResponse).ActionId, response2.(*nas.IncrementAuthorizationResponse).ActionId)
			},
		},
	}

	client := DefaultApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			processOne := func() (interface{}, error) {
				return client.IncrementAuthorization(tc.paymentId, tc.request, &tc.idempotencyKeyRandom1)
			}
			predicateOne := func(data interface{}) bool {
				response := data.(*nas.IncrementAuthorizationResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			processTwo := func() (interface{}, error) {
				return client.IncrementAuthorization(tc.paymentId, tc.request, &tc.idempotencyKeyRandom2)
			}
			predicateTwo := func(data interface{}) bool {
				response := data.(*nas.IncrementAuthorizationResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			retriableOne, errOne := retriable(processOne, predicateOne, 2)
			retriableTwo, errTwo := retriable(processTwo, predicateTwo, 2)
			tc.checker(retriableOne, errOne, retriableTwo, errTwo)
		})
	}
}

func makeCardPaymentPartialAuthorization(t *testing.T, shouldCapture bool, amount int64) *nas.PaymentResponse {
	currentYear := time.Now().Year() + 1

	cardSource := sources.NewRequestCardSource()
	cardSource.Name = "Mr. Test"
	cardSource.Number = "4556447238607884"
	cardSource.ExpiryYear = currentYear
	cardSource.ExpiryMonth = 12
	cardSource.Cvv = "123"
	cardSource.BillingAddress = &common.Address{
		AddressLine1: "CheckoutSdk.com",
		AddressLine2: "90 Tottenham Court Road",
		City:         "London",
		State:        "London",
		Zip:          "W1T 4TJ",
		Country:      common.GB,
	}
	cardSource.Phone = &common.Phone{
		CountryCode: "44",
		Number:      "1234567890",
	}

	customerRequest := &common.CustomerRequest{
		Email: "test@example.com",
		Name:  "Test Customer",
		Phone: &common.Phone{
			CountryCode: "44",
			Number:      "1234567890",
		},
	}

	paymentIndividualSender := nas.NewRequestIndividualSender()
	paymentIndividualSender.FirstName = "Mr"
	paymentIndividualSender.LastName = "Test"
	paymentIndividualSender.Address = &common.Address{
		AddressLine1: "CheckoutSdk.com",
		AddressLine2: "90 Tottenham Court Road",
		City:         "London",
		State:        "London",
		Zip:          "W1T 4TJ",
		Country:      common.GB,
	}

	paymentRequest := nas.PaymentRequest{
		Source:            cardSource,
		Amount:            amount,
		Currency:          common.USD,
		Reference:         uuid.New().String(),
		Description:       "Test Payment",
		Capture:           shouldCapture,
		Customer:          customerRequest,
		Sender:            paymentIndividualSender,
		AuthorizationType: nas.EstimatedAuthorizationType,
		PartialAuthorization: &nas.PartialAuthorization{
			Enabled: true,
		},
		BillingDescriptor: &payments.BillingDescriptor{
			Name: "CheckoutSdk.com",
			City: "London",
		},
	}

	response, err := DefaultApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err, "Expected no error in RequestPayment")
	assert.NotNil(t, response, "Expected response not to be nil")
	return response
}
