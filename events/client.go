package events

import (
	"github.com/checkout/checkout-sdk-go-beta/client"
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
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

func (c *Client) RetrieveAllEventTypes(version ...string) (*EventTypesResponse, error) {
	path := path
	if version != nil {
		path += "?version=" + version[0]
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventTypesResponse
	err = c.apiClient.Get(common.BuildPath(path), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
