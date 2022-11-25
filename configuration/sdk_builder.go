package configuration

import "net/http"

type SdkBuilder struct {
	Environment Environment
	HttpClient  *http.Client
}

func (s *SdkBuilder) GetConfiguration(string, string) *Configuration {
	return new(Configuration)
}
