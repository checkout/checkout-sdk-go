package configuration

import "time"

// Default retry parameters, chosen to give callers opting in sensible resilience
// to transient failures.
const (
	DefaultMaxRetries = 2
	DefaultMinBackoff = 500 * time.Millisecond
	DefaultMaxBackoff = 5 * time.Second
)

// RetryConfiguration controls how the client retries requests that fail with a
// transient error (connection failures, 409, 429 and 5xx responses). It is
// opt-in: a nil RetryConfiguration on Configuration disables retries entirely,
// preserving the historical behaviour of executing each request exactly once.
type RetryConfiguration struct {
	// MaxRetries is the number of additional attempts made after the initial
	// request. A value of 2 therefore allows up to 3 attempts in total.
	MaxRetries int
	// MinBackoff is the base delay applied before the first retry and the floor
	// for the exponential backoff.
	MinBackoff time.Duration
	// MaxBackoff caps the delay between attempts, including any server-provided
	// Retry-After hint.
	MaxBackoff time.Duration
}

// DefaultRetryConfiguration returns a RetryConfiguration with sensible defaults:
// two retries with exponential backoff between 500ms and 5s.
func DefaultRetryConfiguration() *RetryConfiguration {
	return &RetryConfiguration{
		MaxRetries: DefaultMaxRetries,
		MinBackoff: DefaultMinBackoff,
		MaxBackoff: DefaultMaxBackoff,
	}
}
