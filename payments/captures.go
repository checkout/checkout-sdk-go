package payments

import (
	"github.com/checkout/checkout-sdk-go"
)

// CaptureType ...
type CaptureType string

const (
	// NonFinal ...
	NonFinal CaptureType = "NonFinal"
	// Final ...
	Final CaptureType = "Final"
)

func (c CaptureType) String() string {
	return string(c)
}

// CapturesRequest ..
type CapturesRequest struct {
	Amount      uint64            `json:"amount,omitempty"`
	Reference   string            `json:"reference,omitempty"`
	CaptureType CaptureType       `json:"capture_type,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// CapturesResponse ...
type CapturesResponse struct {
	StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	Accepted       *Accepted                `json:"accepted,omitempty"`
}
