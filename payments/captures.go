package payments

import "github.com/checkout/checkout-sdk-go"

// CapturesRequest ..
type CapturesRequest struct {
	Amount    uint64            `json:"amount,omitempty"`
	Reference string            `json:"reference,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// CapturesResponse ...
type CapturesResponse struct {
	StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	Accepted       *Accepted                `json:"accepted,omitempty"`
}
