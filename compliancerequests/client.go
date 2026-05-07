package compliancerequests

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

// Client holds the dependencies for making compliance requests API calls.
type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

// NewClient creates a compliance requests Client using the provided configuration and HTTP client.
func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

// GetComplianceRequest retrieves the compliance request associated with the given payment ID.
func (c *Client) GetComplianceRequest(paymentId string) (*GetComplianceRequestResponse, error) {
	return c.GetComplianceRequestWithContext(context.Background(), paymentId)
}

// GetComplianceRequestWithContext is like GetComplianceRequest but accepts a context for
// cancellation and deadline propagation.
func (c *Client) GetComplianceRequestWithContext(ctx context.Context, paymentId string) (*GetComplianceRequestResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetComplianceRequestResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(complianceRequestsPath, paymentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RespondToComplianceRequest submits the merchant's compliance response for the given payment ID.
func (c *Client) RespondToComplianceRequest(paymentId string, request RespondToComplianceRequestRequest) (*common.MetadataResponse, error) {
	return c.RespondToComplianceRequestWithContext(context.Background(), paymentId, request)
}

// RespondToComplianceRequestWithContext is like RespondToComplianceRequest but accepts a
// context for cancellation and deadline propagation.
func (c *Client) RespondToComplianceRequestWithContext(ctx context.Context, paymentId string, request RespondToComplianceRequestRequest) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(complianceRequestsPath, paymentId), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
