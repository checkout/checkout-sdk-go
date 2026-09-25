package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/checkout/checkout-sdk-go/v3"
	"github.com/checkout/checkout-sdk-go/v3/abc"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/nas"
)

const MaxRetryAttemps = 10

var previousApi *abc.Api
var defaultApi *nas.Api
var oauthApi *nas.Api

const Name = "Name"
const FirstName = "First"
const LastName = "Last"
const Email = "customer@test.checkout.com"
const City = "London"
const Reference = "Reference"
const Description = "Description"
const CardNumber = "4242424242424242"
const Iban = "DE68100100101234567895"
const Bic = "PBNKDEFFXXX"
const Cvv = "100"
const SuccessUrl = "https://test.checkout.com/success"
const FailureUrl = "https://test.checkout.com/failure"
const ExpiryYear = 2027
const ExpiryMonth = 12
const InvalidCustomerId = "cus_xxxxxxxxxxxxxxxxxxxxxxxxxx"

func newRandom() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func PreviousApi() *abc.Api {
	if previousApi == nil {
		var err error
		previousApi, err = checkout.Builder().Previous().
			WithEnvironment(configuration.Sandbox()).
			WithSecretKey(os.Getenv("CHECKOUT_PREVIOUS_SECRET_KEY")).
			WithPublicKey(os.Getenv("CHECKOUT_PREVIOUS_PUBLIC_KEY")).
			Build()
		if err != nil {
			panic(fmt.Sprintf("PreviousApi: failed to build the Previous platform client: %v", err))
		}
	}
	return previousApi
}

func DefaultApi() *nas.Api {
	if defaultApi == nil {
		var err error
		defaultApi, err = checkout.Builder().
			StaticKeys().
			WithEnvironment(configuration.Sandbox()).
			WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
			WithPublicKey(os.Getenv("CHECKOUT_DEFAULT_PUBLIC_KEY")).
			WithEnvironmentSubdomain(os.Getenv("CHECKOUT_MERCHANT_SUBDOMAIN")).
			Build()
		if err != nil {
			panic(fmt.Sprintf("DefaultApi: failed to build the static-keys client: %v", err))
		}
	}
	return defaultApi
}

func OAuthApi() *nas.Api {
	if oauthApi == nil {
		// Build() reaches the OAuth token endpoint synchronously; in CI the endpoint
		// occasionally returns HTML (e.g. proxy challenge / 5xx page) instead of a
		// token. Surface that failure here, with the real error, instead of
		// discarding it and letting every caller hit an unexplained nil-pointer
		// panic later — a swallowed error here previously showed up as a bare
		// "invalid memory address" panic in TestSetupAccountsSuite with no
		// indication of the actual OAuth failure underneath it.
		var err error
		oauthApi, err = checkout.Builder().OAuth().
			WithClientCredentials(
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_ID"),
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET")).
			WithEnvironment(configuration.Sandbox()).
			WithScopes(getOAuthScopes()).
			// The sandbox OAuth clients lack subdomain provisioning, so the token request would
			// come back invalid_client. Opting out explicitly until they are provisioned.
			WithLegacyDomain().
			Build()
		if err != nil {
			panic(fmt.Sprintf("OAuthApi: failed to build the OAuth client (token request failed): %v", err))
		}
	}
	return oauthApi
}

func Address() *common.Address {
	return &common.Address{
		AddressLine1: "Checkout.com",
		AddressLine2: "ABC build",
		City:         "London",
		State:        "London",
		Zip:          "W1T 4TJ",
		Country:      common.GB,
	}
}

func Phone() *common.Phone {
	return &common.Phone{
		CountryCode: "+1",
		Number:      "415 555 2671",
	}
}

func AccountHolder() *common.AccountHolder {
	return &common.AccountHolder{
		FirstName:      FirstName,
		LastName:       LastName,
		Phone:          Phone(),
		BillingAddress: Address(),
	}
}

type customTransport struct {
	schemaVersion string
	base          http.RoundTripper
}

func CustomHttpClient(schemaVersion string) *http.Client {
	return &http.Client{
		Transport: &customTransport{
			schemaVersion: schemaVersion,
			base:          http.DefaultTransport,
		},
	}
}

func (c *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json;schema_version="+c.schemaVersion)
	return c.base.RoundTrip(req)
}

func GenerateRandomString(length int, chars ...string) string {
	defaultChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	charSet := defaultChars
	if len(chars) > 0 {
		charSet = chars[0]
	}

	r := newRandom()
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charSet[r.Intn(len(charSet))])
	}
	return sb.String()
}

func GenerateRandomBusinessRegistrationNumber() string {
	r := newRandom()
	gbPrefixes := []string{"OC", "LP", "SC", "AC", "CE", "GS"}

	if r.Intn(2) == 0 {
		return GenerateRandomDigits(8)
	}
	return gbPrefixes[r.Intn(len(gbPrefixes))] + GenerateRandomDigits(6)
}

func GenerateRandomIdentifier(length int, prefix string) string {
	validChars := "abcdefghijklmnopqrstuvwxyz234567"
	return prefix + GenerateRandomString(length, validChars)
}

func GenerateRandomDigits(length int) string {
	return GenerateRandomString(length, "0123456789")
}

func GenerateRandomAlphanumeric(length int) string {
	return GenerateRandomString(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func GenerateRandomEmail() string {
	rdm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%d@checkout-sdk-go.com", rdm.Intn(9999999999))
}

func GenerateRandomReference() string {
	rdm := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("Reference-%d", rdm.Intn(9999999999))
}

func Wait(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}

func getOAuthScopes() []string {
	// The marketplace scope was retired; accounts is the documented requirement for the sub-entity
	// operations this fixture's suites exercise.
	return []string{configuration.Files, configuration.Flow, configuration.Fx, configuration.Gateway,
		configuration.Accounts, configuration.SessionsApp, configuration.SessionsBrowser,
		configuration.Vault, configuration.PayoutsBankDetails, configuration.Disputes,
		configuration.TransfersCreate, configuration.TransfersView, configuration.Balances,
		configuration.VaultCardMetadata, configuration.FinancialActions, configuration.PaymentsSearch}
}

func Bool(v bool) *bool { return &v }

func retriable(
	callback func() (interface{}, error),
	predicate func(interface{}) bool,
	seconds time.Duration,
) (response interface{}, err error) {
	attempt := 1
	for attempt <= MaxRetryAttemps {
		response, err = callback()
		if err != nil {
			return nil, err
		}
		if response != nil && predicate(response) {
			return response, nil
		}
		attempt++
		Wait(seconds)
	}

	return nil, err
}
