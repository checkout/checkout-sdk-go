package configuration

import (
	"regexp"

	"github.com/checkout/checkout-sdk-go/errors"
)

type StaticKeysBuilder struct {
	SdkBuilder
	PublicKey string
	SecretKey string
}

func (s *StaticKeysBuilder) ValidateSecretKey(regex string) error {
	re := regexp.MustCompile(regex)
	if !re.MatchString(s.SecretKey) {
		return errors.CheckoutArgumentError("Invalid secret key")
	}

	return nil
}

func (s *StaticKeysBuilder) ValidatePublicKey(regex string) error {
	if len(s.PublicKey) == 0 {
		return nil
	}
	re := regexp.MustCompile(regex)
	if !re.MatchString(s.PublicKey) {
		return errors.CheckoutArgumentError("Invalid public key")
	}

	return nil
}
