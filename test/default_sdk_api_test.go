package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/nas"
)

func TestDefaultCheckoutSdks(t *testing.T) {
	var defaultApi, _ = checkout.Builder().
		StaticKeys().
		WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
		WithPublicKey(os.Getenv("CHECKOUT_DEFAULT_PUBLIC_KEY")).
		WithEnvironment(configuration.Sandbox()).
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
