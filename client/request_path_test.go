package client_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/apm/klarna"
	"github.com/checkout/checkout-sdk-go/v3/apm/sepa"
	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/customers"
	"github.com/checkout/checkout-sdk-go/v3/disputes"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

// recordingServer answers every request with an empty JSON object and records the request
// line as sent on the wire.
type recordingServer struct {
	server   *httptest.Server
	mu       sync.Mutex
	requests []string
}

func newRecordingServer(t *testing.T) *recordingServer {
	r := &recordingServer{}
	r.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.mu.Lock()
		r.requests = append(r.requests, req.Method+" "+req.RequestURI)
		r.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	t.Cleanup(r.server.Close)
	return r
}

func (r *recordingServer) sent() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	return append([]string(nil), r.requests...)
}

func (r *recordingServer) configuration() *configuration.Configuration {
	return &configuration.Configuration{
		Credentials: configuration.NewDefaultKeysSdkCredentials("sk_test", "pk_test"),
		Environment: configuration.Sandbox(),
	}
}

func (r *recordingServer) apiClient() *client.ApiClient {
	return &client.ApiClient{
		HttpClient: http.Client{},
		BaseUri:    r.server.URL,
		Log:        log.New(os.Stdout, "test: ", log.LstdFlags),
	}
}

func TestCallerValueWithReservedCharactersStaysInOnePathSegment(t *testing.T) {
	server := newRecordingServer(t)
	customersClient := customers.NewClient(server.configuration(), server.apiClient())
	disputesClient := disputes.NewClient(server.configuration(), server.apiClient())

	_, err := customersClient.Get("cus_x?z#w")
	assert.Nil(t, err)
	_, err = disputesClient.SubmitEvidence("dsp_x?z#w")
	assert.Nil(t, err)

	assert.Equal(t, []string{
		"GET /customers/cus_x%3Fz%23w",
		"POST /disputes/dsp_x%3Fz%23w/evidence",
	}, server.sent())
}

// The API decodes %2F and %5C back into separators, so such values never reach the transport.
func TestCallerValueWithSeparatorsIsRejectedBeforeAnyRequest(t *testing.T) {
	for _, id := range []string{"cus_x/y", "../x/y", `..\x`, "x\r\ny", "x\x00y"} {
		t.Run(id, func(t *testing.T) {
			server := newRecordingServer(t)
			customersClient := customers.NewClient(server.configuration(), server.apiClient())
			disputesClient := disputes.NewClient(server.configuration(), server.apiClient())

			_, err := customersClient.Get(id)
			assert.IsType(t, errors.CheckoutArgumentError(""), err)
			_, err = disputesClient.GetDisputeDetails(id)
			assert.IsType(t, errors.CheckoutArgumentError(""), err)
			assert.Empty(t, server.sent())
		})
	}
}

func TestInvalidPathSegmentIsRejectedBeforeAnyRequest(t *testing.T) {
	for _, id := range []string{"", ".", ".."} {
		t.Run("customer id "+id, func(t *testing.T) {
			server := newRecordingServer(t)
			customersClient := customers.NewClient(server.configuration(), server.apiClient())

			response, err := customersClient.Delete(id)

			assert.Nil(t, response)
			assert.IsType(t, errors.CheckoutArgumentError(""), err)
			assert.Empty(t, server.sent())
		})
	}
}

func TestBaseUrlAndConstantSegmentsAreNotEncoded(t *testing.T) {
	server := newRecordingServer(t)
	klarnaClient := klarna.NewClient(server.configuration(), server.apiClient())
	sepaClient := sepa.NewClient(server.configuration(), server.apiClient())

	_, err := klarnaClient.CapturePayment("pay_x?y", klarna.OrderCaptureRequest{})
	assert.Nil(t, err)
	_, err = sepaClient.GetMandate("man_x?y")
	assert.Nil(t, err)

	assert.Equal(t, []string{
		"POST /klarna-external/orders/pay_x%3Fy/captures",
		"GET /sepa/mandates/man_x%3Fy",
	}, server.sent())
}
