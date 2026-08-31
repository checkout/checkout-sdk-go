package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/nas"
)

func TestOauthCheckoutSdks(t *testing.T) {
	var oauthApi, _ = checkout.Builder().
		OAuth().
		WithClientCredentials(
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_ID"),
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET")).
		WithEnvironment(configuration.Sandbox()).
		// The sandbox OAuth clients lack subdomain provisioning, so the token request would
		// come back invalid_client. Opting out explicitly until they are provisioned.
		WithLegacyDomain().
		Build()

	var oauthApiBad, _ = checkout.Builder().
		OAuth().
		WithClientCredentials(
			"error",
			"error").
		WithEnvironment(configuration.Sandbox()).
		WithLegacyDomain().
		Build()

	cases := []struct {
		name     string
		oauthApi *nas.Api
		checker  func(*nas.Api, error)
	}{
		{
			name:     "should create a oauth checkout sdk api object",
			oauthApi: oauthApi,
			checker: func(token *nas.Api, err error) {
				assert.NotNil(t, oauthApi)
			},
		},
		// Not ready yet to tests with subdomains
		// {
		// 	name:     "should create a oauth checkout sdk api object with valid subdomain",
		// 	oauthApi: oauthApiSubdomain,
		// 	checker: func(token *nas.Api, err error) {
		// 		assert.NotNil(t, oauthApiSubdomain)
		// 	},
		// },
		{
			name:     "should fail a oauth checkout sdk api object",
			oauthApi: oauthApiBad,
			checker: func(token *nas.Api, err error) {
				assert.Nil(t, oauthApiBad)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(tc.oauthApi, nil)
		})
	}

}

func TestOauthCheckoutSdkWithSubdomain(t *testing.T) {
	// This test verifies that OAuth credentials are created with the subdomain-aware authorization URI
	// The failure is expected since we're using fake credentials, but the important part is that
	// the subdomain logic is triggered in the OAuth flow
	_, err := checkout.Builder().
		OAuth().
		WithClientCredentials("client_id", "client_secret").
		WithScopes([]string{configuration.Gateway}).
		WithEnvironment(configuration.Sandbox()).
		WithEnvironmentSubdomain("1234doma").
		Build()

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid_client")
}

func TestOauthShouldFailWithAuthorizationUriAndSubdomain(t *testing.T) {
	api, err := checkout.Builder().
		OAuth().
		WithClientCredentials("client_id", "client_secret").
		WithAuthorizationUri("https://access.sandbox.checkout.com/connect/token").
		WithEnvironment(configuration.Sandbox()).
		WithEnvironmentSubdomain("1234doma").
		Build()

	assert.Nil(t, api)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "authorization URI and environment subdomain cannot both be set")
}

func TestOauthShouldBuildWithAuthorizationUriAndLegacyDomain(t *testing.T) {
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600,"token_type":"Bearer"}`))
	}))
	defer tokenServer.Close()

	api, err := checkout.Builder().
		OAuth().
		WithClientCredentials("client_id", "client_secret").
		WithAuthorizationUri(tokenServer.URL).
		WithEnvironment(configuration.Sandbox()).
		WithLegacyDomain().
		Build()

	assert.Nil(t, err)
	assert.NotNil(t, api)
}
