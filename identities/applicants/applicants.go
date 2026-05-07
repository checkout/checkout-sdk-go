package applicants

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

const applicantsPath = "applicants"
const anonymizePath = "anonymize"

type CreateApplicantRequest struct {
	ExternalApplicantId   string `json:"external_applicant_id,omitempty"`
	Email                 string `json:"email,omitempty"`
	ExternalApplicantName string `json:"external_applicant_name,omitempty"`
}

type UpdateApplicantRequest struct {
	Email                 string `json:"email,omitempty"`
	ExternalApplicantName string `json:"external_applicant_name,omitempty"`
}

type ApplicantResponse struct {
	HttpMetadata          common.HttpMetadata
	Id                    string     `json:"id,omitempty"`
	CreatedOn             *time.Time `json:"created_on,omitempty"`
	ModifiedOn            *time.Time `json:"modified_on,omitempty"`
	ExternalApplicantId   string     `json:"external_applicant_id,omitempty"`
	Email                 string     `json:"email,omitempty"`
	ExternalApplicantName string     `json:"external_applicant_name,omitempty"`
}
