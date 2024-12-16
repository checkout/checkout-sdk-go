package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/abc"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/stretchr/testify/assert"
)

// mockHTTPClient is a helper struct to create a mock HTTP server and capture incoming requests.
type mockHTTPClient struct {
	server           *httptest.Server
	requestsReceived []*http.Request
}

type telemetry struct {
	PrevRequestId       string `json:"prev_request_id"`
	RequestId           string `json:"request_id"`
	PrevRequestDuration int    `json:"prev_request_duration"`
}

// newMockHTTPClient initializes a new mockHTTPClient with a TLS test server, optionally introducing a delay.
func newMockHTTPClient(delay time.Duration) *mockHTTPClient {
	mock := &mockHTTPClient{
		requestsReceived: []*http.Request{},
	}

	// Create a new TLS test server with a handler that records incoming requests and introduces a delay.
	mock.server = httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate server processing delay.
		time.Sleep(delay)
		// Clone and store the request for inspection.
		mock.requestsReceived = append(mock.requestsReceived, r.Clone(context.Background()))
		// Respond with a 200 OK status.
		w.WriteHeader(http.StatusOK)
	}))
	return mock
}

// Client returns an HTTP client configured to communicate with the mock server.
func (m *mockHTTPClient) Client() *http.Client {
	return m.server.Client()
}

// Close shuts down the mock server.
func (m *mockHTTPClient) Close() {
	m.server.Close()
}

// CountRequestsWithHeader counts the number of recorded requests that include a specific header.
func (m *mockHTTPClient) CountRequestsWithHeader(header string) int {
	count := 0
	headerKey := http.CanonicalHeaderKey(header) // Normalize header name.
	for _, req := range m.requestsReceived {
		if _, ok := req.Header[headerKey]; ok {
			count++
		}
	}
	return count
}

// ValidateTelemetryHeadersFormat
func (m *mockHTTPClient) ValidateTelemetryHeadersFormat(header string) error {
	headerKey := http.CanonicalHeaderKey(header) // Normalize header name.
	for _, req := range m.requestsReceived {
		if _, ok := req.Header[headerKey]; ok {
			var telemetryObj telemetry
			headerStr := req.Header[headerKey][0]
			decoder := json.NewDecoder(strings.NewReader(headerStr))
			decoder.DisallowUnknownFields()
			if err := decoder.Decode(&telemetryObj); err != nil {
				return err
			}
		}
	}
	return nil
}

// createMockEnvironment creates a mock environment using the mock server's URL.
func createMockEnvironment(mockServerURL string) *configuration.CheckoutEnv {
	return configuration.NewEnvironment(
		mockServerURL, // baseUri
		mockServerURL, // authorizationUri
		mockServerURL, // filesUri
		mockServerURL, // transfersUri
		mockServerURL, // balancesUri
		true,          // isSandbox
	)
}

// createCheckoutApi initializes the Checkout API client with the given parameters.
func createCheckoutApi(mock *mockHTTPClient, enableTelemetry bool) (*abc.Api, error) {
	// Retrieve public and secret keys from environment variables.
	publicKey := os.Getenv("CHECKOUT_PREVIOUS_PUBLIC_KEY")
	secretKey := os.Getenv("CHECKOUT_PREVIOUS_SECRET_KEY")

	// Ensure that the keys are set.
	if publicKey == "" || secretKey == "" {
		return nil, fmt.Errorf("public and secret keys must be set in environment variables")
	}

	// Build the Checkout API client with the mock environment and client.
	builder := checkout.Builder().
		Previous().
		WithPublicKey(publicKey).
		WithSecretKey(secretKey).
		WithEnvironment(createMockEnvironment(mock.server.URL)).
		WithHttpClient(mock.Client())

	// Enable or disable telemetry based on the parameter.
	if !enableTelemetry {
		builder = builder.WithEnableTelemetry(false)
	}

	return builder.Build()
}

// TestShouldSendTelemetryByDefault verifies that telemetry headers are sent by default.
func TestShouldSendTelemetryByDefault(t *testing.T) {
	mock := newMockHTTPClient(0) // No delay.
	defer mock.Close()

	// Initialize the Checkout API client with telemetry enabled (default).
	checkoutApi, err := createCheckoutApi(mock, true)
	assert.NoError(t, err)
	assert.NotNil(t, checkoutApi)

	// Make multiple API calls to generate telemetry data.
	for i := 0; i < 3; i++ {
		_, err := checkoutApi.Events.RetrieveAllEventTypes()
		assert.NoError(t, err)
	}

	// Telemetry headers are included starting from the second request.
	expectedTelemetryHeaderCount := 2 // Telemetry is sent in the 2nd and 3rd requests.
	telemetryHeaderCount := mock.CountRequestsWithHeader("cko-sdk-telemetry")
	assert.Equal(t, expectedTelemetryHeaderCount, telemetryHeaderCount, "Expected exactly %d requests to contain the telemetry header", expectedTelemetryHeaderCount)

	assert.Nil(t, mock.ValidateTelemetryHeadersFormat("cko-sdk-telemetry"))
}

// TestShouldNotSendTelemetryWhenOptedOut verifies that telemetry headers are not sent when telemetry is disabled.
func TestShouldNotSendTelemetryWhenOptedOut(t *testing.T) {
	mock := newMockHTTPClient(0) // No delay.
	defer mock.Close()

	// Initialize the Checkout API client with telemetry disabled.
	checkoutApi, err := createCheckoutApi(mock, false)
	assert.NoError(t, err)
	assert.NotNil(t, checkoutApi)

	// Make multiple API calls.
	for i := 0; i < 3; i++ {
		_, err := checkoutApi.Events.RetrieveAllEventTypes()
		assert.NoError(t, err)
	}

	// Verify that no telemetry headers were sent.
	telemetryHeaderCount := mock.CountRequestsWithHeader("cko-sdk-telemetry")
	assert.Equal(t, 0, telemetryHeaderCount, "Expected no requests to contain the telemetry header")
}

// TestTelemetryQueueAndBottleneck simulates rapid requests to test telemetry queuing and bottleneck handling.
func TestTelemetryQueueAndBottleneck(t *testing.T) {
	// Introduce a delay to simulate server bottleneck.
	delay := 100 * time.Millisecond
	mock := newMockHTTPClient(delay)
	defer mock.Close()

	// Initialize the Checkout API client with telemetry enabled.
	checkoutApi, err := createCheckoutApi(mock, true)
	assert.NoError(t, err)
	assert.NotNil(t, checkoutApi)

	// Number of requests to simulate.
	numRequests := 10

	// Make multiple API calls rapidly to generate telemetry data.
	for i := 0; i < numRequests; i++ {
		_, err := checkoutApi.Events.RetrieveAllEventTypes()
		assert.NoError(t, err)
	}

	// Wait for all requests to be processed.
	time.Sleep(time.Duration(numRequests)*delay + 1*time.Second)

	// Count the number of requests that included the telemetry header.
	telemetryHeaderCount := mock.CountRequestsWithHeader("cko-sdk-telemetry")

	// Since telemetry starts being sent from the second request, we expect (numRequests - 1) telemetry headers.
	expectedTelemetryHeaderCount := numRequests - 1

	assert.Equal(t, expectedTelemetryHeaderCount, telemetryHeaderCount, "Expected %d requests to contain the telemetry header", expectedTelemetryHeaderCount)
}
