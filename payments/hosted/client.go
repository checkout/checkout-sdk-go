package hosted

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

// GetHostedPaymentsPageDetails sends GET /hosted-payments/{id}: get Hosted Payments Page details.
//
// hostedPaymentId is the {id} path value.
//
// It returns the PaymentHostedDetails, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." hostedPaymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetHostedPaymentsPageDetails(hostedPaymentId string) (*PaymentHostedDetails, error) {
	return c.GetHostedPaymentsPageDetailsWithContext(context.Background(), hostedPaymentId)
}

// GetHostedPaymentsPageDetailsWithContext is GetHostedPaymentsPageDetails with a caller-supplied context.
func (c *Client) GetHostedPaymentsPageDetailsWithContext(ctx context.Context, hostedPaymentId string) (*PaymentHostedDetails, error) {
	hostedPaymentId, err := common.EncodePathSegment(hostedPaymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentHostedDetails
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(HostedPaymentsPath, hostedPaymentId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
