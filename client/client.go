package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
)

type HttpClient interface {
	Get(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error
	GetWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error
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

type ApiClient struct {
	HttpClient          http.Client
	BaseUri             string
	EnableTelemetry     bool
	RequestMetricsQueue common.TelemetryQueue
	Log                 configuration.StdLogger
	Retry               *configuration.RetryConfiguration
}

const (
	CkoRequestId       = "cko-request-id"
	CkoVersion         = "cko-version"
	CkoTelemetryHeader = "cko-sdk-telemetry"
)

func NewApiClient(configuration *configuration.Configuration, baseUri string) *ApiClient {
	return &ApiClient{
		HttpClient:          configuration.HttpClient,
		BaseUri:             baseUri,
		EnableTelemetry:     configuration.EnableTelemetry,
		RequestMetricsQueue: *common.NewTelemetryQueue(),
		Log:                 configuration.Logger,
		Retry:               configuration.Retry,
	}
}

func (a *ApiClient) Get(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return a.GetWithContext(context.Background(), path, authorization, responseMapping)
}

func (a *ApiClient) GetWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return a.invoke(ctx, http.MethodGet, path, authorization, nil, responseMapping, nil)
}

func (a *ApiClient) Post(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.PostWithContext(context.Background(), path, authorization, request, responseMapping, idempotencyKey)
}

func (a *ApiClient) PostWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.invoke(ctx, http.MethodPost, path, authorization, request, responseMapping, a.ensureIdempotencyKey(idempotencyKey))
}

func (a *ApiClient) Put(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.PutWithContext(context.Background(), path, authorization, request, responseMapping, idempotencyKey)
}

func (a *ApiClient) PutWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.invoke(ctx, http.MethodPut, path, authorization, request, responseMapping, a.ensureIdempotencyKey(idempotencyKey))
}

// ensureIdempotencyKey generates a Cko-Idempotency-Key for write requests when
// retries are enabled and the caller did not supply one, so a retried write is
// deduplicated server-side rather than applied twice. A caller-supplied key is
// left untouched, and when retries are disabled the historical behaviour (no
// generated key) is preserved.
func (a *ApiClient) ensureIdempotencyKey(idempotencyKey *string) *string {
	if a.Retry == nil || idempotencyKey != nil {
		return idempotencyKey
	}
	generated := uuid.NewString()
	return &generated
}

func (a *ApiClient) Patch(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	return a.PatchWithContext(context.Background(), path, authorization, request, responseMapping)
}

func (a *ApiClient) PatchWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	return a.invoke(ctx, http.MethodPatch, path, authorization, request, responseMapping, nil)
}

func (a *ApiClient) Delete(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return a.DeleteWithContext(context.Background(), path, authorization, responseMapping)
}

func (a *ApiClient) DeleteWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return a.invoke(ctx, http.MethodDelete, path, authorization, nil, responseMapping, nil)
}

func (a *ApiClient) Upload(path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error {
	return a.UploadWithContext(context.Background(), path, authorization, request, responseMapping)
}

func (a *ApiClient) UploadWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error {
	return a.submit(ctx, path, authorization, request, responseMapping)
}

func (a *ApiClient) PostForm(path string, authorization *configuration.SdkAuthorization, formData url.Values, responseMapping interface{}) error {
	return a.PostFormWithContext(context.Background(), path, authorization, formData, responseMapping)
}

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
	version := rawResponse.Header.Get(CkoVersion)
	body, err := a.readBody(ctx, rawResponse)
	if err != nil {
		return err
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

func (a *ApiClient) readBody(ctx context.Context, response *http.Response) ([]byte, error) {
	// Check if context is already cancelled before reading
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
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
		resp, err := a.send(ctx, req)
		elapsed := time.Since(start)
		if err != nil {
			return err
		}

		lastRequestMetric.PrevRequestDuration = int(elapsed.Milliseconds())
		lastRequestMetric.PrevRequestId = currentRequestId
		a.RequestMetricsQueue.Enqueue(lastRequestMetric)
		return a.handleResponse(ctx, resp, responseMapping)
	} else {
		resp, err := a.send(ctx, req)
		if err != nil {
			return err
		}

		return a.handleResponse(ctx, resp, responseMapping)
	}
}

// send executes req, retrying transient failures when retries are enabled. When
// a.Retry is nil it delegates directly to the underlying client, leaving the
// single-attempt behaviour unchanged regardless of the injected http.Client.
func (a *ApiClient) send(ctx context.Context, req *http.Request) (*http.Response, error) {
	if a.Retry == nil {
		return a.HttpClient.Do(req)
	}

	for attempt := 0; ; attempt++ {
		resp, err := a.HttpClient.Do(req)
		if !shouldRetry(resp, err, attempt, a.Retry) {
			return resp, err
		}

		delay := backoff(attempt, a.Retry, retryAfterDelay(resp))
		drainAndClose(resp)
		if sleepErr := sleep(ctx, delay); sleepErr != nil {
			return nil, sleepErr
		}
		if err := resetBody(req); err != nil {
			return resp, err
		}
	}
}

// resetBody restores a replayable request body for the next attempt. The stdlib
// populates GetBody for the *bytes.Buffer bodies this client constructs, so a
// fresh reader is available on every retry.
func resetBody(req *http.Request) error {
	if req.Body == nil || req.GetBody == nil {
		return nil
	}
	body, err := req.GetBody()
	if err != nil {
		return err
	}
	req.Body = body
	return nil
}
