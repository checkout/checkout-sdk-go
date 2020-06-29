package checkout

// ClientVersion ...
const ClientVersion = "0.0.1"

const (
	sandboxURI    = "https://sandbox.checkout.com"
	productionURI = "https://api.checkout.com"
)

// DefaultConfig ...
var DefaultConfig = Config{
	URI: sandboxURI,
}

// APIResponse ...
type APIResponse struct {
	Status       string
	StatusCode   int
	ResponseBody []byte
}

// Link ...
type Link struct {
	HRef  string
	Title string
}

// HTTPClient ...
type HTTPClient interface {
	Get(param string) (*APIResponse, error)
	Post(param string, request interface{}) (*APIResponse, error)
}
