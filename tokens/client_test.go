package tokens

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
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

func TestRequestCvvToken(t *testing.T) {
	var (
		cvvTokenRequest = CvvTokenRequest{
			Type:      Cvv,
			TokenData: CvvTokenData{Cvv: "100"},
		}

		cvvTokenResponse = CvvTokenResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Type:         Cvv,
			Token:        tokenId,
		}
	)

	cases := []struct {
		name             string
		request          CvvTokenRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CvvTokenResponse, error)
	}{
		{
			name:    "when request is correct then return cvv token",
			request: cvvTokenRequest,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CvvTokenResponse)
						*respMapping = cvvTokenResponse
					})
			},
			checker: func(response *CvvTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, Cvv, response.Type)
				assert.Equal(t, tokenId, response.Token)
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
			checker: func(response *CvvTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
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

			tc.checker(client.RequestCvvToken(tc.request))
		})
	}
}

func TestGetTokenMetadata(t *testing.T) {
	var (
		tokenMetadataResponse = TokenMetadataResponse{
			HttpMetadata:  mocks.HttpMetadataStatusOk,
			Token:         tokenId,
			Type:          "card",
			ExpiryMonth:   expiryMonth,
			ExpiryYear:    expiryYear,
			Scheme:        "Visa",
			Last4:         "4242",
			Bin:           "424242",
			CardType:      "CREDIT",
			CardCategory:  "CONSUMER",
			Issuer:        "JPMORGAN CHASE BANK NA",
			IssuerCountry: "US",
			ProductId:     "A",
			ProductType:   "Visa Traditional",
			BillingAddress: &TokenMetadataBillingAddress{
				City:    "London",
				Country: "GB",
			},
		}
	)

	cases := []struct {
		name             string
		tokenId          string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*TokenMetadataResponse, error)
	}{
		{
			name:    "when token id is valid then return token metadata",
			tokenId: tokenId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*TokenMetadataResponse)
						*respMapping = tokenMetadataResponse
					})
			},
			checker: func(response *TokenMetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, tokenId, response.Token)
				assert.Equal(t, "card", response.Type)
				assert.Equal(t, "CREDIT", response.CardType)
				assert.Equal(t, "CONSUMER", response.CardCategory)
				assert.Equal(t, "4242", response.Last4)
				assert.Equal(t, "424242", response.Bin)
				assert.NotNil(t, response.BillingAddress)
				assert.Equal(t, "London", response.BillingAddress.City)
				assert.Equal(t, "GB", response.BillingAddress.Country)
			},
		},
		{
			name:    "when credentials invalid then return error",
			tokenId: tokenId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *TokenMetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when token not found then return 404",
			tokenId: "tok_unknown",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *TokenMetadataResponse, err error) {
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
			client := NewClient(config, apiClient)

			tc.checker(client.GetTokenMetadata(tc.tokenId))
		})
	}
}

func TestRequestPinToken(t *testing.T) {
	var (
		pinTokenRequest = PinTokenRequest{
			Type:      Pin,
			TokenData: PinTokenData{Pin: "1234"},
		}

		pinTokenResponse = PinTokenResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Type:         Pin,
			Token:        tokenId,
		}
	)

	cases := []struct {
		name             string
		request          PinTokenRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*PinTokenResponse, error)
	}{
		{
			name:    "when request is correct then return pin token",
			request: pinTokenRequest,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*PinTokenResponse)
						*respMapping = pinTokenResponse
					})
			},
			checker: func(response *PinTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, Pin, response.Type)
				assert.Equal(t, tokenId, response.Token)
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
			checker: func(response *PinTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
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

			tc.checker(client.RequestPinToken(tc.request))
		})
	}
}
