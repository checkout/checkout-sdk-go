package addressdocumentverification

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

const (
	addressDocumentVerificationsPath = "address-document-verifications"
	anonymizePath                    = "anonymize"
	attemptsPath                     = "attempts"
	reportPath                       = "pdf-report"
	assetsPath                       = "assets"
)

type CreateAddressDocumentVerificationRequest struct {
	ApplicantId   string `json:"applicant_id"`
	UserJourneyId string `json:"user_journey_id"`
	// DeclaredData is the personal details provided by the applicant. The address document
	// verification request takes the narrower IdvDeclaredData shape.
	// [Optional]
	DeclaredData *identities.DeclaredData `json:"declared_data,omitempty"`
}

type CreateAddressDocumentVerificationAttemptRequest struct {
	Document string `json:"document"`
}

// Address is the address extracted from the document.
type Address struct {
	AddressLine1 string `json:"address_line1,omitempty"`
	AddressLine2 string `json:"address_line2,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Zip          string `json:"zip,omitempty"`

	// Country is the two-letter ISO country code of the address.
	// [Optional]
	// Standard: ISO 3166-1 alpha-2 country code
	// max 2 characters
	// Example: GB
	Country common.Country `json:"country,omitempty"`
}

// AddressDocumentResult is the result of the address document check.
type AddressDocumentResult struct {
	DocumentType string   `json:"document_type,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	FullNames    []string `json:"full_names,omitempty"`
	IssueDate    string   `json:"issue_date,omitempty"`
	Address      *Address `json:"address,omitempty"`
}

// Links holds the HAL links related to the resource.
//
// A union of the three shapes the address document verification endpoints return, so any given
// response leaves the irrelevant members nil:
//   - the verification returns self and applicant
//   - the attempt list returns self, next and previous
//   - a single attempt returns self only
type Links struct {
	// Self is the link to this resource.
	// [Optional]
	Self *common.Link `json:"self,omitempty"`

	// Applicant is the link to the applicant. Returned by the verification only.
	// [Optional]
	Applicant *common.Link `json:"applicant,omitempty"`

	// Next is the link to the next page. Returned by the attempt list only, and absent on the
	// last page.
	// [Optional]
	Next *common.Link `json:"next,omitempty"`

	// Previous is the link to the previous page. Returned by the attempt list only, and absent
	// on the first page.
	// [Optional]
	Previous *common.Link `json:"previous,omitempty"`
}

// addressDocumentVerificationBase holds fields common to all response types.
type addressDocumentVerificationBase struct {
	HttpMetadata  common.HttpMetadata
	Id            string                    `json:"id,omitempty"`
	CreatedOn     *time.Time                `json:"created_on,omitempty"`
	ModifiedOn    *time.Time                `json:"modified_on,omitempty"`
	ResponseCodes []identities.ResponseCode `json:"response_codes,omitempty"`
	Links         *Links                    `json:"_links,omitempty"`
}

type AddressDocumentVerificationResponse struct {
	addressDocumentVerificationBase
	UserJourneyId string                                       `json:"user_journey_id,omitempty"`
	ApplicantId   string                                       `json:"applicant_id,omitempty"`
	Status        identities.AddressDocumentVerificationStatus `json:"status,omitempty"`

	// DeclaredData is the personal details provided by the applicant, echoed back by the API.
	// [Optional]
	DeclaredData *identities.DeclaredData `json:"declared_data,omitempty"`

	// RiskLabels is one or more codes that provide more information about risks associated with
	// the verification.
	// [Optional]
	RiskLabels []identities.RiskLabel `json:"risk_labels,omitempty"`

	AddressDocument *AddressDocumentResult `json:"address_document,omitempty"`
}

type AddressDocumentVerificationAttemptResponse struct {
	addressDocumentVerificationBase
	Status identities.AddressDocumentVerificationAttemptStatus `json:"status,omitempty"`
}

type AddressDocumentVerificationAttemptsResponse struct {
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
	Data  []AddressDocumentVerificationAttemptResponse `json:"data,omitempty"`
	Links *Links                                       `json:"_links,omitempty"`
}

// AddressDocumentVerificationReportResponse represents the response body for
// GET /address-document-verifications/{id}/pdf-report.
type AddressDocumentVerificationReportResponse struct {
	HttpMetadata common.HttpMetadata

	// PdfReport is the pre-signed URL to the PDF report. Replaced signed_url in the 2026-09-02
	// spec: IdvPdf now declares pdf_report as its only property, and requires it.
	// [Required]
	// Format: uri
	PdfReport string `json:"pdf_report,omitempty"`
}

// AddressDocumentVerificationAttemptAsset is a single asset uploaded for an address document
// verification attempt.
type AddressDocumentVerificationAttemptAsset struct {
	// Type is the type of asset.
	// [Required]
	// Enum: "document"
	Type identities.AddressDocumentVerificationAttemptAssetType `json:"type,omitempty"`

	// Links holds the asset_url link, the only link the schema declares, and it is required.
	// [Required]
	Links identities.AttemptAssetLinks `json:"_links,omitempty"`
}

// AddressDocumentVerificationAttemptAssetsResponse represents the response body for
// GET /address-document-verifications/{id}/attempts/{attemptId}/assets.
type AddressDocumentVerificationAttemptAssetsResponse struct {
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
	Data []AddressDocumentVerificationAttemptAsset `json:"data,omitempty"`

	// Links holds the self, next and previous links.
	// [Required]
	Links map[string]common.Link `json:"_links,omitempty"`
}
