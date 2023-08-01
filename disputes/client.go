package disputes

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

func (c *Client) Query(queryFilter QueryFilter) (*QueryResponse, error) {
	return c.QueryWithContext(context.Background(), queryFilter)
}

func (c *Client) QueryWithContext(ctx context.Context, queryFilter QueryFilter) (*QueryResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(disputes), queryFilter)
	if err != nil {
		return nil, err
	}

	var response QueryResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDisputeDetails(disputeId string) (*DisputeResponse, error) {
	return c.GetDisputeDetailsWithContext(context.Background(), disputeId)
}

func (c *Client) GetDisputeDetailsWithContext(ctx context.Context, disputeId string) (*DisputeResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response DisputeResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) Accept(disputeId string) (*common.MetadataResponse, error) {
	return c.AcceptWithContext(context.Background(), disputeId)
}

func (c *Client) AcceptWithContext(ctx context.Context, disputeId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(disputes, disputeId, accept),
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
	return c.PutEvidenceWithContext(context.Background(), disputeId, evidenceRequest)
}

func (c *Client) PutEvidenceWithContext(ctx context.Context, disputeId string, evidenceRequest Evidence) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(disputes, disputeId, evidence),
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
	return c.GetEvidenceWithContext(context.Background(), disputeId)
}

func (c *Client) GetEvidenceWithContext(ctx context.Context, disputeId string) (*EvidenceResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response EvidenceResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId, evidence), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) SubmitEvidence(disputeId string) (*common.MetadataResponse, error) {
	return c.SubmitEvidenceWithContext(context.Background(), disputeId)
}

func (c *Client) SubmitEvidenceWithContext(ctx context.Context, disputeId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(disputes, disputeId, evidence),
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
	return c.UploadFileWithContext(context.Background(), file)
}

func (c *Client) UploadFileWithContext(ctx context.Context, file common.File) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	req, err := common.BuildFileUploadRequest(&file)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.UploadWithContext(ctx, common.BuildPath(files), auth, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFileDetails(fileId string) (*common.FileResponse, error) {
	return c.GetFileDetailsWithContext(context.Background(), fileId)
}

func (c *Client) GetFileDetailsWithContext(ctx context.Context, fileId string) (*common.FileResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.FileResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(files, fileId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDisputeSchemeFiles(disputeId string) (*SchemeFilesResponse, error) {
	return c.GetDisputeSchemeFilesWithContext(context.Background(), disputeId)
}

func (c *Client) GetDisputeSchemeFilesWithContext(ctx context.Context, disputeId string) (*SchemeFilesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response SchemeFilesResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId, schemeFiles), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
