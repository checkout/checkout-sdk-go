package nas

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type CreateCustomerInstrumentRequest struct {
	Id      string        `json:"id,omitempty"`
	Email   string        `json:"email,omitempty"`
	Name    string        `json:"name,omitempty"`
	Phone   *common.Phone `json:"phone,omitempty"`
	Default bool          `json:"nas,omitempty"`
}
