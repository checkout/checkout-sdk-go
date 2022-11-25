package errors

import "fmt"

type ErrorDetails struct {
	RequestID  string   `json:"request_id,omitempty"`
	ErrorType  string   `json:"error_type,omitempty"`
	ErrorCodes []string `json:"error_codes,omitempty"`
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
