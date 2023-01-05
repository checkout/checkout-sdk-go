package hosted

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

func (c *Client) CreateHostedPaymentsPageSession(request HostedPaymentRequest) (*HostedPaymentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response HostedPaymentResponse
	err = c.apiClient.Post(common.BuildPath(HostedPaymentsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetHostedPaymentsPageDetails(hostedPaymentId string) (*HostedPaymentDetails, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response HostedPaymentDetails
	err = c.apiClient.Get(common.BuildPath(HostedPaymentsPath, hostedPaymentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
