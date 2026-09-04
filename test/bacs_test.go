package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/apm/bacs"
	"github.com/checkout/checkout-sdk-go/v3/common"
)

func TestSendBacsNotification(t *testing.T) {
	t.Skip("Requires a merchant enabled for Bacs Direct Debit and an existing Bacs instrument")

	collectionDate := common.APIShortDate(time.Now().AddDate(0, 0, 14))

	request := bacs.NotificationRequest{
		SourceId:          "src_wmlfc3zyhqzehihu7giusaaawu",
		NotificationType:  bacs.AdvanceNotice,
		CollectionDate:    &collectionDate,
		Amount:            4999,
		Currency:          common.GBP,
		CustomerEmail:     "customer@example.com",
		BillingDescriptor: "CHECKOUT",
		SupportEmail:      "support@test.com",
	}

	response, err := DefaultApi().Bacs.SendNotification(request)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.EventId)
}
