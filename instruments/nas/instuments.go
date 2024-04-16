package nas

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
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
	DateOfSignature *time.Time           `json:"date_of_signature,omitempty"`
}

type CreateCustomerInstrumentRequest struct {
	Id      string        `json:"id,omitempty"`
	Email   string        `json:"email,omitempty"`
	Name    string        `json:"name,omitempty"`
	Phone   *common.Phone `json:"phone,omitempty"`
	Default bool          `json:"default,omitempty"`
}
