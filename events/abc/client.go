package abc

import (
	"context"
	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
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
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(eventTypes), auth, nil, &response)
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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
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
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(events), auth, nil, &response)
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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetrieveEvent sends GET /events/{eventId}.
//
// eventId is the {eventId} path value.
//
// It returns the EventResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) RetrieveEvent(eventId string) (*EventResponse, error) {
	return c.RetrieveEventWithContext(context.Background(), eventId)
}

// RetrieveEventWithContext is RetrieveEvent with a caller-supplied context.
func (c *Client) RetrieveEventWithContext(ctx context.Context, eventId string) (*EventResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(events, eventId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetrieveEventNotification sends GET /events/{eventId}/notifications/{notificationId}.
//
// eventId is the {eventId} path value. notificationId is the {notificationId} path value.
//
// It returns the EventNotificationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId or
// notificationId returns a CheckoutArgumentError and sends no request.
func (c *Client) RetrieveEventNotification(
	eventId string,
	notificationId string,
) (*EventNotificationResponse, error) {
	return c.RetrieveEventNotificationWithContext(context.Background(), eventId, notificationId)
}

// RetrieveEventNotificationWithContext is RetrieveEventNotification with a caller-supplied context.
func (c *Client) RetrieveEventNotificationWithContext(
	ctx context.Context,
	eventId string,
	notificationId string,
) (*EventNotificationResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

	notificationId, err = common.EncodePathSegment(notificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventNotificationResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(events, eventId, notifications, notificationId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetryWebhook sends POST /events/{eventId}/webhooks/{webhookId}.
//
// eventId is the {eventId} path value. webhookId is the {webhookId} path value.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId or
// webhookId returns a CheckoutArgumentError and sends no request.
func (c *Client) RetryWebhook(eventId string, webhookId string) (*common.MetadataResponse, error) {
	return c.RetryWebhookWithContext(context.Background(), eventId, webhookId)
}

// RetryWebhookWithContext is RetryWebhook with a caller-supplied context.
func (c *Client) RetryWebhookWithContext(ctx context.Context, eventId string, webhookId string) (*common.MetadataResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

	webhookId, err = common.EncodePathSegment(webhookId)
	if err != nil {
		return nil, err
	}

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

// RetryAllWebhooks sends POST /events/{eventId}/webhooks/retry.
//
// eventId is the {eventId} path value.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) RetryAllWebhooks(eventId string) (*common.MetadataResponse, error) {
	return c.RetryAllWebhooksWithContext(context.Background(), eventId)
}

// RetryAllWebhooksWithContext is RetryAllWebhooks with a caller-supplied context.
func (c *Client) RetryAllWebhooksWithContext(ctx context.Context, eventId string) (*common.MetadataResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

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
