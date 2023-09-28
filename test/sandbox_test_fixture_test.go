package test

import (
	"fmt"
	"math/rand"
	"os"
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
