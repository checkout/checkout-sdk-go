package httpclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/common"
)

var client *HTTPClient

const (
	defaultHTTPTimeout = 30 * time.Second
)

// HTTPClient ...
type HTTPClient struct {
	SecretKey  string
	URI        string
	HTTPClient *http.Client
}

// GetClient ...
func GetClient() *HTTPClient {
	return client
}

// NewClient ...
func NewClient(config checkout.Config) *HTTPClient {
	client = &HTTPClient{
		SecretKey:  config.SecretKey,
		URI:        config.URI,
		HTTPClient: &http.Client{Timeout: defaultHTTPTimeout},
	}
	return client
}

// Get ...
func (c *HTTPClient) Get(param string) (*checkout.StatusResponse, error) {

	request, err := http.NewRequest("GET", c.URI+param, nil)
	if err != nil {
		return nil, err
	}
	c.setHeader(request)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	apiResponse.ResponseBody = responseBody
	if response.StatusCode >= 400 {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}

// Post ...
func (c *HTTPClient) Post(param string, body interface{}) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest("POST", c.URI+param, body)
	if err != nil {
		return nil, err
	}
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return apiResponse, err
	}
	apiResponse.ResponseBody = responseBody
	if response.StatusCode >= 400 {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}

// NewRequest ...
func (c *HTTPClient) NewRequest(method, path string, body interface{}) (*http.Request, error) {

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", path, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	c.setHeader(request)
	return request, nil
}

func (c *HTTPClient) setHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "checkout-sdk-go/"+checkout.ClientVersion)
	req.Header.Add("Authorization", c.SecretKey)
}

func responseToError(apiRes *checkout.StatusResponse, body []byte) *common.Error {
	err := &common.Error{}
	if apiRes.StatusCode == 422 {
		var details common.ErrorDetails
		json.Unmarshal(body, &details)
		err.Data = &details
	}
	err.Status = apiRes.Status
	err.StatusCode = apiRes.StatusCode
	return err
}
