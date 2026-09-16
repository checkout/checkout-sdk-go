package contexts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// Serialization tests for the payment-contexts format: date fields.
//
// PaymentContextsCustomerSummary.{registration_date,first_transaction_date,last_payment_date},
// PaymentContextsTicket.issue_date, PaymentContextsPassenger.date_of_birth and
// PaymentContextsFlightLegDetails.departure_date are all declared
// "type": "string", "format": "date" in the specification.
//
// They were *time.Time, which encoding/json always renders as RFC 3339. These tests pin the
// yyyy-MM-dd wire format and fail if any field is reverted.

func shortDate(t *testing.T, y int, m time.Month, d, hour, min int) *common.APIShortDate {
	t.Helper()
	date := common.APIShortDate(time.Date(y, m, d, hour, min, 0, 0, time.UTC))
	return &date
}

// The times are deliberately not midnight, to prove truncation rather than luck.
func TestPaymentContextsCustomerSummarySerializesShortDates(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsCustomerSummary{
		RegistrationDate:     shortDate(t, 2023, time.May, 1, 13, 59),
		FirstTransactionDate: shortDate(t, 2023, time.July, 1, 1, 2),
		LastPaymentDate:      shortDate(t, 2023, time.August, 1, 23, 59),
		TotalOrderCount:      7,
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "2023-05-01", body["registration_date"])
	assert.Equal(t, "2023-07-01", body["first_transaction_date"])
	assert.Equal(t, "2023-08-01", body["last_payment_date"])
	assert.Equal(t, float64(7), body["total_order_count"])
	assert.NotContains(t, string(raw), "T00:00:00")
}

func TestPaymentContextsCustomerSummaryOmitsDatesWhenUnset(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsCustomerSummary{TotalOrderCount: 1})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	for _, key := range []string{"registration_date", "first_transaction_date", "last_payment_date"} {
		_, present := body[key]
		assert.False(t, present, "%s must be omitted when unset", key)
	}
}

func TestPaymentContextsAirlineSerializesShortDates(t *testing.T) {
	ticket, err := json.Marshal(PaymentContextsTicket{
		Number:    "045-21351455613",
		IssueDate: shortDate(t, 2023, time.May, 20, 8, 15),
	})
	assert.Nil(t, err)
	assert.Contains(t, string(ticket), `"issue_date":"2023-05-20"`)

	passenger, err := json.Marshal(PaymentContextsPassenger{
		FirstName:   "John",
		DateOfBirth: shortDate(t, 1990, time.May, 26, 17, 5),
	})
	assert.Nil(t, err)
	assert.Contains(t, string(passenger), `"date_of_birth":"1990-05-26"`)

	leg, err := json.Marshal(PaymentContextsFlightLegDetails{
		FlightNumber:  "123456",
		DepartureDate: shortDate(t, 2023, time.June, 19, 6, 40),
	})
	assert.Nil(t, err)
	assert.Contains(t, string(leg), `"departure_date":"2023-06-19"`)
}

// The API returns these fields date-only. While they were *time.Time this errored.
func TestPaymentContextsCustomerSummaryDeserializesDateOnlyValues(t *testing.T) {
	payload := `{
		"registration_date":"2023-05-01",
		"first_transaction_date":"2023-07-01",
		"last_payment_date":"2023-08-01"
	}`

	var summary PaymentContextsCustomerSummary
	assert.Nil(t, json.Unmarshal([]byte(payload), &summary))

	assert.NotNil(t, summary.RegistrationDate)
	assert.Equal(t, time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC), time.Time(*summary.RegistrationDate))
	assert.NotNil(t, summary.FirstTransactionDate)
	assert.Equal(t, time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC), time.Time(*summary.FirstTransactionDate))
	assert.NotNil(t, summary.LastPaymentDate)
	assert.Equal(t, time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC), time.Time(*summary.LastPaymentDate))
}
