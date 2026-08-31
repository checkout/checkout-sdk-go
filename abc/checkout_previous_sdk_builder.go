package abc

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

type CheckoutPreviousSdkBuilder struct {
	configuration.StaticKeysBuilder
}

func (b *CheckoutPreviousSdkBuilder) WithEnableTelemetry(telemetry bool) *CheckoutPreviousSdkBuilder {
	b.EnableTelemetry = &telemetry
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithEnvironment(environment configuration.Environment) *CheckoutPreviousSdkBuilder {
	b.Environment = environment
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithEnvironmentSubdomain(subdomain string) *CheckoutPreviousSdkBuilder {
	b.Subdomain = subdomain
	return b
}

// WithLegacyDomain opts out of the merchant-specific subdomain, sending every request to the
// shared hosts instead (api.checkout.com and access.checkout.com, or their sandbox equivalents).
//
// Deprecated: this is an emergency fallback for the rare case where the merchant-specific
// subdomain cannot be used, and will be removed in a future release. Call
// WithEnvironmentSubdomain instead.
// See https://api-reference.checkout.com/#section/Base-URLs
func (b *CheckoutPreviousSdkBuilder) WithLegacyDomain() *CheckoutPreviousSdkBuilder {
	b.UseLegacyDomain = true
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithHttpClient(client *http.Client) *CheckoutPreviousSdkBuilder {
	b.HttpClient = client
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithLogger(logger configuration.StdLogger) *CheckoutPreviousSdkBuilder {
	b.Logger = logger
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithPublicKey(publicKey string) *CheckoutPreviousSdkBuilder {
	b.PublicKey = publicKey
	return b
}

func (b *CheckoutPreviousSdkBuilder) WithSecretKey(secretKey string) *CheckoutPreviousSdkBuilder {
	b.SecretKey = secretKey
	return b
}

func (b *CheckoutPreviousSdkBuilder) Build() (*Api, error) {
	// The Previous (ABC) platform predates merchant-specific subdomains, so it is exempt from
	// the mandatory WithEnvironmentSubdomain/WithLegacyDomain configuration.
	if err := b.ValidateEnvironmentSettings(false); err != nil {
		return nil, err
	}

	environmentSubdomain, err := b.GetEnvironmentSubdomain()
	if err != nil {
		return nil, err
	}

	err = b.ValidateSecretKey(configuration.PreviousSecretKeyPattern)
	if err != nil {
		return nil, err
	}

	err = b.ValidatePublicKey(configuration.PreviousPublicKeyPattern)
	if err != nil {
		return nil, err
	}

	sdkCredentials := configuration.NewPreviousKeysSdkCredentials(b.SecretKey, b.PublicKey)

	newConfiguration := configuration.NewConfiguration(sdkCredentials, b.EnableTelemetry, b.Environment, b.HttpClient, b.Logger)

	if environmentSubdomain != nil {
		newConfiguration = configuration.NewConfigurationWithSubdomain(sdkCredentials, b.EnableTelemetry, b.Environment, environmentSubdomain, b.HttpClient, b.Logger)
	}

	return CheckoutApi(newConfiguration), nil
}
