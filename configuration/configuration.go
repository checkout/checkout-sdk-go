package configuration

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type Configuration struct {
	Credentials          SdkCredentials
	EnableTelemetry      bool
	Environment          Environment
	EnvironmentSubdomain *EnvironmentSubdomain
	HttpClient           http.Client
	Logger               StdLogger
	// Retry is opt-in. When nil (the default) requests are executed once with
	// no retry behaviour. When set, transient failures are retried with
	// exponential backoff. Configure it via the builder's WithRetries or
	// WithRetryConfiguration methods.
	Retry *RetryConfiguration
}

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

func NewConfigurationWithSubdomain(
	credentials SdkCredentials,
	environment Environment,
	environmentSubdomain *EnvironmentSubdomain,
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

	return &Configuration{
		Credentials:          credentials,
		Environment:          environment,
		EnvironmentSubdomain: environmentSubdomain,
		HttpClient:           *client,
		Logger:               logger,
	}
}
