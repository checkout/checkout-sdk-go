package abc

import (
	"github.com/checkout/checkout-sdk-go-beta/apm/ideal"
	"github.com/checkout/checkout-sdk-go-beta/apm/klarna"
	"github.com/checkout/checkout-sdk-go-beta/apm/sepa"
	"github.com/checkout/checkout-sdk-go-beta/client"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
	"github.com/checkout/checkout-sdk-go-beta/customers"
	"github.com/checkout/checkout-sdk-go-beta/disputes"
	"github.com/checkout/checkout-sdk-go-beta/events"
	"github.com/checkout/checkout-sdk-go-beta/instruments/abc"
	payments "github.com/checkout/checkout-sdk-go-beta/payments/abc"
	"github.com/checkout/checkout-sdk-go-beta/sources"
	"github.com/checkout/checkout-sdk-go-beta/tokens"
)

type Api struct {
	Tokens      *tokens.Client
	Events      *events.Client
	Sources     *sources.Client
	Instruments *abc.Client
	Customers   *customers.Client
	Payments    *payments.Client
	Disputes    *disputes.Client

	Ideal  *ideal.Client
	Klarna *klarna.Client
	Sepa   *sepa.Client
}

func CheckoutApi(configuration *configuration.Configuration) *Api {
	apiClient := buildBaseClient(configuration)

	api := Api{}
	api.Tokens = tokens.NewClient(configuration, apiClient)
	api.Events = events.NewClient(configuration, apiClient)
	api.Sources = sources.NewClient(configuration, apiClient)
	api.Instruments = abc.NewClient(configuration, apiClient)
	api.Customers = customers.NewClient(configuration, apiClient)
	api.Payments = payments.NewClient(configuration, apiClient)
	api.Disputes = disputes.NewClient(configuration, apiClient)

	api.Ideal = ideal.NewClient(configuration, apiClient)
	api.Klarna = klarna.NewClient(configuration, apiClient)
	api.Sepa = sepa.NewClient(configuration, apiClient)
	return &api
}

func buildBaseClient(configuration *configuration.Configuration) client.HttpClient {
	return client.NewApiClient(configuration, configuration.Environment.BaseUri())
}
