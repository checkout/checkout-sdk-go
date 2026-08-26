package configuration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRetryConfiguration(t *testing.T) {
	cfg := DefaultRetryConfiguration()

	assert.Equal(t, DefaultMaxRetries, cfg.MaxRetries)
	assert.Equal(t, 500*time.Millisecond, cfg.MinBackoff)
	assert.Equal(t, 5*time.Second, cfg.MaxBackoff)
}
