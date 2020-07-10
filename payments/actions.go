package payments

import (
	"time"

	"github.com/shiuh-yaw-cko/checkout"
)

const (
	// Authorized ...
	Authorized string = "Authorized"
	// Canceled ...
	Canceled string = "Canceled"
	// Captured ...
	Captured string = "Captured"
	// Declined ...
	Declined string = "Declined"
	// Expired ...
	Expired string = "Expired"
	// PartiallyCaptured ...
	PartiallyCaptured string = "Partially Captured"
	// PartiallyRefunded ...
	PartiallyRefunded string = "Partially Refunded"
	// Pending ...
	Pending string = "Pending"
	// Refunded ...
	Refunded string = "Refunded"
	// Voided ...
	Voided string = "Voided"
	// CardVerified ...
	CardVerified string = "Card Verified"
	// Chargeback ...
	Chargeback string = "Chargeback"
)

type (
	// ActionsResponse ...
	ActionsResponse struct {
		StatusResponse *checkout.StatusResponse
		Actions        []*Action
	}
)

type (
	// Action ...
	Action struct {
		ID              string            `json:"id,omitempty"`
		Type            string            `json:"type,omitempty"`
		ProcessedOn     time.Time         `json:"processed_on,omitempty"`
		Amount          uint64            `json:"amount,omitempty"`
		Approved        *bool             `json:"approved,omitempty"`
		AuthCode        string            `json:"auth_code,omitempty"`
		Reference       string            `json:"reference,omitempty"`
		ResponseCode    string            `json:"response_code,omitempty"`
		ResponseSummary *string           `json:"response_summary,omitempty"`
		Processing      *Processing       `json:"processing,omitempty"`
		Metadata        map[string]string `json:"metadata,omitempty"`
	}
	// ActionSummary ...
	ActionSummary struct {
		ID              string  `json:"id,omitempty"`
		Type            string  `json:"type,omitempty"`
		ResponseCode    string  `json:"response_code,omitempty"`
		ResponseSummary *string `json:"response_summary,omitempty"`
	}
)
