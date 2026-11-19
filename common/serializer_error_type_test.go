package common

import (
	stderrors "errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/errors"
)

func assertDecodingErrorCarriesNoBody(t *testing.T, err error) {
	t.Helper()
	if !assert.IsType(t, errors.CheckoutDecodingError{}, err) {
		return
	}
	assert.Nil(t, stderrors.Unwrap(err), "no underlying error")
	for _, text := range []string{err.Error(), fmt.Sprintf("%v", err), fmt.Sprintf("%+v", err), fmt.Sprintf("%#v", err)} {
		assert.NotContains(t, text, bodyMarker)
		assert.NotContains(t, text, "99999999999999999999")
	}
}

func TestUnmarshalDecodingErrorIsTyped(t *testing.T) {
	var response edgeNoContent
	err := unmarshalChecked(t, edgeMetadata("text/html; charset=utf-8", "<html>"+bodyMarker+"</html>"), &response)

	var decodingErr errors.CheckoutDecodingError
	if assert.True(t, stderrors.As(err, &decodingErr), "got %T", err) {
		assert.Equal(t, http.StatusOK, decodingErr.StatusCode)
		assert.Equal(t, "text/html", decodingErr.ContentType)
		assert.Equal(t, `response could not be decoded: status 200, content type "text/html"`, err.Error())
	}
	assertDecodingErrorCarriesNoBody(t, err)
}

func TestUnmarshalDecodingErrorKeepsNoPartOfTheBody(t *testing.T) {
	bodies := []string{
		`{"count":"` + bodyMarker + `"}`,
		`{"count":99999999999999999999123}`,
		`{"id":` + bodyMarker + `}`,
		`{"id":"x"} ` + bodyMarker,
		bodyMarker,
	}
	for _, body := range bodies {
		var response edgeNoContent
		assertDecodingErrorCarriesNoBody(t, unmarshalChecked(t, edgeMetadata("application/json", body), &response))
	}
}

func TestUnmarshalUnsupportedContentIsTyped(t *testing.T) {
	var response edgeIntContent
	err := unmarshalChecked(t, edgeMetadata("application/pdf", bodyMarker), &response)

	assertDecodingErrorCarriesNoBody(t, err)
	assert.Equal(t, `response could not be decoded: status 200, content type "application/pdf"`, err.Error())
}

func TestUnmarshalFailingDecoderGivesTypedError(t *testing.T) {
	var response edgeFailingDecoder
	assertDecodingErrorCarriesNoBody(t, unmarshalChecked(t, edgeMetadata("application/json", `{"id":"`+bodyMarker+`"}`), &response))

	var nilValue edgeNilValueDecoder
	assertDecodingErrorCarriesNoBody(t, unmarshalChecked(t, edgeMetadata("application/json", `{"id":"`+bodyMarker+`"}`), &nilValue))
}
