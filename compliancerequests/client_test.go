package compliancerequests

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

func TestGetComplianceRequest(t *testing.T) {
	var (
		paymentId = "pay_fun26akvvjjerahhctaq2uzhu4"
		detailsResponse = GetComplianceRequestResponse{
			HttpMetadata:  mocks.HttpMetadataStatusOk,
			PaymentId:     paymentId,
			Status:        "pending",
			Amount:        "38.23",
			Currency:      "HKD",
			RecipientName: "Jia Tsang",
		}
	)

	cases := []struct {
		name             string
		paymentId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*GetComplianceRequestResponse, error)
	}{
		{
			name:      "when payment id is valid then return compliance request",
			paymentId: paymentId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*GetComplianceRequestResponse)
						*respMapping = detailsResponse
					})
			},
			checker: func(response *GetComplianceRequestResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, detailsResponse.PaymentId, response.PaymentId)
				assert.Equal(t, detailsResponse.Status, response.Status)
				assert.Equal(t, detailsResponse.Amount, response.Amount)
			},
		},
		{
			name:      "when credentials invalid then return error",
			paymentId: paymentId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *GetComplianceRequestResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:      "when payment not found then return error",
			paymentId: "pay_invalid",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
			},
			checker: func(response *GetComplianceRequestResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildComplianceRequestsClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.GetComplianceRequest(tc.paymentId))
		})
	}
}

func TestRespondToComplianceRequest(t *testing.T) {
	var (
		paymentId = "pay_fun26akvvjjerahhctaq2uzhu4"
	)

	cases := []struct {
		name             string
		paymentId        string
		request          RespondToComplianceRequestRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:      "when request is correct then respond to compliance request",
			paymentId: paymentId,
			request:   buildRespondToComplianceRequestRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*common.MetadataResponse)
						*respMapping = common.MetadataResponse{HttpMetadata: mocks.HttpMetadataStatusAccepted}
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:      "when credentials invalid then return error",
			paymentId: paymentId,
			request:   buildRespondToComplianceRequestRequest(),
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
			name:      "when payment not found then return error",
			paymentId: "pay_invalid",
			request:   buildRespondToComplianceRequestRequest(),
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
			apiClient, credentials, config := buildComplianceRequestsClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.RespondToComplianceRequest(tc.paymentId, tc.request))
		})
	}
}

// common methods

func buildComplianceRequestsClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}

func buildRespondToComplianceRequestRequest() RespondToComplianceRequestRequest {
	return RespondToComplianceRequestRequest{
		Fields: ComplianceRespondedFields{
			Sender: []ComplianceRespondedField{
				{Name: "date_of_birth", Value: "2000-01-01", NotAvailable: false},
			},
		},
		Comments: "Responding to compliance request",
	}
}
