package events

import (
	"fmt"
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	"github.com/checkout/checkout-sdk-go/v3/payments/nas"
)

const (
	// EventTypePaymentApproved is sent when a payment is approved.
	EventTypePaymentApproved = "payment_approved"
	// EventTypePaymentCaptured is sent when a payment is captured.
	EventTypePaymentCaptured = "payment_captured"
	// EventTypePaymentDeclined is sent when a payment is declined.
	EventTypePaymentDeclined = "payment_declined"
	// EventTypePaymentRefunded is sent when a payment is refunded.
	EventTypePaymentRefunded = "payment_refunded"
	// EventTypePaymentVoided is sent when a payment is voided.
	EventTypePaymentVoided = "payment_voided"
	// EventTypePaymentPending is sent when a payment is pending.
	EventTypePaymentPending = "payment_pending"
	// EventTypePaymentExpired is sent when a payment expires.
	EventTypePaymentExpired = "payment_expired"
	// EventTypePaymentCanceled is sent when a payment is canceled.
	EventTypePaymentCanceled = "payment_canceled"
)

// paymentEventTypes are gateway events that share PaymentEventData.
var paymentEventTypes = map[string]struct{}{
	EventTypePaymentApproved: {},
	EventTypePaymentCaptured: {},
	EventTypePaymentDeclined: {},
	EventTypePaymentRefunded: {},
	EventTypePaymentVoided:   {},
	EventTypePaymentPending:  {},
	EventTypePaymentExpired:  {},
	EventTypePaymentCanceled: {},
}

// PaymentEventData is the shared data object for payment gateway webhook events.
// Field names follow the payment event payload: id, action_id, amount, currency,
// approved, status, auth_code, response_code, response_summary, 3ds, source,
// customer, processed_on, reference, and metadata.
// Omitted optional fields stay at their zero value.
type PaymentEventData struct {
	Id              string                      `json:"id,omitempty"`
	ActionId        string                      `json:"action_id,omitempty"`
	Amount          int64                       `json:"amount,omitempty"`
	Currency        common.Currency             `json:"currency,omitempty"`
	Approved        bool                        `json:"approved,omitempty"`
	Status          payments.PaymentStatus      `json:"status,omitempty"`
	AuthCode        string                      `json:"auth_code,omitempty"`
	ResponseCode    string                      `json:"response_code,omitempty"`
	ResponseSummary string                      `json:"response_summary,omitempty"`
	ThreeDs         *payments.ThreeDsEnrollment `json:"3ds,omitempty"`
	Source          *nas.ResponseCardSource     `json:"source,omitempty"`
	Customer        *common.CustomerResponse    `json:"customer,omitempty"`
	ProcessedOn     *time.Time                  `json:"processed_on,omitempty"`
	Reference       string                      `json:"reference,omitempty"`
	Metadata        map[string]interface{}      `json:"metadata,omitempty"`
}

// DecodePayment decodes data for a payment gateway event into PaymentEventData.
//
// ErrUnexpectedEventType is returned when Type is not a payment event constant.
// The envelope and Data map remain available to the caller.
// ErrMissingEventData is returned when data is absent or null.
// A payload with an empty id returns an error wrapping ErrMissingEventField.
// A data field whose JSON type does not match PaymentEventData is returned
// wrapped around the encoding/json error.
func (e *EventResponse) DecodePayment() (*PaymentEventData, error) {
	if e == nil {
		return nil, ErrMissingEventData
	}
	if _, ok := paymentEventTypes[e.Type]; !ok {
		return nil, ErrUnexpectedEventType
	}

	var payload PaymentEventData
	if err := e.DecodeData(&payload); err != nil {
		return nil, err
	}
	if payload.Id == "" {
		return nil, fmt.Errorf("events: missing event data field %q: %w", "id", ErrMissingEventField)
	}
	return &payload, nil
}
