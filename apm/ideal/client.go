package ideal

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

func (c *Client) GetInfo() (*IdealInfo, error) {
	return c.GetInfoWithContext(context.Background())
}

func (c *Client) GetInfoWithContext(ctx context.Context) (*IdealInfo, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response IdealInfo
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idealExternalPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetIssuers() (*IssuerResponse, error) {
	return c.GetIssuersWithContext(context.Background())
}

func (c *Client) GetIssuersWithContext(ctx context.Context) (*IssuerResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response IssuerResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idealExternalPath, issuersPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
