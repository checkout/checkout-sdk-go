package bacs

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

// SendNotification sends a Bacs Direct Debit pre-notification (advance notice) to a payer ahead of
// collecting funds from their account. Calls POST /apms/bacs/notifications.
func (c *Client) SendNotification(request NotificationRequest) (*NotificationResponse, error) {
	return c.SendNotificationWithContext(context.Background(), request)
}

func (c *Client) SendNotificationWithContext(
	ctx context.Context,
	request NotificationRequest,
) (*NotificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response NotificationResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(apmsPath, bacsPath, notificationsPath),
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
