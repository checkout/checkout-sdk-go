package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"

	issuingTesting "github.com/checkout/checkout-sdk-go/issuing/testing"
)

func TestSimulateAuthorization(t *testing.T) {
	t.Skip("Avoid creating cards all the time")

	cases := []struct {
		name    string
		request issuingTesting.CardAuthorizationRequest
		checker func(*issuingTesting.CardAuthorizationResponse, error)
	}{
		{
			name: "when simulate an authorization and this request is correct then should return a response",
			request: issuingTesting.CardAuthorizationRequest{
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
			},
			checker: func(response *issuingTesting.CardAuthorizationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
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

func TestSimulateIncrement(t *testing.T) {
	t.Skip("Avoid creating cards all the time")

	transactionId := cardSimulation(t, *virtualCardResponse)

	cases := []struct {
		name          string
		transactionId string
		request       issuingTesting.CardSimulationRequest
		checker       func(*issuingTesting.CardSimulationResponse, error)
	}{
		{
			name:          "when simulating an increment authorization with valid request then return response",
			transactionId: transactionId.Id,
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, issuingTesting.Authorized, response.Status)
			},
		},
		{
			name:          "when simulating an increment authorization with invalid transactionId then return error",
			transactionId: "not_found",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
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
			tc.checker(client.SimulateIncrement(tc.transactionId, tc.request))
		})
	}
}

func TestSimulateClearance(t *testing.T) {
	t.Skip("Avoid creating cards all the time")

	transactionId := cardSimulation(t, *virtualCardResponse)

	cases := []struct {
		name          string
		transactionId string
		request       issuingTesting.CardSimulationRequest
		checker       func(*common.MetadataResponse, error)
	}{
		{
			name:          "when simulating a clearance authorization with valid request then return response",
			transactionId: transactionId.Id,
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:          "when simulating a clearance authorization with invalid transactionId then return error",
			transactionId: "not_found",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			checker: func(response *common.MetadataResponse, err error) {
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
			tc.checker(client.SimulateClearing(tc.transactionId, tc.request))
		})
	}
}

func TestSimulateReversal(t *testing.T) {
	t.Skip("Avoid creating cards all the time")

	transactionId := cardSimulation(t, *virtualCardResponse)

	cases := []struct {
		name          string
		transactionId string
		request       issuingTesting.CardSimulationRequest
		checker       func(*issuingTesting.CardSimulationResponse, error)
	}{
		{
			name:          "when simulating a reversal authorization with valid request then return response",
			transactionId: transactionId.Id,
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, issuingTesting.Reversed, response.Status)
			},
		},
		{
			name:          "when simulating a reversal authorization with invalid transactionId then return error",
			transactionId: "not_found",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
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
			tc.checker(client.SimulateReversal(tc.transactionId, tc.request))
		})
	}
}
