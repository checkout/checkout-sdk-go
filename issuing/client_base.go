package issuing

import (
	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/configuration"
)

const (
	issuing           = "issuing"
	cardholders       = "cardholders"
	cards             = "cards"
	threeDSEnrollment = "3ds-enrollment"
	activate          = "activate"
	credentials       = "credentials"
	revoke            = "revoke"
	suspend           = "suspend"
	controls          = "controls"
	simulate          = "simulate"
	authorizations    = "authorizations"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}
