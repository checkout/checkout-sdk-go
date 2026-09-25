package contexts

import "encoding/json"

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

	trimmed := trimJSONSpace(aux.Passenger)
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

// trimJSONSpace strips the whitespace JSON permits before a value, so the first meaningful byte
// can be used to tell an array from an object.
func trimJSONSpace(data []byte) []byte {
	for len(data) > 0 {
		switch data[0] {
		case ' ', '\t', '\n', '\r':
			data = data[1:]
		default:
			return data
		}
	}
	return data
}

// MarshalJSON serializes PaymentContextsAirlineData, emitting passenger as a single object when
// there is exactly one passenger and as an array only when there are several.
//
// POST /payment-contexts rejects the array form with passenger_required and accepts a single
// object, verified against the sandbox on 2026-09-25 with a complete airline_data block. This is
// the opposite of what the specification declares. See payments.AirlineData.MarshalJSON for the
// full cross-surface matrix.
func (a PaymentContextsAirlineData) MarshalJSON() ([]byte, error) {
	type alias PaymentContextsAirlineData

	aux := struct {
		alias
		Passenger interface{} `json:"passenger,omitempty"`
	}{alias: alias(a)}

	// Clear the embedded copy so passenger is written once, by the override above.
	aux.alias.Passenger = nil

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
