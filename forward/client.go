package forward

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

func (c *Client) CreateSecret(request CreateSecretRequest) (*SingleSecretResponse, error) {
	return c.CreateSecretWithContext(context.Background(), request)
}

func (c *Client) CreateSecretWithContext(ctx context.Context, request CreateSecretRequest) (*SingleSecretResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response SingleSecretResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(forward, secrets), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ListSecrets() (*ListSecretsResponse, error) {
	return c.ListSecretsWithContext(context.Background())
}

func (c *Client) ListSecretsWithContext(ctx context.Context) (*ListSecretsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ListSecretsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(forward, secrets), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateSecret(name string, request UpdateSecretRequest) (*SingleSecretResponse, error) {
	return c.UpdateSecretWithContext(context.Background(), name, request)
}

func (c *Client) UpdateSecretWithContext(ctx context.Context, name string, request UpdateSecretRequest) (*SingleSecretResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response SingleSecretResponse
	err = c.apiClient.PatchWithContext(ctx, common.BuildPath(forward, secrets, name), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) DeleteSecret(name string) (*common.MetadataResponse, error) {
	return c.DeleteSecretWithContext(context.Background(), name)
}

func (c *Client) DeleteSecretWithContext(ctx context.Context, name string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(ctx, common.BuildPath(forward, secrets, name), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
