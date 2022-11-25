package events

import (
	"encoding/json"
	"github.com/checkout/checkout-sdk-go/common"
)

const path = "event-types"

type (
	EventTypesResponse struct {
		HttpResponse common.HttpMetadata
		EventTypes   []EventTypes
	}

	EventTypes struct {
		Version    string   `json:"version"`
		EventTypes []string `json:"event_types"`
	}
)

func (e *EventTypesResponse) UnmarshalJSON(data []byte) error {
	var eventTypes []EventTypes
	if err := json.Unmarshal(data, &eventTypes); err != nil {
		return err
	}
	e.EventTypes = eventTypes
	return nil
}
