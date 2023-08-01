package links

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

func (c *Client) CreatePaymentLink(request PaymentLinkRequest) (*PaymentLinkResponse, error) {
	return c.CreatePaymentLinkWithContext(context.Background(), request)
}

func (c *Client) CreatePaymentLinkWithContext(ctx context.Context, request PaymentLinkRequest) (*PaymentLinkResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentLinkResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(PaymentLinksPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetPaymentLink(paymentLinkId string) (*PaymentLinkDetails, error) {
	return c.GetPaymentLinkWithContext(context.Background(), paymentLinkId)
}

func (c *Client) GetPaymentLinkWithContext(ctx context.Context, paymentLinkId string) (*PaymentLinkDetails, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentLinkDetails
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(PaymentLinksPath, paymentLinkId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
