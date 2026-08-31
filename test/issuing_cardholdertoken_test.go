package test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/issuing/cardholdertokens"
)

func TestRequestCardholderTokenWithContext(t *testing.T) {
	api := buildIssuingClientApi()

	// Create a fresh cardholder to avoid cross-test coupling.
	created := cardholderRequest(t)
	assert.NotEmpty(t, created.Id, "cardholder must be created before requesting a token")

	cases := []struct {
		name    string
		request cardholdertokens.CardholderTokenRequest
		checker func(*cardholdertokens.CardholderTokenResponse, error)
	}{
		{
			name: "when request is correct then return cardholder token",
			request: cardholdertokens.CardholderTokenRequest{
				GrantType:    "client_credentials",
				ClientId:     os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_ID"),
				ClientSecret: os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_SECRET"),
				CardholderId: created.Id,
			},
			checker: func(response *cardholdertokens.CardholderTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.AccessToken)
				assert.NotEmpty(t, response.TokenType)
				assert.Greater(t, response.ExpiresIn, float64(0))
			},
		},
		{
			name: "when single_use is true then return single-use cardholder token",
			request: cardholdertokens.CardholderTokenRequest{
				GrantType:    "client_credentials",
				ClientId:     os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_ID"),
				ClientSecret: os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_SECRET"),
				CardholderId: created.Id,
				SingleUse:    true,
			},
			checker: func(response *cardholdertokens.CardholderTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotEmpty(t, response.AccessToken)
			},
		},
	}

	client := api.CardholderTokens

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestCardholderTokenWithContext(context.Background(), tc.request))
		})
	}
}
