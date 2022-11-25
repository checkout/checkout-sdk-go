package instruments

import "github.com/checkout/checkout-sdk-go/common"

const Path = "instruments"

type InstrumentType string

const (
	Card        InstrumentType = "card"
	BankAccount InstrumentType = "bank_account"
	Token       InstrumentType = "token"
)

type InstrumentCustomerResponse struct {
	Id      string        `json:"id,omitempty"`
	Email   string        `json:"email,omitempty"`
	Name    string        `json:"name,omitempty"`
	Phone   *common.Phone `json:"phone,omitempty"`
	Default bool          `json:"nas,omitempty"`
}
