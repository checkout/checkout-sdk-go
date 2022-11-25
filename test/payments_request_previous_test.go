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

func TestRequestCardPaymentPrevious(t *testing.T) {

	paymentResponse := makeCardPaymentPrevious(t, false, 10)
	assert.NotNil(t, paymentResponse)

	assert.NotEmpty(t, paymentResponse.Id)
	assert.NotEmpty(t, paymentResponse.ProcessedOn)
	assert.NotEmpty(t, paymentResponse.Reference)
	assert.NotEmpty(t, paymentResponse.ActionId)
	assert.NotEmpty(t, paymentResponse.ResponseCode)
	assert.NotEmpty(t, paymentResponse.SchemeId)
	assert.NotEmpty(t, paymentResponse.ResponseSummary)
	assert.Equal(t, payments.Authorized, paymentResponse.Status)
	assert.Equal(t, 10, paymentResponse.Amount)
	assert.True(t, paymentResponse.Approved)
	assert.NotEmpty(t, paymentResponse.AuthCode)
	assert.NotEmpty(t, paymentResponse.Currency)
	assert.Nil(t, paymentResponse.ThreeDs)

	//Source
	assert.NotEmpty(t, paymentResponse.Source)
	responseCardSource := paymentResponse.Source.ResponseCardSource
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

	//Customer
	assert.NotEmpty(t, paymentResponse.Customer)
	customer := paymentResponse.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)
	assert.NotEmpty(t, customer.Name)

	//Processing
	assert.NotEmpty(t, paymentResponse.Processing)
	processing := paymentResponse.Processing
	assert.NotEmpty(t, processing)
	assert.NotEmpty(t, processing.AcquirerTransactionId)
	assert.NotEmpty(t, processing.RetrievalReferenceNumber)

	//Risk
	assert.False(t, paymentResponse.Risk.Flagged)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["actions"])
	assert.NotEmpty(t, paymentResponse.Links["capture"])
	assert.NotEmpty(t, paymentResponse.Links["void"])

}

func TestMakeCardVerificationPrevious(t *testing.T) {

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

	paymentResponse, err := PreviousApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, paymentResponse)

	assert.NotEmpty(t, paymentResponse.Id)
	assert.NotEmpty(t, paymentResponse.ProcessedOn)
	assert.NotEmpty(t, paymentResponse.Reference)
	assert.NotEmpty(t, paymentResponse.ActionId)
	assert.NotEmpty(t, paymentResponse.ResponseCode)
	assert.NotEmpty(t, paymentResponse.SchemeId)
	assert.NotEmpty(t, paymentResponse.ResponseSummary)
	assert.Equal(t, payments.CardVerified, paymentResponse.Status)
	assert.Equal(t, 0, paymentResponse.Amount)
	assert.True(t, paymentResponse.Approved)
	assert.NotEmpty(t, paymentResponse.AuthCode)
	assert.NotEmpty(t, paymentResponse.Currency)
	assert.Nil(t, paymentResponse.ThreeDs)

	//Source
	assert.NotEmpty(t, paymentResponse.Source)
	responseCardSource := paymentResponse.Source.ResponseCardSource
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

	//Customer
	assert.NotEmpty(t, paymentResponse.Customer)
	customer := paymentResponse.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)
	assert.NotEmpty(t, customer.Name)

	//Processing
	assert.NotEmpty(t, paymentResponse.Processing)
	processing := paymentResponse.Processing
	assert.NotEmpty(t, processing)
	assert.NotEmpty(t, processing.AcquirerTransactionId)
	assert.NotEmpty(t, processing.RetrievalReferenceNumber)

	//Risk
	assert.False(t, paymentResponse.Risk.Flagged)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["actions"])
	assert.Empty(t, paymentResponse.Links["capture"])
	assert.Empty(t, paymentResponse.Links["void"])

}

func TestMakeCard3dsPaymentPrevious(t *testing.T) {

	paymentResponse := make3dsCardPaymentPrevious(t, false)
	assert.NotNil(t, paymentResponse)

	assert.NotEmpty(t, paymentResponse.Id)
	assert.NotEmpty(t, paymentResponse.Reference)
	assert.Equal(t, payments.Pending, paymentResponse.Status)

	//3ds
	assert.NotEmpty(t, paymentResponse.ThreeDs)
	assert.False(t, paymentResponse.ThreeDs.Downgraded)
	assert.Equal(t, payments.Yes, paymentResponse.ThreeDs.Enrolled)

	//Customer
	assert.NotEmpty(t, paymentResponse.Customer)
	customer := paymentResponse.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)
	assert.NotEmpty(t, customer.Name)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["redirect"])

}

func TestMakeCardN3dPaymentPrevious(t *testing.T) {

	paymentResponse := make3dsCardPaymentPrevious(t, true)
	assert.NotNil(t, paymentResponse)

	assert.NotEmpty(t, paymentResponse.Id)
	assert.NotEmpty(t, paymentResponse.ProcessedOn)
	assert.NotEmpty(t, paymentResponse.Reference)
	assert.NotEmpty(t, paymentResponse.ActionId)
	assert.NotEmpty(t, paymentResponse.ResponseCode)
	assert.NotEmpty(t, paymentResponse.SchemeId)
	assert.NotEmpty(t, paymentResponse.ResponseSummary)
	assert.Equal(t, payments.Authorized, paymentResponse.Status)
	assert.Equal(t, 10, paymentResponse.Amount)
	assert.True(t, paymentResponse.Approved)
	assert.NotEmpty(t, paymentResponse.AuthCode)
	assert.NotEmpty(t, paymentResponse.Currency)
	assert.Nil(t, paymentResponse.ThreeDs)

	//Source
	assert.NotEmpty(t, paymentResponse.Source)
	responseCardSource := paymentResponse.Source.ResponseCardSource
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

	//Customer
	assert.NotEmpty(t, paymentResponse.Customer)
	customer := paymentResponse.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)
	assert.NotEmpty(t, customer.Name)

	//Processing
	assert.NotEmpty(t, paymentResponse.Processing)
	processing := paymentResponse.Processing
	assert.NotEmpty(t, processing)
	assert.NotEmpty(t, processing.AcquirerTransactionId)
	assert.NotEmpty(t, processing.RetrievalReferenceNumber)

	//Risk
	assert.False(t, paymentResponse.Risk.Flagged)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["actions"])
	assert.NotEmpty(t, paymentResponse.Links["capture"])
	assert.NotEmpty(t, paymentResponse.Links["void"])

}

func TestRequestCardTokenPaymentPrevious(t *testing.T) {

	paymentResponse := makeCardTokenPaymentPrevious(t)
	assert.NotNil(t, paymentResponse)

	assert.NotEmpty(t, paymentResponse.Id)
	assert.NotEmpty(t, paymentResponse.ProcessedOn)
	assert.NotEmpty(t, paymentResponse.Reference)
	assert.NotEmpty(t, paymentResponse.ActionId)
	assert.NotEmpty(t, paymentResponse.ResponseCode)
	assert.NotEmpty(t, paymentResponse.SchemeId)
	assert.NotEmpty(t, paymentResponse.ResponseSummary)
	assert.Equal(t, payments.Authorized, paymentResponse.Status)
	assert.Equal(t, 10, paymentResponse.Amount)
	assert.True(t, paymentResponse.Approved)
	assert.NotEmpty(t, paymentResponse.AuthCode)
	assert.NotEmpty(t, paymentResponse.Currency)
	assert.Nil(t, paymentResponse.ThreeDs)

	//Source
	assert.NotEmpty(t, paymentResponse.Source)
	responseCardSource := paymentResponse.Source.ResponseCardSource
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

	//Customer
	assert.NotEmpty(t, paymentResponse.Customer)
	customer := paymentResponse.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)
	assert.NotEmpty(t, customer.Name)

	//Processing
	assert.NotEmpty(t, paymentResponse.Processing)
	processing := paymentResponse.Processing
	assert.NotEmpty(t, processing)
	assert.NotEmpty(t, processing.AcquirerTransactionId)
	assert.NotEmpty(t, processing.RetrievalReferenceNumber)

	//Risk
	assert.False(t, paymentResponse.Risk.Flagged)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["actions"])
	assert.NotEmpty(t, paymentResponse.Links["capture"])
	assert.NotEmpty(t, paymentResponse.Links["void"])

}

func TestMakePaymentsIdempotentlyPrevious(t *testing.T) {

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

	idempotencyKey := uuid.New().String()

	paymentResponse1, err := PreviousApi().Payments.RequestPayment(paymentRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, paymentResponse1)

	paymentResponse2, err := PreviousApi().Payments.RequestPayment(paymentRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, paymentResponse2)

	assert.Equal(t, paymentResponse1.ActionId, paymentResponse2.ActionId)

}

func makeCardPaymentPrevious(t *testing.T, shouldCapture bool, amount int) *abc.PaymentResponse {

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

func make3dsCardPaymentPrevious(t *testing.T, attemptN3d bool) *abc.PaymentResponse {

	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	threeDsRequest := &payments.ThreeDsRequest{
		Enabled:    true,
		AttemptN3D: false,
		Version:    "2.0.1",
	}

	if attemptN3d {
		threeDsRequest.AttemptN3D = true
		threeDsRequest.Eci = "05"
		threeDsRequest.Cryptogram = "AgAAAAAAAIR8CQrXcIhbQAAAAAA"
		threeDsRequest.Xid = "MDAwMDAwMDAwMDAwMDAwMzIyNzY"
	}

	paymentRequest := abc.PaymentRequest{
		Source:         cardSource,
		Amount:         10,
		Currency:       common.GBP,
		Reference:      Reference,
		Description:    "description",
		ThreeDsRequest: threeDsRequest,
	}

	response, err := PreviousApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	return response
}

func makeCardTokenPaymentPrevious(t *testing.T) *abc.PaymentResponse {

	tokenSource := sources.NewRequestTokenSource()
	tokenSource.Token = RequestCardTokenPrevious(t).Token

	paymentRequest := abc.PaymentRequest{
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

	response, err := PreviousApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	return response
}
