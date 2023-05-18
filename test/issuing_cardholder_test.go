package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/issuing"
)

var (
	request = issuing.CardholderRequest{
		Type:             issuing.Individual,
		Reference:        "X-123456-N11",
		EntityId:         "ent_mujh2nia2ypezmw5fo2fofk7ka",
		FirstName:        "John",
		MiddleName:       "Fitzgerald",
		LastName:         "Kennedy",
		Email:            "john.kennedy@myemaildomain.com",
		PhoneNumber:      Phone(),
		DateOfBirth:      "1985-05-15",
		BillingAddress:   Address(),
		ResidencyAddress: Address(),
		Document: &issuing.CardholderDocument{
			Type:            "national_identity_card",
			FrontDocumentId: "file_6lbss42ezvoufcb2beo76rvwly",
			BackDocumentId:  "file_aaz5pemp6326zbuvevp6qroqu4",
		},
	}
	cardholderResponse = cardholderRequest(request)
)

func TestCreateCardholder(t *testing.T) {
	cases := []struct {
		name    string
		request issuing.CardholderRequest
		checker func(*issuing.CardholderResponse, error)
	}{
		{
			name:    "when create a cardholder then return it",
			request: request,
			checker: func(response *issuing.CardholderResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.HttpMetadata.StatusCode, http.StatusCreated)
				assert.Equal(t, request.Type, response.Type)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.Reference)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
			},
		},
		{
			name:    "when request don't have the then return error",
			request: issuing.CardholderRequest{},
			checker: func(response *issuing.CardholderResponse, err error) {
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
	cases := []struct {
		name         string
		cardholderId string
		checker      func(*issuing.CardholderDetailsResponse, error)
	}{
		{
			name:         "when request a cardholder then return it",
			cardholderId: cardholderResponse.Id,
			checker: func(response *issuing.CardholderDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.HttpMetadata.StatusCode, http.StatusOK)
				assert.NotNil(t, response.Id)
				assert.Equal(t, request.Type, response.Type)
				assert.Equal(t, request.FirstName, response.FirstName)
				assert.Equal(t, request.MiddleName, response.MiddleName)
				assert.Equal(t, request.LastName, response.LastName)
				assert.Equal(t, request.Email, response.Email)
				assert.NotNil(t, response.PhoneNumber)
				assert.Equal(t, request.DateOfBirth, response.DateOfBirth)
				assert.Equal(t, request.BillingAddress, response.BillingAddress)
				assert.Equal(t, request.ResidencyAddress, response.ResidencyAddress)
				assert.Equal(t, request.Reference, response.Reference)
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
			checker: func(response *issuing.CardholderDetailsResponse, err error) {
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
	cases := []struct {
		name         string
		cardholderId string
		checker      func(*issuing.CardholderCardsResponse, error)
	}{
		{
			name:         "when request all cards from a cardholder then return it",
			cardholderId: cardholderResponse.Id,
			checker: func(response *issuing.CardholderCardsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.HttpMetadata.StatusCode, http.StatusOK)
				for _, card := range response.Cards {
					assert.Equal(t, "cli_p6jeowdtuxku3azxgt2qa7kq7a", card.GetDetails().(issuing.CardDetailsCardholder).ClientId)
				}
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

func cardholderRequest(request issuing.CardholderRequest) *issuing.CardholderResponse {
	response, _ := buildIssuingClientApi().Issuing.CreateCardholder(request)
	return response
}
