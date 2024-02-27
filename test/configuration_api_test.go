package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestShouldCreateConfiguration(t *testing.T) {
	credentials := new(mocks.CredentialsMock)
	environment := configuration.Production()
	subdomain := configuration.NewEnvironmentSubdomain(environment, "123dmain")
	subdomainBad := configuration.NewEnvironmentSubdomain(environment, "baddomain")

	cases := []struct {
		name          string
		configuration *configuration.Configuration
		checker       func(*configuration.Configuration, error)
	}{
		{
			name:          "should create a configuration object",
			configuration: configuration.NewConfiguration(credentials, environment, &http.Client{}, nil),
			checker: func(configuration *configuration.Configuration, err error) {
				assert.NotNil(t, configuration)
			},
		},
		{
			name:          "should create a configuration object with a valid subdomain",
			configuration: configuration.NewConfigurationWithSubdomain(credentials, environment, subdomain, &http.Client{}, nil),
			checker: func(configuration *configuration.Configuration, err error) {
				assert.NotNil(t, configuration)
				assert.Equal(t, "https://123dmain.api.checkout.com", configuration.EnvironmentSubdomain.ApiUrl)
			},
		},
		{
			name:          "should create a configuration object with a invalid subdomain",
			configuration: configuration.NewConfigurationWithSubdomain(credentials, environment, subdomainBad, &http.Client{}, nil),
			checker: func(configuration *configuration.Configuration, err error) {
				assert.NotNil(t, configuration)
				assert.Equal(t, "https://api.checkout.com", configuration.EnvironmentSubdomain.ApiUrl)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(tc.configuration, nil)
		})
	}
}
