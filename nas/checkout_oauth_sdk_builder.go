package nas

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

// CheckoutOAuthSdkBuilder builds an Api authorized with OAuth client credentials.
type CheckoutOAuthSdkBuilder struct {
	configuration.SdkBuilder
	// ClientId is the OAuth client ID.
	ClientId string
	// ClientSecret is the OAuth client secret.
	ClientSecret string
	// AuthorizationUri is the token endpoint; when empty it is derived from the environment or subdomain.
	AuthorizationUri string
	// Scopes are the OAuth scopes requested with the token.
	Scopes []string
}

// WithClientCredentials sets the OAuth client ID and secret. It returns the builder.
func (b *CheckoutOAuthSdkBuilder) WithClientCredentials(id string, secret string) *CheckoutOAuthSdkBuilder {
	b.ClientId = id
	b.ClientSecret = secret
	return b
}

// WithAuthorizationUri sets a custom token endpoint. It returns the builder.
func (b *CheckoutOAuthSdkBuilder) WithAuthorizationUri(uri string) *CheckoutOAuthSdkBuilder {
	b.AuthorizationUri = uri
	return b
}

// WithScopes sets the OAuth scopes requested with the token. It returns the builder.
func (b *CheckoutOAuthSdkBuilder) WithScopes(scopes []string) *CheckoutOAuthSdkBuilder {
	b.Scopes = scopes
	return b
}

// WithEnableTelemetry turns the SDK telemetry header on or off. It returns the builder.
func (b *CheckoutOAuthSdkBuilder) WithEnableTelemetry(telemetry bool) *CheckoutOAuthSdkBuilder {
	b.EnableTelemetry = &telemetry
	return b
}

// WithEnvironment sets the sandbox or production environment. It returns the builder.
func (b *CheckoutOAuthSdkBuilder) WithEnvironment(environment configuration.Environment) *CheckoutOAuthSdkBuilder {
	b.Environment = environment
	return b
}

// WithEnvironmentSubdomain sets the merchant-specific subdomain. It returns the builder.
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

// WithHttpClient sets the client for every request the SDK sends, the token request included.
// A client without Timeout gets a 5 second limit on the token request. When nil or not called,
// API requests use a client with a 10 second timeout and the token request one with 5 seconds.
// It returns the builder.
func (b *CheckoutOAuthSdkBuilder) WithHttpClient(client *http.Client) *CheckoutOAuthSdkBuilder {
	b.HttpClient = client
	return b
}

// WithLogger sets the logger. It returns the builder.
func (b *CheckoutOAuthSdkBuilder) WithLogger(logger configuration.StdLogger) *CheckoutOAuthSdkBuilder {
	b.Logger = logger
	return b
}

// Build validates the settings, requests the first access token and returns the Api. It returns an
// error when a setting is missing or invalid, or when the token request fails.
func (b *CheckoutOAuthSdkBuilder) Build() (*Api, error) {
	if b.ClientId == "" || b.ClientSecret == "" {
		return nil, errors.CheckoutArgumentError("Invalid OAuth 'client_id' or 'client_secret'")
	}

	if err := b.ValidateEnvironmentSettings(true); err != nil {
		return nil, err
	}

	if b.AuthorizationUri != "" && b.Subdomain != "" {
		return nil, errors.CheckoutArgumentError(
			"authorization URI and environment subdomain cannot both be set - the token endpoint " +
				"is derived from your subdomain; combine the authorization URI with " +
				"WithLegacyDomain() if you need a custom token host")
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

	sdkCredentials, err := configuration.NewOAuthSdkCredentialsWithHttpClient(
		b.ClientId,
		b.ClientSecret,
		b.AuthorizationUri,
		b.Scopes,
		b.HttpClient,
		b.Logger)
	if err != nil {
		return nil, err
	}

	newConfiguration := configuration.NewConfiguration(sdkCredentials, b.EnableTelemetry, b.Environment, b.HttpClient, b.Logger)

	if environmentSubdomain != nil {
		newConfiguration = configuration.NewConfigurationWithSubdomain(sdkCredentials, b.EnableTelemetry, b.Environment, environmentSubdomain, b.HttpClient, b.Logger)
	}

	return CheckoutApi(newConfiguration), nil
}
