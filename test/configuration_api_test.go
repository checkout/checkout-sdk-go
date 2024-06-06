package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestShouldCreateConfigurationWithSubdomain(t *testing.T) {
	credentials := new(mocks.CredentialsMock)
	environment := configuration.Sandbox()

	testCases := []struct {
		subdomain   string
		expectedUrl string
	}{
		{"123dmain", "https://123dmain.api.sandbox.checkout.com"},
		{"123domain", "https://123domain.api.sandbox.checkout.com"},
		{"1234domain", "https://1234domain.api.sandbox.checkout.com"},
		{"12345domain", "https://12345domain.api.sandbox.checkout.com"},
	}

	for _, tc := range testCases {
		t.Run("Should create configuration with subdomain "+tc.subdomain, func(t *testing.T) {
			subdomain := configuration.NewEnvironmentSubdomain(environment, tc.subdomain)
			config := configuration.NewConfigurationWithSubdomain(credentials, environment, subdomain, &http.Client{}, nil)

			assert.NotNil(t, config)
			assert.Equal(t, tc.expectedUrl, config.EnvironmentSubdomain.ApiUrl)
		})
	}
}

func TestShouldCreateConfigurationWithBadSubdomain(t *testing.T) {
	credentials := new(mocks.CredentialsMock)
	environment := configuration.Sandbox()

	testCases := []struct {
		subdomain   string
		expectedUrl string
	}{
		{"", "https://api.sandbox.checkout.com"},
		{"123", "https://api.sandbox.checkout.com"},
		{"123bad", "https://api.sandbox.checkout.com"},
		{"12345domainBad", "https://api.sandbox.checkout.com"},
	}

	for _, tc := range testCases {
		t.Run("Should create configuration with bad subdomain "+tc.subdomain, func(t *testing.T) {
			subdomain := configuration.NewEnvironmentSubdomain(environment, tc.subdomain)
			config := configuration.NewConfigurationWithSubdomain(credentials, environment, subdomain, &http.Client{}, nil)

			assert.NotNil(t, config)
			assert.Equal(t, tc.expectedUrl, config.EnvironmentSubdomain.ApiUrl)
		})
	}
}
