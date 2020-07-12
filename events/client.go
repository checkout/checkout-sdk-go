package events

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

// RetrieveEventTypes -
func (c *Client) RetrieveEventTypes() (*Response, error) {
	resp, err := c.API.Get("/event-types?version=2.0")
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var eventTypes []EventType
		err = json.Unmarshal(resp.ResponseBody, &eventTypes)
		response.EventTypes = eventTypes
		return response, err
	}
	return response, err
}

// RetrieveEvents -
func (c *Client) RetrieveEvents(data map[string]string) (*Response, error) {
	var urlPath string = "/events?"
	var query string = ""
	var keyValue string
	for key, val := range data {
		keyValue = fmt.Sprintf("%s=%s&", key, val)
		query = query + keyValue
	}
	queryLength := len(query)
	if queryLength > 0 && query[queryLength-1] == '&' {
		query = query[:queryLength-1]
	}
	fmt.Printf("query: %s", query)
	fmt.Printf("urlPath+query: %s", fmt.Sprintf(urlPath+query))
	resp, err := c.API.Get(fmt.Sprintf(urlPath + query))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var events Events
		err = json.Unmarshal(resp.ResponseBody, &events)
		response.Events = &events
		return response, err
	}
	return response, err

}

// // RetrieveEventWithTime -
// func (c *Client) RetrieveEventWithTime(from time.Time, to time.Time, limit int64) (*Response, error) {
// 	resp, err := c.API.Get(fmt.Sprintf("/events?from=%v&to=%v&limit=%v", from, to, limit))
// 	response := &Response{
// 		StatusResponse: resp,
// 	}
// 	if err != nil {
// 		return response, err
// 	}
// 	if resp.StatusCode == http.StatusOK {
// 		var event Event
// 		err = json.Unmarshal(resp.ResponseBody, &event)
// 		response.Event = &event
// 		return response, err
// 	}
// 	return response, err
// }

// // RetrieveEventWithPayment -
// func (c *Client) RetrieveEventWithPayment(paymentID string) (*Response, error) {
// 	resp, err := c.API.Get(fmt.Sprintf("/events?payment_id=%v", paymentID))
// 	response := &Response{
// 		StatusResponse: resp,
// 	}
// 	if err != nil {
// 		return response, err
// 	}
// 	if resp.StatusCode == http.StatusOK {
// 		var event Event
// 		err = json.Unmarshal(resp.ResponseBody, &event)
// 		response.Event = &event
// 		return response, err
// 	}
// 	return response, err
// }

// RetrieveEvent -
func (c *Client) RetrieveEvent(eventID string) (*Response, error) {
	resp, err := c.API.Get(fmt.Sprintf("/events/%v", eventID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var event Event
		err = json.Unmarshal(resp.ResponseBody, &event)
		response.Event = &event
		return response, err
	}
	return response, err
}

// RetrieveEventNotification -
func (c *Client) RetrieveEventNotification(eventID string, notificationID string) (*Response, error) {
	resp, err := c.API.Get(fmt.Sprintf("/events/%v/notifications/%v", eventID, notificationID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var notification Notification
		err = json.Unmarshal(resp.ResponseBody, &notification)
		response.Notification = &notification
		return response, err
	}
	return response, err
}

// Retry -
func (c *Client) Retry(eventID string, webhookID string) (*Response, error) {
	resp, err := c.API.Post(fmt.Sprintf("/events/%v/webhooks/%v/retry", eventID, webhookID), nil)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		return response, err
	}
	return response, err
}

// RetryAll -
func (c *Client) RetryAll(eventID string) (*Response, error) {
	resp, err := c.API.Post(fmt.Sprintf("/events/%v/webhooks/retry", eventID), nil)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		return response, err
	}
	return response, err
}
