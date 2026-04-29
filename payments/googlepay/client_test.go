package googlepay

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
)

// tests

func TestCreateEnrollment(t *testing.T) {
	var (
		enrollmentResponse = CreateEnrollmentResponse{
			State: Active,
		}
	)

	cases := []struct {
		name             string
		request          CreateEnrollmentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CreateEnrollmentResponse, error)
	}{
		{
			name:    "when request is correct then create enrollment",
			request: buildCreateEnrollmentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CreateEnrollmentResponse)
						*respMapping = enrollmentResponse
					})
			},
			checker: func(response *CreateEnrollmentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, enrollmentResponse.State, response.State)
			},
		},
		{
			name:    "when credentials invalid then return error",
			request: buildCreateEnrollmentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *CreateEnrollmentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:    "when api returns error then propagate error",
			request: buildCreateEnrollmentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusUnprocessableEntity})
			},
			checker: func(response *CreateEnrollmentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildGooglePayClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.CreateEnrollment(tc.request))
		})
	}
}

func TestRegisterDomain(t *testing.T) {
	var (
		entityId = "ent_uzm3uxtssvmuxnyrfdffcyjxeu"
	)

	cases := []struct {
		name             string
		entityId         string
		request          RegisterDomainRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:     "when request is correct then register domain",
			entityId: entityId,
			request:  buildRegisterDomainRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*common.MetadataResponse)
						*respMapping = common.MetadataResponse{HttpMetadata: mocks.HttpMetadataStatusNoContent}
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:     "when credentials invalid then return error",
			entityId: entityId,
			request:  buildRegisterDomainRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:     "when entity not found then return error",
			entityId: "ent_invalid",
			request:  buildRegisterDomainRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
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
			apiClient, credentials, config := buildGooglePayClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.RegisterDomain(tc.entityId, tc.request))
		})
	}
}

func TestGetRegisteredDomains(t *testing.T) {
	var (
		entityId           = "ent_uzm3uxtssvmuxnyrfdffcyjxeu"
		domainListResponse = DomainListResponse{
			Domains: []string{"example.com", "checkout.com"},
		}
	)

	cases := []struct {
		name             string
		entityId         string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*DomainListResponse, error)
	}{
		{
			name:     "when entity id is valid then return registered domains",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*DomainListResponse)
						*respMapping = domainListResponse
					})
			},
			checker: func(response *DomainListResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, domainListResponse.Domains, response.Domains)
				assert.Len(t, response.Domains, 2)
			},
		},
		{
			name:     "when credentials invalid then return error",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *DomainListResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:     "when entity not found then return error",
			entityId: "ent_invalid",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
			},
			checker: func(response *DomainListResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildGooglePayClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.GetRegisteredDomains(tc.entityId))
		})
	}
}

func TestGetEnrollmentState(t *testing.T) {
	var (
		entityId              = "ent_uzm3uxtssvmuxnyrfdffcyjxeu"
		enrollmentStateResponse = EnrollmentStateResponse{
			State: Active,
		}
	)

	cases := []struct {
		name             string
		entityId         string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*EnrollmentStateResponse, error)
	}{
		{
			name:     "when entity id is valid then return enrollment state",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*EnrollmentStateResponse)
						*respMapping = enrollmentStateResponse
					})
			},
			checker: func(response *EnrollmentStateResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, enrollmentStateResponse.State, response.State)
			},
		},
		{
			name:     "when credentials invalid then return error",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *EnrollmentStateResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:     "when entity not found then return error",
			entityId: "ent_invalid",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
			},
			checker: func(response *EnrollmentStateResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildGooglePayClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.GetEnrollmentState(tc.entityId))
		})
	}
}

// common methods

func buildGooglePayClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}

func buildCreateEnrollmentRequest() CreateEnrollmentRequest {
	return CreateEnrollmentRequest{
		EntityId:             "ent_uzm3uxtssvmuxnyrfdffcyjxeu",
		EmailAddress:         "test@gmail.com",
		AcceptTermsOfService: true,
	}
}

func buildRegisterDomainRequest() RegisterDomainRequest {
	return RegisterDomainRequest{
		WebDomain: "some.example.com",
	}
}
