package onboardingsimulator

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v3/client"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
)

type Client struct {
	configuration *configuration.Configuration
	apiClient     client.HttpClient
}

func NewClient(
	configuration *configuration.Configuration,
	apiClient client.HttpClient,
) *Client {
	return &Client{
		configuration: configuration,
		apiClient:     apiClient,
	}
}

// SetRequirementsDue sends POST /simulate/entities/{entityId}/requirements-due: set requirements
// due.
//
// entityId is the {entityId} path value: the ID of the entity. request is sent as the JSON body.
//
// It returns the SimulatorSetRequirementsDueResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) SetRequirementsDue(
	entityId string,
	request SimulatorSetRequirementsDueRequest,
) (*SimulatorSetRequirementsDueResponse, error) {
	return c.SetRequirementsDueWithContext(context.Background(), entityId, request)
}

// SetRequirementsDueWithContext is SetRequirementsDue with a caller-supplied context.
func (c *Client) SetRequirementsDueWithContext(
	ctx context.Context,
	entityId string,
	request SimulatorSetRequirementsDueRequest,
) (*SimulatorSetRequirementsDueResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response SimulatorSetRequirementsDueResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(simulatePath, entitiesPath, entityId, requirementsDuePath),
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

// RunScenario sends POST /simulate/entities/{entityId}/scenarios/{scenarioId}: run scenario.
//
// entityId is the {entityId} path value: the ID of the entity. scenarioId is the {scenarioId} path
// value: the ID of the scenario to run.
//
// It returns the SimulatorRunScenarioResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId or
// scenarioId returns a CheckoutArgumentError and sends no request.
func (c *Client) RunScenario(entityId, scenarioId string) (*SimulatorRunScenarioResponse, error) {
	return c.RunScenarioWithContext(context.Background(), entityId, scenarioId)
}

// RunScenarioWithContext is RunScenario with a caller-supplied context.
func (c *Client) RunScenarioWithContext(
	ctx context.Context,
	entityId, scenarioId string,
) (*SimulatorRunScenarioResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	scenarioId, err = common.EncodePathSegment(scenarioId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response SimulatorRunScenarioResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(simulatePath, entitiesPath, entityId, scenariosPath, scenarioId),
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

// SetEntityStatus sends POST /simulate/entities/{entityId}/status: set entity status.
//
// entityId is the {entityId} path value: the ID of the entity. request is sent as the JSON body.
//
// It returns the SimulatorSetStatusResponse, or the error from the API.
//
// Path values are percent-encoded, each as one path segment. An empty, "." or ".." entityId returns
// a CheckoutArgumentError and sends no request.
func (c *Client) SetEntityStatus(
	entityId string,
	request SimulatorSetStatusRequest,
) (*SimulatorSetStatusResponse, error) {
	return c.SetEntityStatusWithContext(context.Background(), entityId, request)
}

// SetEntityStatusWithContext is SetEntityStatus with a caller-supplied context.
func (c *Client) SetEntityStatusWithContext(
	ctx context.Context,
	entityId string,
	request SimulatorSetStatusRequest,
) (*SimulatorSetStatusResponse, error) {
	entityId, err := common.EncodePathSegment(entityId)
	if err != nil {
		return nil, err
	}

	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response SimulatorSetStatusResponse
	err = c.apiClient.PostWithContext(
		ctx,
		common.BuildPath(simulatePath, entitiesPath, entityId, statusPath),
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

func (c *Client) ListAvailableRequirements() (*SimulatorAvailableRequirementsResponse, error) {
	return c.ListAvailableRequirementsWithContext(context.Background())
}

func (c *Client) ListAvailableRequirementsWithContext(
	ctx context.Context,
) (*SimulatorAvailableRequirementsResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response SimulatorAvailableRequirementsResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(simulatePath, requirementsDuePath),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) ListScenarios() (*SimulatorScenariosResponse, error) {
	return c.ListScenariosWithContext(context.Background())
}

func (c *Client) ListScenariosWithContext(
	ctx context.Context,
) (*SimulatorScenariosResponse, error) {
	auth, err := c.configuration.Credentials.GetAuthorization(configuration.OAuth)
	if err != nil {
		return nil, err
	}

	var response SimulatorScenariosResponse
	err = c.apiClient.GetWithContext(
		ctx,
		common.BuildPath(simulatePath, scenariosPath),
		auth,
		nil,
		&response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
