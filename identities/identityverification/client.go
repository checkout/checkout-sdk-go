package identityverification

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

func (c *Client) CreateIdentityVerificationAndAttempt(request CreateIdentityVerificationAndAttemptRequest) (*IdentityVerificationAndAttemptResponse, error) {
	return c.CreateIdentityVerificationAndAttemptWithContext(context.Background(), request)
}

func (c *Client) CreateIdentityVerificationAndAttemptWithContext(ctx context.Context, request CreateIdentityVerificationAndAttemptRequest) (*IdentityVerificationAndAttemptResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationAndAttemptResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(createAndOpenPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateIdentityVerification(request CreateIdentityVerificationRequest) (*IdentityVerificationResponse, error) {
	return c.CreateIdentityVerificationWithContext(context.Background(), request)
}

func (c *Client) CreateIdentityVerificationWithContext(ctx context.Context, request CreateIdentityVerificationRequest) (*IdentityVerificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(identityVerificationsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetIdentityVerification(verificationId string) (*IdentityVerificationResponse, error) {
	return c.GetIdentityVerificationWithContext(context.Background(), verificationId)
}

func (c *Client) GetIdentityVerificationWithContext(ctx context.Context, verificationId string) (*IdentityVerificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AnonymizeIdentityVerification(verificationId string) (*IdentityVerificationResponse, error) {
	return c.AnonymizeIdentityVerificationWithContext(context.Background(), verificationId)
}

func (c *Client) AnonymizeIdentityVerificationWithContext(ctx context.Context, verificationId string) (*IdentityVerificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, anonymizePath), auth, nil, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateIdentityVerificationAttempt(verificationId string, request CreateIdentityVerificationAttemptRequest) (*IdentityVerificationAttemptResponse, error) {
	return c.CreateIdentityVerificationAttemptWithContext(context.Background(), verificationId, request)
}

func (c *Client) CreateIdentityVerificationAttemptWithContext(ctx context.Context, verificationId string, request CreateIdentityVerificationAttemptRequest) (*IdentityVerificationAttemptResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationAttemptResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, attemptsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIdentityVerificationAttempts gets the details of all attempts for a specific identity verification.
//
// Returns the first page using the API's own defaults. To page through the results, use
// GetIdentityVerificationAttemptsQuery. Beta.
func (c *Client) GetIdentityVerificationAttempts(verificationId string) (*IdentityVerificationAttemptsResponse, error) {
	return c.GetIdentityVerificationAttemptsWithContext(context.Background(), verificationId)
}

// GetIdentityVerificationAttemptsWithContext is the context-aware variant of GetIdentityVerificationAttempts.
func (c *Client) GetIdentityVerificationAttemptsWithContext(ctx context.Context, verificationId string) (*IdentityVerificationAttemptsResponse, error) {
	return c.GetIdentityVerificationAttemptsQueryWithContext(ctx, verificationId, identities.AttemptsQueryFilter{})
}

// GetIdentityVerificationAttemptsQuery gets the details of all attempts for a specific identity verification,
// paginated.
//
// Pass Skip and Limit on the query filter to page through the results. An empty filter behaves
// exactly like GetIdentityVerificationAttempts, because BuildQueryPath appends no query string when
// every value is empty. Beta.
//
// This is a separate method rather than an extra parameter on GetIdentityVerificationAttempts so that
// existing callers keep compiling: Go has no overloads and no optional parameters. The
// Query suffix follows the events client, where RetrieveEvents and RetrieveEventsQuery
// are paired the same way.
func (c *Client) GetIdentityVerificationAttemptsQuery(verificationId string, query identities.AttemptsQueryFilter) (*IdentityVerificationAttemptsResponse, error) {
	return c.GetIdentityVerificationAttemptsQueryWithContext(context.Background(), verificationId, query)
}

// GetIdentityVerificationAttemptsQueryWithContext is the context-aware variant of
// GetIdentityVerificationAttemptsQuery.
func (c *Client) GetIdentityVerificationAttemptsQueryWithContext(ctx context.Context, verificationId string, query identities.AttemptsQueryFilter) (*IdentityVerificationAttemptsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(identityVerificationsPath, verificationId, attemptsPath), query)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationAttemptsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetIdentityVerificationAttempt(verificationId, attemptId string) (*IdentityVerificationAttemptResponse, error) {
	return c.GetIdentityVerificationAttemptWithContext(context.Background(), verificationId, attemptId)
}

func (c *Client) GetIdentityVerificationAttemptWithContext(ctx context.Context, verificationId, attemptId string) (*IdentityVerificationAttemptResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationAttemptResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, attemptsPath, attemptId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetIdentityVerificationReport(verificationId string) (*IdentityVerificationReportResponse, error) {
	return c.GetIdentityVerificationReportWithContext(context.Background(), verificationId)
}

func (c *Client) GetIdentityVerificationReportWithContext(ctx context.Context, verificationId string) (*IdentityVerificationReportResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationReportResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, reportPath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetIdentityVerificationAttemptAssets(verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdentityVerificationAttemptAssetsResponse, error) {
	return c.GetIdentityVerificationAttemptAssetsWithContext(context.Background(), verificationId, attemptId, query)
}

func (c *Client) GetIdentityVerificationAttemptAssetsWithContext(ctx context.Context, verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdentityVerificationAttemptAssetsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(identityVerificationsPath, verificationId, attemptsPath, attemptId, assetsPath), query)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationAttemptAssetsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
