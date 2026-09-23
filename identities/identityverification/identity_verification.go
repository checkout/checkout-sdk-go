package identityverification

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

const (
	createAndOpenPath         = "create-and-open-idv"
	identityVerificationsPath = "identity-verifications"
	anonymizePath             = "anonymize"
	attemptsPath              = "attempts"
	reportPath                = "pdf-report"
	assetsPath                = "assets"
)

// CreateIdentityVerificationRequest represents the request body for POST /identity-verifications.
type CreateIdentityVerificationRequest struct {
	// ApplicantId is the applicant's unique identifier.
	// [Required]
	ApplicantId string `json:"applicant_id"`

	// DeclaredData is the personal details provided by the applicant.
	// [Required]
	DeclaredData *identities.IdentityDeclaredData `json:"declared_data,omitempty"`

	// UserJourneyId is your configuration ID.
	// [Optional]
	UserJourneyId string `json:"user_journey_id,omitempty"`
}

// CreateIdentityVerificationAttemptRequest represents the request body for
// POST /identity-verifications/{id}/attempts.
type CreateIdentityVerificationAttemptRequest struct {
	// RedirectUrl is the URL to redirect the applicant to after the attempt.
	// [Required]
	// Format: uri
	RedirectUrl string `json:"redirect_url"`

	// PhoneNumber is the applicant's mobile phone number, if sharing the attempt URL via SMS.
	// [Optional]
	PhoneNumber *identities.PhoneNumber `json:"phone_number,omitempty"`

	// ClientInformation is the applicant's details.
	// [Optional]
	ClientInformation *identities.IdentityVerificationClientInformation `json:"client_information,omitempty"`
}

// CreateIdentityVerificationAndAttemptRequest represents the request body for
// POST /create-and-open-idv.
type CreateIdentityVerificationAndAttemptRequest struct {
	// DeclaredData is the personal details provided by the applicant.
	// [Required]
	DeclaredData *identities.IdentityDeclaredData `json:"declared_data,omitempty"`

	// RedirectUrl is the URL to redirect the applicant to after the attempt.
	// [Required]
	// Format: uri
	RedirectUrl string `json:"redirect_url"`

	// UserJourneyId is your configuration ID.
	// [Optional]
	UserJourneyId string `json:"user_journey_id,omitempty"`

	// ApplicantId is the applicant's unique identifier.
	// [Optional]
	ApplicantId string `json:"applicant_id,omitempty"`
}

// identityVerificationBase holds fields common to all identity verification response types.
type identityVerificationBase struct {
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

	// DeclaredData is the personal details provided by the applicant.
	//
	// Carried on the shared base, so IdentityVerificationAttemptResponse inherits it too. The
	// attempt schema (IdvAttemptResponse) does not declare declared_data at any level, so on an
	// attempt this field never populates. Retained rather than removed, per the plan's D4
	// decision, and scheduled for removal in a future major. On an attempt, read the declared
	// data from the parent verification instead.
	DeclaredData *identities.IdentityDeclaredData `json:"declared_data,omitempty"`
}

// identityVerificationCore holds fields shared by IdentityVerificationResponse
// and IdentityVerificationAndAttemptResponse.
type identityVerificationCore struct {
	identityVerificationBase
	// UserJourneyId is your configuration ID.
	// [Optional]
	UserJourneyId string `json:"user_journey_id,omitempty"`
	// ApplicantId is the applicant's unique identifier.
	// [Required]
	ApplicantId string `json:"applicant_id,omitempty"`
	// Status is the verification's status.
	// [Required]
	Status identities.IdentityVerificationStatus `json:"status,omitempty"`

	// RiskLabels is one or more codes that provide more information about risks associated with
	// the verification.
	// [Optional]
	RiskLabels []identities.RiskLabel `json:"risk_labels,omitempty"`

	// Documents is the details extracted from the verified documents.
	// [Optional]
	Documents []identities.DocumentDetails `json:"documents,omitempty"`

	// Face is the face image captured during the verification.
	//
	// The JSON tag is `face`, which is what the spec declares. It read `face_image` until the
	// 2026-09-02 pass, a tag the API never sends, so this field silently stayed nil on every
	// response. The face authentication response has always used `face` correctly.
	// [Optional]
	Face *identities.FaceImage `json:"face,omitempty"`

	// VerifiedIdentity is the identity the verification established.
	// [Optional]
	VerifiedIdentity *identities.VerifiedIdentity `json:"verified_identity,omitempty"`

	// Certifications is the certifications associated with the identity verification.
	// [Optional]
	Certifications []identities.Certification `json:"certifications,omitempty"`

	// VerificationPolicyVersion is the version of the verification policy applied.
	// [Optional]
	VerificationPolicyVersion string `json:"verification_policy_version,omitempty"`

	// Links holds the self and applicant HAL links.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

// Links holds the HAL links related to the resource.
//
// A union of the shapes these endpoints return, so any given response leaves the irrelevant
// members nil: the verification returns self and applicant, the attempt list returns self, next
// and previous, and a single attempt returns self and verification_url.
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

	// VerificationUrl is the URL the applicant uses to complete the attempt. Returned by a
	// single attempt only.
	// [Optional]
	VerificationUrl *common.Link `json:"verification_url,omitempty"`
}

type IdentityVerificationResponse struct {
	identityVerificationCore
}

type IdentityVerificationAndAttemptResponse struct {
	identityVerificationCore
	// RedirectUrl is the URL to send the applicant to so they can complete the attempt.
	// [Optional]
	// Format: uri
	RedirectUrl string `json:"redirect_url,omitempty"`
}

type IdentityVerificationAttemptResponse struct {
	identityVerificationBase
	// Status is the attempt's status.
	// [Required]
	Status identities.AttemptVerificationStatus `json:"status,omitempty"`
	// RedirectUrl is the URL the applicant is redirected to after the attempt.
	// [Optional]
	// Format: uri
	RedirectUrl string `json:"redirect_url,omitempty"`

	// PhoneNumber is the applicant's mobile phone number, if the attempt URL was shared via SMS.
	// [Optional]
	PhoneNumber *identities.PhoneNumber `json:"phone_number,omitempty"`

	// ClientInformation is the applicant's details. The identity verification attempt returns the
	// wider IdvClientInformation shape, unlike the face authentication attempt.
	// [Optional]
	ClientInformation *identities.IdentityVerificationClientInformation `json:"client_information,omitempty"`

	// ApplicantSessionInformation is the details of the applicant's session during the attempt.
	// [Optional]
	ApplicantSessionInformation *identities.ApplicantSessionInformation `json:"applicant_session_information,omitempty"`

	// Links holds the self and verification_url HAL links.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

type IdentityVerificationAttemptsResponse struct {
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
	Data []IdentityVerificationAttemptResponse `json:"data,omitempty"`

	// Links holds the self, next and previous HAL links. Without it a caller can request a page
	// with Skip and Limit but cannot walk to the next one.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

// IdentityVerificationReportResponse represents the response body for
// GET /identity-verifications/{id}/pdf-report.
type IdentityVerificationReportResponse struct {
	HttpMetadata common.HttpMetadata

	// PdfReport is the pre-signed URL to the PDF report. Replaced signed_url in the 2026-09-02
	// spec: IdvPdf now declares pdf_report as its only property, and requires it.
	// [Required]
	// Format: uri
	PdfReport string `json:"pdf_report,omitempty"`
}

type IdentityVerificationAttemptAsset struct {
	Type  identities.IdentityVerificationAttemptAssetType `json:"type,omitempty"`
	Links identities.AttemptAssetLinks                    `json:"_links,omitempty"`
}

type IdentityVerificationAttemptAssetsResponse struct {
	HttpMetadata common.HttpMetadata
	TotalCount   int                                `json:"total_count,omitempty"`
	Skip         int                                `json:"skip,omitempty"`
	Limit        int                                `json:"limit,omitempty"`
	Data         []IdentityVerificationAttemptAsset `json:"data,omitempty"`
	Links        map[string]common.Link             `json:"_links,omitempty"`
}
