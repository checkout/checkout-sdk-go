package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"

	cards "github.com/checkout/checkout-sdk-go/issuing/cards"
)

func TestCreateCard(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		checker func(*cards.CardResponse, error)
	}{
		{
			name: "when create a card and this request is correct then should return a response",
			checker: func(response *cards.CardResponse, err error) {
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
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		cardId  string
		checker func(*cards.CardDetailsResponse, error)
	}{
		{
			name:   "when get a card and this request is correct then should return a response",
			cardId: virtualCardId,
			checker: func(response *cards.CardDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, cards.Virtual, response.Type)
				assert.Equal(t, virtualCardResponse.Id, response.Id)
				assert.NotNil(t, response.CardholderId)
				assert.NotNil(t, response.CardProductId)
				assert.NotNil(t, response.ClientId)
				assert.Equal(t, virtualCardResponse.LastFour, response.LastFour)
				assert.Equal(t, virtualCardResponse.ExpiryMonth, response.ExpiryMonth)
				assert.Equal(t, virtualCardResponse.ExpiryYear, response.ExpiryYear)
				assert.NotNil(t, response.Status)
				assert.Equal(t, virtualCardResponse.DisplayName, response.DisplayName)
				assert.Equal(t, virtualCardResponse.BillingCurrency, response.BillingCurrency)
				assert.Equal(t, virtualCardResponse.IssuingCountry, response.IssuingCountry)
				assert.Equal(t, virtualCardResponse.Reference, response.Reference)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
				assert.NotNil(t, response.ExtraData.(*cards.VirtualExtraData).IsSingleUse)
			},
		},
		{
			name:   "when get a card and this card not found then should return an error",
			cardId: "crd_not_found",
			checker: func(response *cards.CardDetailsResponse, err error) {
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
	request := cards.ThreeDSEnrollmentRequest{
		Locale:      "en-US",
		PhoneNumber: Phone(),
	}

	cases := []struct {
		name    string
		cardId  string
		request cards.ThreeDSEnrollmentRequest
		checker func(*cards.ThreeDSEnrollmentResponse, error)
	}{
		{
			name:    "when enroll a card three DS and this request is correct then should return a response",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			checker: func(response *cards.ThreeDSEnrollmentResponse, err error) {
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
	request := cards.ThreeDSUpdateRequest{
		SecurityPair: cards.SecurityPair{
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
		request cards.ThreeDSUpdateRequest
		checker func(*cards.ThreeDSUpdateResponse, error)
	}{
		{
			name:    "when update a card enroll three DS and this request is correct then should return 201",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			checker: func(response *cards.ThreeDSUpdateResponse, err error) {
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
		checker func(*cards.ThreeDSEnrollmentDetailsResponse, error)
	}{
		{
			name:   "when get a card enroll three DS details and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			checker: func(response *cards.ThreeDSEnrollmentDetailsResponse, err error) {
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
	t.Skip("Avoid creating cards all the time")
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
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		cardId  string
		query   cards.CardCredentialsQuery
		checker func(*cards.CardCredentialsResponse, error)
	}{
		{
			name:   "when get card credentials and this request is correct then should return a response",
			cardId: virtualCardId,
			query: cards.CardCredentialsQuery{
				Credentials: "number, cvc2",
			},
			checker: func(response *cards.CardCredentialsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Number)
				assert.NotNil(t, response.Cvc2)
			},
		},
		{
			name:   "when get card credentials and card not fund is correct then should return a response",
			cardId: "crd_not_found",
			query: cards.CardCredentialsQuery{
				Credentials: "number, cvc2",
			},
			checker: func(response *cards.CardCredentialsResponse, err error) {
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
	t.Skip("Avoid creating cards all the time")
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
			if cardDetails.Status == cards.CardActive {
				tc.checker(client.SuspendCard(tc.cardId))
			} else {
				client.ActivateCard(tc.cardId)
				tc.checker(client.SuspendCard(tc.cardId))
			}
		})
	}
}

func TestRevokeCard(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	request := cards.RevokeCardRequest{}

	cases := []struct {
		name    string
		cardId  string
		request cards.RevokeCardRequest
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
