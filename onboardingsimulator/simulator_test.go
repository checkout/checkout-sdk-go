package onboardingsimulator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulatorAvailableRequirementsResponse_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		check   func(*testing.T, *SimulatorAvailableRequirementsResponse, error)
	}{
		{
			name: "when API returns a bare JSON array then Data is populated",
			payload: `[
				{"field":"individual.identification.document","type":"string"},
				{"field":"company.legal_name","type":"string"}
			]`,
			check: func(t *testing.T, r *SimulatorAvailableRequirementsResponse, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, r.Data)
				assert.Len(t, r.Data, 2)
				assert.Equal(t, "individual.identification.document", r.Data[0].Field)
				assert.Equal(t, "string", r.Data[0].Type)
				assert.Equal(t, "company.legal_name", r.Data[1].Field)
			},
		},
		{
			name:    "when API returns an empty array then Data is an empty slice",
			payload: `[]`,
			check: func(t *testing.T, r *SimulatorAvailableRequirementsResponse, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, r.Data)
				assert.Len(t, r.Data, 0)
			},
		},
		{
			name:    "when payload is not a JSON array then error is returned",
			payload: `{"field":"x"}`,
			check: func(t *testing.T, r *SimulatorAvailableRequirementsResponse, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var response SimulatorAvailableRequirementsResponse
			err := json.Unmarshal([]byte(tc.payload), &response)
			tc.check(t, &response, err)
		})
	}
}

func TestSimulatorScenariosResponse_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name    string
		payload string
		check   func(*testing.T, *SimulatorScenariosResponse, error)
	}{
		{
			name: "when API returns a bare JSON array then Data is populated",
			payload: `[
				{"id":"go_active","name":"Go Active","description":"Transitions to active.","action":"set_status","status":"active","requirements_due":[]},
				{"id":"add_requirement","name":"Add Requirement","action":"set_requirements_due","status":"","requirements_due":["individual.identification.document"]}
			]`,
			check: func(t *testing.T, r *SimulatorScenariosResponse, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, r.Data)
				assert.Len(t, r.Data, 2)
				assert.Equal(t, "go_active", r.Data[0].Id)
				assert.Equal(t, "Go Active", r.Data[0].Name)
				assert.Equal(t, "set_status", r.Data[0].Action)
				assert.Equal(t, "active", r.Data[0].Status)
				assert.Equal(t, "add_requirement", r.Data[1].Id)
				assert.Equal(t, []string{"individual.identification.document"}, r.Data[1].RequirementsDue)
			},
		},
		{
			name:    "when API returns an empty array then Data is an empty slice",
			payload: `[]`,
			check: func(t *testing.T, r *SimulatorScenariosResponse, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, r.Data)
				assert.Len(t, r.Data, 0)
			},
		},
		{
			name:    "when payload is not a JSON array then error is returned",
			payload: `{"id":"x"}`,
			check: func(t *testing.T, r *SimulatorScenariosResponse, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var response SimulatorScenariosResponse
			err := json.Unmarshal([]byte(tc.payload), &response)
			tc.check(t, &response, err)
		})
	}
}
