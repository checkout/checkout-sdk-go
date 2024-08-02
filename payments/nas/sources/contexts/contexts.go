package contexts

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	requestPaymentContextsKlarnaSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestPaymentContextsPayPalSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}
)

func NewPaymentContextsKlarnaSource() *requestPaymentContextsKlarnaSource {
	return &requestPaymentContextsKlarnaSource{Type: payments.KlarnaSource}
}

func NewPaymentContextsPayPalSource() *requestPaymentContextsPayPalSource {
	return &requestPaymentContextsPayPalSource{Type: payments.PayPalSource}
}

func (s *requestPaymentContextsKlarnaSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestPaymentContextsPayPalSource) GetType() payments.SourceType {
	return s.Type
}
