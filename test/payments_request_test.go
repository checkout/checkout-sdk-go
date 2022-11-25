package test

import (
	"github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

func TestRequestCardPayment(t *testing.T) {

	paymentResponse := makeCardPayment(t, false, 10)
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
	assert.NotEmpty(t, common.Consumer, responseCardSource.CardCategory)
	assert.NotEmpty(t, common.Credit, responseCardSource.CardType)
	assert.NotEmpty(t, responseCardSource.ExpiryYear)
	assert.NotEmpty(t, responseCardSource.ExpiryMonth)
	assert.NotEmpty(t, responseCardSource.Last4)
	assert.NotEmpty(t, responseCardSource.Name)
	assert.NotEmpty(t, responseCardSource.Fingerprint)
	assert.NotEmpty(t, responseCardSource.ProductId)
	assert.NotEmpty(t, responseCardSource.ProductType)

	//Customer
	assert.NotEmpty(t, paymentResponse.Customer)
	customer := paymentResponse.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)
	assert.NotEmpty(t, customer.Name)
	assert.NotEmpty(t, customer.Email)

	//Processing
	assert.NotEmpty(t, paymentResponse.Processing)
	processing := paymentResponse.Processing
	assert.NotEmpty(t, processing)
	assert.NotEmpty(t, processing.AcquirerTransactionId)
	assert.NotEmpty(t, processing.RetrievalReferenceNumber)

	//Risk
	assert.False(t, paymentResponse.Risk.Flagged)

	//Balances
	assert.NotEmpty(t, paymentResponse.Balances)
	assert.Equal(t, 10, paymentResponse.Balances.TotalAuthorized)
	assert.Equal(t, 0, paymentResponse.Balances.TotalCaptured)
	assert.Equal(t, 0, paymentResponse.Balances.TotalRefunded)
	assert.Equal(t, 0, paymentResponse.Balances.TotalVoided)
	assert.Equal(t, 10, paymentResponse.Balances.AvailableToCapture)
	assert.Equal(t, 0, paymentResponse.Balances.AvailableToRefund)
	assert.Equal(t, 10, paymentResponse.Balances.AvailableToVoid)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["actions"])
	assert.NotEmpty(t, paymentResponse.Links["capture"])
	assert.NotEmpty(t, paymentResponse.Links["void"])

}

func TestMakeCardVerification(t *testing.T) {

	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	paymentRequest := nas.PaymentRequest{
		Source:      cardSource,
		Amount:      0,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: Description,
	}

	paymentResponse, err := DefaultApi().Payments.RequestPayment(paymentRequest, nil)
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
	assert.NotEmpty(t, common.Consumer, responseCardSource.CardCategory)
	assert.NotEmpty(t, common.Credit, responseCardSource.CardType)
	assert.NotEmpty(t, responseCardSource.ExpiryYear)
	assert.NotEmpty(t, responseCardSource.ExpiryMonth)
	assert.NotEmpty(t, responseCardSource.Last4)
	assert.NotEmpty(t, responseCardSource.Name)
	assert.NotEmpty(t, responseCardSource.Fingerprint)
	assert.NotEmpty(t, responseCardSource.ProductId)
	assert.NotEmpty(t, responseCardSource.ProductType)

	//Customer
	assert.Empty(t, paymentResponse.Customer)

	//Processing
	assert.NotEmpty(t, paymentResponse.Processing)
	processing := paymentResponse.Processing
	assert.NotEmpty(t, processing)
	assert.NotEmpty(t, processing.AcquirerTransactionId)
	assert.NotEmpty(t, processing.RetrievalReferenceNumber)

	//Risk
	assert.False(t, paymentResponse.Risk.Flagged)

	//Balances
	assert.NotNil(t, paymentResponse.Balances)
	assert.Equal(t, 0, paymentResponse.Balances.TotalAuthorized)
	assert.Equal(t, 0, paymentResponse.Balances.TotalCaptured)
	assert.Equal(t, 0, paymentResponse.Balances.TotalRefunded)
	assert.Equal(t, 0, paymentResponse.Balances.TotalVoided)
	assert.Equal(t, 0, paymentResponse.Balances.AvailableToCapture)
	assert.Equal(t, 0, paymentResponse.Balances.AvailableToRefund)
	assert.Equal(t, 0, paymentResponse.Balances.AvailableToVoid)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["actions"])
	assert.Empty(t, paymentResponse.Links["capture"])
	assert.Empty(t, paymentResponse.Links["void"])

}

func TestMakeCard3dsPayment(t *testing.T) {

	paymentResponse := make3dsCardPayment(t, false)
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
	assert.NotEmpty(t, customer.Email)

	//Links
	assert.NotEmpty(t, paymentResponse.Links["self"])
	assert.NotEmpty(t, paymentResponse.Links["redirect"])

}

func TestMakeCardN3dPayment(t *testing.T) {

	paymentResponse := make3dsCardPayment(t, true)
	assert.NotNil(t, paymentResponse)

	assert.NotEmpty(t, paymentResponse.Id)
	assert.NotEmpty(t, paymentResponse.ProcessedOn)
	assert.NotEmpty(t, paymentResponse.Reference)
	assert.NotEmpty(t, paymentResponse.ActionId)
	assert.NotEmpty(t, paymentResponse.ResponseCode)
	assert.NotEmpty(t, paymentResponse.SchemeId)
	assert.Equal(t, "Approved", paymentResponse.ResponseSummary)
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
	assert.NotEmpty(t, common.Consumer, responseCardSource.CardCategory)
	assert.NotEmpty(t, common.Credit, responseCardSource.CardType)
	assert.NotEmpty(t, responseCardSource.ExpiryYear)
	assert.NotEmpty(t, responseCardSource.ExpiryMonth)
	assert.NotEmpty(t, responseCardSource.Last4)
	assert.NotEmpty(t, responseCardSource.Name)
	assert.NotEmpty(t, responseCardSource.Fingerprint)
	assert.NotEmpty(t, responseCardSource.ProductId)
	assert.NotEmpty(t, responseCardSource.ProductType)

	//Customer`
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

func TestRequestCardTokenPayment(t *testing.T) {

	paymentResponse := makeCardTokenPayment(t)
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
	assert.NotEmpty(t, common.Consumer, responseCardSource.CardCategory)
	assert.NotEmpty(t, common.Credit, responseCardSource.CardType)
	assert.NotEmpty(t, responseCardSource.ExpiryYear)
	assert.NotEmpty(t, responseCardSource.ExpiryMonth)
	assert.NotEmpty(t, responseCardSource.Last4)
	assert.NotEmpty(t, responseCardSource.Name)
	assert.NotEmpty(t, responseCardSource.Fingerprint)
	assert.NotEmpty(t, responseCardSource.ProductId)
	assert.NotEmpty(t, responseCardSource.ProductType)

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

func TestMakePaymentsIdempotently(t *testing.T) {

	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	paymentRequest := nas.PaymentRequest{
		Source:      cardSource,
		Amount:      0,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
	}

	idempotencyKey := uuid.New().String()

	paymentResponse1, err := DefaultApi().Payments.RequestPayment(paymentRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, paymentResponse1)

	paymentResponse2, err := DefaultApi().Payments.RequestPayment(paymentRequest, &idempotencyKey)
	assert.Nil(t, err)
	assert.NotNil(t, paymentResponse2)

	assert.Equal(t, paymentResponse1.ActionId, paymentResponse2.ActionId)

}

func makeCardPayment(t *testing.T, shouldCapture bool, amount int) *nas.PaymentResponse {

	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	customerRequest := common.CustomerRequest{
		Email: Email,
		Name:  Name,
		Phone: Phone(),
	}

	paymentIndividualSender := nas.NewPaymentIndividualSender()
	paymentIndividualSender.FirstName = FirstName
	paymentIndividualSender.LastName = LastName
	paymentIndividualSender.Address = Address()
	paymentIndividualSender.Identification = &nas.Identification{
		Type:           nas.DrivingLicence,
		Number:         "12345",
		IssuingCountry: common.GT,
	}

	paymentRequest := nas.PaymentRequest{
		Source:      cardSource,
		Amount:      amount,
		Currency:    common.GBP,
		Reference:   Reference,
		Description: "description",
		Capture:     shouldCapture,
		Customer:    &customerRequest,
		Sender:      paymentIndividualSender,
		BillingDescriptor: &payments.BillingDescriptor{
			Name:      Name,
			City:      City,
			Reference: Reference,
		},
	}

	response, err := DefaultApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	return response
}

func make3dsCardPayment(t *testing.T, attemptN3d bool) *nas.PaymentResponse {

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

	customerRequest := common.CustomerRequest{
		Email: Email,
		Name:  Name,
		Phone: Phone(),
	}

	paymentCorporateSender := nas.NewPaymentCorporateSender()
	paymentCorporateSender.CompanyName = Name
	paymentCorporateSender.Address = Address()

	paymentRequest := nas.PaymentRequest{
		Source:         cardSource,
		Amount:         10,
		Currency:       common.GBP,
		Customer:       &customerRequest,
		Reference:      Reference,
		Description:    Description,
		ThreeDsRequest: threeDsRequest,
		Sender:         paymentCorporateSender,
		SuccessUrl:     SuccessUrl,
		FailureUrl:     FailureUrl,
	}

	response, err := DefaultApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	return response
}

func makeCardTokenPayment(t *testing.T) *nas.PaymentResponse {

	tokenSource := sources.NewRequestTokenSource()
	tokenSource.Token = RequestCardToken(t).Token

	paymentRequest := nas.PaymentRequest{
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
		Sender: nas.NewPaymentInstrumentSender(),
	}

	response, err := DefaultApi().Payments.RequestPayment(paymentRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	return response
}
