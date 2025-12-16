package setups

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

func TestCreatePaymentSetup(t *testing.T) {
	var (
		setupResponse = PaymentSetupResponse{
			Id:       "ps_123456789",
			Amount:   1000,
			Currency: common.GBP,
		}
	)

	cases := []struct {
		name             string
		request          PaymentSetupRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*PaymentSetupResponse, error)
	}{
		{
			name: "when request is correct then create payment setup",
			request: PaymentSetupRequest{
				Amount:   1000,
				Currency: common.GBP,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentSetupResponse)
						*respMapping = setupResponse
					})
			},
			checker: func(response *PaymentSetupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, setupResponse.Id, response.Id)
				assert.Equal(t, setupResponse.Amount, response.Amount)
				assert.Equal(t, setupResponse.Currency, response.Currency)
			},
		},
		{
			name: "when credentials invalid then return error",
			request: PaymentSetupRequest{
				Amount:   1000,
				Currency: common.GBP,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentSetupResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
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

			enableTelemetry := true
			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)
			tc.checker(client.CreatePaymentSetup(tc.request))
		})
	}
}

func TestUpdatePaymentSetup(t *testing.T) {
	var (
		setupId       = "ps_123456789"
		setupResponse = PaymentSetupResponse{
			Id:       setupId,
			Amount:   1000,
			Currency: common.GBP,
		}
	)

	cases := []struct {
		name             string
		setupId          string
		request          PaymentSetupRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPut           func(*mock.Mock) mock.Call
		checker          func(*PaymentSetupResponse, error)
	}{
		{
			name:    "when request is correct then update payment setup",
			setupId: setupId,
			request: PaymentSetupRequest{
				Amount:   1000,
				Currency: common.GBP,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPut: func(m *mock.Mock) mock.Call {
				return *m.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentSetupResponse)
						*respMapping = setupResponse
					})
			},
			checker: func(response *PaymentSetupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, setupResponse.Id, response.Id)
				assert.Equal(t, setupResponse.Amount, response.Amount)
				assert.Equal(t, setupResponse.Currency, response.Currency)
			},
		},
		{
			name:    "when credentials invalid then return error",
			setupId: setupId,
			request: PaymentSetupRequest{
				Amount:   1000,
				Currency: common.GBP,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPut: func(m *mock.Mock) mock.Call {
				return *m.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentSetupResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiPut(&apiClient.Mock)

			enableTelemetry := true
			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)
			tc.checker(client.UpdatePaymentSetup(tc.setupId, tc.request))
		})
	}
}

func TestGetPaymentSetup(t *testing.T) {
	var (
		setupId       = "ps_123456789"
		setupResponse = PaymentSetupResponse{
			Id:       setupId,
			Amount:   1000,
			Currency: common.GBP,
		}
	)

	cases := []struct {
		name             string
		setupId          string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*PaymentSetupResponse, error)
	}{
		{
			name:    "when setupId is correct then get payment setup",
			setupId: setupId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*PaymentSetupResponse)
						*respMapping = setupResponse
					})
			},
			checker: func(response *PaymentSetupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, setupResponse.Id, response.Id)
				assert.Equal(t, setupResponse.Amount, response.Amount)
				assert.Equal(t, setupResponse.Currency, response.Currency)
			},
		},
		{
			name:    "when credentials invalid then return error",
			setupId: setupId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentSetupResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
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

			enableTelemetry := true
			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)
			tc.checker(client.GetPaymentSetup(tc.setupId))
		})
	}
}

func TestConfirmPaymentSetup(t *testing.T) {
	var (
		setupId               = "ps_123456789"
		paymentMethodOptionId = "pmo_123456789"
		confirmResponse       = PaymentSetupConfirmResponse{
			Id:     setupId,
			Status: payments.Authorized,
		}
	)

	cases := []struct {
		name                  string
		setupId               string
		paymentMethodOptionId string
		getAuthorization      func(*mock.Mock) mock.Call
		apiPost               func(*mock.Mock) mock.Call
		checker               func(*PaymentSetupConfirmResponse, error)
	}{
		{
			name:                  "when path parameters are correct then confirm payment setup",
			setupId:               setupId,
			paymentMethodOptionId: paymentMethodOptionId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentSetupConfirmResponse)
						*respMapping = confirmResponse
					})
			},
			checker: func(response *PaymentSetupConfirmResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, confirmResponse.Id, response.Id)
				assert.Equal(t, confirmResponse.Status, response.Status)
			},
		},
		{
			name:                  "when credentials invalid then return error",
			setupId:               setupId,
			paymentMethodOptionId: paymentMethodOptionId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentSetupConfirmResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
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

			enableTelemetry := true
			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)
			tc.checker(client.ConfirmPaymentSetup(tc.setupId, tc.paymentMethodOptionId))
		})
	}
}
