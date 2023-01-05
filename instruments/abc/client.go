package abc

import (
	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/instruments"
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

func (c *Client) Create(request CreateInstrumentRequest) (*CreateInstrumentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response CreateInstrumentResponse
	err = c.apiClient.Post(
		common.BuildPath(instruments.Path),
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

func (c *Client) Get(instrumentId string) (*GetInstrumentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response GetInstrumentResponse
	err = c.apiClient.Get(common.BuildPath(instruments.Path, instrumentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Update(instrumentId string, request UpdateInstrumentRequest) (*UpdateInstrumentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response UpdateInstrumentResponse
	err = c.apiClient.Patch(
		common.BuildPath(instruments.Path, instrumentId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Delete(instrumentId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Delete(common.BuildPath(instruments.Path, instrumentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
