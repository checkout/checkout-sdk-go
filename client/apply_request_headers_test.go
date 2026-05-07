package client

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// helpers — header structs used across cases

type testHeaders struct {
	MyHeader   string `json:"x-my-header"`
	OtherValue string `json:"x-other"`
}

type testRequest struct {
	Headers testHeaders
}

type testRequestPtrHeaders struct {
	Headers *testHeaders
}

type noHeaders struct {
	Name string `json:"name"`
}

// TestApplyRequestHeaders_NilInterface verifies a nil interface input is a no-op.
func TestApplyRequestHeaders_NilInterface(t *testing.T) {
	headers := make(http.Header)
	applyRequestHeaders(nil, headers)
	assert.Empty(t, headers)
}

// TestApplyRequestHeaders_NilPointer verifies a typed nil pointer is a no-op.
func TestApplyRequestHeaders_NilPointer(t *testing.T) {
	headers := make(http.Header)
	var req *testRequest
	applyRequestHeaders(req, headers)
	assert.Empty(t, headers)
}

// TestApplyRequestHeaders_NonStruct verifies a non-struct value is a no-op.
func TestApplyRequestHeaders_NonStruct(t *testing.T) {
	headers := make(http.Header)
	applyRequestHeaders("not-a-struct", headers)
	assert.Empty(t, headers)
}

// TestApplyRequestHeaders_NoHeadersField verifies a struct without a Headers field is a no-op.
func TestApplyRequestHeaders_NoHeadersField(t *testing.T) {
	headers := make(http.Header)
	applyRequestHeaders(noHeaders{Name: "test"}, headers)
	assert.Empty(t, headers)
}

// TestApplyRequestHeaders_NilPtrHeaders verifies a nil *Headers pointer is a no-op.
func TestApplyRequestHeaders_NilPtrHeaders(t *testing.T) {
	headers := make(http.Header)
	applyRequestHeaders(&testRequestPtrHeaders{Headers: nil}, headers)
	assert.Empty(t, headers)
}

// TestApplyRequestHeaders_EmptyHeaderValues verifies empty string fields are not set.
func TestApplyRequestHeaders_EmptyHeaderValues(t *testing.T) {
	headers := make(http.Header)
	req := &testRequest{Headers: testHeaders{MyHeader: "", OtherValue: ""}}
	applyRequestHeaders(req, headers)
	assert.Empty(t, headers)
}

// TestApplyRequestHeaders_SetsPopulatedFields verifies populated fields are set correctly.
func TestApplyRequestHeaders_SetsPopulatedFields(t *testing.T) {
	headers := make(http.Header)
	req := &testRequest{Headers: testHeaders{MyHeader: "val-1", OtherValue: "val-2"}}
	applyRequestHeaders(req, headers)
	assert.Equal(t, "val-1", headers.Get("x-my-header"))
	assert.Equal(t, "val-2", headers.Get("x-other"))
}

// TestApplyRequestHeaders_PartialFields verifies only non-empty fields are set.
func TestApplyRequestHeaders_PartialFields(t *testing.T) {
	headers := make(http.Header)
	req := &testRequest{Headers: testHeaders{MyHeader: "present", OtherValue: ""}}
	applyRequestHeaders(req, headers)
	assert.Equal(t, "present", headers.Get("x-my-header"))
	assert.Empty(t, headers.Get("x-other"))
}

// TestApplyRequestHeaders_PtrToHeaders verifies a pointer-to-struct Headers field works.
func TestApplyRequestHeaders_PtrToHeaders(t *testing.T) {
	headers := make(http.Header)
	req := &testRequestPtrHeaders{Headers: &testHeaders{MyHeader: "ptr-val"}}
	applyRequestHeaders(req, headers)
	assert.Equal(t, "ptr-val", headers.Get("x-my-header"))
}

// TestApplyRequestHeaders_TagWithOmitempty verifies "name,omitempty" tags are parsed correctly.
func TestApplyRequestHeaders_TagWithOmitempty(t *testing.T) {
	type headersWithOmitempty struct {
		Value string `json:"x-value,omitempty"`
	}
	type req struct{ Headers headersWithOmitempty }

	headers := make(http.Header)
	applyRequestHeaders(&req{Headers: headersWithOmitempty{Value: "hello"}}, headers)
	assert.Equal(t, "hello", headers.Get("x-value"))
}

// TestApplyRequestHeaders_DashTag verifies fields tagged json:"-" are skipped.
func TestApplyRequestHeaders_DashTag(t *testing.T) {
	type headersWithDash struct {
		Ignored string `json:"-"`
		Kept    string `json:"x-kept"`
	}
	type req struct{ Headers headersWithDash }

	headers := make(http.Header)
	applyRequestHeaders(&req{Headers: headersWithDash{Ignored: "secret", Kept: "visible"}}, headers)
	assert.Empty(t, headers.Get("-"))
	assert.Equal(t, "visible", headers.Get("x-kept"))
}

// TestApplyRequestHeaders_NoJsonTag verifies fields with no json tag are skipped.
func TestApplyRequestHeaders_NoJsonTag(t *testing.T) {
	type headersNoTag struct {
		NoTag string
		Tagged string `json:"x-tagged"`
	}
	type req struct{ Headers headersNoTag }

	headers := make(http.Header)
	applyRequestHeaders(&req{Headers: headersNoTag{NoTag: "ignored", Tagged: "set"}}, headers)
	assert.Empty(t, headers.Get("NoTag"))
	assert.Equal(t, "set", headers.Get("x-tagged"))
}

// TestApplyRequestHeaders_ValuePassedByValue verifies the function works with a non-pointer struct.
func TestApplyRequestHeaders_ValuePassedByValue(t *testing.T) {
	headers := make(http.Header)
	req := testRequest{Headers: testHeaders{MyHeader: "by-value"}}
	applyRequestHeaders(req, headers)
	assert.Equal(t, "by-value", headers.Get("x-my-header"))
}
