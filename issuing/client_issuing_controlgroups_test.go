package issuing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"

	controlgroups "github.com/checkout/checkout-sdk-go/v2/issuing/controlgroups"
)

// # tests

func TestGetControlGroups(t *testing.T) {
	var (
		response = controlgroups.ControlGroupsResponse{
			HttpMetadata:  mocks.HttpMetadataStatusOk,
			ControlGroups: []controlgroups.ControlGroupResponse{{Id: "cgr_test_12345"}},
		}
	)

	cases := []struct {
		name             string
		query            controlgroups.ControlGroupsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*controlgroups.ControlGroupsResponse, error)
	}{
		{
			name:  "when query is correct then should return 200",
			query: buildControlGroupsQuery(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*controlgroups.ControlGroupsResponse)
						*respMapping = response
					})
			},
			checker: assertControlGroupsSuccess(t, &response),
		},
		{
			name:  "when credentials are invalid then return authorization error",
			query: controlgroups.ControlGroupsQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *controlgroups.ControlGroupsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:  "when target not found then return 404",
			query: controlgroups.ControlGroupsQuery{TargetId: "crd_not_found"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *controlgroups.ControlGroupsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.GetControlGroups(tc.query))
		})
	}
}

func TestCreateControlGroup(t *testing.T) {
	var (
		response = controlgroups.ControlGroupResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "cgr_test_12345",
			Description:  "Test control group",
		}
	)

	cases := []struct {
		name             string
		request          controlgroups.CreateControlGroupRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*controlgroups.ControlGroupResponse, error)
	}{
		{
			name:    "when request is correct then should return 201",
			request: buildCreateControlGroupRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*controlgroups.ControlGroupResponse)
						*respMapping = response
					})
			},
			checker: assertControlGroupSuccess(t, &response),
		},
		{
			name:    "when credentials are invalid then return authorization error",
			request: controlgroups.CreateControlGroupRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *controlgroups.ControlGroupResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request is invalid then return 422",
			request: controlgroups.CreateControlGroupRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable",
						Data: &errors.ErrorDetails{
							ErrorType:  "request_invalid",
							ErrorCodes: []string{"target_id_required"},
						},
					})
			},
			checker: func(response *controlgroups.ControlGroupResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.CreateControlGroup(tc.request))
		})
	}
}

func TestGetControlGroupDetails(t *testing.T) {
	var (
		response = controlgroups.ControlGroupResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "cgr_test_12345",
			Description:  "Test control group",
		}
	)

	cases := []struct {
		name             string
		controlGroupId   string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*controlgroups.ControlGroupResponse, error)
	}{
		{
			name:           "when request is correct then should return 200",
			controlGroupId: "cgr_test_12345",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*controlgroups.ControlGroupResponse)
						*respMapping = response
					})
			},
			checker: assertControlGroupSuccess(t, &response),
		},
		{
			name:           "when control group not found then return 404",
			controlGroupId: "cgr_not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *controlgroups.ControlGroupResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.GetControlGroupDetails(tc.controlGroupId))
		})
	}
}

func TestRemoveControlGroup(t *testing.T) {
	var (
		response = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
		}
	)

	cases := []struct {
		name             string
		controlGroupId   string
		getAuthorization func(*mock.Mock) mock.Call
		apiDelete        func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:           "when request is correct then should return 200",
			controlGroupId: "cgr_test_12345",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.MetadataResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:           "when control group not found then return 404",
			controlGroupId: "cgr_not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiDelete(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.RemoveControlGroup(tc.controlGroupId))
		})
	}
}

// # common methods

func buildControlGroupsQuery() controlgroups.ControlGroupsQuery {
	return controlgroups.ControlGroupsQuery{
		TargetId: "crd_fa6psq242dcd6fdn5gifcq1491",
	}
}

func buildCreateControlGroupRequest() controlgroups.CreateControlGroupRequest {
	return controlgroups.CreateControlGroupRequest{
		TargetId:    "crd_fa6psq242dcd6fdn5gifcq1491",
		FailIf:      controlgroups.AllFail,
		Description: "Test control group",
		Controls: []controlgroups.ControlGroupControl{
			{
				ControlType: controlgroups.MccLimitControlType,
				Description: "Block MCC list",
				MccLimit: &controlgroups.MccGroupLimit{
					Type:    controlgroups.Block,
					MccList: []string{"5411", "5422"},
				},
			},
		},
	}
}

func assertControlGroupSuccess(t *testing.T, expected *controlgroups.ControlGroupResponse) func(*controlgroups.ControlGroupResponse, error) {
	return func(response *controlgroups.ControlGroupResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, expected.Id, response.Id)
		assert.Equal(t, expected.Description, response.Description)
	}
}

func assertControlGroupsSuccess(t *testing.T, expected *controlgroups.ControlGroupsResponse) func(*controlgroups.ControlGroupsResponse, error) {
	return func(response *controlgroups.ControlGroupsResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
		assert.NotNil(t, response.ControlGroups)
		assert.Equal(t, len(expected.ControlGroups), len(response.ControlGroups))
	}
}
