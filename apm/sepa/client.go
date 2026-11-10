package sepa

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

// GetMandate sends GET /sepa/mandates/{mandateId}.
//
// mandateId is the {mandateId} path value.
//
// It returns the MandateResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." mandateId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetMandate(mandateId string) (*MandateResponse, error) {
	return c.GetMandateWithContext(context.Background(), mandateId)
}

// GetMandateWithContext is GetMandate with a caller-supplied context.
func (c *Client) GetMandateWithContext(ctx context.Context, mandateId string) (*MandateResponse, error) {
	mandateId, err := common.EncodePathSegment(mandateId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response MandateResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(sepaMandatesPath, mandateId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CancelMandate sends POST /sepa/mandates/{mandateId}/cancel.
//
// mandateId is the {mandateId} path value.
//
// It returns the SepaResource, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." mandateId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CancelMandate(mandateId string) (*SepaResource, error) {
	return c.CancelMandateWithContext(context.Background(), mandateId)
}

// CancelMandateWithContext is CancelMandate with a caller-supplied context.
func (c *Client) CancelMandateWithContext(ctx context.Context, mandateId string) (*SepaResource, error) {
	mandateId, err := common.EncodePathSegment(mandateId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response SepaResource
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(sepaMandatesPath, mandateId, cancelPath),
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

// GetMandateViaPpro sends GET /apms/ppro/sepa/mandates/{mandateId}.
//
// mandateId is the {mandateId} path value.
//
// It returns the MandateResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." mandateId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetMandateViaPpro(mandateId string) (*MandateResponse, error) {
	return c.GetMandateViaPproWithContext(context.Background(), mandateId)
}

// GetMandateViaPproWithContext is GetMandateViaPpro with a caller-supplied context.
func (c *Client) GetMandateViaPproWithContext(ctx context.Context, mandateId string) (*MandateResponse, error) {
	mandateId, err := common.EncodePathSegment(mandateId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response MandateResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(apmsPath, pproPath, sepaMandatesPath, mandateId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CancelMandateViaPpro sends POST /apms/ppro/sepa/mandates/{mandateId}/cancel.
//
// mandateId is the {mandateId} path value.
//
// It returns the SepaResource, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." mandateId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CancelMandateViaPpro(mandateId string) (*SepaResource, error) {
	return c.CancelMandateViaPproWithContext(context.Background(), mandateId)
}

// CancelMandateViaPproWithContext is CancelMandateViaPpro with a caller-supplied context.
func (c *Client) CancelMandateViaPproWithContext(ctx context.Context, mandateId string) (*SepaResource, error) {
	mandateId, err := common.EncodePathSegment(mandateId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response SepaResource
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(apmsPath, pproPath, sepaMandatesPath, mandateId, cancelPath),
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
