package transfers

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

func (c *Client) RetrieveTransfer(transferId string) (*TransferDetails, error) {
	return c.RetrieveTransferWithContext(context.Background(), transferId)
}

func (c *Client) RetrieveTransferWithContext(
	ctx context.Context,
	transferId string,
) (*TransferDetails, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response TransferDetails
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(transfers, transferId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
