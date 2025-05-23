package test

import (
	"fmt"
	"github.com/checkout/checkout-sdk-go/common"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/transfers"
)

func TestInitiateTransferOfFounds(t *testing.T) {
	var idempotencyKey = uuid.New().String()

	cases := []struct {
		name           string
		request        transfers.TransferRequest
		idempotencyKey *string
		checker        func(*transfers.TransferResponse, error)
	}{
		{
			name: "when request is correct then create transfer of founds with idempotency key",
			request: transfers.TransferRequest{
				Reference:    "reference",
				TransferType: transfers.Commission,
				Source: &transfers.TransferSourceRequest{
					Id:       "ent_kidtcgc3ge5unf4a5i6enhnr5m",
					Amount:   100,
					Currency: common.GBP,
				},
				Destination: &transfers.TransferDestinationRequest{Id: "ent_w4jelhppmfiufdnatam37wrfc4"},
			},
			idempotencyKey: &idempotencyKey,
			checker: func(response *transfers.TransferResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
			},
		},
		{
			name:           "when request invalid then return error",
			request:        transfers.TransferRequest{},
			idempotencyKey: &idempotencyKey,
			checker: func(response *transfers.TransferResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "validation_error", chkErr.Data.ErrorType)
			},
		},
		{
			name: "when request without idempotency key then return error",
			request: transfers.TransferRequest{
				Reference:    "reference",
				TransferType: transfers.Commission,
				Source: &transfers.TransferSourceRequest{
					Id:       "ent_kidtcgc3ge5unf4a5i6enhnr5m",
					Amount:   100,
					Currency: common.GBP,
				},
				Destination: &transfers.TransferDestinationRequest{Id: "ent_w4jelhppmfiufdnatam37wrfc4"},
			},
			checker: func(response *transfers.TransferResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "validation_error", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "idempotency_key_required")
			},
		},
	}

	client := OAuthApi().Transfers

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.InitiateTransferOfFounds(tc.request, tc.idempotencyKey))
		})
	}
}

func TestRetrieveTransfer(t *testing.T) {
	transfer := createTransferOfFounds(t)

	cases := []struct {
		name       string
		transferId string
		checker    func(*transfers.TransferDetails, error)
	}{
		{
			name:       "when transfer exists then return transfer details",
			transferId: transfer.Id,
			checker: func(response *transfers.TransferDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, transfer.Id, response.Id)
				assert.NotNil(t, response.Reference)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.TransferType)
				assert.NotNil(t, response.Source)
				assert.Equal(t, int64(100), response.Source.Amount)
				assert.Equal(t, common.GBP, response.Source.Currency)
				assert.NotNil(t, response.Destination)
			},
		},
	}

	client := OAuthApi().Transfers

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrieveTransfer(tc.transferId))
		})
	}
}

func createTransferOfFounds(t *testing.T) *transfers.TransferResponse {
	req := transfers.TransferRequest{
		Reference:    "reference",
		TransferType: transfers.Commission,
		Source: &transfers.TransferSourceRequest{
			Id:       "ent_kidtcgc3ge5unf4a5i6enhnr5m",
			Amount:   100,
			Currency: common.GBP,
		},
		Destination: &transfers.TransferDestinationRequest{Id: "ent_w4jelhppmfiufdnatam37wrfc4"},
	}

	idempotencyKey := uuid.New().String()

	response, err := OAuthApi().Transfers.InitiateTransferOfFounds(req, &idempotencyKey)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating transfer of founds - %s", err.Error()))
	}

	return response
}
