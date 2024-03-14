package payment_sessions

import "github.com/checkout/checkout-sdk-go/common"

const PaymentSessionsPath = "payment-sessions"

type (
	Billing struct {
		Address *common.Address `json:"address,omitempty"`
	}

	PaymentSessionsRequest struct {
		Amount     int64                   `json:"amount,omitempty"`
		Currency   common.Currency         `json:"currency,omitempty"`
		Reference  string                  `json:"reference,omitempty"`
		Billing    *Billing                `json:"billing,omitempty"`
		Customer   *common.CustomerRequest `json:"customer,omitempty"`
		SuccessUrl string                  `json:"success_url,omitempty"`
		FailureUrl string                  `json:"failure_url,omitempty"`
	}
)

type (
	PaymentMethods struct {
		Type        string   `json:"type,omitempty"`
		CardSchemes []string `json:"card_schemes,omitempty"`
	}

	PaymentSessionsResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Links        map[string]common.Link `json:"links,omitempty"`
	}
)
