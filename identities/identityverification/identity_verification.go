package identityverification

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/identities"
)

const (
	createAndOpenPath         = "create-and-open-idv"
	identityVerificationsPath = "identity-verifications"
	anonymizePath             = "anonymize"
	attemptsPath              = "attempts"
	reportPath                = "pdf-report"
)

type CreateIdentityVerificationRequest struct {
	ApplicantId   string                   `json:"applicant_id"`
	DeclaredData  *identities.DeclaredData `json:"declared_data,omitempty"`
	UserJourneyId string                   `json:"user_journey_id,omitempty"`
}

type CreateIdentityVerificationAttemptRequest struct {
	RedirectUrl       string                        `json:"redirect_url"`
	ClientInformation *identities.ClientInformation `json:"client_information,omitempty"`
}

type CreateIdentityVerificationAndAttemptRequest struct {
	DeclaredData  *identities.DeclaredData `json:"declared_data,omitempty"`
	RedirectUrl   string                   `json:"redirect_url"`
	UserJourneyId string                   `json:"user_journey_id,omitempty"`
	ApplicantId   string                   `json:"applicant_id,omitempty"`
}

// identityVerificationBase holds fields common to all identity verification response types.
type identityVerificationBase struct {
	HttpMetadata  common.HttpMetadata
	Id            string                    `json:"id,omitempty"`
	CreatedOn     *time.Time                `json:"created_on,omitempty"`
	ModifiedOn    *time.Time                `json:"modified_on,omitempty"`
	ResponseCodes []identities.ResponseCode `json:"response_codes,omitempty"`
	DeclaredData  *identities.DeclaredData  `json:"declared_data,omitempty"`
}

// identityVerificationCore holds fields shared by IdentityVerificationResponse
// and IdentityVerificationAndAttemptResponse.
type identityVerificationCore struct {
	identityVerificationBase
	UserJourneyId    string                                `json:"user_journey_id,omitempty"`
	ApplicantId      string                                `json:"applicant_id,omitempty"`
	Status           identities.IdentityVerificationStatus `json:"status,omitempty"`
	RiskLabels       []string                              `json:"risk_labels,omitempty"`
	Documents        []identities.DocumentDetails          `json:"documents,omitempty"`
	FaceImage        *identities.FaceImage                 `json:"face_image,omitempty"`
	VerifiedIdentity *identities.VerifiedIdentity          `json:"verified_identity,omitempty"`
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
	Status                      identities.AttemptVerificationStatus    `json:"status,omitempty"`
	RedirectUrl                 string                                  `json:"redirect_url,omitempty"`
	ClientInformation           *identities.ClientInformation           `json:"client_information,omitempty"`
	ApplicantSessionInformation *identities.ApplicantSessionInformation `json:"applicant_session_information,omitempty"`
}

type IdentityVerificationAttemptsResponse struct {
	HttpMetadata common.HttpMetadata
	TotalCount   int                                    `json:"total_count,omitempty"`
	Skip         int                                    `json:"skip,omitempty"`
	Limit        int                                    `json:"limit,omitempty"`
	Data         []IdentityVerificationAttemptResponse  `json:"data,omitempty"`
}

type IdentityVerificationReportResponse struct {
	HttpMetadata common.HttpMetadata
	SignedUrl    string `json:"signed_url,omitempty"`
}
