package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
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
	PublicKey         string
	SecretKey         string
	URI               string
	IdempotencyKey    string
	CancellationToken string
	HTTPClient        *http.Client
}

// GetClient ...
func GetClient() *HTTPClient {
	return client
}

// NewClient ...
func NewClient(config checkout.Config) *HTTPClient {
	client = &HTTPClient{
		PublicKey:      config.PublicKey,
		SecretKey:      config.SecretKey,
		URI:            config.URI,
		IdempotencyKey: config.IdempotencyKey,
		HTTPClient:     &http.Client{Timeout: defaultHTTPTimeout},
	}
	return client
}

// Get ...
func (c *HTTPClient) Get(param string) (*checkout.StatusResponse, error) {

	request, err := http.NewRequest(http.MethodGet, c.URI+param, nil)
	if err != nil {
		return nil, err
	}
	c.setHeader(request)
	c.setCredential(c.URI+param, request)
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
	if response.StatusCode >= http.StatusBadRequest {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}

// Post ...
func (c *HTTPClient) Post(param string, body interface{}) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodPost, c.URI+param, body)
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
	if response.StatusCode >= http.StatusBadRequest {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}

// NewRequest ...
func (c *HTTPClient) NewRequest(method, path string, body interface{}) (*http.Request, error) {

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, path, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	c.setHeader(request)
	c.setCredential(path, request)
	c.setIdempotencyKey(request)
	return request, nil
}

// Upload -
func (c *HTTPClient) Upload(path string, values map[string]io.Reader) (resp *checkout.StatusResponse, err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}

		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	request, err := http.NewRequest("POST", c.URI+path, &b)
	if err != nil {
		return nil, err
	}
	// c.setHeader(request)
	c.setCredential(c.URI+path, request)
	// Don't forget to set the content type, this will contain the boundary.
	request.Header.Set("Content-Type", w.FormDataContentType())

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
	if response.StatusCode >= http.StatusBadRequest {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}

func (c *HTTPClient) setCredential(path string, req *http.Request) {

	if strings.Contains(path, "/tokens") {
		req.Header.Add("Authorization", c.PublicKey)
	} else {
		req.Header.Add("Authorization", c.SecretKey)
	}
}

func (c *HTTPClient) setHeader(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "checkout-sdk-go/"+checkout.ClientVersion)
}

func (c *HTTPClient) setIdempotencyKey(req *http.Request) {
	req.Header.Add("Cko-Idempotency-Key", c.IdempotencyKey)
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

// Delete ...
func (c *HTTPClient) Delete(param string) (*checkout.StatusResponse, error) {

	request, err := http.NewRequest(http.MethodDelete, c.URI+param, nil)
	if err != nil {
		return nil, err
	}
	c.setHeader(request)
	c.setCredential(c.URI+param, request)
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
	if response.StatusCode >= http.StatusBadRequest {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}

// Put ...
func (c *HTTPClient) Put(param string, body interface{}) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodPut, c.URI+param, body)
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
	if response.StatusCode >= http.StatusBadRequest {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}

// Patch ...
func (c *HTTPClient) Patch(param string, body interface{}) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodPatch, c.URI+param, body)
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
	if response.StatusCode >= http.StatusBadRequest {
		err := responseToError(apiResponse, responseBody)
		return apiResponse, err
	}
	return apiResponse, nil
}
