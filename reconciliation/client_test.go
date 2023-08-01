package reconciliation

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestQueryPaymentsReport(t *testing.T) {
	var (
		from = time.Now().Add(-5 * time.Hour)
		to   = time.Now()

		reportData = PaymentReportData{
			Id:                 "pay_1234",
			ProcessingCurrency: common.USD,
			PayoutCurrency:     common.GBP,
			ChannelName:        "www.example.com",
			Reference:          "ORD-5023-4E89",
		}

		reportResponse = PaymentReportsResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Count:        1,
			Data:         []PaymentReportData{reportData},
		}
	)

	cases := []struct {
		name             string
		query            PaymentReportsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*PaymentReportsResponse, error)
	}{
		{
			name: "when query is correct then return payment reports",
			query: PaymentReportsQuery{
				From:      &from,
				To:        &to,
				Reference: "ORD-5023-4E89",
				Limit:     10,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentReportsResponse)
						*respMapping = reportResponse
					})
			},
			checker: func(response *PaymentReportsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, reportResponse.Data, response.Data)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentReportsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.QueryPaymentsReport(tc.query))
		})
	}
}

func TestGetSinglePaymentReport(t *testing.T) {
	var (
		reportData = PaymentReportData{
			Id:                 "pay_1234",
			ProcessingCurrency: common.USD,
			PayoutCurrency:     common.GBP,
			ChannelName:        "www.example.com",
			Reference:          "ORD-5023-4E89",
		}

		reportResponse = PaymentReportsResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Count:        1,
			Data:         []PaymentReportData{reportData},
		}
	)

	cases := []struct {
		name             string
		paymentId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*PaymentReportsResponse, error)
	}{
		{
			name:      "when payment exists then return payment reports",
			paymentId: "pay_1234",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*PaymentReportsResponse)
						*respMapping = reportResponse
					})
			},
			checker: func(response *PaymentReportsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, reportResponse.Data, response.Data)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentReportsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:      "when payment not found then return error",
			paymentId: "not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *PaymentReportsResponse, err error) {
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetSinglePaymentReport(tc.paymentId))
		})
	}
}

func TestQueryStatementsReport(t *testing.T) {
	var (
		from = time.Now().Add(-5 * time.Hour)
		to   = time.Now()

		reportData = StatementReportsData{
			Id:          "190110B107654",
			PeriodStart: "2023-01-09T00:00:00.000",
			PeriodEnd:   "2023-01-09T23:59:59.000",
			Payouts: []PayoutStatement{
				{
					Id:        "pay_1234",
					Currency:  common.GBP,
					NetAmount: 100,
				},
			},
		}

		reportResponse = StatementReportsResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Count:        1,
			Data:         []StatementReportsData{reportData},
		}
	)

	cases := []struct {
		name             string
		query            common.DateRangeQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*StatementReportsResponse, error)
	}{
		{
			name: "when query is correct then return statement reports",
			query: common.DateRangeQuery{
				From: &from,
				To:   &to,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*StatementReportsResponse)
						*respMapping = reportResponse
					})
			},
			checker: func(response *StatementReportsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, reportResponse.Data, response.Data)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *StatementReportsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.QueryStatementsReport(tc.query))
		})
	}
}

func TestRetrieveCVSPaymentsReport(t *testing.T) {
	var (
		from = time.Now().Add(-5 * time.Hour)
		to   = time.Now()

		contentResponse = common.ContentResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Content:      "content",
		}
	)

	cases := []struct {
		name             string
		query            common.DateRangeQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*common.ContentResponse, error)
	}{
		{
			name: "when query is correct then return payment reports content",
			query: common.DateRangeQuery{
				From: &from,
				To:   &to,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.ContentResponse)
						*respMapping = contentResponse
					})
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, contentResponse.Content, response.Content)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.RetrieveCVSPaymentsReport(tc.query))
		})
	}
}

func TestRetrieveCVSSingleStatementReport(t *testing.T) {
	var (
		contentResponse = common.ContentResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Content:      "content",
		}
	)

	cases := []struct {
		name             string
		statementId      string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*common.ContentResponse, error)
	}{
		{
			name:        "when statement exists then return statement reports content",
			statementId: "stm_1234",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.ContentResponse)
						*respMapping = contentResponse
					})
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, contentResponse.Content, response.Content)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:        "when statement not found then return error",
			statementId: "not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *common.ContentResponse, err error) {
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.RetrieveCVSSingleStatementReport(tc.statementId))
		})
	}
}

func TestRetrieveCVSStatementsReport(t *testing.T) {
	var (
		from = time.Now().Add(-5 * time.Hour)
		to   = time.Now()

		contentResponse = common.ContentResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Content:      "content",
		}
	)

	cases := []struct {
		name             string
		query            common.DateRangeQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*common.ContentResponse, error)
	}{
		{
			name: "when query is correct then return statement reports content",
			query: common.DateRangeQuery{
				From: &from,
				To:   &to,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.ContentResponse)
						*respMapping = contentResponse
					})
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, contentResponse.Content, response.Content)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
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

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.RetrieveCVSStatementsReport(tc.query))
		})
	}
}
