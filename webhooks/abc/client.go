package abc

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

func (c *Client) RetrieveWebhooks() (*WebhooksResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhooksResponse
	err = c.apiClient.Get(
		common.BuildPath(webhooks),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RegisterWebhook(request WebhookRequest) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.Post(
		common.BuildPath(webhooks),
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

func (c *Client) RetrieveWebhook(webhookId string) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.Get(
		common.BuildPath(webhooks, webhookId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateWebhook(webhookId string, request WebhookRequest) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.Put(
		common.BuildPath(webhooks, webhookId),
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

func (c *Client) PartiallyUpdateWebhook(webhookId string, request WebhookRequest) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.Patch(
		common.BuildPath(webhooks, webhookId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RemoveWebhook(webhookId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Delete(
		common.BuildPath(webhooks, webhookId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
