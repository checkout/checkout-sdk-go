package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

// schemaVersionHeaders mirrors how the accounts client models the versioned Accept header:
// a Headers field whose json-tagged fields become HTTP headers, excluded from the body via json:"-".
type schemaVersionHeaders struct {
	Accept string `json:"Accept,omitempty"`
}

// TestGetWithContext_TransmitsHeadersWithoutBody verifies end-to-end (real ApiClient + real transport)
// that a GET carrying a per-request Headers source emits the Accept header and sends NO request body.
func TestGetWithContext_TransmitsHeadersWithoutBody(t *testing.T) {
	var gotMethod, gotAccept string
	var gotBodyLen int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotAccept = r.Header.Get("Accept")
		body, _ := ioutil.ReadAll(r.Body)
		gotBodyLen = len(body)
		jsonOK(w)
	}))
	defer server.Close()

	source := struct {
		Headers *schemaVersionHeaders `json:"-"`
	}{&schemaVersionHeaders{Accept: "application/json;schema_version=3.0"}}

	var resp common.IdResponse
	err := newTestClient(server.URL).GetWithContext(context.Background(), "/test", testAuth(), source, &resp)

	assert.Nil(t, err)
	assert.Equal(t, http.MethodGet, gotMethod)
	assert.Equal(t, "application/json;schema_version=3.0", gotAccept)
	assert.Equal(t, 0, gotBodyLen, "GET must not carry a request body")
}

// TestGetWithContext_NilRequestKeepsDefaultAcceptAndNoBody verifies the nil-request path is unchanged:
// the global default Accept is kept and no body is sent.
func TestGetWithContext_NilRequestKeepsDefaultAcceptAndNoBody(t *testing.T) {
	var gotAccept string
	var gotBodyLen int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAccept = r.Header.Get("Accept")
		body, _ := ioutil.ReadAll(r.Body)
		gotBodyLen = len(body)
		jsonOK(w)
	}))
	defer server.Close()

	var resp common.IdResponse
	err := newTestClient(server.URL).GetWithContext(context.Background(), "/test", testAuth(), nil, &resp)

	assert.Nil(t, err)
	assert.Equal(t, "application/json", gotAccept)
	assert.Equal(t, 0, gotBodyLen, "GET must not carry a request body")
}

// TestGetWithContext_HonorsSchemaVersionOverride confirms an alternate version transmits verbatim.
func TestGetWithContext_HonorsSchemaVersionOverride(t *testing.T) {
	var gotAccept string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAccept = r.Header.Get("Accept")
		jsonOK(w)
	}))
	defer server.Close()

	source := struct {
		Headers *schemaVersionHeaders `json:"-"`
	}{&schemaVersionHeaders{Accept: "application/json;schema_version=2.0"}}

	var resp common.IdResponse
	err := newTestClient(server.URL).GetWithContext(context.Background(), "/test", testAuth(), source, &resp)

	assert.Nil(t, err)
	assert.Equal(t, "application/json;schema_version=2.0", gotAccept)
}
