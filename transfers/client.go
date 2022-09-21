package transfers

import (
    "encoding/json"
    "fmt"

    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "transfers"

// Client ...
type Client struct {
    API checkout.HTTPClient
}

// NewClient ...
func NewClient(config checkout.Config) *Client {
    return &Client{
        API: httpclient.NewTransferClient(config),
    }
}

// Initiate ...
func (c *Client) Initiate(request *InitiateRequest, params *checkout.Params) (*InitiateResponse, error) {
    response, err := c.API.Post("/"+path, request, params)
    if err != nil {
        return nil, err
    }

    resp := &InitiateResponse{}
    err = json.Unmarshal(response.ResponseBody, &resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

// Get ...
func (c *Client) Get(id string) (*Transfer, error) {
    response, err := c.API.Get(fmt.Sprintf("/%v/%v", path, id))
    if err != nil {
        return nil, err
    }
    transfer := &Transfer{}
    err = json.Unmarshal(response.ResponseBody, &transfer)
    if err != nil {
        return nil, err
    }

    return transfer, nil
}
