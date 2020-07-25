package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/httpclient"
)

const path = "files"

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
func (c *Client) UploadFile(values map[string]io.Reader) (*Response, error) {

	resp, err := c.API.Upload(fmt.Sprintf("/%v", path), values)
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

	resp, err := c.API.Get(fmt.Sprintf("/%v/%v", path, fileID))
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
