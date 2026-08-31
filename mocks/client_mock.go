package mocks

import (
	"context"
	"net/url"

	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
)

type (
	ApiClientMock struct {
		mock.Mock
		// CapturedGetRequest records the request argument (per-request headers source) passed to the
		// last Get/GetWithContext call. The request is intentionally not forwarded to m.Called so
		// existing GET expectations keep their original argument shape.
		CapturedGetRequest interface{}
	}
)

func (m *ApiClientMock) Get(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	return m.GetWithContext(context.Background(), path, authorization, request, responseMapping)
}

func (m *ApiClientMock) GetWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	m.CapturedGetRequest = request
	args := m.Called(ctx, path, authorization, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Post(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return m.PostWithContext(context.Background(), path, authorization, request, responseMapping, idempotencyKey)
}

func (m *ApiClientMock) PostWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	args := m.Called(ctx, path, authorization, request, responseMapping, idempotencyKey)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Put(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	return m.PutWithContext(context.Background(), path, authorization, request, responseMapping, idempotencyKey)
}

func (m *ApiClientMock) PutWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	args := m.Called(ctx, path, authorization, request, responseMapping, idempotencyKey)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Patch(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	return m.PatchWithContext(context.Background(), path, authorization, request, responseMapping)
}

func (m *ApiClientMock) PatchWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	args := m.Called(ctx, path, authorization, request, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Delete(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return m.DeleteWithContext(context.Background(), path, authorization, responseMapping)
}

func (m *ApiClientMock) DeleteWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	args := m.Called(ctx, path, authorization, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Upload(path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error {
	return m.UploadWithContext(context.Background(), path, authorization, request, responseMapping)
}

func (m *ApiClientMock) UploadWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error {
	args := m.Called(ctx, path, authorization, request, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) PostForm(path string, authorization *configuration.SdkAuthorization, formData url.Values, responseMapping interface{}) error {
	return m.PostFormWithContext(context.Background(), path, authorization, formData, responseMapping)
}

func (m *ApiClientMock) PostFormWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, formData url.Values, responseMapping interface{}) error {
	args := m.Called(ctx, path, authorization, formData, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}
