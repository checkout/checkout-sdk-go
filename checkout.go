package checkout

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"regexp"
	"time"

	"github.com/shiuh-yaw-cko/checkout/common"
)

// ClientVersion ...
const ClientVersion = "0.0.1"

const (
	sandboxURI    = "https://api.sandbox.checkout.com"
	productionURI = "https://api.checkout.com"
)

// DefaultConfig ...
var DefaultConfig = Config{
	URI: sandboxURI,
}

// Create ...
func Create(secretKey string, publicKey *string, idempotencyKey *string) (*Config, error) {

	var config, isSandbox = create(secretKey)
	if idempotencyKey != nil {
		config.IdempotencyKey = idempotencyKey
	}
	if publicKey == nil {
		return &config, nil
	}
	if !isSandbox {
		publicKeyMatch := regexp.MustCompile(common.LivePublicKeyRegex)
		if publicKeyMatch.MatchString(StringValue(publicKey)) {
			config.PublicKey = StringValue(publicKey)
			return &config, nil
		}
		return nil, &common.Error{
			Status: "Configuration Error - Please review your secret key and public key ",
		}
	}
	publicKeyMatch := regexp.MustCompile(common.SandboxPublicKeyRegex)
	if publicKeyMatch.MatchString(StringValue(publicKey)) {
		config.PublicKey = StringValue(publicKey)
		return &config, nil
	}
	return nil, &common.Error{
		Status: "Configuration Error - Please review your secret key and public key ",
	}
}

func create(secretKey string) (Config, bool) {

	liveSecretKeyMatch := regexp.MustCompile(common.LiveSecretKeyRegex)
	if liveSecretKeyMatch.MatchString(secretKey) {
		return Config{
			URI:       productionURI,
			SecretKey: secretKey,
		}, false
	}
	return Config{
		URI:       sandboxURI,
		SecretKey: secretKey,
	}, true
}

// StatusResponse ...
type StatusResponse struct {
	Status       string     `json:"status,omitempty"`
	StatusCode   int        `json:"status_code,omitempty"`
	ResponseBody []byte     `json:"response_body,omitempty"`
	ResponseCSV  [][]string `json:"response_csv,omitempty"`
	Headers      *Headers   `json:"headers,omitempty"`
}

// Headers ...
type Headers struct {
	CKORequestID *string `json:"cko-request-id,omitempty"`
	CKOVersion   *string `json:"cko-version,omitempty"`
}

// HTTPClient ...
type HTTPClient interface {
	Get(param string) (*StatusResponse, error)
	Post(param string, request interface{}) (*StatusResponse, error)
	Put(param string, request interface{}) (*StatusResponse, error)
	Patch(param string, request interface{}) (*StatusResponse, error)
	Delete(param string) (*StatusResponse, error)
	Upload(param, boundary string, body *bytes.Buffer) (*StatusResponse, error)
	Download(path string) (*StatusResponse, error)
}

// NewIdempotencyKey -
func NewIdempotencyKey() string {
	now := time.Now().UnixNano()
	buf := make([]byte, 4)
	rand.Read(buf)
	return fmt.Sprintf("%v_%v", now, base64.URLEncoding.EncodeToString(buf)[:6])
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringSlice returns a slice of string pointers given a slice of strings.
func StringSlice(v []string) []*string {
	out := make([]*string, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}

// BoolValue returns the value of the bool pointer passed in or
// false if the pointer is nil.
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolSlice returns a slice of bool pointers given a slice of bools.
func BoolSlice(v []bool) []*bool {
	out := make([]*bool, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}

// Float64 returns a pointer to the float64 value passed in.
func Float64(v float64) *float64 {
	return &v
}

// Float64Value returns the value of the float64 pointer passed in or
// 0 if the pointer is nil.
func Float64Value(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Slice returns a slice of float64 pointers given a slice of float64s.
func Float64Slice(v []float64) []*float64 {
	out := make([]*float64, len(v))
	for i := range v {
		out[i] = &v[i]
	}
	return out
}
