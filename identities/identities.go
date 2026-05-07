package identities

import "github.com/checkout/checkout-sdk-go/v2/common"

type AmlScreeningStatus string

const (
	AmlCreated             AmlScreeningStatus = "created"
	AmlScreeningInProgress AmlScreeningStatus = "screening_in_progress"
	AmlApproved            AmlScreeningStatus = "approved"
	AmlDeclined            AmlScreeningStatus = "declined"
	AmlReviewRequired      AmlScreeningStatus = "review_required"
)

type FaceAuthenticationStatus string

const (
	FaceAuthApproved          FaceAuthenticationStatus = "approved"
	FaceAuthCaptureInProgress FaceAuthenticationStatus = "capture_in_progress"
	FaceAuthChecksInProgress  FaceAuthenticationStatus = "checks_in_progress"
	FaceAuthCreated           FaceAuthenticationStatus = "created"
	FaceAuthDeclined          FaceAuthenticationStatus = "declined"
	FaceAuthInconclusive      FaceAuthenticationStatus = "inconclusive"
	FaceAuthPending           FaceAuthenticationStatus = "pending"
	FaceAuthRefused           FaceAuthenticationStatus = "refused"
	FaceAuthRetryRequired     FaceAuthenticationStatus = "retry_required"
)

type FaceAuthenticationAttemptStatus string

const (
	FaceAuthAttemptCaptureAborted     FaceAuthenticationAttemptStatus = "capture_aborted"
	FaceAuthAttemptCaptureInProgress  FaceAuthenticationAttemptStatus = "capture_in_progress"
	FaceAuthAttemptChecksInconclusive FaceAuthenticationAttemptStatus = "checks_inconclusive"
	FaceAuthAttemptChecksInProgress   FaceAuthenticationAttemptStatus = "checks_in_progress"
	FaceAuthAttemptCompleted          FaceAuthenticationAttemptStatus = "completed"
	FaceAuthAttemptExpired            FaceAuthenticationAttemptStatus = "expired"
	FaceAuthAttemptPendingRedirection FaceAuthenticationAttemptStatus = "pending_redirection"
	FaceAuthAttemptCaptureRefused     FaceAuthenticationAttemptStatus = "capture_refused"
)

type IdDocumentVerificationStatus string

const (
	IddvCreated                 IdDocumentVerificationStatus = "created"
	IddvQualityChecksInProgress IdDocumentVerificationStatus = "quality_checks_in_progress"
	IddvChecksInProgress        IdDocumentVerificationStatus = "checks_in_progress"
	IddvApproved                IdDocumentVerificationStatus = "approved"
	IddvDeclined                IdDocumentVerificationStatus = "declined"
	IddvRetryRequired           IdDocumentVerificationStatus = "retry_required"
	IddvInconclusive            IdDocumentVerificationStatus = "inconclusive"
)

type IdDocumentVerificationAttemptStatus string

const (
	IddvAttemptChecksInProgress        IdDocumentVerificationAttemptStatus = "checks_in_progress"
	IddvAttemptChecksInconclusive      IdDocumentVerificationAttemptStatus = "checks_inconclusive"
	IddvAttemptCompleted               IdDocumentVerificationAttemptStatus = "completed"
	IddvAttemptQualityChecksAborted    IdDocumentVerificationAttemptStatus = "quality_checks_aborted"
	IddvAttemptQualityChecksInProgress IdDocumentVerificationAttemptStatus = "quality_checks_in_progress"
	IddvAttemptTerminated              IdDocumentVerificationAttemptStatus = "terminated"
)

type IdentityVerificationStatus string

const (
	IdvApproved          IdentityVerificationStatus = "approved"
	IdvCaptureInProgress IdentityVerificationStatus = "capture_in_progress"
	IdvChecksInProgress  IdentityVerificationStatus = "checks_in_progress"
	IdvDeclined          IdentityVerificationStatus = "declined"
	IdvInconclusive      IdentityVerificationStatus = "inconclusive"
	IdvPending           IdentityVerificationStatus = "pending"
	IdvRefused           IdentityVerificationStatus = "refused"
	IdvRetryRequired     IdentityVerificationStatus = "retry_required"
)

type AttemptVerificationStatus string

const (
	AttemptCaptureAborted     AttemptVerificationStatus = "capture_aborted"
	AttemptCaptureInProgress  AttemptVerificationStatus = "capture_in_progress"
	AttemptChecksInconclusive AttemptVerificationStatus = "checks_inconclusive"
	AttemptChecksInProgress   AttemptVerificationStatus = "checks_in_progress"
	AttemptCompleted          AttemptVerificationStatus = "completed"
	AttemptExpired            AttemptVerificationStatus = "expired"
	AttemptPendingRedirection AttemptVerificationStatus = "pending_redirection"
	AttemptCaptureRefused     AttemptVerificationStatus = "capture_refused"
)

type DocumentType string

const (
	DrivingLicence  DocumentType = "Driving licence"
	IdCard          DocumentType = "ID"
	Passport        DocumentType = "Passport"
	ResidencePermit DocumentType = "Residence Permit"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type ResponseCode struct {
	Code    int    `json:"code,omitempty"`
	Summary string `json:"summary,omitempty"`
}

type SearchParameters struct {
	ConfigurationIdentifier string `json:"configuration_identifier,omitempty"`
}

type DeclaredData struct {
	Name string `json:"name,omitempty"`
}

type ClientInformation struct {
	PreSelectedResidenceCountry common.Country `json:"pre_selected_residence_country,omitempty"`
	PreSelectedLanguage         string         `json:"pre_selected_language,omitempty"`
}

type SelectedDocument struct {
	Country      common.Country `json:"country,omitempty"`
	DocumentType DocumentType   `json:"document_type,omitempty"`
}

type ApplicantSessionInformation struct {
	IpAddress         string             `json:"ip_address,omitempty"`
	SelectedDocuments []SelectedDocument `json:"selected_documents,omitempty"`
}

type VerifiedIdentity struct {
	FullName        string         `json:"full_name,omitempty"`
	BirthDate       string         `json:"birth_date,omitempty"`
	FirstNames      string         `json:"first_names,omitempty"`
	LastName        string         `json:"last_name,omitempty"`
	LastNameAtBirth string         `json:"last_name_at_birth,omitempty"`
	BirthPlace      string         `json:"birth_place,omitempty"`
	Nationality     common.Country `json:"nationality,omitempty"`
	Gender          Gender         `json:"gender,omitempty"`
}

type DocumentDetails struct {
	DocumentType            DocumentType   `json:"document_type,omitempty"`
	DocumentIssuingCountry  string         `json:"document_issuing_country,omitempty"`
	FrontImageSignedUrl     string         `json:"front_image_signed_url,omitempty"`
	FullName                string         `json:"full_name,omitempty"`
	BirthDate               string         `json:"birth_date,omitempty"`
	FirstNames              string         `json:"first_names,omitempty"`
	LastName                string         `json:"last_name,omitempty"`
	LastNameAtBirth         string         `json:"last_name_at_birth,omitempty"`
	BirthPlace              string         `json:"birth_place,omitempty"`
	Nationality             common.Country `json:"nationality,omitempty"`
	Gender                  Gender         `json:"gender,omitempty"`
	PersonalNumber          string         `json:"personal_number,omitempty"`
	TaxIdentificationNumber string         `json:"tax_identification_number,omitempty"`
	DocumentNumber          string         `json:"document_number,omitempty"`
	DocumentExpiryDate      string         `json:"document_expiry_date,omitempty"`
	DocumentIssueDate       string         `json:"document_issue_date,omitempty"`
	DocumentIssuePlace      string         `json:"document_issue_place,omitempty"`
	DocumentMrz             string         `json:"document_mrz,omitempty"`
	BackImageSignedUrl      string         `json:"back_image_signed_url,omitempty"`
	SignatureImageSignedUrl string         `json:"signature_image_signed_url,omitempty"`
}

type FaceImage struct {
	ImageSignedUrl string `json:"image_signed_url,omitempty"`
}
