package nas

import (
	"net/http"

	"github.com/checkout/checkout-sdk-go/configuration"
)

type CheckoutDefaultSdkBuilder struct {
	configuration.StaticKeysBuilder
}

func (b *CheckoutDefaultSdkBuilder) WithEnvironment(environment configuration.Environment) *CheckoutDefaultSdkBuilder {
	b.Environment = environment
	return b
}

func (b *CheckoutDefaultSdkBuilder) WithHttpClient(client *http.Client) *CheckoutDefaultSdkBuilder {
	b.HttpClient = client
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
	err := b.ValidateSecretKey(configuration.DefaultSecretKeyPattern)
	if err != nil {
		return nil, err
	}

	err = b.ValidatePublicKey(configuration.DefaultPublicKeyPattern)
	if err != nil {
		return nil, err
	}

	sdkCredentials := configuration.NewDefaultKeysSdkCredentials(b.SecretKey, b.PublicKey)
	newConfiguration := configuration.NewConfiguration(sdkCredentials, b.Environment, b.HttpClient)

	return CheckoutApi(newConfiguration), nil
}
