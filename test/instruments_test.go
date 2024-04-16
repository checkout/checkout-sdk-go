package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/instruments/nas"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/tokens"
)

var (
	instrumentToken *nas.CreateTokenInstrumentResponse
)

func TestSetupInstrument(t *testing.T) {
	cardTokenResponse := RequestCardToken(t)
	_ = createSepaInstrument(t)
	instrumentToken = createTokenInstrument(t, cardTokenResponse)
}

func TestCreateAndGetInstrument(t *testing.T) {
	cases := []struct {
		name       string
		responseId string
		checker    func(*nas.GetInstrumentResponse, error)
	}{
		{
			name:       "when get a created instrument then return it",
			responseId: instrumentToken.Id,
			checker: func(response *nas.GetInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				getCardInstrumentResponse := response.GetCardInstrumentResponse
				assert.Equal(t, common.Card, getCardInstrumentResponse.Type)
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
			},
		},
	}

	client := DefaultApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Get(instrumentToken.Id))
		})
	}
}

func TestShouldGetInstrument(t *testing.T) {
	cases := []struct {
		name       string
		responseId string
		checker    func(*nas.GetInstrumentResponse, error)
	}{
		{
			name:       "when fetching a valid instrument then return instrument data",
			responseId: instrumentToken.Id,
			checker: func(response *nas.GetInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				getCardInstrumentResponse := response.GetCardInstrumentResponse
				assert.Equal(t, common.Card, getCardInstrumentResponse.Type)
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
			},
		},
	}

	client := DefaultApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Get(tc.responseId))
		})
	}
}

func TestShouldGetBankAccountFields(t *testing.T) {
	cases := []struct {
		name     string
		country  common.Country
		currency common.Currency
		query    nas.QueryBankAccountFormatting
		checker  func(*nas.GetBankAccountFieldFormattingResponse, error)
	}{
		{
			name:     "when get an instrument by query account formatting then return instrument data",
			country:  common.GB,
			currency: common.GBP,
			query: nas.QueryBankAccountFormatting{
				AccountHolderType: common.Individual,
				PaymentNetwork:    nas.Ach,
			},
			checker: func(response *nas.GetBankAccountFieldFormattingResponse, err error) {
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
			},
		},
	}

	client := OAuthApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetBankAccountFieldFormatting(string(tc.country), string(tc.currency), tc.query))
		})
	}
}

func TestShouldUpdateInstrument(t *testing.T) {
	cardTokenResponse := RequestCardToken(t)
	updateTokenInstrumentRequest := nas.NewUpdateTokenInstrumentRequest()
	updateTokenInstrumentRequest.Token = cardTokenResponse.Token

	cases := []struct {
		name         string
		instrumentId string
		request      nas.UpdateInstrumentRequest
		checker      func(*nas.UpdateInstrumentResponse, error)
	}{
		{
			name:         "when update an instrument then return instrument data",
			instrumentId: instrumentToken.Id,
			request:      updateTokenInstrumentRequest,
			checker: func(response *nas.UpdateInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.UpdateCardInstrumentResponse)
				assert.NotEmpty(t, response.UpdateCardInstrumentResponse.Fingerprint)
			},
		},
	}

	client := DefaultApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Update(tc.instrumentId, tc.request))
		})
	}
}

func TestShouldDeleteInstrument(t *testing.T) {
	cases := []struct {
		name         string
		instrumentId string
		checkerOne   func(*common.MetadataResponse, error)
		checkerTwo   func(*nas.GetInstrumentResponse, error)
	}{
		{
			name:         "when delete an instrument then return 204",
			instrumentId: instrumentToken.Id,
			checkerOne: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.HttpMetadata.StatusCode, 204)
			},
			checkerTwo: func(response *nas.GetInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := DefaultApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkerOne(client.Delete(tc.instrumentId))
			tc.checkerTwo(client.Get(tc.instrumentId))
		})
	}
}

func createSepaInstrument(t *testing.T) *nas.CreateSepaInstrumentResponse {
	request := nas.NewCreateSepaInstrumentRequest()
	request.InstrumentData = &nas.InstrumentData{
		AccountNumber: "FR7630006000011234567890189",
		Country:       common.FR,
		Currency:      common.EUR,
		PaymentType:   payments.Recurring,
	}
	request.AccountHolder = &common.AccountHolder{
		FirstName:      "Ali",
		LastName:       "Farid",
		BillingAddress: Address(),
		Phone:          Phone(),
	}

	response, err := DefaultApi().Instruments.Create(request)
	assert.Nil(t, err)
	assert.NotNil(t, response.CreateSepaInstrumentResponse)
	assert.Equal(t, common.Sepa, response.CreateSepaInstrumentResponse.Type)
	assert.NotEmpty(t, response.CreateSepaInstrumentResponse.Id)
	assert.NotEmpty(t, response.CreateSepaInstrumentResponse.Fingerprint)
	return response.CreateSepaInstrumentResponse
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
	assert.Equal(t, common.Card, response.CreateTokenInstrumentResponse.Type)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.Id)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.Fingerprint)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.ExpiryMonth)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.ExpiryYear)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.Scheme)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.Last4)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.Bin)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.CardType)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.CardCategory)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.ProductId)
	assert.NotEmpty(t, response.CreateTokenInstrumentResponse.ProductType)
	return response.CreateTokenInstrumentResponse
}
