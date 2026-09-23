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

func (c *Client) GetIdDocumentVerification(verificationId string) (*IdDocumentVerificationResponse, error) {
	return c.GetIdDocumentVerificationWithContext(context.Background(), verificationId)
}

func (c *Client) GetIdDocumentVerificationWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationResponse, error) {
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

func (c *Client) AnonymizeIdDocumentVerification(verificationId string) (*IdDocumentVerificationResponse, error) {
	return c.AnonymizeIdDocumentVerificationWithContext(context.Background(), verificationId)
}

func (c *Client) AnonymizeIdDocumentVerificationWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationResponse, error) {
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

func (c *Client) CreateIdDocumentVerificationAttempt(verificationId string, request CreateIdDocumentVerificationAttemptRequest) (*IdDocumentVerificationAttemptResponse, error) {
	return c.CreateIdDocumentVerificationAttemptWithContext(context.Background(), verificationId, request)
}

func (c *Client) CreateIdDocumentVerificationAttemptWithContext(ctx context.Context, verificationId string, request CreateIdDocumentVerificationAttemptRequest) (*IdDocumentVerificationAttemptResponse, error) {
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
func (c *Client) GetIdDocumentVerificationAttemptsQuery(verificationId string, query identities.AttemptsQueryFilter) (*IdDocumentVerificationAttemptsResponse, error) {
	return c.GetIdDocumentVerificationAttemptsQueryWithContext(context.Background(), verificationId, query)
}

// GetIdDocumentVerificationAttemptsQueryWithContext is the context-aware variant of
// GetIdDocumentVerificationAttemptsQuery.
func (c *Client) GetIdDocumentVerificationAttemptsQueryWithContext(ctx context.Context, verificationId string, query identities.AttemptsQueryFilter) (*IdDocumentVerificationAttemptsResponse, error) {
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

func (c *Client) GetIdDocumentVerificationAttempt(verificationId, attemptId string) (*IdDocumentVerificationAttemptResponse, error) {
	return c.GetIdDocumentVerificationAttemptWithContext(context.Background(), verificationId, attemptId)
}

func (c *Client) GetIdDocumentVerificationAttemptWithContext(ctx context.Context, verificationId, attemptId string) (*IdDocumentVerificationAttemptResponse, error) {
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

func (c *Client) GetIdDocumentVerificationReport(verificationId string) (*IdDocumentVerificationReportResponse, error) {
	return c.GetIdDocumentVerificationReportWithContext(context.Background(), verificationId)
}

func (c *Client) GetIdDocumentVerificationReportWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationReportResponse, error) {
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
func (c *Client) GetIdDocumentVerificationAttemptAssets(verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdDocumentVerificationAttemptAssetsResponse, error) {
	return c.GetIdDocumentVerificationAttemptAssetsWithContext(context.Background(), verificationId, attemptId, query)
}

// GetIdDocumentVerificationAttemptAssetsWithContext is the context-aware variant of
// GetIdDocumentVerificationAttemptAssets.
func (c *Client) GetIdDocumentVerificationAttemptAssetsWithContext(ctx context.Context, verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdDocumentVerificationAttemptAssetsResponse, error) {
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
