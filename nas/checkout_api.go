package nas

import (
	"github.com/checkout/checkout-sdk-go/v2/accounts"
	"github.com/checkout/checkout-sdk-go/v2/apm/ideal"
	"github.com/checkout/checkout-sdk-go/v2/apm/klarna"
	"github.com/checkout/checkout-sdk-go/v2/apm/sepa"
	"github.com/checkout/checkout-sdk-go/v2/balances"
	"github.com/checkout/checkout-sdk-go/v2/client"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/customers"
	"github.com/checkout/checkout-sdk-go/v2/disputes"
	"github.com/checkout/checkout-sdk-go/v2/financial"
	"github.com/checkout/checkout-sdk-go/v2/forex"
	"github.com/checkout/checkout-sdk-go/v2/forward"
	instruments "github.com/checkout/checkout-sdk-go/v2/instruments/nas"
	"github.com/checkout/checkout-sdk-go/v2/issuing"
	"github.com/checkout/checkout-sdk-go/v2/metadata"
	"github.com/checkout/checkout-sdk-go/v2/networktokens"
	"github.com/checkout/checkout-sdk-go/v2/payments/applepay"
	"github.com/checkout/checkout-sdk-go/v2/payments/contexts"
	"github.com/checkout/checkout-sdk-go/v2/payments/googlepay"
	"github.com/checkout/checkout-sdk-go/v2/payments/hosted"
	"github.com/checkout/checkout-sdk-go/v2/payments/links"
	payments "github.com/checkout/checkout-sdk-go/v2/payments/nas"
	"github.com/checkout/checkout-sdk-go/v2/payments/sessions"
	"github.com/checkout/checkout-sdk-go/v2/payments/setups"
	"github.com/checkout/checkout-sdk-go/v2/reports"
	"github.com/checkout/checkout-sdk-go/v2/sessions"
	"github.com/checkout/checkout-sdk-go/v2/tokens"
	"github.com/checkout/checkout-sdk-go/v2/transfers"
	"github.com/checkout/checkout-sdk-go/v2/workflows"
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
	ApplePay        *applepay.Client
	GooglePay       *googlepay.Client
	NetworkTokens   *networktokens.Client

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
	api.ApplePay = applepay.NewClient(configuration, apiClient)
	api.GooglePay = googlepay.NewClient(configuration, apiClient)
	api.NetworkTokens = networktokens.NewClient(configuration, apiClient)

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
