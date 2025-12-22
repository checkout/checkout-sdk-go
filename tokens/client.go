package tokens

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

func (c *Client) RequestCardToken(request CardTokenRequest) (*CardTokenResponse, error) {
	return c.RequestCardTokenWithContext(context.Background(), request)
}

func (c *Client) RequestCardTokenWithContext(ctx context.Context, request CardTokenRequest) (*CardTokenResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response CardTokenResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(tokensPath),
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

func (c *Client) RequestWalletToken(request WalletTokenRequest) (*CardTokenResponse, error) {
	return c.RequestWalletTokenWithContext(context.Background(), request)
}

func (c *Client) RequestWalletTokenWithContext(ctx context.Context, request WalletTokenRequest) (*CardTokenResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response CardTokenResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(tokensPath),
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
