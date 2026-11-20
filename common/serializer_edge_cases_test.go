package common

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/internal/jsondepth"
)

const bodyMarker = "body_marker_7f3a"

type edgeStringContent struct {
	HttpMetadata HttpMetadata
	Content      string `json:"content"`
}

type edgeBytesContent struct {
	HttpMetadata HttpMetadata
	Content      []byte `json:"content"`
}

type edgeIntContent struct {
	HttpMetadata HttpMetadata
	Content      int `json:"content"`
}

type edgeNoContent struct {
	HttpMetadata HttpMetadata
	Id           string `json:"id"`
	Count        int    `json:"count"`
}

type edgeWrongMetadata struct {
	HttpMetadata string
	Id           string `json:"id"`
}

type edgeInner struct {
	HttpMetadata HttpMetadata
	Content      string
}

type edgeEmbeddedPointer struct {
	*edgeInner
	Id string `json:"id"`
}

type edgeUnexported struct {
	httpMetadata HttpMetadata
	Id           string `json:"id"`
}

type edgeFailingDecoder struct {
	HttpMetadata HttpMetadata
}

func (h *edgeFailingDecoder) UnmarshalJSON(data []byte) error {
	var values []string
	_ = values[len(data)]
	return nil
}

type edgeNilValueDecoder struct {
	HttpMetadata HttpMetadata
	Id           string
}

func (h *edgeNilValueDecoder) UnmarshalJSON(data []byte) error {
	h.Id = "partial"
	panic(nil)
}

func edgeMetadata(contentType string, body string) *HttpMetadata {
	header := http.Header{}
	if contentType != "" {
		header.Set("Content-Type", contentType)
	}
	requestId := "req_123"
	return &HttpMetadata{
		Status:       "200 OK",
		StatusCode:   http.StatusOK,
		ResponseBody: []byte(body),
		Headers:      &Headers{Header: header, CKORequestID: &requestId},
	}
}

func unmarshalChecked(t *testing.T, metadata *HttpMetadata, mapping interface{}) (err error) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Unmarshal: unexpected %v", r)
		}
	}()
	return Unmarshal(metadata, mapping)
}

func assertDecodingErrorText(t *testing.T, err error) {
	t.Helper()
	if !assert.Error(t, err) {
		return
	}
	assert.NotContains(t, err.Error(), bodyMarker)
	assert.NotContains(t, err.Error(), "<html")
	assert.NotContains(t, err.Error(), "marker")
}

func TestUnmarshalJsonNullLeavesZeroValueWithMetadata(t *testing.T) {
	for _, body := range []string{"null", " null ", "\nnull\n"} {
		var response edgeNoContent
		err := unmarshalChecked(t, edgeMetadata("application/json", body), &response)

		assert.NoError(t, err, "body %q", body)
		assert.Equal(t, "", response.Id)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode, "body %q", body)
		assert.Equal(t, "req_123", *response.HttpMetadata.Headers.CKORequestID)
	}
}

func TestUnmarshalJsonNullIntoMetadataResponse(t *testing.T) {
	var response MetadataResponse
	err := unmarshalChecked(t, edgeMetadata("application/json", "null"), &response)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
}

func TestUnmarshalJsonNullIntoPointerToPointer(t *testing.T) {
	response := &edgeNoContent{Id: "kept"}
	err := unmarshalChecked(t, edgeMetadata("application/json", "null"), &response)

	assert.NoError(t, err)
	if assert.NotNil(t, response) {
		assert.Equal(t, "kept", response.Id)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
	}
}

func TestUnmarshalNonJsonSuccessBodyIsAnError(t *testing.T) {
	bodies := map[string]string{
		"html page":            "<html><body>" + bodyMarker + "</body></html>",
		"plain text":           bodyMarker,
		"truncated object":     `{"id":"` + bodyMarker,
		"truncated after key":  `{"id":`,
		"trailing garbage":     `{"id":"x"} ` + bodyMarker,
		"single quote":         `{'id':'` + bodyMarker + `'}`,
		"nul byte":             "{\"id\":\"\x00" + bodyMarker + "\"}",
		"invalid utf8 garbage": "\xff\xfe" + bodyMarker,
	}
	for name, body := range bodies {
		t.Run(name, func(t *testing.T) {
			var response edgeNoContent
			err := unmarshalChecked(t, edgeMetadata("application/json", body), &response)
			assertDecodingErrorText(t, err)
		})
	}
}

func TestUnmarshalWrongShapeIsAnError(t *testing.T) {
	bodies := map[string]string{
		"array":                   `["` + bodyMarker + `"]`,
		"string":                  `"` + bodyMarker + `"`,
		"number":                  `12345`,
		"boolean":                 `true`,
		"field of the wrong type": `{"id":"x","count":"` + bodyMarker + `"}`,
		"object where string":     `{"id":{"token":"` + bodyMarker + `"}}`,
	}
	for name, body := range bodies {
		t.Run(name, func(t *testing.T) {
			var response edgeNoContent
			err := unmarshalChecked(t, edgeMetadata("application/json", body), &response)
			assertDecodingErrorText(t, err)
		})
	}
}

func TestUnmarshalMissingContentTypeDecodesJson(t *testing.T) {
	var response edgeNoContent
	err := unmarshalChecked(t, edgeMetadata("", `{"id":"pay_1"}`), &response)

	assert.NoError(t, err)
	assert.Equal(t, "pay_1", response.Id)
}

func TestUnmarshalPdfIntoStringContent(t *testing.T) {
	body := "%PDF-1.4\n%\xe2\xe3\xcf\xd3"
	var response edgeStringContent
	err := unmarshalChecked(t, edgeMetadata("application/pdf", body), &response)

	assert.NoError(t, err)
	assert.Equal(t, body, response.Content)
	assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
}

func TestUnmarshalPdfIntoBytesAndDataContentKeepsBytes(t *testing.T) {
	body := "%PDF-1.4\n%..."

	var bytesResponse edgeBytesContent
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/pdf", body), &bytesResponse))
	assert.Equal(t, []byte(body), bytesResponse.Content)

	var dataResponse ContentResponse
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/pdf", body), &dataResponse))
	assert.Equal(t, []byte(body), dataResponse.Content)
}

func TestUnmarshalCsvIntoStringBytesAndDataContent(t *testing.T) {
	body := "id,amount\npay_1,100"
	for _, contentType := range []string{"text/csv", "text/csv; charset=utf-8", "TEXT/CSV"} {
		var stringResponse edgeStringContent
		assert.NoError(t, unmarshalChecked(t, edgeMetadata(contentType, body), &stringResponse), contentType)
		assert.Equal(t, body, stringResponse.Content, contentType)

		var bytesResponse edgeBytesContent
		assert.NoError(t, unmarshalChecked(t, edgeMetadata(contentType, body), &bytesResponse), contentType)
		assert.Equal(t, []byte(body), bytesResponse.Content, contentType)

		var dataResponse ContentResponse
		assert.NoError(t, unmarshalChecked(t, edgeMetadata(contentType, body), &dataResponse), contentType)
		assert.Equal(t, body, dataResponse.Content, contentType)
	}
}

func TestUnmarshalContentIntoAnUnsupportedFieldIsAnError(t *testing.T) {
	for _, contentType := range []string{"text/csv", "application/pdf"} {
		var intResponse edgeIntContent
		assertDecodingErrorText(t, unmarshalChecked(t, edgeMetadata(contentType, bodyMarker), &intResponse))

		var noContent edgeNoContent
		assertDecodingErrorText(t, unmarshalChecked(t, edgeMetadata(contentType, bodyMarker), &noContent))
	}
}

func TestUnmarshalInvalidMappingIsAnError(t *testing.T) {
	var typedNil *edgeNoContent
	mappings := map[string]interface{}{
		"nil":           nil,
		"typed nil":     typedNil,
		"not a pointer": edgeNoContent{},
	}
	for name, mapping := range mappings {
		for _, contentType := range []string{"application/json", "text/csv", "application/pdf"} {
			t.Run(name+" "+contentType, func(t *testing.T) {
				assertDecodingErrorText(t, unmarshalChecked(t, edgeMetadata(contentType, `{"id":"`+bodyMarker+`"}`), mapping))
			})
		}
		t.Run(name+" empty body", func(t *testing.T) {
			assert.Error(t, unmarshalChecked(t, edgeMetadata("application/json", ""), mapping))
		})
	}
}

func TestUnmarshalIntoNonStructMappings(t *testing.T) {
	var asMap map[string]interface{}
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", `{"id":"x"}`), &asMap))
	assert.Equal(t, "x", asMap["id"])

	var asSlice []string
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", `["a"]`), &asSlice))

	var asString string
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", `null`), &asString))
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", ``), &asString))
}

func TestUnmarshalMetadataFieldOfAnotherTypeIsAnError(t *testing.T) {
	var response edgeWrongMetadata
	assertDecodingErrorText(t, unmarshalChecked(t, edgeMetadata("application/json", `{"id":"x"}`), &response))
}

func TestUnmarshalEmbeddedNilPointerIsHandled(t *testing.T) {
	var response edgeEmbeddedPointer
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", `{"id":"x"}`), &response))
	assert.Equal(t, "x", response.Id)

	var csvResponse edgeEmbeddedPointer
	assert.Error(t, unmarshalChecked(t, edgeMetadata("text/csv", "a,b"), &csvResponse))
}

func TestUnmarshalUnexportedMetadataIsIgnored(t *testing.T) {
	var response edgeUnexported
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", `{"id":"x"}`), &response))
	assert.Equal(t, "x", response.Id)
}

func TestUnmarshalFailingDecoderIsAnError(t *testing.T) {
	var response edgeFailingDecoder
	assertDecodingErrorText(t, unmarshalChecked(t, edgeMetadata("application/json", `{"id":"`+bodyMarker+`"}`), &response))
}

func TestUnmarshalNilMetadataOrHeadersIsHandled(t *testing.T) {
	var response edgeNoContent
	assert.Error(t, unmarshalChecked(t, nil, &response))

	metadata := edgeMetadata("", `{"id":"x"}`)
	metadata.Headers = nil
	assert.NoError(t, unmarshalChecked(t, metadata, &response))
	assert.Equal(t, "x", response.Id)

	metadata = edgeMetadata("", `{"id":"y"}`)
	metadata.Headers.Header = nil
	assert.NoError(t, unmarshalChecked(t, metadata, &response))
	assert.Equal(t, "y", response.Id)
}

func TestUnmarshalErrorNamesTheMediaTypeOnly(t *testing.T) {
	contentTypes := []string{
		"text/html; token=" + bodyMarker,
		"text/html‮" + bodyMarker,
		"text/" + strings.Repeat("x", 5000) + bodyMarker,
	}
	for _, contentType := range contentTypes {
		var response edgeNoContent
		err := unmarshalChecked(t, edgeMetadata(contentType, "<html>"), &response)
		assertDecodingErrorText(t, err)
		if err != nil {
			assert.NotContains(t, err.Error(), "‮")
			assert.True(t, len(err.Error()) < 300, "message length %d", len(err.Error()))
		}
	}
}

func nestedArrays(depth int) string {
	return `{"metadata":{"k":` + strings.Repeat("[", depth) + strings.Repeat("]", depth) + `}}`
}

type edgeMetadataMap struct {
	HttpMetadata HttpMetadata
	Metadata     map[string]interface{} `json:"metadata"`
}

func TestUnmarshalNestingDepth(t *testing.T) {
	var atLimit edgeMetadataMap
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", nestedArrays(jsondepth.MaxDepth-2)), &atLimit))

	for _, depth := range []int{jsondepth.MaxDepth, 1000000} {
		var response edgeMetadataMap
		assertDecodingErrorText(t, unmarshalChecked(t, edgeMetadata("application/json", nestedArrays(depth)), &response))
	}

	var inString edgeNoContent
	body := `{"id":"` + strings.Repeat("[{", 20000) + `\"` + strings.Repeat("[", 20000) + `"}`
	assert.NoError(t, unmarshalChecked(t, edgeMetadata("application/json", body), &inString))
}
