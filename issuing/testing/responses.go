package issuing

import "github.com/checkout/checkout-sdk-go/common"

type TransactionStatus string

const (
	Authorized TransactionStatus = "Authorized"
	Declined   TransactionStatus = "Declined"
	Reversed   TransactionStatus = "Reversed"
)

type (
	CardAuthorizationResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string            `json:"id,omitempty"`
		Status       TransactionStatus `json:"status,omitempty"`
	}

	CardSimulationResponse struct {
		HttpMetadata common.HttpMetadata
		Status       TransactionStatus `json:"status,omitempty"`
	}
)
