package bacs

import "github.com/checkout/checkout-sdk-go/v3/common"

const (
	apmsPath          = "apms"
	bacsPath          = "bacs"
	notificationsPath = "notifications"
)

// NotificationType is the type of pre-notification being sent to the payer.
type NotificationType string

const (
	AdvanceNotice NotificationType = "advance_notice"
)

type (
	// NotificationRequest is a Bacs Direct Debit pre-notification request.
	//
	// SourceId matches the pattern ^(src)_(\w{26})$. CollectionDate is a yyyy-MM-dd date. Amount is
	// in the currency's minor unit with a minimum of 1. Currency is min 3 max 3 characters. Reference
	// is max 50 and BillingDescriptor max 25 characters. CustomerEmail and SupportEmail are email
	// addresses, and SupportPhone is in E.164 format. Reference and SupportPhone are the only
	// optional properties, and are the only two tagged omitempty: the eight required fields always
	// serialize, so a forgotten one surfaces as an explicit zero value rather than a silently
	// absent key.
	NotificationRequest struct {
		SourceId          string               `json:"source_id" binding:"required"`
		NotificationType  NotificationType     `json:"notification_type" binding:"required"`
		CollectionDate    *common.APIShortDate `json:"collection_date" binding:"required"`
		Amount            int64                `json:"amount" binding:"required"`
		Currency          common.Currency      `json:"currency" binding:"required"`
		CustomerEmail     string               `json:"customer_email" binding:"required"`
		BillingDescriptor string               `json:"billing_descriptor" binding:"required"`
		SupportEmail      string               `json:"support_email" binding:"required"`
		Reference         string               `json:"reference,omitempty"`
		SupportPhone      string               `json:"support_phone,omitempty"`
	}

	// NotificationResponse is the Bacs Direct Debit pre-notification response.
	NotificationResponse struct {
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		EventId      string              `json:"event_id,omitempty"`
	}
)
