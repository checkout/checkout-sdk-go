package nas

import (
	"github.com/checkout/checkout-sdk-go-beta/accounts"
	"github.com/checkout/checkout-sdk-go-beta/apm/ideal"
	"github.com/checkout/checkout-sdk-go-beta/apm/klarna"
	"github.com/checkout/checkout-sdk-go-beta/apm/sepa"
	"github.com/checkout/checkout-sdk-go-beta/client"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
	"github.com/checkout/checkout-sdk-go-beta/customers"
	"github.com/checkout/checkout-sdk-go-beta/disputes"
	"github.com/checkout/checkout-sdk-go-beta/forex"
	instruments "github.com/checkout/checkout-sdk-go-beta/instruments/nas"
	payments "github.com/checkout/checkout-sdk-go-beta/payments/nas"
	"github.com/checkout/checkout-sdk-go-beta/sessions"
	"github.com/checkout/checkout-sdk-go-beta/tokens"
)

type Api struct {
	Tokens      *tokens.Client
	Instruments *instruments.Client
	Customers   *customers.Client
	Payments    *payments.Client
	Disputes    *disputes.Client
	Forex       *forex.Client
	Accounts    *accounts.Client
	Sessions    *sessions.Client

	Ideal  *ideal.Client
	Klarna *klarna.Client
	Sepa   *sepa.Client
}

func CheckoutApi(configuration *configuration.Configuration) *Api {
	apiClient := buildBaseClient(configuration)

	api := Api{}
	api.Tokens = tokens.NewClient(configuration, apiClient)
	api.Instruments = instruments.NewClient(configuration, apiClient)
	api.Customers = customers.NewClient(configuration, apiClient)
	api.Payments = payments.NewClient(configuration, apiClient)
	api.Disputes = disputes.NewClient(configuration, apiClient)
	api.Forex = forex.NewClient(configuration, apiClient)
	api.Accounts = accounts.NewClient(configuration, apiClient, buildFilesClient(configuration))
	api.Sessions = sessions.NewClient(configuration, apiClient)

	api.Ideal = ideal.NewClient(configuration, apiClient)
	api.Klarna = klarna.NewClient(configuration, apiClient)
	api.Sepa = sepa.NewClient(configuration, apiClient)
	return &api
}

func buildBaseClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.BaseUri())
}

func buildFilesClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.FilesUri())
}
