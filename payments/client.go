package payments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "payments"

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
func (c *Client) Request(request *Request, params *checkout.Params) (*Response, error) {
	response, err := c.API.Post("/"+path, request, params)
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
	response, err := c.API.Get(fmt.Sprintf("/%v/%v", path, paymentID))
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
	response, err := c.API.Get(fmt.Sprintf("/%v/%v/actions", path, paymentID))
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
func (c *Client) Captures(paymentID string, request *CapturesRequest, params *checkout.Params) (*CapturesResponse, error) {
	response, err := c.API.Post(fmt.Sprintf("/%v/%v/captures", path, paymentID), request, params)
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
func (c *Client) Refunds(paymentID string, request *RefundsRequest, params *checkout.Params) (*RefundsResponse, error) {
	response, err := c.API.Post(fmt.Sprintf("/%v/%v/refunds", path, paymentID), request, params)
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
func (c *Client) Voids(paymentID string, request *VoidsRequest, params *checkout.Params) (*VoidsResponse, error) {
	response, err := c.API.Post(fmt.Sprintf("/%v/%v/voids", path, paymentID), request, params)
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
