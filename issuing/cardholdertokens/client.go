package cardholdertokens

import (
	"context"
	"net/url"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

// Client holds the dependencies for making cardholder token API requests.
type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

// NewClient creates a cardholder tokens Client using the provided configuration and HTTP client.
func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

// RequestCardholderToken issues a short-lived access token bound to a specific cardholder.
// The endpoint is OAuth-style: credentials are passed in the form body, not in an Authorization
// header, and the request uses application/x-www-form-urlencoded encoding.
func (c *Client) RequestCardholderToken(request CardholderTokenRequest) (*CardholderTokenResponse, error) {
	return c.RequestCardholderTokenWithContext(context.Background(), request)
}

// RequestCardholderTokenWithContext is like RequestCardholderToken but accepts a context for
// cancellation and deadline propagation.
func (c *Client) RequestCardholderTokenWithContext(ctx context.Context, request CardholderTokenRequest) (*CardholderTokenResponse, error) {
	formData := url.Values{}
	formData.Set("grant_type", request.GrantType)
	formData.Set("client_id", request.ClientId)
	formData.Set("client_secret", request.ClientSecret)
	formData.Set("cardholder_id", request.CardholderId)
	if request.SingleUse {
		formData.Set("single_use", "true")
	}

	var response CardholderTokenResponse
	err := c.apiClient.PostFormWithContext(ctx, common.BuildPath(cardholderTokenPath), nil, formData, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
