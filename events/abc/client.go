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

func (c *Client) RetrieveAllEventTypes() (*EventTypesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventTypesResponse
	err = c.apiClient.Get(common.BuildPath(eventTypes), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveAllEventTypesQuery(
	query QueryRetrieveAllEventType,
) (*EventTypesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(eventTypes, query)

	if err != nil {
		return nil, err
	}

	var response EventTypesResponse
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEvents() (*EventsPageResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventsPageResponse
	err = c.apiClient.Get(common.BuildPath(events), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEventsQuery(query QueryRetrieveEvents) (*EventsPageResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(events, query)

	if err != nil {
		return nil, err
	}

	var response EventsPageResponse
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEvent(eventId string) (*EventResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventResponse
	err = c.apiClient.Get(common.BuildPath(events, eventId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveEventNotification(
	eventId string,
	notificationId string,
) (*EventNotificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response EventNotificationResponse
	err = c.apiClient.Get(common.BuildPath(events, eventId, notifications, notificationId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetryWebhook(eventId string, webhookId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(common.BuildPath(events, eventId, webhooks, retry),
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
