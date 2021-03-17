package checkout

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"runtime"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/internal/utils"
)

// ClientVersion ...
const ClientVersion = "0.0.1"

const (
	sandboxURI    = "https://api.sandbox.checkout.com"
	productionURI = "https://api.checkout.com"
)

const (
	// UPAPI - Unified Payment API
	UPAPI SupportedAPI = "api"
	// Access - OAuth Authorization
	Access SupportedAPI = "access"
)

const (
	// Sandbox - Sandbox
	Sandbox SupportedEnvironment = "sandbox.checkout.com"
	// Production - Production
	Production SupportedEnvironment = "checkout.com"
	// UnknownPlatform - Production
	UnknownPlatform string = "unknown platform"
)

const (
	// DefaultMaxNetworkRetries is the default maximum number of retries made
	// by a Checkout.com client.
	DefaultMaxNetworkRetries int64 = 2
)

const (
	// CKORequestID ...
	CKORequestID = "cko-request-id"
	// CKOVersion ...
	CKOVersion = "cko-version"
)

// SupportedAPI is an enumeration of supported Checkout.com endpoints.
// Currently supported values are "Unified Payment Gateway".
type SupportedAPI string

// SupportedEnvironment is an enumeration of supported Checkout.com environment.
// Currently supported values are "Sandbox" & "Production".
type SupportedEnvironment string

// Config ...
type Config struct {
	PublicKey         string
	SecretKey         string
	URI               *string
	HTTPClient        *http.Client
	LeveledLogger     LeveledLoggerInterface
	MaxNetworkRetries *int64
}

// DefaultConfig ...
var DefaultConfig = Config{
	URI: utils.String(sandboxURI),
}

const (
	defaultHTTPTimeout = 30 * time.Second
)

var httpClient = &http.Client{
	Timeout: defaultHTTPTimeout,
	Transport: &http.Transport{
		TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
	},
}

// Create ...
func Create(secretKey string, publicKey *string) (*Config, error) {

	var config, isSandbox = create(secretKey)

	if config.HTTPClient == nil {
		config.HTTPClient = httpClient
	}

	if config.LeveledLogger == nil {
		config.LeveledLogger = DefaultLeveledLogger
	}

	if config.MaxNetworkRetries == nil {
		config.MaxNetworkRetries = utils.Int64(DefaultMaxNetworkRetries)
	}

	if publicKey == nil {
		return &config, nil
	}

	if !isSandbox {
		publicKeyMatch := regexp.MustCompile(common.LivePublicKeyRegex)
		if publicKeyMatch.MatchString(utils.StringValue(publicKey)) {
			config.PublicKey = utils.StringValue(publicKey)
			return &config, nil
		}
		return nil, &common.Error{
			Status: "Configuration Error - Please review your secret key and public key ",
		}
	}
	publicKeyMatch := regexp.MustCompile(common.SandboxPublicKeyRegex)
	if publicKeyMatch.MatchString(utils.StringValue(publicKey)) {
		config.PublicKey = utils.StringValue(publicKey)
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
			URI:       utils.String(productionURI),
			SecretKey: secretKey,
		}, false
	}
	return Config{
		URI:       utils.String(sandboxURI),
		SecretKey: secretKey,
	}, true
}

var appInfo *AppInfo
var encodedCheckoutUserAgent string
var encodedUserAgent string

// AppInfo ...
type AppInfo struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Version string `json:"version"`
}

// SetAppInfo sets app information. See AppInfo.
func SetAppInfo(info *AppInfo) {
	if info != nil && info.Name == "" {
		panic(fmt.Errorf("App info name cannot be empty"))
	}
	appInfo = info

	// This is run in init, but we need to reinitialize it now that we have
	// some app info.
	initUserAgent()
}

func (a *AppInfo) formatUserAgent() string {
	str := a.Name
	if a.Version != "" {
		str += "/" + a.Version
	}
	if a.URL != "" {
		str += " (" + a.URL + ")"
	}
	return str
}

type checkoutClientUserAgent struct {
	Application     *AppInfo `json:"application"`
	BindingsVersion string   `json:"bindings_version"`
	Language        string   `json:"lang"`
	LanguageVersion string   `json:"lang_version"`
	Publisher       string   `json:"publisher"`
	Uname           string   `json:"uname"`
}

func initUserAgent() {

	encodedUserAgent = "Checkout/v1 GoBindings/" + ClientVersion
	if appInfo != nil {
		encodedUserAgent += " " + appInfo.formatUserAgent()
	}

	checkoutUserAgent := &checkoutClientUserAgent{
		Application:     appInfo,
		BindingsVersion: ClientVersion,
		Language:        "go",
		LanguageVersion: runtime.Version(),
		Publisher:       "checkout.com",
		Uname:           getUname(),
	}
	marshaled, err := json.Marshal(checkoutUserAgent)
	// Encoding this struct should never be a problem, so we're okay to panic
	// in case it is for some reason.
	if err != nil {
		panic(err)
	}
	encodedCheckoutUserAgent = string(marshaled)
}

func getUname() string {
	path, err := exec.LookPath("uname")
	if err != nil {
		return UnknownPlatform
	}

	cmd := exec.Command(path, "-a")
	var out bytes.Buffer
	cmd.Stderr = nil // goes to os.DevNull
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return UnknownPlatform
	}

	return out.String()
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
	Header       http.Header
	CKORequestID *string `json:"cko-request-id,omitempty"`
	CKOVersion   *string `json:"cko-version,omitempty"`
}

// HTTPClient ...
type HTTPClient interface {
	Get(path string) (*StatusResponse, error)
	Post(path string, request interface{}, params *Params) (*StatusResponse, error)
	Put(path string, request interface{}) (*StatusResponse, error)
	Patch(path string, request interface{}) (*StatusResponse, error)
	Delete(path string) (*StatusResponse, error)
	Upload(path, boundary string, body *bytes.Buffer) (*StatusResponse, error)
	Download(path string) (*StatusResponse, error)
}
