package test

import (
	"github.com/checkout/checkout-sdk-go-beta/payments/nas"
	"github.com/checkout/checkout-sdk-go-beta/payments/nas/sources"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/disputes"
	"github.com/checkout/checkout-sdk-go-beta/payments"
)

func TestGetDisputesWithOAuthSdk(t *testing.T) {
	cases := []struct {
		name      string
		disputeId string
		checker   func(*disputes.DisputeResponse, error)
	}{
		{
			name:      "when dispute exists then return dispute details",
			disputeId: "dsp_35c0fdfabe770k17440a",
			checker: func(response *disputes.DisputeResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, "dsp_35c0fdfabe770k17440a", response.Id)
			},
		},
		{
			name:      "when dispute exists then return dispute details",
			disputeId: "dsp_35c0fdfabe770k17440a",
			checker: func(response *disputes.DisputeResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, "dsp_35c0fdfabe770k17440a", response.Id)
			},
		},
	}

	client := OAuthApi().Disputes

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetDisputeDetails(tc.disputeId))
		})
	}
}

func TestRequestPaymentWithOAuthSdk(t *testing.T) {
	cases := []struct {
		name    string
		request nas.PaymentRequest
		checker func(*nas.PaymentResponse, error)
	}{
		{
			name: "test Card source for request payment - OAuth",
			request: nas.PaymentRequest{
				Source:      getCardSourceRequest(),
				Amount:      10,
				Currency:    common.GBP,
				Reference:   Reference,
				Description: Description,
				Capture:     false,
				Customer: &common.CustomerRequest{
					Email: Email,
					Name:  Name,
					Phone: Phone(),
				},
				Sender: getPaymentIndividualSender(),
				BillingDescriptor: &payments.BillingDescriptor{
					Name:      Name,
					City:      City,
					Reference: Reference,
				},
				ProcessingChannelId: "pc_a6dabcfa2o3ejghb3sjuotdzzy",
				Marketplace: &common.MarketplaceData{
					SubEntityId: "ent_ocw5i74vowfg2edpy66izhts2u",
				},
			},
			checker: func(response *nas.PaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := OAuthApi().Payments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestPayment(tc.request, nil))
		})
	}

}

func getCardSourceRequest() *sources.RequestCardSource {
	cardSource := sources.NewRequestCardSource()
	cardSource.Name = Name
	cardSource.Number = CardNumber
	cardSource.ExpiryYear = ExpiryYear
	cardSource.ExpiryMonth = ExpiryMonth
	cardSource.Cvv = Cvv
	cardSource.BillingAddress = Address()
	cardSource.Phone = Phone()

	return cardSource
}

func getPaymentIndividualSender() *nas.PaymentIndividualSender {
	paymentIndividualSender := nas.NewPaymentIndividualSender()
	paymentIndividualSender.FirstName = FirstName
	paymentIndividualSender.LastName = LastName
	paymentIndividualSender.Address = Address()
	paymentIndividualSender.Identification = &nas.Identification{
		Type:           nas.DrivingLicence,
		Number:         "12345",
		IssuingCountry: common.GT,
	}

	return paymentIndividualSender
}
