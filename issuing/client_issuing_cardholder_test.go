package issuing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"

	cardholders "github.com/checkout/checkout-sdk-go/v2/issuing/cardholders"
)

// # tests

func TestUpdateCardholder(t *testing.T) {
	var (
		response = cardholders.CardholderUpdateResponse{
			HttpMetadata:     mocks.HttpMetadataStatusOk,
			LastModifiedDate: &lastModifiedDate,
			Links:            links,
		}
	)

	cases := []struct {
		name             string
		cardholderId     string
		request          cardholders.CardholderRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPatch         func(*mock.Mock) mock.Call
		checker          func(*cardholders.CardholderUpdateResponse, error)
	}{
		{
			name:         "when request is correct then should return 200",
			cardholderId: "crh_d3ozhf43pcq2xbldn2g45qnb44",
			request:      buildCardholderUpdateRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*cardholders.CardholderUpdateResponse)
						*respMapping = response
					})
			},
			checker: assertUpdateCardholderSuccess(t, &response),
		},
		{
			name:         "when credentials are invalid then return authorization error",
			cardholderId: "crh_d3ozhf43pcq2xbldn2g45qnb44",
			request:      cardholders.CardholderRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *cardholders.CardholderUpdateResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:         "when cardholder not found then return 404",
			cardholderId: "crh_not_found",
			request:      cardholders.CardholderRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *cardholders.CardholderUpdateResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:         "when request is invalid then return 422",
			cardholderId: "crh_d3ozhf43pcq2xbldn2g45qnb44",
			request:      cardholders.CardholderRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable",
						Data: &errors.ErrorDetails{
							ErrorType:  "request_invalid",
							ErrorCodes: []string{"request_body_malformed"},
						},
					})
			},
			checker: func(response *cardholders.CardholderUpdateResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "request_body_malformed")
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
			client := NewClient(config, apiClient)

			tc.checker(client.UpdateCardholder(tc.cardholderId, tc.request))
		})
	}
}

// # common methods

func buildCardholderUpdateRequest() cardholders.CardholderRequest {
	return cardholders.CardholderRequest{
		FirstName:      "John",
		LastName:       "Kennedy",
		Email:          "john.kennedy@myemaildomain.com",
		PhoneNumber:    phone,
		BillingAddress: address,
	}
}

func assertUpdateCardholderSuccess(t *testing.T, expected *cardholders.CardholderUpdateResponse) func(*cardholders.CardholderUpdateResponse, error) {
	return func(response *cardholders.CardholderUpdateResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
		assert.NotNil(t, response.LastModifiedDate)
		assert.NotNil(t, response.Links)
	}
}
