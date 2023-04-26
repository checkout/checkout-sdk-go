package abc

import (
	"time"

	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/abc"
)

const (
	events        = "events"
	eventTypes    = "event-types"
	notifications = "notifications"
	webhooks      = "webhooks"
	retry         = "retry"
)

type EventDataStatus string

const (
	Pending           EventDataStatus = "Pending"
	Authorized        EventDataStatus = "Authorized"
	Voided            EventDataStatus = "Voided"
	PartiallyCaptured EventDataStatus = "Partially Captured"
	Captured          EventDataStatus = "Captured"
	PartiallyRefunded EventDataStatus = "Partially Refunded"
	Refunded          EventDataStatus = "Refunded"
	Declined          EventDataStatus = "Declined"
	Canceled          EventDataStatus = "Canceled"
)

type (
	QueryRetrieveAllEventType struct {
		Version string `url:"version,omitempty"`
	}

	EventTypes struct {
		Version    string   `json:"version"`
		EventTypes []string `json:"event_types"`
	}

	EventTypesResponse struct {
		HttpResponse common.HttpMetadata
		EventTypes   []EventTypes
	}
)

type (
	QueryRetrieveEvents struct {
		PaymentId string `url:"payment_id,omitempty"`
		ChargeId  string `url:"charge_id,omitempty"`
		TrackId   string `url:"track_id,omitempty"`
		Reference string `url:"reference,omitempty"`
		Skip      int    `url:"skip,omitempty"`
		Limit     int    `url:"limit,omitempty"`
	}

	EventsSummaryResponse struct {
		Id        string                 `json:"id,omitempty"`
		Type      string                 `json:"type,omitempty"`
		CreatedOn string                 `json:"created_on,omitempty"`
		Links     map[string]common.Link `json:"_links"`
	}

	EventsPageResponse struct {
		HttpResponse common.HttpMetadata
		TotalCount   int                     `json:"total_count,omitempty"`
		Limit        int                     `json:"limit,omitempty"`
		Skip         int                     `json:"skip,omitempty"`
		Data         []EventsSummaryResponse `json:"data,omitempty"`
	}
)

type (
	EventNotificationSummaryResponse struct {
		Id      string                 `json:"id,omitempty"`
		Url     string                 `json:"url,omitempty"`
		Success bool                   `json:"success,omitempty"`
		Links   map[string]common.Link `json:"_links"`
	}

	EventPaymentData struct {
		Id              string                      `json:"id,omitempty"`
		ActionId        string                      `json:"action_id,omitempty"`
		Amount          int64                       `json:"amount,omitempty"`
		Currency        common.Currency             `json:"currency,omitempty"`
		Approved        bool                        `json:"approved,omitempty"`
		Status          EventDataStatus             `json:"status,omitempty"`
		AuthCode        string                      `json:"auth_code,omitempty"`
		ResponseCode    string                      `json:"response_code,omitempty"`
		ResponseSummary string                      `json:"response_summary,omitempty"`
		ThreeDs         *payments.ThreeDsEnrollment `json:"3ds,omitempty"`
		Source          *abc.ResponseCardSource     `json:"source,omitempty"`
		Customer        *common.CustomerResponse    `json:"customer,omitempty"`
		ProcessedOn     *time.Time                  `json:"processed_on,omitempty"`
		Reference       string                      `json:"reference,omitempty"`
		Metadata        map[string]interface{}      `json:"metadata,omitempty"`
	}

	EventResponse struct {
		HttpResponse  common.HttpMetadata
		Id            string                             `json:"id,omitempty"`
		Type          string                             `json:"type,omitempty"`
		Version       string                             `json:"version,omitempty"`
		CreatedOn     string                             `json:"created_on,omitempty"`
		Data          *EventPaymentData                  `json:"data,omitempty"`
		Notifications []EventNotificationSummaryResponse `json:"notifications,omitempty"`
		Links         map[string]common.Link             `json:"_links"`
	}
)

type (
	AttemptSummary struct {
		StatusCode   int    `json:"status_code,omitempty"`
		ResponseBody string `json:"response_body,omitempty"`
		SendMode     string `json:"send_mode,omitempty"`
		Timestamp    string `json:"timestamp,omitempty"`
	}

	EventNotificationResponse struct {
		HttpResponse common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Url          string                 `json:"url,omitempty"`
		Success      bool                   `json:"success,omitempty"`
		ContentType  string                 `json:"content_type,omitempty"`
		Attempts     []AttemptSummary       `json:"attempts,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}
)

func (e *EventTypesResponse) UnmarshalJSON(data []byte) error {
	var eventTypes []EventTypes
	if err := json.Unmarshal(data, &eventTypes); err != nil {
		return err
	}
	e.EventTypes = eventTypes
	return nil
}
