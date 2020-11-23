package payments

import "github.com/checkout/checkout-sdk-go"

// RefundsRequest ..
type RefundsRequest struct {
	Amount    uint64            `json:"amount,omitempty"`
	Reference string            `json:"reference,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// RefundsResponse ...
type RefundsResponse struct {
	StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	Accepted       *Accepted                `json:"accepted,omitempty"`
}
