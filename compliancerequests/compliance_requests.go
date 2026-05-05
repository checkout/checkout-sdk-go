package compliancerequests

import "github.com/checkout/checkout-sdk-go/v2/common"

const complianceRequestsPath = "compliance-requests"

// RequestedField describes a single data field requested from the merchant by a compliance
// authority, including its name, expected type, and any current value.
type RequestedField struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// RequestedFields groups the compliance data fields requested for the sender and recipient
// of a payment.
type RequestedFields struct {
	Sender    []RequestedField `json:"sender,omitempty"`
	Recipient []RequestedField `json:"recipient,omitempty"`
}

// GetComplianceRequestResponse is the response returned when retrieving a compliance request
// for a payment.
type GetComplianceRequestResponse struct {
	HttpMetadata         common.HttpMetadata
	PaymentId            string           `json:"payment_id,omitempty"`
	ClientId             string           `json:"client_id,omitempty"`
	EntityId             string           `json:"entity_id,omitempty"`
	SegmentId            string           `json:"segment_id,omitempty"`
	Amount               string           `json:"amount,omitempty"`
	Currency             string           `json:"currency,omitempty"`
	RecipientName        string           `json:"recipient_name,omitempty"`
	RequestedOn          string           `json:"requested_on,omitempty"`
	Status               string           `json:"status,omitempty"`
	Fields               *RequestedFields `json:"fields,omitempty"`
	TransactionReference string           `json:"transaction_reference,omitempty"`
	SenderReference      string           `json:"sender_reference,omitempty"`
	LastUpdated          string           `json:"last_updated,omitempty"`
	SenderName           string           `json:"sender_name,omitempty"`
	PaymentType          string           `json:"payment_type,omitempty"`
}

// ComplianceRespondedField holds the merchant's answer to a single compliance data request,
// including an optional flag to indicate the value is not available.
type ComplianceRespondedField struct {
	Name         string      `json:"name,omitempty"`
	Value        interface{} `json:"value,omitempty"`
	NotAvailable bool        `json:"not_available"`
}

// ComplianceRespondedFields groups the merchant's responses for the sender and recipient
// fields of a compliance request.
type ComplianceRespondedFields struct {
	Sender    []ComplianceRespondedField `json:"sender,omitempty"`
	Recipient []ComplianceRespondedField `json:"recipient,omitempty"`
}

// RespondToComplianceRequestRequest is the request body for submitting a compliance
// response for a payment.
type RespondToComplianceRequestRequest struct {
	Fields   ComplianceRespondedFields `json:"fields"`
	Comments string                    `json:"comments,omitempty"`
}
