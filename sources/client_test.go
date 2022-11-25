package sources

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

var (
	sepa  = Sepa
	email = "bruce@wayne-enterprises.com"
	name  = "Bruce Wayne"
	phone = common.Phone{
		CountryCode: "+1",
		Number:      "415 555 2671",
	}
	customer = common.CustomerRequest{
		Email: email,
		Name:  name,
		Phone: &phone,
	}
)

func TestCreateSepaSource(t *testing.T) {
	var (
		httpMetadata = common.HttpMetadata{
			Status:     "201 Created",
			StatusCode: http.StatusCreated,
		}

		sepaResponse = CreateSepaSourceResponse{
			HttpResponse: httpMetadata,
			SourceResponse: &SourceResponse{
				SourceType: sepa,
			},
		}
	)

	cases := []struct {
		name             string
		request          *sepaSourceRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CreateSepaSourceResponse, error)
	}{
		{
			name:    "when request is correct then create sepa source",
			request: getSepaSourceRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*CreateSepaSourceResponse)
						*respMapping = sepaResponse
					})
			},
			checker: func(response *CreateSepaSourceResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpResponse.StatusCode)
				assert.Equal(t, sepaResponse.SourceResponse.SourceType, response.SourceResponse.SourceType)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *CreateSepaSourceResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: NewSepaSourceRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			checker: func(response *CreateSepaSourceResponse, err error) {
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			client := NewClient(configuration, apiClient)

			tc.checker(client.CreateSepaSource(tc.request))
		})
	}
}

func getSepaSourceRequest() *sepaSourceRequest {
	sepaRequest := NewSepaSourceRequest()
	sourceData := SourceData{
		FirstName:   "Bruce",
		LastName:    "Wayne",
		AccountIban: "1234",
	}

	sepaRequest.SourceData = &sourceData
	sepaRequest.Reference = "reference"
	sepaRequest.CustomerRequest = &customer

	return sepaRequest
}
