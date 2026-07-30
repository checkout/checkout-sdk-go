package nas

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

func TestWithRetriesEnablesDefaults(t *testing.T) {
	b := &CheckoutDefaultSdkBuilder{}

	b.WithRetries()

	assert.Equal(t, configuration.DefaultRetryConfiguration(), b.Retry)
}

func TestWithRetryConfigurationSetsCustomValues(t *testing.T) {
	b := &CheckoutOAuthSdkBuilder{}
	custom := &configuration.RetryConfiguration{
		MaxRetries: 5,
		MinBackoff: 10 * time.Millisecond,
		MaxBackoff: time.Second,
	}

	b.WithRetryConfiguration(custom)

	assert.Same(t, custom, b.Retry)
}

func TestRetryDisabledWhenNotConfigured(t *testing.T) {
	b := &CheckoutDefaultSdkBuilder{}

	assert.Nil(t, b.Retry, "retries must be opt-in")
}
