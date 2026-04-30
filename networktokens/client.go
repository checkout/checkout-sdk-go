package networktokens

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

func (c *Client) ProvisionNetworkToken(request ProvisionNetworkTokenRequest) (*NetworkTokenResponse, error) {
	return c.ProvisionNetworkTokenWithContext(context.Background(), request)
}

func (c *Client) ProvisionNetworkTokenWithContext(ctx context.Context, request ProvisionNetworkTokenRequest) (*NetworkTokenResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response NetworkTokenResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(networkTokensPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetNetworkToken(networkTokenId string) (*NetworkTokenResponse, error) {
	return c.GetNetworkTokenWithContext(context.Background(), networkTokenId)
}

func (c *Client) GetNetworkTokenWithContext(ctx context.Context, networkTokenId string) (*NetworkTokenResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response NetworkTokenResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(networkTokensPath, networkTokenId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RequestCryptogram(networkTokenId string, request RequestCryptogramRequest) (*RequestCryptogramResponse, error) {
	return c.RequestCryptogramWithContext(context.Background(), networkTokenId, request)
}

func (c *Client) RequestCryptogramWithContext(ctx context.Context, networkTokenId string, request RequestCryptogramRequest) (*RequestCryptogramResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response RequestCryptogramResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(networkTokensPath, networkTokenId, cryptogramsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteNetworkToken(networkTokenId string, request DeleteNetworkTokenRequest) (*common.MetadataResponse, error) {
	return c.DeleteNetworkTokenWithContext(context.Background(), networkTokenId, request)
}

func (c *Client) DeleteNetworkTokenWithContext(ctx context.Context, networkTokenId string, request DeleteNetworkTokenRequest) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PatchWithContext(ctx, common.BuildPath(networkTokensPath, networkTokenId, deletePath), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
