package configuration

import "github.com/checkout/checkout-sdk-go-beta/errors"

type PlatformType string

const (
	Previous     PlatformType = "PREVIOUS"
	Default      PlatformType = "DEFAULT"
	DefaultOAuth PlatformType = "DEFAULT_OAUTH"
	Custom       PlatformType = "CUSTOM"
)

type AuthorizationType string

const (
	PublicKey        AuthorizationType = "PUBLIC_KEY"
	SecretKey        AuthorizationType = "SECRET_KEY"
	PublicKeyOrOauth AuthorizationType = "PUBLIC_KEY_OR_OAUTH"
	SecretKeyOrOauth AuthorizationType = "SECRET_KEY_OR_OAUTH"
	OAuth            AuthorizationType = "OAUTH"
	CustomAuth       AuthorizationType = "CUSTOM"
)

type (
	SdkCredentials interface {
		GetAuthorization(authorizationType AuthorizationType) (*SdkAuthorization, error)
	}

	SdkAuthorization struct {
		PlatformType PlatformType
		Credential   string
	}
)

func (s *SdkAuthorization) GetAuthorizationHeader() (string, error) {
	switch s.PlatformType {
	case Previous, Custom:
		return s.Credential, nil
	case Default, DefaultOAuth:
		return "Bearer " + s.Credential, nil
	default:
		return "", errors.CheckoutAuthorizationError("Invalid platform type")
	}
}

type StaticKeys struct {
	SecretKey string
	PublicKey string
}
