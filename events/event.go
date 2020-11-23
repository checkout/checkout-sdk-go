package events

import (
	"time"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	// Request -
	Request struct {
		*QueryParameter
		*EventTypeRequest
	}

	// QueryParameter -
	QueryParameter struct {
		From      time.Time `url:"from,omitempty"`
		To        time.Time `url:"to,omitempty"`
		Limit     uint64    `url:"limit,omitempty"`
		PaymentID string    `url:"payment_id,omitempty"`
	}

	// EventTypeRequest -
	EventTypeRequest struct {
		Version string `url:"version,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		EventTypes     []EventType              `json:"event_types,omitempty"`
		Events         *Events                  `json:"events,omitempty"`
		Event          *Event                   `json:"event,omitempty"`
		Notification   *Notification            `json:"notification,omitempty"`
	}
	// EventType -
	EventType struct {
		Version    string   `json:"version,omitempty"`
		EventTypes []string `json:"event_types,omitempty"`
	}
	// Events -
	Events struct {
		TotalCount uint64    `json:"total_count,omitempty"`
		Limit      uint64    `json:"limit,omitempty"`
		Skip       uint64    `json:"skip,omitempty"`
		From       time.Time `json:"from,omitempty"`
		To         time.Time `json:"to,omitempty"`
		Data       []Event   `json:"data,omitempty"`
	}

	// Event -
	Event struct {
		ID            string                 `json:"id,omitempty"`
		Type          string                 `json:"type,omitempty"`
		Version       string                 `json:"version,omitempty"`
		CreatedOn     string                 `json:"created_on,omitempty"`
		Data          *payments.Processed    `json:"data,omitempty"`
		Notifications []Notification         `json:"notifications,omitempty"`
		Links         map[string]common.Link `json:"_links"`
	}
	// Notification -
	Notification struct {
		ID          string                 `json:"id,omitempty"`
		URL         string                 `json:"url,omitempty"`
		Success     *bool                  `json:"success,omitempty"`
		ContentType string                 `json:"content_type,omitempty"`
		Attempts    []NotificationAttempt  `json:"attempts,omitempty"`
		Links       map[string]common.Link `json:"_links"`
	}
	// NotificationAttempt -
	NotificationAttempt struct {
		StatusCode   uint64    `json:"status_code,omitempty"`
		ResponseBody string    `json:"response_body,omitempty"`
		RetryMode    string    `json:"retry_mode,omitempty"`
		Timestamp    time.Time `json:"timestamp,omitempty"`
	}
)
