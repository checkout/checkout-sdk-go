package configuration

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/v2/errors"
)

type SdkBuilder struct {
	EnableTelemetry *bool
	Environment     Environment
	Subdomain       string
	UseLegacyDomain bool
	HttpClient      *http.Client
	Logger          StdLogger
}

func (s *SdkBuilder) GetConfiguration(string, string) *Configuration {
	return new(Configuration)
}

// GetEnvironmentSubdomain builds the merchant-specific subdomain URLs, or returns nil when the
// caller opted out with the legacy domain. It is resolved at build time rather than when the
// subdomain is set, so the order of the With... calls does not matter.
func (s *SdkBuilder) GetEnvironmentSubdomain() (*EnvironmentSubdomain, error) {
	if s.Subdomain == "" {
		return nil, nil
	}
	return NewEnvironmentSubdomain(s.Environment, s.Subdomain)
}

// ValidateEnvironmentSettings enforces an explicit choice of domain. The merchant-specific
// subdomain is how merchants should reach the API, so it is required unless the caller opts out
// with UseLegacyDomain. Setting both, or neither, is a configuration error.
//
// requiresSubdomain is false for the Previous (ABC) platform, which predates merchant-specific
// subdomains.
func (s *SdkBuilder) ValidateEnvironmentSettings(requiresSubdomain bool) error {
	if s.Subdomain != "" && s.UseLegacyDomain {
		return errors.CheckoutArgumentError(
			"WithEnvironmentSubdomain and WithLegacyDomain cannot both be set - provide only your " +
				"merchant-specific subdomain")
	}
	if s.Subdomain == "" && !s.UseLegacyDomain && requiresSubdomain {
		return errors.CheckoutArgumentError(
			"environment subdomain is required - provide your merchant-specific subdomain (typically " +
				"your client ID excluding the cli_ prefix, see " +
				"https://api-reference.checkout.com/#section/Base-URLs), or call WithLegacyDomain " +
				"to opt out only if merchant specific sub domains are causing issues")
	}
	return nil
}
