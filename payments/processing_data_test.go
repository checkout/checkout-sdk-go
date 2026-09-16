package payments

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Verifies the fields recently aligned with the Checkout.com swagger spec
// deserialize correctly from the API wire format.
//   - scheme
//   - partner_fraud_status
//   - partner_merchant_advice_code
//   - accommodation_data (incl. its format: date fields)
//   - airline_data
//   - failure_code
//   - partner_code
//   - partner_response_code
//   - fallback_source_used
//   - scheme_transaction_link_id (Mastercard Transaction Link Identifier)
//   - recommendation_code (Mastercard/Visa recommendation code)
func TestProcessingData_UnmarshalAllNewFields(t *testing.T) {
	payload := `{
		"scheme":"ACCEL",
		"partner_fraud_status":"Accepted",
		"partner_merchant_advice_code":"24",
		"failure_code":"partner_error",
		"partner_code":"902111",
		"partner_response_code":"DECLINED",
		"fallback_source_used":true,
		"scheme_transaction_link_id":"MTL-XYZ-789",
		"recommendation_code":"02",
		"accommodation_data":[{
			"name":"Grand Hotel",
			"check_in_date":"2026-10-01",
			"check_out_date":"2026-10-05",
			"guests":[{"first_name":"Jane","date_of_birth":"1985-07-14"}]
		}],
		"airline_data":[{"ticket":{"number":"045-21351455613","issue_date":"2026-09-20"}}]
	}`

	var data ProcessingData
	err := json.Unmarshal([]byte(payload), &data)

	assert.NoError(t, err)
	assert.Equal(t, "ACCEL", data.Scheme)
	assert.Equal(t, "Accepted", data.PartnerFraudStatus)
	assert.Equal(t, "24", data.PartnerMerchantAdviceCode)
	assert.Equal(t, "partner_error", data.FailureCode)
	assert.Equal(t, "902111", data.PartnerCode)
	assert.Equal(t, "DECLINED", data.PartnerResponseCode)
	assert.True(t, data.FallbackSourceUsed)
	assert.Equal(t, "MTL-XYZ-789", data.SchemeTransactionLinkId)
	assert.Equal(t, "02", data.RecommendationCode)

	assert.Len(t, data.AccommodationData, 1)
	assert.Equal(t, "Grand Hotel", data.AccommodationData[0].Name)

	// The spec declares check_in_date, check_out_date and guests[].date_of_birth as
	// format: date, so the API sends them date-only. Before these fields became
	// common.APIShortDate they were *time.Time, and encoding/json could not parse
	// "2026-10-01" into one -- the whole response failed with
	// `parsing time "2026-10-01" as "2006-01-02T15:04:05Z07:00"`.
	assert.NotNil(t, data.AccommodationData[0].CheckInDate)
	assert.Equal(t,
		time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC),
		time.Time(*data.AccommodationData[0].CheckInDate))
	assert.NotNil(t, data.AccommodationData[0].CheckOutDate)
	assert.Equal(t,
		time.Date(2026, 10, 5, 0, 0, 0, 0, time.UTC),
		time.Time(*data.AccommodationData[0].CheckOutDate))

	assert.Len(t, data.AccommodationData[0].Guests, 1)
	assert.Equal(t, "Jane", data.AccommodationData[0].Guests[0].FirstName)
	assert.NotNil(t, data.AccommodationData[0].Guests[0].DateOfBirth)
	assert.Equal(t,
		time.Date(1985, 7, 14, 0, 0, 0, 0, time.UTC),
		time.Time(*data.AccommodationData[0].Guests[0].DateOfBirth))

	assert.Len(t, data.AirlineData, 1)
	assert.NotNil(t, data.AirlineData[0].Ticket)
	assert.Equal(t, "045-21351455613", data.AirlineData[0].Ticket.Number)
	assert.NotNil(t, data.AirlineData[0].Ticket.IssueDate)
	assert.Equal(t,
		time.Date(2026, 9, 20, 0, 0, 0, 0, time.UTC),
		time.Time(*data.AirlineData[0].Ticket.IssueDate))
}

// Verifies the new fields stay zero-valued (no spurious defaults) when absent
// from the API response.
func TestProcessingData_LeavesNewFieldsZeroWhenAbsent(t *testing.T) {
	payload := `{"locale":"en-GB"}`

	var data ProcessingData
	err := json.Unmarshal([]byte(payload), &data)

	assert.NoError(t, err)
	assert.Equal(t, "en-GB", data.Locale)
	assert.Empty(t, data.Scheme)
	assert.Empty(t, data.PartnerFraudStatus)
	assert.Empty(t, data.PartnerMerchantAdviceCode)
	assert.Empty(t, data.FailureCode)
	assert.Empty(t, data.PartnerCode)
	assert.Empty(t, data.PartnerResponseCode)
	assert.False(t, data.FallbackSourceUsed)
	assert.Empty(t, data.SchemeTransactionLinkId)
	assert.Empty(t, data.RecommendationCode)
	assert.Nil(t, data.AccommodationData)
	assert.Nil(t, data.AirlineData)
}

// Verifies the Mastercard Transaction Link Identifier deserializes from the
// processing object returned in a payment response (PaymentProcessing), and
// serializes back to its snake_case wire name.
func TestPaymentProcessing_SchemeTransactionLinkId(t *testing.T) {
	payload := `{
		"retrieval_reference_number":"RRN001",
		"acquirer_transaction_id":"ACQ001",
		"scheme":"Mastercard",
		"scheme_transaction_link_id":"MTL-XYZ-789"
	}`

	var processing PaymentProcessing
	err := json.Unmarshal([]byte(payload), &processing)

	assert.NoError(t, err)
	assert.Equal(t, "RRN001", processing.RetrievalReferenceNumber)
	assert.Equal(t, "ACQ001", processing.AcquirerTransactionId)
	assert.Equal(t, "Mastercard", processing.Scheme)
	assert.Equal(t, "MTL-XYZ-789", processing.SchemeTransactionLinkId)

	marshalled, err := json.Marshal(PaymentProcessing{SchemeTransactionLinkId: "MTL-001"})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"scheme_transaction_link_id":"MTL-001"`)
}

// Verifies the field stays zero-valued when absent from the payment response
// processing object (it is optional and only populated for Mastercard transactions).
func TestPaymentProcessing_SchemeTransactionLinkIdAbsent(t *testing.T) {
	payload := `{"retrieval_reference_number":"RRN001"}`

	var processing PaymentProcessing
	err := json.Unmarshal([]byte(payload), &processing)

	assert.NoError(t, err)
	assert.Equal(t, "RRN001", processing.RetrievalReferenceNumber)
	assert.Empty(t, processing.SchemeTransactionLinkId)
}

// Verifies scheme_transaction_link_id on the request-side ProcessingSettings
// (PaymentRequestProcessing) serializes to its snake_case wire name and is
// omitted when unset.
func TestProcessingSettings_SchemeTransactionLinkId(t *testing.T) {
	marshalled, err := json.Marshal(ProcessingSettings{SchemeTransactionLinkId: "MTL-001"})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"scheme_transaction_link_id":"MTL-001"`)

	marshalled, err = json.Marshal(ProcessingSettings{})
	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "scheme_transaction_link_id")
}
