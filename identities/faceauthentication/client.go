package faceauthentication

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/identities"
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

func (c *Client) CreateFaceAuthentication(request CreateFaceAuthenticationRequest) (*FaceAuthenticationResponse, error) {
	return c.CreateFaceAuthenticationWithContext(context.Background(), request)
}

func (c *Client) CreateFaceAuthenticationWithContext(ctx context.Context, request CreateFaceAuthenticationRequest) (*FaceAuthenticationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FaceAuthenticationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(faceAuthenticationsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFaceAuthentication(faceAuthenticationId string) (*FaceAuthenticationResponse, error) {
	return c.GetFaceAuthenticationWithContext(context.Background(), faceAuthenticationId)
}

func (c *Client) GetFaceAuthenticationWithContext(ctx context.Context, faceAuthenticationId string) (*FaceAuthenticationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FaceAuthenticationResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(faceAuthenticationsPath, faceAuthenticationId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AnonymizeFaceAuthentication(faceAuthenticationId string) (*FaceAuthenticationResponse, error) {
	return c.AnonymizeFaceAuthenticationWithContext(context.Background(), faceAuthenticationId)
}

func (c *Client) AnonymizeFaceAuthenticationWithContext(ctx context.Context, faceAuthenticationId string) (*FaceAuthenticationResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FaceAuthenticationResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(faceAuthenticationsPath, faceAuthenticationId, anonymizePath), auth, nil, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateFaceAuthenticationAttempt(faceAuthenticationId string, request CreateFaceAuthenticationAttemptRequest) (*FaceAuthenticationAttemptResponse, error) {
	return c.CreateFaceAuthenticationAttemptWithContext(context.Background(), faceAuthenticationId, request)
}

func (c *Client) CreateFaceAuthenticationAttemptWithContext(ctx context.Context, faceAuthenticationId string, request CreateFaceAuthenticationAttemptRequest) (*FaceAuthenticationAttemptResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FaceAuthenticationAttemptResponse
	err = c.apiClient.PostWithContext(ctx, common.BuildPath(faceAuthenticationsPath, faceAuthenticationId, attemptsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFaceAuthenticationAttempts(faceAuthenticationId string) (*FaceAuthenticationAttemptsResponse, error) {
	return c.GetFaceAuthenticationAttemptsWithContext(context.Background(), faceAuthenticationId)
}

func (c *Client) GetFaceAuthenticationAttemptsWithContext(ctx context.Context, faceAuthenticationId string) (*FaceAuthenticationAttemptsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FaceAuthenticationAttemptsResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(faceAuthenticationsPath, faceAuthenticationId, attemptsPath), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFaceAuthenticationAttempt(faceAuthenticationId, attemptId string) (*FaceAuthenticationAttemptResponse, error) {
	return c.GetFaceAuthenticationAttemptWithContext(context.Background(), faceAuthenticationId, attemptId)
}

func (c *Client) GetFaceAuthenticationAttemptWithContext(ctx context.Context, faceAuthenticationId, attemptId string) (*FaceAuthenticationAttemptResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response FaceAuthenticationAttemptResponse
	err = c.apiClient.GetWithContext(ctx, common.BuildPath(faceAuthenticationsPath, faceAuthenticationId, attemptsPath, attemptId), auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetFaceAuthenticationAttemptAssets(faceAuthenticationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*FaceAuthenticationAttemptAssetsResponse, error) {
	return c.GetFaceAuthenticationAttemptAssetsWithContext(context.Background(), faceAuthenticationId, attemptId, query)
}

func (c *Client) GetFaceAuthenticationAttemptAssetsWithContext(ctx context.Context, faceAuthenticationId, attemptId string, query identities.AttemptAssetsQueryFilter) (*FaceAuthenticationAttemptAssetsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	url, err := common.BuildQueryPath(common.BuildPath(faceAuthenticationsPath, faceAuthenticationId, attemptsPath, attemptId, assetsPath), query)
	if err != nil {
		return nil, err
	}

	var response FaceAuthenticationAttemptAssetsResponse
	err = c.apiClient.GetWithContext(ctx, url, auth, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
