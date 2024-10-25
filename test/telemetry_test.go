package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/stretchr/testify/assert"
)

// mockHTTPClient is our custom HTTP client that will track requests
type mockHTTPClient struct {
	server           *httptest.Server
	requestsReceived []*http.Request
}

func newMockHTTPClient() *mockHTTPClient {
	mock := &mockHTTPClient{
		requestsReceived: make([]*http.Request, 0),
	}

	mock.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mock.requestsReceived = append(mock.requestsReceived, r)
		w.WriteHeader(http.StatusOK)
	}))

	return mock
}

// Client returns an HTTP client that will use our mock server
func (m *mockHTTPClient) Client() *http.Client {
	return &http.Client{
		Timeout: 2 * time.Second,
	}
}

// Close shuts down the test server
func (m *mockHTTPClient) Close() {
	m.server.Close()
}

// CountRequestsWithHeader counts how many requests contained a specific header
func (m *mockHTTPClient) CountRequestsWithHeader(header string) int {
	count := 0
	for _, req := range m.requestsReceived {
		if req.Header.Get(header) != "" {
			count++
		}
	}
	return count
}

func TestShouldSendTelemetryByDefault(t *testing.T) {
	mock := newMockHTTPClient()
	defer mock.Close()

	checkoutApi, err := checkout.Builder().
		Previous().
		WithPublicKey("CHECKOUT_PREVIOUS_PUBLIC_KEY").
		WithSecretKey("CHECKOUT_PREVIOUS_SECRET_KEY").
		WithEnvironment(configuration.Sandbox()).
		WithHttpClient(mock.Client()).
		Build()

	assert.NoError(t, err)
	assert.NotNil(t, checkoutApi)

	for i := 0; i < 3; i++ {
		_, err := checkoutApi.Events.RetrieveAllEventTypes()
		assert.NoError(t, err)
	}

	// Verify telemetry header was present in requests
	telemetryHeaderCount := mock.CountRequestsWithHeader("cko-sdk-telemetry")
	assert.Equal(t, 2, telemetryHeaderCount, "Expected exactly 2 requests to contain the telemetry header")
}

func TestShouldNotSendTelemetryWhenOptedOut(t *testing.T) {
	mock := newMockHTTPClient()
	defer mock.Close()

	checkoutApi, err := checkout.Builder().
		Previous().
		WithPublicKey("CHECKOUT_PREVIOUS_PUBLIC_KEY").
		WithSecretKey("CHECKOUT_PREVIOUS_SECRET_KEY").
		WithEnableTelemetry(false).
		WithEnvironment(configuration.Sandbox()).
		WithHttpClient(mock.Client()).
		Build()

	assert.NoError(t, err)
	assert.NotNil(t, checkoutApi)

	for i := 0; i < 3; i++ {
		_, err := checkoutApi.Events.RetrieveAllEventTypes()
		assert.NoError(t, err)
	}

	// Verify no telemetry headers were sent
	telemetryHeaderCount := mock.CountRequestsWithHeader("cko-sdk-telemetry")
	assert.Equal(t, 0, telemetryHeaderCount, "Expected no requests to contain the telemetry header")
}
