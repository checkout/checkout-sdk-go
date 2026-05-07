package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type IssuingDisputeStatus string

const (
	DisputeCreated        IssuingDisputeStatus = "created"
	DisputeCanceled       IssuingDisputeStatus = "canceled"
	DisputeProcessing     IssuingDisputeStatus = "processing"
	DisputeActionRequired IssuingDisputeStatus = "action_required"
	DisputeWon            IssuingDisputeStatus = "won"
	DisputeLost           IssuingDisputeStatus = "lost"
)

type IssuingDisputeStatusReason string

const (
	DisputeExpired                           IssuingDisputeStatusReason = "expired"
	DisputeChargebackPending                 IssuingDisputeStatusReason = "chargeback_pending"
	DisputeChargebackEvidenceInvalidOrInsufficient IssuingDisputeStatusReason = "chargeback_evidence_invalid_or_insufficient"
	DisputeChargebackProcessed              IssuingDisputeStatusReason = "chargeback_processed"
	DisputeChargebackRejected               IssuingDisputeStatusReason = "chargeback_rejected"
	DisputeChargebackReversalPending        IssuingDisputeStatusReason = "chargeback_reversal_pending"
	DisputeChargebackReversed               IssuingDisputeStatusReason = "chargeback_reversed"
	DisputeChargebackResponseAccepted       IssuingDisputeStatusReason = "chargeback_response_accepted"
	DisputePrearbitrationPending            IssuingDisputeStatusReason = "prearbitration_pending"
	DisputePrearbitrationEvidenceInvalid    IssuingDisputeStatusReason = "prearbitration_evidence_invalid_or_insufficient"
	DisputePrearbitrationProcessed         IssuingDisputeStatusReason = "prearbitration_processed"
	DisputePrearbitrationRejected          IssuingDisputeStatusReason = "prearbitration_rejected"
	DisputePrearbitrationReversalPending   IssuingDisputeStatusReason = "prearbitration_reversal_pending"
	DisputePrearbitrationReversed          IssuingDisputeStatusReason = "prearbitration_reversed"
	DisputePrearbitrationResponseAccepted  IssuingDisputeStatusReason = "prearbitration_response_accepted"
	DisputeArbitrationPending              IssuingDisputeStatusReason = "arbitration_pending"
	DisputeArbitrationProcessed            IssuingDisputeStatusReason = "arbitration_processed"
	DisputePresentmentReversed             IssuingDisputeStatusReason = "presentment_reversed"
)

type (
	DisputeEvidence struct {
		Name        string `json:"name,omitempty"`
		Content     string `json:"content,omitempty"`
		Description string `json:"description,omitempty"`
	}

	DisputeFileEvidence struct {
		Name string `json:"name,omitempty"`
		Url  string `json:"url,omitempty"`
	}

	DisputeAmount struct {
		Amount   *int64          `json:"amount,omitempty"`
		Currency common.Currency `json:"currency,omitempty"`
	}

	DisputeMerchant struct {
		Id           string              `json:"id,omitempty"`
		Name         string              `json:"name,omitempty"`
		City         string              `json:"city,omitempty"`
		State        string              `json:"state,omitempty"`
		CountryCode  common.Country      `json:"country_code,omitempty"`
		CategoryCode string              `json:"category_code,omitempty"`
		Evidence     []string            `json:"evidence,omitempty"`
	}

	DisputeReasonChange struct {
		Reason        string `json:"reason,omitempty"`
		Justification string `json:"justification,omitempty"`
	}

	DisputeChargeback struct {
		SubmittedOn   *time.Time           `json:"submitted_on,omitempty"`
		Reason        string               `json:"reason,omitempty"`
		Amount        *DisputeAmount       `json:"amount,omitempty"`
		Evidence      []DisputeFileEvidence `json:"evidence,omitempty"`
		Justification string               `json:"justification,omitempty"`
	}

	DisputeRepresentment struct {
		ReceivedOn *time.Time            `json:"received_on,omitempty"`
		Amount     *DisputeAmount        `json:"amount,omitempty"`
		Evidence   []DisputeFileEvidence `json:"evidence,omitempty"`
	}

	DisputePreArbitration struct {
		SubmittedOn         *time.Time            `json:"submitted_on,omitempty"`
		Evidence            []DisputeFileEvidence `json:"evidence,omitempty"`
		Amount              *DisputeAmount        `json:"amount,omitempty"`
		ReasonChange        *DisputeReasonChange  `json:"reason_change,omitempty"`
		Justification       string                `json:"justification,omitempty"`
		MerchantRespondedOn *time.Time            `json:"merchant_responded_on,omitempty"`
		MerchantEvidence    []DisputeFileEvidence `json:"merchant_evidence,omitempty"`
	}

	DisputeArbitration struct {
		SubmittedOn   *time.Time     `json:"submitted_on,omitempty"`
		Amount        *DisputeAmount `json:"amount,omitempty"`
		Justification string         `json:"justification,omitempty"`
		DecidedOn     *time.Time     `json:"decided_on,omitempty"`
	}
)
