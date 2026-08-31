package abc

import (
	"github.com/checkout/checkout-sdk-go/v3/common"
)

type InstrumentAccountHolder struct {
	BillingAddress *common.Address `json:"billing_address,omitempty"`
	Phone          *common.Phone   `json:"phone,omitempty"`
}
