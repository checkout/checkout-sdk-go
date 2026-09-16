package setups

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
)

// Serialization tests for the Payment Setups format: date fields.
//
// PaymentSetupMerchantAccount.{registration_date,last_modified,first_transaction_date,
// last_transaction_date}, PaymentSetupSubMerchant.registration_date and the
// PaymentSetupAccountFundingTransaction sender/recipient date_of_birth are all declared
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
func TestCustomerMerchantAccountSerializesShortDates(t *testing.T) {
	raw, err := json.Marshal(CustomerMerchantAccount{
		Id:                   "acct_1",
		RegistrationDate:     shortDate(t, 2023, time.May, 1, 13, 59),
		LastModified:         shortDate(t, 2023, time.May, 1, 8, 30),
		FirstTransactionDate: shortDate(t, 2023, time.September, 15, 1, 2),
		LastTransactionDate:  shortDate(t, 2025, time.March, 28, 23, 59),
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "acct_1", body["id"])
	assert.Equal(t, "2023-05-01", body["registration_date"])
	assert.Equal(t, "2023-05-01", body["last_modified"])
	assert.Equal(t, "2023-09-15", body["first_transaction_date"])
	assert.Equal(t, "2025-03-28", body["last_transaction_date"])
	assert.NotContains(t, string(raw), "T00:00:00")
}

func TestCustomerMerchantAccountOmitsDatesWhenUnset(t *testing.T) {
	raw, err := json.Marshal(CustomerMerchantAccount{Id: "acct_1"})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	for _, key := range []string{
		"registration_date", "last_modified", "first_transaction_date", "last_transaction_date",
	} {
		_, present := body[key]
		assert.False(t, present, "%s must be omitted when unset", key)
	}
	assert.Len(t, body, 1)
}

func TestOrderSubMerchantSerializesRegistrationDateAsShortDate(t *testing.T) {
	raw, err := json.Marshal(OrderSubMerchant{
		Id:               "sub_1",
		RegistrationDate: shortDate(t, 2023, time.January, 15, 12, 0),
	})
	assert.Nil(t, err)
	assert.Contains(t, string(raw), `"registration_date":"2023-01-15"`)
}

func TestAccountFundingTransactionSerializesDateOfBirthAsShortDate(t *testing.T) {
	sender, err := json.Marshal(AccountFundingTransactionSender{
		Reference:   "REF-1",
		DateOfBirth: shortDate(t, 2000, time.January, 1, 6, 30),
	})
	assert.Nil(t, err)
	assert.Contains(t, string(sender), `"date_of_birth":"2000-01-01"`)

	recipient, err := json.Marshal(AccountFundingTransactionRecipient{
		FirstName:   "Jane",
		DateOfBirth: shortDate(t, 2000, time.January, 1, 6, 30),
	})
	assert.Nil(t, err)
	assert.Contains(t, string(recipient), `"date_of_birth":"2000-01-01"`)
}

// PaymentSetupIndustry reuses payments.AccommodationData, so Payment Setups inherits the
// accommodation date fix with no separate declaration of its own. This pins that reuse.
func TestPaymentSetupIndustrySerializesAccommodationDatesAsShortDates(t *testing.T) {
	checkIn := common.APIShortDate(time.Date(2025, 4, 11, 15, 0, 0, 0, time.UTC))
	checkOut := common.APIShortDate(time.Date(2025, 4, 18, 10, 0, 0, 0, time.UTC))

	raw, err := json.Marshal(PaymentSetupIndustry{
		AccommodationData: []payments.AccommodationData{{
			Name:         "Checkout Lodge",
			CheckInDate:  &checkIn,
			CheckOutDate: &checkOut,
		}},
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	accommodation := body["accommodation_data"].([]interface{})
	assert.Len(t, accommodation, 1)

	first := accommodation[0].(map[string]interface{})
	assert.Equal(t, "2025-04-11", first["check_in_date"])
	assert.Equal(t, "2025-04-18", first["check_out_date"])
}

// The API returns these fields date-only. While they were *time.Time this errored.
func TestCustomerMerchantAccountDeserializesDateOnlyValues(t *testing.T) {
	payload := `{
		"id":"acct_1",
		"registration_date":"2023-05-01",
		"last_modified":"2023-05-02",
		"first_transaction_date":"2023-09-15",
		"last_transaction_date":"2025-03-28"
	}`

	var account CustomerMerchantAccount
	assert.Nil(t, json.Unmarshal([]byte(payload), &account))

	assert.Equal(t, "acct_1", account.Id)
	assert.NotNil(t, account.RegistrationDate)
	assert.Equal(t, time.Date(2023, 5, 1, 0, 0, 0, 0, time.UTC), time.Time(*account.RegistrationDate))
	assert.NotNil(t, account.LastModified)
	assert.Equal(t, time.Date(2023, 5, 2, 0, 0, 0, 0, time.UTC), time.Time(*account.LastModified))
	assert.NotNil(t, account.FirstTransactionDate)
	assert.Equal(t, time.Date(2023, 9, 15, 0, 0, 0, 0, time.UTC), time.Time(*account.FirstTransactionDate))
	assert.NotNil(t, account.LastTransactionDate)
	assert.Equal(t, time.Date(2025, 3, 28, 0, 0, 0, 0, time.UTC), time.Time(*account.LastTransactionDate))
}
