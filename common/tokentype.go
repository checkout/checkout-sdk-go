package common

// TokenType ...
type TokenType string

const (
	// Token ...
	Token TokenType = "token"
	// Card ...
	Card TokenType = "card"
	// ApplePay ...
	ApplePay TokenType = "ApplePay"
	// GooglePay ...
	GooglePay TokenType = "GooglePay"
)

func (c TokenType) String() string {
	return string(c)
}
