package abc

import (
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/instruments"
)

type (
	CreateInstrumentRequest struct {
		Type          instruments.InstrumentType `json:"type" binding:"required"`
		Token         string                     `json:"token,omitempty"`
		AccountHolder *InstrumentAccountHolder   `json:"account_holder,omitempty"`
		Customer      *InstrumentCustomerRequest `json:"customer,omitempty"`
	}

	CreateInstrumentResponse struct {
		HttpMetadata  common.HttpMetadata
		Type          instruments.InstrumentType `json:"type,omitempty"`
		Id            string                     `json:"id,omitempty"`
		Fingerprint   string                     `json:"fingerprint,omitempty"`
		ExpiryMonth   int                        `json:"expiry_month,omitempty"`
		ExpiryYear    int                        `json:"expiry_year,omitempty"`
		Scheme        string                     `json:"scheme,omitempty"`
		Last4         string                     `json:"last4,omitempty"`
		Bin           string                     `json:"bin,omitempty"`
		CardType      common.CardType            `json:"card_type,omitempty"`
		CardCategory  common.CardCategory        `json:"card_category,omitempty"`
		Issuer        string                     `json:"issuer,omitempty"`
		IssuerCountry common.Country             `json:"issuer_country,omitempty"`
		ProductId     string                     `json:"product_id,omitempty"`
		ProductType   string                     `json:"product_type,omitempty"`
		Customer      *common.CustomerResponse   `json:"customer,omitempty"`
	}
)

type InstrumentCustomerRequest struct {
	Id        string        `json:"id,omitempty"`
	Email     string        `json:"email,omitempty"`
	Name      string        `json:"name,omitempty"`
	Phone     *common.Phone `json:"phone,omitempty"`
	IsDefault bool          `json:"nas,omitempty"`
}
