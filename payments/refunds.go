package payments

import "github.com/shiuh-yaw-cko/checkout"

// RefundsRequest ..
type RefundsRequest struct {
	Amount    uint64            `json:"amount,omitempty"`
	Reference string            `json:"reference,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// RefundsResponse ...
type RefundsResponse struct {
	StatusResponse *checkout.StatusResponse
	Accepted       *Accepted
}
