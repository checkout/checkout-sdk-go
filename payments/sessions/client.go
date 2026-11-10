package payment_sessions

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

func (c *Client) RequestPaymentSessions(request PaymentSessionsRequest) (*PaymentSessionsResponse, error) {
	return c.RequestPaymentSessionsWithContext(context.Background(), request)
}

func (c *Client) RequestPaymentSessionsWithContext(ctx context.Context, request PaymentSessionsRequest) (*PaymentSessionsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSessionsResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(PaymentSessionsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RequestPaymentSessionsWithPayment(request PaymentSessionsWithPaymentRequest) (*PaymentSessionPaymentResponse, error) {
	return c.RequestPaymentSessionsWithPaymentWithContext(context.Background(), request)
}

func (c *Client) RequestPaymentSessionsWithPaymentWithContext(ctx context.Context, request PaymentSessionsWithPaymentRequest) (*PaymentSessionPaymentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSessionPaymentResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(PaymentSessionsCompletePath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SubmitPaymentSession sends POST /payment-sessions/{id}/submit: submit a Payment Session.
//
// sessionId is the {id} path value: the Payment Sessions unique identifier. request is sent as the
// JSON body.
//
// It returns the PaymentSessionPaymentResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." sessionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) SubmitPaymentSession(sessionId string, request SubmitPaymentSessionRequest) (*PaymentSessionPaymentResponse, error) {
	return c.SubmitPaymentSessionWithContext(context.Background(), sessionId, request)
}

// SubmitPaymentSessionWithContext is SubmitPaymentSession with a caller-supplied context.
func (c *Client) SubmitPaymentSessionWithContext(ctx context.Context, sessionId string, request SubmitPaymentSessionRequest) (*PaymentSessionPaymentResponse, error) {
	sessionId, err := common.EncodePathSegment(sessionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSessionPaymentResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(PaymentSessionsPath, sessionId, PaymentSessionsSubmitPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
