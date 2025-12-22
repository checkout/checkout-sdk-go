package tokens

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
)

var (
	typeCard    = Card
	typeApple   = ApplePay
	cardNumber  = "123456789"
	expiryMonth = 01
	expiryYear  = 30
	name        = "Bruce Wayne"
	tokenId     = "tok_1234"
)

func TestRequestCardToken(t *testing.T) {
	var (
		cardTokenRequest = CardTokenRequest{
			Type:        Card,
			Number:      "123456789",
			ExpiryMonth: 01,
			ExpiryYear:  30,
			Name:        "Bruce Wayne",
			CVV:         "111",
		}

		cardTokenResponse = CardTokenResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Type:         typeCard,
			Token:        tokenId,
			ExpiryMonth:  expiryMonth,
			ExpiryYear:   expiryYear,
			Name:         name,
		}
	)

	cases := []struct {
		name             string
		request          CardTokenRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CardTokenResponse, error)
	}{
		{
			name:    "when request is correct then return card token",
			request: cardTokenRequest,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CardTokenResponse)
						*respMapping = cardTokenResponse
					})
			},
			checker: func(response *CardTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, cardTokenResponse.Type, response.Type)
				assert.Equal(t, cardTokenResponse.Token, response.Token)
				assert.Equal(t, cardTokenResponse.ExpiryMonth, response.ExpiryMonth)
				assert.Equal(t, cardTokenResponse.ExpiryYear, response.ExpiryYear)
				assert.Equal(t, cardTokenResponse.Name, response.Name)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *CardTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: CardTokenRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			checker: func(response *CardTokenResponse, err error) {
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

			tc.checker(client.RequestCardToken(tc.request))
		})
	}
}
