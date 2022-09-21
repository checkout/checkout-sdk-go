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
)

// ClientVersion ...
const ClientVersion = "0.0.1"

const (
    sandboxURI    = "https://api.sandbox.checkout.com"
    productionURI = "https://api.checkout.com"

    sandboxTransfersURI    = "https://transfers.sandbox.checkout.com"
    productionTransfersURI = "https://transfers.checkout.com"
)

const (
    // UPAPI - Unified Payment API
    UPAPI SupportedAPI = "api"
    // Access - OAuth Authorization
    Access SupportedAPI = "access"
)

var mbcLiveSecretKeyPattern = regexp.MustCompile(common.LiveSecretKeyRegex)
var fourKeyPattern = regexp.MustCompile(common.FourKeyRegex)
var fourOAuthJwtPattern = regexp.MustCompile(common.FourOAuthJwtPattern)

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
    PublicKey            string
    SecretKey            string
    URI                  *string
    TransferURI          *string
    HTTPClient           *http.Client
    LeveledLogger        LeveledLoggerInterface
    MaxNetworkRetries    *int64
    BearerAuthentication bool
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

// Deprecated: Please use SdkConfig
func Create(secretKey string, publicKey *string) (*Config, error) {

    var config, isSandbox = create(secretKey)

    if config.HTTPClient == nil {
        config.HTTPClient = httpClient
    }

    if config.LeveledLogger == nil {
        config.LeveledLogger = DefaultLeveledLogger
    }

    if config.MaxNetworkRetries == nil {
        config.MaxNetworkRetries = Int64(DefaultMaxNetworkRetries)
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

func SdkConfig(secretKey *string, publicKey *string, env SupportedEnvironment) (*Config, error) {

    var config Config
    config.SecretKey = StringValue(secretKey)
    config.PublicKey = StringValue(publicKey)
    config.BearerAuthentication = shouldApplyBearer(config)

    if env == Sandbox {
        config.URI = String(sandboxURI)
        config.TransferURI = String(sandboxTransfersURI)
    } else if env == Production {
        config.URI = String(productionURI)
        config.TransferURI = String(productionTransfersURI)
    }

    if config.HTTPClient == nil {
        config.HTTPClient = httpClient
    }
    if config.LeveledLogger == nil {
        config.LeveledLogger = DefaultLeveledLogger
    }
    if config.MaxNetworkRetries == nil {
        config.MaxNetworkRetries = Int64(DefaultMaxNetworkRetries)
    }

    return &config, nil
}

func shouldApplyBearer(config Config) bool {
    // SecretKey or PublicKey matches a Four pattern
    if fourKeyPattern.MatchString(config.SecretKey) || fourKeyPattern.MatchString(config.PublicKey) {
        return true
    }
    // SecretKey or PublicKey matches a JWT
    if fourOAuthJwtPattern.MatchString(config.SecretKey) || fourOAuthJwtPattern.MatchString(config.PublicKey) {
        return true
    }
    return false
}

func create(secretKey string) (Config, bool) {

    if mbcLiveSecretKeyPattern.MatchString(secretKey) {
        return Config{
            URI:       String(productionURI),
            SecretKey: secretKey,
        }, false
    }
    return Config{
        URI:       String(sandboxURI),
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

// Int64 returns a pointer to the int64 value passed in.
func Int64(v int64) *int64 {
    return &v
}

// Int64Value returns the value of the int64 pointer passed in or
// 0 if the pointer is nil.
func Int64Value(v *int64) int64 {
    if v != nil {
        return *v
    }
    return 0
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
