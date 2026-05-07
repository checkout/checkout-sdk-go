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

	controlprofiles "github.com/checkout/checkout-sdk-go/v2/issuing/controlprofiles"
)

// # tests

func TestGetAllControlProfiles(t *testing.T) {
	var (
		response = controlprofiles.ControlProfilesResponse{
			HttpMetadata:    mocks.HttpMetadataStatusOk,
			ControlProfiles: []controlprofiles.ControlProfileResponse{{Id: "cp_test_12345", Name: "Default Profile"}},
		}
	)

	cases := []struct {
		name             string
		query            controlprofiles.ControlProfilesQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*controlprofiles.ControlProfilesResponse, error)
	}{
		{
			name:  "when query is correct then should return 200",
			query: buildControlProfilesQuery(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*controlprofiles.ControlProfilesResponse)
						*respMapping = response
					})
			},
			checker: assertControlProfilesSuccess(t, &response),
		},
		{
			name:  "when credentials are invalid then return authorization error",
			query: controlprofiles.ControlProfilesQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *controlprofiles.ControlProfilesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:  "when server has problems then return 500",
			query: controlprofiles.ControlProfilesQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusInternalServerError,
						Status:     "500 Internal Server Error",
					})
			},
			checker: func(response *controlprofiles.ControlProfilesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusInternalServerError, chkErr.StatusCode)
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

			tc.checker(client.GetAllControlProfiles(tc.query))
		})
	}
}

func TestCreateControlProfile(t *testing.T) {
	var (
		response = controlprofiles.ControlProfileResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "cp_test_12345",
			Name:         "Default Profile",
		}
	)

	cases := []struct {
		name             string
		request          controlprofiles.ControlProfileRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*controlprofiles.ControlProfileResponse, error)
	}{
		{
			name:    "when request is correct then should return 201",
			request: buildControlProfileRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*controlprofiles.ControlProfileResponse)
						*respMapping = response
					})
			},
			checker: assertControlProfileSuccess(t, &response),
		},
		{
			name:    "when request is invalid then return 422",
			request: controlprofiles.ControlProfileRequest{},
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
							ErrorCodes: []string{"name_required"},
						},
					})
			},
			checker: func(response *controlprofiles.ControlProfileResponse, err error) {
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

			tc.checker(client.CreateControlProfile(tc.request))
		})
	}
}

func TestGetControlProfileDetails(t *testing.T) {
	var (
		response = controlprofiles.ControlProfileResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "cp_test_12345",
			Name:         "Default Profile",
		}
	)

	cases := []struct {
		name             string
		controlProfileId string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*controlprofiles.ControlProfileResponse, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: "cp_test_12345",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*controlprofiles.ControlProfileResponse)
						*respMapping = response
					})
			},
			checker: assertControlProfileSuccess(t, &response),
		},
		{
			name:             "when control profile not found then return 404",
			controlProfileId: "cp_not_found",
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
			checker: func(response *controlprofiles.ControlProfileResponse, err error) {
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

			tc.checker(client.GetControlProfileDetails(tc.controlProfileId))
		})
	}
}

func TestUpdateControlProfile(t *testing.T) {
	var (
		response = controlprofiles.ControlProfileResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "cp_test_12345",
			Name:         "Updated Profile",
		}
	)

	cases := []struct {
		name             string
		controlProfileId string
		request          controlprofiles.ControlProfileRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPatch         func(*mock.Mock) mock.Call
		checker          func(*controlprofiles.ControlProfileResponse, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: "cp_test_12345",
			request:          controlprofiles.ControlProfileRequest{Name: "Updated Profile"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*controlprofiles.ControlProfileResponse)
						*respMapping = response
					})
			},
			checker: assertControlProfileSuccess(t, &response),
		},
		{
			name:             "when control profile not found then return 404",
			controlProfileId: "cp_not_found",
			request:          controlprofiles.ControlProfileRequest{Name: "Updated Profile"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *controlprofiles.ControlProfileResponse, err error) {
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
			tc.apiPatch(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.UpdateControlProfile(tc.controlProfileId, tc.request))
		})
	}
}

func TestRemoveControlProfile(t *testing.T) {
	var (
		response = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
		}
	)

	cases := []struct {
		name             string
		controlProfileId string
		getAuthorization func(*mock.Mock) mock.Call
		apiDelete        func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: "cp_test_12345",
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
			name:             "when control profile not found then return 404",
			controlProfileId: "cp_not_found",
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

			tc.checker(client.RemoveControlProfile(tc.controlProfileId))
		})
	}
}

func TestAddTargetToControlProfile(t *testing.T) {
	var (
		response = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
		}
	)

	cases := []struct {
		name             string
		controlProfileId string
		targetId         string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: "cp_test_12345",
			targetId:         "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*common.MetadataResponse)
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
			name:             "when control profile not found then return 404",
			controlProfileId: "cp_not_found",
			targetId:         "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			tc.apiPost(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.AddTargetToControlProfile(tc.controlProfileId, tc.targetId))
		})
	}
}

func TestRemoveTargetFromControlProfile(t *testing.T) {
	var (
		response = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
		}
	)

	cases := []struct {
		name             string
		controlProfileId string
		targetId         string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:             "when request is correct then should return 200",
			controlProfileId: "cp_test_12345",
			targetId:         "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*common.MetadataResponse)
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
			name:             "when control profile not found then return 404",
			controlProfileId: "cp_not_found",
			targetId:         "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			tc.apiPost(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.RemoveTargetFromControlProfile(tc.controlProfileId, tc.targetId))
		})
	}
}

// # common methods

func buildControlProfilesQuery() controlprofiles.ControlProfilesQuery {
	return controlprofiles.ControlProfilesQuery{
		TargetId: "crd_fa6psq242dcd6fdn5gifcq1491",
	}
}

func buildControlProfileRequest() controlprofiles.ControlProfileRequest {
	return controlprofiles.ControlProfileRequest{
		Name: "Default Profile",
	}
}

func assertControlProfileSuccess(t *testing.T, expected *controlprofiles.ControlProfileResponse) func(*controlprofiles.ControlProfileResponse, error) {
	return func(response *controlprofiles.ControlProfileResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, expected.Id, response.Id)
		assert.Equal(t, expected.Name, response.Name)
	}
}

func assertControlProfilesSuccess(t *testing.T, expected *controlprofiles.ControlProfilesResponse) func(*controlprofiles.ControlProfilesResponse, error) {
	return func(response *controlprofiles.ControlProfilesResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
		assert.NotNil(t, response.ControlProfiles)
		assert.Equal(t, len(expected.ControlProfiles), len(response.ControlProfiles))
	}
}
