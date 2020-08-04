package checkout

import (
	"bytes"
	"fmt"
	"regexp"

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
func Create(secretKey string, useSandbox bool, publicKey *string, idempotencyKey *string) (*Config, error) {

	var config = create(secretKey, useSandbox)
	if config != nil {
		if idempotencyKey != nil {
			config.IdempotencyKey = idempotencyKey
		}
		if publicKey != nil {
			if useSandbox {
				var publicKeyMatch = regexp.MustCompile(common.SandboxPublicKeyRegex)
				fmt.Println("SandboxPublicKeyRegex: ", common.SandboxPublicKeyRegex)
				if publicKeyMatch.MatchString(StringValue(publicKey)) {
					config.PublicKey = StringValue(publicKey)
					return config, nil
				}
			} else {
				var publicKeyMatch = regexp.MustCompile(common.LivePublicKeyRegex)
				if publicKeyMatch.MatchString(StringValue(publicKey)) {
					config.PublicKey = StringValue(publicKey)
					return config, nil
				}
			}
			return config, &common.Error{
				Status: "Configuration Error - Please review your secret key and public key ",
			}
		}
		return config, &common.Error{
			Status: "Configuration Error - Please review your secret key and public key ",
		}
	}
	return config, &common.Error{
		Status: "Configuration Error - Please review your secret key and public key ",
	}
}

func create(secretKey string, useSandbox bool) *Config {
	if useSandbox {
		var secretKeyMatch = regexp.MustCompile(common.SandboxSecretKeyRegex)
		fmt.Println("SandboxSecretKeyRegex: ", common.SandboxSecretKeyRegex)
		if secretKeyMatch.MatchString(secretKey) {
			return &Config{
				URI:       sandboxURI,
				SecretKey: secretKey,
			}
		}
		return nil
	}
	var secretKeyMatch = regexp.MustCompile(common.LiveSecretKeyRegex)
	if secretKeyMatch.MatchString(secretKey) {
		return &Config{
			URI:       productionURI,
			SecretKey: secretKey,
		}
	}
	return nil
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
