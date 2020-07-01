package payments

import (
	"encoding/json"
	"fmt"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/httpclient"
)

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

// Request ...
func (c *Client) Request(request *checkout.PaymentRequest) (*checkout.Response, error) {
	response, err := c.API.Post("/payments", request)
	req := &checkout.Response{
		APIResponse: response,
	}
	if err != nil {
		return req, err
	}
	if response.StatusCode == 201 {
		var created checkout.Created
		err = json.Unmarshal(response.ResponseBody, &created)
		req.Created = &created
	} else if response.StatusCode == 202 {
		var pending checkout.Pending
		err = json.Unmarshal(response.ResponseBody, &pending)
		req.Pending = &pending
	}
	return req, err
}

// Get ...
func (c *Client) Get(paymentID string) (*checkout.PaymentResponse, error) {
	response, err := c.API.Get(fmt.Sprintf("/payments/%v", paymentID))
	payment := &checkout.PaymentResponse{
		APIResponse: response,
	}
	if err != nil {
		return payment, err
	}
	var paymentDetails checkout.Payment
	err = json.Unmarshal(response.ResponseBody, &paymentDetails)
	payment.Payment = &paymentDetails
	return payment, err
}
