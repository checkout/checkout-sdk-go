package events

import (
	"github.com/checkout/checkout-sdk-go/errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestRetrieveAllEventTypes(t *testing.T) {
	var (
		eventTypes = []EventTypes{
			{
				Version:    "1.0",
				EventTypes: []string{"event.1", "event.2", "event.3"},
			},
			{
				Version:    "2.0",
				EventTypes: []string{"event.4", "event.5"},
			},
		}

		httpMetadata = common.HttpMetadata{
			Status:     "200 OK",
			StatusCode: 200,
		}

		response = EventTypesResponse{
			HttpResponse: httpMetadata,
			EventTypes:   eventTypes,
		}
	)

	cases := []struct {
		name             string
		requestVersions  string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*EventTypesResponse, error)
	}{
		{
			name: "when no event versions sent then return all events",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*EventTypesResponse)
						*respMapping = response
					})
			},
			checker: func(response *EventTypesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
				assert.Equal(t, httpMetadata, response.HttpResponse)
				assert.Equal(t, eventTypes, response.EventTypes)

			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *EventTypesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())

			},
		},
		{
			name: "when request invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnauthorized,
							Status:     "401 Unauthorized",
						})
			},
			checker: func(response *EventTypesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

			},
		},
		// TODO complete with more cases
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			eventsClient := NewClient(configuration, apiClient)

			tc.checker(eventsClient.RetrieveAllEventTypes(tc.requestVersions))
		})
	}
}
