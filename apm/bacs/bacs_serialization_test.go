package bacs

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// Schema validation tests for NotificationRequest against the swagger schema of
// POST /apms/bacs/notifications. Covers all 10 properties.

func TestNotificationRequestSerializesEveryProperty(t *testing.T) {
	collectionDate := common.APIShortDate(time.Date(2026, 7, 15, 0, 0, 0, 0, time.UTC))

	raw, err := json.Marshal(NotificationRequest{
		SourceId:          "src_wmlfc3zyhqzehihu7giusaaawu",
		NotificationType:  AdvanceNotice,
		CollectionDate:    &collectionDate,
		Amount:            4999,
		Currency:          common.GBP,
		Reference:         "INV-12345",
		CustomerEmail:     "customer@example.com",
		BillingDescriptor: "CHECKOUT",
		SupportEmail:      "support@test.com",
		SupportPhone:      "+447700900123",
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "src_wmlfc3zyhqzehihu7giusaaawu", body["source_id"])
	assert.Equal(t, "advance_notice", body["notification_type"])
	assert.Equal(t, "2026-07-15", body["collection_date"])
	assert.Equal(t, float64(4999), body["amount"])
	assert.Equal(t, "GBP", body["currency"])
	assert.Equal(t, "INV-12345", body["reference"])
	assert.Equal(t, "customer@example.com", body["customer_email"])
	assert.Equal(t, "CHECKOUT", body["billing_descriptor"])
	assert.Equal(t, "support@test.com", body["support_email"])
	assert.Equal(t, "+447700900123", body["support_phone"])
	assert.Len(t, body, 10)
}

// The eight required fields carry binding:"required" and no omitempty, matching the SDK convention,
// so a forgotten one serializes as an explicit zero rather than a silently absent key. The two
// optional fields are the only ones omitted when unset.
func TestNotificationRequestAlwaysSerializesTheRequiredFields(t *testing.T) {
	raw, err := json.Marshal(NotificationRequest{})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	for _, key := range []string{
		"source_id", "notification_type", "collection_date", "amount",
		"currency", "customer_email", "billing_descriptor", "support_email",
	} {
		_, present := body[key]
		assert.True(t, present, "required field %s must always serialize", key)
	}

	_, hasReference := body["reference"]
	_, hasSupportPhone := body["support_phone"]
	assert.False(t, hasReference)
	assert.False(t, hasSupportPhone)
	assert.Len(t, body, 8)
}

func TestCollectionDateSerializesAsAShortDate(t *testing.T) {
	// The specification declares collection_date as format: date. common.APIShortDate marshals
	// yyyy-MM-dd; a plain time.Time would emit a full RFC 3339 timestamp.
	collectionDate := common.APIShortDate(time.Date(2026, 7, 15, 13, 45, 0, 0, time.UTC))

	raw, err := json.Marshal(NotificationRequest{CollectionDate: &collectionDate})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))
	assert.Equal(t, "2026-07-15", body["collection_date"])
}

func TestNotificationTypeCarriesTheSingleDeclaredValue(t *testing.T) {
	assert.Equal(t, "advance_notice", string(AdvanceNotice))
}
