package client

import (
	"context"
	stderrors "errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

func TestConnectionErrorIsATransportError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { jsonOK(w) }))
	baseUri := server.URL
	server.Close()

	for _, telemetry := range []bool{false, true} {
		client := newTestClient(baseUri)
		client.EnableTelemetry = telemetry
		client.RequestMetricsQueue = *common.NewTelemetryQueue()

		var resp common.IdResponse
		err := client.Get("/test", testAuth(), nil, &resp)

		var transportErr errors.CheckoutTransportError
		if assert.True(t, stderrors.As(err, &transportErr), "telemetry=%v: got %T", telemetry, err) {
			assert.False(t, transportErr.Timeout())
			var urlErr *url.Error
			assert.True(t, stderrors.As(err, &urlErr), "the net/http error stays reachable")
		}
	}
}

func TestClientTimeoutIsATransportError(t *testing.T) {
	stop := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(time.Second):
			jsonOK(w)
		case <-r.Context().Done():
		case <-stop:
		}
	}))
	defer server.Close()
	defer close(stop)

	client := newTestClient(server.URL)
	client.HttpClient.Timeout = 50 * time.Millisecond

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), nil, &resp)

	if assert.IsType(t, errors.CheckoutTransportError{}, err) {
		assert.True(t, err.(errors.CheckoutTransportError).Timeout())
	}
}

func TestContextDeadlineIsATransportError(t *testing.T) {
	stop := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-time.After(time.Second):
			jsonOK(w)
		case <-r.Context().Done():
		case <-stop:
		}
	}))
	defer server.Close()
	defer close(stop)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	var resp common.IdResponse
	err := newTestClient(server.URL).GetWithContext(ctx, "/test", testAuth(), nil, &resp)

	if assert.IsType(t, errors.CheckoutTransportError{}, err) {
		assert.True(t, err.(errors.CheckoutTransportError).Timeout())
		assert.True(t, stderrors.Is(err, context.DeadlineExceeded))
	}
}

func TestApiErrorIsNotATransportError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	var resp common.IdResponse
	err := newTestClient(server.URL).Get("/test", testAuth(), nil, &resp)

	assert.IsType(t, errors.CheckoutAPIError{}, err)
}

type failingBody struct{}

func (failingBody) Read([]byte) (int, error) {
	return 0, stderrors.New("connection reset while reading")
}
func (failingBody) Close() error { return nil }

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestResponseBodyReadErrorIsATransportError(t *testing.T) {
	for _, telemetry := range []bool{false, true} {
		client := newTestClient("https://api.example.test")
		client.EnableTelemetry = telemetry
		client.RequestMetricsQueue = *common.NewTelemetryQueue()
		client.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: http.Header{}, Body: failingBody{}, Request: r}, nil
		})

		var resp common.IdResponse
		err := client.Get("/test", testAuth(), nil, &resp)

		if assert.IsType(t, errors.CheckoutTransportError{}, err, "telemetry=%v", telemetry) {
			assert.Contains(t, err.Error(), "connection reset while reading")
			assert.False(t, err.(errors.CheckoutTransportError).Timeout())
		}
	}
}

func TestCancelledContextIsATransportError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { jsonOK(w) }))
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var resp common.IdResponse
	err := newTestClient(server.URL).GetWithContext(ctx, "/test", testAuth(), nil, &resp)

	if assert.IsType(t, errors.CheckoutTransportError{}, err) {
		assert.True(t, stderrors.Is(err, context.Canceled))
		assert.False(t, err.(errors.CheckoutTransportError).Timeout())
	}
}

func TestTlsErrorIsATransportError(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { jsonOK(w) }))
	defer server.Close()

	// The default client does not trust the test server certificate.
	var resp common.IdResponse
	err := newTestClient(server.URL).Post("/test", testAuth(), map[string]string{"k": "v"}, &resp, nil)

	if assert.IsType(t, errors.CheckoutTransportError{}, err) {
		var urlErr *url.Error
		assert.True(t, stderrors.As(err, &urlErr))
		assert.NotContains(t, err.Error(), testAuth().Credential, "no credential in the message")
	}
}

// A merchant RoundTripper that returns a CheckoutTransportError is wrapped by net/http in a
// *url.Error; the caller still gets the SDK type at the top.
func TestTransportErrorFromTheMerchantRoundTripperIsATransportError(t *testing.T) {
	client := newTestClient("https://api.example.test")
	client.HttpClient.Transport = roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return nil, errors.CheckoutTransportError{Err: stderrors.New("merchant transport")}
	})

	var resp common.IdResponse
	err := client.Get("/test", testAuth(), nil, &resp)

	assert.IsType(t, errors.CheckoutTransportError{}, err)
	var urlErr *url.Error
	assert.True(t, stderrors.As(err, &urlErr))
}
