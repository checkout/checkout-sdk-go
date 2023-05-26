package test

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/issuing"
)

func TestCreateCard(t *testing.T) {
	cases := []struct {
		name    string
		checker func(*issuing.CardResponse, error)
	}{
		{
			name: "when create a card and this request is correct then should return a response",
			checker: func(response *issuing.CardResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.Equal(t, virtualCardResponse.DisplayName, response.DisplayName)
				assert.NotNil(t, response.LastFour)
				assert.NotNil(t, response.ExpiryMonth)
				assert.NotNil(t, response.ExpiryYear)
				assert.NotNil(t, response.BillingCurrency)
				assert.NotNil(t, response.IssuingCountry)
				assert.Equal(t, virtualCardResponse.Reference, response.Reference)
				assert.NotNil(t, response.CreatedDate)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(cardRequest(cardholderResponse), nil)
		})
	}
}

func TestGetCardDetails(t *testing.T) {
	cases := []struct {
		name    string
		cardId  string
		checker func(*issuing.CardDetailsResponse, error)
	}{
		{
			name:   "when get a card and this request is correct then should return a response",
			cardId: virtualCardId,
			checker: func(response *issuing.CardDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, issuing.Virtual, response.VirtualCardResponse.Type)
				assert.Equal(t, virtualCardResponse.Id, response.VirtualCardResponse.Id)
				assert.NotNil(t, response.VirtualCardResponse.CardholderId)
				assert.NotNil(t, response.VirtualCardResponse.CardProductId)
				assert.NotNil(t, response.VirtualCardResponse.ClientId)
				assert.Equal(t, virtualCardResponse.LastFour, response.VirtualCardResponse.LastFour)
				assert.Equal(t, virtualCardResponse.ExpiryMonth, response.VirtualCardResponse.ExpiryMonth)
				assert.Equal(t, virtualCardResponse.ExpiryYear, response.VirtualCardResponse.ExpiryYear)
				assert.NotNil(t, response.VirtualCardResponse.Status)
				assert.Equal(t, virtualCardResponse.DisplayName, response.VirtualCardResponse.DisplayName)
				assert.Equal(t, virtualCardResponse.BillingCurrency, response.VirtualCardResponse.BillingCurrency)
				assert.Equal(t, virtualCardResponse.IssuingCountry, response.VirtualCardResponse.IssuingCountry)
				assert.Equal(t, virtualCardResponse.Reference, response.VirtualCardResponse.Reference)
				assert.NotNil(t, response.VirtualCardResponse.CreatedDate)
				assert.NotNil(t, response.VirtualCardResponse.LastModifiedDate)
				assert.NotNil(t, response.VirtualCardResponse.IsSingleUse)
			},
		},
		{
			name:   "when get a card and this card not found then should return an error",
			cardId: "crd_not_found",
			checker: func(response *issuing.CardDetailsResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetCardDetails(tc.cardId))
		})
	}
}

func TestEnrollThreeDS(t *testing.T) {
	t.Skip("Client id must be configured for 3ds")
	request := issuing.ThreeDSEnrollmentRequest{
		Locale:      "en-US",
		PhoneNumber: Phone(),
	}

	cases := []struct {
		name    string
		cardId  string
		request issuing.ThreeDSEnrollmentRequest
		checker func(*issuing.ThreeDSEnrollmentResponse, error)
	}{
		{
			name:    "when enroll a card three DS and this request is correct then should return a response",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			checker: func(response *issuing.ThreeDSEnrollmentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 202, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.CreatedDate)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.EnrollThreeDS(tc.cardId, tc.request))
		})
	}
}

func TestUpdateThreeDS(t *testing.T) {
	t.Skip("Client id must be configured for 3ds")
	request := issuing.ThreeDSUpdateRequest{
		SecurityPair: issuing.SecurityPair{
			Question: "Who are you?",
			Answer:   "Bond. James Bond.",
		},
		Password:    "Xtui43FvfiZ",
		Locale:      "en-US",
		PhoneNumber: Phone(),
	}

	cases := []struct {
		name    string
		cardId  string
		request issuing.ThreeDSUpdateRequest
		checker func(*issuing.ThreeDSUpdateResponse, error)
	}{
		{
			name:    "when update a card enroll three DS and this request is correct then should return 201",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			checker: func(response *issuing.ThreeDSUpdateResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 202, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateThreeDS(tc.cardId, tc.request))
		})
	}
}

func TestGetCardThreeDSDetails(t *testing.T) {
	t.Skip("Client id must be configured for 3ds")
	cases := []struct {
		name    string
		cardId  string
		checker func(*issuing.ThreeDSEnrollmentDetailsResponse, error)
	}{
		{
			name:   "when get a card enroll three DS details and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			checker: func(response *issuing.ThreeDSEnrollmentDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetCardThreeDSDetails(tc.cardId))
		})
	}
}

func TestActivateCard(t *testing.T) {
	cases := []struct {
		name    string
		cardId  string
		checker func(*common.IdResponse, error)
	}{
		{
			name:   "when activate a card and this request is correct then should return a response",
			cardId: virtualCardId,
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Links)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client.SuspendCard(tc.cardId)
			tc.checker(client.ActivateCard(tc.cardId))
		})
	}
}

func TestGetCardCredentials(t *testing.T) {
	cases := []struct {
		name    string
		cardId  string
		query   issuing.CardCredentialsQuery
		checker func(*issuing.CardCredentialsResponse, error)
	}{
		{
			name:   "when get card credentials and this request is correct then should return a response",
			cardId: virtualCardId,
			query: issuing.CardCredentialsQuery{
				Credentials: "number, cvc2",
			},
			checker: func(response *issuing.CardCredentialsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Number)
				assert.NotNil(t, response.Cvc2)
			},
		},
		{
			name:   "when get card credentials and card not fund is correct then should return a response",
			cardId: "crd_not_found",
			query: issuing.CardCredentialsQuery{
				Credentials: "number, cvc2",
			},
			checker: func(response *issuing.CardCredentialsResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetCardCredentials(tc.cardId, tc.query))
		})
	}
}

func TestSuspendCard(t *testing.T) {
	cases := []struct {
		name    string
		cardId  string
		checker func(*common.IdResponse, error)
	}{
		{
			name:   "when suspend a card and this request is correct then should return a response",
			cardId: virtualCardId,
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Links)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cardDetails, _ := client.GetCardDetails(tc.cardId)
			if cardDetails.VirtualCardResponse.Status == issuing.CardActive {
				tc.checker(client.SuspendCard(tc.cardId))
			} else {
				client.ActivateCard(tc.cardId)
				tc.checker(client.SuspendCard(tc.cardId))
			}
		})
	}
}

func TestRevokeCard(t *testing.T) {
	request := issuing.RevokeCardRequest{}

	cases := []struct {
		name    string
		cardId  string
		request issuing.RevokeCardRequest
		checker func(*common.IdResponse, error)
	}{
		{
			name:    "when revoke a card and this request is correct then should return a response",
			cardId:  virtualCardId,
			request: request,
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Links)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RevokeCard(tc.cardId, tc.request))
		})
	}
}
