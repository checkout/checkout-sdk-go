package balances

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

func (c *Client) RetrieveEntityBalances(entityId string, query QueryFilter) (*QueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(balances, entityId), query)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	err = c.apiClient.Get(
		url,
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
