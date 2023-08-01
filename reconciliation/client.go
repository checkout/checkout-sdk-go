package reconciliation

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

func (c *Client) QueryPaymentsReport(query PaymentReportsQuery) (*PaymentReportsResponse, error) {
	return c.QueryPaymentsReportWithContext(context.Background(), query)
}

func (c *Client) QueryPaymentsReportWithContext(ctx context.Context, query PaymentReportsQuery) (*PaymentReportsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, payments), query)
	if err != nil {
		return nil, err
	}

	var response PaymentReportsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetSinglePaymentReport(paymentId string) (*PaymentReportsResponse, error) {
	return c.GetSinglePaymentReportWithContext(context.Background(), paymentId)
}

func (c *Client) GetSinglePaymentReportWithContext(ctx context.Context, paymentId string) (*PaymentReportsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response PaymentReportsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(reporting, payments, paymentId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) QueryStatementsReport(query common.DateRangeQuery) (*StatementReportsResponse, error) {
	return c.QueryStatementsReportWithContext(context.Background(), query)
}

func (c *Client) QueryStatementsReportWithContext(ctx context.Context, query common.DateRangeQuery) (*StatementReportsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, statements), query)
	if err != nil {
		return nil, err
	}

	var response StatementReportsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveCVSPaymentsReport(query common.DateRangeQuery) (*common.ContentResponse, error) {
	return c.RetrieveCVSPaymentsReportWithContext(context.Background(), query)
}

func (c *Client) RetrieveCVSPaymentsReportWithContext(ctx context.Context, query common.DateRangeQuery) (*common.ContentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, payments, download), query)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveCVSSingleStatementReport(statementId string) (*common.ContentResponse, error) {
	return c.RetrieveCVSSingleStatementReportWithContext(context.Background(), statementId)
}

func (c *Client) RetrieveCVSSingleStatementReportWithContext(ctx context.Context, statementId string) (*common.ContentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(reporting, statements, statementId, payments, download), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RetrieveCVSStatementsReport(query common.DateRangeQuery) (*common.ContentResponse, error) {
	return c.RetrieveCVSStatementsReportWithContext(context.Background(), query)
}

func (c *Client) RetrieveCVSStatementsReportWithContext(ctx context.Context, query common.DateRangeQuery) (*common.ContentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKey)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reporting, statements, download), query)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
