package customers

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

func (c *Client) Create(request CustomerRequest) (*common.IdResponse, error) {
	return c.CreateWithContext(context.Background(), request)
}

func (c *Client) CreateWithContext(ctx context.Context, request CustomerRequest) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(Path), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Get sends GET /customers/{identifier}: get customer details.
//
// customerId is the {identifier} path value: the customer's ID or email.
//
// It returns the GetCustomerResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." customerId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Get(customerId string) (*GetCustomerResponse, error) {
	return c.GetWithContext(context.Background(), customerId)
}

// GetWithContext is Get with a caller-supplied context.
func (c *Client) GetWithContext(ctx context.Context, customerId string) (*GetCustomerResponse, error) {
	customerId, err := common.EncodePathSegment(customerId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetCustomerResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(Path, customerId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Update sends PATCH /customers/{identifier}: update customer details.
//
// customerId is the {identifier} path value: the customer's ID. request is sent as the JSON body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." customerId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Update(customerId string, request CustomerRequest) (*common.MetadataResponse, error) {
	return c.UpdateWithContext(context.Background(), customerId, request)
}

// UpdateWithContext is Update with a caller-supplied context.
func (c *Client) UpdateWithContext(ctx context.Context, customerId string, request CustomerRequest) (*common.MetadataResponse, error) {
	customerId, err := common.EncodePathSegment(customerId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PatchWithContext(ctx, common.BuildPath(Path, customerId), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Delete sends DELETE /customers/{identifier}: delete a customer.
//
// customerId is the {identifier} path value: the customer's ID.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." customerId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Delete(customerId string) (*common.MetadataResponse, error) {
	return c.DeleteWithContext(context.Background(), customerId)
}

// DeleteWithContext is Delete with a caller-supplied context.
func (c *Client) DeleteWithContext(ctx context.Context, customerId string) (*common.MetadataResponse, error) {
	customerId, err := common.EncodePathSegment(customerId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(ctx, common.BuildPath(Path, customerId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
