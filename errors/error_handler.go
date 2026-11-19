package errors

import (
	"encoding/json"
	"net/http"

	"github.com/checkout/checkout-sdk-go/v3/internal/jsondepth"
)

// HandleError returns the CheckoutAPIError for a response with a status of 400 or above. For a
// status of 500 or above, Status holds the response body. Otherwise the body is decoded into Data
// with the request id; a body that cannot be decoded, or nests JSON deeper than the SDK decodes,
// gives StatusCode 400 and Status "Unparsable error".
func HandleError(statusCode int, status string, requestId string, body []byte) CheckoutAPIError {
	if statusCode >= http.StatusInternalServerError {
		return CheckoutAPIError{
			StatusCode: statusCode,
			Status:     string(body),
		}
	}

	var details ErrorDetails
	if len(body) != 0 {
		if !jsondepth.WithinLimit(body) || json.Unmarshal(body, &details) != nil {
			return CheckoutAPIError{
				StatusCode: http.StatusBadRequest,
				Status:     "Unparsable error",
			}
		}
	}

	details.RequestID = requestId

	return CheckoutAPIError{
		StatusCode: statusCode,
		Status:     status,
		Data:       &details,
	}
}
