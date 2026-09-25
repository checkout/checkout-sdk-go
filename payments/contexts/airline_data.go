package contexts

import (
	"bytes"
	"encoding/json"
)

// UnmarshalJSON deserializes PaymentContextsAirlineData, accepting
// processing.airline_data[].passenger as either an array or a single object.
//
// PayPal is a payment contexts payment method and returns passenger as a bare object where the
// specification declares an array, so both shapes are accepted and a single object is normalized
// to a one-element slice. Callers only ever handle a slice.
//
// See MarshalJSON for the outbound shape: payment contexts rejects the array form.
func (a *PaymentContextsAirlineData) UnmarshalJSON(data []byte) error {
	// The alias sheds the method set, so json.Unmarshal below does not re-enter this method.
	type alias PaymentContextsAirlineData

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

	var single PaymentContextsPassenger
	if err := json.Unmarshal(trimmed, &single); err != nil {
		return err
	}
	a.Passenger = []PaymentContextsPassenger{single}
	return nil
}

// MarshalJSON serializes PaymentContextsAirlineData, emitting passenger as a single object when
// there is exactly one passenger and as an array only when there are several.
//
// POST /payment-contexts rejects the array form with passenger_required and accepts a single
// object, verified against the sandbox on 2026-09-25 with a complete airline_data block. This is
// the opposite of what the specification declares. See payments.AirlineData.MarshalJSON for the
// full cross-surface matrix.
func (a PaymentContextsAirlineData) MarshalJSON() ([]byte, error) {
	// The fields are listed explicitly, in specification order, rather than embedding an
	// alias: an embedded struct is flattened after the fields declared beside it, which
	// would emit passenger last. A new field on PaymentContextsAirlineData must be added here too.
	aux := struct {
		Ticket           *PaymentContextsTicket            `json:"ticket,omitempty"`
		Passenger        interface{}                       `json:"passenger,omitempty"`
		FlightLegDetails []PaymentContextsFlightLegDetails `json:"flight_leg_details,omitempty"`
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
