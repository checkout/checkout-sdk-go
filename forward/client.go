package forward

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

func (c *Client) ForwardAnApiRequest(request ForwardRequest) (*ForwardAnApiResponse, error) {
	return c.ForwardAnApiRequestWithContext(context.Background(), request)
}

func (c *Client) ForwardAnApiRequestWithContext(ctx context.Context, request ForwardRequest) (*ForwardAnApiResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ForwardAnApiResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(forward), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetForwardRequest(forwardId string) (*GetForwardResponse, error) {
	return c.GetForwardRequestWithContext(context.Background(), forwardId)
}

func (c *Client) GetForwardRequestWithContext(ctx context.Context, forwardId string) (*GetForwardResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetForwardResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(forward, forwardId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
