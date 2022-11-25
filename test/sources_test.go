package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go-beta/sources"
)

func TestShouldCreateSepaSource(t *testing.T) {
	request := sources.NewSepaSourceRequest()
	request.BillingAddress = Address()
	request.Reference = "reference"
	request.Phone = Phone()
	request.SourceData = &sources.SourceData{
		FirstName:         FirstName,
		LastName:          LastName,
		AccountIban:       Iban,
		Bic:               Bic,
		BillingDescriptor: "SDK Test",
		MandateType:       sources.Single,
	}

	sepaSourceResponse, err := PreviousApi().Sources.CreateSepaSource(request)
	assert.Nil(t, err)
	assert.NotNil(t, sepaSourceResponse)
}
