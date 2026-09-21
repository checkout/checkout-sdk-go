package identities

import "github.com/checkout/checkout-sdk-go/v3/common"

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

type AddressDocumentVerificationStatus string

const (
	AdvCreated                 AddressDocumentVerificationStatus = "created"
	AdvQualityChecksInProgress AddressDocumentVerificationStatus = "quality_checks_in_progress"
	AdvChecksInProgress        AddressDocumentVerificationStatus = "checks_in_progress"
	AdvApproved                AddressDocumentVerificationStatus = "approved"
	AdvDeclined                AddressDocumentVerificationStatus = "declined"
	AdvRetryRequired           AddressDocumentVerificationStatus = "retry_required"
	AdvInconclusive            AddressDocumentVerificationStatus = "inconclusive"
)

type AddressDocumentVerificationAttemptStatus string

const (
	AdvAttemptChecksInProgress        AddressDocumentVerificationAttemptStatus = "checks_in_progress"
	AdvAttemptChecksInconclusive      AddressDocumentVerificationAttemptStatus = "checks_inconclusive"
	AdvAttemptCompleted               AddressDocumentVerificationAttemptStatus = "completed"
	AdvAttemptQualityChecksAborted    AddressDocumentVerificationAttemptStatus = "quality_checks_aborted"
	AdvAttemptQualityChecksInProgress AddressDocumentVerificationAttemptStatus = "quality_checks_in_progress"
	AdvAttemptTerminated              AddressDocumentVerificationAttemptStatus = "terminated"
)

type IdentityVerificationStatus string

const (
	IdvApproved          IdentityVerificationStatus = "approved"
	IdvCaptureInProgress IdentityVerificationStatus = "capture_in_progress"
	IdvChecksInProgress  IdentityVerificationStatus = "checks_in_progress"
	IdvCreated           IdentityVerificationStatus = "created"
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
	AttemptTerminated         AttemptVerificationStatus = "terminated"
)

// DocumentType is the type of identity document.
type DocumentType string

const (
	DrivingLicence  DocumentType = "Driving licence"
	IdCard          DocumentType = "ID"
	Other           DocumentType = "Other"
	Passport        DocumentType = "Passport"
	ResidencePermit DocumentType = "Residence Permit"
	TravelDocument  DocumentType = "Travel Document"
	Visa            DocumentType = "Visa"
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

// RiskLabel is a code that provides more information about a risk associated with the
// verification.
type RiskLabel string

const (
	MultipleFacesDetected RiskLabel = "multiple_faces_detected"
	MccNotConfident       RiskLabel = "mcc_not_confident"
	RiskyDocumentFormat   RiskLabel = "risky_document_format"
)

// InitialDevice is the type of device the applicant used to start the attempt.
type InitialDevice string

const (
	Desktop InitialDevice = "desktop"
	Mobile  InitialDevice = "mobile"
)

// CertificationType is the certification type.
type CertificationType string

const (
	Diatf CertificationType = "diatf"
)

// Gpg45Profile is the GPG 45 identity profile the verification meets.
type Gpg45Profile string

const (
	Gpg45M1A Gpg45Profile = "M1A"
	Gpg45M1C Gpg45Profile = "M1C"
	Gpg45H1A Gpg45Profile = "H1A"
)

// LevelOfConfidence is the level of confidence in the verified identity.
type LevelOfConfidence string

const (
	LevelOfConfidenceMedium LevelOfConfidence = "medium"
	LevelOfConfidenceHigh   LevelOfConfidence = "high"
)

// DiatfCertificationData is the certification data returned for the diatf certification type.
type DiatfCertificationData struct {
	// Gpg45Profile is the GPG 45 identity profile the verification meets.
	// [Optional]
	// Enum: "M1A" "M1C" "H1A"
	Gpg45Profile Gpg45Profile `json:"gpg45_profile,omitempty"`

	// LevelOfConfidence is the level of confidence in the verified identity.
	// [Optional]
	// Enum: "medium" "high"
	LevelOfConfidence LevelOfConfidence `json:"level_of_confidence,omitempty"`

	// RightToWork is the outcome of the applicant's right to work check.
	// [Optional]
	// Example: GRANTED
	RightToWork string `json:"right_to_work,omitempty"`
}

// Certification is the details of a certification associated with the identity verification.
type Certification struct {
	// Type is the certification type.
	// [Optional]
	// Enum: "diatf"
	Type CertificationType `json:"type,omitempty"`

	// Data is the certification data. The properties returned depend on the certification type.
	// [Optional]
	Data *DiatfCertificationData `json:"data,omitempty"`
}

// PhoneNumber is the applicant's mobile phone number, if sharing the attempt URL via SMS.
type PhoneNumber struct {
	// CountryCode is the international phone country code. This is a dialling prefix, not an ISO
	// country code, so it is deliberately not common.Country.
	// [Required]
	// ^\+(\d+)$
	// Example: +33
	CountryCode string `json:"country_code,omitempty"`

	// Number is the applicant's mobile number, without the country code.
	// [Required]
	// ^\d{1,14}$
	// Example: 5555550102
	Number string `json:"number,omitempty"`
}

// IdvAddress is the applicant's address.
type IdvAddress struct {
	// AddressLine1 is the first line of the address.
	// [Optional]
	// max 250 characters
	// Example: 123 Main Street
	AddressLine1 string `json:"address_line1,omitempty"`

	// AddressLine2 is the second line of the address.
	// [Optional]
	// max 250 characters
	// Example: Apt 4B
	AddressLine2 string `json:"address_line2,omitempty"`

	// City is the city or town.
	// [Optional]
	// max 50 characters
	// Example: London
	City string `json:"city,omitempty"`

	// State is the state, county, or province.
	// [Optional]
	// max 50 characters
	// Example: Greater London
	State string `json:"state,omitempty"`

	// Zip is the postal or ZIP code.
	// [Optional]
	// max 50 characters
	// Example: SW1A 1AA
	Zip string `json:"zip,omitempty"`

	// Country is the two-letter ISO country code of the address.
	// [Optional]
	// Standard: ISO 3166-1 alpha-2 country code
	// max 2 characters
	// Example: GB
	Country common.Country `json:"country,omitempty"`
}

// DeclaredData is the personal details provided by the applicant.
//
// Maps IdvDeclaredData, the shape the address document and ID document verification requests
// accept. The identity verification requests take the larger IdentityDeclaredData instead.
type DeclaredData struct {
	// Name is the applicant's name.
	// [Required]
	// min 2 characters, max 255 characters
	// Example: Hannah Bret
	Name string `json:"name,omitempty"`

	// BirthDate is the applicant's birth date.
	// [Optional]
	// Format: yyyy-MM-dd
	// Example: 1994-10-15
	BirthDate string `json:"birth_date,omitempty"`
}

// IdentityDeclaredData is the personal details provided by the applicant for an identity
// verification.
//
// Maps IdvIdentityDeclaredData, a strict superset of IdvDeclaredData, so it embeds DeclaredData.
// The three extra fields are genuinely sendable, unlike the readOnly response fields the same
// request schemas also declare.
type IdentityDeclaredData struct {
	DeclaredData

	// PhoneNumber is the applicant's mobile phone number, if sharing the attempt URL via SMS.
	// [Optional]
	PhoneNumber *PhoneNumber `json:"phone_number,omitempty"`

	// Email is the applicant's email address. Explicitly nullable in the spec, so the API may
	// return null for it rather than omitting it.
	// [Optional]
	// Format: email
	// Nullable: true
	// Example: hannah.bret@example.com
	Email string `json:"email,omitempty"`

	// Address is the applicant's address.
	// [Optional]
	Address *IdvAddress `json:"address,omitempty"`
}

// ClientInformation is the applicant's details for a face authentication attempt.
//
// Maps FavClientInformation. Deliberately smaller than IdentityVerificationClientInformation: the
// face authentication attempt schema declares neither document field, so sending them here would
// be a request the API rejects.
type ClientInformation struct {
	// PreSelectedResidenceCountry is the applicant's residence country.
	// [Optional]
	// Standard: ISO 3166-1 alpha-2 country code
	// ^[A-Z]{2}
	// Example: FR
	PreSelectedResidenceCountry common.Country `json:"pre_selected_residence_country,omitempty"`

	// PreSelectedLanguage is the language you want to use for the user interface.
	// [Optional]
	// Format: IETF BCP 47 language tag
	// Example: en-US
	PreSelectedLanguage string `json:"pre_selected_language,omitempty"`
}

// IdentityVerificationClientInformation is the applicant's details for an identity verification
// attempt.
//
// Maps IdvClientInformation, a strict superset of FavClientInformation, so it embeds
// ClientInformation.
type IdentityVerificationClientInformation struct {
	ClientInformation

	// PreSelectedDocumentIssuingCountry is the country that issued the applicant's identity
	// document.
	// [Optional]
	// Standard: ISO 3166-1 alpha-2 country code
	// ^[A-Z]{2}
	// Example: FR
	PreSelectedDocumentIssuingCountry common.Country `json:"pre_selected_document_issuing_country,omitempty"`

	// PreSelectedDocumentType is the type of identity document the applicant uses for the attempt.
	// [Optional]
	// Enum: "Driving licence" "ID" "Other" "Passport" "Residence Permit" "Travel Document" "Visa"
	PreSelectedDocumentType DocumentType `json:"pre_selected_document_type,omitempty"`
}

type SelectedDocument struct {
	Country      common.Country `json:"country,omitempty"`
	DocumentType DocumentType   `json:"document_type,omitempty"`
}

// ApplicantSessionInformation is the details of the applicant's session during the attempt.
type ApplicantSessionInformation struct {
	// IpAddress is the IP address the applicant used.
	// [Optional]
	IpAddress string `json:"ip_address,omitempty"`

	// NumberOfSessions is the number of sessions the applicant started for the attempt.
	// [Optional]
	NumberOfSessions int `json:"number_of_sessions,omitempty"`

	// UserAgent is the user agent of the applicant's browser.
	// [Optional]
	UserAgent string `json:"user_agent,omitempty"`

	// InitialDevice is the type of device the applicant used to start the attempt.
	// [Optional]
	// Enum: "desktop" "mobile"
	InitialDevice InitialDevice `json:"initial_device,omitempty"`

	// SelectedDocuments is the documents the applicant selected. Identity verification only.
	// [Optional]
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
	DocumentIssuingCountry  common.Country `json:"document_issuing_country,omitempty"`
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

	// Address is the address extracted from the document. A flat string here, unlike the
	// structured IdvAddress used by the declared data.
	// [Optional]
	// max 1000 characters
	// Example: 123 Main Street, London, SW1A 1AA
	Address string `json:"address,omitempty"`

	// PermitObtainingDate is the date the permit was obtained.
	// [Optional]
	// Format: yyyy-MM-dd
	PermitObtainingDate string `json:"permit_obtaining_date,omitempty"`

	// PermitExpiryDate is the date the permit expires.
	// [Optional]
	// Format: yyyy-MM-dd
	PermitExpiryDate string `json:"permit_expiry_date,omitempty"`

	// PermitTypeDetailed is the detailed permit type.
	// [Optional]
	// max 255 characters
	PermitTypeDetailed string `json:"permit_type_detailed,omitempty"`

	// PermitTypeRemarks is the remarks recorded against the permit type.
	// [Optional]
	// max 255 characters
	PermitTypeRemarks string `json:"permit_type_remarks,omitempty"`
}

type FaceImage struct {
	ImageSignedUrl string `json:"image_signed_url,omitempty"`
}

// FaceAuthenticationAttemptAssetType is the type of asset captured during a face authentication attempt.
type FaceAuthenticationAttemptAssetType string

const (
	FaceImageFavAsset FaceAuthenticationAttemptAssetType = "face_image"
	FaceVideoFavAsset FaceAuthenticationAttemptAssetType = "face_video"
)

// IdentityVerificationAttemptAssetType is the type of asset captured during an identity verification attempt.
type IdentityVerificationAttemptAssetType string

const (
	FaceImageIdvAsset                    IdentityVerificationAttemptAssetType = "face_image"
	FaceVideoIdvAsset                    IdentityVerificationAttemptAssetType = "face_video"
	DocumentFrontImageIdvAsset           IdentityVerificationAttemptAssetType = "document_front_image"
	DocumentBackImageIdvAsset            IdentityVerificationAttemptAssetType = "document_back_image"
	DocumentFrontVideoIdvAsset           IdentityVerificationAttemptAssetType = "document_front_video"
	DocumentBackVideoIdvAsset            IdentityVerificationAttemptAssetType = "document_back_video"
	DocumentSignatureImageIdvAsset       IdentityVerificationAttemptAssetType = "document_signature_image"
	SecondaryDocumentFrontImageIdvAsset  IdentityVerificationAttemptAssetType = "secondary_document_front_image"
	SecondaryDocumentBackImageIdvAsset   IdentityVerificationAttemptAssetType = "secondary_document_back_image"
	SecondaryDocumentFrontVideoIdvAsset  IdentityVerificationAttemptAssetType = "secondary_document_front_video"
	SecondaryDocumentBackVideoIdvAsset   IdentityVerificationAttemptAssetType = "secondary_document_back_video"
	SecondaryDocumentSignatureImageAsset IdentityVerificationAttemptAssetType = "secondary_document_signature_image"
)

// AddressDocumentVerificationAttemptAssetType is the type of asset uploaded for an address
// document verification attempt.
type AddressDocumentVerificationAttemptAssetType string

const (
	DocumentAdvAsset AddressDocumentVerificationAttemptAssetType = "document"
)

// IdDocumentVerificationAttemptAssetType is the type of asset uploaded for an ID document
// verification attempt.
type IdDocumentVerificationAttemptAssetType string

const (
	DocumentFrontImageIddvAsset IdDocumentVerificationAttemptAssetType = "document_front_image"
	DocumentBackImageIddvAsset  IdDocumentVerificationAttemptAssetType = "document_back_image"
)

// AttemptAssetsQueryFilter holds the pagination query parameters for retrieving attempt assets.
//
// Note that omitempty drops a zero Skip, so skip=0 is never sent. That is harmless because 0 is
// the API default, but it means an explicit "start from the beginning" cannot be expressed.
type AttemptAssetsQueryFilter struct {
	// Skip is the number of assets to skip.
	// [Optional]
	// Default: 0
	Skip int `url:"skip,omitempty"`

	// Limit is the maximum number of assets to return.
	// [Optional]
	// Default: 10
	Limit int `url:"limit,omitempty"`
}

// AttemptsQueryFilter holds the pagination query parameters for listing attempts.
//
// Separate from AttemptAssetsQueryFilter because the spec documents the two sets of parameters
// against different resources: these count attempts, those count assets. The same omitempty
// caveat about a zero Skip applies.
type AttemptsQueryFilter struct {
	// Skip is the number of attempts to skip.
	// [Optional]
	// Default: 0
	Skip int `url:"skip,omitempty"`

	// Limit is the maximum number of attempts to return.
	// [Optional]
	// Default: 10
	Limit int `url:"limit,omitempty"`
}

// AttemptAssetLinks holds the links related to an attempt asset.
type AttemptAssetLinks struct {
	AssetUrl common.Link `json:"asset_url,omitempty"`
}
