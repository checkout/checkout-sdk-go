package tokens

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

func (c *Client) RequestCvvToken(request CvvTokenRequest) (*CvvTokenResponse, error) {
	return c.RequestCvvTokenWithContext(context.Background(), request)
}

func (c *Client) RequestCvvTokenWithContext(ctx context.Context, request CvvTokenRequest) (*CvvTokenResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response CvvTokenResponse
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

func (c *Client) RequestPinToken(request PinTokenRequest) (*PinTokenResponse, error) {
	return c.RequestPinTokenWithContext(context.Background(), request)
}

func (c *Client) RequestPinTokenWithContext(ctx context.Context, request PinTokenRequest) (*PinTokenResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response PinTokenResponse
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

func (c *Client) GetTokenMetadata(tokenId string) (*TokenMetadataResponse, error) {
	return c.GetTokenMetadataWithContext(context.Background(), tokenId)
}

func (c *Client) GetTokenMetadataWithContext(ctx context.Context, tokenId string) (*TokenMetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response TokenMetadataResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(tokensPath, tokenId, "metadata"), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
