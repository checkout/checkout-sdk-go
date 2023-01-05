package metadata

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/metadata/sources"
)

const (
	metadata = "metadata"
	card     = "card"
)

type Format string

const (
	Basic       Format = "basic"
	CardPayouts Format = "card_payouts"
)

type SchemeLocalType string

const (
	CartesBancaires SchemeLocalType = "cartes_bancaires"
	Mada            SchemeLocalType = "mada"
	Omannet         SchemeLocalType = "Omannet"
)

type PayoutsTransactionsType string

const (
	NotSupported PayoutsTransactionsType = "not_supported"
	Standard     PayoutsTransactionsType = "standard"
	FastFounds   PayoutsTransactionsType = "fast_founds"
	Unknown      PayoutsTransactionsType = "unknown"
)

type (
	CardMetadataRequest struct {
		Source sources.SourceRequest `json:"source,omitempty"`
		Format Format                `json:"format,omitempty"`
	}
)

type (
	CardMetadataPayouts struct {
		DomesticNonMoneyTransfer    PayoutsTransactionsType `json:"domestic_non_money_transfer,omitempty"`
		CrossBorderNonMoneyTransfer PayoutsTransactionsType `json:"cross_border_non_money_transfer,omitempty"`
		DomesticGambling            PayoutsTransactionsType `json:"domestic_gambling,omitempty"`
		CrossBorderGambling         PayoutsTransactionsType `json:"cross_border_gambling,omitempty"`
		DomesticMoneyTransfer       PayoutsTransactionsType `json:"domestic_money_transfer,omitempty"`
		CrossBorderMoneyTransfer    PayoutsTransactionsType `json:"cross_border_money_transfer,omitempty"`
	}

	CardMetadataResponse struct {
		HttpMetadata      common.HttpMetadata  `json:"http_metadata,omitempty"`
		Bin               string               `json:"bin,omitempty"`
		Scheme            string               `json:"scheme,omitempty"`
		SchemeLocal       SchemeLocalType      `json:"scheme_local,omitempty"`
		CardType          common.CardType      `json:"card_type,omitempty"`
		CardCategory      common.CardCategory  `json:"card_category,omitempty"`
		Issuer            string               `json:"issuer,omitempty"`
		IssuerCountry     common.Country       `json:"issuer_country,omitempty"`
		IssuerCountryName string               `json:"issuer_country_name,omitempty"`
		ProductId         string               `json:"product_id,omitempty"`
		ProductType       string               `json:"product_type,omitempty"`
		CardPayouts       *CardMetadataPayouts `json:"card-payouts,omitempty"`
	}
)
