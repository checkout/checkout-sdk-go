package files

import (
	"encoding/json"
	"fmt"
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
func (c *Client) UploadFile(file *FileUpload) (*Response, error) {
	if file == nil {
		return nil, fmt.Errorf("file cannot be nil, and params.Purpose and params.File must be set")
	}
	bodyBuffer, boundary, err := file.GetBody()
	if err != nil {
		return nil, err
	}
	resp, err := c.API.Upload(fmt.Sprintf("/%v", path), boundary, bodyBuffer)
	response := &Response{
		StatusResponse: resp,
	}
	if err != nil {
		return response, err
	}
	if resp.StatusCode == http.StatusCreated {
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
