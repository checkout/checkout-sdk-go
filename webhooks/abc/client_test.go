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

func TestClientRetrieveWebhooks(t *testing.T) {
	var (
		response = WebhooksResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			WebhookArray: []WebhookResponse{
				{
					Id:          "wh_387ac7a83a054e37ae140105429d76b5",
					Url:         "https://example.com/webhooks",
					Active:      true,
					Headers:     nil,
					ContentType: Json,
					EventTypes: []string{
						"payment_approved",
						"payment_pending",
						"payment_declined",
						"payment_expired",
						"payment_canceled",
						"payment_voided",
						"payment_void_declined",
						"payment_captured",
						"payment_capture_declined",
						"payment_capture_pending",
						"payment_refunded",
						"payment_refund_declined",
						"payment_refund_pending",
					},
					Links: nil,
				},
			},
		}
	)
	cases := []struct {
		name             string
		getAuthorization func(mock *mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*WebhooksResponse, error)
	}{
		{
			name: "when there are webhooks then return all webhooks",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*WebhooksResponse)
						*respMapping = response
					})
			},
			checker: func(response *WebhooksResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
			},
		},
		{
			name: "when webhooks are not configured then return 204",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*WebhooksResponse)
						*respMapping = WebhooksResponse{
							HttpResponse: mocks.HttpMetadataStatusNoContent,
						}
					})
			},
			checker: func(response *WebhooksResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpResponse.StatusCode)
			},
		},
		{
			name: "when request invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnauthorized,
							Status:     "401 Unauthorized",
						})
			},
			checker: func(response *WebhooksResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

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
			eventsClient := NewClient(config, apiClient)

			tc.checker(eventsClient.RetrieveWebhooks())
		})
	}
}

func TestClientRegisterWebhook(t *testing.T) {
	var (
		request = WebhookRequest{
			Url:         "https://example.com/webhooks",
			Active:      true,
			Headers:     nil,
			ContentType: Json,
			EventTypes: []string{
				"payment_approved",
				"payment_pending",
				"payment_declined",
				"payment_expired",
				"payment_canceled",
				"payment_voided",
				"payment_void_declined",
				"payment_captured",
				"payment_capture_declined",
				"payment_capture_pending",
				"payment_refunded",
				"payment_refund_declined",
				"payment_refund_pending",
			},
		}

		response = WebhookResponse{
			HttpResponse: mocks.HttpMetadataStatusCreated,
			Id:           "wh_387ac7a83a054e37ae140105429d76b5",
			Url:          "https://example.com/webhooks",
			Active:       true,
			Headers:      nil,
			ContentType:  Json,
			EventTypes: []string{
				"payment_approved",
				"payment_pending",
				"payment_declined",
				"payment_expired",
				"payment_canceled",
				"payment_voided",
				"payment_void_declined",
				"payment_captured",
				"payment_capture_declined",
				"payment_capture_pending",
				"payment_refunded",
				"payment_refund_declined",
				"payment_refund_pending",
			},
			Links: nil,
		}
	)

	cases := []struct {
		name             string
		request          WebhookRequest
		getAuthorization func(mock *mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*WebhookResponse, error)
	}{
		{
			name:    "when request is correct then create instrument",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext",
					mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*WebhookResponse)
						*respMapping = response
					})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpResponse.StatusCode)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext",
					mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: WebhookRequest{},
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
							Status:     "422 Invalid Request",
							Data: &errors.ErrorDetails{
								ErrorType: "request_invalid",
								ErrorCodes: []string{
									"email_required",
								},
							},
						})
			},
			checker: func(response *WebhookResponse, err error) {
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
			eventsClient := NewClient(config, apiClient)

			tc.checker(eventsClient.RegisterWebhook(tc.request))
		})
	}
}

func TestClientRetrieveWebhook(t *testing.T) {
	var (
		webhookId = "wh_387ac7a83a054e37ae140105429d76b5"

		response = WebhookResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			Url:          "https://example.com/webhooks",
			Active:       true,
			Headers:      nil,
			ContentType:  Json,
			EventTypes: []string{
				"payment_approved",
				"payment_pending",
				"payment_declined",
				"payment_expired",
				"payment_canceled",
				"payment_voided",
				"payment_void_declined",
				"payment_captured",
				"payment_capture_declined",
				"payment_capture_pending",
				"payment_refunded",
				"payment_refund_declined",
				"payment_refund_pending",
			},
		}
	)

	cases := []struct {
		name             string
		webhookId        string
		getAuthorization func(mock *mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*WebhookResponse, error)
	}{
		{
			name:      "when webhook exist then return it",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*WebhookResponse)
						*respMapping = response
					})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
			},
		},
		{
			name:      "when request invalid then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnauthorized,
							Status:     "401 Unauthorized",
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

			},
		},
		{
			name:      "when webhook does not exist then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *WebhookResponse, err error) {
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
			eventsClient := NewClient(config, apiClient)

			tc.checker(eventsClient.RetrieveWebhook(tc.webhookId))
		})
	}
}

func TestClientUpdateWebhook(t *testing.T) {
	var (
		webhookId = "wh_387ac7a83a054e37ae140105429d76b5"

		webhookRequest = WebhookRequest{
			Url:         "https://example.com/webhooks",
			Active:      true,
			Headers:     nil,
			ContentType: Json,
			EventTypes: []string{
				"payment_approved",
				"payment_pending",
				"payment_declined",
				"payment_expired",
				"payment_canceled",
				"payment_voided",
				"payment_void_declined",
				"payment_captured",
				"payment_capture_declined",
				"payment_capture_pending",
				"payment_refunded",
				"payment_refund_declined",
				"payment_refund_pending",
			},
		}

		response = WebhookResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			Url:          "https://example.com/webhooks",
			Active:       true,
			Headers:      nil,
			ContentType:  Json,
			EventTypes: []string{
				"payment_approved",
				"payment_pending",
				"payment_declined",
				"payment_expired",
				"payment_canceled",
				"payment_voided",
				"payment_void_declined",
				"payment_captured",
				"payment_capture_declined",
				"payment_capture_pending",
				"payment_refunded",
				"payment_refund_declined",
				"payment_refund_pending",
			},
		}
	)

	cases := []struct {
		name             string
		webhookId        string
		webhookRequest   WebhookRequest
		getAuthorization func(mock *mock.Mock) mock.Call
		apiUpdate        func(*mock.Mock) mock.Call
		checker          func(*WebhookResponse, error)
	}{
		{
			name:           "when webhook exist and update then return it modified",
			webhookId:      webhookId,
			webhookRequest: webhookRequest,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiUpdate: func(m *mock.Mock) mock.Call {
				return *m.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*WebhookResponse)
						*respMapping = response
					})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
			},
		},
		{
			name:      "when request invalid then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiUpdate: func(m *mock.Mock) mock.Call {
				return *m.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnauthorized,
							Status:     "401 Unauthorized",
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

			},
		},
		{
			name:      "when webhook does not exist then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiUpdate: func(m *mock.Mock) mock.Call {
				return *m.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:      "when webhook url exist then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiUpdate: func(m *mock.Mock) mock.Call {
				return *m.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusConflict,
							Status:     "409 Conflict",
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusConflict, chkErr.StatusCode)
			},
		},
		{
			name:      "when request is not correct then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiUpdate: func(m *mock.Mock) mock.Call {
				return *m.On("PutWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "invalid_request",
								ErrorCodes: []string{"payment_source_required"},
							},
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "payment_source_required")
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
			tc.apiUpdate(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			eventsClient := NewClient(config, apiClient)

			tc.checker(eventsClient.UpdateWebhook(tc.webhookId, tc.webhookRequest))
		})
	}
}

func TestClientPartiallyUpdateWebhook(t *testing.T) {
	var (
		webhookId = "wh_387ac7a83a054e37ae140105429d76b5"

		webhookRequest = WebhookRequest{
			Url:         "https://example.com/webhooks",
			Active:      true,
			Headers:     nil,
			ContentType: Json,
			EventTypes: []string{
				"payment_approved",
				"payment_pending",
				"payment_declined",
				"payment_expired",
				"payment_canceled",
				"payment_voided",
				"payment_void_declined",
				"payment_captured",
				"payment_capture_declined",
				"payment_capture_pending",
				"payment_refunded",
				"payment_refund_declined",
				"payment_refund_pending",
			},
		}

		response = WebhookResponse{
			HttpResponse: mocks.HttpMetadataStatusOk,
			Url:          "https://example.com/webhooks",
			Active:       true,
			Headers:      nil,
			ContentType:  Json,
			EventTypes: []string{
				"payment_approved",
				"payment_pending",
				"payment_declined",
				"payment_expired",
				"payment_canceled",
				"payment_voided",
				"payment_void_declined",
				"payment_captured",
				"payment_capture_declined",
				"payment_capture_pending",
				"payment_refunded",
				"payment_refund_declined",
				"payment_refund_pending",
			},
		}
	)

	cases := []struct {
		name             string
		webhookId        string
		webhookRequest   WebhookRequest
		getAuthorization func(mock *mock.Mock) mock.Call
		apiPatch         func(*mock.Mock) mock.Call
		checker          func(*WebhookResponse, error)
	}{
		{
			name:           "when webhook exist and patch then return it modified",
			webhookId:      webhookId,
			webhookRequest: webhookRequest,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*WebhookResponse)
						*respMapping = response
					})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpResponse.StatusCode)
			},
		},
		{
			name:      "when request invalid then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnauthorized,
							Status:     "401 Unauthorized",
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

			},
		},
		{
			name:      "when webhook does not exist then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:      "when webhook url exist then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusConflict,
							Status:     "409 Conflict",
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusConflict, chkErr.StatusCode)
			},
		},
		{
			name:      "when request is not correct then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "invalid_request",
								ErrorCodes: []string{"payment_source_required"},
							},
						})
			},
			checker: func(response *WebhookResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "payment_source_required")
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
			eventsClient := NewClient(config, apiClient)

			tc.checker(eventsClient.PartiallyUpdateWebhook(tc.webhookId, tc.webhookRequest))
		})
	}
}

func TestClientRemoveWebhook(t *testing.T) {
	var (
		webhookId = "wh_387ac7a83a054e37ae140105429d76b5"

		response = common.MetadataResponse{HttpMetadata: mocks.HttpMetadataStatusOk}
	)

	cases := []struct {
		name             string
		webhookId        string
		getAuthorization func(mock *mock.Mock) mock.Call
		apiDelete        func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:      "when delete a webhook then return 200 Ok",
			webhookId: webhookId,
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
			name:      "when request invalid then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnauthorized,
							Status:     "401 Unauthorized",
						})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

			},
		},
		{
			name:      "when webhook does not exist then return error",
			webhookId: webhookId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiDelete(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			eventsClient := NewClient(config, apiClient)

			tc.checker(eventsClient.RemoveWebhook(tc.webhookId))
		})
	}
}
