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

func TestTokenResponseWithoutTokenOrLifetimeIsRejected(t *testing.T) {
	cases := map[string]string{
		"empty object":           `{}`,
		"null":                   `null`,
		"empty token":            `{"access_token":"","expires_in":3600}`,
		"missing expires_in":     `{"access_token":"token"}`,
		"zero expires_in":        `{"access_token":"token","expires_in":0}`,
		"negative expires_in":    `{"access_token":"token","expires_in":-5}`,
		"overflowing expires_in": `{"access_token":"token","expires_in":10000000000}`,
	}
	for name, body := range cases {
		body := body
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(body))
			}))
			defer server.Close()

			credentials, err := NewOAuthSdkCredentials("client_id", "client_secret", server.URL, nil, nil)

			assert.Nil(t, credentials)
			assert.IsType(t, errors.CheckoutAuthorizationError(""), err)
		})
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// waitForRenewal returns once a token request is in flight on credentials.
func waitForRenewal(t *testing.T, credentials *OAuthSdkCredentials) {
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		credentials.mutex.Lock()
		inFlight := credentials.renewal != nil
		credentials.mutex.Unlock()
		if inFlight {
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatal("no token request started")
}

func TestWaitersGetAnErrorWhenTheTokenRequestPanics(t *testing.T) {
	release := make(chan struct{})
	var calls int32
	credentials := &OAuthSdkCredentials{
		ClientId:         "client_id",
		ClientSecret:     "client_secret",
		AuthorizationUri: "https://access.checkout.example/connect/token",
		HttpClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			if atomic.AddInt32(&calls, 1) > 1 {
				return nil, errors.CheckoutAPIError{Status: "later request"}
			}
			<-release
			panic("transport failure")
		})},
		Log: DefaultLogger(),
	}

	leaderDone := make(chan interface{})
	go func() {
		defer func() { leaderDone <- recover() }()
		_, _ = credentials.GetAuthorization(OAuth)
	}()
	waitForRenewal(t, credentials)

	waiterErrors := make(chan error, 3)
	for i := 0; i < 3; i++ {
		go func() {
			_, err := credentials.GetAuthorization(OAuth)
			waiterErrors <- err
		}()
	}
	close(release)

	assert.Equal(t, "transport failure", <-leaderDone)
	for i := 0; i < 3; i++ {
		select {
		case err := <-waiterErrors:
			assert.Error(t, err)
		case <-time.After(5 * time.Second):
			t.Fatal("a waiter is still blocked after the token request panicked")
		}
	}
}

func TestLockIsNotHeldDuringTheTokenRequest(t *testing.T) {
	started := make(chan struct{}, 1)
	release := make(chan struct{})
	stop := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started <- struct{}{}
		select {
		case <-release:
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(tokenResponse))
		case <-stop:
		}
	}))
	defer server.Close()
	defer close(stop)

	credentials := &OAuthSdkCredentials{
		ClientId:         "client_id",
		ClientSecret:     "client_secret",
		AuthorizationUri: server.URL,
		Log:              DefaultLogger(),
	}
	done := make(chan error, 1)
	go func() {
		_, err := credentials.GetAuthorization(OAuth)
		done <- err
	}()
	select {
	case <-started:
	case <-time.After(5 * time.Second):
		t.Fatal("no token request started")
	}

	locked := make(chan struct{})
	go func() {
		credentials.mutex.Lock()
		credentials.mutex.Unlock()
		close(locked)
	}()
	select {
	case <-locked:
	case <-time.After(2 * time.Second):
		t.Fatal("the lock is held while the token request is in flight")
	}

	close(release)
	assert.NoError(t, <-done)
}

func TestTokenResponseWithTheLargestExpiresInIsAccepted(t *testing.T) {
	var requests int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&requests, 1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"access_token":"token","expires_in":9223372036}`))
	}))
	defer server.Close()

	credentials, err := NewOAuthSdkCredentials("client_id", "client_secret", server.URL, nil, nil)
	if !assert.NoError(t, err) {
		return
	}
	_, err = credentials.GetAuthorization(OAuth)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&requests))
}

func TestTokenRequestWithoutHttpClientIsBounded(t *testing.T) {
	stop := make(chan struct{})
	server := httptest.NewServer(delayedTokenHandler(10*time.Second, stop))
	defer server.Close()
	defer close(stop)

	start := time.Now()
	credentials, err := NewOAuthSdkCredentialsWithHttpClient("client_id", "client_secret", server.URL, nil, nil, nil)
	elapsed := time.Since(start)

	assert.Nil(t, credentials)
	assert.Error(t, err)
	assert.True(t, elapsed >= 4*time.Second && elapsed < 8*time.Second, "token request took %s", elapsed)
}
