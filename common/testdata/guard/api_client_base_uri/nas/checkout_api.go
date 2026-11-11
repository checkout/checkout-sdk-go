package nas

import (
	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
)

func Build(configuration *configuration.Configuration, baseUri string) *client.ApiClient {
	return client.NewApiClient(configuration, baseUri)
}
