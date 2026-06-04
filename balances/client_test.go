package balances

import (
	"fmt"
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

func TestRetrieveEntityBalances(t *testing.T) {
	balancesAt := time.Date(2026, 5, 6, 13, 59, 59, 0, time.UTC)
	asOf := balancesAt

	var (
		response = QueryResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Data: []AccountBalance{
				{
					Descriptor:      "Revenue Account 1",
					HoldingCurrency: common.EUR,
					Balances: Balances{
						Pending:    10,
						Available:  50,
						Payable:    0,
						Collateral: 0,
					}},
			},
		}

		richResponse = QueryResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Data: []AccountBalance{
				{
					CurrencyAccountId: "ca_g5y7d6jo4e2urgforcbf2ey5jm",
					Descriptor:        "Revenue Account 1",
					HoldingCurrency:   common.EUR,
					BalancesAsOf:      &asOf,
					Balances: Balances{
						Pending:     23000,
						Available:   50000,
						Payable:     2700,
						Collateral:  6000,
						Operational: 7000,
						CollateralBreakdown: &CollateralBreakdown{
							FixedReserve:   4000,
							RollingReserve: 2000,
						},
					},
				},
			},
		}
	)

	cases := []struct {
		name             string
		entityId         string
		query            QueryFilter
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*QueryResponse, error)
	}{
		{
			name:     "when request is correct then return a balance details",
			entityId: "ent_w4jelhppmfiufdnatam37wrfc4",
			query:    QueryFilter{Query: fmt.Sprintf("currency:%s", common.EUR)},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*QueryResponse)
						*respMapping = response
					})
			},
			checker: func(response *QueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
		{
			name:     "when response includes new fields then they are mapped",
			entityId: "ent_w4jelhppmfiufdnatam37wrfc4",
			query: QueryFilter{
				Query:                 fmt.Sprintf("currency:%s", common.EUR),
				WithCurrencyAccountId: true,
				BalancesAt:            &balancesAt,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*QueryResponse)
						*respMapping = richResponse
					})
			},
			checker: func(response *QueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Len(t, response.Data, 1)
				acct := response.Data[0]
				assert.Equal(t, "ca_g5y7d6jo4e2urgforcbf2ey5jm", acct.CurrencyAccountId)
				assert.Equal(t, common.EUR, acct.HoldingCurrency)
				assert.NotNil(t, acct.BalancesAsOf)
				assert.Equal(t, int64(23000), acct.Balances.Pending)
				assert.Equal(t, int64(7000), acct.Balances.Operational)
				assert.NotNil(t, acct.Balances.CollateralBreakdown)
				assert.Equal(t, int64(4000), acct.Balances.CollateralBreakdown.FixedReserve)
				assert.Equal(t, int64(2000), acct.Balances.CollateralBreakdown.RollingReserve)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *QueryResponse, err error) {
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
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.RetrieveEntityBalances(tc.entityId, tc.query))
		})
	}
}
