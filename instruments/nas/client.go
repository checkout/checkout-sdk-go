package nas

import (
	"context"

	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/instruments"
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

func (c *Client) CreateWithContext(
	ctx context.Context,
	request CreateInstrumentRequest,
) (*CreateInstrumentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response CreateInstrumentResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(instruments.Path), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Get(instrumentId string) (*GetInstrumentResponse, error) {
	return c.GetWithContext(context.Background(), instrumentId)
}

func (c *Client) GetWithContext(ctx context.Context, instrumentId string) (*GetInstrumentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetInstrumentResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(instruments.Path, instrumentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetBankAccountFieldFormatting(
	country string,
	currency string,
	query QueryBankAccountFormatting,
) (*GetBankAccountFieldFormattingResponse, error) {
	return c.GetBankAccountFieldFormattingWithContext(context.Background(), country, currency, query)
}

func (c *Client) GetBankAccountFieldFormattingWithContext(
	ctx context.Context,
	country string,
	currency string,
	query QueryBankAccountFormatting,
) (*GetBankAccountFieldFormattingResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(instruments.ValidationPath, country, currency), query)
	if err != nil {
		return nil, err
	}

	var response GetBankAccountFieldFormattingResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Update(instrumentId string, request UpdateInstrumentRequest) (*UpdateInstrumentResponse, error) {
	return c.UpdateWithContext(context.Background(), instrumentId, request)
}

func (c *Client) UpdateWithContext(
	ctx context.Context,
	instrumentId string,
	request UpdateInstrumentRequest,
) (*UpdateInstrumentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response UpdateInstrumentResponse
	err = c.apiClient.PatchWithContext(ctx, common.BuildPath(instruments.Path, instrumentId), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Delete(instrumentId string) (*common.MetadataResponse, error) {
	return c.DeleteWithContext(context.Background(), instrumentId)
}

func (c *Client) DeleteWithContext(ctx context.Context, instrumentId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
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
