package reconciliation

import (
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

func (c *Client) QueryPaymentsReport(query PaymentReportsQuery) (*PaymentReportsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, payments), query)
	if err != nil {
		return nil, err
	}

	var response PaymentReportsResponse
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetSinglePaymentReport(paymentId string) (*PaymentReportsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentReportsResponse
	err = c.apiClient.Get(common.BuildPath(reporting, payments, paymentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) QueryStatementsReport(query common.DateRangeQuery) (*StatementReportsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, statements), query)
	if err != nil {
		return nil, err
	}

	var response StatementReportsResponse
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveCVSPaymentsReport(query common.DateRangeQuery) (*common.ContentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, payments, download), query)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveCVSSingleStatementReport(statementId string) (*common.ContentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.Get(common.BuildPath(reporting, statements, statementId, payments, download), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveCVSStatementsReport(query common.DateRangeQuery) (*common.ContentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, statements, download), query)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
