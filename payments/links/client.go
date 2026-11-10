package links

import (
	"context"
	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
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

// GetPaymentLink sends GET /payment-links/{id}: get Payment Link details.
//
// paymentLinkId is the {id} path value.
//
// It returns the PaymentLinkDetails, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentLinkId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetPaymentLink(paymentLinkId string) (*PaymentLinkDetails, error) {
	return c.GetPaymentLinkWithContext(context.Background(), paymentLinkId)
}

// GetPaymentLinkWithContext is GetPaymentLink with a caller-supplied context.
func (c *Client) GetPaymentLinkWithContext(ctx context.Context, paymentLinkId string) (*PaymentLinkDetails, error) {
	paymentLinkId, err := common.EncodePathSegment(paymentLinkId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentLinkDetails
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(PaymentLinksPath, paymentLinkId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
