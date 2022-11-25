package configuration

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/checkout/checkout-sdk-go-beta/errors"
)

type (
	OAuthSdkCredentials struct {
		//HttpClient
		ClientId         string
		ClientSecret     string
		AuthorizationUri string
		Scopes           []string
		AccessToken      *OAuthAccessToken
	}

	OAuthAccessToken struct {
		Token          string
		ExpirationDate time.Time
	}

	OAuthServiceResponse struct {
		AccessToken string        `json:"access_token,omitempty"`
		ExpiresIn   time.Duration `json:"expires_in,omitempty"`
	}
)

func NewOAuthSdkCredentials(clientId, clientSecret, authorizationUri string, scopes []string) (*OAuthSdkCredentials, error) {
	sdkCredentials := OAuthSdkCredentials{
		ClientId:         clientId,
		ClientSecret:     clientSecret,
		AuthorizationUri: authorizationUri,
		Scopes:           scopes,
	}

	err := sdkCredentials.GetAccessToken()
	if err != nil {
		return nil, err
	}

	return &sdkCredentials, nil
}

func (f *OAuthSdkCredentials) GetAuthorization(authorizationType AuthorizationType) (*SdkAuthorization, error) {
	switch authorizationType {
	case PublicKeyOrOauth, SecretKeyOrOauth, OAuth:
		err := f.GetAccessToken()
		if err != nil {
			return nil, err
		}
		return &SdkAuthorization{
			PlatformType: DefaultOAuth,
			Credential:   f.AccessToken.Token,
		}, nil
	default:
		return nil, errors.CheckoutAuthorizationError("Invalid authorization type")
	}
}

func (f *OAuthSdkCredentials) GetAccessToken() error {
	if f.AccessToken != nil && f.AccessToken.IsValid() {
		return nil
	}

	data := url.Values{}
	data.Set("client_id", f.ClientId)
	data.Set("client_secret", f.ClientSecret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", strings.Join(f.Scopes, " "))

	req, err := http.NewRequest(http.MethodPost, f.AuthorizationUri, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{Timeout: time.Duration(5) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var oauthResp OAuthServiceResponse
	err = json.Unmarshal(body, &oauthResp)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		var oauthErr errors.CheckoutOAuthError
		err = json.Unmarshal(body, &oauthErr)
		if err != nil {
			return err
		}
		return errors.CheckoutAuthorizationError(oauthErr.Description)
	}

	accessToken := &OAuthAccessToken{
		Token:          oauthResp.AccessToken,
		ExpirationDate: time.Now().Add(oauthResp.ExpiresIn * time.Second),
	}
	f.AccessToken = accessToken
	return nil
}

func (t *OAuthAccessToken) IsValid() bool {
	if t.Token == "" {
		return false
	}
	return t.ExpirationDate.After(time.Now())
}
