package issuing

import (
	"github.com/checkout/checkout-sdk-go/common"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestCreateCard(t *testing.T) {
	request := NewVirtualCardRequest()
	response := CardResponse{}

	cases := []struct {
		name             string
		request          CardRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CardResponse, error)
	}{
		{
			name:    "when create a card and this request is correct then should return a response",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*CardResponse)
						*respMapping = response
					})
			},
			checker: func(response *CardResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.CreateCard(tc.request))
		})
	}
}

func TestGetCardDetails(t *testing.T) {
	response := CardDetailsResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CardDetailsResponse, error)
	}{
		{
			name:   "when get a card and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*CardDetailsResponse)
						*respMapping = response
					})
			},
			checker: func(response *CardDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.GetCardDetails(tc.cardId))
		})
	}
}

func TestEnrollThreeDS(t *testing.T) {
	request := ThreeDSEnrollmentRequest{}
	response := ThreeDSEnrollmentResponse{}

	cases := []struct {
		name             string
		cardId           string
		request          ThreeDSEnrollmentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*ThreeDSEnrollmentResponse, error)
	}{
		{
			name:    "when enroll a card three DS and this request is correct then should return a response",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*ThreeDSEnrollmentResponse)
						*respMapping = response
					})
			},
			checker: func(response *ThreeDSEnrollmentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.EnrollThreeDS(tc.cardId, tc.request))
		})
	}
}

func TestUpdateThreeDS(t *testing.T) {
	request := ThreeDSUpdateRequest{}
	response := ThreeDSUpdateResponse{}

	cases := []struct {
		name             string
		cardId           string
		request          ThreeDSUpdateRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*ThreeDSUpdateResponse, error)
	}{
		{
			name:    "when update a card enroll three DS and this request is correct then should return 201",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Patch", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*ThreeDSUpdateResponse)
						*respMapping = response
					})
			},
			checker: func(response *ThreeDSUpdateResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.UpdateThreeDS(tc.cardId, tc.request))
		})
	}
}

func TestGetCardThreeDSDetails(t *testing.T) {
	response := ThreeDSEnrollmentDetailsResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*ThreeDSEnrollmentDetailsResponse, error)
	}{
		{
			name:   "when get a card enroll three DS details and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*ThreeDSEnrollmentDetailsResponse)
						*respMapping = response
					})
			},
			checker: func(response *ThreeDSEnrollmentDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.GetCardThreeDSDetails(tc.cardId))
		})
	}
}

func TestActivateCard(t *testing.T) {
	response := common.IdResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:   "when activate a card and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.ActivateCard(tc.cardId))
		})
	}
}

func TestGetCardCredentials(t *testing.T) {
	response := CardCredentialsResponse{}

	cases := []struct {
		name             string
		cardId           string
		query            CardCredentialsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CardCredentialsResponse, error)
	}{
		{
			name:   "when get card credentials and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			query:  CardCredentialsQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*CardCredentialsResponse)
						*respMapping = response
					})
			},
			checker: func(response *CardCredentialsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.GetCardCredentials(tc.cardId, tc.query))
		})
	}
}

func TestRevokeCard(t *testing.T) {
	request := RevokeCardRequest{}
	response := common.IdResponse{}

	cases := []struct {
		name             string
		cardId           string
		request          RevokeCardRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:    "when revoke a card and this request is correct then should return a response",
			cardId:  "",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.RevokeCard(tc.cardId, tc.request))
		})
	}
}

func TestSuspendCard(t *testing.T) {
	response := common.IdResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:   "when suspend a card and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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

			tc.checker(client.SuspendCard(tc.cardId))
		})
	}
}
