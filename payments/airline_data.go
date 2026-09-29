package payments

import (
	"bytes"
	"encoding/json"
)

// UnmarshalJSON deserializes AirlineData, accepting processing.airline_data[].passenger as either
// an array or a single object.
//
// The specification declares passenger as an array on AirlineData, which is the shape reached from
// POST /payments, from the GET /payments/{id} response and from payment contexts. On
// PaymentInterfacesProcessingAirlineData, which payment sessions, hosted payments and payment links
// resolve to, it is declared oneOf[array, object] with the note "PayPal requires a single object".
// Both branches resolve to the same object, so a bare object is normalized to a one-element slice
// and callers only ever handle a slice.
//
// Before this existed, a response carrying passenger data failed the whole call with:
//
//	json: cannot unmarshal array into Go struct field
//	AirlineData.processing.airline_data.passenger of type payments.Passenger
//
// because Passenger was a single *Passenger. Reported internally, pre-3.3.0.
//
// See MarshalJSON for the outbound shape, which is NOT simply "always an array".
func (a *AirlineData) UnmarshalJSON(data []byte) error {
	// The alias sheds the method set, so json.Unmarshal below does not re-enter this method.
	type alias AirlineData

	aux := struct {
		*alias
		Passenger json.RawMessage `json:"passenger,omitempty"`
	}{alias: (*alias)(a)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	a.Passenger = nil

	// Whitespace is legal before a JSON value, so trim it before testing the first byte.
	trimmed := bytes.TrimLeft(aux.Passenger, " \t\n\r")
	if len(trimmed) == 0 || string(trimmed) == "null" {
		return nil
	}

	if trimmed[0] == '[' {
		return json.Unmarshal(trimmed, &a.Passenger)
	}

	var single Passenger
	if err := json.Unmarshal(trimmed, &single); err != nil {
		return err
	}
	a.Passenger = []Passenger{single}
	return nil
}

// MarshalJSON serializes AirlineData, emitting passenger as a single object when there is exactly
// one passenger and as an array only when there are several.
//
// This is driven by the live API, which does not match the specification in either direction.
// Verified against the sandbox on 2026-09-25 with a complete airline_data block:
//
//	surface                  passenger: object   passenger: array
//	POST /payments           201                 201
//	POST /hosted-payments    accepted            422 processing_airline_data_0_passenger_invalid
//	POST /payment-links      accepted            422 processing_airline_data_0_passenger_invalid
//	POST /payment-contexts   201                 422 passenger_required
//
// So a single object is accepted on every request surface, and an array is accepted only on
// POST /payments. The specification says the opposite: it declares AirlineData.passenger as
// array-only (which payment contexts rejects) and PaymentInterfacesProcessingAirlineData.passenger
// as oneOf[array, object] (whose array branch is rejected). ProcessingSettings is shared by
// POST /payments, hosted payments and payment links, so a struct that always emitted an array
// would break the latter two.
//
// Emitting an object for one passenger is therefore safe everywhere. Several passengers can only
// be expressed as an array, which only POST /payments accepts; that is an API limitation, not an
// SDK choice.
func (a AirlineData) MarshalJSON() ([]byte, error) {
	// The fields are listed explicitly, in specification order, rather than embedding an
	// alias: an embedded struct is flattened after the fields declared beside it, which
	// would emit passenger last. A new field on AirlineData must be added here too.
	aux := struct {
		Ticket           *Ticket            `json:"ticket,omitempty"`
		Passenger        interface{}        `json:"passenger,omitempty"`
		FlightLegDetails []FlightLegDetails `json:"flight_leg_details,omitempty"`
	}{
		Ticket:           a.Ticket,
		FlightLegDetails: a.FlightLegDetails,
	}

	switch len(a.Passenger) {
	case 0:
		aux.Passenger = nil
	case 1:
		aux.Passenger = a.Passenger[0]
	default:
		aux.Passenger = a.Passenger
	}

	return json.Marshal(aux)
}
