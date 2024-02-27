package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/abc"
	"github.com/checkout/checkout-sdk-go/configuration"
)

func TestPreviousCheckoutSdks(t *testing.T) {
	var previousApi, _ = checkout.Builder().
		Previous().
		WithSecretKey(os.Getenv("CHECKOUT_PREVIOUS_SECRET_KEY")).
		WithPublicKey(os.Getenv("CHECKOUT_PREVIOUS_PUBLIC_KEY")).
		WithEnvironment(configuration.Sandbox()).
		Build()

	var previousApiSubdomain, _ = checkout.Builder().
		Previous().
		WithSecretKey(os.Getenv("CHECKOUT_PREVIOUS_SECRET_KEY")).
		WithPublicKey(os.Getenv("CHECKOUT_PREVIOUS_PUBLIC_KEY")).
		WithEnvironment(configuration.Sandbox()).
		WithEnvironmentSubdomain("123dmain").
		Build()

	var previousApiBad, _ = checkout.Builder().
		Previous().
		WithSecretKey("error").
		WithPublicKey("error").
		WithEnvironment(configuration.Sandbox()).
		Build()

	cases := []struct {
		name        string
		previousApi *abc.Api
		checker     func(*abc.Api, error)
	}{
		{
			name:        "should create a previous checkout sdk api object",
			previousApi: previousApi,
			checker: func(token *abc.Api, err error) {
				assert.NotNil(t, previousApi)
			},
		},
		{
			name:        "should create a previous checkout sdk api object with valid subdomain",
			previousApi: previousApiSubdomain,
			checker: func(token *abc.Api, err error) {
				assert.NotNil(t, previousApiSubdomain)
			},
		},
		{
			name:        "should fail a previous checkout sdk api object",
			previousApi: previousApiBad,
			checker: func(token *abc.Api, err error) {
				assert.Nil(t, previousApiBad)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(tc.previousApi, nil)
		})
	}

}
