package webhooks

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

// Retrieve ...
func (c *Client) Retrieve() (*Response, error) {
	resp, err := c.API.Get("/webhooks")
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var webhookResponse []WebhookResponse
		err = json.Unmarshal(resp.ResponseBody, &webhookResponse)
		response.ConfiguredWebhooks = &webhookResponse
		return response, err
	}
	return response, err
}

// RegisterWebhook ...
func (c *Client) RegisterWebhook(request *Request) (*Response, error) {
	resp, err := c.API.Post("/webhooks", request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusCreated {
		var webhookResponse WebhookResponse
		err = json.Unmarshal(resp.ResponseBody, &webhookResponse)
		response.Webhook = &webhookResponse
		return response, err
	}
	return response, err
}

// RetrieveWebhook ...
func (c *Client) RetrieveWebhook(webhookID string) (*Response, error) {

	resp, err := c.API.Get(fmt.Sprintf("/webhooks/%v", webhookID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var webhookResponse WebhookResponse
		err = json.Unmarshal(resp.ResponseBody, &webhookResponse)
		response.Webhook = &webhookResponse
		return response, err
	}
	return response, err
}

// UpdateWebhook ...
func (c *Client) UpdateWebhook(webhookID string, request *Request) (*Response, error) {
	resp, err := c.API.Put(fmt.Sprintf("/webhooks/%v", webhookID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var webhookResponse WebhookResponse
		err = json.Unmarshal(resp.ResponseBody, &webhookResponse)
		response.Webhook = &webhookResponse
		return response, err
	}
	return response, err
}

// PartiallyUpdateWebhook ...
func (c *Client) PartiallyUpdateWebhook(webhookID string, request *Request) (*Response, error) {
	resp, err := c.API.Patch(fmt.Sprintf("/webhooks/%v", webhookID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var webhookResponse WebhookResponse
		err = json.Unmarshal(resp.ResponseBody, &webhookResponse)
		response.Webhook = &webhookResponse
		return response, err

	}
	return response, err
}

// RemoveWebhook ...
func (c *Client) RemoveWebhook(webhookID string) (*Response, error) {
	resp, err := c.API.Delete(fmt.Sprintf("/webhooks/%v", webhookID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var webhookResponse WebhookResponse
		err = json.Unmarshal(resp.ResponseBody, &webhookResponse)
		response.Webhook = &webhookResponse
		return response, err
	}
	return response, err
}
