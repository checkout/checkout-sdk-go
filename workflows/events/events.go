package events

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
)

type ActionInvocationStatus string

const (
	Pending    ActionInvocationStatus = "pending"
	Successful ActionInvocationStatus = "successful"
	Failed     ActionInvocationStatus = "failed"
)

type (
	EventTypesResponse struct {
		HttpMetadata common.HttpMetadata
		EventTypes   []WorkflowEventTypes
	}

	EventResponse struct {
		HttpMetadata      common.HttpMetadata
		Id                string                 `json:"id,omitempty"`
		Source            string                 `json:"source,omitempty"`
		Type              string                 `json:"type,omitempty"`
		Timestamp         string                 `json:"timestamp,omitempty"`
		Version           string                 `json:"version,omitempty"`
		Data              map[string]interface{} `json:"data,omitempty"`
		ActionInvocations []ActionInvocation     `json:"action_invocations,omitempty"`
		Links             map[string]common.Link `json:"_links,omitempty"`
	}

	SubjectEventsResponse struct {
		HttpMetadata common.HttpMetadata
		Events       []SubjectEvent `json:"data,omitempty"`
	}
)

func (e *EventTypesResponse) UnmarshalJSON(data []byte) error {
	var eventTypes []WorkflowEventTypes
	if err := json.Unmarshal(data, &eventTypes); err != nil {
		return err
	}
	e.EventTypes = eventTypes
	return nil
}

type WorkflowEventTypes struct {
	Id          string  `json:"id,omitempty"`
	DisplayName string  `json:"display_name,omitempty"`
	Description string  `json:"description,omitempty"`
	Events      []Event `json:"events,omitempty"`
}

type Event struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Description string `json:"description,omitempty"`
}

type ActionInvocation struct {
	WorkflowId       string                 `json:"workflow_id,omitempty"`
	WorkflowActionId string                 `json:"workflow_action_id,omitempty"`
	Status           ActionInvocationStatus `json:"status,omitempty"`
	Links            map[string]common.Link `json:"_links,omitempty"`
}

type SubjectEvent struct {
	Id        string                 `json:"id,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Timestamp string                 `json:"timestamp,omitempty"`
	Links     map[string]common.Link `json:"_links,omitempty"`
}
