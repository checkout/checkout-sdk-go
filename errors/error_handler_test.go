package errors

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	cases := []struct {
		name            string
		inputStatusCode int
		inputStatus     string
		inputRequestId  string
		inputBody       []byte
		expectedError   CheckoutAPIError
	}{
		{
			name:            "when body is nil then return CheckoutAPIError with empty Details",
			inputStatusCode: http.StatusNotFound,
			inputStatus:     "404 Not Found",
			inputRequestId:  "12345",
			expectedError: CheckoutAPIError{
				StatusCode: http.StatusNotFound,
				Status:     "404 Not Found",
				Data: &ErrorDetails{
					RequestID:  "12345",
					ErrorType:  "",
					ErrorCodes: nil,
				},
			},
		},
		{
			name:            "when body is not nil then return CheckoutAPIError with Details",
			inputStatusCode: http.StatusUnprocessableEntity,
			inputStatus:     "422 Unprocessable Entity",
			inputRequestId:  "12345",
			inputBody:       getErrorBody(),
			expectedError: CheckoutAPIError{
				StatusCode: http.StatusUnprocessableEntity,
				Status:     "422 Unprocessable Entity",
				Data: &ErrorDetails{
					RequestID:  "12345",
					ErrorType:  "request_invalid",
					ErrorCodes: []string{"invalid"},
				},
			},
		},
		{
			name:            "when error body is invalid then return unparsable error",
			inputStatusCode: http.StatusNotFound,
			inputStatus:     "404 Not Found",
			inputRequestId:  "12345",
			inputBody:       []byte("unparsable_body"),
			expectedError: CheckoutAPIError{
				StatusCode: http.StatusBadRequest,
				Status:     "Unparsable error",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedError, HandleError(tc.inputStatusCode, tc.inputStatus, tc.inputRequestId, tc.inputBody))
		})
	}
}

func getErrorBody() []byte {
	errorDetails := ErrorDetails{
		RequestID:  "12345",
		ErrorType:  "request_invalid",
		ErrorCodes: []string{"invalid"},
	}

	body, _ := json.Marshal(errorDetails)

	return body
}
