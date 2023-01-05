package mocks

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/common"
)

var (
	HttpMetadataStatusOk = common.HttpMetadata{
		Status:     "200 OK",
		StatusCode: http.StatusOK,
	}

	HttpMetadataStatusCreated = common.HttpMetadata{
		Status:     "201 Created",
		StatusCode: http.StatusCreated,
	}

	HttpMetadataStatusAccepted = common.HttpMetadata{
		Status:     "202 Accepted",
		StatusCode: http.StatusAccepted,
	}

	HttpMetadataStatusNoContent = common.HttpMetadata{
		Status:     "204 No Content",
		StatusCode: http.StatusNoContent,
	}
)
