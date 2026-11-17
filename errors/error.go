package errors

import (
	"context"
	stderrors "errors"
	"fmt"
)

type ErrorDetails struct {
	RequestID  string                 `json:"request_id,omitempty"`
	ErrorType  string                 `json:"error_type,omitempty"`
	ErrorCodes []string               `json:"error_codes,omitempty"`
	Id         string                 `json:"id,omitempty"`
	Links      map[string]interface{} `json:"_links,omitempty"`
	// VariantId is present on Inventory API error responses for the insufficient_stock error
	// code (404/409/422 on the /inventory/* endpoints). Empty for every other domain's errors.
	VariantId string `json:"variant_id,omitempty"`
	// Available is present on Inventory API error responses for the insufficient_stock error
	// code, alongside VariantId. Empty for every other domain's errors.
	Available int64 `json:"available,omitempty"`
}

type (
	CheckoutArgumentError      string
	CheckoutAuthorizationError string

	CheckoutAPIError struct {
		StatusCode int
		Status     string
		Data       *ErrorDetails
	}

	CheckoutOAuthError struct {
		Description string `json:"error"`
	}
)

func (e CheckoutArgumentError) Error() string      { return string(e) }
func (e CheckoutAuthorizationError) Error() string { return string(e) }
func (e CheckoutAPIError) Error() string           { return e.Status }
func (e CheckoutOAuthError) Error() string         { return e.Description }

type (
	UnsupportedTypeError string
	BadRequestError      string
	InternalError        string
)

func (e UnsupportedTypeError) Error() string { return string(e) }
func (e BadRequestError) Error() string      { return string(e) }
func (e InternalError) Error() string        { return string(e) }

func InvalidKey(key string) CheckoutAuthorizationError {
	return CheckoutAuthorizationError(fmt.Sprintf("%s is required for this operation", key))
}

func InvalidAuthorizationType(authType string) CheckoutAuthorizationError {
	return CheckoutAuthorizationError(fmt.Sprintf("Operation requires %s authorization type", authType))
}

// CheckoutTransportError is returned when a request could not be sent or its response could not
// be received: connection failures, TLS errors, timeouts and cancelled contexts. Err is the
// underlying error, so errors.Is(err, context.DeadlineExceeded) and errors.As with *url.Error
// still work.
type CheckoutTransportError struct {
	Err error
}

// Error returns "transport error" followed by the message of the underlying error.
func (e CheckoutTransportError) Error() string {
	if e.Err == nil {
		return "transport error"
	}
	return "transport error: " + e.Err.Error()
}

// Unwrap returns the underlying error, so errors.Is and errors.As reach it.
func (e CheckoutTransportError) Unwrap() error { return e.Err }

// Timeout reports whether the underlying error is a timeout: true when errors.Is(err,
// context.DeadlineExceeded) is true, or when err or an error in its errors.Unwrap chain has a
// Timeout method that returns true.
func (e CheckoutTransportError) Timeout() bool {
	if stderrors.Is(e.Err, context.DeadlineExceeded) {
		return true
	}
	for err := e.Err; err != nil; err = stderrors.Unwrap(err) {
		if timeout, ok := err.(interface{ Timeout() bool }); ok && timeout.Timeout() {
			return true
		}
	}
	return false
}

// NewCheckoutTransportError wraps err in a CheckoutTransportError. It returns nil when err is nil
// and returns err unchanged when err itself is a CheckoutTransportError; one nested deeper, for
// example inside a *url.Error, is wrapped again so the returned error is always the SDK type.
func NewCheckoutTransportError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(CheckoutTransportError); ok {
		return err
	}
	return CheckoutTransportError{Err: err}
}
