package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/abc"
)

func TestRequestPayoutPrevious(t *testing.T) {

	cardDestination := abc.NewRequestCardDestination()
	cardDestination.Name = Name
	cardDestination.FirstName = FirstName
	cardDestination.LastName = LastName
	cardDestination.Number = CardNumber
	cardDestination.ExpiryYear = ExpiryYear
	cardDestination.ExpiryMonth = ExpiryMonth
	cardDestination.BillingAddress = Address()
	cardDestination.Phone = Phone()

	payoutRequest := abc.PayoutRequest{
		Destination:      cardDestination,
		FundTransferType: abc.AA,
		Currency:         common.GBP,
		Reference:        Reference,
		Description:      Description,
		Capture:          false,
		Amount:           5,
	}

	response, err := PreviousApi().Payments.RequestPayout(payoutRequest, nil)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.Reference)
	assert.Equal(t, payments.Pending, response.Status)
	assert.Nil(t, response.ThreeDs)

	//Customer
	assert.NotEmpty(t, response.Customer)
	customer := response.Customer
	assert.NotEmpty(t, customer)
	assert.NotEmpty(t, customer.Id)

	//Links
	assert.NotEmpty(t, response.Links["self"])
}
