package issuing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"

	transactions "github.com/checkout/checkout-sdk-go/v2/issuing/transactions"
)

// # tests

func TestGetListTransactions(t *testing.T) {
	var (
		response = transactions.TransactionsListResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			TotalCount:   &[]int{1}[0],
			Data: []transactions.TransactionResponse{
				{Id: "trx_test_abcdefghijklmnopqr"},
			},
		}
	)

	cases := []struct {
		name             string
		query            transactions.TransactionsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*transactions.TransactionsListResponse, error)
	}{
		{
			name:  "when query is correct then should return 200",
			query: buildTransactionsQuery(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*transactions.TransactionsListResponse)
						*respMapping = response
					})
			},
			checker: assertTransactionsListSuccess(t, &response),
		},
		{
			name:  "when credentials are invalid then return authorization error",
			query: transactions.TransactionsQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *transactions.TransactionsListResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:  "when server has problems then return 500",
			query: transactions.TransactionsQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusInternalServerError,
						Status:     "500 Internal Server Error",
					})
			},
			checker: func(response *transactions.TransactionsListResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusInternalServerError, chkErr.StatusCode)
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

			tc.checker(client.GetListTransactions(tc.query))
		})
	}
}

func TestGetSingleTransaction(t *testing.T) {
	var (
		response = transactions.TransactionResponse{
			HttpMetadata:    mocks.HttpMetadataStatusOk,
			Id:              "trx_test_abcdefghijklmnopqr",
			Status:          transactions.TransactionAuthorized,
			TransactionType: transactions.Purchase,
		}
	)

	cases := []struct {
		name             string
		transactionId    string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*transactions.TransactionResponse, error)
	}{
		{
			name:          "when request is correct then should return 200",
			transactionId: "trx_test_abcdefghijklmnopqr",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*transactions.TransactionResponse)
						*respMapping = response
					})
			},
			checker: assertTransactionSuccess(t, &response),
		},
		{
			name:          "when transaction not found then return 404",
			transactionId: "trx_not_found",
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
			checker: func(response *transactions.TransactionResponse, err error) {
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

			tc.checker(client.GetSingleTransaction(tc.transactionId))
		})
	}
}

// # common methods

func buildTransactionsQuery() transactions.TransactionsQuery {
	return transactions.TransactionsQuery{
		Limit:  10,
		CardId: "crd_fa6psq242dcd6fdn5gifcq1491",
	}
}

func assertTransactionsListSuccess(t *testing.T, expected *transactions.TransactionsListResponse) func(*transactions.TransactionsListResponse, error) {
	return func(response *transactions.TransactionsListResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
		assert.NotNil(t, response.TotalCount)
		assert.NotNil(t, response.Data)
	}
}

func assertTransactionSuccess(t *testing.T, expected *transactions.TransactionResponse) func(*transactions.TransactionResponse, error) {
	return func(response *transactions.TransactionResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
		assert.Equal(t, expected.Id, response.Id)
		assert.Equal(t, expected.Status, response.Status)
		assert.Equal(t, expected.TransactionType, response.TransactionType)
	}
}
