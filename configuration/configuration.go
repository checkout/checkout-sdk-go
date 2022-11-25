package configuration

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/common"
)

type Configuration struct {
	Credentials SdkCredentials
	Environment Environment
	HttpClient  http.Client
}

func NewConfiguration(credentials SdkCredentials, environment Environment, client *http.Client) *Configuration {
	if environment == nil {
		environment = Sandbox()
	}

	if client == nil {
		client = common.BuildDefaultClient()
	}

	return &Configuration{
		Credentials: credentials,
		Environment: environment,
		HttpClient:  *client,
	}
}
