package configuration

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/checkout/checkout-sdk-go/v3/errors"
)

type (
	// OAuthSdkCredentials authorizes requests with an OAuth access token, obtained with the client
	// credentials grant and renewed when it expires. Safe for concurrent use; create it with
	// NewOAuthSdkCredentialsWithHttpClient or NewOAuthSdkCredentials and do not copy it after use.
	OAuthSdkCredentials struct {
		// HttpClient sends the token request. The OAuth builder sets it to the client given to
		// WithHttpClient, so the token request uses the same settings as API calls. When nil, a client
		// with a 5 second timeout is used; a client without Timeout gets a 5 second limit on the token
		// request only.
		HttpClient *http.Client
		// ClientId is the OAuth client ID.
		ClientId string
		// ClientSecret is the OAuth client secret.
		ClientSecret string
		// AuthorizationUri is the token endpoint.
		AuthorizationUri string
		// Scopes are the OAuth scopes requested with the token.
		Scopes []string
		// AccessToken is the cached token. It is written under a lock; use GetAuthorization to read
		// it while other goroutines may renew it.
		AccessToken *OAuthAccessToken
		// Log is the logger used for the token request.
		Log StdLogger

		mutex   sync.Mutex
		renewal *tokenRenewal
	}

	// tokenRenewal is a token request in flight; done is closed once token and err are set.
	tokenRenewal struct {
		done  chan struct{}
		token string
		err   error
	}

	// OAuthAccessToken is a cached OAuth access token.
	OAuthAccessToken struct {
		// Token is the access token sent as the bearer credential.
		Token string
		// ExpirationDate is when the token expires.
		ExpirationDate time.Time
	}

	// OAuthServiceResponse is the access token response of the authorization server.
	OAuthServiceResponse struct {
		// The access token.
		// [Optional]
		AccessToken string `json:"access_token,omitempty"`
		// The token type.
		// [Optional]
		TokenType string `json:"token_type,omitempty"`
		// The lifetime of the access token, in seconds. Holds a number of seconds, not a Duration in
		// nanoseconds.
		// [Optional]
		ExpiresIn time.Duration `json:"expires_in,omitempty"`
	}
)

// NewOAuthSdkCredentials creates the credentials and requests the first access token, using a
// client with a 5 second timeout for the token request. Use NewOAuthSdkCredentialsWithHttpClient
// to send it through your own http.Client: the client given to NewConfiguration is not used for
// the token request. It returns an error when the first token request fails.
func NewOAuthSdkCredentials(
	clientId,
	clientSecret,
	authorizationUri string,
	scopes []string,
	logger StdLogger,
) (*OAuthSdkCredentials, error) {
	return NewOAuthSdkCredentialsWithHttpClient(clientId, clientSecret, authorizationUri, scopes, nil, logger)
}

// NewOAuthSdkCredentialsWithHttpClient creates the credentials and requests the first access token
// through httpClient, which is also used for every renewal. A nil httpClient uses a client with a
// 5 second timeout. It returns an error when the first token request fails.
func NewOAuthSdkCredentialsWithHttpClient(
	clientId,
	clientSecret,
	authorizationUri string,
	scopes []string,
	httpClient *http.Client,
	logger StdLogger,
) (*OAuthSdkCredentials, error) {
	if logger == nil {
		logger = DefaultLogger()
	}

	sdkCredentials := &OAuthSdkCredentials{
		HttpClient:       httpClient,
		ClientId:         clientId,
		ClientSecret:     clientSecret,
		AuthorizationUri: authorizationUri,
		Scopes:           scopes,
		Log:              logger,
	}

	err := sdkCredentials.GetAccessToken()
	if err != nil {
		return nil, err
	}

	return sdkCredentials, nil
}

// GetAuthorization returns the OAuth authorization for authorizationType, renewing the access token
// first when it is missing or expired. It returns an error when authorizationType does not accept
// OAuth or the token request fails.
func (f *OAuthSdkCredentials) GetAuthorization(authorizationType AuthorizationType) (*SdkAuthorization, error) {
	switch authorizationType {
	case PublicKeyOrOauth, SecretKeyOrOauth, OAuth:
		token, err := f.validAccessToken()
		if err != nil {
			return nil, err
		}
		return &SdkAuthorization{
			PlatformType: DefaultOAuth,
			Credential:   token,
		}, nil
	default:
		return nil, errors.CheckoutAuthorizationError("Invalid authorization type")
	}
}

// GetAccessToken requests a new access token when the cached one is missing or expired. Concurrent
// callers share a single token request and its result. It returns the error of that request.
func (f *OAuthSdkCredentials) GetAccessToken() error {
	_, err := f.validAccessToken()
	return err
}

// validAccessToken returns the cached token, renewing it first when needed. Concurrent callers
// share one renewal and its result, error included. The lock is not held during the token request.
func (f *OAuthSdkCredentials) validAccessToken() (string, error) {
	f.mutex.Lock()
	if f.AccessToken != nil && f.AccessToken.IsValid() {
		token := f.AccessToken.Token
		f.mutex.Unlock()
		return token, nil
	}
	if renewal := f.renewal; renewal != nil {
		select {
		case <-renewal.done:
			// A finished renewal is cleared before done is closed, so this value is a copy taken
			// during a renewal: ignore it and renew.
		default:
			f.mutex.Unlock()
			<-renewal.done
			return renewal.token, renewal.err
		}
	}
	renewal := &tokenRenewal{done: make(chan struct{})}
	f.renewal = renewal
	f.mutex.Unlock()

	return f.renew(renewal)
}

// renew sends the token request and publishes its result to the callers waiting on renewal. If
// the request panics, the waiters get an error instead of waiting forever.
func (f *OAuthSdkCredentials) renew(renewal *tokenRenewal) (token string, err error) {
	err = errors.CheckoutAuthorizationError("OAuth token request failed")
	defer func() {
		f.mutex.Lock()
		renewal.token, renewal.err = token, err
		f.renewal = nil
		f.mutex.Unlock()
		close(renewal.done)
	}()

	accessToken, err := f.requestAccessToken()
	if err != nil {
		return "", err
	}
	f.mutex.Lock()
	f.AccessToken = accessToken
	f.mutex.Unlock()
	return accessToken.Token, nil
}

// requestAccessToken sends the token request and returns the new token. It returns an error when
// the response has no access_token, or an expires_in that is not positive or does not fit in a
// time.Duration.
func (f *OAuthSdkCredentials) requestAccessToken() (*OAuthAccessToken, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", f.ClientId)
	data.Set("client_secret", f.ClientSecret)
	data.Set("scope", strings.Join(f.Scopes, " "))

	client := f.tokenHttpClient()
	ctx := context.Background()
	if client.Timeout <= 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, tokenRequestTimeout)
		defer cancel()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, f.AuthorizationUri, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	f.Log.Printf("post: %s", f.AuthorizationUri)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var oauthResp OAuthServiceResponse
	err = json.Unmarshal(body, &oauthResp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		var oauthErr errors.CheckoutOAuthError
		err = json.Unmarshal(body, &oauthErr)
		if err != nil {
			return nil, err
		}
		return nil, errors.CheckoutAuthorizationError(oauthErr.Description)
	}

	if oauthResp.AccessToken == "" || oauthResp.ExpiresIn <= 0 || oauthResp.ExpiresIn > maxExpiresIn {
		return nil, errors.CheckoutAuthorizationError("OAuth token response is invalid")
	}

	accessToken := &OAuthAccessToken{
		Token:          oauthResp.AccessToken,
		ExpirationDate: time.Now().Add(oauthResp.ExpiresIn * time.Second),
	}
	return accessToken, nil
}

// maxExpiresIn is the largest expires_in, in seconds, that fits in a time.Duration.
const maxExpiresIn = time.Duration(math.MaxInt64 / int64(time.Second))

// tokenRequestTimeout bounds the token request when its client sets no Timeout (zero or negative).
const tokenRequestTimeout = time.Duration(5) * time.Second

// tokenHttpClient returns HttpClient, or a client with a 5 second timeout when it is nil.
func (f *OAuthSdkCredentials) tokenHttpClient() *http.Client {
	if f.HttpClient != nil {
		return f.HttpClient
	}
	return &http.Client{Timeout: tokenRequestTimeout}
}

// IsValid reports whether the token is set and not expired.
func (t *OAuthAccessToken) IsValid() bool {
	if t.Token == "" {
		return false
	}
	return t.ExpirationDate.After(time.Now())
}
