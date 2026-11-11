package addressdocumentverification

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

func (c *Client) CreateAddressDocumentVerification(request CreateAddressDocumentVerificationRequest) (*AddressDocumentVerificationResponse, error) {
	return c.CreateAddressDocumentVerificationWithContext(context.Background(), request)
}

func (c *Client) CreateAddressDocumentVerificationWithContext(ctx context.Context, request CreateAddressDocumentVerificationRequest) (*AddressDocumentVerificationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(addressDocumentVerificationsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAddressDocumentVerification sends GET
// /address-document-verifications/{address_document_verification_id}: get an address document
// verification.
//
// verificationId is the {address_document_verification_id} path value: the address document
// verification's unique identifier.
//
// It returns the AddressDocumentVerificationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetAddressDocumentVerification(verificationId string) (*AddressDocumentVerificationResponse, error) {
	return c.GetAddressDocumentVerificationWithContext(context.Background(), verificationId)
}

// GetAddressDocumentVerificationWithContext is GetAddressDocumentVerification with a caller-supplied context.
func (c *Client) GetAddressDocumentVerificationWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(addressDocumentVerificationsPath, verificationId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// AnonymizeAddressDocumentVerification sends POST
// /address-document-verifications/{address_document_verification_id}/anonymize: anonymize an
// address document verification.
//
// verificationId is the {address_document_verification_id} path value: the address document
// verification's unique identifier.
//
// It returns the AddressDocumentVerificationResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) AnonymizeAddressDocumentVerification(verificationId string) (*AddressDocumentVerificationResponse, error) {
	return c.AnonymizeAddressDocumentVerificationWithContext(context.Background(), verificationId)
}

// AnonymizeAddressDocumentVerificationWithContext is AnonymizeAddressDocumentVerification with a caller-supplied context.
func (c *Client) AnonymizeAddressDocumentVerificationWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(addressDocumentVerificationsPath, verificationId, anonymizePath), auth, nil, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CreateAddressDocumentVerificationAttempt sends POST
// /address-document-verifications/{address_document_verification_id}/attempts: create an address
// document verification attempt.
//
// verificationId is the {address_document_verification_id} path value: the address document
// verification's unique identifier. request is sent as the JSON body.
//
// It returns the AddressDocumentVerificationAttemptResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CreateAddressDocumentVerificationAttempt(verificationId string, request CreateAddressDocumentVerificationAttemptRequest) (*AddressDocumentVerificationAttemptResponse, error) {
	return c.CreateAddressDocumentVerificationAttemptWithContext(context.Background(), verificationId, request)
}

// CreateAddressDocumentVerificationAttemptWithContext is CreateAddressDocumentVerificationAttempt with a caller-supplied context.
func (c *Client) CreateAddressDocumentVerificationAttemptWithContext(ctx context.Context, verificationId string, request CreateAddressDocumentVerificationAttemptRequest) (*AddressDocumentVerificationAttemptResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationAttemptResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(addressDocumentVerificationsPath, verificationId, attemptsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAddressDocumentVerificationAttempts gets the details of all attempts for a specific address document verification.
//
// Returns the first page using the API's own defaults. To page through the results, use
// GetAddressDocumentVerificationAttemptsQuery. Beta.
func (c *Client) GetAddressDocumentVerificationAttempts(verificationId string) (*AddressDocumentVerificationAttemptsResponse, error) {
	return c.GetAddressDocumentVerificationAttemptsWithContext(context.Background(), verificationId)
}

// GetAddressDocumentVerificationAttemptsWithContext is the context-aware variant of GetAddressDocumentVerificationAttempts.
func (c *Client) GetAddressDocumentVerificationAttemptsWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationAttemptsResponse, error) {
	return c.GetAddressDocumentVerificationAttemptsQueryWithContext(ctx, verificationId, identities.AttemptsQueryFilter{})
}

// GetAddressDocumentVerificationAttemptsQuery gets the details of all attempts for a specific address document verification,
// paginated.
//
// Pass Skip and Limit on the query filter to page through the results. An empty filter behaves
// exactly like GetAddressDocumentVerificationAttempts, because BuildQueryPath appends no query string when
// every value is empty. Beta.
//
// This is a separate method rather than an extra parameter on GetAddressDocumentVerificationAttempts so that
// existing callers keep compiling: Go has no overloads and no optional parameters. The
// Query suffix follows the events client, where RetrieveEvents and RetrieveEventsQuery
// are paired the same way.
//
// verificationId is the {address_document_verification_id} path value: the address document
// verification's unique identifier.
//
// It returns the AddressDocumentVerificationAttemptsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetAddressDocumentVerificationAttemptsQuery(verificationId string, query identities.AttemptsQueryFilter) (*AddressDocumentVerificationAttemptsResponse, error) {
	return c.GetAddressDocumentVerificationAttemptsQueryWithContext(context.Background(), verificationId, query)
}

// GetAddressDocumentVerificationAttemptsQueryWithContext is the context-aware variant of
// GetAddressDocumentVerificationAttemptsQuery.
func (c *Client) GetAddressDocumentVerificationAttemptsQueryWithContext(ctx context.Context, verificationId string, query identities.AttemptsQueryFilter) (*AddressDocumentVerificationAttemptsResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(addressDocumentVerificationsPath, verificationId, attemptsPath), query)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationAttemptsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAddressDocumentVerificationAttempt sends GET
// /address-document-verifications/{address_document_verification_id}/attempts/{attempt_id}: get an
// address document verification attempt.
//
// verificationId is the {address_document_verification_id} path value: the address document
// verification's unique identifier. attemptId is the {attempt_id} path value: the attempt's unique
// identifier.
//
// It returns the AddressDocumentVerificationAttemptResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// or attemptId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetAddressDocumentVerificationAttempt(verificationId, attemptId string) (*AddressDocumentVerificationAttemptResponse, error) {
	return c.GetAddressDocumentVerificationAttemptWithContext(context.Background(), verificationId, attemptId)
}

// GetAddressDocumentVerificationAttemptWithContext is GetAddressDocumentVerificationAttempt with a caller-supplied context.
func (c *Client) GetAddressDocumentVerificationAttemptWithContext(ctx context.Context, verificationId, attemptId string) (*AddressDocumentVerificationAttemptResponse, error) {
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

	var response AddressDocumentVerificationAttemptResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(addressDocumentVerificationsPath, verificationId, attemptsPath, attemptId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAddressDocumentVerificationReport sends GET
// /address-document-verifications/{address_document_verification_id}/pdf-report: get address
// document verification report.
//
// verificationId is the {address_document_verification_id} path value: the address document
// verification's unique identifier.
//
// It returns the AddressDocumentVerificationReportResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetAddressDocumentVerificationReport(verificationId string) (*AddressDocumentVerificationReportResponse, error) {
	return c.GetAddressDocumentVerificationReportWithContext(context.Background(), verificationId)
}

// GetAddressDocumentVerificationReportWithContext is GetAddressDocumentVerificationReport with a caller-supplied context.
func (c *Client) GetAddressDocumentVerificationReportWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationReportResponse, error) {
	verificationId, err := common.EncodePathSegment(verificationId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationReportResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(addressDocumentVerificationsPath, verificationId, reportPath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAddressDocumentVerificationAttemptAssets gets the assets (the document image) uploaded for an
// address document verification attempt.
//
// Results are paginated: pass Skip and Limit on the query filter to page through them. Beta.
//
// verificationId is the {address_document_verification_id} path value: the address document
// verification's unique identifier. attemptId is the {attempt_id} path value: the attempt's unique
// identifier.
//
// It returns the AddressDocumentVerificationAttemptAssetsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." verificationId
// or attemptId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetAddressDocumentVerificationAttemptAssets(verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*AddressDocumentVerificationAttemptAssetsResponse, error) {
	return c.GetAddressDocumentVerificationAttemptAssetsWithContext(context.Background(), verificationId, attemptId, query)
}

// GetAddressDocumentVerificationAttemptAssetsWithContext is the context-aware variant of
// GetAddressDocumentVerificationAttemptAssets.
func (c *Client) GetAddressDocumentVerificationAttemptAssetsWithContext(ctx context.Context, verificationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*AddressDocumentVerificationAttemptAssetsResponse, error) {
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

	url, err := common.BuildQueryPath(common.BuildPath(addressDocumentVerificationsPath, verificationId, attemptsPath, attemptId, assetsPath), query)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationAttemptAssetsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
