package hosted

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

func TestCreateHostedPaymentsPageSession(t *testing.T) {
	var (
		hostedPaymentResponse = HostedPaymentResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "pay_1234",
			Reference:    "reference",
		}
	)

	cases := []struct {
		name             string
		request          HostedPaymentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*HostedPaymentResponse, error)
	}{
		{
			name: "when request is correct then create hosted payment session",
			request: HostedPaymentRequest{
				Amount:      1000,
				Currency:    common.GBP,
				PaymentType: payments.Regular,
				Reference:   "reference",
				Description: "description",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*HostedPaymentResponse)
						*respMapping = hostedPaymentResponse
					})
			},
			checker: func(response *HostedPaymentResponse, err error) {
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
			checker: func(response *HostedPaymentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: HostedPaymentRequest{},
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
			checker: func(response *HostedPaymentResponse, err error) {
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.CreateHostedPaymentsPageSession(tc.request))
		})
	}
}

func TestGetHostedPaymentsPageDetails(t *testing.T) {
	var (
		hostedPaymentDetails = HostedPaymentDetails{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "pay_1234",
			Status:       PaymentPending,
			PaymentId:    "12345",
			Amount:       1000,
			Currency:     common.GBP,
			Reference:    "reference",
		}
	)

	cases := []struct {
		name             string
		hostedPaymentId  string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*HostedPaymentDetails, error)
	}{
		{
			name:            "when hostedPaymentId is correct then return hosted payment session details",
			hostedPaymentId: "pay_1234",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*HostedPaymentDetails)
						*respMapping = hostedPaymentDetails
					})
			},
			checker: func(response *HostedPaymentDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, hostedPaymentDetails.Id, response.Id)
				assert.Equal(t, hostedPaymentDetails.Status, response.Status)
				assert.Equal(t, hostedPaymentDetails.PaymentId, response.PaymentId)
				assert.Equal(t, hostedPaymentDetails.Amount, response.Amount)
				assert.Equal(t, hostedPaymentDetails.Currency, response.Currency)
				assert.Equal(t, hostedPaymentDetails.Reference, response.Reference)
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
			checker: func(response *HostedPaymentDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:            "when hostedPaymentId not found then return error",
			hostedPaymentId: "not_found",
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
			checker: func(response *HostedPaymentDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
				assert.Equal(t, "404 Not Found", chkErr.Status)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetHostedPaymentsPageDetails(tc.hostedPaymentId))
		})
	}
}
