package agenticcommerce

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

func (c *Client) CreateDelegatedPaymentToken(request CreateDelegatedPaymentTokenRequest, idempotencyKey *string) (*CreateDelegatedPaymentTokenResponse, error) {
	return c.CreateDelegatedPaymentTokenWithContext(context.Background(), request, idempotencyKey)
}

func (c *Client) CreateDelegatedPaymentTokenWithContext(ctx context.Context, request CreateDelegatedPaymentTokenRequest, idempotencyKey *string) (*CreateDelegatedPaymentTokenResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response CreateDelegatedPaymentTokenResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(delegatePaymentPath), auth, request, &response, idempotencyKey)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
