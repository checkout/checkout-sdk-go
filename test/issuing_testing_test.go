package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"

	issuingTesting "github.com/checkout/checkout-sdk-go/issuing/testing"
)

func TestSimulateAuthorization(t *testing.T) {
	t.Skip("Avoid creating cards all the time")
	request := issuingTesting.CardAuthorizationRequest{
		Card: issuingTesting.CardSimulation{
			Id:          virtualCardResponse.Id,
			ExpiryMonth: virtualCardResponse.ExpiryMonth,
			ExpiryYear:  virtualCardResponse.ExpiryYear,
		},
		Transaction: issuingTesting.TransactionSimulation{
			Type:     issuingTesting.Purchase,
			Amount:   100,
			Currency: common.GBP,
		},
	}

	cases := []struct {
		name    string
		request issuingTesting.CardAuthorizationRequest
		checker func(*issuingTesting.CardAuthorizationResponse, error)
	}{
		{
			name:    "when simulate an authorization and this request is correct then should return a response",
			request: request,
			checker: func(response *issuingTesting.CardAuthorizationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.Equal(t, issuingTesting.Declined, response.Status)
			},
		},
	}

	client := buildIssuingClientApi().Issuing

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.SimulateAuthorization(tc.request))
		})
	}
}
