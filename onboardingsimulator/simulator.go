package onboardingsimulator

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

const (
	simulatePath       = "simulate"
	entitiesPath       = "entities"
	requirementsDuePath = "requirements-due"
	scenariosPath      = "scenarios"
	statusPath         = "status"
)

type SimulatorEntityStatus string

const (
	Draft           SimulatorEntityStatus = "draft"
	RequirementsDue SimulatorEntityStatus = "requirements_due"
	Pending         SimulatorEntityStatus = "pending"
	Active          SimulatorEntityStatus = "active"
	Restricted      SimulatorEntityStatus = "restricted"
	Rejected        SimulatorEntityStatus = "rejected"
	Inactive        SimulatorEntityStatus = "inactive"
)

type (
	SimulatorAvailableRequirement struct {
		Field string `json:"field,omitempty"`
		Type  string `json:"type,omitempty"`
	}

	SimulatorScenario struct {
		Id              string   `json:"id,omitempty"`
		Name            string   `json:"name,omitempty"`
		Description     string   `json:"description,omitempty"`
		Action          string   `json:"action,omitempty"`
		Status          string   `json:"status,omitempty"`
		RequirementsDue []string `json:"requirements_due,omitempty"`
	}
)

type (
	SimulatorSetRequirementsDueRequest struct {
		Fields []string `json:"fields"`
	}

	SimulatorSetStatusRequest struct {
		Status SimulatorEntityStatus `json:"status"`
	}
)

type (
	SimulatorSetRequirementsDueResponse struct {
		HttpMetadata    common.HttpMetadata
		EntityId        string   `json:"entity_id,omitempty"`
		PreviousStatus  string   `json:"previous_status,omitempty"`
		CurrentStatus   string   `json:"current_status,omitempty"`
		RequirementsDue []string `json:"requirements_due,omitempty"`
	}

	SimulatorRunScenarioResponse struct {
		HttpMetadata    common.HttpMetadata
		EntityId        string   `json:"entity_id,omitempty"`
		ScenarioId      string   `json:"scenario_id,omitempty"`
		ScenarioName    string   `json:"scenario_name,omitempty"`
		PreviousStatus  string   `json:"previous_status,omitempty"`
		CurrentStatus   string   `json:"current_status,omitempty"`
		RequirementsDue []string `json:"requirements_due,omitempty"`
	}

	SimulatorSetStatusResponse struct {
		HttpMetadata   common.HttpMetadata
		EntityId       string `json:"entity_id,omitempty"`
		PreviousStatus string `json:"previous_status,omitempty"`
		CurrentStatus  string `json:"current_status,omitempty"`
	}

	SimulatorAvailableRequirementsResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []SimulatorAvailableRequirement
	}

	SimulatorScenariosResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []SimulatorScenario
	}
)

// UnmarshalJSON parses the JSON array returned by GET /simulate/requirements-due
// into the Data slice (the API returns a bare array, not an object).
func (r *SimulatorAvailableRequirementsResponse) UnmarshalJSON(data []byte) error {
	var items []SimulatorAvailableRequirement
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}
	r.Data = items
	return nil
}

// UnmarshalJSON parses the JSON array returned by GET /simulate/scenarios
// into the Data slice (the API returns a bare array, not an object).
func (r *SimulatorScenariosResponse) UnmarshalJSON(data []byte) error {
	var items []SimulatorScenario
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}
	r.Data = items
	return nil
}
