package disputes

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const (
	disputes    = "disputes"
	accept      = "accept"
	evidence    = "evidence"
	files       = "files"
	schemeFiles = "schemefiles"
)

type (
	Dispute struct {
		Id                 string                `json:"id,omitempty"`
		Category           DisputeCategory       `json:"category,omitempty"`
		Status             DisputeStatus         `json:"status,omitempty"`
		Amount             int64                 `json:"amount,omitempty"`
		Currency           common.Currency       `json:"currency,omitempty"`
		ReasonCode         string                `json:"reason_code,omitempty"`
		ResolvedReason     DisputeResolvedReason `json:"resolved_reason,omitempty"`
		RelevantEvidence   []RelevantEvidence    `json:"relevant_evidence,omitempty"`
		EvidenceRequiredBy time.Time             `json:"evidence_required_by,omitempty"`
		ReceivedOn         time.Time             `json:"received_on,omitempty"`
		LastUpdate         time.Time             `json:"last_update,omitempty"`
		Payment            *PaymentDispute       `json:"payment,omitempty"`

		// Not available on Previous
		EntityId    string                 `json:"entity_id,omitempty"`
		SubEntityId string                 `json:"sub_entity_id,omitempty"`
		Links       map[string]common.Link `json:"_links,omitempty"`
	}

	PaymentDispute struct {
		Id          string          `json:"id,omitempty"`
		Amount      int64           `json:"amount,omitempty"`
		Currency    common.Currency `json:"currency,omitempty"`
		Method      string          `json:"method,omitempty"`
		Arn         string          `json:"arn,omitempty"`
		ProcessedOn time.Time       `json:"processed_on,omitempty"`

		// Not available on Previous
		ActionId                 string                    `json:"actionId,omitempty"`
		ProcessingChannelId      string                    `json:"processing_channel_id,omitempty"`
		Mcc                      string                    `json:"mcc,omitempty"`
		ThreeDSVersionEnrollment *ThreeDSVersionEnrollment `json:"3ds,omitempty"`
		Eci                      string                    `json:"eci,omitempty"`
		HasRefund                bool                      `json:"has_refund,omitempty"`
	}

	ThreeDSVersionEnrollment struct {
		Enrolled string `json:"enrolled,omitempty"`
		Version  string `json:"version,omitempty"`
	}
)

type (
	DisputeResponse struct {
		HttpMetadata common.HttpMetadata
		Dispute
	}
)

// Query
type (
	QueryFilter struct {
		Limit uint8     `url:"limit,omitempty"` //min=1 - max=250
		Skip  int       `url:"skip,omitempty"`
		From  time.Time `url:"from,omitempty" layout:"2006-01-02T15:04:05Z"`
		To    time.Time `url:"to,omitempty" layout:"2006-01-02T15:04:05Z"`

		Id               string `url:"id,omitempty"`
		Statuses         string `url:"statuses,omitempty"` //One or more comma-separated statuses. This works like a logical OR operator
		PaymentId        string `url:"payment_id,omitempty"`
		PaymentReference string `url:"payment_reference,omitempty"`
		PaymentArn       string `url:"payment_arn,omitempty"`
		ThisChannelOnly  bool   `url:"this_channel_only,omitempty"`

		// Not available on Previous
		EntityIds    string `url:"entity_ids,omitempty"`    //One or more comma-separated client entities. This works like a logical OR operator
		SubEntityIds string `url:"subEntity_ids,omitempty"` //One or more comma-separated client entities. This works like a logical OR operator
		PaymentMcc   string `url:"payment_mcc,omitempty"`
	}

	QueryResponse struct {
		HttpMetadata common.HttpMetadata

		Limit uint8     `json:"limit,omitempty"` //min=1 - max=250
		Skip  int       `json:"skip,omitempty"`
		From  time.Time `json:"from,omitempty" time_format:"2006-01-02T15:04:05Z"`
		To    time.Time `json:"to,omitempty" time_format:"2006-01-02T15:04:05Z"`

		Id               string `json:"id,omitempty"`
		Statuses         string `json:"statuses,omitempty"` //One or more comma-separated statuses. This works like a logical OR operator
		PaymentId        string `json:"payment_id,omitempty"`
		PaymentReference string `json:"payment_reference,omitempty"`
		PaymentArn       string `json:"payment_arn,omitempty"`
		ThisChannelOnly  bool   `json:"this_channel_only,omitempty"`

		TotalCount int              `json:"total_count,omitempty"`
		Data       []DisputeSummary `json:"data,omitempty"`

		// Not available on Previous
		EntityIds    string `url:"entity_ids,omitempty"`    //One or more comma-separated client entities. This works like a logical OR operator//One or more comma-separated client entities. This works like a logical OR operator
		SubEntityIds string `url:"subEntity_ids,omitempty"` //One or more comma-separated client entities. This works like a logical OR operator//One or more comma-separated client entities. This works like a logical OR operator
		PaymentMcc   string `url:"payment_mcc,omitempty"`
	}

	DisputeSummary struct {
		Id                 string          `json:"id,omitempty"`
		Category           DisputeCategory `json:"category,omitempty"`
		Status             DisputeStatus   `json:"status,omitempty"`
		Amount             int64           `json:"amount,omitempty"`
		Currency           common.Currency `json:"currency,omitempty"`
		ReasonCode         string          `json:"reason_code,omitempty"`
		PaymentId          string          `json:"payment_id,omitempty"`
		PaymentActionId    string          `json:"payment_action_id,omitempty"`
		PaymentReference   string          `json:"payment_reference,omitempty"`
		PaymentArn         string          `json:"payment_arn,omitempty"`
		PaymentMethod      string          `json:"payment_method,omitempty"`
		EvidenceRequiredBy time.Time       `json:"evidence_required_by,omitempty"`
		ReceivedOn         time.Time       `json:"received_on,omitempty"`
		LastUpdate         time.Time       `json:"last_update,omitempty"`

		// Not available on Previous
		EntityId    string `json:"entity_id,omitempty"`
		SubEntityId string `json:"sub_entity_id,omitempty"`
		PaymentMcc  string `json:"payment_mcc,omitempty"`
	}
)

// Evidence
type (
	Evidence struct {
		ProofOfDeliveryOrServiceFile           string                 `json:"proof_of_delivery_or_service_file,omitempty"`
		ProofOfDeliveryOrServiceText           string                 `json:"proof_of_delivery_or_service_text,omitempty"` // max 500
		InvoiceOrReceiptFile                   string                 `json:"invoice_or_receipt_file,omitempty"`
		InvoiceOrReceiptText                   string                 `json:"invoice_or_receipt_text,omitempty"`
		InvoiceShowingDistinctTransactionsFile string                 `json:"invoice_showing_distinct_transactions_file,omitempty"`
		InvoiceShowingDistinctTransactionsText string                 `json:"invoice_showing_distinct_transactions_text,omitempty"` // max 500
		CustomerCommunicationFile              string                 `json:"customer_communication_file,omitempty"`
		CustomerCommunicationText              string                 `json:"customer_communication_text,omitempty"` // max 500
		RefundOrCancellationPolicyFile         string                 `json:"refund_or_cancellation_policy_file,omitempty"`
		RefundOrCancellationPolicyText         string                 `json:"refund_or_cancellation_policy_text,omitempty"` // max 500
		RecurringTransactionAgreementFile      string                 `json:"recurring_transaction_agreement_file,omitempty"`
		RecurringTransactionAgreementText      string                 `json:"recurring_transaction_agreement_text,omitempty"` // max 500
		AdditionalEvidenceFile                 string                 `json:"additional_evidence_file,omitempty"`
		AdditionalEvidenceText                 string                 `json:"additional_evidence_text,omitempty"` // max 500
		ProofOfDeliveryOrServiceDateFile       string                 `json:"proof_of_delivery_or_service_date_file,omitempty"`
		ProofOfDeliveryOrServiceDateText       string                 `json:"proof_of_delivery_or_service_date_text,omitempty"` // max 500
		Links                                  map[string]common.Link `json:"_links,omitempty"`
	}

	EvidenceResponse struct {
		HttpMetadata common.HttpMetadata
		Evidence
	}
)

// Files
type (
	SchemeFilesResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Files        []SchemeFile           `json:"files,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	SchemeFile struct {
		DisputeStatus DisputeStatus `json:"dispute_status,omitempty"`
		File          string        `json:"file,omitempty"`
	}
)
