package nas

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

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
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

// GetBankAccountFieldFormatting sends GET /validation/bank-accounts/{country}/{currency}: get bank
// account field formatting.
//
// country is the {country} path value: the two-letter ISO country code. currency is the {currency}
// path value: the three-letter ISO currency code. query is sent as the query string.
//
// It returns the GetBankAccountFieldFormattingResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." country or
// currency returns a CheckoutArgumentError and sends no request.
func (c *Client) GetBankAccountFieldFormatting(
	country string,
	currency string,
	query QueryBankAccountFormatting,
) (*GetBankAccountFieldFormattingResponse, error) {
	return c.GetBankAccountFieldFormattingWithContext(context.Background(), country, currency, query)
}

// GetBankAccountFieldFormattingWithContext is GetBankAccountFieldFormatting with a caller-supplied context.
func (c *Client) GetBankAccountFieldFormattingWithContext(
	ctx context.Context,
	country string,
	currency string,
	query QueryBankAccountFormatting,
) (*GetBankAccountFieldFormattingResponse, error) {
	country, err := common.EncodePathSegment(country)
	if err != nil {
		return nil, err
	}

	currency, err = common.EncodePathSegment(currency)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(instruments.ValidationPath, country, currency), query)
	if err != nil {
		return nil, err
	}

	var response GetBankAccountFieldFormattingResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
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
func (c *Client) UpdateWithContext(
	ctx context.Context,
	instrumentId string,
	request UpdateInstrumentRequest,
) (*UpdateInstrumentResponse, error) {
	instrumentId, err := common.EncodePathSegment(instrumentId)
	if err != nil {
		return nil, err
	}

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

// RevokeInstrument sends PATCH /instruments/{instrumentId}/revoke.
//
// instrumentId is the {instrumentId} path value.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." instrumentId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) RevokeInstrument(instrumentId string) (*common.MetadataResponse, error) {
	return c.RevokeInstrumentWithContext(context.Background(), instrumentId)
}

// RevokeInstrumentWithContext is RevokeInstrument with a caller-supplied context.
func (c *Client) RevokeInstrumentWithContext(ctx context.Context, instrumentId string) (*common.MetadataResponse, error) {
	instrumentId, err := common.EncodePathSegment(instrumentId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PatchWithContext(ctx, common.BuildPath(instruments.Path, instrumentId, "revoke"), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
