package sources

import (
	"encoding/json"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "sources"

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

// AddPaymentSource -
func (c *Client) AddPaymentSource(request *Request) (*Response, error) {
	response, err := c.API.Post("/"+path, request, nil)
	resp := &Response{
		StatusResponse: response,
	}
	if err != nil {
		return resp, err
	}
	if response.StatusCode == http.StatusCreated {
		var source Source
		err = json.Unmarshal(response.ResponseBody, &source)
		resp.Source = &source
	}
	return resp, err
}
