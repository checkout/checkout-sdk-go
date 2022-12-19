package configuration

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/common"
)

type Configuration struct {
	Credentials SdkCredentials
	Environment Environment
	HttpClient  http.Client
	Logger      StdLogger
}

func NewConfiguration(
	credentials SdkCredentials,
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

	return &Configuration{
		Credentials: credentials,
		Environment: environment,
		HttpClient:  *client,
		Logger:      logger,
	}
}
