package networktokens

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

// Client holds the dependencies for making network tokens API requests.
type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

// NewClient creates a network tokens Client using the provided configuration and HTTP client.
func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

// ProvisionNetworkToken provisions a new network token for the given card source.
func (c *Client) ProvisionNetworkToken(request ProvisionNetworkTokenRequest) (*NetworkTokenResponse, error) {
	return c.ProvisionNetworkTokenWithContext(context.Background(), request)
}

// ProvisionNetworkTokenWithContext is like ProvisionNetworkToken but accepts a context for
// cancellation and deadline propagation.
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

// GetNetworkToken retrieves the network token identified by networkTokenId.
func (c *Client) GetNetworkToken(networkTokenId string) (*NetworkTokenResponse, error) {
	return c.GetNetworkTokenWithContext(context.Background(), networkTokenId)
}

// GetNetworkTokenWithContext is like GetNetworkToken but accepts a context for cancellation
// and deadline propagation.
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

// RequestCryptogram generates a cryptogram for the network token identified by networkTokenId.
func (c *Client) RequestCryptogram(networkTokenId string, request RequestCryptogramRequest) (*RequestCryptogramResponse, error) {
	return c.RequestCryptogramWithContext(context.Background(), networkTokenId, request)
}

// RequestCryptogramWithContext is like RequestCryptogram but accepts a context for cancellation
// and deadline propagation.
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

// DeleteNetworkToken deletes the network token identified by networkTokenId.
func (c *Client) DeleteNetworkToken(networkTokenId string, request DeleteNetworkTokenRequest) (*common.MetadataResponse, error) {
	return c.DeleteNetworkTokenWithContext(context.Background(), networkTokenId, request)
}

// DeleteNetworkTokenWithContext is like DeleteNetworkToken but accepts a context for
// cancellation and deadline propagation.
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
