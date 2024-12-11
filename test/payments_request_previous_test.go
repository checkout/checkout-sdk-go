package test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/abc"
	"github.com/checkout/checkout-sdk-go/payments/abc/sources"
)

func TestRequestPaymentListPrevious(t *testing.T) {
	t.Skip("unavailable")
	paymentResponse := makeCardPaymentPrevious(t, false, 10)

	queryRequest := payments.QueryRequest{
		Limit:     1,
		Skip:      0,
		Reference: paymentResponse.Reference,
	}

	cases := []struct {
		name         string
		queryRequest payments.QueryRequest
		checker      func(*abc.GetPaymentListResponse, error)
	}{
		{
			name:         "when payment is valid then return payment list",
			queryRequest: queryRequest,
			checker: func(response *abc.GetPaymentListResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, 1, response.Limit)
				assert.Equal(t, 0, response.Skip)
				assert.NotNil(t, response.TotalCount)
				assert.NotNil(t, response.Data)
				assert.NotNil(t, response.Data[0].Id)
				assert.NotNil(t, response.Data[0].RequestedOn)
				assert.NotNil(t, response.Data[0].Source)

				paymentCommonAssertionsPrevious(t, paymentResponse)
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestPaymentList(tc.queryRequest))
		})
	}
}

func TestRequestPaymentPrevious(t *testing.T) {
	t.Skip("unavailable")
	tokenSource := sources.NewRequestTokenSource()
	tokenSource.Token = RequestCardTokenPrevious(t).Token

	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	paymentRequestAuthorized := abc.PaymentRequest{
		Source:      cardSource,
		Amount:      0,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
	}

	paymentRequestCardVerified := abc.PaymentRequest{
		Source:      tokenSource,
		Amount:      10,
		Currency:    common.USD,
		Reference:   Reference,
		Description: "description",
		Customer: &common.CustomerRequest{
			Email: Email,
		},
		BillingDescriptor: &payments.BillingDescriptor{
			Name:      Name,
			City:      "London",
			Reference: Reference,
		},
	}

	paymentRequest3DSTrue := abc.PaymentRequest{
		Source:      cardSource,
		Amount:      10,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
		ThreeDsRequest: &payments.ThreeDsRequest{
			Enabled:    true,
			AttemptN3D: true,
			Eci:        "05",
			Cryptogram: "AgAAAAAAAIR8CQrXcIhbQAAAAAA",
			Xid:        "MDAwMDAwMDAwMDAwMDAwMzIyNzY",
			Version:    "2.0.1",
		},
	}

	paymentRequest3DSFalse := abc.PaymentRequest{
		Source:      cardSource,
		Amount:      10,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
		ThreeDsRequest: &payments.ThreeDsRequest{
			Enabled:    true,
			AttemptN3D: false,
			Version:    "2.0.1",
		},
	}

	cases := []struct {
		name           string
		paymentRequest abc.PaymentRequest
		checker        func(*abc.PaymentResponse, error)
	}{
		{
			name:           "when get a payment card source request then return a payment response",
			paymentRequest: paymentRequestAuthorized,
			checker: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)

				assert.NotEmpty(t, response.Id)
				assert.NotEmpty(t, response.ProcessedOn)
				assert.NotEmpty(t, response.Reference)
				assert.NotEmpty(t, response.ActionId)
				assert.NotEmpty(t, response.ResponseCode)
				assert.NotEmpty(t, response.SchemeId)
				assert.NotEmpty(t, response.ResponseSummary)
				assert.Equal(t, payments.CardVerified, response.Status)
				assert.Equal(t, int64(0), response.Amount)
				assert.True(t, response.Approved)
				assert.NotEmpty(t, response.AuthCode)
				assert.NotEmpty(t, response.Currency)
				assert.Nil(t, response.ThreeDs)

				assertSourcePrevious(t, response)

				assertCustomerPrevious(t, response)

				assertProcessingPrevious(t, response)

				assertRiskPrevious(t, response)

				assert.NotEmpty(t, response.Links["self"])
				assert.NotEmpty(t, response.Links["actions"])
				assert.Empty(t, response.Links["capture"])
				assert.Empty(t, response.Links["void"])
			},
		},
		{
			name:           "when request is valid with attemptN3d then return a payment response",
			paymentRequest: paymentRequest3DSTrue,
			checker: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				paymentCommonAssertionsPrevious(t, response)
			},
		},
		{
			name:           "when request valid without attemptN3d then return a payment response",
			paymentRequest: paymentRequest3DSFalse,
			checker: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)

				assert.NotEmpty(t, response.Id)
				assert.NotEmpty(t, response.Reference)
				assert.Equal(t, payments.Pending, response.Status)

				//3ds
				assert.NotEmpty(t, response.ThreeDs)
				assert.False(t, response.ThreeDs.Downgraded)
				assert.Equal(t, payments.Yes, response.ThreeDs.Enrolled)

				//Customer
				assert.NotEmpty(t, response.Customer)
				customer := response.Customer
				assert.NotEmpty(t, customer)
				assert.NotEmpty(t, customer.Id)
				assert.NotEmpty(t, customer.Name)

				//Links
				assert.NotEmpty(t, response.Links["self"])
				assert.NotEmpty(t, response.Links["redirect"])
			},
		},
		{
			name:           "when request valid then return a payment response with card verified",
			paymentRequest: paymentRequestCardVerified,
			checker: func(response *abc.PaymentResponse, err error) {
				assert.Nil(t, err)
				paymentCommonAssertionsPrevious(t, response)
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestPayment(tc.paymentRequest, nil))
		})
	}
}

func TestRequestPaymentIdempotentlyPrevious(t *testing.T) {
	t.Skip("unavailable")
	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	paymentRequest := abc.PaymentRequest{
		Source:      cardSource,
		Amount:      0,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
	}

	idempotencyKeyRandom1 := uuid.New().String()

	idempotencyKeyRandom2 := uuid.New().String()

	cases := []struct {
		name                  string
		paymentRequest        abc.PaymentRequest
		idempotencyKeyRandom1 string
		idempotencyKeyRandom2 string
		checker               func(interface{}, error, interface{}, error)
	}{
		{
			name:                  "when get a request payment with idempotencyKey then return a payment response",
			paymentRequest:        paymentRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom1,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.Equal(t, response1.(*abc.PaymentResponse).ActionId, response2.(*abc.PaymentResponse).ActionId)
			},
		},
		{
			name:                  "when get a request payment with idempotencyKey different then return a payment response",
			paymentRequest:        paymentRequest,
			idempotencyKeyRandom1: idempotencyKeyRandom1,
			idempotencyKeyRandom2: idempotencyKeyRandom2,
			checker: func(response1 interface{}, err1 error, response2 interface{}, err2 error) {
				assert.Nil(t, err1)
				assert.NotNil(t, response1)
				assert.Nil(t, err2)
				assert.NotNil(t, response2)
				assert.NotEqual(t, response1.(*abc.PaymentResponse).ActionId, response2.(*abc.PaymentResponse).ActionId)
			},
		},
	}

	client := PreviousApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			processOne := func() (interface{}, error) {
				return client.RequestPayment(tc.paymentRequest, &tc.idempotencyKeyRandom1)
			}
			predicateOne := func(data interface{}) bool {
				response := data.(*abc.PaymentResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			processTwo := func() (interface{}, error) {
				return client.RequestPayment(tc.paymentRequest, &tc.idempotencyKeyRandom2)
			}
			predicateTwo := func(data interface{}) bool {
				response := data.(*abc.PaymentResponse)
				return response.Links != nil && len(response.Links) >= 0
			}

			retriableOne, errOne := retriable(processOne, predicateOne, 2)
			retriableTwo, errTwo := retriable(processTwo, predicateTwo, 2)
			tc.checker(retriableOne, errOne, retriableTwo, errTwo)
		})
	}
}

func makeCardPaymentPrevious(t *testing.T, shouldCapture bool, amount int64) *abc.PaymentResponse {
	t.Skip("unavailable")
	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	paymentRequest := abc.PaymentRequest{
		Source:      cardSource,
		Amount:      amount,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
		Capture:     shouldCapture,
	}

	response, err := PreviousApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	return response
}

func paymentCommonAssertionsPrevious(t *testing.T, response *abc.PaymentResponse) {
	t.Skip("unavailable")
	assert.NotNil(t, response)

	assertAuthorizedPaymentPrevious(t, response)

	assertSourcePrevious(t, response)

	assertCustomerPrevious(t, response)

	assertProcessingPrevious(t, response)

	assertRiskPrevious(t, response)

	assertLinksPrevious(t, response)
}

func assertAuthorizedPaymentPrevious(t *testing.T, response *abc.PaymentResponse) {
	t.Skip("unavailable")
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.ProcessedOn)
	assert.NotEmpty(t, response.Reference)
	assert.NotEmpty(t, response.ActionId)
	assert.NotEmpty(t, response.ResponseCode)
	assert.NotEmpty(t, response.SchemeId)
	assert.NotEmpty(t, response.ResponseSummary)
	assert.Equal(t, payments.Authorized, response.Status)
	assert.Equal(t, int64(10), response.Amount)
	assert.True(t, response.Approved)
	assert.NotEmpty(t, response.AuthCode)
	assert.NotEmpty(t, response.Currency)
	assert.Nil(t, response.ThreeDs)
}

func assertSourcePrevious(t *testing.T, response *abc.PaymentResponse) {
	t.Skip("unavailable")
	assert.NotEmpty(t, response.Source)
	responseCardSource := response.Source.ResponseCardSource
	assert.NotEmpty(t, payments.CardSource, responseCardSource.Type)
	assert.NotEmpty(t, responseCardSource.Id)
	assert.NotEmpty(t, responseCardSource.AvsCheck)
	assert.NotEmpty(t, responseCardSource.CvvCheck)
	assert.NotEmpty(t, responseCardSource.Bin)
	assert.NotEmpty(t, responseCardSource.ExpiryYear)
	assert.NotEmpty(t, responseCardSource.ExpiryMonth)
	assert.NotEmpty(t, responseCardSource.Last4)
	assert.NotEmpty(t, responseCardSource.Name)
	assert.NotEmpty(t, responseCardSource.FastFunds)
	assert.NotEmpty(t, responseCardSource.Fingerprint)
	assert.True(t, responseCardSource.Payouts)
}

func assertCustomerPrevious(t *testing.T, response *abc.PaymentResponse) {
	t.Skip("unavailable")
	assert.NotEmpty(t, response.Customer)
	customer := response.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)
	assert.NotEmpty(t, customer.Name)
}

func assertProcessingPrevious(t *testing.T, response *abc.PaymentResponse) {
	t.Skip("unavailable")
	assert.NotEmpty(t, response.Processing)
	processing := response.Processing
	assert.NotEmpty(t, processing)
	assert.NotEmpty(t, processing.AcquirerTransactionId)
	assert.NotEmpty(t, processing.RetrievalReferenceNumber)
}

func assertRiskPrevious(t *testing.T, response *abc.PaymentResponse) {
	t.Skip("unavailable")
	assert.False(t, response.Risk.Flagged)
}

func assertLinksPrevious(t *testing.T, response *abc.PaymentResponse) {
	t.Skip("unavailable")
	assert.NotEmpty(t, response.Links["self"])
	assert.NotEmpty(t, response.Links["actions"])
	assert.NotEmpty(t, response.Links["capture"])
	assert.NotEmpty(t, response.Links["void"])
}
