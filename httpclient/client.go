package httpclient

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/common"
)

var client *HTTPClient

const (
	defaultHTTPTimeout = 30 * time.Second
)

const (
	ckoRequestID = "cko-request-id"
	ckoVersion   = "cko-version"
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

type nopReadCloser struct {
	io.Reader
}

func (nopReadCloser) Close() error { return nil }

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
		Headers: &checkout.Headers{
			CKORequestID: response.Header.Get(ckoRequestID),
			CKOVersion:   response.Header.Get(ckoVersion),
		},
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
		Headers: &checkout.Headers{
			CKORequestID: response.Header.Get(ckoRequestID),
			CKOVersion:   response.Header.Get(ckoVersion),
		},
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
func (c *HTTPClient) Upload(path string, boundary string, body *bytes.Buffer) (resp *checkout.StatusResponse, err error) {

	contentType := "multipart/form-data; boundary=" + boundary
	req, err := http.NewRequest(http.MethodPost, c.URI+path, nil)
	if err != nil {
		return nil, err
	}
	if body != nil {
		reader := bytes.NewReader(body.Bytes())
		req.Body = nopReadCloser{reader}
		req.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body.Bytes())
			return nopReadCloser{reader}, nil
		}
	}
	c.setCredential(c.URI+path, req)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", "checkout-sdk-go/"+checkout.ClientVersion)
	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: response.Header.Get(ckoRequestID),
			CKOVersion:   response.Header.Get(ckoVersion),
		},
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

// Download -
func (c *HTTPClient) Download(path string) (resp *checkout.StatusResponse, err error) {

	req, err := http.NewRequest(http.MethodGet, c.URI+path, nil)
	// Setting headers if needed
	c.setCredential(c.URI+path, req)
	contentType := "text/csv;"
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", "checkout-sdk-go/"+checkout.ClientVersion)

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: response.Header.Get(ckoRequestID),
			CKOVersion:   response.Header.Get(ckoVersion),
		},
	}
	reader := csv.NewReader(response.Body)
	data, err := reader.ReadAll()
	if err != nil {
		return apiResponse, err
	}
	apiResponse.ResponseCSV = data
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
		Headers: &checkout.Headers{
			CKORequestID: response.Header.Get(ckoRequestID),
			CKOVersion:   response.Header.Get(ckoVersion),
		},
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
		Headers: &checkout.Headers{
			CKORequestID: response.Header.Get(ckoRequestID),
			CKOVersion:   response.Header.Get(ckoVersion),
		},
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
		Headers: &checkout.Headers{
			CKORequestID: response.Header.Get(ckoRequestID),
			CKOVersion:   response.Header.Get(ckoVersion),
		},
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
