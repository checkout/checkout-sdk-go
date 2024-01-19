package contexts

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
	"github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/nas/sources/contexts"
)

func TestCreateAPaymentContext(t *testing.T) {
	var (
		paymentContextsRequestResponse = PaymentContextsRequestResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "pct_y3oqhf46pyzuxjbcn2giaqnb44",
			PartnerMetadata: &PaymentContextsPartnerMetadata{
				OrderId:    "test_order_123",
				CustomerId: "cus_123",
			},
			Links: map[string]common.Link{
				"self": {
					HRef: &[]string{"https://api.checkout.com/payment-contexts/pct_y3oqhf46pyzuxjbcn2giaqnb44"}[0],
				},
			},
		}
	)

	cases := []struct {
		name             string
		request          PaymentContextsRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*PaymentContextsRequestResponse, error)
	}{
		{
			name: "when request is correct then create a payment context",
			request: PaymentContextsRequest{
				Source:      contexts.NewPaymentContextsPaypalSource(),
				Amount:      2000,
				Currency:    common.EUR,
				PaymentType: payments.Regular,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentContextsRequestResponse)
						*respMapping = paymentContextsRequestResponse
					})
			},
			checker: func(response *PaymentContextsRequestResponse, err error) {
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
			checker: func(response *PaymentContextsRequestResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: PaymentContextsRequest{},
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
			checker: func(response *PaymentContextsRequestResponse, err error) {
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.RequestPaymentContexts(tc.request))
		})
	}
}

func TestGetAPaymentContext(t *testing.T) {
	var (
		paymentContextResponse = &PaymentContextsResponse{
			Source: &nas.SourceResponse{
				PaymentContextsPayPayResponseSource: &nas.PaymentContextsPayPayResponseSource{
					Type: payments.PayPalSource,
				},
			},
			Amount:      2000,
			Currency:    common.EUR,
			PaymentType: payments.Regular,
			Capture:     true,
		}

		paymentContextDetailsResponse = PaymentContextDetailsResponse{
			HttpMetadata:   mocks.HttpMetadataStatusOk,
			PaymentRequest: paymentContextResponse,
		}
	)

	cases := []struct {
		name             string
		paymentContextId string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*PaymentContextDetailsResponse, error)
	}{
		{
			name:             "when paymentContextId is correct then return payment context details",
			paymentContextId: "pay_1234",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*PaymentContextDetailsResponse)
						*respMapping = paymentContextDetailsResponse
					})
			},
			checker: func(response *PaymentContextDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, paymentContextDetailsResponse.PaymentRequest.Amount, response.PaymentRequest.Amount)
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
			checker: func(response *PaymentContextDetailsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:             "when paymentContextId not found then return error",
			paymentContextId: "not_found",
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
			checker: func(response *PaymentContextDetailsResponse, err error) {
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetPaymentContextDetails(tc.paymentContextId))
		})
	}
}
