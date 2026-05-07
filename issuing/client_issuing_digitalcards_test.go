package issuing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"

	digitalcards "github.com/checkout/checkout-sdk-go/v2/issuing/digitalcards"
)

// # tests

func TestGetDigitalCard(t *testing.T) {
	var (
		response = digitalcards.GetDigitalCardResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "dcr_5ngxzsynm2me3oxf73esbhda6q",
			CardId:       "crd_fa6psq242dcd6fdn5gifcq1491",
			LastFour:     "4242",
			Status:       digitalcards.DigitalCardActive,
		}
	)

	cases := []struct {
		name             string
		digitalCardId    string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*digitalcards.GetDigitalCardResponse, error)
	}{
		{
			name:          "when request is correct then should return 200",
			digitalCardId: "dcr_5ngxzsynm2me3oxf73esbhda6q",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*digitalcards.GetDigitalCardResponse)
						*respMapping = response
					})
			},
			checker: assertGetDigitalCardSuccess(t, &response),
		},
		{
			name:          "when credentials are invalid then return authorization error",
			digitalCardId: "dcr_5ngxzsynm2me3oxf73esbhda6q",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *digitalcards.GetDigitalCardResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:          "when digital card not found then return 404",
			digitalCardId: "dcr_not_found",
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
			checker: func(response *digitalcards.GetDigitalCardResponse, err error) {
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

			tc.checker(client.GetDigitalCard(tc.digitalCardId))
		})
	}
}

// # common methods

func assertGetDigitalCardSuccess(t *testing.T, expected *digitalcards.GetDigitalCardResponse) func(*digitalcards.GetDigitalCardResponse, error) {
	return func(response *digitalcards.GetDigitalCardResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
		assert.Equal(t, expected.Id, response.Id)
		assert.Equal(t, expected.CardId, response.CardId)
		assert.Equal(t, expected.LastFour, response.LastFour)
		assert.Equal(t, expected.Status, response.Status)

	}
}
