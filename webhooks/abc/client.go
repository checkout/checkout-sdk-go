package abc

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

func (c *Client) RetrieveWebhooks() (*WebhooksResponse, error) {
	return c.RetrieveWebhooksWithContext(context.Background())
}

func (c *Client) RetrieveWebhooksWithContext(ctx context.Context) (*WebhooksResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhooksResponse
	err = c.apiClient.GetWithContext(
		ctx,
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
	return c.RegisterWebhookWithContext(context.Background(), request)
}

func (c *Client) RegisterWebhookWithContext(ctx context.Context, request WebhookRequest) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.PostWithContext(
		ctx,
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
	return c.RetrieveWebhookWithContext(context.Background(), webhookId)
}

func (c *Client) RetrieveWebhookWithContext(ctx context.Context, webhookId string) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.GetWithContext(
		ctx,
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
	return c.UpdateWebhookWithContext(context.Background(), webhookId, request)
}

func (c *Client) UpdateWebhookWithContext(ctx context.Context, webhookId string, request WebhookRequest) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.PutWithContext(
		ctx,
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
	return c.PartiallyUpdateWebhookWithContext(context.Background(), webhookId, request)
}

func (c *Client) PartiallyUpdateWebhookWithContext(ctx context.Context, webhookId string, request WebhookRequest) (*WebhookResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response WebhookResponse
	err = c.apiClient.PatchWithContext(
		ctx,
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
	return c.RemoveWebhookWithContext(context.Background(), webhookId)
}

func (c *Client) RemoveWebhookWithContext(ctx context.Context, webhookId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(webhooks, webhookId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
