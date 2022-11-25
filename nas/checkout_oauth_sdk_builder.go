package nas

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
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

func (b *CheckoutOAuthSdkBuilder) WithEnvironment(environment configuration.Environment) *CheckoutOAuthSdkBuilder {
	b.Environment = environment
	return b
}

func (b *CheckoutOAuthSdkBuilder) WithHttpClient(client *http.Client) *CheckoutOAuthSdkBuilder {
	b.HttpClient = client
	return b
}

func (b *CheckoutOAuthSdkBuilder) Build() (*Api, error) {
	if b.ClientId == "" || b.ClientSecret == "" {
		return nil, errors.CheckoutArgumentError("Invalid OAuth 'client_id' or 'client_secret'")
	}

	if b.AuthorizationUri == "" {
		b.AuthorizationUri = b.SdkBuilder.Environment.AuthorizationUri()
	}

	sdkCredentials, err := configuration.NewOAuthSdkCredentials(
		b.ClientId,
		b.ClientSecret,
		b.AuthorizationUri,
		b.Scopes)
	if err != nil {
		return nil, err
	}

	newConfiguration := configuration.NewConfiguration(sdkCredentials, b.Environment, b.HttpClient)

	return CheckoutApi(newConfiguration), nil
}
