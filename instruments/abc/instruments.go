package abc

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type InstrumentAccountHolder struct {
	BillingAddress *common.Address `json:"billing_address,omitempty"`
	Phone          *common.Phone   `json:"phone,omitempty"`
}
