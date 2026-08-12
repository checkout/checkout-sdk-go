package configuration

import (
	"net/url"
	"regexp"

	"github.com/checkout/checkout-sdk-go/v2/errors"
)

type Environment interface {
	BaseUri() string
	AuthorizationUri() string
	FilesUri() string
	TransfersUri() string
	BalancesUri() string
	ForwardUri() string
	IdentityUri() string
	IsSandbox() bool
}

type EnvironmentSubdomain struct {
	ApiUrl           string
	AuthorizationUrl string
}

func NewEnvironmentSubdomain(environment Environment, subdomain string) (*EnvironmentSubdomain, error) {
	apiUrl, err := createUrlWithSubdomain(environment.BaseUri(), subdomain)
	if err != nil {
		return nil, err
	}
	authorizationUrl, err := createUrlWithSubdomain(environment.AuthorizationUri(), subdomain)
	if err != nil {
		return nil, err
	}
	return &EnvironmentSubdomain{
		ApiUrl:           apiUrl,
		AuthorizationUrl: authorizationUrl,
	}, nil
}

// createUrlWithSubdomain applies subdomain transformation to any given URI, prepending the
// subdomain to the host. It returns an error when the subdomain is not a valid merchant-specific
// subdomain, rather than quietly falling back to the shared host.
func createUrlWithSubdomain(originalUrl, subdomain string) (string, error) {
	regex := regexp.MustCompile("^(?:pl-)?[a-z0-9]+$")

	if !regex.MatchString(subdomain) {
		return "", errors.CheckoutArgumentError(
			"invalid environment subdomain - provide your merchant-specific subdomain, the first " +
				"8 characters of your client ID (see " +
				"https://api-reference.checkout.com/#section/Base-URLs)")
	}

	merchantUrl, err := url.Parse(originalUrl)
	if err != nil {
		return "", err
	}
	merchantUrl.Host = subdomain + "." + merchantUrl.Host
	return merchantUrl.String(), nil
}

type CheckoutEnv struct {
	baseUri          string
	authorizationUri string
	filesUri         string
	transfersUri     string
	balancesUri      string
	forwardUri       string
	identityUri      string
	isSandbox        bool
}

func (e *CheckoutEnv) BaseUri() string {
	return e.baseUri
}

func (e *CheckoutEnv) AuthorizationUri() string {
	return e.authorizationUri
}

func (e *CheckoutEnv) FilesUri() string {
	return e.filesUri
}

func (e *CheckoutEnv) TransfersUri() string {
	return e.transfersUri
}

func (e *CheckoutEnv) BalancesUri() string {
	return e.balancesUri
}

func (e *CheckoutEnv) ForwardUri() string {
	return e.forwardUri
}

func (e *CheckoutEnv) IdentityUri() string {
	return e.identityUri
}

func (e *CheckoutEnv) IsSandbox() bool {
	return e.isSandbox
}

func NewEnvironment(
	baseUri string,
	authorizationUri string,
	filesUri string,
	transfersUri string,
	balancesUri string,
	forwardUri string,
	identityUri string,
	isSandbox bool,
) *CheckoutEnv {
	return &CheckoutEnv{
		baseUri:          baseUri,
		authorizationUri: authorizationUri,
		filesUri:         filesUri,
		transfersUri:     transfersUri,
		balancesUri:      balancesUri,
		forwardUri:       forwardUri,
		identityUri:      identityUri,
		isSandbox:        isSandbox}
}

func Sandbox() *CheckoutEnv {
	return NewEnvironment(
		"https://api.sandbox.checkout.com",
		"https://access.sandbox.checkout.com/connect/token",
		"https://files.sandbox.checkout.com",
		"https://transfers.sandbox.checkout.com",
		"https://balances.sandbox.checkout.com",
		"https://forward.sandbox.checkout.com",
		"https://identity-verification.sandbox.checkout.com",
		true)
}

func Production() *CheckoutEnv {
	return NewEnvironment(
		"https://api.checkout.com",
		"https://access.checkout.com/connect/token",
		"https://files.checkout.com/",
		"https://transfers.checkout.com/",
		"https://balances.checkout.com/",
		"https://forward.checkout.com",
		"https://identity-verification.checkout.com",
		false)
}
