package configuration

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// Configuration holds what every client of an Api needs to send requests.
type Configuration struct {
	// Credentials authorize each request.
	Credentials SdkCredentials
	// EnableTelemetry sends the SDK telemetry header with each request when true.
	EnableTelemetry bool
	// Environment is the sandbox or production environment.
	Environment Environment
	// EnvironmentSubdomain is the merchant-specific subdomain, when set.
	EnvironmentSubdomain *EnvironmentSubdomain
	// HttpClient sends the API requests.
	HttpClient http.Client
	// Logger receives the request log lines.
	Logger StdLogger
}

// NewConfiguration creates the configuration of an Api. client sends the API requests; when nil, a
// client with a 10 second timeout is used. OAuth credentials send their token request through their
// own HttpClient: create them with NewOAuthSdkCredentialsWithHttpClient to use the same client.
func NewConfiguration(
	credentials SdkCredentials,
	enableTelemetry *bool,
	environment Environment,
	client *http.Client,
	logger StdLogger,
) *Configuration {
	if environment == nil {
		environment = Sandbox()
	}

	if client == nil {
		client = common.BuildDefaultClient()
	}

	if logger == nil {
		logger = DefaultLogger()
	}

	telemetry := true
	if enableTelemetry != nil {
		telemetry = *enableTelemetry
	}

	return &Configuration{
		Credentials:     credentials,
		EnableTelemetry: telemetry,
		Environment:     environment,
		HttpClient:      *client,
		Logger:          logger,
	}
}

// NewConfigurationWithSubdomain is NewConfiguration with the merchant-specific environmentSubdomain.
func NewConfigurationWithSubdomain(
	credentials SdkCredentials,
	enableTelemetry *bool,
	environment Environment,
	environmentSubdomain *EnvironmentSubdomain,
	client *http.Client,
	logger StdLogger,
) *Configuration {
	config := NewConfiguration(credentials, enableTelemetry, environment, client, logger)
	config.EnvironmentSubdomain = environmentSubdomain
	return config
}
