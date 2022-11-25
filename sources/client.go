package sources

import (
	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
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

func (c *Client) CreateSepaSource(request *sepaSourceRequest) (*CreateSepaSourceResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response CreateSepaSourceResponse
	err = c.apiClient.Post(
		common.BuildPath(path),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
