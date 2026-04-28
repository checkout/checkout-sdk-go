package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"

	controlprofiles "github.com/checkout/checkout-sdk-go/v2/issuing/controlprofiles"
)

// # tests

func TestGetAllControlProfiles(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		query   controlprofiles.ControlProfilesQuery
		checker func(*controlprofiles.ControlProfilesResponse, error)
	}{
		{
			name:  "when query is correct then should return 200",
			query: controlProfilesQuery(),
			checker: func(response *controlprofiles.ControlProfilesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.ControlProfiles)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetAllControlProfiles(tc.query))
		})
	}
}

func TestCreateControlProfile(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		checker func(*controlprofiles.ControlProfileResponse, error)
	}{
		{
			name: "when request is correct then should return 201",
			checker: func(response *controlprofiles.ControlProfileResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Name)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(controlProfileRequest(t), nil)
		})
	}
}

func TestGetControlProfileDetails(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name             string
		controlProfileId string
		checker          func(*controlprofiles.ControlProfileResponse, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: controlProfileResponse.Id,
			checker: func(response *controlprofiles.ControlProfileResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, controlProfileResponse.Id, response.Id)
				assert.NotNil(t, response.Name)
			},
		},
		{
			name:             "when control profile id not found then return error",
			controlProfileId: "cp_not_found",
			checker: func(response *controlprofiles.ControlProfileResponse, err error) {
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
			tc.checker(client.GetControlProfileDetails(tc.controlProfileId))
		})
	}
}

func TestUpdateControlProfile(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	request := controlprofiles.ControlProfileRequest{
		Name: "Updated Profile Name",
	}
	cases := []struct {
		name             string
		controlProfileId string
		request          controlprofiles.ControlProfileRequest
		checker          func(*controlprofiles.ControlProfileResponse, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: controlProfileResponse.Id,
			request:          request,
			checker: func(response *controlprofiles.ControlProfileResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, request.Name, response.Name)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateControlProfile(tc.controlProfileId, tc.request))
		})
	}
}

func TestRemoveControlProfile(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name             string
		controlProfileId string
		checker          func(interface{}, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: controlProfileResponse.Id,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RemoveControlProfile(tc.controlProfileId))
		})
	}
}

func TestAddTargetToControlProfile(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name             string
		controlProfileId string
		targetId         string
		checker          func(interface{}, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: controlProfileResponse.Id,
			targetId:         virtualCardId,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.AddTargetToControlProfile(tc.controlProfileId, tc.targetId))
		})
	}
}

func TestRemoveTargetFromControlProfile(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name             string
		controlProfileId string
		targetId         string
		checker          func(interface{}, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: controlProfileResponse.Id,
			targetId:         virtualCardId,
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RemoveTargetFromControlProfile(tc.controlProfileId, tc.targetId))
		})
	}
}

// # common methods

func controlProfilesQuery() controlprofiles.ControlProfilesQuery {
	return controlprofiles.ControlProfilesQuery{
		TargetId: virtualCardId,
	}
}

func controlProfileRequest(t *testing.T) *controlprofiles.ControlProfileResponse {
	request := controlprofiles.ControlProfileRequest{
		Name: "Default Control Profile",
	}

	response, err := buildIssuingClientApi().Issuing.CreateControlProfile(request)
	if err != nil {
		t.Fatalf("error creating control profile: %s", err.Error())
	}

	return response
}

func assertControlProfileResponse(t *testing.T, response *controlprofiles.ControlProfileResponse) {
	assert.NotNil(t, response)
	assert.NotNil(t, response.Id)
	assert.NotNil(t, response.Name)
}
