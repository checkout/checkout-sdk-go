package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"

	controlgroups "github.com/checkout/checkout-sdk-go/v2/issuing/controlgroups"
)

// # tests

func TestGetControlGroups(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		query   controlgroups.ControlGroupsQuery
		checker func(*controlgroups.ControlGroupsResponse, error)
	}{
		{
			name:  "when query is correct then should return 200",
			query: controlGroupsQuery(),
			checker: func(response *controlgroups.ControlGroupsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.ControlGroups)
			},
		},
		{
			name:  "when target id not found then return error",
			query: controlgroups.ControlGroupsQuery{TargetId: "crd_not_found"},
			checker: func(response *controlgroups.ControlGroupsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetControlGroups(tc.query))
		})
	}
}

func TestCreateControlGroup(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		checker func(*controlgroups.ControlGroupResponse, error)
	}{
		{
			name: "when request is correct then should return 201",
			checker: func(response *controlgroups.ControlGroupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(controlGroupRequest(t), nil)
		})
	}
}

func TestGetControlGroupDetails(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name           string
		controlGroupId string
		checker        func(*controlgroups.ControlGroupResponse, error)
	}{
		{
			name:           "when request is correct then should return 200",
			controlGroupId: controlGroupResponse.Id,
			checker: func(response *controlgroups.ControlGroupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, controlGroupResponse.Id, response.Id)
				assert.NotNil(t, response.Description)
				assert.NotNil(t, response.Controls)
			},
		},
		{
			name:           "when control group id not found then return error",
			controlGroupId: "cgr_not_found",
			checker: func(response *controlgroups.ControlGroupResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetControlGroupDetails(tc.controlGroupId))
		})
	}
}

func TestRemoveControlGroup(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name           string
		controlGroupId string
		checker        func(interface{}, error)
	}{
		{
			name:           "when request is correct then should return 200",
			controlGroupId: controlGroupResponse.Id,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RemoveControlGroup(tc.controlGroupId))
		})
	}
}

// # common methods

func controlGroupsQuery() controlgroups.ControlGroupsQuery {
	return controlgroups.ControlGroupsQuery{
		TargetId: virtualCardId,
	}
}

func controlGroupRequest(t *testing.T) *controlgroups.ControlGroupResponse {
	request := controlgroups.CreateControlGroupRequest{
		TargetId:    virtualCardId,
		FailIf:      controlgroups.AllFail,
		Description: "Max spend of 500€ per week for restaurants",
		Controls: []controlgroups.ControlGroupControl{
			{
				ControlType: controlgroups.MccLimitControlType,
				Description: "Block restaurant MCCs",
				MccLimit: &controlgroups.MccGroupLimit{
					Type:    controlgroups.Block,
					MccList: []string{"5812", "5814"},
				},
			},
		},
	}

	response, err := buildIssuingClientApi().Issuing.CreateControlGroup(request)
	if err != nil {
		t.Fatalf("error creating control group: %s", err.Error())
	}

	return response
}
