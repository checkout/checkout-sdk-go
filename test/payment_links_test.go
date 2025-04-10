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
	"github.com/checkout/checkout-sdk-go/payments/links"
)

func TestCreatePaymentLink(t *testing.T) {
	cases := []struct {
		name    string
		request links.PaymentLinkRequest
		checker func(*links.PaymentLinkResponse, error)
	}{
		{
			name:    "when request is valid then create payment link",
			request: *getPaymentLinkRequest(),
			checker: func(response *links.PaymentLinkResponse, err error) {
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
			request: links.PaymentLinkRequest{},
			checker: func(response *links.PaymentLinkResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().Links

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreatePaymentLink(tc.request))
		})
	}
}

func TestGetPaymentLink(t *testing.T) {
	paymentLink, err := DefaultApi().Links.CreatePaymentLink(*getPaymentLinkRequest())
	if err != nil {
		assert.Fail(t, fmt.Sprintf("Error creating payment link: %s", err.Error()))
	}

	cases := []struct {
		name             string
		paymentSessionId string
		checker          func(*links.PaymentLinkDetails, error)
	}{
		{
			name:             "when fetching existing payment link then return link details",
			paymentSessionId: paymentLink.Id,
			checker: func(response *links.PaymentLinkDetails, err error) {
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
				assert.Equal(t, links.Active, response.Status)
			},
		},
		{
			name:             "when paymentSessionId is inexistent then return error",
			paymentSessionId: "not_found",
			checker: func(response *links.PaymentLinkDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().Links

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetPaymentLink(tc.paymentSessionId))
		})
	}
}

func getPaymentLinkRequest() *links.PaymentLinkRequest {
	now = time.Now()

	return &links.PaymentLinkRequest{
		Amount:      200,
		Currency:    common.GBP,
		PaymentType: payments.Regular,
		BillingDescriptor: &payments.BillingDescriptor{
			Name: Name,
			City: "London",
		},
		Reference:   Reference,
		Description: Description,
		ExpiresIn:   604800,
		Customer: &common.CustomerRequest{
			Email: GenerateRandomEmail(),
			Name:  Name,
			Phone: Phone(),
		},
		Shipping: &payments.ShippingDetailsFlowHostedLinks{
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
		Processing: &payments.ProcessingSettings{Aft: true},
		Capture:    true,
		CaptureOn:  &now,
		Products: []payments.Product{
			{
				Name:     "Gold Necklace",
				Quantity: 1,
				Price:    200,
			},
		},
		ThreeDs: &payments.ThreeDsRequestFlowHostedLinks{
			Enabled:            false,
			AttemptN3D:         false,
			ChallengeIndicator: common.NoChallengeRequested,
			Exemption:          payments.LowValue,
			AllowUpgrade:       true,
		},
		Risk:                &payments.RiskRequest{Enabled: false},
		ReturnUrl:           "https://example.com/payments/success",
		Locale:              "en-GB",
		AllowPaymentMethods: []payments.SourceType{payments.CardSource, payments.IdealSource},
	}
}
