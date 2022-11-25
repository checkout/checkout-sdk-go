package test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPaymentActions(t *testing.T) {

	paymentResponse := makeCardPayment(t, false, 10)

	Wait(time.Duration(3))

	response, err := DefaultApi().Payments.GetPaymentActions(paymentResponse.Id)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Actions)

	for _, action := range response.Actions {
		assert.NotEmpty(t, action.Amount)
		assert.True(t, action.Approved)
		assert.Nil(t, action.Links)
		assert.NotEmpty(t, action.ProcessedOn)
		assert.NotEmpty(t, action.Reference)
		assert.NotEmpty(t, action.ResponseCode)
		assert.NotEmpty(t, action.ResponseSummary)
		assert.NotEmpty(t, action.Type)
	}

}
