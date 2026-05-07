package identityverification

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
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId), auth, &response)
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

func (c *Client) GetIdentityVerificationAttempts(verificationId string) (*IdentityVerificationAttemptsResponse, error) {
	return c.GetIdentityVerificationAttemptsWithContext(context.Background(), verificationId)
}

func (c *Client) GetIdentityVerificationAttemptsWithContext(ctx context.Context, verificationId string) (*IdentityVerificationAttemptsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdentityVerificationAttemptsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, attemptsPath), auth, &response)
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
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, attemptsPath, attemptId), auth, &response)
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
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(identityVerificationsPath, verificationId, reportPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
