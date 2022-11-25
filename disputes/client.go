package disputes

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

func (c *Client) Query(queryFilter QueryFilter) (*QueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(path, queryFilter)
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

func (c *Client) GetDisputeDetails(disputeId string) (*DisputeResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response DisputeResponse
	err = c.apiClient.Get(common.BuildPath(path, disputeId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Accept(disputeId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
		common.BuildPath(path, disputeId, accept),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) PutEvidence(disputeId string, evidenceRequest Evidence) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Put(
		common.BuildPath(path, disputeId, evidence),
		auth,
		evidenceRequest,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetEvidence(disputeId string) (*EvidenceResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response EvidenceResponse
	err = c.apiClient.Get(common.BuildPath(path, disputeId, evidence), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SubmitEvidence(disputeId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
		common.BuildPath(path, disputeId, evidence),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UploadFile(file common.File) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	req, err := common.BuildFileUploadRequest(&file)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Upload(common.BuildPath(files), auth, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFileDetails(fileId string) (*common.FileResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.FileResponse
	err = c.apiClient.Get(common.BuildPath(files, fileId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
