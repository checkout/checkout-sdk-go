package errors

import (
	"context"
	stderrors "errors"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCheckoutTransportError(t *testing.T) {
	assert.Nil(t, NewCheckoutTransportError(nil))

	err := NewCheckoutTransportError(context.DeadlineExceeded)
	assert.IsType(t, CheckoutTransportError{}, err)
	assert.Equal(t, "transport error: context deadline exceeded", err.Error())
	assert.True(t, stderrors.Is(err, context.DeadlineExceeded))
	assert.True(t, err.(CheckoutTransportError).Timeout())

	assert.Equal(t, err, NewCheckoutTransportError(err), "an existing transport error is not wrapped again")

	nested := fmt.Errorf("outer: %w", err)
	wrapped := NewCheckoutTransportError(nested)
	if assert.IsType(t, CheckoutTransportError{}, wrapped, "a transport error nested deeper is wrapped at the top") {
		assert.Equal(t, nested, wrapped.(CheckoutTransportError).Err)
	}

	canceled := NewCheckoutTransportError(context.Canceled).(CheckoutTransportError)
	assert.False(t, canceled.Timeout())
	assert.Equal(t, "transport error", CheckoutTransportError{}.Error())
}

type timeoutError struct{ timeout bool }

func (e timeoutError) Error() string { return "timeout error" }
func (e timeoutError) Timeout() bool { return e.timeout }

// Timeout looks at every wrapped error, also below a wrapper with its own Timeout method.
func TestCheckoutTransportErrorTimeoutLooksAtTheWholeChain(t *testing.T) {
	below := &url.Error{Op: "Get", URL: "https://api.example.test", Err: fmt.Errorf("transport: %w", timeoutError{timeout: true})}
	assert.True(t, CheckoutTransportError{Err: below}.Timeout())

	deadline := &url.Error{Op: "Get", URL: "https://api.example.test", Err: fmt.Errorf("transport: %w", context.DeadlineExceeded)}
	assert.True(t, CheckoutTransportError{Err: deadline}.Timeout())

	notTimeout := &url.Error{Op: "Get", URL: "https://api.example.test", Err: fmt.Errorf("transport: %w", timeoutError{})}
	assert.False(t, CheckoutTransportError{Err: notTimeout}.Timeout())
}

type deadlineByIs struct{}

func (deadlineByIs) Error() string        { return "deadline" }
func (deadlineByIs) Is(target error) bool { return target == context.DeadlineExceeded }

// Timeout agrees with errors.Is for the deadline, also when it is matched through an Is method.
func TestCheckoutTransportErrorTimeoutAgreesWithErrorsIs(t *testing.T) {
	byIs := &url.Error{Op: "Get", URL: "https://api.example.test", Err: fmt.Errorf("transport: %w", deadlineByIs{})}
	assert.True(t, stderrors.Is(byIs, context.DeadlineExceeded))
	assert.True(t, CheckoutTransportError{Err: byIs}.Timeout())
}
