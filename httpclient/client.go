package httpclient

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

var client *HTTPClient

// HTTPClient ...
type HTTPClient struct {
	HTTPClient           *http.Client
	PublicKey            string
	SecretKey            string
	URI                  string
	LeveledLogger        checkout.LeveledLoggerInterface
	MaxNetworkRetries    int64
	networkRetriesSleep  bool
	BearerAuthentication bool
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
		HTTPClient:           config.HTTPClient,
		PublicKey:            config.PublicKey,
		SecretKey:            config.SecretKey,
		URI:                  checkout.StringValue(config.URI),
		LeveledLogger:        config.LeveledLogger,
		MaxNetworkRetries:    *config.MaxNetworkRetries,
		networkRetriesSleep:  true,
		BearerAuthentication: config.BearerAuthentication,
	}
	return client
}

// Get ...
func (c *HTTPClient) Get(path string) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodGet, c.URI+path, nil)
	if err != nil {
		return nil, err
	}
	c.setContentType(request)
	c.setUserAgent(request)
	c.setAuthorization(c.URI+path, request)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	requestID := response.Header.Get(checkout.CKORequestID)
	version := response.Header.Get(checkout.CKOVersion)
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			Header:       response.Header,
			CKORequestID: &requestID,
			CKOVersion:   &version,
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
func (c *HTTPClient) Post(path string, body interface{}, params *checkout.Params) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodPost, c.URI+path, body)
	if err != nil {
		return nil, err
	}
	c.setContentType(request)
	c.setUserAgent(request)
	c.setAuthorization(c.URI+path, request)
	c.setIdempotencyKey(request, params)

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	requestID := response.Header.Get(checkout.CKORequestID)
	version := response.Header.Get(checkout.CKOVersion)
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: &requestID,
			CKOVersion:   &version,
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

	if body != nil {
		requestBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		request, err := http.NewRequest(method, path, bytes.NewBuffer(requestBody))
		if err != nil {
			return nil, err
		}
		return request, nil
	}
	request, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

// Upload -
func (c *HTTPClient) Upload(path string, boundary string, body *bytes.Buffer) (resp *checkout.StatusResponse, err error) {

	contentType := "multipart/form-data; boundary=" + boundary
	request, err := c.NewRequest(http.MethodPost, c.URI+path, nil)
	if err != nil {
		return nil, err
	}
	if body != nil {
		reader := bytes.NewReader(body.Bytes())
		request.Body = nopReadCloser{reader}
		request.GetBody = func() (io.ReadCloser, error) {
			reader := bytes.NewReader(body.Bytes())
			return nopReadCloser{reader}, nil
		}
	}
	c.setAuthorization(c.URI+path, request)
	c.setUserAgent(request)
	request.Header.Add("Content-Type", contentType)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	requestID := response.Header.Get(checkout.CKORequestID)
	version := response.Header.Get(checkout.CKOVersion)
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: &requestID,
			CKOVersion:   &version,
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

	request, err := c.NewRequest(http.MethodGet, c.URI+path, nil)
	// Setting headers if needed
	c.setAuthorization(c.URI+path, request)
	c.setUserAgent(request)
	request.Header.Add("Content-Type", "text/csv;")

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	requestID := response.Header.Get(checkout.CKORequestID)
	version := response.Header.Get(checkout.CKOVersion)
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: &requestID,
			CKOVersion:   &version,
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

func (c *HTTPClient) setAuthorization(path string, req *http.Request) {

	if strings.Contains(path, "/tokens") {
		req.Header.Add("Authorization", getAuthorizationHeader(c.PublicKey, c.BearerAuthentication))
	} else {
		req.Header.Add("Authorization", getAuthorizationHeader(c.SecretKey, c.BearerAuthentication))
	}
}

func getAuthorizationHeader(key string, bearerAuthentication bool) string {
	if bearerAuthentication {
		return "Bearer " + key
	}
	return key
}

func (c *HTTPClient) setUserAgent(req *http.Request) {
	req.Header.Add("User-Agent", "checkout-sdk-go/"+checkout.ClientVersion)
}

func (c *HTTPClient) setContentType(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
}

func isHTTPWriteMethod(method string) bool {
	return method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch || method == http.MethodDelete
}

func (c *HTTPClient) setIdempotencyKey(req *http.Request, params *checkout.Params) {
	if params != nil {
		if params.IdempotencyKey != nil {
			idempotencyKey := strings.TrimSpace(*params.IdempotencyKey)
			req.Header.Add("Cko-Idempotency-Key", idempotencyKey)
		} else if isHTTPWriteMethod(req.Method) {
			req.Header.Add("Cko-Idempotency-Key", checkout.NewIdempotencyKey())
		}
		for k, v := range params.Headers {
			for _, line := range v {
				// Use Set to override the default value possibly set before
				req.Header.Set(k, line)
			}
		}
	}
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
func (c *HTTPClient) Delete(path string) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodDelete, c.URI+path, nil)
	if err != nil {
		return nil, err
	}
	c.setContentType(request)
	c.setUserAgent(request)
	c.setAuthorization(c.URI+path, request)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	requestID := response.Header.Get(checkout.CKORequestID)
	version := response.Header.Get(checkout.CKOVersion)
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: &requestID,
			CKOVersion:   &version,
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
func (c *HTTPClient) Put(path string, body interface{}) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodPut, c.URI+path, body)
	if err != nil {
		return nil, err
	}
	c.setContentType(request)
	c.setUserAgent(request)
	c.setAuthorization(c.URI+path, request)

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	requestID := response.Header.Get(checkout.CKORequestID)
	version := response.Header.Get(checkout.CKOVersion)
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: &requestID,
			CKOVersion:   &version,
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
func (c *HTTPClient) Patch(path string, body interface{}) (*checkout.StatusResponse, error) {

	request, err := c.NewRequest(http.MethodPatch, c.URI+path, body)
	if err != nil {
		return nil, err
	}
	c.setContentType(request)
	c.setUserAgent(request)
	c.setAuthorization(c.URI+path, request)

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	requestID := response.Header.Get(checkout.CKORequestID)
	version := response.Header.Get(checkout.CKOVersion)
	apiResponse := &checkout.StatusResponse{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers: &checkout.Headers{
			CKORequestID: &requestID,
			CKOVersion:   &version,
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
