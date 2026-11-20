package client

import (
	"bytes"
	"context"
	"encoding/json"
	goerrors "errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
	"github.com/checkout/checkout-sdk-go/v3/internal/download"
	"github.com/checkout/checkout-sdk-go/v3/internal/loopback"
)

type HttpClient interface {
	Get(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error
	GetWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error
	Post(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error
	PostWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error
	Put(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error
	PutWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error
	Patch(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error
	PatchWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error
	Delete(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error
	DeleteWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error
	Upload(path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error
	UploadWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error
	PostForm(path string, authorization *configuration.SdkAuthorization, formData url.Values, responseMapping interface{}) error
	PostFormWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, formData url.Values, responseMapping interface{}) error
}

// ApiClient sends the requests of the SDK clients to one base URI and decodes their responses.
type ApiClient struct {
	// HttpClient sends the requests. Each request uses a copy of it with the SDK's own policy;
	// HttpClient itself is not changed. A custom Transport applies its own behaviour.
	HttpClient http.Client
	// BaseUri is prefixed to every request path.
	BaseUri string
	// EnableTelemetry sends the SDK telemetry header with each request when true.
	EnableTelemetry bool
	// RequestMetricsQueue holds the metrics of the previous request for the telemetry header.
	RequestMetricsQueue common.TelemetryQueue
	// Log receives the request log lines.
	Log configuration.StdLogger
}

const (
	CkoRequestId       = "cko-request-id"
	CkoVersion         = "cko-version"
	CkoTelemetryHeader = "cko-sdk-telemetry"
)

// NewApiClient creates an ApiClient for baseUri with the HttpClient, telemetry setting and Logger of
// configuration. The ApiClient keeps a copy of configuration.HttpClient.
func NewApiClient(configuration *configuration.Configuration, baseUri string) *ApiClient {
	return &ApiClient{
		HttpClient:          configuration.HttpClient,
		BaseUri:             baseUri,
		EnableTelemetry:     configuration.EnableTelemetry,
		RequestMetricsQueue: *common.NewTelemetryQueue(),
		Log:                 configuration.Logger,
	}
}

// Get is GetWithContext with context.Background().
func (a *ApiClient) Get(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	return a.GetWithContext(context.Background(), path, authorization, request, responseMapping)
}

// GetWithContext sends a GET request to path with authorization and decodes the response into
// responseMapping. request, when not nil, only supplies per-request headers: a GET has no body.
// It returns an errors.CheckoutAPIError for a response that is not a 2xx, and an
// errors.CheckoutTransportError when sending the request or reading its response fails.
func (a *ApiClient) GetWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	if request != nil {
		// request carries headers only (e.g. the Accept schema-version header); a GET must not send a body,
		// so build it with an empty body while still resolving headers from request.
		req, err := a.buildRequest(ctx, http.MethodGet, path, authorization, "application/json", new(bytes.Buffer), nil, request)
		if err != nil {
			return err
		}

		a.Log.Printf("%s: %s", http.MethodGet, path)
		return a.doRequest(ctx, req, responseMapping)
	}

	// No per-request headers: behave exactly as before, request = nil, no body, no headers to resolve.
	return a.invoke(ctx, http.MethodGet, path, authorization, request, responseMapping, nil)
}

// Post is PostWithContext with context.Background().
func (a *ApiClient) Post(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.PostWithContext(context.Background(), path, authorization, request, responseMapping, idempotencyKey)
}

// PostWithContext sends request as a JSON body in a POST to path with authorization and the
// idempotency key when not nil, and decodes the response into responseMapping.
// It returns an errors.CheckoutAPIError for a response that is not a 2xx, and an
// errors.CheckoutTransportError when sending the request or reading its response fails.
func (a *ApiClient) PostWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.invoke(ctx, http.MethodPost, path, authorization, request, responseMapping, idempotencyKey)
}

// Put is PutWithContext with context.Background().
func (a *ApiClient) Put(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.PutWithContext(context.Background(), path, authorization, request, responseMapping, idempotencyKey)
}

// PutWithContext sends request as a JSON body in a PUT to path with authorization and the
// idempotency key when not nil, and decodes the response into responseMapping.
// It returns an errors.CheckoutAPIError for a response that is not a 2xx, and an
// errors.CheckoutTransportError when sending the request or reading its response fails.
func (a *ApiClient) PutWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.invoke(ctx, http.MethodPut, path, authorization, request, responseMapping, idempotencyKey)
}

// Patch is PatchWithContext with context.Background().
func (a *ApiClient) Patch(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	return a.PatchWithContext(context.Background(), path, authorization, request, responseMapping)
}

// PatchWithContext sends request as a JSON body in a PATCH to path with authorization and decodes
// the response into responseMapping.
// It returns an errors.CheckoutAPIError for a response that is not a 2xx, and an
// errors.CheckoutTransportError when sending the request or reading its response fails.
func (a *ApiClient) PatchWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	return a.invoke(ctx, http.MethodPatch, path, authorization, request, responseMapping, nil)
}

// Delete is DeleteWithContext with context.Background().
func (a *ApiClient) Delete(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return a.DeleteWithContext(context.Background(), path, authorization, responseMapping)
}

// DeleteWithContext sends a DELETE request to path with authorization and decodes the response
// into responseMapping.
// It returns an errors.CheckoutAPIError for a response that is not a 2xx, and an
// errors.CheckoutTransportError when sending the request or reading its response fails.
func (a *ApiClient) DeleteWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return a.invoke(ctx, http.MethodDelete, path, authorization, nil, responseMapping, nil)
}

// Upload is UploadWithContext with context.Background().
func (a *ApiClient) Upload(path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error {
	return a.UploadWithContext(context.Background(), path, authorization, request, responseMapping)
}

// UploadWithContext sends request as a multipart form in a POST to path with authorization and
// decodes the response into responseMapping.
// It returns an errors.CheckoutAPIError for a response that is not a 2xx, and an
// errors.CheckoutTransportError when sending the request or reading its response fails.
func (a *ApiClient) UploadWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error {
	return a.submit(ctx, path, authorization, request, responseMapping)
}

// PostForm is PostFormWithContext with context.Background().
func (a *ApiClient) PostForm(path string, authorization *configuration.SdkAuthorization, formData url.Values, responseMapping interface{}) error {
	return a.PostFormWithContext(context.Background(), path, authorization, formData, responseMapping)
}

// PostFormWithContext sends formData as a URL-encoded form in a POST to path with authorization
// and decodes the response into responseMapping.
// It returns an errors.CheckoutAPIError for a response that is not a 2xx, and an
// errors.CheckoutTransportError when sending the request or reading its response fails.
func (a *ApiClient) PostFormWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, formData url.Values, responseMapping interface{}) error {
	return a.submitForm(ctx, path, authorization, formData, responseMapping)
}

func (a *ApiClient) invoke(
	ctx context.Context,
	method string,
	path string,
	authorization *configuration.SdkAuthorization,
	request interface{},
	responseMapping interface{},
	idempotencyKey *string,
) error {
	body, err := common.Marshal(request)
	if err != nil {
		return err
	}

	req, err := a.buildRequest(ctx, method, path, authorization, "application/json", body, idempotencyKey, request)
	if err != nil {
		return err
	}

	a.Log.Printf("%s: %s", method, path)
	return a.doRequest(ctx, req, responseMapping)
}

func applyRequestHeaders(request interface{}, headers http.Header) {
	v, ok := resolveToStruct(reflect.ValueOf(request))
	if !ok {
		return
	}
	headersField, ok := resolveToStruct(v.FieldByName("Headers"))
	if !ok {
		return
	}
	setHeadersFromFields(headersField, headers)
}

func resolveToStruct(v reflect.Value) (reflect.Value, bool) {
	if !v.IsValid() {
		return v, false
	}
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return v, false
		}
		v = v.Elem()
	}
	return v, v.Kind() == reflect.Struct
}

func setHeadersFromFields(v reflect.Value, headers http.Header) {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		name := headerNameFromTag(t.Field(i).Tag.Get("json"))
		if name == "" {
			continue
		}
		if value := v.Field(i); value.Kind() == reflect.String && value.String() != "" {
			headers.Set(name, value.String())
		}
	}
}

func headerNameFromTag(tag string) string {
	if tag == "" || tag == "-" {
		return ""
	}
	name := strings.SplitN(tag, ",", 2)[0]
	if name == "-" {
		return ""
	}
	return name
}

// 2026/04/27 DRY - At some point avoid this hardcoding and use reflection to build up the body and content type in buildRequest directly
// we do not need here to know about FileUploadRequest object!!!
func (a *ApiClient) submit(
	ctx context.Context,
	path string,
	authorization *configuration.SdkAuthorization,
	request *common.FileUploadRequest,
	responseMapping interface{},
) error {
	req, err := a.buildRequest(ctx, http.MethodPost, path, authorization, request.W.FormDataContentType(), request.B, nil, request)
	if err != nil {
		return err
	}

	a.Log.Printf("post: %s", path)
	return a.doRequest(ctx, req, responseMapping)
}

func (a *ApiClient) submitForm(
	ctx context.Context,
	path string,
	authorization *configuration.SdkAuthorization,
	formData url.Values,
	responseMapping interface{},
) error {
	body := bytes.NewBufferString(formData.Encode())
	req, err := a.buildRequest(ctx, http.MethodPost, path, authorization, "application/x-www-form-urlencoded", body, nil, nil)
	if err != nil {
		return err
	}

	a.Log.Printf("post: %s", path)
	return a.doRequest(ctx, req, responseMapping)
}

func (a *ApiClient) buildRequest(
	ctx context.Context,
	method string,
	path string,
	authorization *configuration.SdkAuthorization,
	contentType string,
	body *bytes.Buffer,
	idempotencyKey *string,
	request interface{},
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, a.BaseUri+path, body)
	if err != nil {
		return nil, err
	}

	var authorizationHeader string
	if authorization != nil {
		authorizationHeader, err = authorization.GetAuthorizationHeader()
		if err != nil {
			return nil, err
		}
	}

	req.Header = a.getHeaders(contentType, authorizationHeader, idempotencyKey, request)

	return req, nil
}

func (a *ApiClient) handleResponse(ctx context.Context, rawResponse *http.Response, responseMapping interface{}) error {
	requestId := rawResponse.Header.Get(CkoRequestId)
	if isRedirect(rawResponse.StatusCode) {
		if rawResponse.Body != nil {
			_ = rawResponse.Body.Close()
		}
		return statusOnlyError(rawResponse.StatusCode, requestId)
	}

	version := rawResponse.Header.Get(CkoVersion)
	body, err := a.readBody(ctx, rawResponse)
	if err != nil {
		return errors.NewCheckoutTransportError(withoutDownloadUrl(ctx, err))
	}

	if rawResponse.StatusCode >= http.StatusBadRequest {
		return errors.HandleError(rawResponse.StatusCode, rawResponse.Status, requestId, body)
	}

	metadata := &common.HttpMetadata{
		Status:       rawResponse.Status,
		StatusCode:   rawResponse.StatusCode,
		ResponseBody: body,
		Headers: &common.Headers{
			Header:       rawResponse.Header,
			CKORequestID: &requestId,
			CKOVersion:   &version,
		},
	}

	return common.Unmarshal(metadata, responseMapping)
}

func (a *ApiClient) getHeaders(contentType string, authorization string, idempotencyKey *string, request interface{}) http.Header {
	headers := make(http.Header)

	headers.Set("User-Agent", "checkout-sdk-go/"+SDK_VERSION)
	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", contentType)
	if authorization != "" {
		headers.Set("Authorization", authorization)
	}
	if idempotencyKey != nil {
		headers.Set("Cko-Idempotency-Key", *idempotencyKey)
	}

	applyRequestHeaders(request, headers)

	return headers
}

// readBody returns the response body, or nil when the response has none. It returns the context
// error when ctx is already done, and the read error when the body cannot be read.
func (a *ApiClient) readBody(ctx context.Context, response *http.Response) ([]byte, error) {
	// Check if context is already cancelled before reading
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if response.Body == nil {
		return nil, nil
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		if errTemp := Body.Close(); errTemp != nil {
			err = errTemp
		}
	}(response.Body)

	return body, err
}

func (a *ApiClient) doRequest(ctx context.Context, req *http.Request, responseMapping interface{}) error {
	if a.EnableTelemetry {
		currentRequestId := uuid.New().String()
		var lastRequestMetric common.RequestMetrics
		lastRequestMetric, ok := a.RequestMetricsQueue.Dequeue()
		if ok {
			lastRequestMetric.RequestId = currentRequestId
			lastRequestMetricStr, err := json.Marshal(lastRequestMetric)
			if err != nil {
				return err
			}
			req.Header.Set(CkoTelemetryHeader, string(lastRequestMetricStr))
		}
		start := time.Now()
		resp, err := a.httpClientFor(ctx).Do(req)
		elapsed := time.Since(start)
		if apiErr, refused := refusedRedirect(err); refused {
			return apiErr
		}
		if err != nil {
			return errors.NewCheckoutTransportError(withoutDownloadUrl(ctx, err))
		}

		lastRequestMetric.PrevRequestDuration = int(elapsed.Milliseconds())
		lastRequestMetric.PrevRequestId = currentRequestId
		a.RequestMetricsQueue.Enqueue(lastRequestMetric)
		return a.handleResponse(ctx, resp, responseMapping)
	} else {
		resp, err := a.httpClientFor(ctx).Do(req)
		if apiErr, refused := refusedRedirect(err); refused {
			return apiErr
		}
		if err != nil {
			return errors.NewCheckoutTransportError(withoutDownloadUrl(ctx, err))
		}

		return a.handleResponse(ctx, resp, responseMapping)
	}
}

// httpClientFor returns the client that sends one request: a copy of HttpClient that keeps its
// settings, such as its Transport and Timeout, and follows no redirect, whatever CheckRedirect
// sets. When ctx comes from download.Mark, the copy follows the download redirect only (see
// followDownloadRedirect). HttpClient itself is not changed.
func (a *ApiClient) httpClientFor(ctx context.Context) *http.Client {
	client := a.HttpClient
	marked := download.IsMarked(ctx)
	next := client.Transport
	if next == nil {
		next = http.DefaultTransport
	}
	client.Transport = &redirectGuard{next: next, download: marked}
	if marked {
		client.CheckRedirect = followDownloadRedirect
	} else {
		client.CheckRedirect = refuseRedirect
	}
	return &client
}

// redirectGuard turns a 3xx response into a redirectRefusal before net/http handles it, so every
// 3xx gives the same error whatever its Location. The only response let through is the download
// redirect (see followDownloadRedirect) with a Location that parses.
type redirectGuard struct {
	next     http.RoundTripper
	download bool
}

// CancelRequest forwards to the wrapped Transport when it supports it, so that a Client.Timeout
// still cancels a Transport that relies on CancelRequest.
func (g *redirectGuard) CancelRequest(req *http.Request) {
	if canceler, ok := g.next.(interface{ CancelRequest(*http.Request) }); ok {
		canceler.CancelRequest(req)
	}
}

func (g *redirectGuard) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := g.next.RoundTrip(req)
	if err != nil || resp == nil {
		return resp, err
	}
	// The download URL answering an error gives the status only, never its body.
	downloadHop := g.download && req.Response != nil
	if !isRedirect(resp.StatusCode) && !(downloadHop && resp.StatusCode >= http.StatusBadRequest) {
		return resp, nil
	}
	if g.download && req.Response == nil && resp.StatusCode == http.StatusFound && hasValidLocation(req, resp) {
		return resp, nil
	}
	refusal := &redirectRefusal{statusCode: resp.StatusCode, requestId: resp.Header.Get(CkoRequestId)}
	if resp.Body != nil {
		_ = resp.Body.Close()
	}
	return nil, refusal
}

// hasValidLocation reports whether the Location of resp is present and parses against the request
// URL, as net/http resolves it.
func hasValidLocation(req *http.Request, resp *http.Response) bool {
	location := resp.Header.Get("Location")
	if location == "" {
		return false
	}
	_, err := req.URL.Parse(location)
	return err == nil
}

// redirectRefusal carries a refused 3xx, or an error of the download URL, from redirectGuard to
// doRequest, which turns it into the CheckoutAPIError of statusOnlyError. Its text names no Location.
type redirectRefusal struct {
	statusCode int
	requestId  string
}

func (r *redirectRefusal) Error() string { return "redirect not followed" }

// withoutDownloadUrl keeps a request marked by download.Mark from returning its download URL in an
// error: the URL of a *url.Error in err loses its query, fragment and user information. The error
// keeps its type, Timeout and Unwrap.
func withoutDownloadUrl(ctx context.Context, err error) error {
	if !download.IsMarked(ctx) {
		return err
	}
	var urlErr *url.Error
	if goerrors.As(err, &urlErr) {
		parsed, parseErr := url.Parse(urlErr.URL)
		if parseErr != nil {
			urlErr.URL = "(download URL)"
		} else {
			parsed.RawQuery, parsed.Fragment, parsed.User = "", "", nil
			urlErr.URL = parsed.String()
		}
	}
	return err
}

// refusedRedirect returns the CheckoutAPIError for err when err comes from redirectGuard.
func refusedRedirect(err error) (errors.CheckoutAPIError, bool) {
	var refusal *redirectRefusal
	if goerrors.As(err, &refusal) {
		return statusOnlyError(refusal.statusCode, refusal.requestId), true
	}
	return errors.CheckoutAPIError{}, false
}

// refuseRedirect makes the client return the redirect response instead of following it.
func refuseRedirect(*http.Request, []*http.Request) error {
	return http.ErrUseLastResponse
}

// followDownloadRedirect follows a single 302 to an https URL (http only for a loopback host, for
// local tests), as a GET without the Authorization header, without the SDK's own headers (any
// Cko- header, the telemetry header included), without a Referer and without a body (net/http
// sends the request that follows a 302 as a GET with no body; a body is refused anyway). Any other
// redirect, a second one included, is returned as the response instead of being followed. Only
// GetReportFile, a GET, marks its context for it.
func followDownloadRedirect(req *http.Request, via []*http.Request) error {
	if len(via) != 1 || req.Response == nil || req.Response.StatusCode != http.StatusFound ||
		(req.Body != nil && req.Body != http.NoBody) || !isDownloadScheme(req.URL) {
		return http.ErrUseLastResponse
	}
	req.Header.Del("Authorization")
	req.Header.Del("Referer")
	for name := range req.Header {
		if len(name) >= len(ckoHeaderPrefix) && strings.EqualFold(name[:len(ckoHeaderPrefix)], ckoHeaderPrefix) {
			delete(req.Header, name)
		}
	}
	req.URL.User = nil
	return nil
}

// ckoHeaderPrefix starts the name of every header the SDK sets for the API only.
const ckoHeaderPrefix = "cko-"

// isDownloadScheme reports whether the download URL may be requested: https, or http for a
// loopback host.
func isDownloadScheme(u *url.URL) bool {
	switch u.Scheme {
	case "https":
		return true
	case "http":
		return loopback.IsHost(u.Hostname())
	}
	return false
}

func isRedirect(statusCode int) bool {
	return statusCode >= http.StatusMultipleChoices && statusCode < http.StatusBadRequest
}

// statusOnlyError is the error for a 3xx response and for an error of the download URL (see
// redirectGuard). It carries
// the status code, the standard status text and the request id only: the server reason phrase, the
// Location and the body are not used.
func statusOnlyError(statusCode int, requestId string) errors.CheckoutAPIError {
	return errors.CheckoutAPIError{
		StatusCode: statusCode,
		Status:     strings.TrimSpace(fmt.Sprintf("%d %s", statusCode, http.StatusText(statusCode))),
		Data:       &errors.ErrorDetails{RequestID: requestId},
	}
}
