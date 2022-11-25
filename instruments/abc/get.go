package abc

import (
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/instruments"
)

type (
	GetInstrumentResponse struct {
		HttpMetadata  common.HttpMetadata
		Type          instruments.InstrumentType              `json:"type,omitempty"`
		Id            string                                  `json:"id,omitempty"`
		Fingerprint   string                                  `json:"fingerprint,omitempty"`
		ExpiryMonth   int                                     `json:"expiry_month,omitempty"`
		ExpiryYear    int                                     `json:"expiry_year,omitempty"`
		Name          string                                  `json:"name,omitempty"`
		Scheme        string                                  `json:"scheme,omitempty"`
		Last4         string                                  `json:"last4,omitempty"`
		Bin           string                                  `json:"bin,omitempty"`
		CardType      common.CardType                         `json:"card_type,omitempty"`
		CardCategory  common.CardCategory                     `json:"card_category,omitempty"`
		Issuer        string                                  `json:"issuer,omitempty"`
		IssuerCountry common.Country                          `json:"issuer_country,omitempty"`
		ProductId     string                                  `json:"product_id,omitempty"`
		ProductType   string                                  `json:"product_type,omitempty"`
		AccountHolder *InstrumentAccountHolder                `json:"account_holder,omitempty"`
		Customer      *instruments.InstrumentCustomerResponse `json:"customer,omitempty"`
	}
)
