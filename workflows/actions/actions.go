package actions

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

type WorkflowActionType string

const (
	Webhook WorkflowActionType = "webhook"
)

type WorkflowActionStatus string

const (
	Pending    WorkflowActionStatus = "pending"
	Successful WorkflowActionStatus = "successful"
	Failed     WorkflowActionStatus = "failed"
)

// Requests
type (
	ActionsRequest interface {
		GetType() WorkflowActionType
	}

	WorkflowActions struct {
		Type WorkflowActionType `json:"type,omitempty"`
	}

	webhookActionRequest struct {
		WorkflowActions
		Url       string            `json:"url,omitempty"`
		Headers   map[string]string `json:"headers,omitempty"`
		Signature *WebhookSignature `json:"signature,omitempty"`
	}
)

func NewWebhookActionRequest() *webhookActionRequest {
	return &webhookActionRequest{WorkflowActions: WorkflowActions{Type: Webhook}}
}

func (a *webhookActionRequest) GetType() WorkflowActionType {
	return a.Type
}

type WebhookSignature struct {
	Method string `json:"method,omitempty"`
	Key    string `json:"key,omitempty"`
}

// Responses
type (
	ActionsResponse struct {
		Id   string
		Type WorkflowActionType
		Actions
		Links map[string]common.Link `json:"_links,omitempty"`
	}

	ActionInvocationsResponse struct {
		HttpMetadata      common.HttpMetadata
		WorkflowId        string                     `json:"workflow_id,omitempty"`
		EventId           string                     `json:"event_id,omitempty"`
		WorkflowActionId  string                     `json:"workflow_action_id,omitempty"`
		ActionType        WorkflowActionType         `json:"action_type,omitempty"`
		Status            WorkflowActionStatus       `json:"status,omitempty"`
		ActionInvocations []WorkflowActionInvocation `json:"action_invocations,omitempty"`
	}
)

type Actions struct {
	*WebhookAction
}

type WebhookAction struct {
	Url       string            `json:"url,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
	Signature *WebhookSignature `json:"signature,omitempty"`
}

type WorkflowActionInvocation struct {
	InvocationId string     `json:"invocation_id,omitempty"`
	Timestamp    *time.Time `json:"timestamp,omitempty"`
	Retry        bool       `json:"retry,omitempty"`
	Succeeded    bool       `json:"succeeded,omitempty"`
	FinalAttempt bool       `json:"final,omitempty"`
	// Deprecated: This property will be removed in the future, and should not be used. Use ResultDetails instead.
	Result        map[string]interface{} `json:"result,omitempty"`
	ResultDetails map[string]interface{} `json:"result_details,omitempty"`
}
