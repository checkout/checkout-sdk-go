package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/nas"
)

func TestDefaultCheckoutSdks(t *testing.T) {
	var defaultApi, _ = checkout.Builder().
		StaticKeys().
		WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
		WithPublicKey(os.Getenv("CHECKOUT_DEFAULT_PUBLIC_KEY")).
		WithEnvironment(configuration.Sandbox()).
		// The sandbox OAuth clients are not provisioned for the merchant-specific subdomain,
		// so the token request would come back invalid_client. Opting out explicitly until
		// they are.
		WithLegacyDomain().
		Build()

	var defaultApiSubdomain, _ = checkout.Builder().
		StaticKeys().
		WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
		WithPublicKey(os.Getenv("CHECKOUT_DEFAULT_PUBLIC_KEY")).
		WithEnvironment(configuration.Sandbox()).
		WithEnvironmentSubdomain("123dmain").
		Build()

	var defaultApiBad, _ = checkout.Builder().
		StaticKeys().
		WithSecretKey("error").
		WithPublicKey("error").
		WithEnvironment(configuration.Sandbox()).
		WithLegacyDomain().
		Build()

	cases := []struct {
		name       string
		defaultApi *nas.Api
		checker    func(*nas.Api, error)
	}{
		{
			name:       "should create a default checkout sdk api object",
			defaultApi: defaultApi,
			checker: func(token *nas.Api, err error) {
				assert.NotNil(t, defaultApi)
			},
		},
		{
			name:       "should create a default checkout sdk api object with valid subdomain",
			defaultApi: defaultApiSubdomain,
			checker: func(token *nas.Api, err error) {
				assert.NotNil(t, defaultApiSubdomain)
			},
		},
		{
			name:       "should fail a default checkout sdk api object",
			defaultApi: defaultApiBad,
			checker: func(token *nas.Api, err error) {
				assert.Nil(t, defaultApiBad)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(tc.defaultApi, nil)
		})
	}

}

func TestShouldFailWithoutSubdomainOrLegacyDomain(t *testing.T) {
	api, err := checkout.Builder().
		StaticKeys().
		WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
		WithEnvironment(configuration.Sandbox()).
		Build()

	assert.Nil(t, api)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "environment subdomain is required")
}

func TestShouldFailWithBothSubdomainAndLegacyDomain(t *testing.T) {
	api, err := checkout.Builder().
		StaticKeys().
		WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
		WithEnvironment(configuration.Sandbox()).
		WithEnvironmentSubdomain("1234doma").
		WithLegacyDomain().
		Build()

	assert.Nil(t, api)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "cannot both be set")
}

func TestShouldFailWithInvalidSubdomain(t *testing.T) {
	api, err := checkout.Builder().
		StaticKeys().
		WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
		WithEnvironment(configuration.Sandbox()).
		WithEnvironmentSubdomain("not a subdomain").
		Build()

	assert.Nil(t, api)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "invalid environment subdomain")
}

func TestShouldCreateSdkWithLegacyDomain(t *testing.T) {
	api, err := checkout.Builder().
		StaticKeys().
		WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
		WithPublicKey(os.Getenv("CHECKOUT_DEFAULT_PUBLIC_KEY")).
		WithEnvironment(configuration.Sandbox()).
		WithLegacyDomain().
		Build()

	assert.Nil(t, err)
	assert.NotNil(t, api)
}

func TestShouldCreatePreviousSdkWithoutSubdomain(t *testing.T) {
	api, err := checkout.Builder().Previous().
		WithSecretKey(os.Getenv("CHECKOUT_PREVIOUS_SECRET_KEY")).
		WithPublicKey(os.Getenv("CHECKOUT_PREVIOUS_PUBLIC_KEY")).
		WithEnvironment(configuration.Sandbox()).
		Build()

	assert.Nil(t, err)
	assert.NotNil(t, api)
}
