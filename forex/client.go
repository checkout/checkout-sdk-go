package forex

import (
	"context"

	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
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

func (c *Client) RequestQuote(request QuoteRequest) (*QuoteResponse, error) {
	return c.RequestQuoteWithContext(context.Background(), request)
}

func (c *Client) RequestQuoteWithContext(ctx context.Context, request QuoteRequest) (*QuoteResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response QuoteResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(forex, quotes),
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

func (c *Client) GetRates(queryFilter RatesQuery) (*RatesResponse, error) {
	return c.GetRatesWithContext(context.Background(), queryFilter)
}

func (c *Client) GetRatesWithContext(ctx context.Context, queryFilter RatesQuery) (*RatesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(forex, rates), queryFilter)
	if err != nil {
		return nil, err
	}

	var response RatesResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
