package files

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

// UploadFile -
func (c *Client) UploadFile(request *Request) (*Response, error) {

	resp, err := c.API.Post(fmt.Sprintf("/files"), request)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var file File
		err = json.Unmarshal(resp.ResponseBody, &file)
		response.File = &file
		return response, err
	}
	return response, err
}

// GetFile -
func (c *Client) GetFile(fileID string) (*Response, error) {

	resp, err := c.API.Get(fmt.Sprintf("/files/%v", fileID))
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusOK {
		var file File
		err = json.Unmarshal(resp.ResponseBody, &file)
		response.File = &file
		return response, err
	}
	return response, err
}
