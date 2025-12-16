package nas

import (
	"github.com/checkout/checkout-sdk-go/accounts"
	"github.com/checkout/checkout-sdk-go/apm/ideal"
	"github.com/checkout/checkout-sdk-go/apm/klarna"
	"github.com/checkout/checkout-sdk-go/apm/sepa"
	"github.com/checkout/checkout-sdk-go/balances"
	"github.com/checkout/checkout-sdk-go/client"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/customers"
	"github.com/checkout/checkout-sdk-go/disputes"
	"github.com/checkout/checkout-sdk-go/financial"
	"github.com/checkout/checkout-sdk-go/forex"
	"github.com/checkout/checkout-sdk-go/forward"
	instruments "github.com/checkout/checkout-sdk-go/instruments/nas"
	"github.com/checkout/checkout-sdk-go/issuing"
	"github.com/checkout/checkout-sdk-go/metadata"
	"github.com/checkout/checkout-sdk-go/payments/contexts"
	"github.com/checkout/checkout-sdk-go/payments/hosted"
	"github.com/checkout/checkout-sdk-go/payments/links"
	payments "github.com/checkout/checkout-sdk-go/payments/nas"
	"github.com/checkout/checkout-sdk-go/payments/sessions"
	"github.com/checkout/checkout-sdk-go/payments/setups"
	"github.com/checkout/checkout-sdk-go/reports"
	"github.com/checkout/checkout-sdk-go/sessions"
	"github.com/checkout/checkout-sdk-go/tokens"
	"github.com/checkout/checkout-sdk-go/transfers"
	"github.com/checkout/checkout-sdk-go/workflows"
)

type Api struct {
	Accounts        *accounts.Client
	Balances        *balances.Client
	Customers       *customers.Client
	Disputes        *disputes.Client
	Financial       *financial.Client
	Forex           *forex.Client
	Hosted          *hosted.Client
	Instruments     *instruments.Client
	Links           *links.Client
	Metadata        *metadata.Client
	Payments        *payments.Client
	Sessions        *sessions.Client
	Tokens          *tokens.Client
	Transfers       *transfers.Client
	WorkFlows       *workflows.Client
	Reports         *reports.Client
	Issuing         *issuing.Client
	Contexts        *contexts.Client
	PaymentSessions *payment_sessions.Client
	PaymentSetups   *setups.Client
	Forward         *forward.Client

	Ideal  *ideal.Client
	Klarna *klarna.Client
	Sepa   *sepa.Client
}

func CheckoutApi(configuration *configuration.Configuration) *Api {
	apiClient := buildBaseClient(configuration)

	api := Api{}
	api.Accounts = accounts.NewClient(configuration, apiClient, buildFilesClient(configuration))
	api.Balances = balances.NewClient(configuration, buildBalancesClient(configuration))
	api.Customers = customers.NewClient(configuration, apiClient)
	api.Disputes = disputes.NewClient(configuration, apiClient)
	api.Instruments = instruments.NewClient(configuration, apiClient)
	api.Financial = financial.NewClient(configuration, apiClient)
	api.Forex = forex.NewClient(configuration, apiClient)
	api.Hosted = hosted.NewClient(configuration, apiClient)
	api.Links = links.NewClient(configuration, apiClient)
	api.Metadata = metadata.NewClient(configuration, apiClient)
	api.Payments = payments.NewClient(configuration, apiClient)
	api.Sessions = sessions.NewClient(configuration, apiClient)
	api.Tokens = tokens.NewClient(configuration, apiClient)
	api.Transfers = transfers.NewClient(configuration, buildTransfersClient(configuration))
	api.WorkFlows = workflows.NewClient(configuration, apiClient)
	api.Reports = reports.NewClient(configuration, apiClient)
	api.Issuing = issuing.NewClient(configuration, apiClient)
	api.Contexts = contexts.NewClient(configuration, apiClient)
	api.PaymentSessions = payment_sessions.NewClient(configuration, apiClient)
	api.PaymentSetups = setups.NewClient(configuration, apiClient)
	api.Forward = forward.NewClient(configuration, apiClient)

	api.Ideal = ideal.NewClient(configuration, apiClient)
	api.Klarna = klarna.NewClient(configuration, apiClient)
	api.Sepa = sepa.NewClient(configuration, apiClient)
	return &api
}

func buildBaseClient(configuration *configuration.Configuration) client.HttpClient {
	if configuration.EnvironmentSubdomain != nil {
		return client.NewApiClient(configuration, configuration.EnvironmentSubdomain.ApiUrl)
	}
	return client.NewApiClient(configuration, configuration.Environment.BaseUri())
}

func buildFilesClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.FilesUri())
}

func buildBalancesClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.BalancesUri())
}

func buildTransfersClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.TransfersUri())
}
