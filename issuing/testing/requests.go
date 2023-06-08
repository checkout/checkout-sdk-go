package issuing

import "github.com/checkout/checkout-sdk-go/common"

type TransactionType string

const (
	Purchase TransactionType = "purchase"
)

type (
	CardSimulation struct {
		Id          string `json:"id,omitempty"`
		ExpiryMonth int    `json:"expiry_month,omitempty"`
		ExpiryYear  int    `json:"expiry_year,omitempty"`
	}

	TransactionSimulation struct {
		Type     TransactionType `json:"type,omitempty"`
		Amount   int             `json:"amount,omitempty"`
		Currency common.Currency `json:"currency,omitempty"`
	}

	CardAuthorizationRequest struct {
		Card        CardSimulation        `json:"card,omitempty"`
		Transaction TransactionSimulation `json:"transaction,omitempty"`
	}

	CardSimulationRequest struct {
		Amount int `json:"amount,omitempty"`
	}
)
