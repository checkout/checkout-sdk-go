package iddocumentverification

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
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId), auth, &response)
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

func (c *Client) GetIdDocumentVerificationAttempts(verificationId string) (*IdDocumentVerificationAttemptsResponse, error) {
	return c.GetIdDocumentVerificationAttemptsWithContext(context.Background(), verificationId)
}

func (c *Client) GetIdDocumentVerificationAttemptsWithContext(ctx context.Context, verificationId string) (*IdDocumentVerificationAttemptsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationAttemptsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId, attemptsPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetIdDocumentVerificationAttempt(verificationId string, attemptId string) (*IdDocumentVerificationAttemptResponse, error) {
	return c.GetIdDocumentVerificationAttemptWithContext(context.Background(), verificationId, attemptId)
}

func (c *Client) GetIdDocumentVerificationAttemptWithContext(ctx context.Context, verificationId string, attemptId string) (*IdDocumentVerificationAttemptResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response IdDocumentVerificationAttemptResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId, attemptsPath, attemptId), auth, &response)
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
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(idDocumentVerificationsPath, verificationId, reportPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
