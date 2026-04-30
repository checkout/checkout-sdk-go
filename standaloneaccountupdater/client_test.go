package standaloneaccountupdater

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
)

// tests

func TestGetUpdatedCardCredentials(t *testing.T) {
	var (
		credentialsResponse = GetUpdatedCardCredentialsResponse{
			HttpMetadata:        mocks.HttpMetadataStatusOk,
			AccountUpdateStatus: CardUpdated,
			Card: &AccountUpdaterCardDetails{
				Bin:         "424242",
				Last4:       "4242",
				ExpiryMonth: 12,
				ExpiryYear:  2027,
				Fingerprint: "fngrprnt_test123",
			},
		}
	)

	cases := []struct {
		name             string
		request          GetUpdatedCardCredentialsRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*GetUpdatedCardCredentialsResponse, error)
	}{
		{
			name:    "when source is a valid card then return updated card credentials",
			request: buildGetUpdatedCardCredentialsWithCard(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*GetUpdatedCardCredentialsResponse)
						*respMapping = credentialsResponse
					})
			},
			checker: func(response *GetUpdatedCardCredentialsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, credentialsResponse.AccountUpdateStatus, response.AccountUpdateStatus)
				assert.NotNil(t, response.Card)
				assert.Equal(t, credentialsResponse.Card.Last4, response.Card.Last4)
				assert.Equal(t, credentialsResponse.Card.ExpiryYear, response.Card.ExpiryYear)
			},
		},
		{
			name:    "when credentials invalid then return error",
			request: buildGetUpdatedCardCredentialsWithCard(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *GetUpdatedCardCredentialsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:    "when api returns error then propagate error",
			request: buildGetUpdatedCardCredentialsWithCard(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusUnprocessableEntity})
			},
			checker: func(response *GetUpdatedCardCredentialsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildStandaloneAccountUpdaterClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.GetUpdatedCardCredentials(tc.request))
		})
	}
}

// common methods

func buildStandaloneAccountUpdaterClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}

func buildGetUpdatedCardCredentialsWithCard() GetUpdatedCardCredentialsRequest {
	return GetUpdatedCardCredentialsRequest{
		SourceOptions: AccountUpdaterSourceOptions{
			Card: &AccountUpdaterCard{
				Number:      "4242424242424242",
				ExpiryMonth: 12,
				ExpiryYear:  2025,
			},
		},
	}
}
