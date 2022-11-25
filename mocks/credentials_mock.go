package mocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go-beta/configuration"
)

type (
	CredentialsMock struct{ mock.Mock }
)

func (m *CredentialsMock) GetAuthorization(authorizationType configuration.AuthorizationType) (*configuration.SdkAuthorization, error) {
	args := m.Called(authorizationType)

	if args.Get(0) != nil {
		return args.Get(0).(*configuration.SdkAuthorization), nil
	}

	return nil, args.Get(1).(error)
}
