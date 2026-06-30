package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
)

func TestShouldCreateConfigurationWithSubdomain(t *testing.T) {
	credentials := new(mocks.CredentialsMock)
	environment := configuration.Sandbox()

	testCases := []struct {
		subdomain       string
		expectedApiUrl  string
		expectedAuthUrl string
	}{
		{"a", "https://a.api.sandbox.checkout.com", "https://a.access.sandbox.checkout.com/connect/token"},
		{"ab", "https://ab.api.sandbox.checkout.com", "https://ab.access.sandbox.checkout.com/connect/token"},
		{"abc", "https://abc.api.sandbox.checkout.com", "https://abc.access.sandbox.checkout.com/connect/token"},
		{"abc1", "https://abc1.api.sandbox.checkout.com", "https://abc1.access.sandbox.checkout.com/connect/token"},
		{"12345domain", "https://12345domain.api.sandbox.checkout.com", "https://12345domain.access.sandbox.checkout.com/connect/token"},
		{"a1b2c3d4", "https://a1b2c3d4.api.sandbox.checkout.com", "https://a1b2c3d4.access.sandbox.checkout.com/connect/token"},
		{"12345678", "https://12345678.api.sandbox.checkout.com", "https://12345678.access.sandbox.checkout.com/connect/token"},
		{"abcdefgh", "https://abcdefgh.api.sandbox.checkout.com", "https://abcdefgh.access.sandbox.checkout.com/connect/token"},
		{"1234doma", "https://1234doma.api.sandbox.checkout.com", "https://1234doma.access.sandbox.checkout.com/connect/token"},
		{"pl-vkuhvk4v", "https://pl-vkuhvk4v.api.sandbox.checkout.com", "https://pl-vkuhvk4v.access.sandbox.checkout.com/connect/token"},
		{"pl-abc123", "https://pl-abc123.api.sandbox.checkout.com", "https://pl-abc123.access.sandbox.checkout.com/connect/token"},
	}

	for _, tc := range testCases {
		t.Run("Should create configuration with subdomain "+tc.subdomain, func(t *testing.T) {
			subdomain := configuration.NewEnvironmentSubdomain(environment, tc.subdomain)
			config := configuration.NewConfigurationWithSubdomain(credentials, environment, subdomain, &http.Client{}, nil)

			assert.NotNil(t, config)
			assert.Equal(t, tc.expectedApiUrl, config.EnvironmentSubdomain.ApiUrl)
			assert.Equal(t, tc.expectedAuthUrl, config.EnvironmentSubdomain.AuthorizationUrl)
		})
	}
}

func TestShouldCreateConfigurationWithBadSubdomain(t *testing.T) {
	credentials := new(mocks.CredentialsMock)
	environment := configuration.Sandbox()

	testCases := []struct {
		subdomain       string
		expectedApiUrl  string
		expectedAuthUrl string
	}{
		{"", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"  ", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{" - ", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"a b", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"ab c1", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"foo-", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"-foo", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"ABC", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"test-123", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"foo-bar", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
		{"pl-", "https://api.sandbox.checkout.com", "https://access.sandbox.checkout.com/connect/token"},
	}

	for _, tc := range testCases {
		t.Run("Should create configuration with bad subdomain "+tc.subdomain, func(t *testing.T) {
			subdomain := configuration.NewEnvironmentSubdomain(environment, tc.subdomain)
			config := configuration.NewConfigurationWithSubdomain(credentials, environment, subdomain, &http.Client{}, nil)

			assert.NotNil(t, config)
			assert.Equal(t, tc.expectedApiUrl, config.EnvironmentSubdomain.ApiUrl)
			assert.Equal(t, tc.expectedAuthUrl, config.EnvironmentSubdomain.AuthorizationUrl)
		})
	}
}

func TestShouldCreateConfigurationWithSubdomainForProduction(t *testing.T) {
	credentials := new(mocks.CredentialsMock)
	environment := configuration.Production()
	subdomain := "1234prod"

	subdomain_env := configuration.NewEnvironmentSubdomain(environment, subdomain)
	config := configuration.NewConfigurationWithSubdomain(credentials, environment, subdomain_env, &http.Client{}, nil)

	assert.NotNil(t, config)
	assert.Equal(t, "https://1234prod.api.checkout.com", config.EnvironmentSubdomain.ApiUrl)
	assert.Equal(t, "https://1234prod.access.checkout.com/connect/token", config.EnvironmentSubdomain.AuthorizationUrl)
	assert.Equal(t, "https://forward.checkout.com", config.Environment.ForwardUri())
	assert.Equal(t, "https://identity-verification.checkout.com", config.Environment.IdentityUri())
}

func TestShouldHaveCorrectSandboxUrls(t *testing.T) {
	env := configuration.Sandbox()

	assert.Equal(t, "https://api.sandbox.checkout.com", env.BaseUri())
	assert.Equal(t, "https://files.sandbox.checkout.com", env.FilesUri())
	assert.Equal(t, "https://transfers.sandbox.checkout.com", env.TransfersUri())
	assert.Equal(t, "https://balances.sandbox.checkout.com", env.BalancesUri())
	assert.Equal(t, "https://forward.sandbox.checkout.com", env.ForwardUri())
	assert.Equal(t, "https://identity-verification.sandbox.checkout.com", env.IdentityUri())
}

func TestShouldHaveCorrectProductionUrls(t *testing.T) {
	env := configuration.Production()

	assert.Equal(t, "https://api.checkout.com", env.BaseUri())
	assert.Equal(t, "https://files.checkout.com/", env.FilesUri())
	assert.Equal(t, "https://transfers.checkout.com/", env.TransfersUri())
	assert.Equal(t, "https://balances.checkout.com/", env.BalancesUri())
	assert.Equal(t, "https://forward.checkout.com", env.ForwardUri())
	assert.Equal(t, "https://identity-verification.checkout.com", env.IdentityUri())
}
