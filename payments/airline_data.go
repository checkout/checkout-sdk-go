package payments

import "encoding/json"

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

	trimmed := trimJSONSpace(aux.Passenger)
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
	type alias AirlineData

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
