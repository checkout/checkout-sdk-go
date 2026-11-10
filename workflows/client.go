package workflows

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/workflows/actions"
	"github.com/checkout/checkout-sdk-go/v3/workflows/conditions"
	"github.com/checkout/checkout-sdk-go/v3/workflows/events"
	"github.com/checkout/checkout-sdk-go/v3/workflows/reflows"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

func NewClient(configuration *configuration.Configuration, apiClient client.HttpClient) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

func (c *Client) GetWorkflows() (*GetWorkflowsResponse, error) {
	return c.GetWorkflowsWithContext(context.Background())
}

func (c *Client) GetWorkflowsWithContext(
	ctx context.Context,
) (*GetWorkflowsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetWorkflowsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateWorkflow(request CreateWorkflowRequest) (*common.IdResponse, error) {
	return c.CreateWorkflowWithContext(context.Background(), request)
}

func (c *Client) CreateWorkflowWithContext(
	ctx context.Context,
	request CreateWorkflowRequest,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetWorkflow sends GET /workflows/{workflowId}: get a workflow.
//
// workflowId is the {workflowId} path value: the workflow identifier.
//
// It returns the GetWorkflowResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetWorkflow(workflowId string) (*GetWorkflowResponse, error) {
	return c.GetWorkflowWithContext(context.Background(), workflowId)
}

// GetWorkflowWithContext is GetWorkflow with a caller-supplied context.
func (c *Client) GetWorkflowWithContext(
	ctx context.Context,
	workflowId string,
) (*GetWorkflowResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetWorkflowResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RemoveWorkflow sends DELETE /workflows/{workflowId}: remove a workflow.
//
// workflowId is the {workflowId} path value: the workflow identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) RemoveWorkflow(workflowId string) (*common.MetadataResponse, error) {
	return c.RemoveWorkflowWithContext(context.Background(), workflowId)
}

// RemoveWorkflowWithContext is RemoveWorkflow with a caller-supplied context.
func (c *Client) RemoveWorkflowWithContext(
	ctx context.Context,
	workflowId string,
) (*common.MetadataResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateWorkflow sends PATCH /workflows/{workflowId}: patch a workflow.
//
// workflowId is the {workflowId} path value: the workflow identifier. request is sent as the JSON
// body.
//
// It returns the UpdateWorkflowResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateWorkflow(
	workflowId string,
	request UpdateWorkflowRequest,
) (*UpdateWorkflowResponse, error) {
	return c.UpdateWorkflowWithContext(context.Background(), workflowId, request)
}

// UpdateWorkflowWithContext is UpdateWorkflow with a caller-supplied context.
func (c *Client) UpdateWorkflowWithContext(
	ctx context.Context,
	workflowId string,
	request UpdateWorkflowRequest,
) (*UpdateWorkflowResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response UpdateWorkflowResponse
	err = c.apiClient.PatchWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId),
		auth,
		request,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// AddWorkflowAction sends POST /workflows/{workflowId}/actions: add a workflow action.
//
// workflowId is the {workflowId} path value: the workflow identifier. request is sent as the JSON
// body.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) AddWorkflowAction(
	workflowId string,
	request actions.ActionsRequest,
) (*common.IdResponse, error) {
	return c.AddWorkflowActionWithContext(context.Background(), workflowId, request)
}

// AddWorkflowActionWithContext is AddWorkflowAction with a caller-supplied context.
func (c *Client) AddWorkflowActionWithContext(
	ctx context.Context,
	workflowId string,
	request actions.ActionsRequest,
) (*common.IdResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId, ActionsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateWorkflowAction sends PUT /workflows/{workflowId}/actions/{workflowActionId}: update a
// workflow action.
//
// workflowId is the {workflowId} path value: the workflow identifier. actionId is the
// {workflowActionId} path value: the workflow action identifier. request is sent as the JSON body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId or
// actionId returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateWorkflowAction(
	workflowId string,
	actionId string,
	request actions.ActionsRequest,
) (*common.MetadataResponse, error) {
	return c.UpdateWorkflowActionWithContext(context.Background(), workflowId, actionId, request)
}

// UpdateWorkflowActionWithContext is UpdateWorkflowAction with a caller-supplied context.
func (c *Client) UpdateWorkflowActionWithContext(
	ctx context.Context,
	workflowId string,
	actionId string,
	request actions.ActionsRequest,
) (*common.MetadataResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	actionId, err = common.EncodePathSegment(actionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId, ActionsPath, actionId),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// RemoveWorkflowAction sends DELETE /workflows/{workflowId}/actions/{workflowActionId}: remove a
// workflow action.
//
// workflowId is the {workflowId} path value: the workflow identifier. actionId is the
// {workflowActionId} path value: the workflow action identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId or
// actionId returns a CheckoutArgumentError and sends no request.
func (c *Client) RemoveWorkflowAction(workflowId, actionId string) (*common.MetadataResponse, error) {
	return c.RemoveWorkflowActionWithContext(context.Background(), workflowId, actionId)
}

// RemoveWorkflowActionWithContext is RemoveWorkflowAction with a caller-supplied context.
func (c *Client) RemoveWorkflowActionWithContext(
	ctx context.Context,
	workflowId string,
	actionId string,
) (*common.MetadataResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	actionId, err = common.EncodePathSegment(actionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId, ActionsPath, actionId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// AddWorkflowCondition sends POST /workflows/{workflowId}/conditions: add a workflow condition.
//
// workflowId is the {workflowId} path value: the workflow identifier. request is sent as the JSON
// body.
//
// It returns the common.IdResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) AddWorkflowCondition(
	workflowId string,
	request conditions.ConditionsRequest,
) (*common.IdResponse, error) {
	return c.AddWorkflowConditionWithContext(context.Background(), workflowId, request)
}

// AddWorkflowConditionWithContext is AddWorkflowCondition with a caller-supplied context.
func (c *Client) AddWorkflowConditionWithContext(
	ctx context.Context,
	workflowId string,
	request conditions.ConditionsRequest,
) (*common.IdResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.IdResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId, ConditionsPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateWorkflowCondition sends PUT /workflows/{workflowId}/conditions/{workflowConditionId}:
// update a workflow condition.
//
// workflowId is the {workflowId} path value: the workflow identifier. conditionId is the
// {workflowConditionId} path value: the workflow condition identifier. request is sent as the JSON
// body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId or
// conditionId returns a CheckoutArgumentError and sends no request.
func (c *Client) UpdateWorkflowCondition(
	workflowId string,
	conditionId string,
	request conditions.ConditionsRequest,
) (*common.MetadataResponse, error) {
	return c.UpdateWorkflowConditionWithContext(context.Background(), workflowId, conditionId, request)
}

// UpdateWorkflowConditionWithContext is UpdateWorkflowCondition with a caller-supplied context.
func (c *Client) UpdateWorkflowConditionWithContext(
	ctx context.Context,
	workflowId string,
	conditionId string,
	request conditions.ConditionsRequest,
) (*common.MetadataResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	conditionId, err = common.EncodePathSegment(conditionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PutWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId, ConditionsPath, conditionId),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// RemoveWorkflowCondition sends DELETE /workflows/{workflowId}/conditions/{workflowConditionId}:
// remove a workflow condition.
//
// workflowId is the {workflowId} path value: the workflow identifier. conditionId is the
// {workflowConditionId} path value: the workflow condition identifier.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId or
// conditionId returns a CheckoutArgumentError and sends no request.
func (c *Client) RemoveWorkflowCondition(workflowId, conditionId string) (*common.MetadataResponse, error) {
	return c.RemoveWorkflowConditionWithContext(context.Background(), workflowId, conditionId)
}

// RemoveWorkflowConditionWithContext is RemoveWorkflowCondition with a caller-supplied context.
func (c *Client) RemoveWorkflowConditionWithContext(
	ctx context.Context,
	workflowId string,
	conditionId string,
) (*common.MetadataResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	conditionId, err = common.EncodePathSegment(conditionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.DeleteWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId, ConditionsPath, conditionId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// TestWorkflow sends POST /workflows/{workflowId}/test: test a workflow.
//
// workflowId is the {workflowId} path value: the workflow identifier. request is sent as the JSON
// body.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." workflowId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) TestWorkflow(workflowId string, request events.EventTypesRequest) (*common.MetadataResponse, error) {
	return c.TestWorkflowWithContext(context.Background(), workflowId, request)
}

// TestWorkflowWithContext is TestWorkflow with a caller-supplied context.
func (c *Client) TestWorkflowWithContext(
	ctx context.Context,
	workflowId string,
	request events.EventTypesRequest,
) (*common.MetadataResponse, error) {
	workflowId, err := common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, workflowId, TestPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetEventTypes() (*events.EventTypesResponse, error) {
	return c.GetEventTypesWithContext(context.Background())
}

func (c *Client) GetEventTypesWithContext(ctx context.Context) (*events.EventTypesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response events.EventTypesResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventTypesPath),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetEvent sends GET /workflows/events/{eventId}: get an event.
//
// eventId is the {eventId} path value: the event identifier.
//
// It returns the events.EventResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) GetEvent(eventId string) (*events.EventResponse, error) {
	return c.GetEventWithContext(context.Background(), eventId)
}

// GetEventWithContext is GetEvent with a caller-supplied context.
func (c *Client) GetEventWithContext(
	ctx context.Context,
	eventId string,
) (*events.EventResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response events.EventResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, eventId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetActionInvocations sends GET /workflows/events/{eventId}/actions/{workflowActionId}: get action
// invocations.
//
// eventId is the {eventId} path value: the event identifier. actionId is the {workflowActionId}
// path value: the workflow action identifier.
//
// It returns the actions.ActionInvocationsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId or
// actionId returns a CheckoutArgumentError and sends no request.
func (c *Client) GetActionInvocations(eventId, actionId string) (*actions.ActionInvocationsResponse, error) {
	return c.GetActionInvocationsWithContext(context.Background(), eventId, actionId)
}

// GetActionInvocationsWithContext is GetActionInvocations with a caller-supplied context.
func (c *Client) GetActionInvocationsWithContext(
	ctx context.Context,
	eventId string,
	actionId string,
) (*actions.ActionInvocationsResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

	actionId, err = common.EncodePathSegment(actionId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response actions.ActionInvocationsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, eventId, ActionsPath, actionId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ReflowByEvent sends POST /workflows/events/{eventId}/reflow: reflow by event.
//
// eventId is the {eventId} path value: the unique identifier for the event to be reflowed.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) ReflowByEvent(eventId string) (*common.MetadataResponse, error) {
	return c.ReflowByEventWithContext(context.Background(), eventId)
}

// ReflowByEventWithContext is ReflowByEvent with a caller-supplied context.
func (c *Client) ReflowByEventWithContext(
	ctx context.Context,
	eventId string,
) (*common.MetadataResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, eventId, ReflowPath),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ReflowByEventAndWorkflow sends POST /workflows/events/{eventId}/workflow/{workflowId}/reflow:
// reflow by event and workflow.
//
// eventId is the {eventId} path value: the unique identifier for the event to be reflowed.
// workflowId is the {workflowId} path value: the identifier of the workflow whose actions you want
// to trigger.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." eventId or
// workflowId returns a CheckoutArgumentError and sends no request.
func (c *Client) ReflowByEventAndWorkflow(eventId, workflowId string) (*common.MetadataResponse, error) {
	return c.ReflowByEventAndWorkflowWithContext(context.Background(), eventId, workflowId)
}

// ReflowByEventAndWorkflowWithContext is ReflowByEventAndWorkflow with a caller-supplied context.
func (c *Client) ReflowByEventAndWorkflowWithContext(
	ctx context.Context,
	eventId string,
	workflowId string,
) (*common.MetadataResponse, error) {
	eventId, err := common.EncodePathSegment(eventId)
	if err != nil {
		return nil, err
	}

	workflowId, err = common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, eventId, WorkflowPath, workflowId, ReflowPath),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) Reflow(request reflows.ReflowRequest) (*common.MetadataResponse, error) {
	return c.ReflowWithContext(context.Background(), request)
}

func (c *Client) ReflowWithContext(
	ctx context.Context,
	request reflows.ReflowRequest,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, ReflowPath),
		auth,
		request,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// GetSubjectEvents sends GET /workflows/events/subject/{subjectId}: get subject events.
//
// subjectId is the {subjectId} path value: the event identifier.
//
// It returns the events.SubjectEventsResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." subjectId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) GetSubjectEvents(subjectId string) (*events.SubjectEventsResponse, error) {
	return c.GetSubjectEventsWithContext(context.Background(), subjectId)
}

// GetSubjectEventsWithContext is GetSubjectEvents with a caller-supplied context.
func (c *Client) GetSubjectEventsWithContext(
	ctx context.Context,
	subjectId string,
) (*events.SubjectEventsResponse, error) {
	subjectId, err := common.EncodePathSegment(subjectId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response events.SubjectEventsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, SubjectPath, subjectId),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ReflowBySubject sends POST /workflows/events/subject/{subjectId}/reflow: reflow by subject.
//
// subjectId is the {subjectId} path value: the subject identifier (for example, a payment ID or a
// dispute ID). The events associated with these subjects will be reflowed.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." subjectId
// returns a CheckoutArgumentError and sends no request.
func (c *Client) ReflowBySubject(subjectId string) (*common.MetadataResponse, error) {
	return c.ReflowBySubjectWithContext(context.Background(), subjectId)
}

// ReflowBySubjectWithContext is ReflowBySubject with a caller-supplied context.
func (c *Client) ReflowBySubjectWithContext(
	ctx context.Context,
	subjectId string,
) (*common.MetadataResponse, error) {
	subjectId, err := common.EncodePathSegment(subjectId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, SubjectPath, subjectId, ReflowPath),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// ReflowBySubjectAndWorkflow sends POST
// /workflows/events/subject/{subjectId}/workflow/{workflowId}/reflow: reflow by subject and
// workflow.
//
// subjectId is the {subjectId} path value: the subject identifier (for example, a payment ID or a
// dispute ID). The events associated with these subjects will be reflowed. workflowId is the {workflowId} path value: the identifier of the workflow whose
// actions you want to trigger.
//
// It returns the common.MetadataResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." subjectId or
// workflowId returns a CheckoutArgumentError and sends no request.
func (c *Client) ReflowBySubjectAndWorkflow(subjectId, workflowId string) (*common.MetadataResponse, error) {
	return c.ReflowBySubjectAndWorkflowWithContext(context.Background(), subjectId, workflowId)
}

// ReflowBySubjectAndWorkflowWithContext is ReflowBySubjectAndWorkflow with a caller-supplied context.
func (c *Client) ReflowBySubjectAndWorkflowWithContext(
	ctx context.Context,
	subjectId string,
	workflowId string,
) (*common.MetadataResponse, error) {
	subjectId, err := common.EncodePathSegment(subjectId)
	if err != nil {
		return nil, err
	}

	workflowId, err = common.EncodePathSegment(workflowId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response common.MetadataResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, SubjectPath, subjectId, WorkflowPath, workflowId, ReflowPath),
		auth,
		nil,
		&response,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
