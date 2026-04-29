package standaloneaccountupdater

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
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

func (c *Client) GetUpdatedCardCredentials(request GetUpdatedCardCredentialsRequest) (*GetUpdatedCardCredentialsResponse, error) {
	return c.GetUpdatedCardCredentialsWithContext(context.Background(), request)
}

func (c *Client) GetUpdatedCardCredentialsWithContext(ctx context.Context, request GetUpdatedCardCredentialsRequest) (*GetUpdatedCardCredentialsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response GetUpdatedCardCredentialsResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(accountUpdaterPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
