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

func (c *Client) Get(customerID string) (*GetCustomerResponse, error) {
	resp, err := c.API.Get(fmt.Sprintf("/%v/%v", path, customerID))
	response := &GetCustomerResponse{
		StatusResponse: resp,
	}

	if err != nil {
		return nil, err
	}

	var customer Customer
	err = json.Unmarshal(resp.ResponseBody, &customer)

	if resp.StatusCode == http.StatusNoContent {
		return nil, err
	}

	response.Customer = &customer

	return response, err
}
