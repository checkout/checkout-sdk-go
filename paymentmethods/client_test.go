package paymentmethods

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

func TestGetAvailablePaymentMethods(t *testing.T) {
	var (
		methodsResponse = GetPaymentMethodsResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Methods: []PaymentMethod{
				{Type: Visa, PartnerMerchantId: "merchant_123"},
				{Type: Klarna, PartnerMerchantId: "merchant_456"},
			},
		}
	)

	cases := []struct {
		name             string
		query            GetPaymentMethodsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*GetPaymentMethodsResponse, error)
	}{
		{
			name:  "when processing channel id is valid then return payment methods",
			query: buildGetPaymentMethodsQuery(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*GetPaymentMethodsResponse)
						*respMapping = methodsResponse
					})
			},
			checker: func(response *GetPaymentMethodsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, len(methodsResponse.Methods), len(response.Methods))
				assert.Equal(t, methodsResponse.Methods[0].Type, response.Methods[0].Type)
				assert.Equal(t, methodsResponse.Methods[0].PartnerMerchantId, response.Methods[0].PartnerMerchantId)
			},
		},
		{
			name:  "when credentials invalid then return error",
			query: buildGetPaymentMethodsQuery(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *GetPaymentMethodsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:  "when processing channel not found then return error",
			query: GetPaymentMethodsQuery{ProcessingChannelId: "pc_invalid"},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
			},
			checker: func(response *GetPaymentMethodsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildPaymentMethodsClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.GetAvailablePaymentMethods(tc.query))
		})
	}
}

// common methods

func buildPaymentMethodsClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}

func buildGetPaymentMethodsQuery() GetPaymentMethodsQuery {
	return GetPaymentMethodsQuery{ProcessingChannelId: "pc_5jp2az55l3cuths25t5p3xhwru"}
}
