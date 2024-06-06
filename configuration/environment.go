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
	ApiUrl string
}

func NewEnvironmentSubdomain(environment Environment, subdomain string) *EnvironmentSubdomain {
	apiUrl := addSubdomainToApiUrlEnvironment(environment, subdomain)
	return &EnvironmentSubdomain{ApiUrl: apiUrl}
}

func addSubdomainToApiUrlEnvironment(environment Environment, subdomain string) string {
	apiUrl := environment.BaseUri()

	newEnvironment := apiUrl

	regex := regexp.MustCompile("^[0-9a-z]{8,11}$")

	if regex.MatchString(subdomain) {
		merchantApiUrl, _ := url.Parse(apiUrl)
		merchantApiUrl.Host = subdomain + "." + merchantApiUrl.Host

		newEnvironment = merchantApiUrl.String()
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
