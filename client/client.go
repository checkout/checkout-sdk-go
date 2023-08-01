package client

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
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
}

type ApiClient struct {
	HttpClient http.Client
	BaseUri    string
	Log        configuration.StdLogger
}

const (
	CkoRequestId = "cko-request-id"
	CkoVersion   = "cko-version"
)

func NewApiClient(configuration *configuration.Configuration, baseUri string) *ApiClient {
	return &ApiClient{
		HttpClient: configuration.HttpClient,
		BaseUri:    baseUri,
		Log:        configuration.Logger,
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
	return a.invoke(ctx, http.MethodPost, path, authorization, request, responseMapping, idempotencyKey)
}

func (a *ApiClient) Put(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.PutWithContext(context.Background(), path, authorization, request, responseMapping, idempotencyKey)
}

func (a *ApiClient) PutWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return a.invoke(ctx, http.MethodPut, path, authorization, request, responseMapping, idempotencyKey)
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

	req, err := a.buildRequest(ctx, method, path, authorization, "application/json", body, idempotencyKey)
	if err != nil {
		return err
	}

	a.Log.Printf("%s: %s", method, path)
	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return err
	}

	return a.handleResponse(resp, responseMapping)
}

func (a *ApiClient) submit(
	ctx context.Context,
	path string,
	authorization *configuration.SdkAuthorization,
	request *common.FileUploadRequest,
	responseMapping interface{},
) error {
	req, err := a.buildRequest(
		ctx,
		http.MethodPost,
		path,
		authorization,
		request.W.FormDataContentType(),
		request.B,
		nil,
	)
	if err != nil {
		return err
	}

	a.Log.Printf("post: %s", path)
	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return err
	}

	return a.handleResponse(resp, responseMapping)
}

func (a *ApiClient) buildRequest(
	ctx context.Context,
	method string,
	path string,
	authorization *configuration.SdkAuthorization,
	contentType string,
	body *bytes.Buffer,
	idempotencyKey *string,
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, a.BaseUri+path, body)
	if err != nil {
		return nil, err
	}

	authorizationHeader, err := authorization.GetAuthorizationHeader()
	if err != nil {
		return nil, err
	}

	req.Header = a.getHeaders(contentType, authorizationHeader, idempotencyKey)

	return req, nil
}

func (a *ApiClient) handleResponse(rawResponse *http.Response, responseMapping interface{}) error {
	requestId := rawResponse.Header.Get(CkoRequestId)
	version := rawResponse.Header.Get(CkoVersion)
	body, err := a.readBody(rawResponse)
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

func (a *ApiClient) getHeaders(contentType string, authorization string, idempotencyKey *string) http.Header {
	headers := make(http.Header)

	headers.Set("User-Agent", "checkout-sdk-go/"+SDK_VERSION)
	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", contentType)
	headers.Set("Authorization", authorization)
	if idempotencyKey != nil {
		headers.Set("Cko-Idempotency-Key", *idempotencyKey)
	}

	return headers
}

func (a *ApiClient) readBody(response *http.Response) ([]byte, error) {
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
