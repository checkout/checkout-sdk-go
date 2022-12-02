package links

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

func TestCreatePaymentLink(t *testing.T) {
	var (
		payementLinkResponse = PaymentLinkResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "pay_1234",
			Reference:    "reference",
		}
	)

	cases := []struct {
		name             string
		request          PaymentLinkRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*PaymentLinkResponse, error)
	}{
		{
			name: "when request is correct then create payment link",
			request: PaymentLinkRequest{
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
						respMapping := args.Get(3).(*PaymentLinkResponse)
						*respMapping = payementLinkResponse
					})
			},
			checker: func(response *PaymentLinkResponse, err error) {
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
			checker: func(response *PaymentLinkResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: PaymentLinkRequest{},
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
			checker: func(response *PaymentLinkResponse, err error) {
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			client := NewClient(configuration, apiClient)

			tc.checker(client.CreatePaymentLink(tc.request))
		})
	}
}

func TestGetPaymentLink(t *testing.T) {
	var (
		paymentLinkDetails = PaymentLinkDetails{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "pay_1234",
			Status:       Active,
			PaymentId:    "12345",
			Amount:       1000,
			Currency:     common.GBP,
			Reference:    "reference",
		}
	)

	cases := []struct {
		name             string
		paymentLinkId    string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*PaymentLinkDetails, error)
	}{
		{
			name:          "when hostedPaymentId is correct then return payment link details",
			paymentLinkId: "pay_1234",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*PaymentLinkDetails)
						*respMapping = paymentLinkDetails
					})
			},
			checker: func(response *PaymentLinkDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, paymentLinkDetails.Id, response.Id)
				assert.Equal(t, paymentLinkDetails.Status, response.Status)
				assert.Equal(t, paymentLinkDetails.PaymentId, response.PaymentId)
				assert.Equal(t, paymentLinkDetails.Amount, response.Amount)
				assert.Equal(t, paymentLinkDetails.Currency, response.Currency)
				assert.Equal(t, paymentLinkDetails.Reference, response.Reference)
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
			checker: func(response *PaymentLinkDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:          "when paymentLinkId not found then return error",
			paymentLinkId: "not_found",
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
			checker: func(response *PaymentLinkDetails, err error) {
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetPaymentLink(tc.paymentLinkId))
		})
	}
}
