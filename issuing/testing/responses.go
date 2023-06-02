package issuing

import "github.com/checkout/checkout-sdk-go/common"

type TransactionStatus string

const (
	Authorized TransactionStatus = "Authorized"
	Declined   TransactionStatus = "Declined"
)

type (
	CardAuthorizationResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string            `json:"id,omitempty"`
		Status       TransactionStatus `json:"status,omitempty"`
	}
)
