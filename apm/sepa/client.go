package sepa

import (
	"github.com/checkout/checkout-sdk-go-beta/client"
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
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

func (c *Client) GetMandate(mandateId string) (*MandateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response MandateResponse
	err = c.apiClient.Get(common.BuildPath(sepaMandatesPath, mandateId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CancelMandate(mandateId string) (*SepaResource, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response SepaResource
	err = c.apiClient.Post(
		common.BuildPath(sepaMandatesPath, mandateId, cancelPath),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetMandateViaPpro(mandateId string) (*MandateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response MandateResponse
	err = c.apiClient.Get(common.BuildPath(pproPath, sepaMandatesPath, mandateId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CancelMandateViaPpro(mandateId string) (*SepaResource, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response SepaResource
	err = c.apiClient.Post(
		common.BuildPath(pproPath, sepaMandatesPath, mandateId, cancelPath),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
