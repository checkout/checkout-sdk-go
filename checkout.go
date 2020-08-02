package checkout

import "bytes"

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
func Create(secretKey string, useSandbox bool, publicKey *string) Config {

	var config = create(secretKey, useSandbox)
	if publicKey != nil {
		config.PublicKey = *publicKey
	}
	return config
}

func create(secretKey string, useSandbox bool) Config {
	if useSandbox {
		return Config{
			URI:       sandboxURI,
			SecretKey: secretKey,
		}
	}
	return Config{
		URI:       productionURI,
		SecretKey: secretKey,
	}
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
	CKORequestID string `json:"cko-request-id,omitempty"`
	CKOVersion   string `json:"cko-version,omitempty"`
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

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}
