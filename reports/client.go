package reports

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

func (c *Client) GetAllReports(query QueryFilter) (*QueryResponse, error) {
	return c.GetAllReportsWithContext(context.Background(), query)
}

func (c *Client) GetAllReportsWithContext(ctx context.Context, query QueryFilter) (*QueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reports), query)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetReportDetails sends GET /reports/{id}: get report details.
//
// reportId is the {id} path value: the ID of the report to retrieve.
//
// It returns the ReportResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." reportId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetReportDetails(reportId string) (*ReportResponse, error) {
	return c.GetReportDetailsWithContext(context.Background(), reportId)
}

// GetReportDetailsWithContext is GetReportDetails with a caller-supplied context.
func (c *Client) GetReportDetailsWithContext(ctx context.Context, reportId string) (*ReportResponse, error) {
	reportId, err := common.EncodePathSegment(reportId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ReportResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(reports, reportId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetReportFile sends GET /reports/{id}/files/{fileId}: get report file.
//
// reportId is the {id} path value: the ID of the report that the file belongs to. fileId is the
// {fileId} path value: the ID of the file to retrieve.
//
// It returns the common.ContentResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." reportId or
// fileId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetReportFile(reportId, fileId string) (*common.ContentResponse, error) {
	return c.GetReportFileWithContext(context.Background(), reportId, fileId)
}

// GetReportFileWithContext is GetReportFile with a caller-supplied context.
func (c *Client) GetReportFileWithContext(ctx context.Context, reportId, fileId string) (*common.ContentResponse, error) {
	reportId, err := common.EncodePathSegment(reportId)
	if err != nil {
		return nil, err
	}

	fileId, err = common.EncodePathSegment(fileId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(reports, reportId, files, fileId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
