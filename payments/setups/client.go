package setups

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

func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

// CreatePaymentSetup creates a Payment Setup.
// Beta
//
// Creates a Payment Setup.
//
// To maximize the amount of information the payment setup can use, we
// recommend that you create a payment setup as early as possible in the
// customer's journey. For example, the first time they land on the basket
// page.
func (c *Client) CreatePaymentSetup(request PaymentSetupRequest) (*PaymentSetupResponse, error) {
	return c.CreatePaymentSetupWithContext(context.Background(), request)
}

func (c *Client) CreatePaymentSetupWithContext(ctx context.Context, request PaymentSetupRequest) (*PaymentSetupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(PaymentSetupsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdatePaymentSetup updates a Payment Setup.
// Beta
//
// Updates a Payment Setup.
//
// You should update the payment setup whenever there are significant changes
// in the data relevant to the customer's transaction. For example, when the
// customer makes a change that impacts the total payment amount.
func (c *Client) UpdatePaymentSetup(setupId string, request PaymentSetupRequest) (*PaymentSetupResponse, error) {
	return c.UpdatePaymentSetupWithContext(context.Background(), setupId, request)
}

func (c *Client) UpdatePaymentSetupWithContext(ctx context.Context, setupId string, request PaymentSetupRequest) (*PaymentSetupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupResponse
	err = c.apiClient.PutWithContext(ctx, common.BuildPath(PaymentSetupsPath, setupId), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetPaymentSetup retrieves a Payment Setup.
// Beta
//
// Retrieves a Payment Setup.
func (c *Client) GetPaymentSetup(setupId string) (*PaymentSetupResponse, error) {
	return c.GetPaymentSetupWithContext(context.Background(), setupId)
}

func (c *Client) GetPaymentSetupWithContext(ctx context.Context, setupId string) (*PaymentSetupResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(PaymentSetupsPath, setupId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// ConfirmPaymentSetup confirms a Payment Setup to begin processing the
// payment request with your chosen payment method option.
// Beta
//
// Confirm a Payment Setup to begin processing the payment request with your
// chosen payment method option.
func (c *Client) ConfirmPaymentSetup(setupId string, paymentMethodOptionId string) (*PaymentSetupConfirmResponse, error) {
	return c.ConfirmPaymentSetupWithContext(context.Background(), setupId, paymentMethodOptionId)
}

func (c *Client) ConfirmPaymentSetupWithContext(ctx context.Context, setupId string, paymentMethodOptionId string) (*PaymentSetupConfirmResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response PaymentSetupConfirmResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(PaymentSetupsPath, setupId, ConfirmPath, paymentMethodOptionId), auth, nil, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
