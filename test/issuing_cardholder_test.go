package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/errors"

	cardholders "github.com/checkout/checkout-sdk-go/issuing/cardholders"
)

func TestCreateCardholder(t *testing.T) {
	t.Skip("Avoid creating cards all the time")

	cases := []struct {
		name    string
		request cardholders.CardholderRequest
		checker func(*cardholders.CardholderResponse, error)
	}{
		{
			name:    "when create a cardholder then return it",
			request: cardholder,
			checker: func(response *cardholders.CardholderResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.HttpMetadata.StatusCode, http.StatusCreated)
				assert.Equal(t, cardholder.Type, response.Type)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.Reference)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
			},
		},
		{
			name:    "when request don't have the then return error",
			request: cardholders.CardholderRequest{},
			checker: func(response *cardholders.CardholderResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "cardholder_type_required")
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateCardholder(tc.request))
		})
	}
}

func TestGetCardholderDetails(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name         string
		cardholderId string
		checker      func(*cardholders.CardholderDetailsResponse, error)
	}{
		{
			name:         "when request a cardholder then return it",
			cardholderId: cardholderResponse.Id,
			checker: func(response *cardholders.CardholderDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.HttpMetadata.StatusCode, http.StatusOK)
				assert.NotNil(t, response.Id)
				assert.Equal(t, cardholder.Type, response.Type)
				assert.Equal(t, cardholder.FirstName, response.FirstName)
				assert.Equal(t, cardholder.MiddleName, response.MiddleName)
				assert.Equal(t, cardholder.LastName, response.LastName)
				assert.Equal(t, cardholder.Email, response.Email)
				assert.NotNil(t, response.PhoneNumber)
				assert.Equal(t, cardholder.DateOfBirth, response.DateOfBirth)
				assert.Equal(t, cardholder.BillingAddress, response.BillingAddress)
				assert.Equal(t, cardholder.ResidencyAddress, response.ResidencyAddress)
				assert.Equal(t, cardholder.Reference, response.Reference)
				assert.Equal(t, "ent_mujh2nia2ypezmw5fo2fofk7ka", response.AccountEntityId)
				assert.Equal(t, "", response.ParentSubEntityId)
				assert.Equal(t, "ent_mujh2nia2ypezmw5fo2fofk7ka", response.EntityId)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
			},
		},
		{
			name:         "when cardholder id not found the then return error",
			cardholderId: "crd_not_found",
			checker: func(response *cardholders.CardholderDetailsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetCardholder(tc.cardholderId))
		})
	}
}

func TestGetCardholderCards(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name         string
		cardholderId string
		checker      func(*cardholders.CardholderCardsResponse, error)
	}{
		{
			name:         "when request all cards from a cardholder then return it",
			cardholderId: cardholderResponse.Id,
			checker: func(response *cardholders.CardholderCardsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.HttpMetadata.StatusCode, http.StatusOK)
				/*for _, card := range response.Cards {
					assert.Equal(t, "cli_p6jeowdtuxku3azxgt2qa7kq7a", card.VirtualCardResponse.ClientId)
				}*/
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetCardholderCards(tc.cardholderId))
		})
	}
}
