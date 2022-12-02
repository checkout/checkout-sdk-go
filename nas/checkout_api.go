package nas

import (
	"github.com/checkout/checkout-sdk-go/accounts"
	"github.com/checkout/checkout-sdk-go/apm/ideal"
	"github.com/checkout/checkout-sdk-go/apm/klarna"
	"github.com/checkout/checkout-sdk-go/apm/sepa"
	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/customers"
	"github.com/checkout/checkout-sdk-go/disputes"
	"github.com/checkout/checkout-sdk-go/forex"
	instruments "github.com/checkout/checkout-sdk-go/instruments/nas"
	"github.com/checkout/checkout-sdk-go/metadata"
	"github.com/checkout/checkout-sdk-go/payments/hosted"
	"github.com/checkout/checkout-sdk-go/payments/links"
	payments "github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/sessions"
	"github.com/checkout/checkout-sdk-go/tokens"
	"github.com/checkout/checkout-sdk-go/workflows"
)

type Api struct {
	Tokens      *tokens.Client
	Instruments *instruments.Client
	Customers   *customers.Client
	Payments    *payments.Client
	Hosted      *hosted.Client
	Links       *links.Client
	Disputes    *disputes.Client
	Forex       *forex.Client
	Accounts    *accounts.Client
	Sessions    *sessions.Client
	Metadata    *metadata.Client
	WorkFlows   *workflows.Client

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
	api.Hosted = hosted.NewClient(configuration, apiClient)
	api.Links = links.NewClient(configuration, apiClient)
	api.Disputes = disputes.NewClient(configuration, apiClient)
	api.Forex = forex.NewClient(configuration, apiClient)
	api.Accounts = accounts.NewClient(configuration, apiClient, buildFilesClient(configuration))
	api.Sessions = sessions.NewClient(configuration, apiClient)
	api.Metadata = metadata.NewClient(configuration, apiClient)
	api.WorkFlows = workflows.NewClient(configuration, apiClient)

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
