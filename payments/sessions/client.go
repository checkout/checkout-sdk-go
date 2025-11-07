package payment_sessions

import (
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

func (c *Client) RequestPaymentSessions(request PaymentSessionsRequest) (*PaymentSessionsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentSessionsResponse
	err = c.apiClient.Post(common.BuildPath(PaymentSessionsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RequestPaymentSessionsWithPayment(request PaymentSessionsWithPaymentRequest) (*PaymentSessionPaymentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentSessionPaymentResponse
	err = c.apiClient.Post(common.BuildPath(PaymentSessionsCompletePath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SubmitPaymentSession(sessionId string, request SubmitPaymentSessionRequest) (*PaymentSessionPaymentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentSessionPaymentResponse
	err = c.apiClient.Post(common.BuildPath(PaymentSessionsPath, sessionId, PaymentSessionsSubmitPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
