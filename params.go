package checkout

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// Params is the structure that contains the common properties
// of any *Params structure.
type Params struct {
	// Headers may be used to provide extra header lines on the HTTP request.
	Headers        http.Header `form:"-"`
	IdempotencyKey *string     `form:"-"` // Passed as header
}

// GetParams -
func (p *Params) GetParams() *Params {
	return p
}

// SetIdempotencyKey sets a value for the Idempotency-Key header.
func (p *Params) SetIdempotencyKey(val string) {
	p.IdempotencyKey = &val
}

// ParamsContainer is a general interface for which all parameter structs
// should comply. They achieve this by embedding a Params struct and inheriting
// its implementation of this interface.
type ParamsContainer interface {
	GetParams() *Params
}

// NewIdempotencyKey -
func NewIdempotencyKey() string {
	now := time.Now().UnixNano()
	buf := make([]byte, 4)
	rand.Read(buf)
	return fmt.Sprintf("%v_%v", now, base64.URLEncoding.EncodeToString(buf)[:6])
}
