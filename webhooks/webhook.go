package webhooks

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type WebhookEvent string

const (
	CardVerifiedEvent             WebhookEvent = "card_verified"
	CardVerificationDeclinedEvent WebhookEvent = "card_verification_declined"
	DisputeCanceledEvent          WebhookEvent = "dispute_canceled"
	DisputeEvidenceRequiredEvent  WebhookEvent = "dispute_evidence_required"
	DisputeExpiredEvent           WebhookEvent = "dispute_expired"
	DisputeLostEvent              WebhookEvent = "dispute_lost"
	DisputeResolvedEvent          WebhookEvent = "dispute_resolved"
	DisputeWonEvent               WebhookEvent = "dispute_won"
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
)

type (
	// Request -
	Request struct {
		*Webhook
	}
	// Webhook ...
	Webhook struct {
		URL         string                    `json:"url,omitempty"`
		Active      *bool                     `json:"active,omitempty"`
		Headers     *Headers                  `json:"headers,omitempty"`
		ContentType common.WebhookContentType `json:"content_type,omitempty"`
		EventTypes  []WebhookEvent            `json:"event_types,omitempty"`
	}
	// Headers ...
	Headers struct {
		Authorization string `json:"Authorization,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse     *checkout.StatusResponse `json:"api_response,omitempty"`
		ConfiguredWebhooks []WebhookResponse        `json:"webhooks,omitempty"`
		Webhook            *WebhookResponse         `json:"webhook,omitempty"`
	}

	// WebhookResponse -
	WebhookResponse struct {
		ID          string                    `json:"id,omitempty"`
		URL         string                    `json:"url,omitempty"`
		Active      *bool                     `json:"active,omitempty"`
		Headers     *Headers                  `json:"headers,omitempty"`
		ContentType common.WebhookContentType `json:"content_type,omitempty"`
		EventTypes  []WebhookEvent            `json:"event_types,omitempty"`
		Links       map[string]common.Link    `json:"_links,omitempty"`
		Version     string                    `json:"version,omitempty"`
	}
)
