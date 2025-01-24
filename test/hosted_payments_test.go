package test

import (
	"fmt"
	"net/http"
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
				assert.Equal(t, Reference, response.Reference)
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
		Amount:      1000,
		Currency:    common.GBP,
		PaymentType: payments.Regular,
		BillingDescriptor: &payments.BillingDescriptor{
			Name: Name,
			City: "London",
		},
		DisplayName: "**Test Hosted Payment**",
		Reference:   Reference,
		Description: Description,
		Customer: &common.CustomerRequest{
			Email: GenerateRandomEmail(),
			Name:  Name,
			Phone: Phone(),
		},
		Shipping: &payments.ShippingDetails{
			Address: Address(),
			Phone:   Phone(),
		},
		Billing: &payments.BillingInformation{
			Address: Address(),
			Phone:   Phone(),
		},
		Recipient: &payments.PaymentRecipient{
			DateOfBirth:   "1985-05-15",
			AccountNumber: "1234567",
			CountryCode:   common.GB,
			Zip:           "12345",
			FirstName:     FirstName,
			LastName:      LastName,
		},
		Processing:          &payments.ProcessingSettings{Aft: true},
		AllowPaymentMethods: []payments.SourceType{payments.CardSource, payments.IdealSource},
		Products: []payments.Product{
			{
				Name:     "Gold Necklace",
				Quantity: 1,
				Price:    1000,
			},
		},
		Risk:       &payments.RiskRequest{Enabled: false},
		SuccessUrl: "https://example.com/payments/success",
		CancelUrl:  "https://example.com/payments/cancel",
		FailureUrl: "https://example.com/payments/failure",
		Locale:     "en-GB",
		ThreeDs: &payments.ThreeDsRequest{
			Enabled:            false,
			AttemptN3D:         false,
			ChallengeIndicator: common.NoChallengeRequested,
		},
		Capture:   true,
		CaptureOn: &now,
	}
}
