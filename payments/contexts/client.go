package contexts

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

func (c *Client) RequestPaymentContexts(request PaymentContextsRequest) (*PaymentContextsRequestResponse, error) {
	return c.RequestPaymentContextsWithContext(context.Background(), request)
}

func (c *Client) RequestPaymentContextsWithContext(ctx context.Context, request PaymentContextsRequest) (*PaymentContextsRequestResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentContextsRequestResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(PaymentContextsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetPaymentContextDetails(paymentContextId string) (*PaymentContextDetailsResponse, error) {
	return c.GetPaymentContextDetailsWithContext(context.Background(), paymentContextId)
}

func (c *Client) GetPaymentContextDetailsWithContext(ctx context.Context, paymentContextId string) (*PaymentContextDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentContextDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(PaymentContextsPath, paymentContextId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
