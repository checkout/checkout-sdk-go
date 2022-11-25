package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
)

type (
	ApiClientMock struct{ mock.Mock }
)

func (m *ApiClientMock) Get(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	args := m.Called(path, authorization, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Post(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	args := m.Called(path, authorization, request, responseMapping, idempotencyKey)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Put(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}, idempotencyKey *string) error {
	args := m.Called(path, authorization, request, responseMapping, idempotencyKey)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Patch(path string, authorization *configuration.SdkAuthorization, request interface{}, responseMapping interface{}) error {
	args := m.Called(path, authorization, request, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Delete(path string, authorization *configuration.SdkAuthorization, responseMapping interface{}) error {
	args := m.Called(path, authorization, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}

func (m *ApiClientMock) Upload(path string, authorization *configuration.SdkAuthorization, request *common.FileUploadRequest, responseMapping interface{}) error {
	args := m.Called(path, authorization, request, responseMapping)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}
