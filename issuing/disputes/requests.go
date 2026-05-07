package issuing

type (
	CreateDisputeRequest struct {
		TransactionId      string            `json:"transaction_id,omitempty"`
		Reason             string            `json:"reason,omitempty"`
		Evidence           []DisputeEvidence `json:"evidence,omitempty"`
		Amount             *int64            `json:"amount,omitempty"`
		PresentmentMessageId string          `json:"presentment_message_id,omitempty"`
		Justification      string            `json:"justification,omitempty"`
	}

	EscalateDisputeRequest struct {
		Justification      string                `json:"justification,omitempty"`
		AdditionalEvidence []DisputeEvidence     `json:"additional_evidence,omitempty"`
		Amount             *int64                `json:"amount,omitempty"`
		ReasonChange       *DisputeReasonChange  `json:"reason_change,omitempty"`
	}
)
