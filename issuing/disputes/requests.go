package issuing

type (
	CreateDisputeRequest struct {
		TransactionId        string                      `json:"transaction_id,omitempty"`
		Reason               string                      `json:"reason,omitempty"`
		Evidence             []DisputeEvidence           `json:"evidence,omitempty"`
		Amount               *int64                      `json:"amount,omitempty"`
		PresentmentMessageId string                      `json:"presentment_message_id,omitempty"`
		Justification        string                      `json:"justification,omitempty"`
		FraudDetails         *IssuingDisputeFraudDetails `json:"fraud_details,omitempty"`
	}

	EscalateDisputeRequest struct {
		Justification      string                      `json:"justification,omitempty"`
		AdditionalEvidence []DisputeEvidence           `json:"additional_evidence,omitempty"`
		Amount             *int64                      `json:"amount,omitempty"`
		ReasonChange       *DisputeReasonChange        `json:"reason_change,omitempty"`
		FraudDetails       *IssuingDisputeFraudDetails `json:"fraud_details,omitempty"`
	}

	// AmendDisputeRequest is the request body for POST /issuing/disputes/{disputeId}/amend.
	AmendDisputeRequest struct {
		Reason                    string                      `json:"reason,omitempty"`
		Amount                    *int64                      `json:"amount,omitempty"`
		Evidence                  []DisputeEvidence           `json:"evidence,omitempty"`
		FraudDetails              *IssuingDisputeFraudDetails `json:"fraud_details,omitempty"`
		ReasonChangeJustification string                      `json:"reason_change_justification,omitempty"`
		ActionResponse            string                      `json:"action_response,omitempty"`
	}
)
