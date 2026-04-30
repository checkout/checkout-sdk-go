package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/errors"

	digitalcards "github.com/checkout/checkout-sdk-go/v2/issuing/digitalcards"
)

// # tests

func TestGetDigitalCard(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	cases := []struct {
		name          string
		digitalCardId string
		checker       func(*digitalcards.GetDigitalCardResponse, error)
	}{
		{
			name:          "when request is correct then should return 200",
			digitalCardId: "dcr_5ngxzsynm2me3oxf73esbhda6q",
			checker: func(response *digitalcards.GetDigitalCardResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.CardId)
				assert.NotNil(t, response.Status)
			},
		},
		{
			name:          "when digital card id not found then return error",
			digitalCardId: "dcr_not_found",
			checker: func(response *digitalcards.GetDigitalCardResponse, err error) {
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
			tc.checker(client.GetDigitalCard(tc.digitalCardId))
		})
	}
}
