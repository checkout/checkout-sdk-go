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
		httpMetadata = common.HttpMetadata{
			Status:     "201 Created",
			StatusCode: http.StatusCreated,
		}

		quote = QuoteResponse{
			HttpMetadata:        httpMetadata,
			Id:                  "qte_id",
			SourceCurrency:      common.GBP,
			SourceAmount:        30000,
			DestinationCurrency: common.USD,
			DestinationAmount:   35700,
			Rate:                1.19,
			ExpiresOn:           time.Now(),
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
							Data: &errors.ErrorDetails{
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			client := NewClient(configuration, apiClient)

			tc.checker(client.RequestQuote(tc.request))
		})
	}
}
