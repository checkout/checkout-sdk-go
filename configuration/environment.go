package configuration

import (
	"net/url"
	"regexp"
)

type Environment interface {
	BaseUri() string
	AuthorizationUri() string
	FilesUri() string
	TransfersUri() string
	BalancesUri() string
	IsSandbox() bool
}

type EnvironmentSubdomain struct {
	ApiUrl           string
	AuthorizationUrl string
}

func NewEnvironmentSubdomain(environment Environment, subdomain string) *EnvironmentSubdomain {
	apiUrl := createUrlWithSubdomain(environment.BaseUri(), subdomain)
	authorizationUrl := createUrlWithSubdomain(environment.AuthorizationUri(), subdomain)
	return &EnvironmentSubdomain{
		ApiUrl:           apiUrl,
		AuthorizationUrl: authorizationUrl,
	}
}

// createUrlWithSubdomain applies subdomain transformation to any given URI.
// If the subdomain is valid (alphanumeric pattern), prepends it to the host.
// Otherwise, returns the original URI unchanged.
func createUrlWithSubdomain(originalUrl string, subdomain string) string {
	newEnvironment := originalUrl

	regex := regexp.MustCompile("^[0-9a-z]+$")

	if regex.MatchString(subdomain) {
		merchantUrl, err := url.Parse(originalUrl)
		if err != nil {
			return newEnvironment
		}
		merchantUrl.Host = subdomain + "." + merchantUrl.Host
		newEnvironment = merchantUrl.String()
	}

	return newEnvironment
}

type CheckoutEnv struct {
	baseUri          string
	authorizationUri string
	filesUri         string
	transfersUri     string
	balancesUri      string
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

func (e *CheckoutEnv) IsSandbox() bool {
	return e.isSandbox
}

func NewEnvironment(
	baseUri string,
	authorizationUri string,
	filesUri string,
	transfersUri string,
	balancesUri string,
	isSandbox bool,
) *CheckoutEnv {
	return &CheckoutEnv{
		baseUri:          baseUri,
		authorizationUri: authorizationUri,
		filesUri:         filesUri,
		transfersUri:     transfersUri,
		balancesUri:      balancesUri,
		isSandbox:        isSandbox}
}

func Sandbox() *CheckoutEnv {
	return NewEnvironment("https://api.sandbox.checkout.com",
		"https://access.sandbox.checkout.com/connect/token",
		"https://files.sandbox.checkout.com",
		"https://transfers.sandbox.checkout.com",
		"https://balances.sandbox.checkout.com",
		true)
}

func Production() *CheckoutEnv {
	return NewEnvironment(
		"https://api.checkout.com",
		"https://access.checkout.com/connect/token",
		"https://files.checkout.com/",
		"https://transfers.checkout.com/",
		"https://balances.checkout.com/",
		false)
}
