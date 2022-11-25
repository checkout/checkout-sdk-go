package sessions

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
)

func TestGetAuthorization(t *testing.T) {
	cases := []struct {
		name              string
		sessionSecret     string
		authorizationType configuration.AuthorizationType
		checker           func(*configuration.SdkAuthorization, error)
	}{
		{
			name:              "when session secret is empty then return error",
			authorizationType: configuration.CustomAuth,
			checker: func(authorization *configuration.SdkAuthorization, err error) {
				assert.Nil(t, authorization)
				assert.NotNil(t, err)
				assert.IsType(t, reflect.TypeOf(errors.CheckoutAuthorizationError("")), reflect.TypeOf(err))
				assert.Equal(t, "session_secret is required for this operation", err.Error())
			},
		},
		{
			name:              "when authorization type is invalid then return error",
			sessionSecret:     "secret",
			authorizationType: configuration.OAuth,
			checker: func(authorization *configuration.SdkAuthorization, err error) {
				assert.Nil(t, authorization)
				assert.NotNil(t, err)
				assert.IsType(t, reflect.TypeOf(errors.CheckoutAuthorizationError("")), reflect.TypeOf(err))
				assert.Equal(t, fmt.Sprintf("Operation requires %s authorization type", configuration.CustomAuth), err.Error())
			},
		},
		{
			name:              "when session secret and authorization type is valid then SDK Authorization",
			sessionSecret:     "secret",
			authorizationType: configuration.CustomAuth,
			checker: func(authorization *configuration.SdkAuthorization, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, authorization)
				assert.Equal(t, configuration.Custom, authorization.PlatformType)
				assert.Equal(t, "secret", authorization.Credential)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(NewSessionSecretCredentials(tc.sessionSecret).GetAuthorization(tc.authorizationType))
		})
	}
}
