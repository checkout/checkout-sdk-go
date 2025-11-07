package payment_sessions

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
	"github.com/checkout/checkout-sdk-go/payments"
)

func TestCreateAPaymentSessions(t *testing.T) {
	var (
		paymentSessionsResponse = PaymentSessionsResponse{
			HttpMetadata:         mocks.HttpMetadataStatusCreated,
			Id:                   "ps_2Un6I6lRpIAiIEwQIyxWVnV9CqQ",
			PaymentSessionToken:  "string",
			PaymentSessionSecret: "pss_9823241e-2cec-4c98-b23d-7b29ow4e2e34",
			Links: map[string]common.Link{
				"self": {
					HRef: &[]string{"https://api.sandbox.checkout.com/payment-sessions/ps_2Un6I6lRpIAiIEwQIyxWVnV9CqQ"}[0],
				},
			},
		}
	)

	cases := []struct {
		name             string
		request          PaymentSessionsRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*PaymentSessionsResponse, error)
	}{
		{
			name: "when request is correct then create a payment sessions",
			request: PaymentSessionsRequest{
				Amount:    2000,
				Currency:  common.GBP,
				Reference: "ORD-123A",
				Billing: &payments.BillingInformation{Address: &common.Address{
					Country: common.GB,
				}},
				Customer: &common.CustomerRequest{
					Email: "bruce@wayne-enterprises.com",
					Name:  "Bruce Wayne",
				},
				SuccessUrl: "https://example.com/payments/success",
				FailureUrl: "https://example.com/payments/failure",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentSessionsResponse)
						*respMapping = paymentSessionsResponse
					})
			},
			checker: func(response *PaymentSessionsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentSessionsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: PaymentSessionsRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Invalid Request",
							Data:       &errors.ErrorDetails{ErrorType: "request_invalid"},
						})
			},
			checker: func(response *PaymentSessionsResponse, err error) {
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

			tc.checker(client.RequestPaymentSessions(tc.request))
		})
	}
}

func TestRequestPaymentSessionsWithPayment(t *testing.T) {
	var (
		paymentSessionPaymentResponse = PaymentSessionPaymentResponse{
			HttpMetadata:         mocks.HttpMetadataStatusCreated,
			Id:                   "pay_mbabizu24mvu3mela5njyhpit4",
			Status:               "Approved",
			Type:                 CardPMT,
			PaymentSessionId:     "ps_2Un6I6lRpIAiIEwQIyxWVnV9CqQ",
			PaymentSessionSecret: "pss_9823241e-2cec-4c98-b23d-7b29ow4e2e34",
		}
	)

	cases := []struct {
		name             string
		request          PaymentSessionsWithPaymentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*PaymentSessionPaymentResponse, error)
	}{
		{
			name: "when request is correct then create a payment sessions with payment",
			request: PaymentSessionsWithPaymentRequest{
				SessionData: "session_data_token",
				Amount:      2000,
				Currency:    common.GBP,
				Reference:   "ORD-123A",
				Billing: &payments.BillingInformation{Address: &common.Address{
					Country: common.GB,
				}},
				Customer: &common.CustomerRequest{
					Email: "bruce@wayne-enterprises.com",
					Name:  "Bruce Wayne",
				},
				SuccessUrl: "https://example.com/payments/success",
				FailureUrl: "https://example.com/payments/failure",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentSessionPaymentResponse)
						*respMapping = paymentSessionPaymentResponse
					})
			},
			checker: func(response *PaymentSessionPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
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

			tc.checker(client.RequestPaymentSessionsWithPayment(tc.request))
		})
	}
}

func TestSubmitPaymentSession(t *testing.T) {
	var (
		paymentSessionPaymentResponse = PaymentSessionPaymentResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "pay_mbabizu24mvu3mela5njyhpit4",
			Status:       "Approved",
			Type:         CardPMT,
		}
	)

	cases := []struct {
		name             string
		sessionId        string
		request          SubmitPaymentSessionRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*PaymentSessionPaymentResponse, error)
	}{
		{
			name:      "when request is correct then submit payment session",
			sessionId: "ps_2Un6I6lRpIAiIEwQIyxWVnV9CqQ",
			request: SubmitPaymentSessionRequest{
				SessionData: "session_data_token",
				Amount:      2000,
				Reference:   "ORD-123A",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentSessionPaymentResponse)
						*respMapping = paymentSessionPaymentResponse
					})
			},
			checker: func(response *PaymentSessionPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
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

			tc.checker(client.SubmitPaymentSession(tc.sessionId, tc.request))
		})
	}
}
