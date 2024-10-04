package customers

import (
	"context"
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

func (c *Client) Create(request CustomerRequest) (*common.IdResponse, error) {
	return c.CreateWithContext(context.Background(), request)
}

func (c *Client) CreateWithContext(ctx context.Context, request CustomerRequest) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(Path), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Get(customerId string) (*GetCustomerResponse, error) {
	return c.GetWithContext(context.Background(), customerId)
}

func (c *Client) GetWithContext(ctx context.Context, customerId string) (*GetCustomerResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetCustomerResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(Path, customerId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Update(customerId string, request CustomerRequest) (*common.MetadataResponse, error) {
	return c.UpdateWithContext(context.Background(), customerId, request)
}

func (c *Client) UpdateWithContext(ctx context.Context, customerId string, request CustomerRequest) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PatchWithContext(ctx, common.BuildPath(Path, customerId), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Delete(customerId string) (*common.MetadataResponse, error) {
	return c.DeleteWithContext(context.Background(), customerId)
}

func (c *Client) DeleteWithContext(ctx context.Context, customerId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(ctx, common.BuildPath(Path, customerId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
