package sepa

import (
	"context"
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

func (c *Client) GetMandate(mandateId string) (*MandateResponse, error) {
	return c.GetMandateWithContext(context.Background(), mandateId)
}

func (c *Client) GetMandateWithContext(ctx context.Context, mandateId string) (*MandateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response MandateResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(sepaMandatesPath, mandateId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CancelMandate(mandateId string) (*SepaResource, error) {
	return c.CancelMandateWithContext(context.Background(), mandateId)
}

func (c *Client) CancelMandateWithContext(ctx context.Context, mandateId string) (*SepaResource, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response SepaResource
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.GetMandateViaPproWithContext(context.Background(), mandateId)
}

func (c *Client) GetMandateViaPproWithContext(ctx context.Context, mandateId string) (*MandateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response MandateResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(apmsPath, pproPath, sepaMandatesPath, mandateId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CancelMandateViaPpro(mandateId string) (*SepaResource, error) {
	return c.CancelMandateViaPproWithContext(context.Background(), mandateId)
}

func (c *Client) CancelMandateViaPproWithContext(ctx context.Context, mandateId string) (*SepaResource, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response SepaResource
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(apmsPath, pproPath, sepaMandatesPath, mandateId, cancelPath),
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
