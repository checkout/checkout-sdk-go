package payments

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// airlineJSON is the swagger AirlineData example values, shared byte for byte with the equivalent
// test in the other SDKs so that seven languages assert against one wire shape.
const airlineJSON = `{
	"ticket": {
		"number": "045-21351455613",
		"issue_date": "2023-05-20",
		"issuing_carrier_code": "AI",
		"travel_package_indicator": "B",
		"travel_agency_name": "World Tours",
		"travel_agency_code": "01"
	},
	"passenger": [{
		"first_name": "John",
		"last_name": "White",
		"date_of_birth": "1990-05-26",
		"address": {"country": "US"}
	}],
	"flight_leg_details": [{
		"flight_number": "101",
		"carrier_code": "BA",
		"class_of_travelling": "J",
		"departure_airport": "LHR",
		"departure_date": "2023-06-19",
		"departure_time": "15:30",
		"arrival_airport": "LAX",
		"stop_over_code": "x",
		"fare_basis_code": "SPRSVR"
	}]
}`

func shortDateValue(t *testing.T, d *common.APIShortDate) string {
	t.Helper()
	assert.NotNil(t, d)
	return time.Time(*d).Format("2006-01-02")
}

// TestAirlineData_UnmarshalPassengerArray covers the shape the API actually returns.
func TestAirlineData_UnmarshalPassengerArray(t *testing.T) {
	var airline AirlineData
	assert.Nil(t, json.Unmarshal([]byte(airlineJSON), &airline))

	assert.NotNil(t, airline.Ticket)
	assert.Equal(t, "045-21351455613", airline.Ticket.Number)
	assert.Equal(t, "2023-05-20", shortDateValue(t, airline.Ticket.IssueDate))
	assert.Equal(t, "AI", airline.Ticket.IssuingCarrierCode)
	assert.Equal(t, "B", airline.Ticket.TravelPackageIndicator)
	assert.Equal(t, "World Tours", airline.Ticket.TravelAgencyName)
	assert.Equal(t, "01", airline.Ticket.TravelAgencyCode)

	assert.Len(t, airline.Passenger, 1)
	assert.Equal(t, "John", airline.Passenger[0].FirstName)
	assert.Equal(t, "White", airline.Passenger[0].LastName)
	assert.Equal(t, "1990-05-26", shortDateValue(t, airline.Passenger[0].DateOfBirth))
	assert.NotNil(t, airline.Passenger[0].Address)
	assert.Equal(t, common.US, airline.Passenger[0].Address.Country)

	assert.Len(t, airline.FlightLegDetails, 1)
	leg := airline.FlightLegDetails[0]
	assert.Equal(t, "101", leg.FlightNumber)
	assert.Equal(t, "BA", leg.CarrierCode)
	// class_of_travelling, double l. The SDK used to ship class_of_traveling.
	assert.Equal(t, "J", leg.ClassOfTravelling)
	assert.Equal(t, "LHR", leg.DepartureAirport)
	assert.Equal(t, "2023-06-19", shortDateValue(t, leg.DepartureDate))
	assert.Equal(t, "15:30", leg.DepartureTime)
	assert.Equal(t, "LAX", leg.ArrivalAirport)
	// stop_over_code, three tokens. The SDK used to ship stopover_code.
	assert.Equal(t, "x", leg.StopOverCode)
	assert.Equal(t, "SPRSVR", leg.FareBasisCode)
}

// TestAirlineData_UnmarshalPassengerSingleObject covers the shape PayPal sends. The specification
// allows it on PaymentInterfacesProcessingAirlineData with the note "PayPal requires a single
// object".
func TestAirlineData_UnmarshalPassengerSingleObject(t *testing.T) {
	payload := `{
		"ticket": {"number": "045-21351455613"},
		"passenger": {
			"first_name": "John",
			"last_name": "White",
			"date_of_birth": "1990-05-26",
			"address": {"country": "US"}
		}
	}`

	var airline AirlineData
	assert.Nil(t, json.Unmarshal([]byte(payload), &airline))

	// Normalized to a one-element slice, so callers only handle one shape.
	assert.Len(t, airline.Passenger, 1)
	assert.Equal(t, "John", airline.Passenger[0].FirstName)
	assert.Equal(t, "White", airline.Passenger[0].LastName)
	assert.Equal(t, "1990-05-26", shortDateValue(t, airline.Passenger[0].DateOfBirth))
	assert.Equal(t, common.US, airline.Passenger[0].Address.Country)
	assert.Equal(t, "045-21351455613", airline.Ticket.Number)
}

func TestAirlineData_UnmarshalPassengerVariants(t *testing.T) {
	cases := []struct {
		name     string
		payload  string
		expected int
		nilSlice bool
	}{
		{"null passenger", `{"passenger":null}`, 0, true},
		{"absent passenger", `{"ticket":{"number":"045"}}`, 0, true},
		{"empty array", `{"passenger":[]}`, 0, false},
		{"two passengers", `{"passenger":[{"first_name":"John"},{"first_name":"Jane"}]}`, 2, false},
		{"leading whitespace before object", "{\"passenger\": \n\t {\"first_name\":\"John\"}}", 1, false},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var airline AirlineData
			assert.Nil(t, json.Unmarshal([]byte(tc.payload), &airline))
			assert.Len(t, airline.Passenger, tc.expected)
			if tc.nilSlice {
				assert.Nil(t, airline.Passenger)
			} else {
				assert.NotNil(t, airline.Passenger)
			}
		})
	}
}

// TestAirlineData_UnmarshalPropagatesElementErrors guards that the tolerant reader does not
// swallow a genuinely malformed payload.
func TestAirlineData_UnmarshalPropagatesElementErrors(t *testing.T) {
	var airline AirlineData
	err := json.Unmarshal([]byte(`{"passenger":"John"}`), &airline)
	assert.NotNil(t, err)
}

// TestAirlineData_MarshalPassengerCardinality pins the outbound shape against the live API, which
// accepts a single object on every request surface but an array only on POST /payments. See
// AirlineData.MarshalJSON for the sandbox-verified matrix.
func TestAirlineData_MarshalPassengerCardinality(t *testing.T) {
	// One passenger serializes as an object, the shape hosted payments, payment links and
	// payment contexts all require.
	raw, err := json.Marshal(AirlineData{
		Passenger: []Passenger{{FirstName: "John", LastName: "White"}},
	})
	assert.Nil(t, err)
	assert.Contains(t, string(raw), `"passenger":{`)
	assert.NotContains(t, string(raw), `"passenger":[`)

	// Several passengers can only be expressed as an array.
	raw, err = json.Marshal(AirlineData{
		Passenger: []Passenger{{FirstName: "John"}, {FirstName: "Jane"}},
	})
	assert.Nil(t, err)
	assert.Contains(t, string(raw), `"passenger":[{`)

	// No passengers omits the key entirely.
	raw, err = json.Marshal(AirlineData{Ticket: &Ticket{Number: "045"}})
	assert.Nil(t, err)
	assert.NotContains(t, string(raw), "passenger")
}

func TestAirlineData_RoundTrip(t *testing.T) {
	original := AirlineData{
		Ticket: &Ticket{
			Number:                 "045-21351455613",
			TravelPackageIndicator: "B",
		},
		Passenger: []Passenger{
			{FirstName: "John"},
			{FirstName: "Jane"},
		},
		FlightLegDetails: []FlightLegDetails{{
			FlightNumber:      "101",
			ClassOfTravelling: "J",
			StopOverCode:      "x",
		}},
	}

	raw, err := json.Marshal(original)
	assert.Nil(t, err)

	var result AirlineData
	assert.Nil(t, json.Unmarshal(raw, &result))

	assert.Equal(t, "045-21351455613", result.Ticket.Number)
	assert.Equal(t, "B", result.Ticket.TravelPackageIndicator)
	assert.Len(t, result.Passenger, 2)
	assert.Equal(t, "John", result.Passenger[0].FirstName)
	assert.Equal(t, "Jane", result.Passenger[1].FirstName)
	assert.Equal(t, "101", result.FlightLegDetails[0].FlightNumber)
	assert.Equal(t, "J", result.FlightLegDetails[0].ClassOfTravelling)
	assert.Equal(t, "x", result.FlightLegDetails[0].StopOverCode)
}

// TestAirlineData_MarshalsSpecKeyNames asserts on the serialized bytes, so a future rename cannot
// pass silently. Each of these keys was wrong at some point and each was dropped by the gateway.
func TestAirlineData_MarshalsSpecKeyNames(t *testing.T) {
	raw, err := json.Marshal(AirlineData{
		Ticket: &Ticket{TravelPackageIndicator: "B"},
		FlightLegDetails: []FlightLegDetails{{
			FlightNumber:      "101",
			ClassOfTravelling: "J",
			StopOverCode:      "x",
		}},
	})
	assert.Nil(t, err)

	body := string(raw)
	assert.Contains(t, body, `"class_of_travelling":"J"`)
	assert.Contains(t, body, `"stop_over_code":"x"`)
	assert.Contains(t, body, `"flight_number":"101"`)
	assert.Contains(t, body, `"travel_package_indicator":"B"`)

	// The keys the SDK used to send, which the API does not define.
	assert.NotContains(t, body, "class_of_traveling\"")
	assert.NotContains(t, body, `"stopover_code"`)
}

// TestPassengerAddress_MarshalsOnlyCountry pins the narrowed address: the specification defines
// exactly one property on passenger.address.
func TestPassengerAddress_MarshalsOnlyCountry(t *testing.T) {
	raw, err := json.Marshal(Passenger{
		FirstName: "John",
		Address:   &PassengerAddress{Country: common.US},
	})
	assert.Nil(t, err)

	body := string(raw)
	assert.Contains(t, body, `"address":{"country":"US"}`)
	assert.NotContains(t, body, "address_line1")
	assert.NotContains(t, body, `"zip"`)
}

// TestAccommodationData_UnmarshalFullSubTree covers the fields this row added or retyped.
func TestAccommodationData_UnmarshalFullSubTree(t *testing.T) {
	payload := `{
		"name": "The Sea View Hotel",
		"booking_reference": "HOTEL123",
		"check_in_date": "2023-06-20",
		"check_out_date": "2023-06-23",
		"address": {"address_line1": "123 Beach Road", "zip": "10001"},
		"state": "FL",
		"country": "USA",
		"city": "Los Angeles",
		"number_of_rooms": 2,
		"guests": [{"first_name": "Jane", "last_name": "Doe", "date_of_birth": "1985-07-14"}],
		"room": [{"rate": "70", "number_of_nights_at_room_rate": "3"}],
		"property_phone": [{"country_code": "44", "number": "7123456789"}],
		"customer_service_phone": [{"country_code": "44", "number": "7987654321"}]
	}`

	var stay AccommodationData
	assert.Nil(t, json.Unmarshal([]byte(payload), &stay))

	assert.Equal(t, "The Sea View Hotel", stay.Name)
	assert.Equal(t, "HOTEL123", stay.BookingReference)
	assert.Equal(t, "2023-06-20", shortDateValue(t, stay.CheckInDate))
	assert.Equal(t, "2023-06-23", shortDateValue(t, stay.CheckOutDate))
	assert.Equal(t, "123 Beach Road", stay.Address.AddressLine1)
	assert.Equal(t, "Los Angeles", stay.City)
	assert.Equal(t, 2, stay.NumberOfRooms)

	// state and country are free-form strings. Typed as the common.Country enum, country could
	// not carry "USA", a three-letter code.
	assert.Equal(t, "FL", stay.State)
	assert.Equal(t, "USA", stay.Country)

	assert.Len(t, stay.Guests, 1)
	assert.Equal(t, "Jane", stay.Guests[0].FirstName)
	assert.Len(t, stay.Room, 1)
	assert.Equal(t, "70", stay.Room[0].Rate)
	assert.Equal(t, "3", stay.Room[0].NumberOfNightsAtRoomRate)

	// property_phone and customer_service_phone were missing from the struct entirely.
	assert.Len(t, stay.PropertyPhone, 1)
	assert.Equal(t, "44", stay.PropertyPhone[0].CountryCode)
	assert.Equal(t, "7123456789", stay.PropertyPhone[0].Number)
	assert.Len(t, stay.CustomerServicePhone, 1)
	assert.Equal(t, "7987654321", stay.CustomerServicePhone[0].Number)
}
