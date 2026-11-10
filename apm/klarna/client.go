package klarna

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/payments"
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

func (c *Client) CreateCreditSession(request CreditSessionRequest) (*CreditSessionResponse, error) {
	return c.CreateCreditSessionWithContext(context.Background(), request)
}

func (c *Client) CreateCreditSessionWithContext(ctx context.Context, request CreditSessionRequest) (*CreditSessionResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response CreditSessionResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(c.getBaseUrl(), creditSessionPath),
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

// GetCreditSession sends GET /klarna/credit-sessions/{sessionId} (klarna-external in sandbox).
//
// sessionId is the {sessionId} path value.
//
// It returns the CreditSession, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." sessionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetCreditSession(sessionId string) (*CreditSession, error) {
	return c.GetCreditSessionWithContext(context.Background(), sessionId)
}

// GetCreditSessionWithContext is GetCreditSession with a caller-supplied context.
func (c *Client) GetCreditSessionWithContext(ctx context.Context, sessionId string) (*CreditSession, error) {
	sessionId, err := common.EncodePathSegment(sessionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response CreditSession
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(c.getBaseUrl(), creditSessionPath, sessionId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CapturePayment sends POST /klarna/orders/{paymentId}/captures (klarna-external in sandbox).
//
// paymentId is the {paymentId} path value. request is sent as the JSON body.
//
// It returns the CaptureResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CapturePayment(paymentId string, request OrderCaptureRequest) (*CaptureResponse, error) {
	return c.CapturePaymentWithContext(context.Background(), paymentId, request)
}

// CapturePaymentWithContext is CapturePayment with a caller-supplied context.
func (c *Client) CapturePaymentWithContext(ctx context.Context, paymentId string, request OrderCaptureRequest) (*CaptureResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response CaptureResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(c.getBaseUrl(), ordersPath, paymentId, capturesPath),
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

// VoidPayment sends POST /klarna/orders/{paymentId}/voids (klarna-external in sandbox).
//
// paymentId is the {paymentId} path value. request is sent as the JSON body.
//
// It returns the payments.VoidResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." paymentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) VoidPayment(paymentId string, request payments.VoidRequest) (*payments.VoidResponse, error) {
	return c.VoidPaymentWithContext(context.Background(), paymentId, request)
}

// VoidPaymentWithContext is VoidPayment with a caller-supplied context.
func (c *Client) VoidPaymentWithContext(ctx context.Context, paymentId string, request payments.VoidRequest) (*payments.VoidResponse, error) {
	paymentId, err := common.EncodePathSegment(paymentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response payments.VoidResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(c.getBaseUrl(), ordersPath, paymentId, voidsPath),
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

func (c *Client) getBaseUrl() string {
	if c.configuration.Environment.IsSandbox() {
		return "klarna-external"
	}

	return "klarna"
}
