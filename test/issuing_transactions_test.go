package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"

	transactions "github.com/checkout/checkout-sdk-go/v2/issuing/transactions"
)

// # tests

func TestGetListTransactions(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name    string
		query   transactions.TransactionsQuery
		checker func(*transactions.TransactionsListResponse, error)
	}{
		{
			name:  "when query is correct then should return 200",
			query: listTransactionsQuery(),
			checker: func(response *transactions.TransactionsListResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.TotalCount)
				assert.NotNil(t, response.Data)
			},
		},
		{
			name:  "when card id not found then return empty result",
			query: transactions.TransactionsQuery{CardId: "crd_not_found"},
			checker: func(response *transactions.TransactionsListResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 0, *response.TotalCount)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetListTransactions(tc.query))
		})
	}
}

func TestGetSingleTransaction(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name          string
		transactionId string
		checker       func(*transactions.TransactionResponse, error)
	}{
		{
			name:          "when request is correct then should return 200",
			transactionId: transactionResponse.Id,
			checker: func(response *transactions.TransactionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, transactionResponse.Id, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.TransactionType)
			},
		},
		{
			name:          "when transaction id not found then return error",
			transactionId: "trx_not_found",
			checker: func(response *transactions.TransactionResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetSingleTransaction(tc.transactionId))
		})
	}
}

// # common methods

func listTransactionsQuery() transactions.TransactionsQuery {
	return transactions.TransactionsQuery{
		Limit:  10,
		CardId: virtualCardId,
	}
}

func assertTransactionResponse(t *testing.T, response *transactions.TransactionResponse) {
	assert.NotNil(t, response)
	assert.NotNil(t, response.Id)
	assert.NotNil(t, response.Status)
	assert.NotNil(t, response.TransactionType)
}
