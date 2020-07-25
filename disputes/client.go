package disputes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
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

// GetDisputes ...
func (c *Client) GetDisputes(request *Request) (*Response, error) {
	value, _ := query.Values(request.QueryParameter)
	var query string = value.Encode()
	var urlPath string = "/disputes" + "?"
	resp, err := c.API.Get(urlPath + query)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var disputes Disputes
		err = json.Unmarshal(resp.ResponseBody, &disputes)
		response.Disputes = &disputes
		return response, err
	}
	return response, err
}

// GetDispute ...
func (c *Client) GetDispute(disputeID string) (*Response, error) {
	resp, err := c.API.Get(fmt.Sprintf("/disputes/%v", disputeID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var dispute Dispute
		err = json.Unmarshal(resp.ResponseBody, &dispute)
		response.Dispute = &dispute
		return response, err
	}
	return response, err
}

// AcceptDispute -
func (c *Client) AcceptDispute(disputeID string) (*Response, error) {
	resp, err := c.API.Post(fmt.Sprintf("/disputes/%v/accept", disputeID), nil)
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

// ProvideDisputeEvidence ...
func (c *Client) ProvideDisputeEvidence(disputeID string, request *Request) (*Response, error) {
	resp, err := c.API.Put(fmt.Sprintf("/disputes/%v", disputeID), request)
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

// GetDisputeEvidence ...
func (c *Client) GetDisputeEvidence(disputeID string) (*Response, error) {
	resp, err := c.API.Get(fmt.Sprintf("/disputes/%v/evidence", disputeID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var evidence map[string]string
		err = json.Unmarshal(resp.ResponseBody, &evidence)
		response.Evidence = &evidence
		return response, err
	}
	return response, err
}

// SubmitDisputeEvidence -
func (c *Client) SubmitDisputeEvidence(disputeID string) (*Response, error) {
	resp, err := c.API.Post(fmt.Sprintf("/disputes/%v/evidence", disputeID), nil)
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
