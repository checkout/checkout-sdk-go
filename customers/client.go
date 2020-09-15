package customers

import (
	"fmt"
	"net/http"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/httpclient"
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
