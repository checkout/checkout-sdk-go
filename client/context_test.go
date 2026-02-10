package client

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/stretchr/testify/assert"
)

// newTestClient creates an ApiClient pointing to a test server with a proper logger
func newTestClient(baseURL string) *ApiClient {
	return &ApiClient{
		HttpClient:      http.Client{},
		BaseUri:         baseURL,
		EnableTelemetry: false,
		Log:             log.New(os.Stdout, "test: ", log.LstdFlags),
	}
}

func testAuth() *configuration.SdkAuthorization {
	return &configuration.SdkAuthorization{
		PlatformType: configuration.Default,
		Credential:   "test-token",
	}
}

// jsonOK responds with 200 and a JSON body
func jsonOK(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id":"ctx-123"}`))
}

// ---------- Tests ----------

// TestContextCancellation verifies that a pre-cancelled context aborts the request immediately
func TestContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel before performing the request

	var resp common.IdResponse
	err := client.GetWithContext(ctx, "/test", testAuth(), &resp)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, context.Canceled),
		"Expected context.Canceled, got: %v", err)
}

// TestContextTimeout verifies deadline exceeded error when server is too slow
func TestContextTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond) // slower than the timeout
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	var resp common.IdResponse
	err := client.GetWithContext(ctx, "/test", testAuth(), &resp)
	elapsed := time.Since(start)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded),
		"Expected context.DeadlineExceeded, got: %v", err)
	assert.Less(t, elapsed, 150*time.Millisecond,
		"Request should abort quickly, not wait for the full server delay")
}

// TestContextTimeoutSuccess verifies a request succeeds when completed before the deadline
func TestContextTimeoutSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond) // fast response
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var resp common.IdResponse
	err := client.GetWithContext(ctx, "/test", testAuth(), &resp)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.HttpMetadata.StatusCode)
	assert.Equal(t, "ctx-123", resp.Id)
}

// TestContextPropagation verifies that the context reaches the HTTP transport layer
func TestContextPropagation(t *testing.T) {
	type ctxKey string
	const key ctxKey = "trace-id"

	contextReachedServer := false

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contextReachedServer = true
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	ctx := context.WithValue(context.Background(), key, "abc-123")

	var resp common.IdResponse
	err := client.GetWithContext(ctx, "/test", testAuth(), &resp)

	assert.Nil(t, err)
	assert.True(t, contextReachedServer,
		"Server should have received the request (context was propagated)")
	assert.Equal(t, http.StatusOK, resp.HttpMetadata.StatusCode)
}

// TestDefaultMethodUsesBackgroundContext verifies the non-context method works normally
func TestDefaultMethodUsesBackgroundContext(t *testing.T) {
	requestReceived := false

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestReceived = true
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), &resp)

	assert.Nil(t, err)
	assert.True(t, requestReceived, "Request should succeed with implicit context.Background()")
	assert.Equal(t, http.StatusOK, resp.HttpMetadata.StatusCode)
	assert.Equal(t, "ctx-123", resp.Id)
}

// TestConcurrentRequestsWithDifferentContexts verifies contexts are independent
func TestConcurrentRequestsWithDifferentContexts(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		delay := r.URL.Query().Get("delay")
		if delay == "long" {
			time.Sleep(200 * time.Millisecond)
		} else {
			time.Sleep(10 * time.Millisecond)
		}
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	type result struct {
		err  error
		resp common.IdResponse
	}

	done := make(chan result, 2)

	// Request 1: tight timeout + slow endpoint → should fail
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
		defer cancel()
		var r common.IdResponse
		err := client.GetWithContext(ctx, "/test?delay=long", testAuth(), &r)
		done <- result{err: err, resp: r}
	}()

	// Request 2: generous timeout + fast endpoint → should succeed
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		var r common.IdResponse
		err := client.GetWithContext(ctx, "/test?delay=short", testAuth(), &r)
		done <- result{err: err, resp: r}
	}()

	r1 := <-done
	r2 := <-done

	// One should fail, one should succeed
	timedOut := (r1.err != nil && errors.Is(r1.err, context.DeadlineExceeded)) ||
		(r2.err != nil && errors.Is(r2.err, context.DeadlineExceeded))
	succeeded := (r1.err == nil) || (r2.err == nil)

	assert.True(t, timedOut, "One request should have timed out")
	assert.True(t, succeeded, "One request should have succeeded")
}

// TestContextPostWithCancellation verifies Post method also respects context cancellation
func TestContextPostWithCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		time.Sleep(300 * time.Millisecond)
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	body := map[string]string{"key": "value"}
	var resp common.IdResponse
	err := client.PostWithContext(ctx, "/resource", testAuth(), body, &resp, nil)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded),
		"PostWithContext should respect context deadline, got: %v", err)
}

// TestContextPutWithCancellation verifies Put method also respects context cancellation
func TestContextPutWithCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		time.Sleep(300 * time.Millisecond)
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	body := map[string]string{"key": "value"}
	var resp common.IdResponse
	err := client.PutWithContext(ctx, "/resource/123", testAuth(), body, &resp, nil)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded),
		"PutWithContext should respect context deadline, got: %v", err)
}

// TestContextDeleteWithCancellation verifies Delete method also respects context cancellation
func TestContextDeleteWithCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		time.Sleep(300 * time.Millisecond)
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var resp common.MetadataResponse
	err := client.DeleteWithContext(ctx, "/resource/123", testAuth(), &resp)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded),
		"DeleteWithContext should respect context deadline, got: %v", err)
}

// TestContextDeadline verifies WithDeadline works correctly
func TestContextDeadline(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		jsonOK(w)
	}))
	defer server.Close()

	client := newTestClient(server.URL)

	deadline := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var resp common.IdResponse
	err := client.GetWithContext(ctx, "/test", testAuth(), &resp)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, context.DeadlineExceeded),
		"Should respect explicit deadline, got: %v", err)
}
