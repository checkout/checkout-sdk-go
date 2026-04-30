package compliancerequests

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
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

func (c *Client) GetComplianceRequest(paymentId string) (*GetComplianceRequestResponse, error) {
	return c.GetComplianceRequestWithContext(context.Background(), paymentId)
}

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

func (c *Client) RespondToComplianceRequest(paymentId string, request RespondToComplianceRequestRequest) (*common.MetadataResponse, error) {
	return c.RespondToComplianceRequestWithContext(context.Background(), paymentId, request)
}

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
