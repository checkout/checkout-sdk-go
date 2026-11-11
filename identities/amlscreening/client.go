package amlscreening

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

func (c *Client) CreateAmlScreening(request CreateAmlScreeningRequest) (*AmlScreeningResponse, error) {
	return c.CreateAmlScreeningWithContext(context.Background(), request)
}

func (c *Client) CreateAmlScreeningWithContext(ctx context.Context, request CreateAmlScreeningRequest) (*AmlScreeningResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AmlScreeningResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(amlScreeningPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAmlScreening sends GET /aml-verifications/{aml_verification_id}: get an AML screening.
//
// screeningId is the {aml_verification_id} path value: the AML screening's unique identifier.
//
// It returns the AmlScreeningResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." screeningId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetAmlScreening(screeningId string) (*AmlScreeningResponse, error) {
	return c.GetAmlScreeningWithContext(context.Background(), screeningId)
}

// GetAmlScreeningWithContext is GetAmlScreening with a caller-supplied context.
func (c *Client) GetAmlScreeningWithContext(ctx context.Context, screeningId string) (*AmlScreeningResponse, error) {
	screeningId, err := common.EncodePathSegment(screeningId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AmlScreeningResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(amlScreeningPath, screeningId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
