package customers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "customers"

// Client ...
type Client struct {
	API checkout.HTTPClient
}

// NewClient ...
func NewClient(config checkout.Config) *Client {
	return &Client{
		API: httpclient.NewClient(config),
	}
}

// Create a customer
func (c *Client) Create(request *Request) (*Response, error) {
	resp, err := c.API.Post("/"+path, request, nil)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusCreated {
		var instrumentResponse CustomerResponse
		err = json.Unmarshal(resp.ResponseBody, &instrumentResponse)
		response.Customer = &instrumentResponse
		return response, err
	}
	return response, err
}

// Update customer details
func (c *Client) Update(customerID string, request *Request) (*Response, error) {
	resp, err := c.API.Patch(fmt.Sprintf("/%v/%v", path, customerID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusNoContent {
		return response, err
	}
	return response, err
}

// Get customer details
func (c *Client) Get(customerID string) (*Response, error) {

	resp, err := c.API.Get(fmt.Sprintf("/%v/%v", path, customerID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var customerResponse CustomerResponse
		err = json.Unmarshal(resp.ResponseBody, &customerResponse)
		response.Customer = &customerResponse
		return response, err
	}
	return response, err
}
