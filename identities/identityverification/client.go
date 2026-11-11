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

// GetIdentityVerification sends GET /identity-verifications/{identity_verification_id}: get an
// identity verification.
//
// verificationId is the {identity_verification_id} path value: the identity verification's unique
// identifier.
//
// It returns the IdentityVerificationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdentityVerification(verificationId string) (*IdentityVerificationResponse, error) {
	return c.GetIdentityVerificationWithContext(context.Background(), verificationId)
}

// GetIdentityVerificationWithContext is GetIdentityVerification with a caller-supplied context.
func (c *Client) GetIdentityVerificationWithContext(ctx context.Context, verificationId string) (*IdentityVerificationResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

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

// AnonymizeIdentityVerification sends POST
// /identity-verifications/{identity_verification_id}/anonymize: anonymize an identity verification.
//
// verificationId is the {identity_verification_id} path value: the identity verification's unique
// identifier.
//
// It returns the IdentityVerificationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) AnonymizeIdentityVerification(verificationId string) (*IdentityVerificationResponse, error) {
	return c.AnonymizeIdentityVerificationWithContext(context.Background(), verificationId)
}

// AnonymizeIdentityVerificationWithContext is AnonymizeIdentityVerification with a caller-supplied context.
func (c *Client) AnonymizeIdentityVerificationWithContext(ctx context.Context, verificationId string) (*IdentityVerificationResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

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

// CreateIdentityVerificationAttempt sends POST
// /identity-verifications/{identity_verification_id}/attempts: create an identity verification
// attempt.
//
// verificationId is the {identity_verification_id} path value: the identity verification's unique
// identifier. request is sent as the JSON body.
//
// It returns the IdentityVerificationAttemptResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CreateIdentityVerificationAttempt(verificationId string, request CreateIdentityVerificationAttemptRequest) (*IdentityVerificationAttemptResponse, error) {
	return c.CreateIdentityVerificationAttemptWithContext(context.Background(), verificationId, request)
}

// CreateIdentityVerificationAttemptWithContext is CreateIdentityVerificationAttempt with a caller-supplied context.
func (c *Client) CreateIdentityVerificationAttemptWithContext(ctx context.Context, verificationId string, request CreateIdentityVerificationAttemptRequest) (*IdentityVerificationAttemptResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

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
//
// verificationId is the {identity_verification_id} path value: the identity verification's unique
// identifier.
//
// It returns the IdentityVerificationAttemptsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdentityVerificationAttemptsQuery(verificationId string, query identities.AttemptsQueryFilter) (*IdentityVerificationAttemptsResponse, error) {
	return c.GetIdentityVerificationAttemptsQueryWithContext(context.Background(), verificationId, query)
}

// GetIdentityVerificationAttemptsQueryWithContext is the context-aware variant of
// GetIdentityVerificationAttemptsQuery.
func (c *Client) GetIdentityVerificationAttemptsQueryWithContext(ctx context.Context, verificationId string, query identities.AttemptsQueryFilter) (*IdentityVerificationAttemptsResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

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

// GetIdentityVerificationAttempt sends GET
// /identity-verifications/{identity_verification_id}/attempts/{attempt_id}: get an identity
// verification attempt.
//
// verificationId is the {identity_verification_id} path value: the identity verification's unique
// identifier. attemptId is the {attempt_id} path value: the attempt's unique identifier.
//
// It returns the IdentityVerificationAttemptResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// or attemptId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdentityVerificationAttempt(verificationId, attemptId string) (*IdentityVerificationAttemptResponse, error) {
	return c.GetIdentityVerificationAttemptWithContext(context.Background(), verificationId, attemptId)
}

// GetIdentityVerificationAttemptWithContext is GetIdentityVerificationAttempt with a caller-supplied context.
func (c *Client) GetIdentityVerificationAttemptWithContext(ctx context.Context, verificationId, attemptId string) (*IdentityVerificationAttemptResponse, error) {
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

	var response IdentityVerificationAttemptResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, attemptsPath, attemptId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetIdentityVerificationReport sends GET
// /identity-verifications/{identity_verification_id}/pdf-report: get identity verification report.
//
// verificationId is the {identity_verification_id} path value: the identity verification's unique
// identifier.
//
// It returns the IdentityVerificationReportResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdentityVerificationReport(verificationId string) (*IdentityVerificationReportResponse, error) {
	return c.GetIdentityVerificationReportWithContext(context.Background(), verificationId)
}

// GetIdentityVerificationReportWithContext is GetIdentityVerificationReport with a caller-supplied context.
func (c *Client) GetIdentityVerificationReportWithContext(ctx context.Context, verificationId string) (*IdentityVerificationReportResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

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

// GetIdentityVerificationAttemptAssets sends GET
// /identity-verifications/{identity_verification_id}/attempts/{attempt_id}/assets: get identity
// verification attempt assets.
//
// verificationId is the {identity_verification_id} path value: the identity verification's unique
// identifier. attemptId is the {attempt_id} path value: the attempt's unique identifier. query is
// sent as the query string.
//
// It returns the IdentityVerificationAttemptAssetsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// or attemptId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetIdentityVerificationAttemptAssets(verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdentityVerificationAttemptAssetsResponse, error) {
	return c.GetIdentityVerificationAttemptAssetsWithContext(context.Background(), verificationId, attemptId, query)
}

// GetIdentityVerificationAttemptAssetsWithContext is GetIdentityVerificationAttemptAssets with a caller-supplied context.
func (c *Client) GetIdentityVerificationAttemptAssetsWithContext(ctx context.Context, verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*IdentityVerificationAttemptAssetsResponse, error) {
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
