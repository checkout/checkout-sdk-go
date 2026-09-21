package client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// cardUpdateHeaders mirrors cards.CardUpdateHeaders. Duplicated here rather than imported because
// the issuing package imports client, and importing back would be a cycle.
type cardUpdateHeaders struct {
	ReturnEncryptedCvv string `json:"return-encrypted-cvv,omitempty"`
	EncryptionKey      string `json:"Encryption-Key,omitempty"`
}

// TestPatchWithContext_TransmitsHeadersAndBody verifies end-to-end, through the real ApiClient and
// transport, that a PATCH carrying a Headers source emits both headers and still sends the body.
// PatchWithContext has no headers parameter, so the only route is the reflected Headers field.
// This is what proves part D needs no change to the HttpClient interface.
func TestPatchWithContext_TransmitsHeadersAndBody(t *testing.T) {
	var gotMethod, gotCvv, gotKey, gotBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		// Look up by the spec's lower-case spelling. http.Header.Get is case insensitive, which
		// is the whole reason the canonicalisation below is harmless.
		gotCvv = r.Header.Get("return-encrypted-cvv")
		gotKey = r.Header.Get("Encryption-Key")
		body, _ := ioutil.ReadAll(r.Body)
		gotBody = string(body)
		jsonOK(w)
	}))
	defer server.Close()

	source := struct {
		Reference string             `json:"reference,omitempty"`
		Headers   *cardUpdateHeaders `json:"-"`
	}{
		"X-123456-N11",
		&cardUpdateHeaders{ReturnEncryptedCvv: "true", EncryptionKey: "MIIBIjAN"},
	}

	var resp common.IdResponse
	err := newTestClient(server.URL).PatchWithContext(context.Background(), "/test", testAuth(), source, &resp)

	assert.Nil(t, err)
	assert.Equal(t, http.MethodPatch, gotMethod)
	assert.Equal(t, "true", gotCvv)
	assert.Equal(t, "MIIBIjAN", gotKey)
	assert.Contains(t, gotBody, `"reference":"X-123456-N11"`)
	assert.NotContains(t, gotBody, "return-encrypted-cvv", "headers must not leak into the body")
	assert.NotContains(t, gotBody, "Headers")
}

// TestPatchWithContext_HeaderNamesAreCanonicalisedOnTheWire documents a Go constraint rather than a
// preference. net/http applies textproto.CanonicalMIMEHeaderKey when writing a request, so
// return-encrypted-cvv always leaves as Return-Encrypted-Cvv no matter how the json tag is spelled,
// and direct map assignment does not avoid it either. RFC 7230 makes header names case insensitive,
// so the API reads it correctly. This test exists so nobody spends time trying to "fix" the casing.
func TestPatchWithContext_HeaderNamesAreCanonicalisedOnTheWire(t *testing.T) {
	var canonical, lowerCase string
	var sawExactLowerCaseKey bool
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		canonical = r.Header.Get("Return-Encrypted-Cvv")
		lowerCase = r.Header.Get("return-encrypted-cvv")
		_, sawExactLowerCaseKey = r.Header["return-encrypted-cvv"]
		jsonOK(w)
	}))
	defer server.Close()

	source := struct {
		Headers *cardUpdateHeaders `json:"-"`
	}{&cardUpdateHeaders{ReturnEncryptedCvv: "true"}}

	var resp common.IdResponse
	err := newTestClient(server.URL).PatchWithContext(context.Background(), "/test", testAuth(), source, &resp)

	assert.Nil(t, err)
	assert.Equal(t, "true", canonical, "Go emits the canonical form")
	assert.Equal(t, "true", lowerCase, "a case insensitive lookup still finds it")
	assert.False(t, sawExactLowerCaseKey, "the exact lower-case map key never survives the wire")
}

// An absent Headers field must leave the request untouched, which is the ordinary UpdateCard path.
func TestPatchWithContext_NoHeadersWhenTheSourceHasNone(t *testing.T) {
	var gotCvv, gotKey, gotBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotCvv = r.Header.Get("return-encrypted-cvv")
		gotKey = r.Header.Get("Encryption-Key")
		body, _ := ioutil.ReadAll(r.Body)
		gotBody = string(body)
		jsonOK(w)
	}))
	defer server.Close()

	source := struct {
		Reference string             `json:"reference,omitempty"`
		Headers   *cardUpdateHeaders `json:"-"`
	}{"X-123456-N11", nil}

	var resp common.IdResponse
	err := newTestClient(server.URL).PatchWithContext(context.Background(), "/test", testAuth(), source, &resp)

	assert.Nil(t, err)
	assert.Empty(t, gotCvv)
	assert.Empty(t, gotKey)
	assert.Contains(t, gotBody, `"reference":"X-123456-N11"`)
}

// The encryption key can be sent on its own; only the pairing rule is the API's business.
func TestPatchWithContext_SendsTheKeyWithoutTheFlag(t *testing.T) {
	var gotCvv, gotKey string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotCvv = r.Header.Get("return-encrypted-cvv")
		gotKey = r.Header.Get("Encryption-Key")
		jsonOK(w)
	}))
	defer server.Close()

	source := struct {
		Headers *cardUpdateHeaders `json:"-"`
	}{&cardUpdateHeaders{EncryptionKey: "MIIBIjAN"}}

	var resp common.IdResponse
	err := newTestClient(server.URL).PatchWithContext(context.Background(), "/test", testAuth(), source, &resp)

	assert.Nil(t, err)
	assert.Empty(t, gotCvv)
	assert.Equal(t, "MIIBIjAN", gotKey)
}
