package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/hosted"
)

func TestCreateHostedPaymentsPageSession(t *testing.T) {
	cases := []struct {
		name    string
		request hosted.HostedPaymentRequest
		checker func(*hosted.HostedPaymentResponse, error)
	}{
		{
			name:    "when request is valid then create hosted payment session",
			request: *getHostedPaymentRequest(),
			checker: func(response *hosted.HostedPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Reference)
				assert.NotNil(t, response.Links)
				assert.Equal(t, "Hosted Payment Reference", response.Reference)
				assert.NotNil(t, response.Links["redirect"])
			},
		},
		{

			name:    "when request is invalid then return error",
			request: hosted.HostedPaymentRequest{},
			checker: func(response *hosted.HostedPaymentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().Hosted

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateHostedPaymentsPageSession(tc.request))
		})
	}
}

func TestGetHostedPaymentsPageDetails(t *testing.T) {
	hostedPaymentSession, err := DefaultApi().Hosted.CreateHostedPaymentsPageSession(*getHostedPaymentRequest())
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Error creating hosted payment session: %s", err.Error()))
	}

	cases := []struct {
		name            string
		hostedPaymentId string
		checker         func(*hosted.HostedPaymentDetails, error)
	}{
		{
			name:            "when fetching existing hosted payment session then return session details",
			hostedPaymentId: hostedPaymentSession.Id,
			checker: func(response *hosted.HostedPaymentDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.PaymentId)
				assert.NotNil(t, response.Amount)
				assert.NotNil(t, response.Currency)
				assert.NotNil(t, response.Reference)

				assert.Equal(t, common.GBP, response.Currency)
				assert.Equal(t, hosted.PaymentPending, response.Status)
			},
		},
		{
			name:            "when hostedPaymentId is inexistent then return error",
			hostedPaymentId: "not_found",
			checker: func(response *hosted.HostedPaymentDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().Hosted

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetHostedPaymentsPageDetails(tc.hostedPaymentId))
		})
	}
}

func getHostedPaymentRequest() *hosted.HostedPaymentRequest {
	now = time.Now()

	return &hosted.HostedPaymentRequest{
		Amount:    1000,
		Currency:  common.GBP,
		PaymentIp: "0.0.0.0",
		BillingDescriptor: &payments.BillingDescriptor{
			Name:      "Company Name",
			City:      "City",
			Reference: "Billing Descriptor Reference",
		},
		Reference:           "Hosted Payment Reference",
		Description:         "Hosted Payment Description",
		DisplayName:         "Company Name",
		ProcessingChannelId: os.Getenv("CHECKOUT_PROCESSING_CHANNEL_ID"),
		AmountAllocations: []common.AmountAllocations{
			{
				Id:        "ent_w4jelhppmfiufdnatam37wrfc4",
				Amount:    1000,
				Reference: "Entity Reference",
				Commission: &common.Commission{
					Amount:     1000,
					Percentage: 1.125,
				},
			},
		},
		Customer: &common.CustomerRequest{
			Email: GenerateRandomEmail(),
			Name:  "Customer Name",
		},
		Shipping: &payments.ShippingDetails{
			Address: &common.Address{
				AddressLine1: "Address",
				AddressLine2: "Road",
				City:         "City",
				State:        "State",
				Zip:          "Zip Code",
				Country:      common.GB,
			},
			Phone: &common.Phone{
				CountryCode: "1",
				Number:      "1234567890",
			},
		},
		Billing: &payments.BillingInformation{
			Address: Address(),
			Phone:   Phone(),
		},
		Recipient: &payments.PaymentRecipient{
			DateOfBirth:   "1980-01-01",
			AccountNumber: "1234567890",
			Address:       Address(),
			Zip:           "12345",
			FirstName:     "Recipient First Name",
			LastName:      "Recipient Last Name",
		},
		Processing: &payments.ProcessingSettings{
			Aft: true,
		},
		AllowPaymentMethods: []payments.SourceType{
			payments.CardSource,
			payments.GooglepaySource,
			payments.ApplepaySource,
		},
		DisabledPaymentMethods: []payments.SourceType{
			payments.EpsSource,
			payments.IdealSource,
			payments.KnetSource,
		},
		Products: []payments.Product{
			{
				Reference: "Product Reference",
				Name:      "Product Name",
				Quantity:  1,
				Price:     1000,
			},
		},
		Risk: &payments.RiskRequest{
			Enabled: false,
		},
		SuccessUrl: "https://example.com/payments/success",
		CancelUrl:  "https://example.com/payments/cancel",
		FailureUrl: "https://example.com/payments/failure",
		Locale:     payments.EnGBLT,
		ThreeDs: &payments.ThreeDsRequest{
			Enabled:            false,
			AttemptN3D:         false,
			ChallengeIndicator: common.NoPreference,
			AllowUpgrade:       true,
			Exemption:          payments.LowValue,
		},
		Capture:   true,
		CaptureOn: &now,
		Instruction: &payments.PaymentInstruction{
			Purpose: payments.DonationsPPT,
		},
		PaymentMethodConfiguration: &payments.PaymentMethodConfiguration{
			Applepay: &payments.Applepay{
				AccountHolder: &common.AccountHolder{
					FirstName: "Account Holder First Name",
					LastName:  "Account Holder Last Name",
					Type:      common.Individual,
				},
			},
			Card: &payments.Card{
				AccountHolder: &common.AccountHolder{
					FirstName: "Account Holder First Name",
					LastName:  "Account Holder Last Name",
					Type:      common.Individual,
				},
			},
			Googlepay: &payments.Googlepay{
				AccountHolder: &common.AccountHolder{
					FirstName: "Account Holder First Name",
					LastName:  "Account Holder Last Name",
					Type:      common.Individual,
				},
			},
		},
	}
}
