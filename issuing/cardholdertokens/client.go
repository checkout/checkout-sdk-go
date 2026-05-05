package cardholdertokens

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

type Client struct {
	configuration *configuration.Configuration
	httpClient    *http.Client
}

func NewClient(cfg *configuration.Configuration) *Client {
	return &Client{
		configuration: cfg,
		httpClient:    &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) RequestCardholderToken(request CardholderTokenRequest) (*CardholderTokenResponse, error) {
	return c.RequestCardholderTokenWithContext(context.Background(), request)
}

func (c *Client) RequestCardholderTokenWithContext(ctx context.Context, request CardholderTokenRequest) (*CardholderTokenResponse, error) {
	var baseUri string
	if c.configuration.EnvironmentSubdomain != nil {
		baseUri = c.configuration.EnvironmentSubdomain.ApiUrl
	}
	if baseUri == "" {
		baseUri = c.configuration.Environment.BaseUri()
	}

	data := url.Values{}
	data.Set("grant_type", request.GrantType)
	data.Set("client_id", request.ClientId)
	data.Set("client_secret", request.ClientSecret)
	data.Set("cardholder_id", request.CardholderId)
	if request.SingleUse {
		data.Set("single_use", "true")
	}

	reqURL := fmt.Sprintf("%s/%s", strings.TrimRight(baseUri, "/"), cardholderTokenPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResp CardholderTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}
