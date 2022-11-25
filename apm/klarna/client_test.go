package klarna

import (
	"github.com/checkout/checkout-sdk-go-beta/payments"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
	"github.com/checkout/checkout-sdk-go-beta/errors"
	"github.com/checkout/checkout-sdk-go-beta/mocks"
)

func TestCreateSession(t *testing.T) {
	var (
		httpMetadata = common.HttpMetadata{
			Status:     "201 Created",
			StatusCode: http.StatusCreated,
		}

		session = CreditSessionResponse{
			HttpMetadata: httpMetadata,
			SessionId:    "session_id",
			ClientToken:  "client_token",
		}
	)

	cases := []struct {
		name             string
		request          CreditSessionRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CreditSessionResponse, error)
	}{
		{
			name: "when request is correct then create klarna session",
			request: CreditSessionRequest{
				PurchaseCountry: common.GB,
				Currency:        common.GBP,
				Locale:          "en-GB",
				Amount:          1000,
				TaxAmount:       1,
				Products:        getKlarnaProduct(),
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*CreditSessionResponse)
						*respMapping = session
					})
			},
			checker: func(response *CreditSessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, session.SessionId, response.SessionId)
				assert.Equal(t, session.ClientToken, response.ClientToken)
			},
		},
		{
			name:    "when request is invalid then return error",
			request: CreditSessionRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable Entity",
							Data: &errors.ErrorDetails{
								ErrorType: "request_invalid",
								ErrorCodes: []string{
									"amount_required",
								},
							},
						})
			},
			checker: func(response *CreditSessionResponse, err error) {
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

			tc.checker(client.CreateCreditSession(tc.request))
		})
	}
}

func getKlarnaProduct() []map[string]interface{} {
	return []map[string]interface{}{
		{
			"name":             "test product",
			"quantity":         1,
			"unit_price":       1000,
			"tax_rate":         0,
			"total_amount":     1000,
			"total_tax_amount": 0,
		},
	}
}

func TestGetCreditSession(t *testing.T) {
	var (
		sessionId = "session_id"

		httpMetadata = common.HttpMetadata{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		}

		session = CreditSession{
			HttpMetadata:    httpMetadata,
			ClientToken:     "client_token",
			PurchaseCountry: string(common.GB),
			Currency:        string(common.GBP),
			Amount:          100,
			TaxAmount:       0,
		}
	)

	cases := []struct {
		name             string
		sessionId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*CreditSession, error)
	}{
		{
			name:      "when session exists then return session",
			sessionId: sessionId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*CreditSession)
						*respMapping = session
					})
			},
			checker: func(response *CreditSession, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, session.ClientToken, response.ClientToken)
				assert.Equal(t, session.PurchaseCountry, response.PurchaseCountry)
				assert.Equal(t, session.Currency, response.Currency)
				assert.Equal(t, session.Amount, response.Amount)
			},
		},
		{
			name:      "when session not found then return error",
			sessionId: "invalid_session_id",
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
						},
					)
			},
			checker: func(response *CreditSession, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{})
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCreditSession(tc.sessionId))
		})
	}
}

func TestCapturePayment(t *testing.T) {
	var (
		httpMetadata = common.HttpMetadata{
			Status:     "202 Accepted",
			StatusCode: http.StatusAccepted,
		}

		captureResponse = CaptureResponse{
			HttpMetadata: httpMetadata,
			ActionId:     "action_id",
			Reference:    "reference",
		}
	)

	cases := []struct {
		name             string
		paymentId        string
		request          OrderCaptureRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CaptureResponse, error)
	}{
		{
			name:      "when request is correct then capture payment",
			paymentId: "1234",
			request: OrderCaptureRequest{
				Type:      payments.KlarnaSource,
				Amount:    1000,
				Reference: "reference",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*CaptureResponse)
						*respMapping = captureResponse
					})
			},
			checker: func(response *CaptureResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
				assert.Equal(t, captureResponse.ActionId, response.ActionId)
				assert.Equal(t, captureResponse.Reference, response.Reference)
			},
		},
		{
			name:      "when request is invalid then return error",
			paymentId: "1234",
			request:   OrderCaptureRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable Entity",
							Data: &errors.ErrorDetails{
								ErrorType: "request_invalid",
								ErrorCodes: []string{
									"amount_required",
								},
							},
						})
			},
			checker: func(response *CaptureResponse, err error) {
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

			tc.checker(client.CapturePayment(tc.paymentId, tc.request))
		})
	}
}

func TestVoidPayment(t *testing.T) {
	var (
		httpMetadata = common.HttpMetadata{
			Status:     "202 Accepted",
			StatusCode: http.StatusAccepted,
		}

		voidResponse = payments.VoidResponse{
			HttpMetadata: httpMetadata,
			ActionId:     "action_id",
		}
	)

	cases := []struct {
		name             string
		paymentId        string
		request          payments.VoidRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*payments.VoidResponse, error)
	}{
		{
			name:      "when request is correct then void payment",
			paymentId: "1234",
			request: payments.VoidRequest{
				Reference: "reference",
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*payments.VoidResponse)
						*respMapping = voidResponse
					})
			},
			checker: func(response *payments.VoidResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
				assert.Equal(t, voidResponse.ActionId, response.ActionId)
				assert.Equal(t, voidResponse.Reference, response.Reference)
			},
		},
		{
			name:      "when request is invalid then return error",
			paymentId: "1234",
			request:   payments.VoidRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable Entity",
							Data: &errors.ErrorDetails{
								ErrorType: "request_invalid",
								ErrorCodes: []string{
									"amount_required",
								},
							},
						})
			},
			checker: func(response *payments.VoidResponse, err error) {
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

			tc.checker(client.VoidPayment(tc.paymentId, tc.request))
		})
	}
}
