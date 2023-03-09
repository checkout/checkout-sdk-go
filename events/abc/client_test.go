package abc

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestRetrieveAllEventTypes(t *testing.T) {
	var (
		eventTypesAll = []EventTypes{
			{
				Version:    "1.0",
				EventTypes: []string{"event.1", "event.2", "event.3"},
			},
			{
				Version:    "2.0",
				EventTypes: []string{"event.4", "event.5"},
			},
		}

		eventTypes1Dot0 = []EventTypes{
			{
				Version:    "1.0",
				EventTypes: []string{"event.1", "event.2", "event.3"},
			},
		}

		response = EventTypesResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			EventTypes:   eventTypesAll,
		}
	)

	cases := []struct {
		name             string
		query            QueryRetrieveAllEventType
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*EventTypesResponse, error)
	}{
		{
			name: "when version is send then return events of this version",
			query: QueryRetrieveAllEventType{
				Version: "1.0",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*EventTypesResponse)
						*respMapping = EventTypesResponse{
							HttpResponse: mocks.HttpMetadataStatusOk,
							EventTypes:   eventTypes1Dot0,
						}
					})
			},
			checker: func(response *EventTypesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
				assert.Equal(t, eventTypes1Dot0, response.EventTypes)

			},
		},
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
				assert.Equal(t, eventTypesAll, response.EventTypes)

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
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			eventsClient := NewClient(configuration, apiClient)

			tc.checker(eventsClient.RetrieveAllEventTypesQuery(tc.query))
		})
	}
}

func TestRetrieveEvents(t *testing.T) {
	var (
		response = EventsPageResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			TotalCount:   1,
			Limit:        5,
			Skip:         0,
			Data: []EventsSummaryResponse{
				{
					Id:        "evt_3nup2pts3emebenhtw6ky4frim",
					Type:      "payment_approved",
					CreatedOn: "2021-06-25T09:40:12Z",
					Links:     nil,
				},
			},
		}
	)

	cases := []struct {
		name             string
		query            QueryRetrieveEvents
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*EventsPageResponse, error)
	}{
		{
			name: "when events exists then return events",
			query: QueryRetrieveEvents{
				PaymentId: "pay_ok2ynq6ubn3ufmo6jsdfmdvy5q",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*EventsPageResponse)
						*respMapping = response
					})
			},
			checker: func(response *EventsPageResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
				assert.NotNil(t, response.TotalCount)
				assert.NotNil(t, response.Limit)
				assert.NotNil(t, response.Skip)
				assert.NotNil(t, response.Data)
			},
		},
		{
			name: "when event exists but do not have data then return 204 No Content",
			query: QueryRetrieveEvents{
				PaymentId: "pay_ok2ynq6ubn3ufmo6jsdfmdvy5q",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*EventsPageResponse)
						*respMapping = EventsPageResponse{
							HttpResponse: mocks.HttpMetadataStatusNoContent,
						}
					})
			},
			checker: func(response *EventsPageResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpResponse.StatusCode)
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
			checker: func(response *EventsPageResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())

			},
		},
		{
			name:  "when request is not correct then return error",
			query: QueryRetrieveEvents{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "request_invalid",
								ErrorCodes: []string{"payment_id_invalid"},
							},
						})
			},
			checker: func(response *EventsPageResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "payment_id_invalid")
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			eventsClient := NewClient(configuration, apiClient)

			tc.checker(eventsClient.RetrieveEventsQuery(tc.query))
		})
	}
}

func TestRetrieveEvent(t *testing.T) {
	var (
		eventId = "evt_az5sblvku4ge3dwpztvyizgcau"

		response = EventResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			Id:           "evt_az5sblvku4ge3dwpztvyizgcau",
			Type:         "payment_approved",
			Version:      "2.0",
			CreatedOn:    "2019-08-24T14:15:22Z",
		}
	)

	cases := []struct {
		name             string
		eventId          string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*EventResponse, error)
	}{
		{
			name:    "when event exists then return event",
			eventId: eventId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*EventResponse)
						*respMapping = response
					})
			},
			checker: func(response *EventResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
				assert.Equal(t, eventId, response.Id)
				assert.NotNil(t, response.Type)
				assert.NotNil(t, response.Version)
				assert.NotNil(t, response.CreatedOn)
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
			checker: func(response *EventResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())

			},
		},
		{
			name:    "when event does not exist then return error",
			eventId: "evt_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *EventResponse, err error) {
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			eventsClient := NewClient(configuration, apiClient)

			tc.checker(eventsClient.RetrieveEvent(tc.eventId))
		})
	}
}

func TestRetrieveEventNotification(t *testing.T) {
	var (
		eventId = "evt_az5sblvku4ge3dwpztvyizgcau"

		response = EventNotificationResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			Id:           "ntf_az5sblvku4ge3dwpztvyizgcau",
			Url:          "https://example.com/webhooks",
			Success:      false,
			ContentType:  "json",
		}
	)

	cases := []struct {
		name             string
		eventId          string
		notificationId   string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*EventNotificationResponse, error)
	}{
		{
			name:    "when event exists then return event",
			eventId: eventId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*EventNotificationResponse)
						*respMapping = response
					})
			},
			checker: func(response *EventNotificationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Url)
				assert.NotNil(t, response.Success)
				assert.NotNil(t, response.ContentType)
			},
		},
		{
			name: "when credentials invalid then return errorr",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *EventNotificationResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())

			},
		},
		{
			name:           "when event does not exist then return error",
			eventId:        "evt_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			notificationId: "not_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *EventNotificationResponse, err error) {
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			eventsClient := NewClient(configuration, apiClient)

			tc.checker(eventsClient.RetrieveEventNotification(tc.eventId, tc.notificationId))
		})
	}
}

func TestRetryWebhook(t *testing.T) {
	eventId := "evt_az5sblvku4ge3dwpztvyizgcau"
	webhookId := "wh_az5sblvku4ge3dwpztvyizgcau"

	var (
		metadataResponse = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusAccepted,
		}
	)

	cases := []struct {
		name             string
		eventId          string
		webhookId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:      "when retry a webhook then return 202 status",
			eventId:   eventId,
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.MetadataResponse)
						*respMapping = metadataResponse
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())

			},
		},
		{
			name:    "when webhook does not exist then return error",
			eventId: "evt_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
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
		{
			name:      "when webhook does not exist then return error",
			webhookId: "evt_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			eventsClient := NewClient(configuration, apiClient)

			tc.checker(eventsClient.RetryWebhook(tc.eventId, tc.webhookId))
		})
	}
}

func TestRetryAllWebhooks(t *testing.T) {
	eventId := "evt_az5sblvku4ge3dwpztvyizgcau"

	var (
		metadataResponse = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusAccepted,
		}
	)

	cases := []struct {
		name             string
		eventId          string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:    "when retry all webhooks then return 202 status",
			eventId: eventId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.MetadataResponse)
						*respMapping = metadataResponse
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())

			},
		},
		{
			name:    "when webhooks by event does not exist then return error",
			eventId: "evt_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			eventsClient := NewClient(configuration, apiClient)

			tc.checker(eventsClient.RetryAllWebhooks(tc.eventId))
		})
	}
}
