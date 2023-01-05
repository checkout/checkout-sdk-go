package conditions

import "github.com/checkout/checkout-sdk-go/common"

type WorkflowConditionType string

const (
	Event             WorkflowConditionType = "event"
	Entity            WorkflowConditionType = "entity"
	ProcessingChannel WorkflowConditionType = "processing_channel"
)

// Requests
type (
	ConditionsRequest interface {
		GetType() WorkflowConditionType
	}

	WorkflowConditions struct {
		Type WorkflowConditionType `json:"type,omitempty"`
	}

	entityConditionRequest struct {
		WorkflowConditions
		Entities []string `json:"entities,omitempty"`
	}

	eventConditionRequest struct {
		WorkflowConditions
		Events map[string][]string `json:"events,omitempty"`
	}

	processingChannelConditionRequest struct {
		WorkflowConditions
		ProcessingChannels []string `json:"processing_channels,omitempty"`
	}
)

func NewEntityConditionRequest() *entityConditionRequest {
	return &entityConditionRequest{WorkflowConditions: WorkflowConditions{Type: Entity}}
}

func NewEventConditionRequest() *eventConditionRequest {
	return &eventConditionRequest{WorkflowConditions: WorkflowConditions{Type: Event}}
}

func NewProcessingChannelConditionRequest() *processingChannelConditionRequest {
	return &processingChannelConditionRequest{WorkflowConditions: WorkflowConditions{Type: ProcessingChannel}}
}

func (c *entityConditionRequest) GetType() WorkflowConditionType {
	return c.Type
}

func (c *eventConditionRequest) GetType() WorkflowConditionType {
	return c.Type
}

func (c *processingChannelConditionRequest) GetType() WorkflowConditionType {
	return c.Type
}

// Responses
type (
	ConditionsResponse struct {
		Id   string                `json:"id,omitempty"`
		Type WorkflowConditionType `json:"type,omitempty"`
		Conditions
		Links map[string]common.Link `json:"_links,omitempty"`
	}

	Conditions struct {
		*EntitiesCondition
		*EventsCondition
		*ProcessingChannelCondition
	}

	EntitiesCondition struct {
		Entities []string `json:"entities,omitempty"`
	}

	EventsCondition struct {
		Events map[string][]string `json:"events,omitempty"`
	}

	ProcessingChannelCondition struct {
		ProcessingChannels []string `json:"processing_channels,omitempty"`
	}
)
