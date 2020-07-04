package checkout

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

// APIResponse ...
type APIResponse struct {
	Status       string `json:"status,omitempty"`
	StatusCode   int    `json:"status_code,omitempty"`
	ResponseBody []byte `json:"response_body,omitempty"`
}

// Link ...git
type Link struct {
	HRef string `json:"href,omitempty"`
}

// HTTPClient ...
type HTTPClient interface {
	Get(param string) (*APIResponse, error)
	Post(param string, request interface{}) (*APIResponse, error)
}
