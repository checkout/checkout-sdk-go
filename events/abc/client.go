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

func (c *Client) RetrieveAllEventTypes() (*EventTypesResponse, error) {
	return c.RetrieveAllEventTypesWithContext(context.Background())
}

func (c *Client) RetrieveAllEventTypesWithContext(ctx context.Context) (*EventTypesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventTypesResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(eventTypes), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveAllEventTypesQuery(
	query QueryRetrieveAllEventType,
) (*EventTypesResponse, error) {
	return c.RetrieveAllEventTypesQueryWithContext(context.Background(), query)
}

func (c *Client) RetrieveAllEventTypesQueryWithContext(
	ctx context.Context,
	query QueryRetrieveAllEventType,
) (*EventTypesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(eventTypes), query)

	if err != nil {
		return nil, err
	}

	var response EventTypesResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEvents() (*EventsPageResponse, error) {
	return c.RetrieveEventsWithContext(context.Background())
}

func (c *Client) RetrieveEventsWithContext(ctx context.Context) (*EventsPageResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventsPageResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(events), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEventsQuery(query QueryRetrieveEvents) (*EventsPageResponse, error) {
	return c.RetrieveEventsQueryWithContext(context.Background(), query)
}

func (c *Client) RetrieveEventsQueryWithContext(ctx context.Context, query QueryRetrieveEvents) (*EventsPageResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(events), query)

	if err != nil {
		return nil, err
	}

	var response EventsPageResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEvent(eventId string) (*EventResponse, error) {
	return c.RetrieveEventWithContext(context.Background(), eventId)
}

func (c *Client) RetrieveEventWithContext(ctx context.Context, eventId string) (*EventResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(events, eventId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEventNotification(
	eventId string,
	notificationId string,
) (*EventNotificationResponse, error) {
	return c.RetrieveEventNotificationWithContext(context.Background(), eventId, notificationId)
}

func (c *Client) RetrieveEventNotificationWithContext(
	ctx context.Context,
	eventId string,
	notificationId string,
) (*EventNotificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventNotificationResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(events, eventId, notifications, notificationId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetryWebhook(eventId string, webhookId string) (*common.MetadataResponse, error) {
	return c.RetryWebhookWithContext(context.Background(), eventId, webhookId)
}

func (c *Client) RetryWebhookWithContext(ctx context.Context, eventId string, webhookId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(events, eventId, webhooks, webhookId),
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

func (c *Client) RetryAllWebhooks(eventId string) (*common.MetadataResponse, error) {
	return c.RetryAllWebhooksWithContext(context.Background(), eventId)
}

func (c *Client) RetryAllWebhooksWithContext(ctx context.Context, eventId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(events, eventId, webhooks, retry),
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
