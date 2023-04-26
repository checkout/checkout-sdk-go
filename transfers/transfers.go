package transfers

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const (
	transfers = "transfers"
)

type TransferType string

const (
	Commission TransferType = "commission"
	Promotion  TransferType = "promotion"
	Refund     TransferType = "refund"
)

// Request
type (
	TransferSourceRequest struct {
		Id     string `json:"id,omitempty"`
		Amount int64  `json:"amount,omitempty"`
	}

	TransferDestinationRequest struct {
		Id string `json:"id,omitempty"`
	}

	TransferRequest struct {
		Reference    string                      `json:"reference,omitempty"`
		TransferType TransferType                `json:"transfer_type,omitempty"`
		Source       *TransferSourceRequest      `json:"source,omitempty"`
		Destination  *TransferDestinationRequest `json:"destination,omitempty"`
	}
)

// Response
type (
	TransferResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Status       string                 `json:"status,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	TransferSourceResponse struct {
		EntityId string          `json:"entity_id,omitempty"`
		Amount   int64           `json:"amount,omitempty"`
		Currency common.Currency `json:"currency,omitempty"`
	}

	TransferDestinationResponse struct {
		EntityId string `json:"entity_id,omitempty"`
	}

	TransferDetails struct {
		HttpMetadata common.HttpMetadata
		Id           string                       `json:"id,omitempty"`
		Reference    string                       `json:"reference,omitempty"`
		Status       string                       `json:"status,omitempty"`
		TransferType TransferType                 `json:"transfer_type,omitempty"`
		RequestedOn  *time.Time                   `json:"requested_on,omitempty"`
		ReasonCodes  []string                     `json:"reason_codes,omitempty"`
		Source       *TransferSourceResponse      `json:"source,omitempty"`
		Destination  *TransferDestinationResponse `json:"destination,omitempty"`
		Links        map[string]common.Link       `json:"_links"`
	}
)
