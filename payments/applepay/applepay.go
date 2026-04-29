package applepay

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

const (
	ApplePayCertificatesPath    = "applepay/certificates"
	ApplePayEnrollmentsPath     = "applepay/enrollments"
	ApplePaySigningRequestsPath = "applepay/signing-requests"
)

type ProtocolVersion string

const (
	EcV1  ProtocolVersion = "ec_v1"
	RsaV1 ProtocolVersion = "rsa_v1"
)

type UploadCertificateRequest struct {
	Content string `json:"content"`
}

type UploadCertificateResponse struct {
	HttpMetadata  common.HttpMetadata
	Id            string     `json:"id"`
	PublicKeyHash string     `json:"public_key_hash"`
	ValidFrom     *time.Time `json:"valid_from"`
	ValidUntil    *time.Time `json:"valid_until"`
}

type EnrollDomainRequest struct {
	Domain string `json:"domain"`
}

type GenerateCertificateSigningRequest struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version,omitempty"`
}

type GenerateCertificateSigningRequestResponse struct {
	HttpMetadata common.HttpMetadata
	Content      string `json:"content"`
}
