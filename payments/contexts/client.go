package contexts

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

// GetPaymentContextDetails sends GET /payment-contexts/{id}: get Payment Context details.
//
// paymentContextId is the {id} path value: payment Context's unique identifier.
//
// It returns the PaymentContextDetailsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentContextId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetPaymentContextDetails(paymentContextId string) (*PaymentContextDetailsResponse, error) {
	return c.GetPaymentContextDetailsWithContext(context.Background(), paymentContextId)
}

// GetPaymentContextDetailsWithContext is GetPaymentContextDetails with a caller-supplied context.
func (c *Client) GetPaymentContextDetailsWithContext(ctx context.Context, paymentContextId string) (*PaymentContextDetailsResponse, error) {
	paymentContextId, err := common.EncodePathSegment(paymentContextId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentContextDetailsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(PaymentContextsPath, paymentContextId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
