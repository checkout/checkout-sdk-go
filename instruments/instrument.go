package instruments

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type (
	// Request -
	Request struct {
		Type          string                `json:"type" binding:"required"`
		Token         string                `json:"token" binding:"required"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
		Customer      *Customer             `json:"customer,omitempty"`
	}

	// Customer -
	Customer struct {
		ID      string        `json:"id,omitempty"`
		Email   string        `json:"email,omitempty"`
		Name    string        `json:"name,omitempty"`
		Phone   *common.Phone `json:"phone,omitempty"`
		Default *bool         `json:"default,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Instrument     *InstrumentResponse      `json:"instrument_response,omitempty"`
	}

	// InstrumentResponse -
	InstrumentResponse struct {
		*Instrument
		Customer *Customer `json:"customer,omitempty"`
	}

	Instrument struct {
		ID            string               `json:"id"`
		Type          string               `json:"type,omitempty"`
		Fingerprint   string               `json:"fingerprint,omitempty"`
		ExpiryMonth   int                  `json:"expiry_month,omitempty"`
		ExpiryYear    int                  `json:"expiry_year,omitempty"`
		Name          string               `json:"name,omitempty"`
		Scheme        string               `json:"scheme,omitempty"`
		Last4         string               `json:"last4,omitempty"`
		Bin           string               `json:"bin,omitempty"`
		CardType      string               `json:"card_type,omitempty"`
		CardCategory  string               `json:"card_category,omitempty"`
		Issuer        string               `json:"issuer,omitempty"`
		IssuerCountry string               `json:"issuer_country,omitempty"`
		ProductId     string               `json:"product_id,omitempty"`
		ProductType   string               `json:"product_type,omitempty"`
		AccountHolder common.AccountHolder `json:"account_holder,omitempty"`
	}
)
