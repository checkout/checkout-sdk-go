package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/instruments"
	"github.com/checkout/checkout-sdk-go/instruments/nas"
	"github.com/checkout/checkout-sdk-go/tokens"
)

func TestCreateAndGetInstrument(t *testing.T) {

	cardTokenResponse := RequestCardToken(t)
	response := createTokenInstrument(t, cardTokenResponse)
	assert.Equal(t, instruments.Card, response.Type)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.Fingerprint)
	assert.NotEmpty(t, response.ExpiryMonth)
	assert.NotEmpty(t, response.ExpiryYear)
	assert.NotEmpty(t, response.Scheme)
	assert.NotEmpty(t, response.Last4)
	assert.NotEmpty(t, response.Bin)
	assert.NotEmpty(t, response.CardType)
	assert.NotEmpty(t, response.CardCategory)
	assert.NotEmpty(t, response.ProductId)
	assert.NotEmpty(t, response.ProductType)

	getResponse, err := DefaultApi().Instruments.Get(response.Id)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	getCardInstrumentResponse := getResponse.GetCardInstrumentResponse
	assert.Equal(t, instruments.Card, getCardInstrumentResponse.Type)
	assert.NotEmpty(t, getCardInstrumentResponse.Id)
	assert.NotEmpty(t, getCardInstrumentResponse.Fingerprint)
	assert.NotEmpty(t, getCardInstrumentResponse.AccountHolder)
	assert.NotEmpty(t, getCardInstrumentResponse.ExpiryMonth)
	assert.NotEmpty(t, getCardInstrumentResponse.ExpiryYear)
	assert.NotEmpty(t, getCardInstrumentResponse.Name)
	assert.NotEmpty(t, getCardInstrumentResponse.Scheme)
	assert.NotEmpty(t, getCardInstrumentResponse.Last4)
	assert.NotEmpty(t, getCardInstrumentResponse.Bin)
	assert.NotEmpty(t, getCardInstrumentResponse.CardType)
	assert.NotEmpty(t, getCardInstrumentResponse.CardCategory)
	assert.NotEmpty(t, getCardInstrumentResponse.ProductId)
	assert.NotEmpty(t, getCardInstrumentResponse.ProductType)
}

func TestShouldGetInstrument(t *testing.T) {

	token := RequestCardToken(t)
	createInstrumentResponse := createTokenInstrument(t, token)

	response, err := DefaultApi().Instruments.Get(createInstrumentResponse.Id)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	getCardInstrumentResponse := response.GetCardInstrumentResponse
	assert.Equal(t, instruments.Card, getCardInstrumentResponse.Type)
	assert.NotEmpty(t, getCardInstrumentResponse.Id)
	assert.NotEmpty(t, getCardInstrumentResponse.Fingerprint)
	assert.NotEmpty(t, getCardInstrumentResponse.AccountHolder)
	assert.NotEmpty(t, getCardInstrumentResponse.ExpiryMonth)
	assert.NotEmpty(t, getCardInstrumentResponse.ExpiryYear)
	assert.NotEmpty(t, getCardInstrumentResponse.Name)
	assert.NotEmpty(t, getCardInstrumentResponse.Scheme)
	assert.NotEmpty(t, getCardInstrumentResponse.Last4)
	assert.NotEmpty(t, getCardInstrumentResponse.Bin)
	assert.NotEmpty(t, getCardInstrumentResponse.CardType)
	assert.NotEmpty(t, getCardInstrumentResponse.CardCategory)
	assert.NotEmpty(t, getCardInstrumentResponse.ProductId)
	assert.NotEmpty(t, getCardInstrumentResponse.ProductType)
}

func TestShouldGetBankAccountFields(t *testing.T) {

	response, err := OAuthApi().Instruments.GetBankAccountFieldFormatting("GB", "GBP", nas.QueryBankAccountFormatting{
		AccountHolderType: common.Individual,
		PaymentNetwork:    nas.Ach,
	})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	getBankAccountFieldsResponse := response.Sections[0]
	assert.Equal(t, "Account holder details", getBankAccountFieldsResponse.Name)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].Id)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].Section)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].Display)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].HelpText)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].Type)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].Required)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].ValidationRegex)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].MinLength)
	assert.NotNil(t, getBankAccountFieldsResponse.Fields[0].MaxLength)
}

func TestShouldUpdateInstrument(t *testing.T) {

	token := RequestCardToken(t)
	createInstrumentResponse := createTokenInstrument(t, token)

	cardTokenResponse := RequestCardToken(t)

	updateTokenInstrumentRequest := nas.NewUpdateTokenInstrumentRequest()
	updateTokenInstrumentRequest.Token = cardTokenResponse.Token

	response, err := DefaultApi().Instruments.Update(createInstrumentResponse.Id, updateTokenInstrumentRequest)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.UpdateCardInstrumentResponse)
	assert.NotEmpty(t, response.UpdateCardInstrumentResponse.Fingerprint)
}

func TestShouldDeleteInstrument(t *testing.T) {

	token := RequestCardToken(t)
	createInstrumentResponse := createTokenInstrument(t, token)

	response, err := DefaultApi().Instruments.Delete(createInstrumentResponse.Id)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.HttpMetadata.StatusCode, 204)

	getResponse, err := DefaultApi().Instruments.Get(createInstrumentResponse.Id)
	assert.Nil(t, getResponse)
	assert.NotNil(t, err)
	chkErr := err.(errors.CheckoutAPIError)
	assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
}

func createTokenInstrument(t *testing.T, token *tokens.CardTokenResponse) *nas.CreateTokenInstrumentResponse {
	request := nas.NewCreateTokenInstrumentRequest()
	request.Token = token.Token
	request.AccountHolder = &common.AccountHolder{
		FirstName:      FirstName,
		LastName:       LastName,
		BillingAddress: Address(),
		Phone:          Phone(),
	}

	response, err := DefaultApi().Instruments.Create(request)
	assert.Nil(t, err)
	assert.NotNil(t, response.CreateTokenInstrumentResponse)
	return response.CreateTokenInstrumentResponse
}
