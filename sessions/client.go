package sessions

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/sessions/channels"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

func NewClient(
	configuration *configuration.Configuration,
	apiClient client.HttpClient,
) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

func (c *Client) RequestSession(request SessionRequest) (*SessionResponse, error) {
	return c.RequestSessionWithContext(context.Background(), request)
}

func (c *Client) RequestSessionWithContext(
	ctx context.Context,
	request SessionRequest,
) (*SessionResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response SessionDetails
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(SessionsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var sessionResponse SessionResponse
	sessionResponse.MapResponse(&response)
	return &sessionResponse, nil
}

// GetSessionDetails sends GET /sessions/{id}: get session details.
//
// sessionId is the {id} path value: session ID. sessionSecret, when not empty, authorizes the
// request instead of the configured OAuth credentials.
//
// It returns the GetSessionResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." sessionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetSessionDetails(sessionId string, sessionSecret string) (*GetSessionResponse, error) {
	return c.GetSessionDetailsWithContext(context.Background(), sessionId, sessionSecret)
}

// GetSessionDetailsWithContext is GetSessionDetails with a caller-supplied context.
func (c *Client) GetSessionDetailsWithContext(
	ctx context.Context,
	sessionId string,
	sessionSecret string,
) (*GetSessionResponse, error) {
	sessionId, err := common.EncodePathSegment(sessionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response GetSessionResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(SessionsPath, sessionId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateSession sends PUT /sessions/{id}/collect-data: update a session.
//
// sessionId is the {id} path value: session ID. request is sent as the JSON body. sessionSecret,
// when not empty, authorizes the request instead of the configured OAuth credentials.
//
// It returns the GetSessionResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." sessionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateSession(
	sessionId string,
	request channels.Channel,
	sessionSecret string,
) (*GetSessionResponse, error) {
	return c.UpdateSessionWithContext(context.Background(), sessionId, request, sessionSecret)
}

// UpdateSessionWithContext is UpdateSession with a caller-supplied context.
func (c *Client) UpdateSessionWithContext(
	ctx context.Context,
	sessionId string,
	request channels.Channel,
	sessionSecret string,
) (*GetSessionResponse, error) {
	sessionId, err := common.EncodePathSegment(sessionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response GetSessionResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(SessionsPath, sessionId, CollectDataPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// CompleteSession sends POST /sessions/{id}/complete: complete a session.
//
// sessionId is the {id} path value: session ID. sessionSecret, when not empty, authorizes the
// request instead of the configured OAuth credentials.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." sessionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) CompleteSession(sessionId, sessionSecret string) (*common.MetadataResponse, error) {
	return c.CompleteSessionWithContext(context.Background(), sessionId, sessionSecret)
}

// CompleteSessionWithContext is CompleteSession with a caller-supplied context.
func (c *Client) CompleteSessionWithContext(
	ctx context.Context,
	sessionId string,
	sessionSecret string,
) (*common.MetadataResponse, error) {
	sessionId, err := common.EncodePathSegment(sessionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(SessionsPath, sessionId, CompletePath),
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

// Update3dsMethodCompletion sends PUT /sessions/{id}/issuer-fingerprint: update session 3DS Method
// completion indicator.
//
// sessionId is the {id} path value: session ID. request is sent as the JSON body. sessionSecret,
// when not empty, authorizes the request instead of the configured OAuth credentials.
//
// It returns the Update3dsMethodCompletionResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." sessionId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) Update3dsMethodCompletion(
	sessionId string,
	request ThreeDsMethodCompletionRequest,
	sessionSecret string,
) (*Update3dsMethodCompletionResponse, error) {
	return c.Update3dsMethodCompletionWithContext(context.Background(), sessionId, request, sessionSecret)
}

// Update3dsMethodCompletionWithContext is Update3dsMethodCompletion with a caller-supplied context.
func (c *Client) Update3dsMethodCompletionWithContext(
	ctx context.Context,
	sessionId string,
	request ThreeDsMethodCompletionRequest,
	sessionSecret string,
) (*Update3dsMethodCompletionResponse, error) {
	sessionId, err := common.EncodePathSegment(sessionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response Update3dsMethodCompletionResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(SessionsPath, sessionId, IssuerFingerprintPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) customSdkAuthorization(sessionSecret string) (*configuration.SdkAuthorization, error) {
	if sessionSecret == "" {
		return c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	}
	return NewSessionSecretCredentials(sessionSecret).GetAuthorization(configuration.CustomAuth)
}
