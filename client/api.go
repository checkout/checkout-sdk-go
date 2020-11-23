package client

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/disputes"
	"github.com/checkout/checkout-sdk-go/events"
	"github.com/checkout/checkout-sdk-go/files"
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
}

// Init -
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

// New -
func New(secretKey string, useSandbox bool, publicKey *string) *API {

	api := API{}
	api.Init(secretKey, useSandbox, publicKey)
	return &api
}
