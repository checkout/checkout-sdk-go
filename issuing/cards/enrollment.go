package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type (
	SecurityPair struct {
		Question string `json:"question,omitempty"`
		Answer   string `json:"answer,omitempty"`
	}

	ThreeDSEnrollmentRequest struct {
		Locale      string        `json:"locale,omitempty"`
		PhoneNumber *common.Phone `json:"phone_number,omitempty"`
	}

	passwordThreeDSEnrollmentRequest struct {
		Password string `json:"password,omitempty"`
		ThreeDSEnrollmentRequest
	}

	securityQuestionThreeDSEnrollmentRequest struct {
		SecurityPair SecurityPair `json:"security_pair"`
		ThreeDSEnrollmentRequest
	}
)

func NewPasswordThreeDSEnrollmentRequest() *passwordThreeDSEnrollmentRequest {
	return &passwordThreeDSEnrollmentRequest{}
}

func NewSecurityQuestionThreeDSEnrollmentRequest() *securityQuestionThreeDSEnrollmentRequest {
	return &securityQuestionThreeDSEnrollmentRequest{}
}

type (
	ThreeDSUpdateRequest struct {
		SecurityPair SecurityPair  `json:"security_pair,omitempty"`
		Password     string        `json:"password,omitempty"`
		Locale       string        `json:"locale,omitempty"`
		PhoneNumber  *common.Phone `json:"phone_number,omitempty"`
	}

	ThreeDSUpdateResponse struct {
		HttpMetadata     common.HttpMetadata
		LastModifiedDate *time.Time             `json:"last_modified_date,omitempty"`
		Links            map[string]common.Link `json:"_links,omitempty"`
	}
)

type (
	ThreeDSEnrollmentResponse struct {
		HttpMetadata common.HttpMetadata
		CreatedDate  *time.Time             `json:"created_date,omitempty"`
		Links        map[string]common.Link `json:"_links,omitempty"`
	}

	ThreeDSEnrollmentDetailsResponse struct {
		HttpMetadata     common.HttpMetadata
		Locale           string                 `json:"locale,omitempty"`
		PhoneNumber      *common.Phone          `json:"phone_number,omitempty"`
		CreatedDate      *time.Time             `json:"created_date,omitempty"`
		LastModifiedDate *time.Time             `json:"last_modified_date,omitempty"`
		Links            map[string]common.Link `json:"_links,omitempty"`
	}
)
