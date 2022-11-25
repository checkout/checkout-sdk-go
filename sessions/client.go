package sessions

import (
	"github.com/checkout/checkout-sdk-go-beta/client"
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
	"github.com/checkout/checkout-sdk-go-beta/sessions/channels"
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

func (c *Client) RequestSession(request SessionRequest) (*SessionResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response SessionDetails
	err = c.apiClient.Post(
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

func (c *Client) GetSessionDetails(sessionId string, sessionSecret string) (*GetSessionResponse, error) {
	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response GetSessionResponse
	err = c.apiClient.Get(common.BuildPath(SessionsPath, sessionId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateSession(
	sessionId string,
	request channels.Channel,
	sessionSecret string,
) (*GetSessionResponse, error) {
	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response GetSessionResponse
	err = c.apiClient.Put(
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

func (c *Client) CompleteSession(sessionId, sessionSecret string) (*common.MetadataResponse, error) {
	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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

func (c *Client) Update3dsMethodCompletion(
	sessionId string,
	request ThreeDsMethodCompletionRequest,
	sessionSecret string,
) (*Update3dsMethodCompletionResponse, error) {
	auth, err := c.customSdkAuthorization(sessionSecret)
	if err != nil {
		return nil, err
	}

	var response Update3dsMethodCompletionResponse
	err = c.apiClient.Put(
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
