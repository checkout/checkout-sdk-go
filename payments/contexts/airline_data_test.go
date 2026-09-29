package contexts

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
)

// TestPaymentContextsAirlineData_TicketIsAnObject pins the cardinality of
// airline_data[].ticket. It was a slice, so the SDK sent ticket as an array.
func TestPaymentContextsAirlineData_TicketIsAnObject(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsAirlineData{
		Ticket: &PaymentContextsTicket{
			Number:                 "045-21351455613",
			TravelPackageIndicator: "B",
		},
		Passenger: []PaymentContextsPassenger{{FirstName: "John"}},
		FlightLegDetails: []PaymentContextsFlightLegDetails{{
			FlightNumber:      "101",
			ClassOfTravelling: "J",
			StopOverCode:      "x",
		}},
	})
	assert.Nil(t, err)

	body := string(raw)
	assert.Contains(t, body, `"ticket":{`)
	assert.NotContains(t, body, `"ticket":[`)
	// One passenger serializes as an object; see MarshalJSON.
	assert.Contains(t, body, `"passenger":{`)
	assert.Contains(t, body, `"class_of_travelling":"J"`)
	assert.Contains(t, body, `"stop_over_code":"x"`)
	assert.Contains(t, body, `"flight_number":"101"`)
}

func TestPaymentContextsAirlineData_UnmarshalTicketFromObject(t *testing.T) {
	payload := `{
		"ticket": {"number": "045-21351455613", "issue_date": "2023-05-20"},
		"passenger": [{"first_name": "John"}]
	}`

	var airline PaymentContextsAirlineData
	assert.Nil(t, json.Unmarshal([]byte(payload), &airline))

	assert.NotNil(t, airline.Ticket)
	assert.Equal(t, "045-21351455613", airline.Ticket.Number)
	assert.Len(t, airline.Passenger, 1)
	assert.Equal(t, "John", airline.Passenger[0].FirstName)
}

// PayPal is a payment contexts payment method and returns passenger as a bare object.
func TestPaymentContextsAirlineData_UnmarshalPassengerSingleObject(t *testing.T) {
	payload := `{
		"ticket": {"number": "045"},
		"passenger": {"first_name": "John", "date_of_birth": "1990-05-26", "address": {"country": "US"}}
	}`

	var airline PaymentContextsAirlineData
	assert.Nil(t, json.Unmarshal([]byte(payload), &airline))

	assert.Len(t, airline.Passenger, 1)
	assert.Equal(t, "John", airline.Passenger[0].FirstName)
	assert.NotNil(t, airline.Passenger[0].Address)
	assert.Equal(t, common.US, airline.Passenger[0].Address.Country)
}

func TestPaymentContextsAirlineData_UnmarshalPassengerVariants(t *testing.T) {
	cases := []struct {
		name     string
		payload  string
		expected int
	}{
		{"null passenger", `{"passenger":null}`, 0},
		{"absent passenger", `{"ticket":{"number":"045"}}`, 0},
		{"empty array", `{"passenger":[]}`, 0},
		{"two passengers", `{"passenger":[{"first_name":"John"},{"first_name":"Jane"}]}`, 2},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var airline PaymentContextsAirlineData
			assert.Nil(t, json.Unmarshal([]byte(tc.payload), &airline))
			assert.Len(t, airline.Passenger, tc.expected)
		})
	}
}

// POST /payment-contexts rejects the array form of passenger and requires a single object.
func TestPaymentContextsAirlineData_MarshalPassengerCardinality(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsAirlineData{
		Passenger: []PaymentContextsPassenger{{FirstName: "John"}},
	})
	assert.Nil(t, err)
	assert.Contains(t, string(raw), `"passenger":{`)
	assert.NotContains(t, string(raw), `"passenger":[`)

	raw, err = json.Marshal(PaymentContextsAirlineData{
		Passenger: []PaymentContextsPassenger{{FirstName: "John"}, {FirstName: "Jane"}},
	})
	assert.Nil(t, err)
	assert.Contains(t, string(raw), `"passenger":[{`)
}

// An empty slice and a nil slice must both drop the key: the API rejects "passenger":[] and
// "passenger":null with processing_airline_data_0_passenger_invalid, and only absence is accepted.
func TestPaymentContextsAirlineData_MarshalOmitsPassengerWhenEmpty(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsAirlineData{
		Ticket:    &PaymentContextsTicket{Number: "045"},
		Passenger: []PaymentContextsPassenger{},
	})
	assert.Nil(t, err)
	assert.NotContains(t, string(raw), "passenger")

	raw, err = json.Marshal(PaymentContextsAirlineData{
		Ticket: &PaymentContextsTicket{Number: "045"},
	})
	assert.Nil(t, err)
	assert.NotContains(t, string(raw), "passenger")
}

// Field order follows the specification: ticket, passenger, flight_leg_details.
func TestPaymentContextsAirlineData_MarshalFieldOrder(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsAirlineData{
		Ticket:           &PaymentContextsTicket{Number: "045"},
		Passenger:        []PaymentContextsPassenger{{FirstName: "John"}},
		FlightLegDetails: []PaymentContextsFlightLegDetails{{FlightNumber: "101"}},
	})
	assert.Nil(t, err)

	body := string(raw)
	assert.True(t, strings.Index(body, `"ticket"`) < strings.Index(body, `"passenger"`), body)
	assert.True(t, strings.Index(body, `"passenger"`) < strings.Index(body, `"flight_leg_details"`), body)
}

// TestPaymentContextsPassenger_AddressIsNarrow pins the narrowed address: the specification
// defines exactly one property on passenger.address.
func TestPaymentContextsPassenger_AddressIsNarrow(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsPassenger{
		FirstName: "John",
		Address:   &payments.PassengerAddress{Country: common.US},
	})
	assert.Nil(t, err)

	body := string(raw)
	assert.Contains(t, body, `"address":{"country":"US"}`)
	assert.NotContains(t, body, "address_line1")
	assert.NotContains(t, body, `"zip"`)
}

// TestPaymentContextsProcessing_SpecFields covers the four fields that were missing from the
// struct: discount_amount, tax_amount, custom_payment_method_ids and accommodation_data.
func TestPaymentContextsProcessing_SpecFields(t *testing.T) {
	payload := `{
		"discount_amount": 5,
		"shipping_amount": 300,
		"tax_amount": 3000,
		"invoice_id": "INV-1",
		"brand_name": "Acme Corporation",
		"locale": "en-US",
		"custom_payment_method_ids": ["cpm_001", "cpm_002"],
		"airline_data": [{"ticket": {"number": "045"}}],
		"accommodation_data": [{
			"name": "The Sea View Hotel",
			"state": "FL",
			"country": "USA",
			"room": [{"rate": "70", "number_of_nights_at_room_rate": "3"}]
		}]
	}`

	var processing PaymentContextsProcessing
	assert.Nil(t, json.Unmarshal([]byte(payload), &processing))

	assert.Equal(t, 5, processing.DiscountAmount)
	assert.Equal(t, 300, processing.ShippingAmount)
	assert.Equal(t, 3000, processing.TaxAmount)
	assert.Equal(t, "INV-1", processing.InvoiceId)
	assert.Equal(t, "Acme Corporation", processing.BrandName)
	assert.Equal(t, "en-US", processing.Locale)
	assert.Equal(t, []string{"cpm_001", "cpm_002"}, processing.CustomPaymentMethodIds)

	assert.Len(t, processing.AirlineData, 1)
	assert.Equal(t, "045", processing.AirlineData[0].Ticket.Number)

	// accommodation_data uses the shared payments.AccommodationData, because payment contexts,
	// POST /payments and the GET /payments/{id} response all resolve it to one schema.
	assert.Len(t, processing.AccommodationData, 1)
	assert.Equal(t, "The Sea View Hotel", processing.AccommodationData[0].Name)
	assert.Equal(t, "FL", processing.AccommodationData[0].State)
	assert.Equal(t, "USA", processing.AccommodationData[0].Country)
	assert.Equal(t, "3", processing.AccommodationData[0].Room[0].NumberOfNightsAtRoomRate)
}

func TestPaymentContextsProcessing_MarshalsSpecKeyNames(t *testing.T) {
	raw, err := json.Marshal(PaymentContextsProcessing{
		DiscountAmount:         5,
		TaxAmount:              3000,
		CustomPaymentMethodIds: []string{"cpm_001"},
		AccommodationData: []payments.AccommodationData{{
			Name:    "The Sea View Hotel",
			State:   "FL",
			Country: "USA",
		}},
	})
	assert.Nil(t, err)

	body := string(raw)
	assert.Contains(t, body, `"discount_amount":5`)
	assert.Contains(t, body, `"tax_amount":3000`)
	assert.Contains(t, body, `"custom_payment_method_ids":["cpm_001"]`)
	assert.Contains(t, body, `"accommodation_data":[{`)
	assert.Contains(t, body, `"state":"FL"`)
	assert.Contains(t, body, `"country":"USA"`)
}
