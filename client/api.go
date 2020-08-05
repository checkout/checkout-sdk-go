package client

import (
	"github.com/shiuh-yaw-cko/checkout"
	"github.com/shiuh-yaw-cko/checkout/disputes"
	"github.com/shiuh-yaw-cko/checkout/events"
	"github.com/shiuh-yaw-cko/checkout/files"
	"github.com/shiuh-yaw-cko/checkout/payments"
	"github.com/shiuh-yaw-cko/checkout/reconciliation"
	"github.com/shiuh-yaw-cko/checkout/sources"
	"github.com/shiuh-yaw-cko/checkout/tokens"
	"github.com/shiuh-yaw-cko/checkout/webhooks"
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
func (a *API) Init(secretKey string, useSandbox bool, publicKey *string, idempotencyKey *string) {

	config, err := checkout.Create(secretKey, useSandbox, publicKey, idempotencyKey)
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
func New(secretKey string, useSandbox bool, publicKey *string, idempotencyKey *string) *API {

	api := API{}
	api.Init(secretKey, useSandbox, publicKey, idempotencyKey)
	return &api
}
