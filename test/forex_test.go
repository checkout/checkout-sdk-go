package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/forex"
)

func TestRequestQuote(t *testing.T) {
	t.Skip("unavailable")
	cases := []struct {
		name    string
		request forex.QuoteRequest
		checker func(*forex.QuoteResponse, error)
	}{
		{
			name: "when request is correct then should request quote",
			request: forex.QuoteRequest{
				SourceCurrency:      common.GBP,
				SourceAmount:        30000,
				DestinationCurrency: common.USD,
				ProcessingChannelId: "pc_abcdefghijklmnopqrstuvwxyz",
			},
			checker: func(response *forex.QuoteResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Rate)
				assert.NotNil(t, response.ExpiresOn)
			},
		},
		{
			name: "when request is not correct then return error",
			request: forex.QuoteRequest{
				ProcessingChannelId: "pc_abcdefghijklmnopqrstuvwxyz",
			},
			checker: func(response *forex.QuoteResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "source_currency_required")
			},
		},
	}

	client := OAuthApi().Forex

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestQuote(tc.request))
		})
	}
}

func TestGetRates(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")
	cases := []struct {
		name    string
		request forex.RatesQuery
		checker func(*forex.RatesResponse, error)
	}{
		{
			name: "when request is correct then should request quote",
			request: forex.RatesQuery{
				Product:             "card_payouts",
				Source:              forex.Visa,
				CurrencyPairs:       "GBPEUR,USDNOK,JPNCAD",
				ProcessingChannelId: "pc_a6dabcfa2o3ejghb3sjuotdzzy",
			},
			checker: func(response *forex.RatesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Product)
				assert.NotNil(t, response.Source)
				assert.NotNil(t, response.Rates)
			},
		},
		{
			name:    "when request is not correct then return error",
			request: forex.RatesQuery{},
			checker: func(response *forex.RatesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "processing_error", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "product_required")
			},
		},
	}

	client := OAuthApi().Forex

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetRates(tc.request))
		})
	}
}
