package customers

import (
	"github.com/checkout/checkout-sdk-go/common"
)

const Path = "customers"

type (
	CustomerRequest struct {
		Email       string                 `json:"email,omitempty"`
		Name        string                 `json:"name,omitempty"`
		Phone       *common.Phone          `json:"phone,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Instruments []string               `json:"instruments,omitempty"`
	}

	GetCustomerResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                     `json:"id,omitempty"`
		DefaultId    string                     `json:"nas,omitempty"`
		Email        string                     `json:"email,omitempty"`
		Name         string                     `json:"name,omitempty"`
		Phone        *common.Phone              `json:"phone,omitempty"`
		Metadata     map[string]interface{}     `json:"metadata,omitempty"`
		Instruments  []common.InstrumentDetails `json:"instruments,omitempty"`
	}
)
