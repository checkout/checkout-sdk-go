package applepay

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
)

// tests

func TestUploadPaymentProcessingCertificate(t *testing.T) {
	var (
		certificateResponse = UploadCertificateResponse{
			Id:            "aplc_hefptsiydvkexnzzb35zrlqgfq",
			PublicKeyHash: "tqYV+tmG9aMh+l/K6cicUnPqkb1gUiLjSTM9gEz6Nl0=",
		}
	)

	cases := []struct {
		name             string
		request          UploadCertificateRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*UploadCertificateResponse, error)
	}{
		{
			name:    "when request is correct then upload certificate",
			request: buildUploadCertificateRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*UploadCertificateResponse)
						*respMapping = certificateResponse
					})
			},
			checker: func(response *UploadCertificateResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, certificateResponse.Id, response.Id)
				assert.Equal(t, certificateResponse.PublicKeyHash, response.PublicKeyHash)
			},
		},
		{
			name:    "when credentials invalid then return error",
			request: buildUploadCertificateRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *UploadCertificateResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:    "when api returns error then propagate error",
			request: buildUploadCertificateRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusUnprocessableEntity})
			},
			checker: func(response *UploadCertificateResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildApplePayClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.UploadPaymentProcessingCertificate(tc.request))
		})
	}
}

func TestEnrollDomain(t *testing.T) {
	cases := []struct {
		name             string
		request          EnrollDomainRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:    "when request is correct then enroll domain",
			request: buildEnrollDomainRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			name:    "when credentials invalid then return error",
			request: buildEnrollDomainRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			name:    "when api returns error then propagate error",
			request: buildEnrollDomainRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusUnprocessableEntity})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildApplePayClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.EnrollDomain(tc.request))
		})
	}
}

func TestGenerateCertificateSigningRequest(t *testing.T) {
	var (
		signingRequestResponse = GenerateCertificateSigningRequestResponse{
			Content: "-----BEGIN CERTIFICATE REQUEST-----\nMIIBSDCB8AIBADCBj...\n-----END CERTIFICATE REQUEST-----",
		}
	)

	cases := []struct {
		name             string
		request          GenerateCertificateSigningRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*GenerateCertificateSigningRequestResponse, error)
	}{
		{
			name:    "when request uses ec_v1 protocol then generate certificate signing request",
			request: GenerateCertificateSigningRequest{ProtocolVersion: EcV1},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*GenerateCertificateSigningRequestResponse)
						*respMapping = signingRequestResponse
					})
			},
			checker: func(response *GenerateCertificateSigningRequestResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, signingRequestResponse.Content, response.Content)
			},
		},
		{
			name:    "when request uses rsa_v1 protocol then generate certificate signing request",
			request: GenerateCertificateSigningRequest{ProtocolVersion: RsaV1},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*GenerateCertificateSigningRequestResponse)
						*respMapping = signingRequestResponse
					})
			},
			checker: func(response *GenerateCertificateSigningRequestResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.Content)
			},
		},
		{
			name:    "when credentials invalid then return error",
			request: GenerateCertificateSigningRequest{ProtocolVersion: EcV1},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *GenerateCertificateSigningRequestResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:    "when api returns error then propagate error",
			request: GenerateCertificateSigningRequest{ProtocolVersion: EcV1},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusUnprocessableEntity})
			},
			checker: func(response *GenerateCertificateSigningRequestResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildApplePayClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.GenerateCertificateSigningRequest(tc.request))
		})
	}
}

// common methods

func buildApplePayClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}

func buildUploadCertificateRequest() UploadCertificateRequest {
	return UploadCertificateRequest{
		Content: "MIIEfTCCBCOgAwIBAgIID/asezaWNycwCgYIKoZIzj0EAwIwgYAxNDAy...",
	}
}

func buildEnrollDomainRequest() EnrollDomainRequest {
	return EnrollDomainRequest{
		Domain: "https://example.com",
	}
}
