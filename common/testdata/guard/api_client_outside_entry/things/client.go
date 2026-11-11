package things

import (
	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
)

func Build(configuration *configuration.Configuration) *client.ApiClient {
	return client.NewApiClient(configuration, configuration.Environment.BaseUri())
}
