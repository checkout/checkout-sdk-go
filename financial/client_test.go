package financial

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestGetFinancialActions(t *testing.T) {
	var (
		financialAction = FinancialAction{
			PaymentId:           "pay_1234",
			ActionId:            "act_1234",
			ActionType:          "Capture",
			EntityId:            "ent_1234",
			SubEntityId:         "ent_4567",
			CurrencyAccountId:   "ca_1234",
			PaymentMethod:       "MASTERCARD",
			ProcessingChannelId: "pc_1234",
			Breakdown:           []ActionBreakdown{},
		}

		queryResponse = QueryResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Limit:        5,
			Count:        1,
			Data:         []FinancialAction{financialAction},
		}

		pagingError = errors.ErrorDetails{
			RequestID:  "0HL80RJLS76I7",
			ErrorType:  "request_invalid",
			ErrorCodes: []string{"paging_limit_invalid"},
		}
	)

	cases := []struct {
		name             string
		query            QueryFilter
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*QueryResponse, error)
	}{
		{
			name: "when query is correct then return financial actions",
			query: QueryFilter{
				Limit:     10,
				PaymentId: "pay_1234",
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
						*respMapping = queryResponse
					})
			},
			checker: func(response *QueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, queryResponse.Limit, response.Limit)
				for _, action := range queryResponse.Data {
					assert.Equal(t, financialAction.PaymentId, action.PaymentId)
					assert.Equal(t, financialAction.ActionId, action.ActionId)
					assert.Equal(t, financialAction.EntityId, action.EntityId)
					assert.Equal(t, financialAction.SubEntityId, action.SubEntityId)
				}
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
			checker: func(response *QueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name: "when invalid paging then return error",
			query: QueryFilter{
				Limit: 255,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable Entity",
							Data:       &pagingError,
						})
			},
			checker: func(response *QueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "422 Unprocessable Entity", chkErr.Status)
				assert.Equal(t, &pagingError, chkErr.Data)
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

			tc.checker(client.GetFinancialActions(tc.query))
		})
	}
}
