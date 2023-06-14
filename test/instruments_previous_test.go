package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/instruments/abc"
	"github.com/checkout/checkout-sdk-go/tokens"
)

var (
	instrumentTokenPrevious *abc.CreateInstrumentResponse
)

func TestSetupInstrumentsPrevious(t *testing.T) {
	cardTokenResponse := RequestCardTokenPrevious(t)
	instrumentTokenPrevious = createTokenInstrumentPrevious(t, cardTokenResponse)
}

func TestGetInstrumentPrevious(t *testing.T) {
	cases := []struct {
		name       string
		responseId string
		checker    func(*abc.GetInstrumentResponse, error)
	}{
		{
			name:       "when the request is valid then response is not nil",
			responseId: instrumentTokenPrevious.Id,
			checker: func(response *abc.GetInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, common.Card, response.Type)
				assert.NotEmpty(t, response.Id)
				assert.NotEmpty(t, response.Fingerprint)
				assert.NotEmpty(t, response.ExpiryMonth)
				assert.NotEmpty(t, response.ExpiryYear)
				assert.NotEmpty(t, response.Scheme)
				assert.NotEmpty(t, response.Last4)
				assert.NotEmpty(t, response.Bin)
				assert.NotEmpty(t, response.CardType)
				assert.NotEmpty(t, response.CardCategory)
				assert.NotEmpty(t, response.ProductId)
				assert.NotEmpty(t, response.ProductType)
				assert.NotEmpty(t, response.Customer)
				assert.NotEmpty(t, response.Customer.Id)
				assert.NotEmpty(t, response.Customer.Name)
				assert.NotEmpty(t, response.Customer.Email)
			},
		},
	}

	client := PreviousApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Get(tc.responseId))
		})
	}
}

func TestUpdateInstrumentPrevious(t *testing.T) {
	updateRequest := abc.UpdateInstrumentRequest{
		ExpiryMonth: 12,
		ExpiryYear:  2026,
		Name:        "New Name",
	}

	cases := []struct {
		name          string
		responseId    string
		updateRequest abc.UpdateInstrumentRequest
		checkerUpdate func(*abc.UpdateInstrumentResponse, error)
		checkerGet    func(*abc.GetInstrumentResponse, error)
	}{
		{
			name:          "when update instrument request then this request is updated",
			responseId:    instrumentTokenPrevious.Id,
			updateRequest: updateRequest,
			checkerUpdate: func(response *abc.UpdateInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, common.Card, response.Type)
				assert.NotEmpty(t, response.Fingerprint)
			},
			checkerGet: func(response *abc.GetInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, common.Card, response.Type)
				assert.Equal(t, 12, response.ExpiryMonth)
				assert.Equal(t, 2026, response.ExpiryYear)
				assert.Equal(t, "New Name", response.Name)
			},
		},
	}

	client := PreviousApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkerUpdate(client.Update(tc.responseId, tc.updateRequest))
			tc.checkerGet(client.Get(tc.responseId))
		})
	}
}

func TestDeleteInstrumentPrevious(t *testing.T) {
	cases := []struct {
		name          string
		responseId    string
		checkerDelete func(*common.MetadataResponse, error)
		checkerGet    func(*abc.GetInstrumentResponse, error)
	}{
		{
			name:       "when delete a instrument request then return 204",
			responseId: instrumentTokenPrevious.Id,
			checkerDelete: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 204, response.HttpMetadata.StatusCode)
			},
			checkerGet: func(response *abc.GetInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := PreviousApi().Instruments

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkerDelete(client.Delete(tc.responseId))
			tc.checkerGet(client.Get(tc.responseId))
		})
	}
}

func createTokenInstrumentPrevious(t *testing.T, token *tokens.CardTokenResponse) *abc.CreateInstrumentResponse {
	request := abc.CreateInstrumentRequest{
		Type:  common.Token,
		Token: token.Token,
		Customer: &abc.InstrumentCustomerRequest{
			Email:     Email,
			Name:      Name,
			Phone:     Phone(),
			IsDefault: true,
		},
	}
	response, err := PreviousApi().Instruments.Create(request)

	testCreateInstrument(t, err, response)

	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating token instrument - %s", err.Error()))
	}
	return response
}

func testCreateInstrument(t *testing.T, err error, response *abc.CreateInstrumentResponse) {
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, common.Card, response.Type)
	assert.NotEmpty(t, response.Id)
	assert.NotEmpty(t, response.Fingerprint)
	assert.NotEmpty(t, response.ExpiryMonth)
	assert.NotEmpty(t, response.ExpiryYear)
	assert.NotEmpty(t, response.Scheme)
	assert.NotEmpty(t, response.Last4)
	assert.NotEmpty(t, response.Bin)
	assert.NotEmpty(t, response.CardType)
	assert.NotEmpty(t, response.CardCategory)
	assert.NotEmpty(t, response.ProductId)
	assert.NotEmpty(t, response.ProductType)
	assert.NotEmpty(t, response.Customer)
	assert.NotEmpty(t, response.Customer.Id)
	assert.NotEmpty(t, response.Customer.Name)
	assert.NotEmpty(t, response.Customer.Email)
}
