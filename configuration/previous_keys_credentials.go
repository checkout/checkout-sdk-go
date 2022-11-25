package configuration

import "github.com/checkout/checkout-sdk-go/errors"

type PreviousKeysSdkCredentials struct {
	StaticKeys
}

func NewPreviousKeysSdkCredentials(secretKey string, publicKey string) *PreviousKeysSdkCredentials {
	return &PreviousKeysSdkCredentials{StaticKeys{
		SecretKey: secretKey,
		PublicKey: publicKey,
	}}
}

func (c *PreviousKeysSdkCredentials) GetAuthorization(authorizationType AuthorizationType) (*SdkAuthorization, error) {
	switch authorizationType {
	case SecretKey, SecretKeyOrOauth:
		return &SdkAuthorization{
			PlatformType: Previous,
			Credential:   c.SecretKey,
		}, nil
	case PublicKey, PublicKeyOrOauth:
		return &SdkAuthorization{
			PlatformType: Previous,
			Credential:   c.PublicKey,
		}, nil
	default:
		return nil, errors.CheckoutAuthorizationError("Invalid authorization type")
	}
}
