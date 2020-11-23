package sources

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type (
	// Request -
	Request struct {
		*SEPA
		*ACH
	}

	// SEPA -
	SEPA struct {
		Type           string          `json:"type" binding:"required"`
		Reference      string          `json:"reference,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
		Customer       *Customer       `json:"customer,omitempty"`
		BillingAddress *common.Address `json:"billing_address" binding:"required"`
		SourceData     *SEPASourceData `json:"source_data,omitempty"`
	}

	// ACH -
	ACH struct {
		Type           string          `json:"type" binding:"required"`
		Reference      string          `json:"reference,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
		Customer       *Customer       `json:"customer,omitempty"`
		BillingAddress *common.Address `json:"billing_address" binding:"required"`
		SourceData     *ACHSourceData  `json:"source_data,omitempty"`
	}

	// Customer -
	Customer struct {
		ID    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
	}

	// SEPASourceData -
	SEPASourceData struct {
		FirstName         string `json:"first_name,omitempty"`
		LastName          string `json:"last_name,omitempty"`
		AccountIBAN       string `json:"account_iban,omitempty"`
		BIC               string `json:"bic,omitempty"`
		BillingDescriptor string `json:"billing_descriptor,omitempty"`
		MandateType       string `json:"mandate_type,omitempty"`
	}

	// ACHSourceData -
	ACHSourceData struct {
		AccountType       string `json:"account_type,omitempty"`
		AccountNumber     string `json:"account_number,omitempty"`
		RoutingNumber     string `json:"routing_number,omitempty"`
		AccountHolderName string `json:"account_holder_name,omitempty"`
		BillingDescriptor string `json:"billing_descriptor,omitempty"`
		CompanyName       string `json:"company_name,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Source         *Source                  `json:"source,omitempty"`
	}
	// Source -
	Source struct {
		ID           string                 `json:"id,omitempty"`
		Type         string                 `json:"type,omitempty"`
		ResponseCode string                 `json:"response_code,omitempty"`
		Customer     *Customer              `json:"response_data,omitempty"`
		ResponseData *ResponseData          `json:"uploaded_on,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	// ResponseData -
	ResponseData struct {
		MandateReference string `json:"mandate_reference,omitempty"`
	}
)
