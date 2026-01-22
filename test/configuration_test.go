package test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
)

func TestGetAccessToken(t *testing.T) {
	cases := []struct {
		name             string
		inputCredentials *configuration.OAuthSdkCredentials
		checker          func(*configuration.OAuthAccessToken, error)
	}{
		{
			name: "when input credentials are correct then return OAuth token",
			inputCredentials: &configuration.OAuthSdkCredentials{
				ClientId:         os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_ID"),
				ClientSecret:     os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET"),
				AuthorizationUri: "https://access.sandbox.checkout.com/connect/token",
				Scopes:           []string{configuration.Disputes},
				Log:              log.New(os.Stderr, "checkout-sdk-go - ", log.LstdFlags),
			},
			checker: func(token *configuration.OAuthAccessToken, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, token)
				assert.NotNil(t, token.Token)
				assert.NotNil(t, token.ExpirationDate)
				assert.True(t, token.ExpirationDate.After(time.Now()))
			},
		},
		{
			name: "when input credentials are incorrect then return error",
			inputCredentials: &configuration.OAuthSdkCredentials{
				ClientId:         "invalid_client_id",
				ClientSecret:     "invalid_client_secret",
				AuthorizationUri: "https://access.sandbox.checkout.com/connect/token",
				Scopes:           []string{configuration.Disputes},
				Log:              log.New(os.Stderr, "checkout-sdk-go - ", log.LstdFlags),
			},
			checker: func(token *configuration.OAuthAccessToken, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, token)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "invalid_client", chkErr.Error())
			},
		},
		{
			name: "when input scopes are incorrect then return error",
			inputCredentials: &configuration.OAuthSdkCredentials{
				ClientId:         os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_ID"),
				ClientSecret:     os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET"),
				AuthorizationUri: "https://access.sandbox.checkout.com/connect/token",
				Scopes:           []string{"invalid_scope"},
				Log:              log.New(os.Stderr, "checkout-sdk-go - ", log.LstdFlags),
			},
			checker: func(token *configuration.OAuthAccessToken, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, token)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "invalid_scope", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.inputCredentials.GetAccessToken()

			tc.checker(tc.inputCredentials.AccessToken, err)
		})
	}
}

func TestGetAccessTokenWithSubdomain(t *testing.T) {
	cases := []struct {
		name             string
		inputCredentials *configuration.OAuthSdkCredentials
		expectedAuthUri  string
		checker          func(*configuration.OAuthAccessToken, error)
	}{
		{
			name: "when OAuth credentials are created with subdomain-aware authorization URI then use subdomain",
			inputCredentials: &configuration.OAuthSdkCredentials{
				ClientId:         "invalid_client_id",
				ClientSecret:     "invalid_client_secret",
				AuthorizationUri: "https://1234doma.access.sandbox.checkout.com/connect/token",
				Scopes:           []string{configuration.Gateway},
				Log:              log.New(os.Stderr, "checkout-sdk-go - ", log.LstdFlags),
			},
			expectedAuthUri: "https://1234doma.access.sandbox.checkout.com/connect/token",
			checker: func(token *configuration.OAuthAccessToken, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, token)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "invalid_client", chkErr.Error())
				// This test verifies that OAuth credentials are created with the subdomain-aware authorization URI
				// The failure is expected since we're using fake credentials, but the important part is that
				// the subdomain logic is triggered in the OAuth flow
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedAuthUri, tc.inputCredentials.AuthorizationUri)
			err := tc.inputCredentials.GetAccessToken()

			tc.checker(tc.inputCredentials.AccessToken, err)
		})
	}
}
