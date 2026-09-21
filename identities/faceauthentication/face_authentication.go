package faceauthentication

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

const (
	faceAuthenticationsPath = "face-authentications"
	anonymizePath           = "anonymize"
	attemptsPath            = "attempts"
	assetsPath              = "assets"
)

type CreateFaceAuthenticationRequest struct {
	ApplicantId   string `json:"applicant_id"`
	UserJourneyId string `json:"user_journey_id"`
}

// CreateFaceAuthenticationAttemptRequest represents the request body for
// POST /face-authentications/{id}/attempts.
type CreateFaceAuthenticationAttemptRequest struct {
	// RedirectUrl is the URL to redirect the applicant to after the attempt.
	// [Required]
	// Format: uri
	RedirectUrl string `json:"redirect_url"`

	// PhoneNumber is the applicant's mobile phone number, if sharing the attempt URL via SMS.
	// [Optional]
	PhoneNumber *identities.PhoneNumber `json:"phone_number,omitempty"`

	// ClientInformation is the applicant's details. The face authentication attempt takes the
	// narrower FavClientInformation shape, which declares neither document field.
	// [Optional]
	ClientInformation *identities.ClientInformation `json:"client_information,omitempty"`
}

// faceAuthenticationBase holds fields common to all face authentication response types.
type faceAuthenticationBase struct {
	HttpMetadata  common.HttpMetadata
	Id            string                    `json:"id,omitempty"`
	CreatedOn     *time.Time                `json:"created_on,omitempty"`
	ModifiedOn    *time.Time                `json:"modified_on,omitempty"`
	ResponseCodes []identities.ResponseCode `json:"response_codes,omitempty"`
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

type FaceAuthenticationResponse struct {
	faceAuthenticationBase
	UserJourneyId string                              `json:"user_journey_id,omitempty"`
	ApplicantId   string                              `json:"applicant_id,omitempty"`
	Status        identities.FaceAuthenticationStatus `json:"status,omitempty"`

	// RiskLabels is one or more codes that provide more information about risks associated with
	// the verification.
	// [Optional]
	RiskLabels []identities.RiskLabel `json:"risk_labels,omitempty"`

	Face *identities.FaceImage `json:"face,omitempty"`

	// Links holds the self and applicant HAL links.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

type FaceAuthenticationAttemptResponse struct {
	faceAuthenticationBase
	Status      identities.FaceAuthenticationAttemptStatus `json:"status,omitempty"`
	RedirectUrl string                                     `json:"redirect_url,omitempty"`

	// PhoneNumber is the applicant's mobile phone number, if the attempt URL was shared via SMS.
	// [Optional]
	PhoneNumber *identities.PhoneNumber `json:"phone_number,omitempty"`

	ClientInformation           *identities.ClientInformation           `json:"client_information,omitempty"`
	ApplicantSessionInformation *identities.ApplicantSessionInformation `json:"applicant_session_information,omitempty"`

	// Links holds the self and verification_url HAL links.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

type FaceAuthenticationAttemptsResponse struct {
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
	Data []FaceAuthenticationAttemptResponse `json:"data,omitempty"`

	// Links holds the self, next and previous HAL links. Without it a caller can request a page
	// with Skip and Limit but cannot walk to the next one.
	// [Optional]
	Links *Links `json:"_links,omitempty"`
}

type FaceAuthenticationAttemptAsset struct {
	Type  identities.FaceAuthenticationAttemptAssetType `json:"type,omitempty"`
	Links identities.AttemptAssetLinks                  `json:"_links,omitempty"`
}

type FaceAuthenticationAttemptAssetsResponse struct {
	HttpMetadata common.HttpMetadata
	TotalCount   int                              `json:"total_count,omitempty"`
	Skip         int                              `json:"skip,omitempty"`
	Limit        int                              `json:"limit,omitempty"`
	Data         []FaceAuthenticationAttemptAsset `json:"data,omitempty"`
	Links        map[string]common.Link           `json:"_links,omitempty"`
}
