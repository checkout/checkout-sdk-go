package bacs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
	"github.com/checkout/checkout-sdk-go/v3/mocks"
)

func TestSendNotification(t *testing.T) {
	var (
		notification = NotificationResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			EventId:      "evt_lzr4csdtddwetactr6phd3kea4",
		}
	)

	cases := []struct {
		name             string
		request          NotificationRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*NotificationResponse, error)
	}{
		{
			name:    "when request is valid then send the pre-notification",
			request: NotificationRequest{SourceId: "src_wmlfc3zyhqzehihu7giusaaawu"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext",
					mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						// Positional args are (ctx, path, auth, request, response, idempotencyKey).
						assert.Equal(t, "/apms/bacs/notifications", args.Get(1))
						respMapping := args.Get(4).(*NotificationResponse)
						*respMapping = notification
					})
			},
			checker: func(response *NotificationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, notification.EventId, response.EventId)
			},
		},
		{
			name:    "when request is invalid then return error",
			request: NotificationRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext",
					mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Invalid data was sent",
						})
			},
			checker: func(response *NotificationResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
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

			tc.checker(client.SendNotification(tc.request))
		})
	}
}
