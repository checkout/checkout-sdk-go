package metadata

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/metadata/sources"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestRequestQuote(t *testing.T) {
	var (
		response = CardMetadataResponse{
			HttpMetadata:      mocks.HttpMetadataStatusOk,
			Bin:               "45434720",
			Scheme:            "visa",
			SchemeLocal:       CartesBancaires,
			CardType:          common.Credit,
			CardCategory:      common.Consumer,
			Issuer:            "STATE BANK OF MAURITIUS",
			IssuerCountry:     common.MU,
			IssuerCountryName: "Mauritius",
			ProductId:         "F",
			ProductType:       "Visa Classic",
			CardPayouts: &CardMetadataPayouts{
				DomesticNonMoneyTransfer:    NotSupported,
				CrossBorderNonMoneyTransfer: NotSupported,
				DomesticGambling:            NotSupported,
				CrossBorderGambling:         NotSupported,
				DomesticMoneyTransfer:       NotSupported,
				CrossBorderMoneyTransfer:    NotSupported,
			},
		}
	)

	cases := []struct {
		name             string
		request          CardMetadataRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CardMetadataResponse, error)
	}{
		{
			name: "when request is correct then should request card metadata",
			request: CardMetadataRequest{
				Source: sources.NewRequestBinSource("4539467987109256"),
				Format: Basic,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CardMetadataResponse)
						*respMapping = response
					})
			},
			checker: func(response *CardMetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
			},
		},
		{
			name: "when request is not correct then return error",
			request: CardMetadataRequest{
				Source: sources.NewRequestBinSource("4539467987109256"),
				Format: Basic,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			checker: func(response *CardMetadataResponse, err error) {
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
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.RequestCardMetadata(tc.request))
		})
	}
}
