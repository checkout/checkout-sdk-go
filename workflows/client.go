package workflows

import (
	"context"

	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/workflows/actions"
	"github.com/checkout/checkout-sdk-go/workflows/conditions"
	"github.com/checkout/checkout-sdk-go/workflows/events"
	"github.com/checkout/checkout-sdk-go/workflows/reflows"
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
		&response,
	)
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

func (c *Client) GetWorkflow(workflowId string) (*GetWorkflowResponse, error) {
	return c.GetWorkflowWithContext(context.Background(), workflowId)
}

func (c *Client) GetWorkflowWithContext(
	ctx context.Context,
	workflowId string,
) (*GetWorkflowResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetWorkflowResponse
	err = c.apiClient.GetWithContext(
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

func (c *Client) RemoveWorkflow(workflowId string) (*common.MetadataResponse, error) {
	return c.RemoveWorkflowWithContext(context.Background(), workflowId)
}

func (c *Client) RemoveWorkflowWithContext(
	ctx context.Context,
	workflowId string,
) (*common.MetadataResponse, error) {
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

func (c *Client) UpdateWorkflow(
	workflowId string,
	request UpdateWorkflowRequest,
) (*UpdateWorkflowResponse, error) {
	return c.UpdateWorkflowWithContext(context.Background(), workflowId, request)
}

func (c *Client) UpdateWorkflowWithContext(
	ctx context.Context,
	workflowId string,
	request UpdateWorkflowRequest,
) (*UpdateWorkflowResponse, error) {
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

func (c *Client) AddWorkflowAction(
	workflowId string,
	request actions.ActionsRequest,
) (*common.IdResponse, error) {
	return c.AddWorkflowActionWithContext(context.Background(), workflowId, request)
}

func (c *Client) AddWorkflowActionWithContext(
	ctx context.Context,
	workflowId string,
	request actions.ActionsRequest,
) (*common.IdResponse, error) {
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

func (c *Client) UpdateWorkflowAction(
	workflowId string,
	actionId string,
	request actions.ActionsRequest,
) (*common.MetadataResponse, error) {
	return c.UpdateWorkflowActionWithContext(context.Background(), workflowId, actionId, request)
}

func (c *Client) UpdateWorkflowActionWithContext(
	ctx context.Context,
	workflowId string,
	actionId string,
	request actions.ActionsRequest,
) (*common.MetadataResponse, error) {
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

func (c *Client) RemoveWorkflowAction(workflowId, actionId string) (*common.MetadataResponse, error) {
	return c.RemoveWorkflowActionWithContext(context.Background(), workflowId, actionId)
}

func (c *Client) RemoveWorkflowActionWithContext(
	ctx context.Context,
	workflowId string,
	actionId string,
) (*common.MetadataResponse, error) {
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

func (c *Client) AddWorkflowCondition(
	workflowId string,
	request conditions.ConditionsRequest,
) (*common.IdResponse, error) {
	return c.AddWorkflowConditionWithContext(context.Background(), workflowId, request)
}

func (c *Client) AddWorkflowConditionWithContext(
	ctx context.Context,
	workflowId string,
	request conditions.ConditionsRequest,
) (*common.IdResponse, error) {
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

func (c *Client) UpdateWorkflowCondition(
	workflowId string,
	conditionId string,
	request conditions.ConditionsRequest,
) (*common.MetadataResponse, error) {
	return c.UpdateWorkflowConditionWithContext(context.Background(), workflowId, conditionId, request)
}

func (c *Client) UpdateWorkflowConditionWithContext(
	ctx context.Context,
	workflowId string,
	conditionId string,
	request conditions.ConditionsRequest,
) (*common.MetadataResponse, error) {
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

func (c *Client) RemoveWorkflowCondition(workflowId, conditionId string) (*common.MetadataResponse, error) {
	return c.RemoveWorkflowConditionWithContext(context.Background(), workflowId, conditionId)
}

func (c *Client) RemoveWorkflowConditionWithContext(
	ctx context.Context,
	workflowId string,
	conditionId string,
) (*common.MetadataResponse, error) {
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

func (c *Client) TestWorkflow(workflowId string, request events.EventTypesRequest) (*common.MetadataResponse, error) {
	return c.TestWorkflowWithContext(context.Background(), workflowId, request)
}

func (c *Client) TestWorkflowWithContext(
	ctx context.Context,
	workflowId string,
	request events.EventTypesRequest,
) (*common.MetadataResponse, error) {
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
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetEvent(eventId string) (*events.EventResponse, error) {
	return c.GetEventWithContext(context.Background(), eventId)
}

func (c *Client) GetEventWithContext(
	ctx context.Context,
	eventId string,
) (*events.EventResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response events.EventResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, eventId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) GetActionInvocations(eventId, actionId string) (*actions.ActionInvocationsResponse, error) {
	return c.GetActionInvocationsWithContext(context.Background(), eventId, actionId)
}

func (c *Client) GetActionInvocationsWithContext(
	ctx context.Context,
	eventId string,
	actionId string,
) (*actions.ActionInvocationsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response actions.ActionInvocationsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, eventId, ActionsPath, actionId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ReflowByEvent(eventId string) (*common.MetadataResponse, error) {
	return c.ReflowByEventWithContext(context.Background(), eventId)
}

func (c *Client) ReflowByEventWithContext(
	ctx context.Context,
	eventId string,
) (*common.MetadataResponse, error) {
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

func (c *Client) ReflowByEventAndWorkflow(eventId, workflowId string) (*common.MetadataResponse, error) {
	return c.ReflowByEventAndWorkflowWithContext(context.Background(), eventId, workflowId)
}

func (c *Client) ReflowByEventAndWorkflowWithContext(
	ctx context.Context,
	eventId string,
	workflowId string,
) (*common.MetadataResponse, error) {
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

func (c *Client) GetSubjectEvents(subjectId string) (*events.SubjectEventsResponse, error) {
	return c.GetSubjectEventsWithContext(context.Background(), subjectId)
}

func (c *Client) GetSubjectEventsWithContext(
	ctx context.Context,
	subjectId string,
) (*events.SubjectEventsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}
	var response events.SubjectEventsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(WorkflowsPath, EventsPath, SubjectPath, subjectId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ReflowBySubject(subjectId string) (*common.MetadataResponse, error) {
	return c.ReflowBySubjectWithContext(context.Background(), subjectId)
}

func (c *Client) ReflowBySubjectWithContext(
	ctx context.Context,
	subjectId string,
) (*common.MetadataResponse, error) {
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

func (c *Client) ReflowBySubjectAndWorkflow(subjectId, workflowId string) (*common.MetadataResponse, error) {
	return c.ReflowBySubjectAndWorkflowWithContext(context.Background(), subjectId, workflowId)
}

func (c *Client) ReflowBySubjectAndWorkflowWithContext(
	ctx context.Context,
	subjectId string,
	workflowId string,
) (*common.MetadataResponse, error) {
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
