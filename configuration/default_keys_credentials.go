package configuration

import "github.com/checkout/checkout-sdk-go-beta/errors"

type DefaultKeysSdkCredentials struct {
	StaticKeys
}

func NewDefaultKeysSdkCredentials(secretKey string, publicKey string) *DefaultKeysSdkCredentials {
	return &DefaultKeysSdkCredentials{StaticKeys{
		SecretKey: secretKey,
		PublicKey: publicKey,
	}}
}

func (f *DefaultKeysSdkCredentials) GetAuthorization(authorizationType AuthorizationType) (*SdkAuthorization, error) {
	switch authorizationType {
	case SecretKey, SecretKeyOrOauth:
		return &SdkAuthorization{
			PlatformType: Default,
			Credential:   f.SecretKey,
		}, nil
	case PublicKey, PublicKeyOrOauth:
		return &SdkAuthorization{
			PlatformType: Default,
			Credential:   f.PublicKey,
		}, nil
	default:
		return nil, errors.CheckoutAuthorizationError("Invalid authorization type")
	}
}
