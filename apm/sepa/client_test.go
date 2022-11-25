package sepa

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
	"github.com/checkout/checkout-sdk-go-beta/errors"
	"github.com/checkout/checkout-sdk-go-beta/mocks"
)

func TestGetMandate(t *testing.T) {
	var (
		httpMetadata = common.HttpMetadata{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		}

		mandate = MandateResponse{
			HttpMetadata:     httpMetadata,
			MandateReference: "reference",
			CustomerId:       "cus_1234",
			FirstName:        "Bruce",
			LastName:         "Wayne",
			City:             "Gotham",
		}
	)

	cases := []struct {
		name             string
		mandateId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*MandateResponse, error)
	}{
		{
			name:      "when mandate exists then return mandate info",
			mandateId: "mandate_id",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*MandateResponse)
						*respMapping = mandate
					})
			},
			checker: func(response *MandateResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, mandate.MandateReference, response.MandateReference)
				assert.Equal(t, mandate.CustomerId, response.CustomerId)
			},
		},
		{
			name:      "when mandate not found then return error",
			mandateId: "not_found",
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
			checker: func(response *MandateResponse, err error) {
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetMandate(tc.mandateId))
		})
	}
}

func TestCancelMandate(t *testing.T) {
	var (
		link = "https://test-link.com"

		httpMetadata = common.HttpMetadata{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		}

		cancelResponse = SepaResource{
			HttpMetadata: httpMetadata,
			Links: map[string]common.Link{
				"payment": {
					HRef: &link,
				},
			},
		}
	)

	cases := []struct {
		name             string
		mandateId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*SepaResource, error)
	}{
		{
			name:      "when request is correct then cancel mandate",
			mandateId: "1234",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*SepaResource)
						*respMapping = cancelResponse
					})
			},
			checker: func(response *SepaResource, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, cancelResponse.Links, response.Links)
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

			tc.checker(client.CancelMandate(tc.mandateId))
		})
	}
}
