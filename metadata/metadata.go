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
	Accel           SchemeLocalType = "accel"
	CartesBancaires SchemeLocalType = "cartes_bancaires"
	Mada            SchemeLocalType = "mada"
	Nyce            SchemeLocalType = "nyce"
	Omannet         SchemeLocalType = "omannet"
	Pulse           SchemeLocalType = "pulse"
	Star            SchemeLocalType = "star"
	Upi             SchemeLocalType = "upi"
)

type PayoutsTransactionsType string

const (
	NotSupported PayoutsTransactionsType = "not_supported"
	Standard     PayoutsTransactionsType = "standard"
	FastFunds    PayoutsTransactionsType = "fast_funds"
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

	PinlessDebitSchemeMetadata struct {
		NetworkId               string `json:"network_id,omitempty"`
		NetworkDescription      string `json:"network_description,omitempty"`
		BillPayIndicator        bool   `json:"bill_pay_indicator,omitempty"`
		EcommerceIndicator      bool   `json:"ecommerce_indicator,omitempty"`
		InterchangeFeeIndicator string `json:"interchange_fee_indicator,omitempty"`
		MoneyTransferIndicator  bool   `json:"money_transfer_indicator,omitempty"`
		TokenIndicator          bool   `json:"token_indicator,omitempty"`
	}

	SchemeMetadata struct {
		Accel *PinlessDebitSchemeMetadata `json:"accel,omitempty"`
		Pulse *PinlessDebitSchemeMetadata `json:"pulse,omitempty"`
		Nyce  *PinlessDebitSchemeMetadata `json:"nyce,omitempty"`
		Star  *PinlessDebitSchemeMetadata `json:"star,omitempty"`
	}

	CardMetadataResponse struct {
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		Bin          string              `json:"bin,omitempty"`
		Scheme       string              `json:"scheme,omitempty"`
		// Deprecated: This property will be removed in the future, and should not be used. Use LocalSchemes instead.
		SchemeLocal        SchemeLocalType      `json:"scheme_local,omitempty"`
		LocalSchemes       []SchemeLocalType    `json:"local_schemes,omitempty"`
		CardType           common.CardType      `json:"card_type,omitempty"`
		CardCategory       common.CardCategory  `json:"card_category,omitempty"`
		Currency           common.Currency      `json:"currency,omitempty"`
		Issuer             string               `json:"issuer,omitempty"`
		IssuerCountry      common.Country       `json:"issuer_country,omitempty"`
		IssuerCountryName  string               `json:"issuer_country_name,omitempty"`
		ProductId          string               `json:"product_id,omitempty"`
		ProductType        string               `json:"product_type,omitempty"`
		SubproductId       string               `json:"subproduct_id,omitempty"`
		RegulatedIndicator bool                 `json:"regulated_indicator,omitempty"`
		RegulatedType      string               `json:"regulated_type,omitempty"`
		CardPayouts        *CardMetadataPayouts `json:"card_payouts,omitempty"`
		SchemeMetadata     *SchemeMetadata      `json:"scheme_metadata,omitempty"`
	}
)
