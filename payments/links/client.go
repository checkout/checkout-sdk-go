package links

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

func (c *Client) CreatePaymentLink(request PaymentLinkRequest) (*PaymentLinkResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentLinkResponse
	err = c.apiClient.Post(common.BuildPath(PaymentLinksPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetPaymentLink(paymentLinkId string) (*PaymentLinkDetails, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentLinkDetails
	err = c.apiClient.Get(common.BuildPath(PaymentLinksPath, paymentLinkId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
