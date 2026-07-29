package addressdocumentverification

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
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

func (c *Client) GetAddressDocumentVerification(verificationId string) (*AddressDocumentVerificationResponse, error) {
	return c.GetAddressDocumentVerificationWithContext(context.Background(), verificationId)
}

func (c *Client) GetAddressDocumentVerificationWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationResponse, error) {
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

func (c *Client) AnonymizeAddressDocumentVerification(verificationId string) (*AddressDocumentVerificationResponse, error) {
	return c.AnonymizeAddressDocumentVerificationWithContext(context.Background(), verificationId)
}

func (c *Client) AnonymizeAddressDocumentVerificationWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationResponse, error) {
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

func (c *Client) CreateAddressDocumentVerificationAttempt(verificationId string, request CreateAddressDocumentVerificationAttemptRequest) (*AddressDocumentVerificationAttemptResponse, error) {
	return c.CreateAddressDocumentVerificationAttemptWithContext(context.Background(), verificationId, request)
}

func (c *Client) CreateAddressDocumentVerificationAttemptWithContext(ctx context.Context, verificationId string, request CreateAddressDocumentVerificationAttemptRequest) (*AddressDocumentVerificationAttemptResponse, error) {
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

func (c *Client) GetAddressDocumentVerificationAttempts(verificationId string) (*AddressDocumentVerificationAttemptsResponse, error) {
	return c.GetAddressDocumentVerificationAttemptsWithContext(context.Background(), verificationId)
}

func (c *Client) GetAddressDocumentVerificationAttemptsWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationAttemptsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response AddressDocumentVerificationAttemptsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(addressDocumentVerificationsPath, verificationId, attemptsPath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetAddressDocumentVerificationAttempt(verificationId, attemptId string) (*AddressDocumentVerificationAttemptResponse, error) {
	return c.GetAddressDocumentVerificationAttemptWithContext(context.Background(), verificationId, attemptId)
}

func (c *Client) GetAddressDocumentVerificationAttemptWithContext(ctx context.Context, verificationId, attemptId string) (*AddressDocumentVerificationAttemptResponse, error) {
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

func (c *Client) GetAddressDocumentVerificationReport(verificationId string) (*AddressDocumentVerificationReportResponse, error) {
	return c.GetAddressDocumentVerificationReportWithContext(context.Background(), verificationId)
}

func (c *Client) GetAddressDocumentVerificationReportWithContext(ctx context.Context, verificationId string) (*AddressDocumentVerificationReportResponse, error) {
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
