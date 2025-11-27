package setups

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

func (c *Client) CreatePaymentSetup(request PaymentSetupRequest) (*PaymentSetupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupResponse
	err = c.apiClient.Post(common.BuildPath(PaymentSetupsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdatePaymentSetup(setupId string, request PaymentSetupRequest) (*PaymentSetupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupResponse
	err = c.apiClient.Put(common.BuildPath(PaymentSetupsPath, setupId), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetPaymentSetup(setupId string) (*PaymentSetupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupResponse
	err = c.apiClient.Get(common.BuildPath(PaymentSetupsPath, setupId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ConfirmPaymentSetup(setupId string, paymentMethodOptionId string, request PaymentSetupConfirmRequest) (*PaymentSetupConfirmResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupConfirmResponse
	err = c.apiClient.Post(common.BuildPath(PaymentSetupsPath, setupId, ConfirmPath, paymentMethodOptionId), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}