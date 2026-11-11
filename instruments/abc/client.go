package abc

import (
	"context"
	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/instruments"
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

func (c *Client) Create(request CreateInstrumentRequest) (*CreateInstrumentResponse, error) {
	return c.CreateWithContext(context.Background(), request)
}

func (c *Client) CreateWithContext(ctx context.Context, request CreateInstrumentRequest) (*CreateInstrumentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response CreateInstrumentResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(instruments.Path),
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

// Get sends GET /instruments/{id}: get instrument details.
//
// instrumentId is the {id} path value: the instrument ID.
//
// It returns the GetInstrumentResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." instrumentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Get(instrumentId string) (*GetInstrumentResponse, error) {
	return c.GetWithContext(context.Background(), instrumentId)
}

// GetWithContext is Get with a caller-supplied context.
func (c *Client) GetWithContext(ctx context.Context, instrumentId string) (*GetInstrumentResponse, error) {
	instrumentId, err := common.EncodePathSegment(instrumentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response GetInstrumentResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(instruments.Path, instrumentId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Update sends PATCH /instruments/{id}: update an instrument.
//
// instrumentId is the {id} path value: the instrument ID. request is sent as the JSON body.
//
// It returns the UpdateInstrumentResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." instrumentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Update(instrumentId string, request UpdateInstrumentRequest) (*UpdateInstrumentResponse, error) {
	return c.UpdateWithContext(context.Background(), instrumentId, request)
}

// UpdateWithContext is Update with a caller-supplied context.
func (c *Client) UpdateWithContext(ctx context.Context, instrumentId string, request UpdateInstrumentRequest) (*UpdateInstrumentResponse, error) {
	instrumentId, err := common.EncodePathSegment(instrumentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response UpdateInstrumentResponse
	err = c.apiClient.PatchWithContext(
		ctx,
		common.BuildPath(instruments.Path, instrumentId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Delete sends DELETE /instruments/{id}: delete an instrument.
//
// instrumentId is the {id} path value: the ID of the payment instrument to be deleted.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." instrumentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Delete(instrumentId string) (*common.MetadataResponse, error) {
	return c.DeleteWithContext(context.Background(), instrumentId)
}

// DeleteWithContext is Delete with a caller-supplied context.
func (c *Client) DeleteWithContext(ctx context.Context, instrumentId string) (*common.MetadataResponse, error) {
	instrumentId, err := common.EncodePathSegment(instrumentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(ctx, common.BuildPath(instruments.Path, instrumentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
