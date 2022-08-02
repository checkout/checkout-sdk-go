package workflows

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/httpclient"
)

const path = "workflows"

type Client struct {
	API checkout.HTTPClient
}

func NewClient(config checkout.Config) *Client {
	return &Client{
		API: httpclient.NewClient(config),
	}
}

func (c *Client) RetrieveAll() (*Response, error) {
	resp, err := c.API.Get("/" + path)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var res []WorkflowResponse
		err = json.Unmarshal(resp.ResponseBody, &res)
		response.Workflows = res
		return response, err
	}
	return response, err
}

func (c *Client) Register(request *Request) (*Response, error) {
	resp, err := c.API.Post("/"+path, request, nil)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusCreated {
		var res WorkflowResponse
		err = json.Unmarshal(resp.ResponseBody, &res)
		response.Workflow = &res
		return response, err
	}
	return response, err
}

func (c *Client) RetrieveSingle(workflowID string) (*Response, error) {
	resp, err := c.API.Get(fmt.Sprintf("/%v/%v", path, workflowID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var res WorkflowResponse
		err = json.Unmarshal(resp.ResponseBody, &res)
		response.Workflow = &res
		return response, err
	}
	return response, err
}

func (c *Client) Update(workflowID string, request *Request) (*Response, error) {
	resp, err := c.API.Put(fmt.Sprintf("/%v/%v", path, workflowID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var res WorkflowResponse
		err = json.Unmarshal(resp.ResponseBody, &res)
		response.Workflow = &res
		return response, err
	}
	return response, err
}

func (c *Client) PartiallyUpdate(workflowID string, request *Request) (*Response, error) {
	resp, err := c.API.Patch(fmt.Sprintf("/%v/%v", path, workflowID), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var res WorkflowResponse
		err = json.Unmarshal(resp.ResponseBody, &res)
		response.Workflow = &res
		return response, err

	}
	return response, err
}

func (c *Client) Remove(workflowID string) (*Response, error) {
	resp, err := c.API.Delete(fmt.Sprintf("/%v/%v", path, workflowID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var res WorkflowResponse
		err = json.Unmarshal(resp.ResponseBody, &res)
		response.Workflow = &res
		return response, err
	}
	return response, err
}

func (c *Client) Test(workflowID string, request *TestRequest) (*Response, error) {
	resp, err := c.API.Post(fmt.Sprintf("/%v/%v", path, workflowID), request, nil)
	response := &Response{
		StatusResponse: resp,
	}
	return response, err
}
