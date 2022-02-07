package payments

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"time"
)

// AuthorizationRequest ..
type AuthorizationRequest struct {
	Amount    uint64                 `json:"amount,omitempty"`
	Reference string                 `json:"reference,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// AuthorizationResponse ...
type AuthorizationResponse struct {
	ActionID        string                   `json:"action_id,omitempty"`
	Amount          uint64                   `json:"amount,omitempty"`
	Currency        string                   `json:"currency,omitempty"`
	Approved        *bool                    `json:"approved,omitempty"`
	Status          common.PaymentAction     `json:"status,omitempty"`
	AuthCode        string                   `json:"auth_code,omitempty"`
	ResponseCode    string                   `json:"response_code,omitempty"`
	ResponseSummary string                   `json:"response_summary,omitempty"`
	ExpiresOn       time.Time                `json:"expires_on,omitempty"`
	Balances        *Balances                `json:"balances,omitempty"`
	ProcessedOn     time.Time                `json:"processed_on,omitempty"`
	Reference       string                   `json:"reference,omitempty"`
	Processing      *PaymentProcessing       `json:"processing,omitempty"` // review
	ECI             string                   `json:"eci,omitempty"`
	SchemeID        string                   `json:"scheme_id,omitempty"`
	Risk            RiskAssessment           `json:"risk,omitempty"`
	Links           map[string]common.Link   `json:"_links,omitempty"`
	StatusResponse  *checkout.StatusResponse `json:"api_response,omitempty"`
}

type Balances struct {
	TotalAuthorized    uint64 `json:"total_authorized,omitempty"`
	TotalVoided        uint64 `json:"total_voided,omitempty"`
	AvailableToVoid    uint64 `json:"available_to_void,omitempty"`
	TotalCaptured      uint64 `json:"total_captured,omitempty"`
	AvailableToCapture uint64 `json:"available_to_capture,omitempty"`
	TotalRefunded      uint64 `json:"total_refunded,omitempty"`
	AvailableToRefund  uint64 `json:"available_to_refund,omitempty"`
}

type PaymentProcessing struct {
	RetrievalReferenceNumber string `json:"retrieval_reference_number,omitempty"`
	AcquirerTransactionId    string `json:"acquirer_transaction_id,omitempty"`
	RecommendationCode       string `json:"recommendation_code,omitempty"`
}
