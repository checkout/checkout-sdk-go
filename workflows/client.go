package workflows

import (
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetWorkflowsResponse
	err = c.apiClient.Get(common.BuildPath(WorkflowsPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CreateWorkflow(request CreateWorkflowRequest) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(common.BuildPath(WorkflowsPath), auth, request, &response, nil)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetWorkflow(workflowId string) (*GetWorkflowResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response GetWorkflowResponse
	err = c.apiClient.Get(common.BuildPath(WorkflowsPath, workflowId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) RemoveWorkflow(workflowId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Delete(common.BuildPath(WorkflowsPath, workflowId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) UpdateWorkflow(workflowId string, request UpdateWorkflowRequest) (*UpdateWorkflowResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response UpdateWorkflowResponse
	err = c.apiClient.Patch(common.BuildPath(WorkflowsPath, workflowId), auth, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) AddWorkflowAction(
	workflowId string,
	request actions.ActionsRequest,
) (*common.IdResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(
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
	workflowId, actionId string,
	request actions.ActionsRequest,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Put(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Delete(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.IdResponse
	err = c.apiClient.Post(
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
	workflowId, conditionId string,
	request conditions.ConditionsRequest,
) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Put(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Delete(
		common.BuildPath(WorkflowsPath, workflowId, ConditionsPath, conditionId),
		auth,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetEventTypes() (*events.EventTypesResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response events.EventTypesResponse
	err = c.apiClient.Get(common.BuildPath(WorkflowsPath, EventTypesPath), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetEvent(eventId string) (*events.EventResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response events.EventResponse
	err = c.apiClient.Get(common.BuildPath(WorkflowsPath, EventsPath, eventId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetActionInvocations(eventId, actionId string) (*actions.ActionInvocationsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response actions.ActionInvocationsResponse
	err = c.apiClient.Get(common.BuildPath(WorkflowsPath, EventsPath, eventId, ActionsPath, actionId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ReflowByEvent(eventId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response events.SubjectEventsResponse
	err = c.apiClient.Get(common.BuildPath(WorkflowsPath, EventsPath, SubjectPath, subjectId), auth, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ReflowBySubject(subjectId string) (*common.MetadataResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.SecretKeyOrOauth)
	if err != nil {
		return nil, err
	}

	var response common.MetadataResponse
	err = c.apiClient.Post(
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
