package paymentmethods

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

func (c *Client) GetAvailablePaymentMethods(query GetPaymentMethodsQuery) (*GetPaymentMethodsResponse, error) {
	return c.GetAvailablePaymentMethodsWithContext(context.Background(), query)
}

func (c *Client) GetAvailablePaymentMethodsWithContext(ctx context.Context, query GetPaymentMethodsQuery) (*GetPaymentMethodsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(paymentMethodsPath), query)
	if err != nil {
		return nil, err
	}

	var response GetPaymentMethodsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
