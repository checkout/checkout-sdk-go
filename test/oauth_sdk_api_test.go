package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/nas"
)

func TestOauthCheckoutSdks(t *testing.T) {
	var oauthApi, _ = checkout.Builder().
		OAuth().
		WithClientCredentials(
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_ID"),
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET")).
		WithEnvironment(configuration.Sandbox()).
		Build()

	var oauthApiSubdomain, _ = checkout.Builder().
		OAuth().
		WithClientCredentials(
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_ID"),
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET")).
		WithEnvironment(configuration.Sandbox()).
		WithEnvironmentSubdomain("123dmain").
		Build()

	var oauthApiBad, _ = checkout.Builder().
		OAuth().
		WithClientCredentials(
			"error",
			"error").
		WithEnvironment(configuration.Sandbox()).
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
		{
			name:     "should create a oauth checkout sdk api object with valid subdomain",
			oauthApi: oauthApiSubdomain,
			checker: func(token *nas.Api, err error) {
				assert.NotNil(t, oauthApiSubdomain)
			},
		},
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
