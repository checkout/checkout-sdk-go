package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/tokens"
)

func TestShouldRequestCardToken(t *testing.T) {
	response := RequestCardToken(t)
	assert.Equal(t, tokens.Card, response.Type)
	assert.NotEmpty(t, response.Token)
	assert.NotEmpty(t, response.ExpiresOn)
	assert.NotEmpty(t, response.ExpiryMonth)
	assert.NotEmpty(t, response.ExpiryYear)
	assert.NotEmpty(t, response.Scheme)
	assert.NotEmpty(t, response.Last4)
	assert.NotEmpty(t, response.IssuerCountry, common.GB)
	assert.NotEmpty(t, response.ProductType)
	assert.NotEmpty(t, response.Name)
}

func RequestCardToken(t *testing.T) *tokens.CardTokenResponse {
	request := tokens.CardTokenRequest{
		Type:        tokens.Card,
		Number:      CardNumber,
		ExpiryMonth: ExpiryMonth,
		ExpiryYear:  ExpiryYear,
		Name:        Name,
		CVV:         Cvv,
	}
	response, err := DefaultApi().Tokens.RequestCardToken(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Token)
	return response
}
