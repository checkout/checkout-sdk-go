package onboardingsimulator

import (
	"context"

	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
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

func (c *Client) SetRequirementsDue(
	entityId string,
	request SimulatorSetRequirementsDueRequest,
) (*SimulatorSetRequirementsDueResponse, error) {
	return c.SetRequirementsDueWithContext(context.Background(), entityId, request)
}

func (c *Client) SetRequirementsDueWithContext(
	ctx context.Context,
	entityId string,
	request SimulatorSetRequirementsDueRequest,
) (*SimulatorSetRequirementsDueResponse, error) {
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

func (c *Client) RunScenario(entityId, scenarioId string) (*SimulatorRunScenarioResponse, error) {
	return c.RunScenarioWithContext(context.Background(), entityId, scenarioId)
}

func (c *Client) RunScenarioWithContext(
	ctx context.Context,
	entityId, scenarioId string,
) (*SimulatorRunScenarioResponse, error) {
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

func (c *Client) SetEntityStatus(
	entityId string,
	request SimulatorSetStatusRequest,
) (*SimulatorSetStatusResponse, error) {
	return c.SetEntityStatusWithContext(context.Background(), entityId, request)
}

func (c *Client) SetEntityStatusWithContext(
	ctx context.Context,
	entityId string,
	request SimulatorSetStatusRequest,
) (*SimulatorSetStatusResponse, error) {
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
		&response,
	)
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
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
