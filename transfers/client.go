package transfers

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

func NewClient(
	configuration *configuration.Configuration,
	apiClient client.HttpClient,
) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

// Deprecated: Use InitiateTransferOfFunds instead.
func (c *Client) InitiateTransferOfFounds(
	request TransferRequest,
	idempotencyKey *string,
) (*TransferResponse, error) {
	return c.InitiateTransferOfFundsWithContext(context.Background(), request, idempotencyKey)
}

// Deprecated: Use InitiateTransferOfFundsWithContext instead.
func (c *Client) InitiateTransferOfFoundsWithContext(
	ctx context.Context,
	request TransferRequest,
	idempotencyKey *string,
) (*TransferResponse, error) {
	return c.InitiateTransferOfFundsWithContext(ctx, request, idempotencyKey)
}

func (c *Client) InitiateTransferOfFunds(
	request TransferRequest,
	idempotencyKey *string,
) (*TransferResponse, error) {
	return c.InitiateTransferOfFundsWithContext(context.Background(), request, idempotencyKey)
}

func (c *Client) InitiateTransferOfFundsWithContext(
	ctx context.Context,
	request TransferRequest,
	idempotencyKey *string,
) (*TransferResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response TransferResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(transfers),
		auth,
		request,
		&response,
		idempotencyKey,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetrieveTransfer sends GET /transfers/{id}: retrieve a transfer.
//
// transferId is the {id} path value: the transfer identifier.
//
// It returns the TransferDetails, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." transferId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) RetrieveTransfer(transferId string) (*TransferDetails, error) {
	return c.RetrieveTransferWithContext(context.Background(), transferId)
}

// RetrieveTransferWithContext is RetrieveTransfer with a caller-supplied context.
func (c *Client) RetrieveTransferWithContext(
	ctx context.Context,
	transferId string,
) (*TransferDetails, error) {
	transferId, err := common.EncodePathSegment(transferId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response TransferDetails
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(transfers, transferId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
