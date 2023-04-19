package accounts

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type InstrumentStatus string

const (
	Verified          InstrumentStatus = "verified"
	Unverified        InstrumentStatus = "unverified"
	InstrumentPending InstrumentStatus = "pending"
)

type (
	PaymentInstrument struct {
		Type          common.InstrumentType `json:"type,omitempty"`
		Label         string                `json:"label,omitempty"`
		AccountType   common.AccountType    `json:"account_type,omitempty"`
		AccountNumber string                `json:"account_number,omitempty"`
		BankCode      string                `json:"bank_code,omitempty"`
		BranchCode    string                `json:"branch_code,omitempty"`
		Iban          string                `json:"iban,omitempty"`
		Bban          string                `json:"bban,omitempty"`
		SwiftBic      string                `json:"swift_bic,omitempty"`
		Currency      common.Currency       `json:"currency,omitempty"`
		Country       common.Country        `json:"country,omitempty"`
		Document      *InstrumentDocument   `json:"document,omitempty"`
		AccountHolder *AccountHolder        `json:"account_holder,omitempty"`
		Bank          *common.BankDetails   `json:"bank,omitempty"`
	}
)

func NewAccountsPaymentInstrument() *PaymentInstrument {
	return &PaymentInstrument{Type: common.BankAccount}
}
