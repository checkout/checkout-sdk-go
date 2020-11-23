package instruments

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	// Request -
	Request struct {
		*Instrument
		*Source
	}
	// Instrument -
	Instrument struct {
		Type  string `json:"type" binding:"required"`
		Token string `json:"token" binding:"required"`
	}
	// Source -
	Source struct {
		ExpiryMonth   uint64         `json:"expiry_month,omitempty"`
		ExpiryYear    uint64         `json:"expiry_year,omitempty"`
		Name          string         `json:"name,omitempty"`
		AccountHolder *AccountHolder `json:"account_holder,omitempty"`
		Customer      *Customer      `json:"customer,omitempty"`
	}
	// AccountHolder -
	AccountHolder struct {
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}
	// Customer -
	Customer struct {
		ID      string `json:"id,omitempty"`
		Default *bool  `json:"default,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse     *checkout.StatusResponse `json:"api_response,omitempty"`
		Source             *payments.SourceResponse `json:"source,omitempty"`
		InstrumentResponse *InstrumentResponse      `json:"instrument_response,omitempty"`
	}

	// InstrumentResponse -
	InstrumentResponse struct {
		Type        string `json:"type" binding:"required"`
		Fingerprint string `json:"fingerprint,omitempty"`
	}
)
