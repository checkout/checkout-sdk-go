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
	// ApplicantId is the applicant's unique identifier.
	// [Required]
	ApplicantId string `json:"applicant_id"`

	// UserJourneyId is your configuration ID.
	// [Required]
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
	HttpMetadata common.HttpMetadata

	// Id is the resource's unique identifier.
	// [Optional]
	Id string `json:"id,omitempty"`

	// CreatedOn is when the resource was created.
	// [Optional]
	// Format: date-time
	CreatedOn *time.Time `json:"created_on,omitempty"`

	// ModifiedOn is when the resource was last modified.
	// [Optional]
	// Format: date-time
	ModifiedOn *time.Time `json:"modified_on,omitempty"`

	// ResponseCodes is the codes explaining the verification outcome.
	// [Optional]
	ResponseCodes []identities.ResponseCode `json:"response_codes,omitempty"`
}

// Links holds the HAL links related to the resource.
//
// A union of the shapes the ID document verification endpoints return, so any given response
// leaves the irrelevant members nil: the verification returns self and applicant, the attempt list
// returns self, next and previous, and a single attempt returns self only.
type Links struct {
	// Self is the link to this resource.
	// [Optional]
	Self *common.Link `json:"self,omitempty"`

	// Applicant is the link to the applicant. Returned by the verification only.
	// [Optional]
	Applicant *common.Link `json:"applicant,omitempty"`

	// Next is the link to the next page. Returned by the attempt list only.
	// [Optional]
	Next *common.Link `json:"next,omitempty"`

	// Previous is the link to the previous page. Returned by the attempt list only.
	// [Optional]
	Previous *common.Link `json:"previous,omitempty"`
}

type IdDocumentVerificationResponse struct {
	idDocumentVerificationBase
	// UserJourneyId is your configuration ID.
	// [Optional]
	UserJourneyId string `json:"user_journey_id,omitempty"`

	// ApplicantId is the applicant's unique identifier.
	// [Required]
	ApplicantId string `json:"applicant_id,omitempty"`

	// Status is the verification's status.
	// [Required]
	Status identities.IdDocumentVerificationStatus `json:"status,omitempty"`
	// DeclaredData is the personal details provided by the applicant.
	// [Optional]
	DeclaredData *identities.DeclaredData `json:"declared_data,omitempty"`
	// Document is the details extracted from the verified document.
	// [Optional]
	Document *identities.DocumentDetails `json:"document,omitempty"`

	// Links holds the self and applicant HAL links.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

type IdDocumentVerificationAttemptResponse struct {
	idDocumentVerificationBase
	// Status is the attempt's status.
	// [Required]
	Status identities.IdDocumentVerificationAttemptStatus `json:"status,omitempty"`

	// Links holds the self HAL link.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

type IdDocumentVerificationAttemptsResponse struct {
	HttpMetadata common.HttpMetadata

	// TotalCount is the total number of attempts.
	// [Required]
	TotalCount int `json:"total_count,omitempty"`

	// Skip is the number of attempts skipped.
	// [Required]
	Skip int `json:"skip,omitempty"`

	// Limit is the maximum number of attempts returned.
	// [Required]
	Limit int `json:"limit,omitempty"`

	// Data is the list of attempts for the current page.
	// [Required]
	Data []IdDocumentVerificationAttemptResponse `json:"data,omitempty"`

	// Links holds the self, next and previous HAL links. Without it a caller can request a page
	// with Skip and Limit but cannot walk to the next one.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
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
