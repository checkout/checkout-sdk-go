package payments

import "github.com/shiuh-yaw-cko/checkout"

// CapturesRequest ..
type CapturesRequest struct {
	Amount    uint64            `json:"amount,omitempty"`
	Reference string            `json:"reference,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// CapturesResponse ...
type CapturesResponse struct {
	StatusResponse *checkout.StatusResponse
	Accepted       *Accepted
}
