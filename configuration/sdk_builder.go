package configuration

import "net/http"

type SdkBuilder struct {
	EnableTelemetry      *bool
	Environment          Environment
	EnvironmentSubdomain *EnvironmentSubdomain
	HttpClient           *http.Client
	Logger               StdLogger
}

func (s *SdkBuilder) GetConfiguration(string, string) *Configuration {
	return new(Configuration)
}
