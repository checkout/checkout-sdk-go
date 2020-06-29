package checkout

// ErrorDetails ...
type ErrorDetails struct {
	RequestID  string   `json:"request_id,"`
	ErrorType  string   `json:"error_type"`
	ErrorCodes []string `json:"error_codes"`
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
