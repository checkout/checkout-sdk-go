package reports

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

func (c *Client) GetAllReports(query QueryFilter) (*QueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(reports), query)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	err = c.apiClient.Get(url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetReportDetails(reportId string) (*ReportResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response ReportResponse
	err = c.apiClient.Get(common.BuildPath(reports, reportId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetReportFile(reportId, fileId string) (*common.ContentResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.ContentResponse
	err = c.apiClient.Get(common.BuildPath(reports, reportId, files, fileId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
