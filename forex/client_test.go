package forex

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestRequestQuote(t *testing.T) {
	var (
		now = time.Now()

		quote = QuoteResponse{
			HttpMetadata:        mocks.HttpMetadataStatusOk,
			Id:                  "qte_id",
			SourceCurrency:      common.GBP,
			SourceAmount:        30000,
			DestinationCurrency: common.USD,
			DestinationAmount:   35700,
			Rate:                1.19,
			ExpiresOn:           &now,
			IsSingleUse:         false,
		}
	)

	cases := []struct {
		name             string
		request          QuoteRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*QuoteResponse, error)
	}{
		{
			name: "when request is correct then should request quote",
			request: QuoteRequest{
				SourceCurrency:      common.GBP,
				SourceAmount:        30000,
				DestinationCurrency: common.USD,
				ProcessingChannelId: "pc_abcdefghijklmnopqrstuvwxyz",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*QuoteResponse)
						*respMapping = quote
					})
			},
			checker: func(response *QuoteResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Rate)
				assert.NotNil(t, response.ExpiresOn)
			},
		},
		{
			name: "when request is not correct then return error",
			request: QuoteRequest{
				ProcessingChannelId: "pc_abcdefghijklmnopqrstuvwxyz",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: errors.ErrorDetails{
								ErrorType:  "request_invalid",
								ErrorCodes: []string{"source_currency_required"},
							},
						})
			},
			checker: func(response *QuoteResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "source_currency_required")
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

			tc.checker(client.RequestQuote(tc.request))
		})
	}
}

func TestGetRates(t *testing.T) {
	ratesResponse := RatesResponse{
		Product: "card_payouts",
		Source:  Visa,
		Rates: []Rate{
			{
				ExchangeRate: 1.14208777,
				CurrencyPair: "GBPEUR",
			},
			{
				ExchangeRate: 0.83708142,
				CurrencyPair: "USDGBP",
			},
		},
	}

	invalidDataError := errors.ErrorDetails{
		RequestID:  "0HL80RJLS76I7",
		ErrorType:  "request_invalid",
		ErrorCodes: []string{"payment_source_required"},
	}

	cases := []struct {
		name             string
		query            RatesQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*RatesResponse, error)
	}{
		{
			name: "when query params are valid then return rates",
			query: RatesQuery{
				Product:             "card_payouts",
				Source:              Visa,
				CurrencyPairs:       "GBPEUR,USDNOK,JPNCAD",
				ProcessingChannelId: "pc_vxt6yftthv4e5flqak6w2i7rim",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*RatesResponse)
						*respMapping = ratesResponse
					})
			},
			checker: func(response *RatesResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Product)
				assert.NotNil(t, response.Source)
				assert.NotNil(t, response.Rates)
				assert.Equal(t, ratesResponse.Product, response.Product)
				assert.Equal(t, ratesResponse.Source, response.Source)
				assert.Equal(t, ratesResponse.Rates, response.Rates)
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
			checker: func(response *RatesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name: "when invalid data then return error",
			query: RatesQuery{
				Product:             "card_payouts",
				CurrencyPairs:       "GBPEUR,USDNOK,JPNCAD",
				ProcessingChannelId: "pc_vxt6yftthv4e5flqak6w2i7rim",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable Entity",
							Data:       invalidDataError,
						})
			},
			checker: func(response *RatesResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "422 Unprocessable Entity", chkErr.Status)
				assert.Equal(t, invalidDataError, chkErr.Data)
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

			tc.checker(client.GetRates(tc.query))
		})
	}
}
