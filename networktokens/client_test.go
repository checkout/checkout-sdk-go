package networktokens

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
	"github.com/checkout/checkout-sdk-go/v2/payments"
)

// tests

func TestProvisionNetworkToken(t *testing.T) {
	var (
		tokenResponse = NetworkTokenResponse{
			Card: &NetworkTokenCard{
				Last4:       "6378",
				ExpiryMonth: "5",
				ExpiryYear:  "2025",
			},
			NetworkToken: &NetworkTokenDetails{
				Id:    "nt_xgu3isllqfyu7ktpk5z2yxbwna",
				State: Active,
				Type:  payments.Vts,
			},
		}
	)

	cases := []struct {
		name             string
		request          ProvisionNetworkTokenRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*NetworkTokenResponse, error)
	}{
		{
			name:    "when request is correct then provision network token",
			request: buildProvisionWithIdSource(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*NetworkTokenResponse)
						*respMapping = tokenResponse
					})
			},
			checker: func(response *NetworkTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Card)
				assert.NotNil(t, response.NetworkToken)
				assert.Equal(t, tokenResponse.NetworkToken.Id, response.NetworkToken.Id)
				assert.Equal(t, tokenResponse.NetworkToken.State, response.NetworkToken.State)
			},
		},
		{
			name:    "when credentials invalid then return error",
			request: buildProvisionWithIdSource(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *NetworkTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:    "when api returns error then propagate error",
			request: buildProvisionWithIdSource(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusUnprocessableEntity})
			},
			checker: func(response *NetworkTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildNetworkTokensClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.ProvisionNetworkToken(tc.request))
		})
	}
}

func TestGetNetworkToken(t *testing.T) {
	var (
		networkTokenId = "nt_xgu3isllqfyu7ktpk5z2yxbwna"
		tokenResponse  = NetworkTokenResponse{
			Card: &NetworkTokenCard{
				Last4:       "6378",
				ExpiryMonth: "5",
				ExpiryYear:  "2025",
			},
			NetworkToken: &NetworkTokenDetails{
				Id:    networkTokenId,
				State: Active,
				Type:  payments.Vts,
			},
		}
	)

	cases := []struct {
		name             string
		networkTokenId   string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*NetworkTokenResponse, error)
	}{
		{
			name:           "when token id is valid then return network token",
			networkTokenId: networkTokenId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*NetworkTokenResponse)
						*respMapping = tokenResponse
					})
			},
			checker: func(response *NetworkTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Card)
				assert.NotNil(t, response.NetworkToken)
				assert.Equal(t, tokenResponse.NetworkToken.Id, response.NetworkToken.Id)
				assert.Equal(t, tokenResponse.NetworkToken.State, response.NetworkToken.State)
			},
		},
		{
			name:           "when credentials invalid then return error",
			networkTokenId: networkTokenId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *NetworkTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:           "when token not found then return error",
			networkTokenId: "nt_invalid",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
			},
			checker: func(response *NetworkTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildNetworkTokensClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.GetNetworkToken(tc.networkTokenId))
		})
	}
}

func TestRequestCryptogram(t *testing.T) {
	var (
		networkTokenId      = "nt_xgu3isllqfyu7ktpk5z2yxbwna"
		cryptogramResponse  = RequestCryptogramResponse{
			Cryptogram: "AhJ3hnYAoAbVz5zg1e17MAACAAA=",
			Eci:        "7",
		}
	)

	cases := []struct {
		name             string
		networkTokenId   string
		request          RequestCryptogramRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*RequestCryptogramResponse, error)
	}{
		{
			name:           "when request is correct then return cryptogram",
			networkTokenId: networkTokenId,
			request:        buildRequestCryptogramRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*RequestCryptogramResponse)
						*respMapping = cryptogramResponse
					})
			},
			checker: func(response *RequestCryptogramResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, cryptogramResponse.Cryptogram, response.Cryptogram)
				assert.Equal(t, cryptogramResponse.Eci, response.Eci)
			},
		},
		{
			name:           "when credentials invalid then return error",
			networkTokenId: networkTokenId,
			request:        buildRequestCryptogramRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *RequestCryptogramResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:           "when token not found then return error",
			networkTokenId: "nt_invalid",
			request:        buildRequestCryptogramRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
			},
			checker: func(response *RequestCryptogramResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildNetworkTokensClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.RequestCryptogram(tc.networkTokenId, tc.request))
		})
	}
}

func TestDeleteNetworkToken(t *testing.T) {
	var (
		networkTokenId = "nt_xgu3isllqfyu7ktpk5z2yxbwna"
	)

	cases := []struct {
		name             string
		networkTokenId   string
		request          DeleteNetworkTokenRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPatch         func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:           "when request is correct then delete network token",
			networkTokenId: networkTokenId,
			request:        buildDeleteNetworkTokenRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*common.MetadataResponse)
						*respMapping = common.MetadataResponse{HttpMetadata: mocks.HttpMetadataStatusNoContent}
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:           "when credentials invalid then return error",
			networkTokenId: networkTokenId,
			request:        buildDeleteNetworkTokenRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:           "when token not found then return error",
			networkTokenId: "nt_invalid",
			request:        buildDeleteNetworkTokenRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusNotFound})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildNetworkTokensClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPatch(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.DeleteNetworkToken(tc.networkTokenId, tc.request))
		})
	}
}

// common methods

func buildNetworkTokensClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}

func buildProvisionWithIdSource() ProvisionNetworkTokenRequest {
	src := NewIdSource()
	src.Id = "src_wmlfc3zyhqzehihu7giusaaawu"
	return ProvisionNetworkTokenRequest{Source: src}
}

func buildRequestCryptogramRequest() RequestCryptogramRequest {
	return RequestCryptogramRequest{TransactionType: Ecom}
}

func buildDeleteNetworkTokenRequest() DeleteNetworkTokenRequest {
	return DeleteNetworkTokenRequest{
		InitiatedBy: TokenRequestor,
		Reason:      Other,
	}
}
