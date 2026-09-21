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
	HttpMetadata  common.HttpMetadata
	Id            string                    `json:"id,omitempty"`
	CreatedOn     *time.Time                `json:"created_on,omitempty"`
	ModifiedOn    *time.Time                `json:"modified_on,omitempty"`
	ResponseCodes []identities.ResponseCode `json:"response_codes,omitempty"`

	// DeclaredData is the personal details provided by the applicant.
	DeclaredData *identities.IdentityDeclaredData `json:"declared_data,omitempty"`
}

// identityVerificationCore holds fields shared by IdentityVerificationResponse
// and IdentityVerificationAndAttemptResponse.
type identityVerificationCore struct {
	identityVerificationBase
	UserJourneyId string                                `json:"user_journey_id,omitempty"`
	ApplicantId   string                                `json:"applicant_id,omitempty"`
	Status        identities.IdentityVerificationStatus `json:"status,omitempty"`

	// RiskLabels is one or more codes that provide more information about risks associated with
	// the verification.
	// [Optional]
	RiskLabels []identities.RiskLabel `json:"risk_labels,omitempty"`

	Documents []identities.DocumentDetails `json:"documents,omitempty"`

	// Face is the face image captured during the verification.
	//
	// The JSON tag is `face`, which is what the spec declares. It read `face_image` until the
	// 2026-09-02 pass, a tag the API never sends, so this field silently stayed nil on every
	// response. The face authentication response has always used `face` correctly.
	// [Optional]
	Face *identities.FaceImage `json:"face,omitempty"`

	VerifiedIdentity *identities.VerifiedIdentity `json:"verified_identity,omitempty"`

	// Certifications is the certifications associated with the identity verification.
	// [Optional]
	Certifications []identities.Certification `json:"certifications,omitempty"`

	// VerificationPolicyVersion is the version of the verification policy applied.
	// [Optional]
	VerificationPolicyVersion string `json:"verification_policy_version,omitempty"`
}

type IdentityVerificationResponse struct {
	identityVerificationCore
}

type IdentityVerificationAndAttemptResponse struct {
	identityVerificationCore
	RedirectUrl string `json:"redirect_url,omitempty"`
}

type IdentityVerificationAttemptResponse struct {
	identityVerificationBase
	Status      identities.AttemptVerificationStatus `json:"status,omitempty"`
	RedirectUrl string                               `json:"redirect_url,omitempty"`

	// PhoneNumber is the applicant's mobile phone number, if the attempt URL was shared via SMS.
	// [Optional]
	PhoneNumber *identities.PhoneNumber `json:"phone_number,omitempty"`

	// ClientInformation is the applicant's details. The identity verification attempt returns the
	// wider IdvClientInformation shape, unlike the face authentication attempt.
	// [Optional]
	ClientInformation *identities.IdentityVerificationClientInformation `json:"client_information,omitempty"`

	ApplicantSessionInformation *identities.ApplicantSessionInformation `json:"applicant_session_information,omitempty"`
}

type IdentityVerificationAttemptsResponse struct {
	HttpMetadata common.HttpMetadata
	TotalCount   int                                   `json:"total_count,omitempty"`
	Skip         int                                   `json:"skip,omitempty"`
	Limit        int                                   `json:"limit,omitempty"`
	Data         []IdentityVerificationAttemptResponse `json:"data,omitempty"`
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
