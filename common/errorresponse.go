package common

// ErrorDetails ...
type ErrorDetails struct {
	RequestID  string   `json:"request_id,omitempty"`
	ErrorType  string   `json:"error_type,omitempty"`
	ErrorCodes []string `json:"error_codes,omitempty"`
}

// Error ...
type Error struct {
	Data       *ErrorDetails
	Status     string
	StatusCode int
}

func (e *Error) Error() string {
	return e.Status
}
