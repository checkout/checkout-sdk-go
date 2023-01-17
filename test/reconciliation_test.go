package test

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/abc"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/reconciliation"
)

var (
	layout        = "2006-01-02T15:04:05Z"
	now           = time.Now()
	nowMinusMonth = now.AddDate(0, -1, 0)
	from, _       = time.Parse(layout, nowMinusMonth.Format(layout))
	to, _         = time.Parse(layout, now.Format(layout))
)

func TestQueryPaymentsReport(t *testing.T) {
	t.Skip("only works in production")

	cases := []struct {
		name    string
		query   reconciliation.PaymentReportsQuery
		checker func(*reconciliation.PaymentReportsResponse, error)
	}{
		{
			name: "when query is correct then return payment reports",
			query: reconciliation.PaymentReportsQuery{
				From:  from,
				To:    to,
				Limit: 10,
			},
			checker: func(response *reconciliation.PaymentReportsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Data)
			},
		},
	}

	client := getProdApi().Reconciliation

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.QueryPaymentsReport(tc.query))
		})
	}
}

func TestGetSinglePaymentReport(t *testing.T) {
	t.Skip("only works in production")

	cases := []struct {
		name      string
		paymentId string
		checker   func(*reconciliation.PaymentReportsResponse, error)
	}{
		{
			name:      "when payment exists then return payment reports",
			paymentId: "pay_oe5vaxisis4krciobenmrv4xze",
			checker: func(response *reconciliation.PaymentReportsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Data)
			},
		},
		{
			name:      "when payment doesn't exist then return error",
			paymentId: "not_found",
			checker: func(response *reconciliation.PaymentReportsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := getProdApi().Reconciliation

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetSinglePaymentReport(tc.paymentId))
		})
	}
}

func TestQueryStatementsReport(t *testing.T) {
	t.Skip("only works in production")

	cases := []struct {
		name    string
		query   common.DateRangeQuery
		checker func(*reconciliation.StatementReportsResponse, error)
	}{
		{
			name: "when query is correct then return statements reports",
			query: common.DateRangeQuery{
				From: from,
				To:   to,
			},
			checker: func(response *reconciliation.StatementReportsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Data)
			},
		},
	}

	client := getProdApi().Reconciliation

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.QueryStatementsReport(tc.query))
		})
	}
}

func TestRetrieveCVSPaymentsReport(t *testing.T) {
	t.Skip("only works in production")

	cases := []struct {
		name    string
		query   common.DateRangeQuery
		checker func(*common.ContentResponse, error)
	}{
		{
			name: "when query is correct then return payment reports",
			query: common.DateRangeQuery{
				From: from,
				To:   to,
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Content)
			},
		},
	}

	client := getProdApi().Reconciliation

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrieveCVSPaymentsReport(tc.query))
		})
	}
}

func TestRetrieveCVSSingleStatementReport(t *testing.T) {
	t.Skip("only works in production")

	cases := []struct {
		name        string
		statementId string
		checker     func(*common.ContentResponse, error)
	}{
		{
			name:        "when statement exists then return statement reports",
			statementId: "221222B100981",
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Content)
			},
		},
	}

	client := getProdApi().Reconciliation

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrieveCVSSingleStatementReport(tc.statementId))
		})
	}
}

func TestRetrieveCVSStatementsReport(t *testing.T) {
	t.Skip("only works in production")

	cases := []struct {
		name    string
		query   common.DateRangeQuery
		checker func(*common.ContentResponse, error)
	}{
		{
			name: "when query is correct then return statement reports",
			query: common.DateRangeQuery{
				From: from,
				To:   to,
			},
			checker: func(response *common.ContentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Content)
			},
		},
	}

	client := getProdApi().Reconciliation

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrieveCVSStatementsReport(tc.query))
		})
	}
}

func getProdApi() *abc.Api {
	api, _ := checkout.Builder().Previous().
		WithEnvironment(configuration.Production()).
		WithSecretKey(os.Getenv("CHECKOUT_PREVIOUS_SECRET_KEY_PROD")).
		Build()

	return api
}
