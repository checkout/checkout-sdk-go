package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/financial"
)

func TestGetFinancialActions(t *testing.T) {
	t.Skip("Skipping tests because this suite is unstable")
	payment := makeCardPayment(t, true, 100)

	cases := []struct {
		name    string
		query   financial.QueryFilter
		checker func(interface{}, error)
	}{
		{
			name: "when query filters are valid then return financial actions",
			query: financial.QueryFilter{
				PaymentId: payment.Id,
				Limit:     5,
			},
			checker: func(response interface{}, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				queryResponse := response.(*financial.QueryResponse)
				assert.Equal(t, http.StatusOK, queryResponse.HttpMetadata.StatusCode)
				assert.True(t, queryResponse.Count > 0)
				for _, action := range queryResponse.Data {
					assert.NotNil(t, action)
					assert.NotNil(t, action.ActionId)
					assert.Equal(t, action.PaymentId, payment.Id)
				}
			},
		},
	}

	client := OAuthApi().Financial

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			process := func() (interface{}, error) { return client.GetFinancialActions(tc.query) }
			predicate := func(data interface{}) bool {
				response := data.(*financial.QueryResponse)
				return response.Count > 0
			}

			tc.checker(retriable(process, predicate, 5))
		})
	}
}
