package configuration

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/errors"
)

const tokenResponse = `{"access_token":"token","expires_in":3600,"token_type":"Bearer"}`

// tokenHandler answers every token request with a valid token and counts the requests.
func tokenHandler(requests *int32) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(requests, 1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(tokenResponse))
	})
}

// The test server's certificate is trusted only by server.Client(), so a successful token request
// proves which client sent it.
func TestTokenRequestUsesTheGivenHttpClient(t *testing.T) {
	var requests int32
	server := httptest.NewTLSServer(tokenHandler(&requests))
	defer server.Close()

	credentials, err := NewOAuthSdkCredentialsWithHttpClient(
		"client_id", "client_secret", server.URL, nil, server.Client(), nil)

	assert.NoError(t, err)
	assert.NotNil(t, credentials)
	assert.Equal(t, "token", credentials.AccessToken.Token)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requests))
}

func TestTokenRequestUsesTheHttpClientField(t *testing.T) {
	var requests int32
	server := httptest.NewTLSServer(tokenHandler(&requests))
	defer server.Close()

	credentials := &OAuthSdkCredentials{
		ClientId:         "client_id",
		ClientSecret:     "client_secret",
		AuthorizationUri: server.URL,
		HttpClient:       server.Client(),
		Log:              DefaultLogger(),
	}

	err := credentials.GetAccessToken()

	assert.NoError(t, err)
	assert.Equal(t, "token", credentials.AccessToken.Token)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requests))
}

func TestTokenRequestWithoutHttpClientDoesNotTrustCustomCa(t *testing.T) {
	var requests int32
	server := httptest.NewTLSServer(tokenHandler(&requests))
	defer server.Close()

	credentials, err := NewOAuthSdkCredentials("client_id", "client_secret", server.URL, nil, nil)

	assert.Nil(t, credentials)
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "certificate")
	}
	assert.Equal(t, int32(0), atomic.LoadInt32(&requests))
}

// On expiry, concurrent callers share one token renewal.
// Run with -race to also check the cached token is never read and written concurrently.
func TestConcurrentGetAuthorizationSendsOneTokenRequest(t *testing.T) {
	var requests int32
	server := httptest.NewServer(tokenHandler(&requests))
	defer server.Close()

	credentials := &OAuthSdkCredentials{
		ClientId:         "client_id",
		ClientSecret:     "client_secret",
		AuthorizationUri: server.URL,
		AccessToken:      &OAuthAccessToken{Token: "expired", ExpirationDate: time.Now().Add(-time.Minute)},
		Log:              DefaultLogger(),
	}

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			authorization, err := credentials.GetAuthorization(OAuth)
			if assert.NoError(t, err) {
				assert.Equal(t, "token", authorization.Credential)
			}
		}()
	}
	wg.Wait()

	assert.Equal(t, int32(1), atomic.LoadInt32(&requests))
}

// delayedTokenHandler answers with a valid token after delay, or returns without answering when
// the client gives up or stop is closed.
func delayedTokenHandler(delay time.Duration, stop <-chan struct{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(delay):
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(tokenResponse))
		case <-r.Context().Done():
		case <-stop:
		}
	})
}

// A client without Timeout would wait forever, and every caller waits behind the renewal, so the
// token request keeps the 5 second bound.
func TestTokenRequestWithClientWithoutTimeoutIsBounded(t *testing.T) {
	t.Parallel()
	stop := make(chan struct{})
	server := httptest.NewServer(delayedTokenHandler(10*time.Second, stop))
	defer server.Close()
	defer close(stop)

	start := time.Now()
	credentials, err := NewOAuthSdkCredentialsWithHttpClient(
		"client_id", "client_secret", server.URL, nil, &http.Client{}, nil)
	elapsed := time.Since(start)

	assert.Nil(t, credentials)
	assert.Error(t, err)
	assert.True(t, elapsed < 8*time.Second, "token request took %s", elapsed)
}

func TestTokenRequestKeepsTheClientTimeout(t *testing.T) {
	t.Parallel()
	stop := make(chan struct{})
	server := httptest.NewServer(delayedTokenHandler(6*time.Second, stop))
	defer server.Close()
	defer close(stop)

	credentials, err := NewOAuthSdkCredentialsWithHttpClient(
		"client_id", "client_secret", server.URL, nil, &http.Client{Timeout: 20 * time.Second}, nil)

	assert.NoError(t, err)
	if assert.NotNil(t, credentials) {
		assert.Equal(t, "token", credentials.AccessToken.Token)
	}
}

func TestTokenRequestWithNegativeClientTimeoutIsBounded(t *testing.T) {
	t.Parallel()
	stop := make(chan struct{})
	server := httptest.NewServer(delayedTokenHandler(10*time.Second, stop))
	defer server.Close()
	defer close(stop)

	start := time.Now()
	credentials, err := NewOAuthSdkCredentialsWithHttpClient(
		"client_id", "client_secret", server.URL, nil, &http.Client{Timeout: -1}, nil)
	elapsed := time.Since(start)

	assert.Nil(t, credentials)
	assert.Error(t, err)
	assert.True(t, elapsed < 8*time.Second, "token request took %s", elapsed)
}

// Callers waiting for a renewal get its result, also when it fails, instead of each retrying in turn.
func TestConcurrentCallersShareAFailedRenewal(t *testing.T) {
	t.Parallel()
	var requests int32
	stop := make(chan struct{})
	slow := delayedTokenHandler(10*time.Second, stop)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requests, 1)
		slow.ServeHTTP(w, r)
	}))
	defer server.Close()
	defer close(stop)

	credentials := &OAuthSdkCredentials{
		ClientId:         "client_id",
		ClientSecret:     "client_secret",
		AuthorizationUri: server.URL,
		HttpClient:       &http.Client{},
		Log:              DefaultLogger(),
	}

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := credentials.GetAuthorization(OAuth)
			assert.Error(t, err)
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)

	assert.True(t, elapsed < 8*time.Second, "callers waited %s", elapsed)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requests))
}

// A finished renewal is never left on its own credentials, only on a copy of them; it must not pin
// its result there.
func TestFinishedRenewalDoesNotPinItsResult(t *testing.T) {
	var requests int32
	server := httptest.NewServer(tokenHandler(&requests))
	defer server.Close()

	finished := &tokenRenewal{done: make(chan struct{}), err: errors.CheckoutAuthorizationError("old")}
	close(finished.done)
	credentials := &OAuthSdkCredentials{
		ClientId:         "client_id",
		ClientSecret:     "client_secret",
		AuthorizationUri: server.URL,
		Log:              DefaultLogger(),
		renewal:          finished,
	}

	authorization, err := credentials.GetAuthorization(OAuth)

	assert.NoError(t, err)
	if assert.NotNil(t, authorization) {
		assert.Equal(t, "token", authorization.Credential)
	}
	assert.Equal(t, int32(1), atomic.LoadInt32(&requests))
}
