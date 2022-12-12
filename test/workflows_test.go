package test

import (
	"fmt"
	"github.com/checkout/checkout-sdk-go/workflows/reflows"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/workflows"
	"github.com/checkout/checkout-sdk-go/workflows/actions"
	"github.com/checkout/checkout-sdk-go/workflows/conditions"
	"github.com/checkout/checkout-sdk-go/workflows/events"
)

var (
	entities   = []string{"ent_kidtcgc3ge5unf4a5i6enhnr5m"}
	eventsList = map[string][]string{
		"gateway": {"card_verification_declined",
			"card_verified",
			"payment_approved",
			"payment_authorization_increment_declined",
			"payment_authorization_incremented",
			"payment_capture_declined",
			"payment_captured",
			"payment_declined",
			"payment_refund_declined",
			"payment_refunded",
			"payment_void_declined",
			"payment_voided"},
		"dispute": {"dispute_canceled",
			"dispute_evidence_required",
			"dispute_expired",
			"dispute_lost",
			"dispute_resolved",
			"dispute_won"},
	}
	processingChannel = []string{"pc_5jp2az55l3cuths25t5p3xhwru"}
)

func TestCreateWorkflow(t *testing.T) {
	cases := []struct {
		name    string
		request workflows.CreateWorkflowRequest
		checker func(*common.IdResponse, error)
	}{
		{
			name: "when request is valid then should create workflow",
			request: workflows.CreateWorkflowRequest{
				Name:   "Test",
				Active: true,
				Conditions: []conditions.ConditionsRequest{
					getEntityCondition(),
					getEventCondition(),
					getProcessingChannelCondition(),
				},
				Actions: []actions.ActionsRequest{
					getWebhookAction(),
				},
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
			},
		},
		{
			name: "when conditions missing then should return error",
			request: workflows.CreateWorkflowRequest{
				Name:   "Test",
				Active: true,
				Actions: []actions.ActionsRequest{
					getWebhookAction(),
				},
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Contains(t, chkErr.Data.ErrorCodes, "conditions_required")
			},
		},
		{
			name: "when actions missing then should return error",
			request: workflows.CreateWorkflowRequest{
				Name:   "Test",
				Active: true,
				Conditions: []conditions.ConditionsRequest{
					getEntityCondition(),
					getEventCondition(),
					getProcessingChannelCondition(),
				},
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Contains(t, chkErr.Data.ErrorCodes, "actions_required")
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			response, err := client.CreateWorkflow(tc.request)

			tc.checker(response, err)

			if response != nil {
				if _, err := client.RemoveWorkflow(response.Id); err != nil {
					assert.Fail(t, fmt.Sprintf("Failed to remove workflow: %s", err.Error()))
				}
			}
		})
	}
}

func TestGetWorkflows(t *testing.T) {
	createWorkflow(t)

	cases := []struct {
		name    string
		checker func(*workflows.GetWorkflowsResponse, error)
	}{
		{
			name: "when retrieving all workflows then it returns existing workflows",
			checker: func(response *workflows.GetWorkflowsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Workflows)
				assert.NotEmpty(t, response.Workflows)
				for _, workflow := range response.Workflows {
					assert.NotNil(t, workflow.Id)
					assert.NotNil(t, workflow.Name)
					assert.True(t, workflow.Active)
				}
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetWorkflows())
		})
	}
}

func TestGetWorkflow(t *testing.T) {
	workflowId := createWorkflow(t).Id

	cases := []struct {
		name       string
		workflowId string
		checker    func(*workflows.GetWorkflowResponse, error)
	}{
		{
			name:       "when workflow exists return workflow details",
			workflowId: workflowId,
			checker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Name)
				assert.True(t, response.Active)
				assert.Equal(t, workflowId, response.Id)
				assert.Equal(t, "Test", response.Name)

				assert.NotNil(t, response.Actions)
				assert.NotEmpty(t, response.Actions)
				for _, action := range response.Actions {
					assert.NotNil(t, action.Id)
					assert.NotNil(t, action.Url)
					assert.NotNil(t, action.Signature)
				}

				assert.NotNil(t, response.Conditions)
				assert.NotEmpty(t, response.Conditions)
				for _, condition := range response.Conditions {
					switch condition.Type {
					case conditions.Entity:
						assert.NotNil(t, condition.Entities)
						assert.Equal(t, entities, condition.Entities)
					case conditions.Event:
						assert.NotNil(t, condition.Events)
						assert.Equal(t, eventsList, condition.Events)
					case conditions.ProcessingChannel:
						assert.NotNil(t, condition.ProcessingChannels)
						assert.Equal(t, processingChannel, condition.ProcessingChannels)
					default:
						assert.Fail(t, fmt.Sprintf("Condition Type not supported: %s", condition.Type))
					}
				}
			},
		},
		{
			name:       "when workflow is inexistent then return error",
			workflowId: "not_found",
			checker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetWorkflow(tc.workflowId))
		})
	}
}

func TestUpdateWorkflow(t *testing.T) {
	workflow := createWorkflow(t)

	cases := []struct {
		name       string
		workflowId string
		request    workflows.UpdateWorkflowRequest
		checker    func(*workflows.UpdateWorkflowResponse, error)
	}{
		{
			name:       "when updating existing workflow then it should modify data",
			workflowId: workflow.Id,
			request: workflows.UpdateWorkflowRequest{
				Name:   "New Name",
				Active: false,
			},
			checker: func(response *workflows.UpdateWorkflowResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, "New Name", response.Name)
				assert.Equal(t, false, response.Active)
			},
		},
		{
			name:       "when updating inexistent workflow then return error",
			workflowId: "not_found",
			request:    workflows.UpdateWorkflowRequest{},
			checker: func(response *workflows.UpdateWorkflowResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateWorkflow(tc.workflowId, tc.request))
		})
	}
}

func TestUpdateWorkflowAction(t *testing.T) {
	workflow := getWorkflow(t, createWorkflow(t).Id)

	webhookActionRequest := actions.NewWebhookActionRequest()
	webhookActionRequest.Url = "https://new-url.com"
	webhookActionRequest.Headers = map[string]string{}
	webhookActionRequest.Signature = &actions.WebhookSignature{
		Method: "HMACSHA256",
		Key:    "8V8x0dLK%AyD*DNS8JJr",
	}

	cases := []struct {
		name       string
		workflowId string
		actionId   string
		request    actions.ActionsRequest
		putChecker func(*common.MetadataResponse, error)
		getChecker func(*workflows.GetWorkflowResponse, error)
	}{
		{
			name:       "when updating existing workflow action then it should modify data",
			workflowId: workflow.Id,
			actionId:   workflow.Actions[0].Id,
			request:    webhookActionRequest,
			putChecker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
			getChecker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.Equal(t, webhookActionRequest.Url, response.Actions[0].Url)
				assert.Equal(t, webhookActionRequest.Headers, response.Actions[0].Headers)
				assert.NotNil(t, response.Actions[0].Signature)
			},
		},
		{
			name:       "when updating inexistent workflow then return error",
			workflowId: "not_found",
			actionId:   "action_id",
			request:    webhookActionRequest,
			putChecker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
			getChecker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:       "when updating inexistent action in workflow then return error",
			workflowId: workflow.Id,
			actionId:   "not_found",
			request:    webhookActionRequest,
			putChecker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
			getChecker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.Equal(t, webhookActionRequest.Url, response.Actions[0].Url)
				assert.Equal(t, webhookActionRequest.Headers, response.Actions[0].Headers)
				assert.NotNil(t, response.Actions[0].Signature)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.putChecker(client.UpdateWorkflowAction(tc.workflowId, tc.actionId, tc.request))

			tc.getChecker(client.GetWorkflow(tc.workflowId))
		})
	}
}

func TestUpdateWorkflowConditions(t *testing.T) {
	workflow := getWorkflow(t, createWorkflow(t).Id)

	var eventConditionId string
	for _, condition := range workflow.Conditions {
		if condition.Type == conditions.Event {
			eventConditionId = condition.Id
		}
	}

	eventConditionRequest := conditions.NewEventConditionRequest()
	eventConditionRequest.Events = map[string][]string{
		"gateway": {"card_verification_declined",
			"card_verified",
			"payment_approved",
			"payment_authorization_increment_declined",
			"payment_authorization_incremented",
			"payment_capture_declined",
			"payment_captured",
			"payment_declined",
			"payment_refund_declined",
			"payment_refunded",
			"payment_void_declined",
			"payment_voided"},
	}

	cases := []struct {
		name        string
		workflowId  string
		conditionId string
		request     conditions.ConditionsRequest
		putChecker  func(*common.MetadataResponse, error)
		getChecker  func(*workflows.GetWorkflowResponse, error)
	}{
		{
			name:        "when updating existing workflow condition then it should modify data",
			workflowId:  workflow.Id,
			conditionId: eventConditionId,
			request:     eventConditionRequest,
			putChecker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
			getChecker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Conditions)
				assert.NotEmpty(t, response.Conditions)
				for _, condition := range response.Conditions {
					if condition.Type == conditions.Event {
						assert.Equal(t, eventConditionRequest.Events, condition.Events)
					}
				}
			},
		},
		{
			name:        "when updating inexistent workflow then return error",
			workflowId:  "not_found",
			conditionId: "condition_id",
			request:     eventConditionRequest,
			putChecker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
			getChecker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:        "when updating inexistent condition in workflow then return error",
			workflowId:  workflow.Id,
			conditionId: "not_found",
			request:     eventConditionRequest,
			putChecker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
			getChecker: func(response *workflows.GetWorkflowResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Conditions)
				assert.NotEmpty(t, response.Conditions)
				for _, condition := range response.Conditions {
					if condition.Type == conditions.Event {
						assert.Equal(t, eventConditionRequest.Events, condition.Events)
					}
				}
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.putChecker(client.UpdateWorkflowCondition(tc.workflowId, tc.conditionId, tc.request))

			tc.getChecker(client.GetWorkflow(tc.workflowId))
		})
	}
}

func TestGetEventTypes(t *testing.T) {
	cases := []struct {
		name    string
		checker func(*events.EventTypesResponse, error)
	}{
		{
			name: "when retrieving all event types then it returns existing event types",
			checker: func(response *events.EventTypesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.EventTypes)
				assert.NotEmpty(t, response.EventTypes)
				for _, eventType := range response.EventTypes {
					assert.NotNil(t, eventType.Id)
					assert.NotNil(t, eventType.Description)
					assert.NotNil(t, eventType.DisplayName)

					for _, event := range eventType.Events {
						assert.NotNil(t, event.Id)
						assert.NotNil(t, event.Description)
						assert.NotNil(t, event.DisplayName)
					}
				}
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetEventTypes())
		})
	}
}

func TestGetSubjectEvents(t *testing.T) {
	payment := makeCardPayment(t, true, 10)

	cases := []struct {
		name      string
		subjectId string
		checker   func(interface{}, error)
	}{
		{
			name:      "when fetching existing subject with events then return event details",
			subjectId: payment.Id,
			checker: func(rawResponse interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, rawResponse)
				response := rawResponse.(*events.SubjectEventsResponse)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			process := func() (interface{}, error) { return client.GetSubjectEvents(tc.subjectId) }
			predicate := func(data interface{}) bool {
				response := data.(*events.SubjectEventsResponse)
				return response.Events != nil && len(response.Events) >= 2
			}

			tc.checker(retriable(process, predicate, 2))
		})
	}
}

func TestGetEvent(t *testing.T) {
	payment := makeCardPayment(t, true, 10)
	subjectEvents := getSubjectEvents(t, payment.Id, 2)
	var captureEventId string
	for _, event := range subjectEvents {
		if event.Type == "payment_captured" {
			captureEventId = event.Id
			break
		}
	}

	cases := []struct {
		name    string
		eventId string
		checker func(*events.EventResponse, error)
	}{
		{
			name:    "when fetching existing event then return event details",
			eventId: captureEventId,
			checker: func(response *events.EventResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Type)
				assert.NotNil(t, response.Source)
				assert.NotNil(t, response.Timestamp)
				assert.Equal(t, "payment_captured", response.Type)
				assert.Equal(t, "gateway", response.Source)
			},
		},
		{
			name:    "when fetching inexisting event then return error",
			eventId: "not_found",
			checker: func(response *events.EventResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetEvent(tc.eventId))
		})
	}
}

func TestReflowByEvent(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")

	payment := makeCardPayment(t, false, 10)
	subjectEvents := getSubjectEvents(t, payment.Id, 1)

	cases := []struct {
		name    string
		eventId string
		checker func(*common.MetadataResponse, error)
	}{
		{
			name:    "when reflowing by existing event then it should reflow workflow",
			eventId: subjectEvents[0].Id,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:    "when reflowing by inexisting event then it should return error",
			eventId: "not_found",
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.ReflowByEvent(tc.eventId))
		})
	}
}

func TestReflowBySubject(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")

	payment := makeCardPayment(t, false, 10)

	cases := []struct {
		name      string
		subjectId string
		checker   func(interface{}, error)
	}{
		{
			name:      "when reflowing by existing subject then it should reflow workflow",
			subjectId: payment.Id,
			checker: func(rawResponse interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, rawResponse)
				response := rawResponse.(*common.MetadataResponse)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:      "when reflowing by inexisting subject then it should return error",
			subjectId: "not_found",
			checker: func(rawResponse interface{}, err error) {
				assert.Nil(t, rawResponse)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			process := func() (interface{}, error) { return client.ReflowBySubject(tc.subjectId) }
			predicate := func(data interface{}) bool {
				response := data.(*common.MetadataResponse)
				return response.HttpMetadata.StatusCode == http.StatusAccepted
			}

			tc.checker(retriable(process, predicate, 1))
		})
	}
}

func TestReflowByEventAndWorkflow(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")

	workflow := createWorkflow(t)

	payment := makeCardPayment(t, false, 10)
	subjectEvents := getSubjectEvents(t, payment.Id, 1)

	cases := []struct {
		name       string
		eventId    string
		workflowId string
		checker    func(*common.MetadataResponse, error)
	}{
		{
			name:       "when reflowing by existing event and workflow then it should reflow workflow",
			eventId:    subjectEvents[0].Id,
			workflowId: workflow.Id,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:       "when reflowing by inexisting event then it should return error",
			eventId:    "not_found",
			workflowId: workflow.Id,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
		{
			name:       "when reflowing by inexisting workflow then it should return error",
			eventId:    subjectEvents[0].Id,
			workflowId: "not_found",
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.ReflowByEventAndWorkflow(tc.eventId, tc.workflowId))
		})
	}
}

func TestReflowBySubjectAndWorkflow(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")
	
	workflow := createWorkflow(t)

	payment := makeCardPayment(t, false, 10)

	cases := []struct {
		name       string
		subjectId  string
		workflowId string
		checker    func(interface{}, error)
	}{
		{
			name:       "when reflowing by existing subject and workflow then it should reflow workflow",
			subjectId:  payment.Id,
			workflowId: workflow.Id,
			checker: func(rawResponse interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, rawResponse)
				response := rawResponse.(*common.MetadataResponse)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:       "when reflowing by inexisting subject then it should return error",
			subjectId:  "not_found",
			workflowId: workflow.Id,
			checker: func(rawResponse interface{}, err error) {
				assert.Nil(t, rawResponse)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
		{
			name:       "when reflowing by inexisting workflow then it should return error",
			subjectId:  payment.Id,
			workflowId: "not_found",
			checker: func(rawResponse interface{}, err error) {
				assert.Nil(t, rawResponse)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			process := func() (interface{}, error) {
				return client.ReflowBySubjectAndWorkflow(tc.subjectId, tc.workflowId)
			}
			predicate := func(data interface{}) bool {
				response := data.(*common.MetadataResponse)
				return response.HttpMetadata.StatusCode == http.StatusAccepted
			}

			tc.checker(retriable(process, predicate, 1))
		})
	}
}

func TestReflow(t *testing.T) {
	workflow := createWorkflow(t)

	payment := makeCardPayment(t, false, 10)
	subjectEvents := getSubjectEvents(t, payment.Id, 1)

	cases := []struct {
		name    string
		request reflows.ReflowRequest
		checker func(*common.MetadataResponse, error)
	}{
		{
			name: "when reflowing by EVENT and workflow it should reflow workflow",
			request: &reflows.ReflowByEventsRequest{
				Events:          []string{subjectEvents[0].Id},
				ReflowWorkflows: reflows.ReflowWorkflows{Workflows: []string{workflow.Id}},
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name: "when reflowing by SUBJECT and workflow it should reflow workflow",
			request: &reflows.ReflowBySubjectsRequest{
				Subjects:        []string{payment.Id},
				ReflowWorkflows: reflows.ReflowWorkflows{Workflows: []string{workflow.Id}},
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:    "when invalid data it should return error",
			request: &reflows.ReflowBySubjectsRequest{},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Reflow(tc.request))
		})
	}
}

func TestRemoveWorkflow(t *testing.T) {
	cases := []struct {
		name       string
		workflowId string
		checker    func(*common.MetadataResponse, error)
	}{
		{
			name:       "when workflow exists then delete workflow",
			workflowId: createWorkflow(t).Id,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:       "when workflow is inexistent then return error",
			workflowId: "not_found",
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RemoveWorkflow(tc.workflowId))
		})
	}
}

func TestGetActionInvocations(t *testing.T) {
	workflow := getWorkflow(t, createWorkflow(t).Id)
	payment := makeCardPayment(t, false, 10)
	subjectEvents := getSubjectEvents(t, payment.Id, 1)

	cases := []struct {
		name     string
		eventId  string
		actionId string
		checker  func(*actions.ActionInvocationsResponse, error)
	}{
		{
			name:     "when fetching existing eventId and actionId then return action invocations",
			eventId:  subjectEvents[0].Id,
			actionId: workflow.Actions[0].Id,
			checker: func(response *actions.ActionInvocationsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.WorkflowId)
				assert.NotNil(t, response.EventId)
				assert.NotNil(t, response.WorkflowActionId)
				assert.NotNil(t, response.ActionType)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.ActionInvocations)

				assert.Equal(t, subjectEvents[0].Id, response.EventId)
				assert.Equal(t, workflow.Actions[0].Id, response.WorkflowActionId)
			},
		},
		{
			name:     "when eventId is inexistent then return error",
			eventId:  "not_found",
			actionId: workflow.Actions[0].Id,
			checker: func(response *actions.ActionInvocationsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
		{
			name:     "when actionId is inexistent then return error",
			eventId:  subjectEvents[0].Id,
			actionId: "not_found",
			checker: func(response *actions.ActionInvocationsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				errChk := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, errChk.StatusCode)
			},
		},
	}

	client := DefaultApi().WorkFlows

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetActionInvocations(tc.eventId, tc.actionId))
		})
	}
}

func getEntityCondition() conditions.ConditionsRequest {
	entityCondition := conditions.NewEntityConditionRequest()
	entityCondition.Entities = entities

	return entityCondition
}

func getEventCondition() conditions.ConditionsRequest {
	eventCondition := conditions.NewEventConditionRequest()
	eventCondition.Events = eventsList

	return eventCondition
}

func getProcessingChannelCondition() conditions.ConditionsRequest {
	processingChannelCondition := conditions.NewProcessingChannelConditionRequest()
	processingChannelCondition.ProcessingChannels = processingChannel

	return processingChannelCondition
}

func getWebhookAction() actions.ActionsRequest {
	webhookAction := actions.NewWebhookActionRequest()
	webhookAction.Url = "https://google.com/fail"
	webhookAction.Headers = map[string]string{}
	webhookAction.Signature = &actions.WebhookSignature{
		Method: "HMACSHA256",
		Key:    "8V8x0dLK%AyD*DNS8JJr",
	}

	return webhookAction
}

func createWorkflow(t *testing.T) *common.IdResponse {
	request := workflows.CreateWorkflowRequest{
		Name:   "Test",
		Active: true,
		Conditions: []conditions.ConditionsRequest{
			getEntityCondition(),
			getEventCondition(),
			getProcessingChannelCondition(),
		},
		Actions: []actions.ActionsRequest{
			getWebhookAction(),
		},
	}

	response, err := DefaultApi().WorkFlows.CreateWorkflow(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating workflow - %s", err.Error()))
	}

	t.Cleanup(func() {
		if _, err = DefaultApi().WorkFlows.RemoveWorkflow(response.Id); err != nil {
			assert.Fail(t, fmt.Sprintf("Failed to remove workflow: %s", err.Error()))
		}
	})

	return response
}

func getWorkflow(t *testing.T, workflowId string) *workflows.GetWorkflowResponse {
	response, err := DefaultApi().WorkFlows.GetWorkflow(workflowId)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error getting workflow - %s", err.Error()))
	}

	return response
}

func getSubjectEvents(t *testing.T, subjectId string, eventsQuantity int) []events.SubjectEvent {
	process := func() (interface{}, error) {
		return DefaultApi().WorkFlows.GetSubjectEvents(subjectId)
	}
	predicate := func(data interface{}) bool {
		response := data.(*events.SubjectEventsResponse)
		return response.Events != nil && len(response.Events) >= eventsQuantity
	}

	response, err := retriable(process, predicate, 2)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error getting subject events - %s", err.Error()))
	}

	return response.(*events.SubjectEventsResponse).Events
}
