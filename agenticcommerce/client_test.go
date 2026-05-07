package agenticcommerce

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"
)

// tests

func TestCreateDelegatedPaymentToken(t *testing.T) {
	expiresAt := time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC)
	var (
		tokenResponse = CreateDelegatedPaymentTokenResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "vt_abc123def456ghi789",
			Metadata:     map[string]string{"psp": "checkout.com"},
		}
	)

	cases := []struct {
		name             string
		request          CreateDelegatedPaymentTokenRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CreateDelegatedPaymentTokenResponse, error)
	}{
		{
			name:    "when request is correct then create delegated payment token",
			request: buildCreateDelegatedPaymentTokenRequest(expiresAt),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CreateDelegatedPaymentTokenResponse)
						*respMapping = tokenResponse
					})
			},
			checker: func(response *CreateDelegatedPaymentTokenResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, tokenResponse.Id, response.Id)
				assert.NotNil(t, response.Metadata)
			},
		},
		{
			name:    "when credentials invalid then return error",
			request: buildCreateDelegatedPaymentTokenRequest(expiresAt),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *CreateDelegatedPaymentTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization", chkErr.Error())
			},
		},
		{
			name:    "when api returns error then propagate error",
			request: buildCreateDelegatedPaymentTokenRequest(expiresAt),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{StatusCode: http.StatusUnprocessableEntity})
			},
			checker: func(response *CreateDelegatedPaymentTokenResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient, credentials, config := buildAgenticCommerceClientConfig()
			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			client := NewClient(config, apiClient)
			tc.checker(client.CreateDelegatedPaymentToken(tc.request, nil))
		})
	}
}

// common methods

func buildAgenticCommerceClientConfig() (*mocks.ApiClientMock, *mocks.CredentialsMock, *configuration.Configuration) {
	apiClient := new(mocks.ApiClientMock)
	credentials := new(mocks.CredentialsMock)
	environment := new(mocks.EnvironmentMock)
	enableTelemetry := true
	config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
	return apiClient, credentials, config
}

func buildCreateDelegatedPaymentTokenRequest(expiresAt time.Time) CreateDelegatedPaymentTokenRequest {
	card := NewDelegatedPaymentMethodCard()
	card.CardNumberType = Fpan
	card.Number = "4242424242424242"
	card.ExpMonth = "11"
	card.ExpYear = "2026"
	card.Metadata = map[string]string{"issuing_bank": "test"}

	return CreateDelegatedPaymentTokenRequest{
		PaymentMethod: *card,
		Allowance: DelegatedPaymentAllowance{
			Reason:            OneTime,
			MaxAmount:         10000,
			Currency:          common.USD,
			MerchantId:        "cli_vkuhvk4vjn2edkps7dfsq6emqm",
			CheckoutSessionId: "1PQrsT",
			ExpiresAt:         &expiresAt,
		},
		RiskSignals: []DelegatedPaymentRiskSignal{
			{Type: "card_testing", Score: 10, Action: "blocked"},
		},
		Metadata: map[string]string{"campaign": "q4"},
		Headers: DelegatedPaymentHeaders{
			Signature: "eyJtZX...",
			Timestamp: "2026-03-11T10:30:00Z",
		},
	}
}
