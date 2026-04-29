package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/payments/applepay"
)

// tests

func TestUploadPaymentProcessingCertificate(t *testing.T) {
	t.Skip("Requires a valid Apple Pay payment processing certificate")
	cases := []struct {
		name    string
		request applepay.UploadCertificateRequest
		checker func(*applepay.UploadCertificateResponse, error)
	}{
		{
			name:    "when certificate is valid then upload certificate",
			request: buildApplePayUploadCertificateRequest(),
			checker: func(response *applepay.UploadCertificateResponse, err error) {
				assert.Nil(t, err)
				assertUploadCertificateResponse(t, response)
			},
		},
	}

	client := DefaultApi().ApplePay

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UploadPaymentProcessingCertificate(tc.request))
		})
	}
}

func TestEnrollDomain(t *testing.T) {
	t.Skip("Requires OAuth credentials with vault:apme-enrollment scope and domain verification")
	cases := []struct {
		name    string
		request applepay.EnrollDomainRequest
		checker func(*common.MetadataResponse, error)
	}{
		{
			name:    "when domain is valid then enroll domain",
			request: buildApplePayEnrollDomainRequest(),
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := OAuthApi().ApplePay

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.EnrollDomain(tc.request))
		})
	}
}

func TestGenerateCertificateSigningRequest(t *testing.T) {
	cases := []struct {
		name    string
		request applepay.GenerateCertificateSigningRequest
		checker func(*applepay.GenerateCertificateSigningRequestResponse, error)
	}{
		{
			name:    "when request uses ec_v1 protocol then return signing request",
			request: applepay.GenerateCertificateSigningRequest{ProtocolVersion: applepay.EcV1},
			checker: func(response *applepay.GenerateCertificateSigningRequestResponse, err error) {
				assert.Nil(t, err)
				assertSigningRequestResponse(t, response)
			},
		},
		{
			name:    "when request uses rsa_v1 protocol then return signing request",
			request: applepay.GenerateCertificateSigningRequest{ProtocolVersion: applepay.RsaV1},
			checker: func(response *applepay.GenerateCertificateSigningRequestResponse, err error) {
				assert.Nil(t, err)
				assertSigningRequestResponse(t, response)
			},
		},
	}

	client := DefaultApi().ApplePay

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GenerateCertificateSigningRequest(tc.request))
		})
	}
}

// common methods

func buildApplePayUploadCertificateRequest() applepay.UploadCertificateRequest {
	return applepay.UploadCertificateRequest{
		Content: "MIIEfTCCBCOgAwIBAgIID/asezaWNycwCgYIKoZIzj0EAwIwgYAxNDAy...",
	}
}

func buildApplePayEnrollDomainRequest() applepay.EnrollDomainRequest {
	return applepay.EnrollDomainRequest{
		Domain: "checkout-test-domain.com",
	}
}

func assertUploadCertificateResponse(t *testing.T, response *applepay.UploadCertificateResponse) {
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.PublicKeyHash)
	assert.NotNil(t, response.ValidFrom)
	assert.NotNil(t, response.ValidUntil)
	assert.True(t, response.ValidUntil.After(*response.ValidFrom))
}

func assertSigningRequestResponse(t *testing.T, response *applepay.GenerateCertificateSigningRequestResponse) {
	assert.NotNil(t, response)
	assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
	assert.NotEmpty(t, response.Content)
	assert.Contains(t, response.Content, "BEGIN CERTIFICATE REQUEST")
	assert.Contains(t, response.Content, "END CERTIFICATE REQUEST")
}
