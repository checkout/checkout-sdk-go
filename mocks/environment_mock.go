package mocks

import "github.com/stretchr/testify/mock"

type (
	EnvironmentMock struct{ mock.Mock }
)

func (m *EnvironmentMock) BaseUri() string {
	return ""
}

func (m *EnvironmentMock) AuthorizationUri() string {
	return ""
}

func (m *EnvironmentMock) FilesUri() string {
	return ""
}

func (m *EnvironmentMock) TransfersUri() string {
	return ""
}

func (m *EnvironmentMock) BalancesUri() string {
	return ""
}

func (m *EnvironmentMock) IsSandbox() bool {
	return true
}
