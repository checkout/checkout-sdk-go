package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/apm/klarna"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
)

func TestCreateCreditSession(t *testing.T) {
	t.Skip("unavailable")
	cases := []struct {
		name    string
		request klarna.CreditSessionRequest
		checker func(*klarna.CreditSessionResponse, error)
	}{
		{
			name: "when request is correct then create klarna session",
			request: klarna.CreditSessionRequest{
				PurchaseCountry: common.GB,
				Currency:        common.GBP,
				Locale:          "en-GB",
				Amount:          1000,
				TaxAmount:       1,
				Products:        getKlarnaProduct(),
			},
			checker: func(response *klarna.CreditSessionResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.SessionId)
				assert.NotNil(t, response.ClientToken)
				assert.NotNil(t, response.PaymentMethodCategories)
			},
		},
		{
			name:    "when request is missing information then return error",
			request: klarna.CreditSessionRequest{},
			checker: func(response *klarna.CreditSessionResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
			},
		},
	}

	client := PreviousApi().Klarna

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateCreditSession(tc.request))
		})
	}
}

func TestGetCreditSession(t *testing.T) {
	t.Skip("unavailable")
	var (
		sessionId = createCreditSession(t).SessionId
	)

	cases := []struct {
		name      string
		sessionId string
		checker   func(*klarna.CreditSession, error)
	}{
		{
			name:      "when session exists then return session",
			sessionId: sessionId,
			checker: func(response *klarna.CreditSession, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.ClientToken)
				assert.NotNil(t, response.PurchaseCountry)
				assert.NotNil(t, response.Currency)
				assert.NotNil(t, response.Amount)
				assert.NotNil(t, response.TaxAmount)
			},
		},
		{
			name:      "when session not found then return error",
			sessionId: "invalid_session_id",
			checker: func(response *klarna.CreditSession, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := PreviousApi().Klarna

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetCreditSession(tc.sessionId))
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

func createCreditSession(t *testing.T) *klarna.CreditSessionResponse {
	t.Skip("unavailable")
	request := klarna.CreditSessionRequest{
		PurchaseCountry: common.GB,
		Currency:        common.GBP,
		Locale:          "en-GB",
		Amount:          1000,
		TaxAmount:       1,
		Products:        getKlarnaProduct(),
	}

	response, err := PreviousApi().Klarna.CreateCreditSession(request)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating klarna session - %s", err.Error()))
	}

	return response
}
