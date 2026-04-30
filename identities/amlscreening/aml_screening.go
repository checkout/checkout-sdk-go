package amlscreening

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/identities"
)

const amlScreeningPath = "aml-verifications"

type CreateAmlScreeningRequest struct {
	ApplicantId      string                      `json:"applicant_id"`
	SearchParameters identities.SearchParameters `json:"search_parameters"`
	Monitored        *bool                       `json:"monitored,omitempty"`
}

type AmlScreeningResponse struct {
	HttpMetadata     common.HttpMetadata
	Id               string                        `json:"id,omitempty"`
	CreatedOn        *time.Time                    `json:"created_on,omitempty"`
	ModifiedOn       *time.Time                    `json:"modified_on,omitempty"`
	ApplicantId      string                        `json:"applicant_id,omitempty"`
	Status           identities.AmlScreeningStatus `json:"status,omitempty"`
	SearchParameters *identities.SearchParameters  `json:"search_parameters,omitempty"`
	Monitored        *bool                         `json:"monitored,omitempty"`
}
