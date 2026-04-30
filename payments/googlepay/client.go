package googlepay

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

func (c *Client) CreateEnrollment(request CreateEnrollmentRequest) (*CreateEnrollmentResponse, error) {
	return c.CreateEnrollmentWithContext(context.Background(), request)
}

func (c *Client) CreateEnrollmentWithContext(ctx context.Context, request CreateEnrollmentRequest) (*CreateEnrollmentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response CreateEnrollmentResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(GooglePayEnrollmentsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RegisterDomain(entityId string, request RegisterDomainRequest) (*common.MetadataResponse, error) {
	return c.RegisterDomainWithContext(context.Background(), entityId, request)
}

func (c *Client) RegisterDomainWithContext(ctx context.Context, entityId string, request RegisterDomainRequest) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(GooglePayEnrollmentsPath, entityId, domainPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetRegisteredDomains(entityId string) (*DomainListResponse, error) {
	return c.GetRegisteredDomainsWithContext(context.Background(), entityId)
}

func (c *Client) GetRegisteredDomainsWithContext(ctx context.Context, entityId string) (*DomainListResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response DomainListResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(GooglePayEnrollmentsPath, entityId, domainsPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetEnrollmentState(entityId string) (*EnrollmentStateResponse, error) {
	return c.GetEnrollmentStateWithContext(context.Background(), entityId)
}

func (c *Client) GetEnrollmentStateWithContext(ctx context.Context, entityId string) (*EnrollmentStateResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response EnrollmentStateResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(GooglePayEnrollmentsPath, entityId, statePath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
