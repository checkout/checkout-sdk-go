package amlscreening

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

func (c *Client) CreateAmlScreening(request CreateAmlScreeningRequest) (*AmlScreeningResponse, error) {
	return c.CreateAmlScreeningWithContext(context.Background(), request)
}

func (c *Client) CreateAmlScreeningWithContext(ctx context.Context, request CreateAmlScreeningRequest) (*AmlScreeningResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AmlScreeningResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(amlScreeningPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAmlScreening(screeningId string) (*AmlScreeningResponse, error) {
	return c.GetAmlScreeningWithContext(context.Background(), screeningId)
}

func (c *Client) GetAmlScreeningWithContext(ctx context.Context, screeningId string) (*AmlScreeningResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AmlScreeningResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(amlScreeningPath, screeningId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
