package customers

import (
	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/instruments"
)

type (
	// Request -
	Request struct {
		*Customer
	}

	// Customer request
	Customer struct {
		Email    string                 `json:"email" binding:"required"`
		Name     string                 `json:"name,omitempty"`
		Phone    *common.Phone          `json:"phone,omitempty"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}
)

type (
	// Response for customer
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Customer       *CustomerResponse        `json:"customer,omitempty"`
	}

	// CustomerResponse - When the create endpoint is called, only ID fill be filled
	CustomerResponse struct {
		ID          string                   `json:"id"`
		Email       string                   `json:"email,omitempty"`
		Default     string                   `json:"default,omitempty"`
		Name        string                   `json:"name,omitempty"`
		Phone       *common.Phone            `json:"phone,omitempty"`
		Metadata    map[string]interface{}   `json:"metadata,omitempty"`
		Instruments []instruments.Instrument `json:"instruments,omitempty"`
	}
)
