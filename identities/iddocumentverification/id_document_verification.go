package iddocumentverification

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

const (
	idDocumentVerificationsPath = "id-document-verifications"
	anonymizePath               = "anonymize"
	attemptsPath                = "attempts"
	reportPath                  = "pdf-report"
	assetsPath                  = "assets"
)

type CreateIdDocumentVerificationRequest struct {
	ApplicantId   string `json:"applicant_id"`
	UserJourneyId string `json:"user_journey_id"`
	// DeclaredData is the personal details provided by the applicant. The ID document
	// verification request takes the narrower IdvDeclaredData shape.
	// [Optional]
	DeclaredData *identities.DeclaredData `json:"declared_data,omitempty"`
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
	TotalCount   int                                     `json:"total_count,omitempty"`
	Skip         int                                     `json:"skip,omitempty"`
	Limit        int                                     `json:"limit,omitempty"`
	Data         []IdDocumentVerificationAttemptResponse `json:"data,omitempty"`
}

// IdDocumentVerificationReportResponse represents the response body for
// GET /id-document-verifications/{id}/pdf-report.
type IdDocumentVerificationReportResponse struct {
	HttpMetadata common.HttpMetadata

	// PdfReport is the pre-signed URL to the PDF report. Replaced signed_url in the 2026-09-02
	// spec: IdvPdf now declares pdf_report as its only property, and requires it.
	// [Required]
	// Format: uri
	PdfReport string `json:"pdf_report,omitempty"`
}

// IdDocumentVerificationAttemptAsset is a single asset uploaded for an ID document verification
// attempt.
type IdDocumentVerificationAttemptAsset struct {
	// Type is the type of asset.
	// [Required]
	// Enum: "document_front_image" "document_back_image"
	Type identities.IdDocumentVerificationAttemptAssetType `json:"type,omitempty"`

	// Links holds the asset_url link, the only link the schema declares, and it is required.
	// [Required]
	Links identities.AttemptAssetLinks `json:"_links,omitempty"`
}

// IdDocumentVerificationAttemptAssetsResponse represents the response body for
// GET /id-document-verifications/{id}/attempts/{attemptId}/assets.
type IdDocumentVerificationAttemptAssetsResponse struct {
	HttpMetadata common.HttpMetadata

	// TotalCount is the total number of assets.
	// [Required]
	TotalCount int `json:"total_count,omitempty"`

	// Skip is the number of assets skipped.
	// [Required]
	Skip int `json:"skip,omitempty"`

	// Limit is the maximum number of assets returned.
	// [Required]
	Limit int `json:"limit,omitempty"`

	// Data is the list of assets for the current page. May be empty: the schema allows minItems 0.
	// [Required]
	Data []IdDocumentVerificationAttemptAsset `json:"data,omitempty"`

	// Links holds the self, next and previous links.
	// [Required]
	Links map[string]common.Link `json:"_links,omitempty"`
}
