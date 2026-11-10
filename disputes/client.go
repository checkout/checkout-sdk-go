package disputes

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
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetDisputeDetails sends GET /disputes/{dispute_id}: get dispute details.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the DisputeResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetDisputeDetails(disputeId string) (*DisputeResponse, error) {
	return c.GetDisputeDetailsWithContext(context.Background(), disputeId)
}

// GetDisputeDetailsWithContext is GetDisputeDetails with a caller-supplied context.
func (c *Client) GetDisputeDetailsWithContext(ctx context.Context, disputeId string) (*DisputeResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response DisputeResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Accept sends POST /disputes/{dispute_id}/accept: accept dispute.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Accept(disputeId string) (*common.MetadataResponse, error) {
	return c.AcceptWithContext(context.Background(), disputeId)
}

// AcceptWithContext is Accept with a caller-supplied context.
func (c *Client) AcceptWithContext(ctx context.Context, disputeId string) (*common.MetadataResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

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

// PutEvidence sends PUT /disputes/{dispute_id}/evidence: provide dispute evidence.
//
// disputeId is the {dispute_id} path value: the dispute identifier. evidenceRequest is sent as the
// JSON body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) PutEvidence(disputeId string, evidenceRequest Evidence) (*common.MetadataResponse, error) {
	return c.PutEvidenceWithContext(context.Background(), disputeId, evidenceRequest)
}

// PutEvidenceWithContext is PutEvidence with a caller-supplied context.
func (c *Client) PutEvidenceWithContext(ctx context.Context, disputeId string, evidenceRequest Evidence) (*common.MetadataResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

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

// GetEvidence sends GET /disputes/{dispute_id}/evidence: get dispute evidence.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the EvidenceResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetEvidence(disputeId string) (*EvidenceResponse, error) {
	return c.GetEvidenceWithContext(context.Background(), disputeId)
}

// GetEvidenceWithContext is GetEvidence with a caller-supplied context.
func (c *Client) GetEvidenceWithContext(ctx context.Context, disputeId string) (*EvidenceResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response EvidenceResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId, evidence), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SubmitEvidence sends POST /disputes/{dispute_id}/evidence: submit dispute evidence.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) SubmitEvidence(disputeId string) (*common.MetadataResponse, error) {
	return c.SubmitEvidenceWithContext(context.Background(), disputeId)
}

// SubmitEvidenceWithContext is SubmitEvidence with a caller-supplied context.
func (c *Client) SubmitEvidenceWithContext(ctx context.Context, disputeId string) (*common.MetadataResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

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

// GetCompiledSubmittedEvidence sends GET /disputes/{dispute_id}/evidence/submitted: get dispute
// submitted evidence.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the DisputeCompiledSubmittedEvidenceResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetCompiledSubmittedEvidence(disputeId string) (*DisputeCompiledSubmittedEvidenceResponse, error) {
	return c.GetCompiledSubmittedEvidenceWithContext(context.Background(), disputeId)
}

// GetCompiledSubmittedEvidenceWithContext is GetCompiledSubmittedEvidence with a caller-supplied context.
func (c *Client) GetCompiledSubmittedEvidenceWithContext(ctx context.Context, disputeId string) (*DisputeCompiledSubmittedEvidenceResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response DisputeCompiledSubmittedEvidenceResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId, evidence, submitted), auth, nil, &response)
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

// GetFileDetails sends GET /files/{file_id}: get file information.
//
// fileId is the {file_id} path value: the file identifier. It is always prefixed by file_.
//
// It returns the common.FileResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." fileId returns a
// CheckoutArgumentError and sends no request.
func (c *Client) GetFileDetails(fileId string) (*common.FileResponse, error) {
	return c.GetFileDetailsWithContext(context.Background(), fileId)
}

// GetFileDetailsWithContext is GetFileDetails with a caller-supplied context.
func (c *Client) GetFileDetailsWithContext(ctx context.Context, fileId string) (*common.FileResponse, error) {
	fileId, err := common.EncodePathSegment(fileId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.FileResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(files, fileId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// SubmitArbitrationEvidence sends POST /disputes/{dispute_id}/evidence/arbitration: submit dispute
// arbitration evidence.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) SubmitArbitrationEvidence(disputeId string) (*common.MetadataResponse, error) {
	return c.SubmitArbitrationEvidenceWithContext(context.Background(), disputeId)
}

// SubmitArbitrationEvidenceWithContext is SubmitArbitrationEvidence with a caller-supplied context.
func (c *Client) SubmitArbitrationEvidenceWithContext(ctx context.Context, disputeId string) (*common.MetadataResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(disputes, disputeId, evidence, arbitration),
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

// GetSubmittedArbitrationEvidence sends GET /disputes/{dispute_id}/evidence/arbitration/submitted:
// get dispute submitted arbitration evidence.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the DisputeCompiledSubmittedEvidenceResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetSubmittedArbitrationEvidence(disputeId string) (*DisputeCompiledSubmittedEvidenceResponse, error) {
	return c.GetSubmittedArbitrationEvidenceWithContext(context.Background(), disputeId)
}

// GetSubmittedArbitrationEvidenceWithContext is GetSubmittedArbitrationEvidence with a caller-supplied context.
func (c *Client) GetSubmittedArbitrationEvidenceWithContext(ctx context.Context, disputeId string) (*DisputeCompiledSubmittedEvidenceResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response DisputeCompiledSubmittedEvidenceResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId, evidence, arbitration, submitted), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetDisputeSchemeFiles sends GET /disputes/{dispute_id}/schemefiles: get dispute scheme files.
//
// disputeId is the {dispute_id} path value: the dispute identifier.
//
// It returns the SchemeFilesResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." disputeId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetDisputeSchemeFiles(disputeId string) (*SchemeFilesResponse, error) {
	return c.GetDisputeSchemeFilesWithContext(context.Background(), disputeId)
}

// GetDisputeSchemeFilesWithContext is GetDisputeSchemeFiles with a caller-supplied context.
func (c *Client) GetDisputeSchemeFilesWithContext(ctx context.Context, disputeId string) (*SchemeFilesResponse, error) {
	disputeId, err := common.EncodePathSegment(disputeId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response SchemeFilesResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(disputes, disputeId, schemeFiles), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
