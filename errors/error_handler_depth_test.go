package errors

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/internal/jsondepth"
)

func TestHandleErrorWithDeeplyNestedBody(t *testing.T) {
	nested := func(depth int) []byte {
		return []byte(`{"error_type":"request_invalid","_links":{"a":` + strings.Repeat("[", depth) + strings.Repeat("]", depth) + `}}`)
	}

	within := HandleError(http.StatusUnprocessableEntity, "422 Unprocessable Entity", "req_1", nested(jsondepth.MaxDepth-2))
	assert.Equal(t, http.StatusUnprocessableEntity, within.StatusCode)
	if assert.NotNil(t, within.Data) {
		assert.Equal(t, "request_invalid", within.Data.ErrorType)
	}

	for _, depth := range []int{jsondepth.MaxDepth, 1000000} {
		err := HandleError(http.StatusUnprocessableEntity, "422 Unprocessable Entity", "req_1", nested(depth))
		assert.Equal(t, CheckoutAPIError{StatusCode: http.StatusBadRequest, Status: "Unparsable error"}, err, "depth %d", depth)
	}
}
