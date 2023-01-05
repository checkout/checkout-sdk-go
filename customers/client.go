package customers

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

func (c *Client) Create(request CustomerRequest) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(common.BuildPath(Path), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Get(customerId string) (*GetCustomerResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetCustomerResponse
	err = c.apiClient.Get(common.BuildPath(common.BuildPath(Path), customerId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Update(customerId string, request CustomerRequest) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Patch(common.BuildPath(Path, customerId), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Delete(customerId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Delete(common.BuildPath(Path, customerId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
