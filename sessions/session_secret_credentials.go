package sessions

import (
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
)

type SessionSecretCredentials struct {
	SessionSecret string
}

func NewSessionSecretCredentials(sessionSecret string) *SessionSecretCredentials {
	return &SessionSecretCredentials{SessionSecret: sessionSecret}
}

func (c *SessionSecretCredentials) GetAuthorization(authorizationType configuration.AuthorizationType) (*configuration.SdkAuthorization, error) {
	switch authorizationType {
	case configuration.CustomAuth:
		if c.SessionSecret != "" {
			return &configuration.SdkAuthorization{
				PlatformType: configuration.Custom,
				Credential:   c.SessionSecret,
			}, nil
		}
		return nil, errors.InvalidKey("session_secret")
	default:
		return nil, errors.InvalidAuthorizationType(string(configuration.CustomAuth))
	}
}
