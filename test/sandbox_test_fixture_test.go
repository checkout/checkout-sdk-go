package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/abc"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/nas"
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
		previousApi, _ = checkout.Builder().Previous().
			WithEnvironment(configuration.Sandbox()).
			WithSecretKey(os.Getenv("CHECKOUT_PREVIOUS_SECRET_KEY")).
			WithPublicKey(os.Getenv("CHECKOUT_PREVIOUS_PUBLIC_KEY")).
			Build()
	}
	return previousApi
}

func DefaultApi() *nas.Api {
	if defaultApi == nil {
		defaultApi, _ = checkout.Builder().
			StaticKeys().
			WithEnvironment(configuration.Sandbox()).
			WithSecretKey(os.Getenv("CHECKOUT_DEFAULT_SECRET_KEY")).
			WithPublicKey(os.Getenv("CHECKOUT_DEFAULT_PUBLIC_KEY")).
			Build()
	}
	return defaultApi
}

func OAuthApi() *nas.Api {
	if oauthApi == nil {
		oauthApi, _ = checkout.Builder().OAuth().
			WithClientCredentials(
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_ID"),
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET")).
			WithEnvironment(configuration.Sandbox()).
			WithScopes(getOAuthScopes()).
			Build()
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
	return []string{configuration.Files, configuration.Flow, configuration.Fx, configuration.Gateway,
		configuration.Marketplace, configuration.SessionsApp, configuration.SessionsBrowser,
		configuration.Vault, configuration.PayoutsBankDetails, configuration.Disputes,
		configuration.TransfersCreate, configuration.TransfersView, configuration.Balances,
		configuration.VaultCardMetadata, configuration.FinancialActions}
}

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
