package payments

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// Serialization tests for the processing.accommodation_data and processing.airline_data
// date fields.
//
// The specification declares check_in_date, check_out_date, guests[].date_of_birth,
// ticket.issue_date, passenger.date_of_birth and flight_leg_details[].departure_date as
// "type": "string", "format": "date" -- so the wire value is yyyy-MM-dd.
//
// These fields used to be *time.Time. Go's encoding/json marshals time.Time as RFC 3339
// and offers no per-field override, so the SDK always emitted a full timestamp
// ("2026-10-01T00:00:00Z"). Tamara rejects that and the merchant saw a gateway 500 on
// Hosted Payments Pages. common.APIShortDate marshals yyyy-MM-dd instead.
//
// Every test below fails if a field is reverted to *time.Time.

func shortDate(t *testing.T, y int, m time.Month, d, hour, min int) *common.APIShortDate {
	t.Helper()
	date := common.APIShortDate(time.Date(y, m, d, hour, min, 0, 0, time.UTC))
	return &date
}

// The times passed in are deliberately NOT midnight: the assertion proves the time
// component is truncated, rather than passing by luck because the caller happened to
// supply a zero time.
func TestAccommodationDataSerializesDatesAsShortDates(t *testing.T) {
	raw, err := json.Marshal(AccommodationData{
		Name:             "Grand Hotel",
		BookingReference: "HOTEL123",
		CheckInDate:      shortDate(t, 2026, time.October, 1, 13, 45),
		CheckOutDate:     shortDate(t, 2026, time.October, 5, 11, 30),
		Guests: []Guest{{
			FirstName:   "Jane",
			LastName:    "Doe",
			DateOfBirth: shortDate(t, 1985, time.July, 14, 23, 59),
		}},
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "Grand Hotel", body["name"])
	assert.Equal(t, "HOTEL123", body["booking_reference"])
	assert.Equal(t, "2026-10-01", body["check_in_date"])
	assert.Equal(t, "2026-10-05", body["check_out_date"])

	guests, ok := body["guests"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, guests, 1)
	assert.Equal(t, "1985-07-14", guests[0].(map[string]interface{})["date_of_birth"])
}

// Guards the omitempty half of the change: *common.APIShortDate + omitempty omits the key
// when nil. A non-pointer APIShortDate would always serialize, because encoding/json never
// treats a struct value as empty.
func TestAccommodationDataOmitsDatesWhenUnset(t *testing.T) {
	raw, err := json.Marshal(AccommodationData{Name: "Grand Hotel"})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "Grand Hotel", body["name"])
	for _, key := range []string{"check_in_date", "check_out_date", "guests"} {
		_, present := body[key]
		assert.False(t, present, "%s must be omitted when unset", key)
	}
	assert.Len(t, body, 1)
}

// The airline_data dates share the same format: date declaration and the same defect.
func TestAirlineDataSerializesDatesAsShortDates(t *testing.T) {
	raw, err := json.Marshal(AirlineData{
		Ticket: &Ticket{
			Number:    "045-21351455613",
			IssueDate: shortDate(t, 2026, time.September, 20, 8, 15),
		},
		Passenger: &Passenger{
			FirstName:   "John",
			DateOfBirth: shortDate(t, 1990, time.May, 26, 17, 5),
		},
		FlightLegDetails: []FlightLegDetails{{
			FlightNumber:  "123456",
			DepartureDate: shortDate(t, 2026, time.June, 19, 6, 40),
		}},
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "2026-09-20", body["ticket"].(map[string]interface{})["issue_date"])
	assert.Equal(t, "1990-05-26", body["passenger"].(map[string]interface{})["date_of_birth"])

	legs, ok := body["flight_leg_details"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, legs, 1)
	assert.Equal(t, "2026-06-19", legs[0].(map[string]interface{})["departure_date"])
}

// The merchant's reported payload: processing.accommodation_data on a request body.
// ProcessingSettings is the struct that Hosted Payments Pages, Payment Links and Payment
// Sessions all embed as "processing", so this covers every request surface at once.
func TestProcessingSettingsSerializesAccommodationDatesAsShortDates(t *testing.T) {
	raw, err := json.Marshal(ProcessingSettings{
		AccommodationData: []AccommodationData{{
			Name:         "Grand Hotel",
			CheckInDate:  shortDate(t, 2026, time.October, 1, 0, 0),
			CheckOutDate: shortDate(t, 2026, time.October, 5, 0, 0),
			Guests: []Guest{{
				FirstName:   "Jane",
				DateOfBirth: shortDate(t, 1985, time.July, 14, 0, 0),
			}},
		}},
	})
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	accommodation, ok := body["accommodation_data"].([]interface{})
	assert.True(t, ok)
	assert.Len(t, accommodation, 1)

	first := accommodation[0].(map[string]interface{})
	assert.Equal(t, "2026-10-01", first["check_in_date"])
	assert.Equal(t, "2026-10-05", first["check_out_date"])

	guests := first["guests"].([]interface{})
	assert.Equal(t, "1985-07-14", guests[0].(map[string]interface{})["date_of_birth"])

	// Nothing in the payload may carry a time component.
	assert.NotContains(t, string(raw), "T00:00:00")
}

// The response side of the same schema. AccommodationData is reused on ProcessingData,
// which hangs off the payment-details response, and the API returns these fields
// date-only. While they were *time.Time this failed outright with
// `parsing time "2026-10-01" as "2006-01-02T15:04:05Z07:00": cannot parse "" as "T"`,
// so GetPaymentDetails broke for any lodging payment carrying accommodation_data --
// whether or not the merchant ever sent the field.
func TestAccommodationDataDeserializesDateOnlyValues(t *testing.T) {
	payload := `{
		"name":"Grand Hotel",
		"check_in_date":"2026-10-01",
		"check_out_date":"2026-10-05",
		"guests":[{"first_name":"Jane","date_of_birth":"1985-07-14"}]
	}`

	var data AccommodationData
	assert.Nil(t, json.Unmarshal([]byte(payload), &data))

	assert.Equal(t, "Grand Hotel", data.Name)
	assert.NotNil(t, data.CheckInDate)
	assert.Equal(t, time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC), time.Time(*data.CheckInDate))
	assert.NotNil(t, data.CheckOutDate)
	assert.Equal(t, time.Date(2026, 10, 5, 0, 0, 0, 0, time.UTC), time.Time(*data.CheckOutDate))

	assert.Len(t, data.Guests, 1)
	assert.NotNil(t, data.Guests[0].DateOfBirth)
	assert.Equal(t, time.Date(1985, 7, 14, 0, 0, 0, 0, time.UTC), time.Time(*data.Guests[0].DateOfBirth))
}

// A date-only round trip must be lossless: what the API sends is what the SDK sends back.
func TestAccommodationDataDateRoundTrip(t *testing.T) {
	payload := `{"check_in_date":"2026-10-01","check_out_date":"2026-10-05"}`

	var data AccommodationData
	assert.Nil(t, json.Unmarshal([]byte(payload), &data))

	raw, err := json.Marshal(data)
	assert.Nil(t, err)
	assert.JSONEq(t, payload, string(raw))
}

// A2: common.APIShortDate also accepts the compact yyyyMMdd form, matching the Java SDK's
// LocalDate deserializer. This confirms the tolerance reaches the migrated fields, not just the
// type in isolation -- a response using the compact form parses, and re-serializes in the
// canonical yyyy-MM-dd the specification declares.
func TestAccommodationDataAcceptsCompactDates(t *testing.T) {
	payload := `{
		"check_in_date":"20261001",
		"check_out_date":"20261005",
		"guests":[{"first_name":"Jane","date_of_birth":"19850714"}]
	}`

	var data AccommodationData
	assert.Nil(t, json.Unmarshal([]byte(payload), &data))

	assert.Equal(t, time.Date(2026, 10, 1, 0, 0, 0, 0, time.UTC), time.Time(*data.CheckInDate))
	assert.Equal(t, time.Date(2026, 10, 5, 0, 0, 0, 0, time.UTC), time.Time(*data.CheckOutDate))
	assert.Equal(t, time.Date(1985, 7, 14, 0, 0, 0, 0, time.UTC), time.Time(*data.Guests[0].DateOfBirth))

	raw, err := json.Marshal(data)
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))
	assert.Equal(t, "2026-10-01", body["check_in_date"])
	assert.Equal(t, "2026-10-05", body["check_out_date"])
}

// The rejection still holds on a migrated field: a date-time on a format: date field is a
// contract break and must surface as an error rather than be silently truncated.
func TestAccommodationDataRejectsDateTimeValues(t *testing.T) {
	var data AccommodationData
	err := json.Unmarshal([]byte(`{"check_in_date":"2026-10-01T00:00:00Z"}`), &data)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "APIShortDate only accepts")
}
