package klarna

import (
	"context"
	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/payments"
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

func (c *Client) GetCreditSession(sessionId string) (*CreditSession, error) {
	return c.GetCreditSessionWithContext(context.Background(), sessionId)
}

func (c *Client) GetCreditSessionWithContext(ctx context.Context, sessionId string) (*CreditSession, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.PublicKey)
	if err != nil {
		return nil, err
	}

	var response CreditSession
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(c.getBaseUrl(), creditSessionPath, sessionId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CapturePayment(paymentId string, request OrderCaptureRequest) (*CaptureResponse, error) {
	return c.CapturePaymentWithContext(context.Background(), paymentId, request)
}

func (c *Client) CapturePaymentWithContext(ctx context.Context, paymentId string, request OrderCaptureRequest) (*CaptureResponse, error) {
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

func (c *Client) VoidPayment(paymentId string, request payments.VoidRequest) (*payments.VoidResponse, error) {
	return c.VoidPaymentWithContext(context.Background(), paymentId, request)
}

func (c *Client) VoidPaymentWithContext(ctx context.Context, paymentId string, request payments.VoidRequest) (*payments.VoidResponse, error) {
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
