package instruments

import "github.com/checkout/checkout-sdk-go/common"

const Path = "instruments"

const ValidationPath = "validation/bank-accounts"

type InstrumentCustomerResponse struct {
	Id      string        `json:"id,omitempty"`
	Email   string        `json:"email,omitempty"`
	Name    string        `json:"name,omitempty"`
	Phone   *common.Phone `json:"phone,omitempty"`
	Default bool          `json:"default,omitempty"`
}
