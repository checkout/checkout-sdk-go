package events

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	// ErrMissingEventData is returned when an event has no data object.
	ErrMissingEventData = errors.New("events: event data is missing")

	// ErrUnexpectedEventType is returned when an event type cannot be decoded
	// into the requested payload.
	ErrUnexpectedEventType = errors.New("events: unexpected event type")

	// ErrMissingEventField is returned when a required payload field is empty.
	ErrMissingEventField = errors.New("events: required event data field is missing")
)

var errNilDecodeDestination = errors.New("events: decode destination is nil")

// RawData returns a copy of the original JSON value of the event data field.
// The result is nil when the event was not decoded from JSON, or when data was absent.
func (e *EventResponse) RawData() json.RawMessage {
	if e == nil || len(e.rawData) == 0 {
		return nil
	}
	out := make(json.RawMessage, len(e.rawData))
	copy(out, e.rawData)
	return out
}

// DecodeData unmarshals the event data into dest, which must be a non-nil pointer.
//
// When the event was decoded from JSON, the original data bytes are used.
// When the event was built in memory with only Data set, that map is encoded
// and then unmarshaled into dest so existing struct literals keep working.
//
// ErrMissingEventData is returned when data is absent or JSON null.
// A JSON value that does not match dest is returned wrapped around the encoding/json error.
func (e *EventResponse) DecodeData(dest interface{}) error {
	if isNilDestination(dest) {
		return errNilDecodeDestination
	}
	if e == nil {
		return ErrMissingEventData
	}

	raw, err := e.eventDataBytes()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, dest); err != nil {
		return fmt.Errorf("events: decode event data: %w", err)
	}
	return nil
}

// UnmarshalJSON decodes an event envelope and keeps both the generic Data map
// and the original data JSON. A data value that is not a JSON object or null
// fails the whole unmarshal, matching the previous map-based behavior.
func (e *EventResponse) UnmarshalJSON(data []byte) error {
	type eventResponse EventResponse
	var aux struct {
		eventResponse
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	*e = EventResponse(aux.eventResponse)
	if len(aux.Data) > 0 {
		e.rawData = append(json.RawMessage(nil), aux.Data...)
	}
	if isNullRaw(aux.Data) {
		return nil
	}
	if err := json.Unmarshal(aux.Data, &e.Data); err != nil {
		return fmt.Errorf("events: event data: %w", err)
	}
	return nil
}

func (e *EventResponse) eventDataBytes() (json.RawMessage, error) {
	if !isNullRaw(e.rawData) {
		return e.rawData, nil
	}
	if e.Data == nil {
		return nil, ErrMissingEventData
	}
	raw, err := json.Marshal(e.Data)
	if err != nil {
		return nil, fmt.Errorf("events: encode event data: %w", err)
	}
	return raw, nil
}

func isNullRaw(raw json.RawMessage) bool {
	trimmed := bytes.TrimSpace(raw)
	return len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null"))
}

func isNilDestination(dest interface{}) bool {
	if dest == nil {
		return true
	}
	value := reflect.ValueOf(dest)
	switch value.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return value.IsNil()
	default:
		return false
	}
}
