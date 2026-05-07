package faceauthentication

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/identities"
)

const (
	faceAuthenticationsPath = "face-authentications"
	anonymizePath           = "anonymize"
	attemptsPath            = "attempts"
)

type CreateFaceAuthenticationRequest struct {
	ApplicantId  string `json:"applicant_id"`
	UserJourneyId string `json:"user_journey_id"`
}

type CreateFaceAuthenticationAttemptRequest struct {
	RedirectUrl       string                      `json:"redirect_url"`
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

type FaceAuthenticationResponse struct {
	faceAuthenticationBase
	UserJourneyId string                              `json:"user_journey_id,omitempty"`
	ApplicantId   string                              `json:"applicant_id,omitempty"`
	Status        identities.FaceAuthenticationStatus `json:"status,omitempty"`
	RiskLabels    []string                            `json:"risk_labels,omitempty"`
	Face          *identities.FaceImage               `json:"face,omitempty"`
}

type FaceAuthenticationAttemptResponse struct {
	faceAuthenticationBase
	Status                      identities.FaceAuthenticationAttemptStatus `json:"status,omitempty"`
	RedirectUrl                 string                                     `json:"redirect_url,omitempty"`
	ClientInformation           *identities.ClientInformation              `json:"client_information,omitempty"`
	ApplicantSessionInformation *identities.ApplicantSessionInformation    `json:"applicant_session_information,omitempty"`
}

type FaceAuthenticationAttemptsResponse struct {
	HttpMetadata common.HttpMetadata
	TotalCount   int                                  `json:"total_count,omitempty"`
	Skip         int                                  `json:"skip,omitempty"`
	Limit        int                                  `json:"limit,omitempty"`
	Data         []FaceAuthenticationAttemptResponse  `json:"data,omitempty"`
}
