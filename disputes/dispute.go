package disputes

import (
	"time"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/common"
)

type (
	// Request -
	Request struct {
		*QueryParameter
	}

	// QueryParameter -
	QueryParameter struct {
		Limit            uint64    `url:"limit,omitempty"`
		Skip             uint64    `url:"skip,omitempty"`
		From             time.Time `url:"from,omitempty"`
		To               time.Time `url:"to,omitempty"`
		ID               string    `url:"id,omitempty"`
		Statuses         string    `url:"statuses,omitempty"`
		PaymentID        string    `url:"payment_id,omitempty"`
		PaymentReference string    `url:"payment_reference,omitempty"`
		PaymentARN       string    `url:"payment_arn,omitempty"`
		ThisChannelOnly  *bool     `url:"this_channel_only,omitempty"`
	}
)
type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Disputes       *Disputes                `json:"disputes,omitempty"`
		Dispute        *Dispute                 `json:"dispute,omitempty"`
		Evidence       *map[string]string       `json:"evidence,omitempty"`
	}

	// Disputes -
	Disputes struct {
		Limit            uint64           `json:"limit,omitempty"`
		Skip             uint64           `json:"skip,omitempty"`
		From             time.Time        `json:"from,omitempty"`
		To               time.Time        `json:"to,omitempty"`
		Statuses         string           `json:"statuses,omitempty"`
		ID               string           `json:"id,omitempty"`
		PaymentID        string           `json:"payment_id,omitempty"`
		PaymentReference string           `json:"payment_reference,omitempty"`
		PaymentARN       string           `json:"payment_arn,omitempty"`
		ThisChannelOnly  *bool            `json:"this_channel_only,omitempty"`
		TotalCount       uint64           `json:"total_count,omitempty"`
		Data             []DisputeSummary `json:"data,omitempty"`
	}

	// DisputeSummary -
	DisputeSummary struct {
		ID                 string                 `json:"id,omitempty"`
		Category           string                 `json:"category,omitempty"`
		Status             string                 `json:"status,omitempty"`
		Amount             uint64                 `json:"amount,omitempty"`
		Currency           string                 `json:"currency,omitempty"`
		PaymentID          string                 `url:"payment_id,omitempty"`
		PaymentReference   string                 `url:"payment_reference,omitempty"`
		PaymentARN         string                 `url:"payment_arn,omitempty"`
		PaymentMethod      string                 `json:"payment_method,omitempty"`
		EvidenceRequiredBy time.Time              `json:"evidence_required_by,omitempty"`
		ReceivedOn         time.Time              `json:"received_on,omitempty"`
		LastUpdate         time.Time              `json:"last_update,omitempty"`
		Links              map[string]common.Link `json:"_links,omitempty"`
	}
	// Dispute -
	Dispute struct {
		ID                 string                 `json:"id,omitempty"`
		Category           string                 `json:"category,omitempty"`
		Currency           string                 `json:"currency,omitempty"`
		ReasonCode         string                 `json:"reason_code,omitempty"`
		RelevantEvidence   []string               `json:"relevant_evidence,omitempty"`
		EvidenceRequiredBy time.Time              `json:"evidence_required_by,omitempty"`
		ReceivedOn         time.Time              `json:"received_on,omitempty"`
		LastUpdate         time.Time              `json:"last_update,omitempty"`
		Payment            *Payment               `json:"payment,omitempty"`
		Links              map[string]common.Link `json:"_links,omitempty"`
	}
	// Payment -
	Payment struct {
		ID          string    `json:"id,omitempty"`
		Amount      uint64    `json:"amount,omitempty"`
		Currency    string    `json:"currency,omitempty"`
		Method      string    `json:"method,omitempty"`
		ARN         string    `json:"arn,omitempty"`
		ProcessedOn time.Time `json:"processed_on,omitempty"`
	}

	// Evidence -
	Evidence struct {
		Links map[string]string `json:"-,omitempty"`
	}
)
