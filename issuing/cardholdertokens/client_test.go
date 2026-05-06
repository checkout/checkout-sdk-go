package cardholdertokens

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
)

func TestRequestCardholderToken(t *testing.T) {
	var (
		request = CardholderTokenRequest{
			GrantType:    "client_credentials",
			ClientId:     "ack_abc123",
			ClientSecret: "secret",
			CardholderId: "crh_d3oqhf46pyzuxjbcn2giaqnb44",
			SingleUse:    true,
		}

		response = CardholderTokenResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			AccessToken:  "eyJhbGciOiJIUzI1NiJ9.token",
			TokenType:    "Bearer",
			ExpiresIn:    3600,
			Scope:        "issuing:cardholder:read",
		}
	)

	cases := []struct {
		name     string
		request  CardholderTokenRequest
		apiPost  func(*mock.Mock) mock.Call
		checker  func(*CardholderTokenResponse, error)
		formData url.Values
	}{
		{
			name:    "when request is correct then return cardholder token",
			request: request,
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostFormWithContext", mock.Anything, common.BuildPath(cardholderTokenPath), (*configuration.SdkAuthorization)(nil), mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CardholderTokenResponse)
						*respMapping = response
						formData := args.Get(3).(url.Values)
						assert.Equal(t, "client_credentials", formData.Get("grant_type"))
						assert.Equal(t, "ack_abc123", formData.Get("client_id"))
						assert.Equal(t, "secret", formData.Get("client_secret"))
						assert.Equal(t, "crh_d3oqhf46pyzuxjbcn2giaqnb44", formData.Get("cardholder_id"))
						assert.Equal(t, "true", formData.Get("single_use"))
					})
			},
			checker: func(response *CardholderTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, "eyJhbGciOiJIUzI1NiJ9.token", response.AccessToken)
				assert.Equal(t, "Bearer", response.TokenType)
				assert.Equal(t, float64(3600), response.ExpiresIn)
			},
		},
		{
			name: "when single_use is false then form omits the field",
			request: CardholderTokenRequest{
				GrantType:    "client_credentials",
				ClientId:     "ack_abc123",
				ClientSecret: "secret",
				CardholderId: "crh_d3oqhf46pyzuxjbcn2giaqnb44",
				SingleUse:    false,
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostFormWithContext", mock.Anything, common.BuildPath(cardholderTokenPath), (*configuration.SdkAuthorization)(nil), mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						formData := args.Get(3).(url.Values)
						assert.Empty(t, formData.Get("single_use"))
					})
			},
			checker: func(response *CardholderTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
		{
			name:    "when api returns error then propagate it",
			request: request,
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostFormWithContext", mock.Anything, common.BuildPath(cardholderTokenPath), (*configuration.SdkAuthorization)(nil), mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusUnauthorized,
						Status:     "401 Unauthorized",
					})
			},
			checker: func(response *CardholderTokenResponse, err error) {
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

			tc.apiPost(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			c := NewClient(config, apiClient)

			tc.checker(c.RequestCardholderToken(tc.request))
		})
	}
}
