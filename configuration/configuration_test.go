package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func boolPtr(v bool) *bool { return &v }

func buildSubdomain(t *testing.T) *EnvironmentSubdomain {
	subdomain, err := NewEnvironmentSubdomain(Sandbox(), "vkuhvk4v")
	assert.NoError(t, err)
	assert.NotNil(t, subdomain)
	return subdomain
}

func TestNewConfigurationWithSubdomainKeepsTelemetryEnabled(t *testing.T) {
	config := NewConfigurationWithSubdomain(nil, boolPtr(true), Sandbox(), buildSubdomain(t), nil, nil)

	assert.True(t, config.EnableTelemetry)
	assert.NotNil(t, config.EnvironmentSubdomain)
}

func TestNewConfigurationWithSubdomainDefaultsTelemetryOn(t *testing.T) {
	config := NewConfigurationWithSubdomain(nil, nil, Sandbox(), buildSubdomain(t), nil, nil)

	assert.True(t, config.EnableTelemetry)
	assert.NotNil(t, config.EnvironmentSubdomain)
}

func TestNewConfigurationWithSubdomainRespectsTelemetryOptOut(t *testing.T) {
	config := NewConfigurationWithSubdomain(nil, boolPtr(false), Sandbox(), buildSubdomain(t), nil, nil)

	assert.False(t, config.EnableTelemetry)
	assert.NotNil(t, config.EnvironmentSubdomain)
}
