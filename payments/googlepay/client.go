package googlepay

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

// RegisterDomain sends POST /googlepay/enrollments/{entity_id}/domain: register a web domain for an
// enrolled entity.
//
// entityId is the {entity_id} path value: unique identifier of the entity. request is sent as the
// JSON body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) RegisterDomain(entityId string, request RegisterDomainRequest) (*common.MetadataResponse, error) {
	return c.RegisterDomainWithContext(context.Background(), entityId, request)
}

// RegisterDomainWithContext is RegisterDomain with a caller-supplied context.
func (c *Client) RegisterDomainWithContext(ctx context.Context, entityId string, request RegisterDomainRequest) (*common.MetadataResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

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

// GetRegisteredDomains sends GET /googlepay/enrollments/{entity_id}/domains: get registered domains
// for an entity.
//
// entityId is the {entity_id} path value: unique identifier of the entity.
//
// It returns the DomainListResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetRegisteredDomains(entityId string) (*DomainListResponse, error) {
	return c.GetRegisteredDomainsWithContext(context.Background(), entityId)
}

// GetRegisteredDomainsWithContext is GetRegisteredDomains with a caller-supplied context.
func (c *Client) GetRegisteredDomainsWithContext(ctx context.Context, entityId string) (*DomainListResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response DomainListResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(GooglePayEnrollmentsPath, entityId, domainsPath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetEnrollmentState sends GET /googlepay/enrollments/{entity_id}/state: get enrollment state.
//
// entityId is the {entity_id} path value: unique identifier of the entity.
//
// It returns the EnrollmentStateResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetEnrollmentState(entityId string) (*EnrollmentStateResponse, error) {
	return c.GetEnrollmentStateWithContext(context.Background(), entityId)
}

// GetEnrollmentStateWithContext is GetEnrollmentState with a caller-supplied context.
func (c *Client) GetEnrollmentStateWithContext(ctx context.Context, entityId string) (*EnrollmentStateResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response EnrollmentStateResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(GooglePayEnrollmentsPath, entityId, statePath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
