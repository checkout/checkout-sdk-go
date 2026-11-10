package iddocumentverification

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/identities"
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

func (c *Client) CreateIdDocumentVerification(request CreateIdDocumentVerificationRequest) (*IdDocumentVerificationResponse, error) {
	return c.CreateIdDocumentVerificationWithContext(context.Background(), request)
}

func (c *Client) CreateIdDocumentVerificationWithContext(ctx context.Context, request CreateIdDocumentVerificationRequest) (*IdDocumentVerificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(idDocumentVerificationsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIdDocumentVerification sends GET /id-document-verifications/{id_document_verification_id}: get
// an ID document verification.
//
// verificationId is the {id_document_verification_id} path value: the ID document verification's
// unique identifier.
//
// It returns the IdDocumentVerificationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdDocumentVerification(verificationId string) (*IdDocumentVerificationResponse, error) {
	return c.GetIdDocumentVerificationWithContext(context.Background(), verificationId)
}

// GetIdDocumentVerificationWithContext is GetIdDocumentVerification with a caller-supplied context.
func (c *Client) GetIdDocumentVerificationWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// AnonymizeIdDocumentVerification sends POST
// /id-document-verifications/{id_document_verification_id}/anonymize: anonymize an ID document
// verification.
//
// verificationId is the {id_document_verification_id} path value: the ID document verification's
// unique identifier.
//
// It returns the IdDocumentVerificationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) AnonymizeIdDocumentVerification(verificationId string) (*IdDocumentVerificationResponse, error) {
	return c.AnonymizeIdDocumentVerificationWithContext(context.Background(), verificationId)
}

// AnonymizeIdDocumentVerificationWithContext is AnonymizeIdDocumentVerification with a caller-supplied context.
func (c *Client) AnonymizeIdDocumentVerificationWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId, anonymizePath), auth, nil, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateIdDocumentVerificationAttempt sends POST
// /id-document-verifications/{id_document_verification_id}/attempts: create an ID document
// verification attempt.
//
// verificationId is the {id_document_verification_id} path value: the ID document verification's
// unique identifier. request is sent as the JSON body.
//
// It returns the IdDocumentVerificationAttemptResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CreateIdDocumentVerificationAttempt(verificationId string, request CreateIdDocumentVerificationAttemptRequest) (*IdDocumentVerificationAttemptResponse, error) {
	return c.CreateIdDocumentVerificationAttemptWithContext(context.Background(), verificationId, request)
}

// CreateIdDocumentVerificationAttemptWithContext is CreateIdDocumentVerificationAttempt with a caller-supplied context.
func (c *Client) CreateIdDocumentVerificationAttemptWithContext(ctx context.Context, verificationId string, request CreateIdDocumentVerificationAttemptRequest) (*IdDocumentVerificationAttemptResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationAttemptResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId, attemptsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIdDocumentVerificationAttempts gets the details of all attempts for a specific ID document verification.
//
// Returns the first page using the API's own defaults. To page through the results, use
// GetIdDocumentVerificationAttemptsQuery. Beta.
func (c *Client) GetIdDocumentVerificationAttempts(verificationId string) (*IdDocumentVerificationAttemptsResponse, error) {
	return c.GetIdDocumentVerificationAttemptsWithContext(context.Background(), verificationId)
}

// GetIdDocumentVerificationAttemptsWithContext is the context-aware variant of GetIdDocumentVerificationAttempts.
func (c *Client) GetIdDocumentVerificationAttemptsWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationAttemptsResponse, error) {
	return c.GetIdDocumentVerificationAttemptsQueryWithContext(ctx, verificationId, identities.AttemptsQueryFilter{})
}

// GetIdDocumentVerificationAttemptsQuery gets the details of all attempts for a specific ID document verification,
// paginated.
//
// Pass Skip and Limit on the query filter to page through the results. An empty filter behaves
// exactly like GetIdDocumentVerificationAttempts, because BuildQueryPath appends no query string when
// every value is empty. Beta.
//
// This is a separate method rather than an extra parameter on GetIdDocumentVerificationAttempts so that
// existing callers keep compiling: Go has no overloads and no optional parameters. The
// Query suffix follows the events client, where RetrieveEvents and RetrieveEventsQuery
// are paired the same way.
//
// verificationId is the {id_document_verification_id} path value: the ID document verification's
// unique identifier.
//
// It returns the IdDocumentVerificationAttemptsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdDocumentVerificationAttemptsQuery(verificationId string, query identities.AttemptsQueryFilter) (*IdDocumentVerificationAttemptsResponse, error) {
	return c.GetIdDocumentVerificationAttemptsQueryWithContext(context.Background(), verificationId, query)
}

// GetIdDocumentVerificationAttemptsQueryWithContext is the context-aware variant of
// GetIdDocumentVerificationAttemptsQuery.
func (c *Client) GetIdDocumentVerificationAttemptsQueryWithContext(ctx context.Context, verificationId string, query identities.AttemptsQueryFilter) (*IdDocumentVerificationAttemptsResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(idDocumentVerificationsPath, verificationId, attemptsPath), query)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationAttemptsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIdDocumentVerificationAttempt sends GET
// /id-document-verifications/{id_document_verification_id}/attempts/{attempt_id}: get an ID
// document verification attempt.
//
// verificationId is the {id_document_verification_id} path value: the ID document verification's
// unique identifier. attemptId is the {attempt_id} path value: the attempt's unique identifier.
//
// It returns the IdDocumentVerificationAttemptResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// or attemptId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdDocumentVerificationAttempt(verificationId, attemptId string) (*IdDocumentVerificationAttemptResponse, error) {
	return c.GetIdDocumentVerificationAttemptWithContext(context.Background(), verificationId, attemptId)
}

// GetIdDocumentVerificationAttemptWithContext is GetIdDocumentVerificationAttempt with a caller-supplied context.
func (c *Client) GetIdDocumentVerificationAttemptWithContext(ctx context.Context, verificationId, attemptId string) (*IdDocumentVerificationAttemptResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	attemptId, err = common.EncodePathSegment(attemptId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationAttemptResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId, attemptsPath, attemptId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIdDocumentVerificationReport sends GET
// /id-document-verifications/{id_document_verification_id}/pdf-report: get ID document verification
// report.
//
// verificationId is the {id_document_verification_id} path value: the ID document verification's
// unique identifier.
//
// It returns the IdDocumentVerificationReportResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdDocumentVerificationReport(verificationId string) (*IdDocumentVerificationReportResponse, error) {
	return c.GetIdDocumentVerificationReportWithContext(context.Background(), verificationId)
}

// GetIdDocumentVerificationReportWithContext is GetIdDocumentVerificationReport with a caller-supplied context.
func (c *Client) GetIdDocumentVerificationReportWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationReportResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationReportResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId, reportPath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIdDocumentVerificationAttemptAssets gets the assets (the front and back images of the
// document) uploaded for an ID document verification attempt.
//
// Results are paginated: pass Skip and Limit on the query filter to page through them. Beta.
//
// verificationId is the {id_document_verification_id} path value: the ID document verification's
// unique identifier. attemptId is the {attempt_id} path value: the attempt's unique identifier.
//
// It returns the IdDocumentVerificationAttemptAssetsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// or attemptId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdDocumentVerificationAttemptAssets(verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdDocumentVerificationAttemptAssetsResponse, error) {
	return c.GetIdDocumentVerificationAttemptAssetsWithContext(context.Background(), verificationId, attemptId, query)
}

// GetIdDocumentVerificationAttemptAssetsWithContext is the context-aware variant of
// GetIdDocumentVerificationAttemptAssets.
func (c *Client) GetIdDocumentVerificationAttemptAssetsWithContext(ctx context.Context, verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdDocumentVerificationAttemptAssetsResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	attemptId, err = common.EncodePathSegment(attemptId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(idDocumentVerificationsPath, verificationId, attemptsPath, attemptId, assetsPath), query)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationAttemptAssetsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
