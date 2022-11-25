package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/instruments"
	"github.com/checkout/checkout-sdk-go/instruments/abc"
	"github.com/checkout/checkout-sdk-go/tokens"
)

func TestCreateAndGetInstrumentPrevious(t *testing.T) {

	cardTokenResponse := RequestCardTokenPrevious(t)

	createResponse := createTokenInstrumentPrevious(t, cardTokenResponse)
	assert.Equal(t, instruments.Card, createResponse.Type)
	assert.NotEmpty(t, createResponse.Id)
	assert.NotEmpty(t, createResponse.Fingerprint)
	assert.NotEmpty(t, createResponse.ExpiryMonth)
	assert.NotEmpty(t, createResponse.ExpiryYear)
	assert.NotEmpty(t, createResponse.Scheme)
	assert.NotEmpty(t, createResponse.Last4)
	assert.NotEmpty(t, createResponse.Bin)
	assert.NotEmpty(t, createResponse.CardType)
	assert.NotEmpty(t, createResponse.CardCategory)
	assert.NotEmpty(t, createResponse.ProductId)
	assert.NotEmpty(t, createResponse.ProductType)
	assert.NotEmpty(t, createResponse.Customer)
	assert.NotEmpty(t, createResponse.Customer.Id)
	assert.NotEmpty(t, createResponse.Customer.Name)
	assert.NotEmpty(t, createResponse.Customer.Email)

	getResponse, err := PreviousApi().Instruments.Get(createResponse.Id)
	assert.Nil(t, err)
	assert.NotNil(t, getResponse)
	assert.Equal(t, instruments.Card, getResponse.Type)
	assert.NotEmpty(t, getResponse.Id)
	assert.NotEmpty(t, getResponse.Fingerprint)
	assert.NotEmpty(t, getResponse.ExpiryMonth)
	assert.NotEmpty(t, getResponse.ExpiryYear)
	assert.NotEmpty(t, getResponse.Scheme)
	assert.NotEmpty(t, getResponse.Last4)
	assert.NotEmpty(t, getResponse.Bin)
	assert.NotEmpty(t, getResponse.CardType)
	assert.NotEmpty(t, getResponse.CardCategory)
	assert.NotEmpty(t, getResponse.ProductId)
	assert.NotEmpty(t, getResponse.ProductType)
	assert.NotEmpty(t, getResponse.Customer)
	assert.NotEmpty(t, getResponse.Customer.Id)
	assert.NotEmpty(t, getResponse.Customer.Name)
	assert.NotEmpty(t, getResponse.Customer.Email)
}

func TestCreateAndUpdateInstrumentPrevious(t *testing.T) {

	cardTokenResponse := RequestCardTokenPrevious(t)
	createResponse := createTokenInstrumentPrevious(t, cardTokenResponse)

	updateRequest := abc.UpdateInstrumentRequest{
		ExpiryMonth: 12,
		ExpiryYear:  2026,
		Name:        "New Name",
	}

	updateResponse, err := PreviousApi().Instruments.Update(createResponse.Id, updateRequest)
	assert.Nil(t, err)
	assert.NotNil(t, updateResponse)
	assert.Equal(t, instruments.Card, updateResponse.Type)
	assert.NotEmpty(t, updateResponse.Fingerprint)

	getResponse, err := PreviousApi().Instruments.Get(createResponse.Id)
	assert.Nil(t, err)
	assert.NotNil(t, getResponse)
	assert.Equal(t, instruments.Card, getResponse.Type)
	assert.Equal(t, 12, getResponse.ExpiryMonth)
	assert.Equal(t, 2026, getResponse.ExpiryYear)
	assert.Equal(t, "New Name", getResponse.Name)
}

func TestCreateAndDeleteInstrumentPrevious(t *testing.T) {

	cardTokenResponse := RequestCardTokenPrevious(t)
	createResponse := createTokenInstrumentPrevious(t, cardTokenResponse)

	deleteResponse, err := PreviousApi().Instruments.Delete(createResponse.Id)
	assert.Nil(t, err)
	assert.NotNil(t, deleteResponse)
	assert.Equal(t, 204, deleteResponse.HttpMetadata.StatusCode)

	getResponse, err := PreviousApi().Instruments.Get(createResponse.Id)
	assert.Nil(t, getResponse)
	assert.NotNil(t, err)
	chkErr := err.(errors.CheckoutAPIError)
	assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
}

func createTokenInstrumentPrevious(t *testing.T, token *tokens.CardTokenResponse) *abc.CreateInstrumentResponse {
	request := abc.CreateInstrumentRequest{
		Type:  instruments.Token,
		Token: token.Token,
		Customer: &abc.InstrumentCustomerRequest{
			Email:     Email,
			Name:      Name,
			Phone:     Phone(),
			IsDefault: true,
		},
	}
	response, err := PreviousApi().Instruments.Create(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	return response
}
