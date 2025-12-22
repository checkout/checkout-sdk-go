package balances

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestRetrieveEntityBalances(t *testing.T) {
	var (
		response = QueryResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Data: []AccountBalance{
				{
					Descriptor:      "Revenue Account 1",
					HoldingCurrency: string(common.EUR),
					Balances: Balances{
						Pending:    10,
						Available:  50,
						Payable:    0,
						Collateral: 0,
					}},
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
