package customers

import (
	"github.com/checkout/checkout-sdk-go"
)

type (
	// Request -
	Request struct {
		*Customer
	}
	// Customer -
	Customer struct {
		Email       string       `json:"email,omitempty"`
		Name        string       `json:"name,omitempty"`
		Default     string       `json:"default,omitempty"`
		Instruments []Instrument `json:"instruments,omitempty"`
	}

	Instrument struct {
		ID            string `json:"id,omitempty"`
		Type          string `json:"type,omitempty"`
		Fingerprint   string `json:"fingerprint,omitempty"`
		ExpiryMonth   int    `json:"expiry_month,omitempty"`
		ExpiryYear    int    `json:"expiry_year,omitempty"`
		Scheme        string `json:"scheme,omitempty"`
		Last4         string `json:"last4,omitempty"`
		Bin           string `json:"bin,omitempty"`
		CardType      string `json:"card_type,omitempty"`
		CardCategory  string `json:"card_category,omitempty"`
		Issuer        string `json:"issuer,omitempty"`
		IssuerCountry string `json:"issuer_country,omitempty"`
		ProductID     string `json:"product_id,omitempty"`
		ProductType   string `json:"product_type,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	}

	GetCustomerResponse struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Customer       *Customer                `json:"customer,omitempty"`
	}
)
