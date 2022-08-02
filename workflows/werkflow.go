package workflows

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type WebhookEvent string

const (
	// Gateway
	CardVerifiedEvent             WebhookEvent = "card_verified"
	CardVerificationDeclinedEvent WebhookEvent = "card_verification_declined"
	PaymentApprovedEvent          WebhookEvent = "payment_approved"
	PaymentPendingEvent           WebhookEvent = "payment_pending"
	PaymentDeclinedEvent          WebhookEvent = "payment_declined"
	PaymentExpiredEvent           WebhookEvent = "payment_expired"
	PaymentVoidedEvent            WebhookEvent = "payment_voided"
	PaymentCanceledEvent          WebhookEvent = "payment_canceled"
	PaymentVoidDeclinedEvent      WebhookEvent = "payment_void_declined"
	PaymentCapturedEvent          WebhookEvent = "payment_captured"
	PaymentCaptureDeclinedEvent   WebhookEvent = "payment_capture_declined"
	PaymentCapturePendingEvent    WebhookEvent = "payment_capture_pending"
	PaymentRefundedEvent          WebhookEvent = "payment_refunded"
	PaymentRefundDeclinedEvent    WebhookEvent = "payment_refund_declined"
	PaymentRefundPendingEvent     WebhookEvent = "payment_refund_pending"
	PaymentChargebackEvent        WebhookEvent = "payment_chargeback"
	PaymentRetrievalEvent         WebhookEvent = "payment_retrieval"
	SourceUpdatedEvent            WebhookEvent = "source_updated"
	PaymentPaidEvent              WebhookEvent = "payment_paid"

	// Dispute
	DisputeCanceledEvent          WebhookEvent = "dispute_canceled"
	DisputeEvidenceRequiredEvent  WebhookEvent = "dispute_evidence_required"
	DisputeExpiredEvent           WebhookEvent = "dispute_expired"
	DisputeLostEvent              WebhookEvent = "dispute_lost"
	DisputeResolvedEvent          WebhookEvent = "dispute_resolved"
	DisputeWonEvent               WebhookEvent = "dispute_won"
)

type ConditionType string

const (
	EventType         ConditionType = "event"
	EntityType        ConditionType = "entity"
	ProcessingChannel ConditionType = "processing_channel"
)

type ActionType string

const (
	WebhookType ActionType = "webhook"
)

type (
	Request struct {
		*Workflow
	}
	Workflow struct {
		Name         string                   `json:"name,omitempty"`
		Active       bool                     `json:"active,omitempty"`
		Conditions   []Condition              `json:"conditions,omitempty"`
		Actions      []Action                 `json:"actions,omitempty"`
	}
	Condition struct {
		Type   ConditionType `json:"type,omitempty"`
		Events Events        `json:"events,omitempty"`
	}
	Events struct {
		Gateway []WebhookEvent `json:"gateway,omitempty"`
		Dispute []WebhookEvent `json:"dispute,omitempty"`
	}
	Action struct {
		Type      ActionType `json:"type,omitempty"`
		URL       string     `json:"url,omitempty"`
		Headers   *Headers   `json:"headers,omitempty"`
		Signature *Signature `json:"signature,omitempty"`

	}
	Headers struct {
		Authorization string `json:"Authorization,omitempty"`
	}
	Signature struct {
		Method string `json:"method,omitempty"`
		Key    string `json:"key,omitempty"`
	}
)

type (
	Response struct {
		StatusResponse     *checkout.StatusResponse `json:"api_response,omitempty"`
		Workflows          []WorkflowResponse       `json:"workflows,omitempty"`
		Workflow           *WorkflowResponse        `json:"workflow,omitempty"`
	}

	WorkflowResponse struct {
		ID          string                    `json:"id,omitempty"`
		Name        string                    `json:"name,omitempty"`
		Active      bool                      `json:"active,omitempty"`
		Conditions  []ConditionResponse       `json:"conditions,omitempty"`
		Actions     []ActionResponse          `json:"actions,omitempty"`
	}

	ConditionResponse struct {
		ID                 string         `json:"id,omitempty"`
		Type               ConditionType  `json:"type,omitempty"`
		Events             Events         `json:"events,omitempty"`
		Entities           []string       `json:"entities,omitempty"`
		ProcessingChannels []string       `json:"processing_channels,omitempty"`
		Links              []common.Link  `json:"links,omitempty"`
	}

	ActionResponse struct {
		ID         string        `json:"id,omitempty"`
		Type       ActionType    `json:"type,omitempty"`
		URL        string        `json:"url,omitempty"`
		Headers    Headers       `json:"headers,omitempty"`
		Signature  Signature     `json:"signature,omitempty"`
		Links      []common.Link `json:"links,omitempty"`
	}
)

type TestRequest struct {
	EventTypes []WebhookEvent `json:"event_types"`
}
