package nas

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
)

type CheckoutDefaultSdkBuilder struct {
	configuration.StaticKeysBuilder
}

func (b *CheckoutDefaultSdkBuilder) WithEnableTelemetry(telemetry bool) *CheckoutDefaultSdkBuilder {
	b.EnableTelemetry = &telemetry
	return b
}

func (b *CheckoutDefaultSdkBuilder) WithEnvironment(environment configuration.Environment) *CheckoutDefaultSdkBuilder {
	b.Environment = environment
	return b
}

func (b *CheckoutDefaultSdkBuilder) WithEnvironmentSubdomain(subdomain string) *CheckoutDefaultSdkBuilder {
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
func (b *CheckoutDefaultSdkBuilder) WithLegacyDomain() *CheckoutDefaultSdkBuilder {
	b.UseLegacyDomain = true
	return b
}

func (b *CheckoutDefaultSdkBuilder) WithHttpClient(client *http.Client) *CheckoutDefaultSdkBuilder {
	b.HttpClient = client
	return b
}

func (b *CheckoutDefaultSdkBuilder) WithLogger(logger configuration.StdLogger) *CheckoutDefaultSdkBuilder {
	b.Logger = logger
	return b
}

func (b *CheckoutDefaultSdkBuilder) WithPublicKey(publicKey string) *CheckoutDefaultSdkBuilder {
	b.PublicKey = publicKey
	return b
}

func (b *CheckoutDefaultSdkBuilder) WithSecretKey(secretKey string) *CheckoutDefaultSdkBuilder {
	b.SecretKey = secretKey
	return b
}

func (b *CheckoutDefaultSdkBuilder) Build() (*Api, error) {
	if err := b.ValidateEnvironmentSettings(true); err != nil {
		return nil, err
	}

	environmentSubdomain, err := b.GetEnvironmentSubdomain()
	if err != nil {
		return nil, err
	}

	err = b.ValidateSecretKey(configuration.DefaultSecretKeyPattern)
	if err != nil {
		return nil, err
	}

	err = b.ValidatePublicKey(configuration.DefaultPublicKeyPattern)
	if err != nil {
		return nil, err
	}

	sdkCredentials := configuration.NewDefaultKeysSdkCredentials(b.SecretKey, b.PublicKey)

	newConfiguration := configuration.NewConfiguration(sdkCredentials, b.EnableTelemetry, b.Environment, b.HttpClient, b.Logger)

	if environmentSubdomain != nil {
		newConfiguration = configuration.NewConfigurationWithSubdomain(sdkCredentials, b.EnableTelemetry, b.Environment, environmentSubdomain, b.HttpClient, b.Logger)
	}

	return CheckoutApi(newConfiguration), nil
}
