package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/balances"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
)

func TestRetrieveEntityBalances(t *testing.T) {
	cases := []struct {
		name     string
		entityId string
		query    balances.QueryFilter
		checker  func(*balances.QueryResponse, error)
	}{
		{
			name:     "when request is correct then return a balance details",
			entityId: "ent_kidtcgc3ge5unf4a5i6enhnr5m",
			query:    balances.QueryFilter{Query: fmt.Sprintf("currency:%s", common.GBP)},
			checker: func(response *balances.QueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
		{
			name:     "when entityId is inexistant then return error",
			entityId: "ent_w4jelhppmfiufdnatam37wrfc4",
			query:    balances.QueryFilter{Query: fmt.Sprintf("currency:%s", common.GBP)},
			checker: func(response *balances.QueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := OAuthApi().Balances

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrieveEntityBalances(tc.entityId, tc.query))
		})
	}
}
