package nas

import (
	"github.com/checkout/checkout-sdk-go/common"
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

type CreateCustomerInstrumentRequest struct {
	Id      string        `json:"id,omitempty"`
	Email   string        `json:"email,omitempty"`
	Name    string        `json:"name,omitempty"`
	Phone   *common.Phone `json:"phone,omitempty"`
	Default bool          `json:"default,omitempty"`
}
