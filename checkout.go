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
	Status       string `json:"status"`
	StatusCode   int    `json:"status_code"`
	ResponseBody []byte `json:"response_body"`
}

// Link ...
type Link struct {
	HRef string `json:"href"`
}

// HTTPClient ...
type HTTPClient interface {
	Get(param string) (*APIResponse, error)
	Post(param string, request interface{}) (*APIResponse, error)
}
