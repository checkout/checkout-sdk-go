package payments

import (
	"encoding/json"
	"fmt"
	"net/http"

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
func (c *Client) Request(request *Request) (*Response, error) {
	response, err := c.API.Post("/payments", request)
	resp := &Response{
		StatusResponse: response,
	}
	if err != nil {
		return resp, err
	}
	if response.StatusCode == http.StatusCreated {
		var processed Processed
		err = json.Unmarshal(response.ResponseBody, &processed)
		resp.Processed = &processed
	} else if response.StatusCode == http.StatusAccepted {
		var pending PaymentPending
		err = json.Unmarshal(response.ResponseBody, &pending)
		resp.Pending = &pending
	}
	return resp, err
}

// Get ...
func (c *Client) Get(paymentID string) (*PaymentResponse, error) {
	response, err := c.API.Get(fmt.Sprintf("/payments/%v", paymentID))
	payment := &PaymentResponse{
		StatusResponse: response,
	}
	if err != nil {
		return payment, err
	}
	var paymentDetails Payment
	err = json.Unmarshal(response.ResponseBody, &paymentDetails)
	payment.Payment = &paymentDetails
	return payment, err
}

// Actions ...
func (c *Client) Actions(paymentID string) (*ActionsResponse, error) {
	response, err := c.API.Get(fmt.Sprintf("/payments/%v/actions", paymentID))
	act := &ActionsResponse{
		StatusResponse: response,
	}
	if err != nil {
		return act, err
	}
	actions := make([]*Action, 0)
	err = json.Unmarshal(response.ResponseBody, &actions)
	act.Actions = actions
	return act, err
}

// Captures ...
func (c *Client) Captures(paymentID string, request *CapturesRequest) (*CapturesResponse, error) {
	response, err := c.API.Post(fmt.Sprintf("/payments/%v/captures", paymentID), request)
	cap := &CapturesResponse{
		StatusResponse: response,
	}
	if err != nil {
		return cap, err
	}

	var accepted Accepted
	err = json.Unmarshal(response.ResponseBody, &accepted)
	cap.Accepted = &accepted
	return cap, err
}

// Refunds ...
func (c *Client) Refunds(paymentID string, request *RefundsRequest) (*RefundsResponse, error) {
	response, err := c.API.Post(fmt.Sprintf("/payments/%v/refunds", paymentID), request)
	ref := &RefundsResponse{
		StatusResponse: response,
	}
	if err != nil {
		return ref, err
	}

	var accepted Accepted
	err = json.Unmarshal(response.ResponseBody, &accepted)
	ref.Accepted = &accepted
	return ref, err
}

// Voids ...
func (c *Client) Voids(paymentID string, request *VoidsRequest) (*VoidsResponse, error) {
	response, err := c.API.Post(fmt.Sprintf("/payments/%v/voids", paymentID), request)
	void := &VoidsResponse{
		StatusResponse: response,
	}
	if err != nil {
		return void, err
	}

	var accepted Accepted
	err = json.Unmarshal(response.ResponseBody, &accepted)
	void.Accepted = &accepted
	return void, err
}
