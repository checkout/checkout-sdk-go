package sources

import (
	"github.com/checkout/checkout-sdk-go/common"
)

const path = "sources"

type SourceType string
type MandateType string

const (
	Sepa SourceType = "Sepa"
)

const (
	Single    MandateType = "single"
	Recurring MandateType = "recurring"
)

// Request
type (
	SourceData struct {
		FirstName         string      `json:"first_name,omitempty"`
		LastName          string      `json:"last_name,omitempty"`
		AccountIban       string      `json:"account_iban,omitempty"`
		Bic               string      `json:"bic,omitempty"`
		BillingDescriptor string      `json:"billing_descriptor,omitempty"`
		MandateType       MandateType `json:"mandate_type,omitempty"`
	}

	sepaSourceRequest struct {
		Type            SourceType              `json:"type" binding:"required"`
		BillingAddress  *common.Address         `json:"billing_address,omitempty"`
		SourceData      *SourceData             `json:"source_data,omitempty"`
		Reference       string                  `json:"reference,omitempty"`
		Phone           *common.Phone           `json:"phone,omitempty"`
		CustomerRequest *common.CustomerRequest `json:"customer,omitempty"`
	}
)

func NewSepaSourceRequest() *sepaSourceRequest {
	return &sepaSourceRequest{
		Type: Sepa,
	}
}

// Response
type (
	CreateSepaSourceResponse struct {
		HttpResponse   common.HttpMetadata
		SourceResponse *SourceResponse
		ResponseData   map[string]string `json:"response_data,omitempty"`
	}

	SourceResponse struct {
		SourceType   SourceType               `json:"type,omitempty"`
		Id           string                   `json:"id,omitempty"`
		ResponseCode string                   `json:"response_code,omitempty"`
		Customer     *common.CustomerResponse `json:"customer,omitempty"`
		Links        map[string]common.Link   `json:"_links"`
	}
)
