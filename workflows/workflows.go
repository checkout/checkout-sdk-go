package workflows

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/workflows/actions"
	"github.com/checkout/checkout-sdk-go/workflows/conditions"
)

const (
	WorkflowsPath  = "workflows"
	ActionsPath    = "actions"
	ConditionsPath = "conditions"
	EventTypesPath = "event-types"
	EventsPath     = "events"
	SubjectPath    = "subject"
	ReflowPath     = "reflow"
	WorkflowPath   = "workflow"
)

// Requests
type (
	CreateWorkflowRequest struct {
		Name       string                         `json:"name,omitempty"`
		Active     bool                           `json:"active,omitempty"`
		Conditions []conditions.ConditionsRequest `json:"conditions,omitempty"`
		Actions    []actions.ActionsRequest       `json:"actions,omitempty"`
	}

	UpdateWorkflowRequest struct {
		Name   string `json:"name,omitempty"`
		Active bool   `json:"active,omitempty"`
	}
)

type Workflow struct {
	Id     string                 `json:"id,omitempty"`
	Name   string                 `json:"name,omitempty"`
	Active bool                   `json:"active,omitempty"`
	Links  map[string]common.Link `json:"_links,omitempty"`
}

// Responses
type (
	GetWorkflowsResponse struct {
		HttpMetadata common.HttpMetadata
		Workflows    []Workflow `json:"data,omitempty"`
	}

	GetWorkflowResponse struct {
		HttpMetadata common.HttpMetadata
		Workflow
		Conditions []conditions.ConditionsResponse `json:"conditions,omitempty"`
		Actions    []actions.ActionsResponse       `json:"actions,omitempty"`
		Links      map[string]common.Link          `json:"_links,omitempty"`
	}

	UpdateWorkflowResponse struct {
		HttpMetadata common.HttpMetadata
		Name         string `json:"name,omitempty"`
		Active       bool   `json:"active,omitempty"`
	}
)
