package client

import (
	stderrors "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

const responseMarker = "body_marker_4b2d"

type stringContentResponse struct {
	HttpMetadata common.HttpMetadata
	Content      string `json:"content"`
}

func respondWith(status int, contentType string, body string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contentType != "" {
			w.Header().Set("Content-Type", contentType)
		}
		w.Header().Set(CkoRequestId, "req_456")
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
}

func callChecked(t *testing.T, do func() error) (err error) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("request: unexpected %v", r)
		}
	}()
	return do()
}

func TestSuccessWithJsonNullBodyGivesZeroValueAndMetadata(t *testing.T) {
	server := respondWith(http.StatusOK, "application/json", "null")
	defer server.Close()

	client := newTestClient(server.URL)
	var resp common.IdResponse
	err := callChecked(t, func() error { return client.Get("/test", testAuth(), nil, &resp) })

	assert.NoError(t, err)
	assert.Equal(t, "", resp.Id)
	assert.Equal(t, http.StatusOK, resp.HttpMetadata.StatusCode)
	assert.Equal(t, "req_456", *resp.HttpMetadata.Headers.CKORequestID)
}

func TestSuccessWithHtmlBodyIsADecodingError(t *testing.T) {
	server := respondWith(http.StatusOK, "text/html; charset=utf-8", "<html>"+responseMarker+"</html>")
	defer server.Close()

	client := newTestClient(server.URL)
	var resp common.IdResponse
	err := callChecked(t, func() error { return client.Post("/test", testAuth(), nil, &resp, nil) })

	var decodingErr errors.CheckoutDecodingError
	if assert.True(t, stderrors.As(err, &decodingErr), "got %T", err) {
		assert.Equal(t, http.StatusOK, decodingErr.StatusCode)
		assert.NotContains(t, err.Error(), responseMarker)
	}
}

func TestSuccessWithPdfIntoStringContent(t *testing.T) {
	body := "%PDF-1.4\n%..."
	server := respondWith(http.StatusOK, "application/pdf", body)
	defer server.Close()

	client := newTestClient(server.URL)
	var resp stringContentResponse
	err := callChecked(t, func() error { return client.Post("/test", testAuth(), nil, &resp, nil) })

	assert.NoError(t, err)
	assert.Equal(t, body, resp.Content)
}

func TestSuccessWithCsvIntoJsonResponseIsADecodingError(t *testing.T) {
	server := respondWith(http.StatusOK, "text/csv", "id\n"+responseMarker)
	defer server.Close()

	client := newTestClient(server.URL)
	var resp common.IdResponse
	err := callChecked(t, func() error { return client.Get("/test", testAuth(), nil, &resp) })

	assert.IsType(t, errors.CheckoutDecodingError{}, err)
	assert.NotContains(t, err.Error(), responseMarker)
}

func TestClientErrorWithNullOrHtmlBodyIsHandled(t *testing.T) {
	for _, body := range []string{"null", "<html>" + responseMarker + "</html>", `{"error_codes":null}`, `[]`} {
		server := respondWith(http.StatusUnprocessableEntity, "application/json", body)
		client := newTestClient(server.URL)
		var resp common.IdResponse
		err := callChecked(t, func() error { return client.Get("/test", testAuth(), nil, &resp) })
		server.Close()

		var apiErr errors.CheckoutAPIError
		if assert.True(t, stderrors.As(err, &apiErr), "body %q: got %T", body, err) {
			assert.NotContains(t, err.Error(), responseMarker, "body %q", body)
		}
	}
}

func TestSuccessIntoNilMappingIsAnError(t *testing.T) {
	server := respondWith(http.StatusOK, "application/json", `{"id":"x"}`)
	defer server.Close()

	client := newTestClient(server.URL)
	err := callChecked(t, func() error { return client.Get("/test", testAuth(), nil, nil) })

	assert.IsType(t, errors.CheckoutDecodingError{}, err)
}

type noBodyTransport struct{}

func (noBodyTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: http.StatusOK, Status: "200 OK", Header: http.Header{}, Request: r}, nil
}

func TestSuccessWithoutBodyFromCustomTransport(t *testing.T) {
	client := newTestClient("http://127.0.0.1")
	client.HttpClient.Transport = noBodyTransport{}

	var resp common.IdResponse
	err := callChecked(t, func() error { return client.Get("/test", testAuth(), nil, &resp) })

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.HttpMetadata.StatusCode)
}
