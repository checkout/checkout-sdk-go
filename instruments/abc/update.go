package abc

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type (
	UpdateInstrumentRequest struct {
		ExpiryMonth   int                              `json:"expiry_month,omitempty"`
		ExpiryYear    int                              `json:"expiry_year,omitempty"`
		Name          string                           `json:"name,omitempty"`
		AccountHolder *InstrumentAccountHolder         `json:"account_holder,omitempty"`
		Customer      *InstrumentCustomerUpdateRequest `json:"customer,omitempty"`
	}

	UpdateInstrumentResponse struct {
		HttpMetadata common.HttpMetadata
		Type         common.InstrumentType `json:"type" binding:"required"`
		Fingerprint  string                `json:"fingerprint,omitempty"`
	}
)

type (
	InstrumentCustomerUpdateRequest struct {
		Id        string `json:"id,omitempty"`
		IsDefault bool   `json:"default,omitempty"`
	}
)
