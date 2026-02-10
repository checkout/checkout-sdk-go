package hosted

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

func (c *Client) CreateHostedPaymentsPageSession(request PaymentHostedRequest) (*PaymentHostedResponse, error) {
	return c.CreateHostedPaymentsPageSessionWithContext(context.Background(), request)
}

func (c *Client) CreateHostedPaymentsPageSessionWithContext(ctx context.Context, request PaymentHostedRequest) (*PaymentHostedResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentHostedResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(HostedPaymentsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetHostedPaymentsPageDetails(hostedPaymentId string) (*PaymentHostedDetails, error) {
	return c.GetHostedPaymentsPageDetailsWithContext(context.Background(), hostedPaymentId)
}

func (c *Client) GetHostedPaymentsPageDetailsWithContext(ctx context.Context, hostedPaymentId string) (*PaymentHostedDetails, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentHostedDetails
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(HostedPaymentsPath, hostedPaymentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
