package nas

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
)

type CheckoutOAuthSdkBuilder struct {
	configuration.SdkBuilder
	ClientId         string
	ClientSecret     string
	AuthorizationUri string
	Scopes           []string
}

func (b *CheckoutOAuthSdkBuilder) WithClientCredentials(id string, secret string) *CheckoutOAuthSdkBuilder {
	b.ClientId = id
	b.ClientSecret = secret
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithAuthorizationUri(uri string) *CheckoutOAuthSdkBuilder {
	b.AuthorizationUri = uri
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithScopes(scopes []string) *CheckoutOAuthSdkBuilder {
	b.Scopes = scopes
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithEnableTelemetry(telemetry bool) *CheckoutOAuthSdkBuilder {
	b.EnableTelemetry = &telemetry
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithEnvironment(environment configuration.Environment) *CheckoutOAuthSdkBuilder {
	b.Environment = environment
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithEnvironmentSubdomain(subdomain string) *CheckoutOAuthSdkBuilder {
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
func (b *CheckoutOAuthSdkBuilder) WithLegacyDomain() *CheckoutOAuthSdkBuilder {
	b.UseLegacyDomain = true
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithHttpClient(client *http.Client) *CheckoutOAuthSdkBuilder {
	b.HttpClient = client
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithLogger(logger configuration.StdLogger) *CheckoutOAuthSdkBuilder {
	b.Logger = logger
	return b
}

func (b *CheckoutOAuthSdkBuilder) Build() (*Api, error) {
	if b.ClientId == "" || b.ClientSecret == "" {
		return nil, errors.CheckoutArgumentError("Invalid OAuth 'client_id' or 'client_secret'")
	}

	if err := b.ValidateEnvironmentSettings(true); err != nil {
		return nil, err
	}

	environmentSubdomain, subdomainErr := b.GetEnvironmentSubdomain()
	if subdomainErr != nil {
		return nil, subdomainErr
	}

	if b.AuthorizationUri == "" {
		if environmentSubdomain != nil {
			b.AuthorizationUri = environmentSubdomain.AuthorizationUrl
		} else {
			b.AuthorizationUri = b.SdkBuilder.Environment.AuthorizationUri()
		}
	}

	sdkCredentials, err := configuration.NewOAuthSdkCredentials(
		b.ClientId,
		b.ClientSecret,
		b.AuthorizationUri,
		b.Scopes,
		b.Logger)
	if err != nil {
		return nil, err
	}

	newConfiguration := configuration.NewConfiguration(sdkCredentials, b.EnableTelemetry, b.Environment, b.HttpClient, b.Logger)

	if environmentSubdomain != nil {
		newConfiguration = configuration.NewConfigurationWithSubdomain(sdkCredentials, b.Environment, environmentSubdomain, b.HttpClient, b.Logger)
	}

	return CheckoutApi(newConfiguration), nil
}
