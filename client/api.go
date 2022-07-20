package client

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/customers"
	"github.com/checkout/checkout-sdk-go/disputes"
	"github.com/checkout/checkout-sdk-go/events"
	"github.com/checkout/checkout-sdk-go/files"
	"github.com/checkout/checkout-sdk-go/instruments"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/reconciliation"
	"github.com/checkout/checkout-sdk-go/sources"
	"github.com/checkout/checkout-sdk-go/tokens"
	"github.com/checkout/checkout-sdk-go/webhooks"
)

// API -
type API struct {
	Payments       *payments.Client
	Sources        *sources.Client
	Tokens         *tokens.Client
	Events         *events.Client
	Webhooks       *webhooks.Client
	Disputes       *disputes.Client
	Files          *files.Client
	Reconciliation *reconciliation.Client
	Instruments    *instruments.Client
	Customers      *customers.Client
}

// Deprecated: This initialization method does not support the new Configuration entrypoint. To use the new entrypoint
// please use CheckoutApi
func (a *API) Init(secretKey string, useSandbox bool, publicKey *string) {

	config, err := checkout.Create(secretKey, publicKey)
	if err != nil {
		return
	}
	a.Payments = payments.NewClient(*config)
	a.Sources = sources.NewClient(*config)
	a.Tokens = tokens.NewClient(*config)
	a.Events = events.NewClient(*config)
	a.Webhooks = webhooks.NewClient(*config)
	a.Disputes = disputes.NewClient(*config)
	a.Files = files.NewClient(*config)
	a.Reconciliation = reconciliation.NewClient(*config)
}

// Deprecated: This initialization method does not support the new Configuration entrypoint. To use the new entrypoint
// please use CheckoutApi
func New(secretKey string, useSandbox bool, publicKey *string) *API {

	api := API{}
	api.Init(secretKey, useSandbox, publicKey)
	return &api
}

func CheckoutApi(secretKey *string, publicKey *string, environment checkout.SupportedEnvironment) *API {
	config, err := checkout.SdkConfig(secretKey, publicKey, environment)
	if err != nil {
		panic(err)
	}
	api := API{}
	api.Payments = payments.NewClient(*config)
	api.Sources = sources.NewClient(*config)
	api.Tokens = tokens.NewClient(*config)
	api.Events = events.NewClient(*config)
	api.Webhooks = webhooks.NewClient(*config)
	api.Disputes = disputes.NewClient(*config)
	api.Files = files.NewClient(*config)
	api.Reconciliation = reconciliation.NewClient(*config)
	api.Instruments = instruments.NewClient(*config)
	api.Customers = customers.NewClient(*config)
	return &api
}
