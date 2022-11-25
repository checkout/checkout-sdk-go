package sepa

import "github.com/checkout/checkout-sdk-go-beta/common"

const (
	sepaMandatesPath = "sepa/mandates"
	pproPath         = "ppro"
	cancelPath       = "cancel"
)

type (
	MandateResponse struct {
		HttpMetadata        common.HttpMetadata    `json:"http_metadata,omitempty"`
		MandateReference    string                 `json:"mandate_reference,omitempty"`
		CustomerId          string                 `json:"customer_id,omitempty"`
		FirstName           string                 `json:"first_name,omitempty"`
		LastName            string                 `json:"last_name,omitempty"`
		AddressLine1        string                 `json:"address_line1,omitempty"`
		City                string                 `json:"city,omitempty"`
		Zip                 string                 `json:"zip,omitempty"`
		Country             common.Country         `json:"country,omitempty"`
		MaskedAccountIban   string                 `json:"masked_account_iban,omitempty"`
		AccountCurrencyCode string                 `json:"account_currency_code,omitempty"`
		AccountCountryCode  common.Country         `json:"account_country_code,omitempty"`
		MandateState        string                 `json:"mandate_state,omitempty"`
		BillingDescriptor   string                 `json:"billing_descriptor,omitempty"`
		MandateType         string                 `json:"mandate_type,omitempty"`
		Links               map[string]common.Link `json:"_links,omitempty"`
	}

	SepaResource struct {
		HttpMetadata common.HttpMetadata    `json:"http_metadata,omitempty"`
		Links        map[string]common.Link `json:"_links,omitempty"`
	}
)
