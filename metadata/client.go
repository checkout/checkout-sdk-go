package metadata

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

func (c *Client) RequestCardMetadata(request CardMetadataRequest) (*CardMetadataResponse, error) {
	return c.RequestCardMetadataWithContext(context.Background(), request)
}

func (c *Client) RequestCardMetadataWithContext(
	ctx context.Context,
	request CardMetadataRequest,
) (*CardMetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CardMetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(metadata, card),
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
