package instruments

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
	"github.com/checkout/checkout-sdk-go/payments"
)

const path = "instruments"

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

// Create an instrument
func (c *Client) Create(request *Request) (*Response, error) {
	resp, err := c.API.Post("/"+path, request, nil)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusCreated {
		var source payments.SourceResponse
		err = json.Unmarshal(resp.ResponseBody, &source)
		response.Source = &source
		return response, err
	}
	return response, err
}

// Get instrument details
func (c *Client) Get(sourceID string) (*Response, error) {

	resp, err := c.API.Get(fmt.Sprintf("/%v/%v", path, sourceID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var source payments.SourceResponse
		err = json.Unmarshal(resp.ResponseBody, &source)
		response.Source = &source
		return response, err
	}
	return response, err
}

// Update instrument details
func (c *Client) Update(sourceID string, request *Request) (*Response, error) {
	resp, err := c.API.Patch(fmt.Sprintf("/%v/%v", path, sourceID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var instrumentResponse InstrumentResponse
		err = json.Unmarshal(resp.ResponseBody, &instrumentResponse)
		response.InstrumentResponse = &instrumentResponse
		return response, err
	}
	return response, err
}
