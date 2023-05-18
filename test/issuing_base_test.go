package test

import (
	"os"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/nas"
)

var (
	issuingClientApi *nas.Api
)

func buildIssuingClientApi() *nas.Api {
	if issuingClientApi == nil {
		issuingClientApi, _ = checkout.Builder().OAuth().
			WithClientCredentials(
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_ID"),
				os.Getenv("CHECKOUT_DEFAULT_OAUTH_ISSUING_CLIENT_SECRET")).
			WithEnvironment(configuration.Sandbox()).
			WithScopes([]string{
				configuration.Vault,
				configuration.IssuingClient,
				configuration.IssuingCardMgmt,
				configuration.IssuingControlsRead,
				configuration.IssuingControlsWrite}).
			Build()
	}

	return issuingClientApi
}
