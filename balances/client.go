package balances

import (
	"context"
	"strings"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
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

func (c *Client) RetrieveEntityBalances(entityId string, query QueryFilter) (*QueryResponse, error) {
	return c.RetrieveEntityBalancesWithContext(context.Background(), entityId, query)
}

func (c *Client) RetrieveEntityBalancesWithContext(
	ctx context.Context,
	entityId string,
	query QueryFilter,
) (*QueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(balances, entityId), query)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	err = c.apiClient.GetWithContext(
		ctx,
		url,
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RetrieveTopUpInstructions retrieves the bank details required to top up a sub-account, along
// with the payment reference that attributes an incoming payment to that sub-account.
//
// Note: The sub-account is referred to as currency account in the API.
//
// GET /entities/{entityId}/currency-accounts/{currencyAccountId}/top-up-instructions
//
// entityId is the ID of the entity that owns the sub-account, or of an entity above it in your
// hierarchy; a platform can use its own entity ID to reach the sub-accounts of any entity beneath
// it. currencyAccountId is the ID of the sub-account to retrieve top-up instructions for.
//
// Both arguments are required. A blank value for either returns a CheckoutArgumentError without
// making a request, matching the java and .NET SDKs' validateParams semantics.
func (c *Client) RetrieveTopUpInstructions(entityId, currencyAccountId string) (*TopUpInstructionsResponse, error) {
	return c.RetrieveTopUpInstructionsWithContext(context.Background(), entityId, currencyAccountId)
}

// RetrieveTopUpInstructionsWithContext is RetrieveTopUpInstructions with a caller-supplied context.
func (c *Client) RetrieveTopUpInstructionsWithContext(
	ctx context.Context,
	entityId, currencyAccountId string,
) (*TopUpInstructionsResponse, error) {
	if strings.TrimSpace(entityId) == "" {
		return nil, errors.CheckoutArgumentError("entityId cannot be blank")
	}
	if strings.TrimSpace(currencyAccountId) == "" {
		return nil, errors.CheckoutArgumentError("currencyAccountId cannot be blank")
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response TopUpInstructionsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(entities, entityId, currencyAccounts, currencyAccountId, topUpInstructions),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
