package customers

import (
	instruments "github.com/checkout/checkout-sdk-go/instruments/nas"

	"github.com/checkout/checkout-sdk-go/common"
)

const Path = "customers"

type (
	CustomerRequest struct {
		Email       string                 `json:"email,omitempty"`
		Name        string                 `json:"name,omitempty"`
		Phone       *common.Phone          `json:"phone,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		DefaultId   string                 `json:"default,omitempty"`
		Instruments []string               `json:"instruments,omitempty"`
	}

	GetCustomerResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                              `json:"id,omitempty"`
		Email        string                              `json:"email,omitempty"`
		Default      string                              `json:"default,omitempty"`
		Name         string                              `json:"name,omitempty"`
		Phone        *common.Phone                       `json:"phone,omitempty"`
		Metadata     map[string]interface{}              `json:"metadata,omitempty"`
		Instruments  []instruments.GetInstrumentResponse `json:"instruments,omitempty"`
	}
)
