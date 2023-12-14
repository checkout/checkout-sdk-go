package contexts

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

func (c *Client) RequestPaymentContexts(request PaymentContextsRequest) (*PaymentContextsRequestResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentContextsRequestResponse
	err = c.apiClient.Post(common.BuildPath(PaymentContextsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetPaymentContextDetails(paymentContextId string) (*PaymentContextDetailsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentContextDetailsResponse
	err = c.apiClient.Get(common.BuildPath(PaymentContextsPath, paymentContextId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
