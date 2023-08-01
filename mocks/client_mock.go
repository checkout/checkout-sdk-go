package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
)

type (
	ApiClientMock struct{ mock.Mock }
)

func (m *ApiClientMock) Get(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	return m.GetWithContext(context.Background(), path, authorization, responseMapping)
}

func (m *ApiClientMock) GetWithContext(ctx context.Context, path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
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
