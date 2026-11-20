package applepay

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

func signingRequestMetadata(contentType string, body string) *common.HttpMetadata {
	header := http.Header{}
	header.Set("Content-Type", contentType)
	return &common.HttpMetadata{StatusCode: http.StatusOK, ResponseBody: []byte(body), Headers: &common.Headers{Header: header}}
}

func TestSigningRequestResponseFromEveryContentType(t *testing.T) {
	csr := "-----BEGIN CERTIFICATE REQUEST-----\nMIIBSDCB8AIBADCBj...\n-----END CERTIFICATE REQUEST-----"
	jsonBody, _ := json.Marshal(map[string]string{"content": csr})

	cases := map[string]string{
		"application/json": string(jsonBody),
		"application/pdf":  csr,
		"text/csv":         csr,
	}
	for contentType, body := range cases {
		t.Run(contentType, func(t *testing.T) {
			var response GenerateCertificateSigningRequestResponse
			err := common.Unmarshal(signingRequestMetadata(contentType, body), &response)
			assert.NoError(t, err)
			assert.Equal(t, csr, response.Content)
			assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
		})
	}
}

func TestSigningRequestResponseFromJsonNull(t *testing.T) {
	var response GenerateCertificateSigningRequestResponse
	err := common.Unmarshal(signingRequestMetadata("application/json", "null"), &response)
	assert.NoError(t, err)
	assert.Equal(t, "", response.Content)
}
