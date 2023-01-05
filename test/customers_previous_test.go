package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/customers"
	"github.com/checkout/checkout-sdk-go/errors"
)

func TestCreateCustomerPrevious(t *testing.T) {
	cases := []struct {
		name    string
		request customers.CustomerRequest
		checker func(*common.IdResponse, error)
	}{
		{
			name: "when correct data then create new customer",
			request: customers.CustomerRequest{
				Email: GenerateRandomEmail(),
				Name:  FirstName + LastName,
				Phone: Phone(),
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Id)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
			},
		},
		{
			name: "when invalid email then return error",
			request: customers.CustomerRequest{
				Email: "bad_email",
				Name:  FirstName + LastName,
				Phone: Phone(),
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
			},
		},
	}

	client := PreviousApi().Customers

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Create(tc.request))
		})
	}
}

func TestGetCustomerPrevious(t *testing.T) {
	custId := createCustomerPrevious()

	cases := []struct {
		name       string
		customerId string
		checker    func(*customers.GetCustomerResponse, error)
	}{
		{
			name:       "when customer exists then return customer info",
			customerId: custId,
			checker: func(response *customers.GetCustomerResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, custId, response.Id)
			},
		},
		{
			name:       "when customer doesn't exists then return error",
			customerId: InvalidCustomerId,
			checker: func(response *customers.GetCustomerResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := PreviousApi().Customers

	for _, tc := range cases {
		tc.checker(client.Get(tc.customerId))
	}
}

func TestUpdateCustomerPrevious(t *testing.T) {
	custId := createCustomerPrevious()

	cases := []struct {
		name       string
		customerId string
		request    customers.CustomerRequest
		checker    func(*common.MetadataResponse, error)
	}{
		{
			name:       "when customer exists then return 204 Customer updated successfully",
			customerId: custId,
			request: customers.CustomerRequest{
				Name: "New Name",
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode) // StatusNoContent == 204
			},
		},
		{
			name:       "when customer doesn't exists then return error",
			customerId: InvalidCustomerId,
			request: customers.CustomerRequest{
				Name: "New Name",
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := PreviousApi().Customers

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Update(tc.customerId, tc.request))
		})
	}
}

func TestDeleteCustomerPrevious(t *testing.T) {
	custId := createCustomerPrevious()

	cases := []struct {
		name       string
		customerId string
		checker    func(*common.MetadataResponse, error)
	}{
		{
			name:       "when customer exists then delete customer and return 204 Customer Deleted Successfully",
			customerId: custId,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:       "when customer doesn't exists then return error",
			customerId: InvalidCustomerId,
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := PreviousApi().Customers

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.Delete(tc.customerId))
		})
	}
}

func createCustomerPrevious() string {
	request := customers.CustomerRequest{
		Email: GenerateRandomEmail(),
		Name:  Name,
		Phone: Phone(),
	}
	response, _ := PreviousApi().Customers.Create(request)

	return response.Id
}
