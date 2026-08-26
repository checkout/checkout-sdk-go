package client

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

// fastRetry returns a retry configuration with negligible backoff so tests do
// not sleep for real durations.
func fastRetry() *configuration.RetryConfiguration {
	return &configuration.RetryConfiguration{
		MaxRetries: 2,
		MinBackoff: time.Millisecond,
		MaxBackoff: 5 * time.Millisecond,
	}
}

// newRetryClient builds an ApiClient with retries enabled pointing at baseURL.
func newRetryClient(baseURL string, retry *configuration.RetryConfiguration) *ApiClient {
	return &ApiClient{
		HttpClient:      http.Client{},
		BaseUri:         baseURL,
		EnableTelemetry: false,
		Log:             log.New(os.Stdout, "test: ", log.LstdFlags),
		Retry:           retry,
	}
}

// roundTripFunc adapts a function to an http.RoundTripper.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestRetrySucceedsAfterTransientFailure(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonOK(w)
	}))
	defer server.Close()

	client := newRetryClient(server.URL, fastRetry())

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, int32(2), atomic.LoadInt32(&calls), "expected one retry after the 500")
	assert.Equal(t, "ctx-123", resp.Id)
}

func TestRetryExhaustsAndReturnsLastError(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusBadGateway)
	}))
	defer server.Close()

	client := newRetryClient(server.URL, fastRetry())

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), &resp)

	assert.NotNil(t, err)
	assert.Equal(t, int32(3), atomic.LoadInt32(&calls), "expected MaxRetries+1 attempts")
}

func TestNoRetryOnClientError(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	client := newRetryClient(server.URL, fastRetry())

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), &resp)

	assert.NotNil(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&calls), "4xx (except 409) must not be retried")
}

func TestRetryDisabledByDefault(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Retry is nil: historical single-attempt behaviour must be preserved.
	client := newTestClient(server.URL)

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), &resp)

	assert.NotNil(t, err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&calls), "retries must be opt-in")
}

func TestRetryGeneratesStableIdempotencyKeyOnPost(t *testing.T) {
	var keys []string
	var bodies []string
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys = append(keys, r.Header.Get("Cko-Idempotency-Key"))
		buf := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(buf)
		bodies = append(bodies, string(buf))
		if atomic.AddInt32(&calls, 1) == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		jsonOK(w)
	}))
	defer server.Close()

	client := newRetryClient(server.URL, fastRetry())

	body := map[string]string{"amount": "100"}
	var resp common.IdResponse
	err := client.Post("/payments", testAuth(), body, &resp, nil)

	assert.Nil(t, err)
	assert.Len(t, keys, 2)
	assert.NotEmpty(t, keys[0], "an idempotency key should be generated when retries are enabled")
	assert.Equal(t, keys[0], keys[1], "the same key must be sent on every attempt")
	assert.Equal(t, bodies[0], bodies[1], "the request body must be replayed on retry")
}

func TestRetryPreservesCallerIdempotencyKey(t *testing.T) {
	var keys []string
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keys = append(keys, r.Header.Get("Cko-Idempotency-Key"))
		if atomic.AddInt32(&calls, 1) == 1 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonOK(w)
	}))
	defer server.Close()

	client := newRetryClient(server.URL, fastRetry())

	supplied := "caller-supplied-key"
	var resp common.IdResponse
	err := client.Post("/payments", testAuth(), map[string]string{"k": "v"}, &resp, &supplied)

	assert.Nil(t, err)
	assert.Equal(t, []string{supplied, supplied}, keys, "a caller-supplied key must not be overwritten")
}

func TestNoIdempotencyKeyWhenRetriesDisabled(t *testing.T) {
	var key string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key = r.Header.Get("Cko-Idempotency-Key")
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	var resp common.IdResponse
	err := client.Post("/payments", testAuth(), map[string]string{"k": "v"}, &resp, nil)

	assert.Nil(t, err)
	assert.Empty(t, key, "no key should be generated when retries are disabled")
}

func TestRetryStopsWhenContextCancelledDuringBackoff(t *testing.T) {
	var calls int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// Long backoff so the deadline fires while we are waiting to retry.
	client := newRetryClient(server.URL, &configuration.RetryConfiguration{
		MaxRetries: 5,
		MinBackoff: 200 * time.Millisecond,
		MaxBackoff: time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var resp common.IdResponse
	err := client.GetWithContext(ctx, "/test", testAuth(), &resp)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded), "expected deadline error, got: %v", err)
	assert.Equal(t, int32(1), atomic.LoadInt32(&calls), "should not retry once the context is done")
}

func TestRetryOnConnectionError(t *testing.T) {
	var calls int32
	transport := roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if atomic.AddInt32(&calls, 1) == 1 {
			return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: errors.New("connection reset by peer")}
		}
		rec := httptest.NewRecorder()
		jsonOK(rec)
		return rec.Result(), nil
	})

	client := &ApiClient{
		HttpClient: http.Client{Transport: transport},
		BaseUri:    "http://example.test",
		Log:        log.New(os.Stdout, "test: ", log.LstdFlags),
		Retry:      fastRetry(),
	}

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, int32(2), atomic.LoadInt32(&calls), "connection errors should be retried")
}

func TestShouldRetry(t *testing.T) {
	cfg := &configuration.RetryConfiguration{MaxRetries: 2}
	tests := []struct {
		name    string
		status  int
		err     error
		attempt int
		want    bool
	}{
		{name: "500 retried", status: http.StatusInternalServerError, want: true},
		{name: "409 retried", status: http.StatusConflict, want: true},
		{name: "429 retried", status: http.StatusTooManyRequests, want: true},
		{name: "400 not retried", status: http.StatusBadRequest, want: false},
		{name: "404 not retried", status: http.StatusNotFound, want: false},
		{name: "200 not retried", status: http.StatusOK, want: false},
		{name: "attempts exhausted", status: http.StatusInternalServerError, attempt: 2, want: false},
		{name: "context canceled not retried", err: context.Canceled, want: false},
		{name: "deadline not retried", err: context.DeadlineExceeded, want: false},
		{name: "generic connection error retried", err: errors.New("broken pipe"), want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var resp *http.Response
			if tt.err == nil {
				resp = &http.Response{StatusCode: tt.status}
			}
			got := shouldRetry(resp, tt.err, tt.attempt, cfg)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBackoffGrowsAndIsCapped(t *testing.T) {
	cfg := &configuration.RetryConfiguration{
		MinBackoff: 500 * time.Millisecond,
		MaxBackoff: 5 * time.Second,
	}

	// No Retry-After: value stays within jittered exponential bounds and is capped.
	for attempt := 0; attempt < 6; attempt++ {
		d := backoff(attempt, cfg, 0)
		base := cfg.MinBackoff + cfg.MinBackoff*time.Duration(attempt*attempt)
		if base > cfg.MaxBackoff {
			base = cfg.MaxBackoff
		}
		assert.GreaterOrEqual(t, d, time.Duration(float64(base)*0.75), "attempt %d below jitter floor", attempt)
		assert.LessOrEqual(t, d, base, "attempt %d above jittered value", attempt)
	}

	// Retry-After takes precedence and is capped at MaxBackoff.
	assert.Equal(t, 2*time.Second, backoff(0, cfg, 2*time.Second))
	assert.Equal(t, cfg.MaxBackoff, backoff(0, cfg, 30*time.Second))
}

func TestRetryAfterDelay(t *testing.T) {
	resp := &http.Response{Header: http.Header{"Retry-After": []string{"3"}}}
	assert.Equal(t, 3*time.Second, retryAfterDelay(resp))

	assert.Equal(t, time.Duration(0), retryAfterDelay(&http.Response{Header: http.Header{}}))
	assert.Equal(t, time.Duration(0), retryAfterDelay(nil))
}
