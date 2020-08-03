package common

// PaymentType ...
type PaymentType string

const (
	// Regular ...
	Regular PaymentType = "Regular"
	// Recurring ...
	Recurring PaymentType = "Recurring"
	// MOTO ...
	MOTO PaymentType = "MOTO"
)

func (c PaymentType) String() string {
	return string(c)
}
