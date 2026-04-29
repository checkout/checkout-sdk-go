package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/standaloneaccountupdater"
)

// tests

func TestGetUpdatedCardCredentials(t *testing.T) {
	t.Skip("use on demand")
	cases := []struct {
		name    string
		request standaloneaccountupdater.GetUpdatedCardCredentialsRequest
		checker func(*standaloneaccountupdater.GetUpdatedCardCredentialsResponse, error)
	}{
		{
			name:    "when source is a valid card then return updated card credentials",
			request: buildGetUpdatedCardCredentialsWithCardIntegration(),
			checker: func(response *standaloneaccountupdater.GetUpdatedCardCredentialsResponse, err error) {
				assert.Nil(t, err)
				assertUpdatedCardCredentialsResponse(t, response)
			},
		},
		{
			name:    "when source is a valid instrument then return updated card credentials",
			request: buildGetUpdatedCardCredentialsWithInstrument(),
			checker: func(response *standaloneaccountupdater.GetUpdatedCardCredentialsResponse, err error) {
				assert.Nil(t, err)
				assertUpdatedCardCredentialsResponse(t, response)
			},
		},
	}

	client := OAuthApi().StandaloneAccountUpdater

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetUpdatedCardCredentials(tc.request))
		})
	}
}

// common methods

func buildGetUpdatedCardCredentialsWithCardIntegration() standaloneaccountupdater.GetUpdatedCardCredentialsRequest {
	return standaloneaccountupdater.GetUpdatedCardCredentialsRequest{
		SourceOptions: standaloneaccountupdater.AccountUpdaterSourceOptions{
			Card: &standaloneaccountupdater.AccountUpdaterCard{
				Number:      "4242424242424242",
				ExpiryMonth: 12,
				ExpiryYear:  2025,
			},
		},
	}
}

func buildGetUpdatedCardCredentialsWithInstrument() standaloneaccountupdater.GetUpdatedCardCredentialsRequest {
	return standaloneaccountupdater.GetUpdatedCardCredentialsRequest{
		SourceOptions: standaloneaccountupdater.AccountUpdaterSourceOptions{
			Instrument: &standaloneaccountupdater.AccountUpdaterInstrument{
				Id: "src_wmlfc3zyhqzehihu7giusaaawu",
			},
		},
	}
}

func assertUpdatedCardCredentialsResponse(t *testing.T, response *standaloneaccountupdater.GetUpdatedCardCredentialsResponse) {
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.AccountUpdateStatus)
	if response.AccountUpdateStatus == standaloneaccountupdater.CardUpdated ||
		response.AccountUpdateStatus == standaloneaccountupdater.CardExpiryUpdated {
		assert.NotNil(t, response.Card)
		assert.NotZero(t, response.Card.ExpiryMonth)
		assert.NotZero(t, response.Card.ExpiryYear)
	}
	if response.AccountUpdateStatus == standaloneaccountupdater.UpdateFailed {
		assert.NotEmpty(t, response.AccountUpdateFailureCode)
	}
}
