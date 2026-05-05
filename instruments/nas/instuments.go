package nas

import (
	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/payments"
)

type PaymentNetwork string

const (
	Local   PaymentNetwork = "local"
	Sepa    PaymentNetwork = "sepa"
	Fps     PaymentNetwork = "fps"
	Ach     PaymentNetwork = "ach"
	Fedwire PaymentNetwork = "fedwire"
	Swift   PaymentNetwork = "swift"
)

type InstrumentData struct {
	AccountNumber   string               `json:"account_number,omitempty"`
	Country         common.Country       `json:"country,omitempty"`
	Currency        common.Currency      `json:"currency,omitempty"`
	PaymentType     payments.PaymentType `json:"payment_type,omitempty"`
	MandateId       string               `json:"mandate_id,omitempty"`
	DateOfSignature *common.APIShortDate `json:"date_of_signature,omitempty"`
}

type AchInstrumentData struct {
	AccountType   string          `json:"account_type,omitempty"`
	AccountNumber string          `json:"account_number,omitempty"`
	BankCode      string          `json:"bank_code,omitempty"`
	Currency      common.Currency `json:"currency,omitempty"`
	Country       common.Country  `json:"country,omitempty"`
}

type ProvisionNetworkToken struct {
	Provision bool `json:"provision,omitempty"`
}

type NetworkTokenResponse struct {
	Id    string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
}

type CreateCustomerInstrumentRequest struct {
	Id      string        `json:"id,omitempty"`
	Email   string        `json:"email,omitempty"`
	Name    string        `json:"name,omitempty"`
	Phone   *common.Phone `json:"phone,omitempty"`
	Default bool          `json:"default,omitempty"`
}
