package contexts

import "github.com/checkout/checkout-sdk-go/payments"

type (
	requestPaymentContextsPaypalSource struct {
		Type payments.SourceType `json:"type,omitempty"`
	}
)

func NewPaymentContextsPaypalSource() *requestPaymentContextsPaypalSource {
	return &requestPaymentContextsPaypalSource{Type: payments.PayPalSource}
}

func (s *requestPaymentContextsPaypalSource) GetType() payments.SourceType {
	return s.Type
}
