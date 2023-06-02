package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
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
		SecurityPair SecurityPair
		Password     string
		Locale       string
		PhoneNumber  *common.Phone
	}

	ThreeDSUpdateResponse struct {
		HttpMetadata     common.HttpMetadata
		LastModifiedDate *time.Time `json:"last_modified_date,omitempty"`
	}
)

type (
	ThreeDSEnrollmentResponse struct {
		HttpMetadata common.HttpMetadata
		CreatedDate  *time.Time `json:"created_date,omitempty"`
	}

	ThreeDSEnrollmentDetailsResponse struct {
		HttpMetadata     common.HttpMetadata
		Locale           string        `json:"locale,omitempty"`
		Phone            *common.Phone `json:"phone,omitempty"`
		CreatedDate      *time.Time    `json:"created_date,omitempty"`
		LastModifiedDate *time.Time    `json:"last_modified_date,omitempty"`
	}
)
