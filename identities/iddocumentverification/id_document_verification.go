package iddocumentverification

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/identities"
)

const (
	idDocumentVerificationsPath = "id-document-verifications"
	anonymizePath               = "anonymize"
	attemptsPath                = "attempts"
	reportPath                  = "pdf-report"
)

type CreateIdDocumentVerificationRequest struct {
	ApplicantId   string                 `json:"applicant_id"`
	UserJourneyId string                 `json:"user_journey_id"`
	DeclaredData  *identities.DeclaredData `json:"declared_data,omitempty"`
}

type CreateIdDocumentVerificationAttemptRequest struct {
	DocumentFront string `json:"document_front"`
	DocumentBack  string `json:"document_back,omitempty"`
}

// idDocumentVerificationBase holds fields common to all ID document verification response types.
type idDocumentVerificationBase struct {
	HttpMetadata  common.HttpMetadata
	Id            string                    `json:"id,omitempty"`
	CreatedOn     *time.Time                `json:"created_on,omitempty"`
	ModifiedOn    *time.Time                `json:"modified_on,omitempty"`
	ResponseCodes []identities.ResponseCode `json:"response_codes,omitempty"`
}

type IdDocumentVerificationResponse struct {
	idDocumentVerificationBase
	UserJourneyId string                                  `json:"user_journey_id,omitempty"`
	ApplicantId   string                                  `json:"applicant_id,omitempty"`
	Status        identities.IdDocumentVerificationStatus `json:"status,omitempty"`
	DeclaredData  *identities.DeclaredData                `json:"declared_data,omitempty"`
	Document      *identities.DocumentDetails             `json:"document,omitempty"`
}

type IdDocumentVerificationAttemptResponse struct {
	idDocumentVerificationBase
	Status identities.IdDocumentVerificationAttemptStatus `json:"status,omitempty"`
}

type IdDocumentVerificationAttemptsResponse struct {
	HttpMetadata common.HttpMetadata
	TotalCount   int                                      `json:"total_count,omitempty"`
	Skip         int                                      `json:"skip,omitempty"`
	Limit        int                                      `json:"limit,omitempty"`
	Data         []IdDocumentVerificationAttemptResponse  `json:"data,omitempty"`
}

type IdDocumentVerificationReportResponse struct {
	HttpMetadata common.HttpMetadata
	SignedUrl    string `json:"signed_url,omitempty"`
}
