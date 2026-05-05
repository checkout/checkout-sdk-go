package agenticcommerce

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

const delegatePaymentPath = "agentic_commerce/delegate_payment"

// CardNumberType indicates whether the card number supplied in a delegated payment is a
// full primary account number or a network token.
type CardNumberType string

const (
	// Fpan represents a full primary account number (the physical card number).
	Fpan CardNumberType = "fpan"
	// NetworkToken represents a payment network token that replaces the PAN.
	NetworkToken CardNumberType = "network_token"
)

// DisplayCardFundingType classifies the card's funding source as reported to the merchant.
type DisplayCardFundingType string

const (
	// Credit indicates a credit card.
	Credit DisplayCardFundingType = "credit"
	// Debit indicates a debit card.
	Debit DisplayCardFundingType = "debit"
	// Prepaid indicates a prepaid card.
	Prepaid DisplayCardFundingType = "prepaid"
)

// AllowanceReason describes why a delegated payment allowance was granted.
type AllowanceReason string

const (
	// OneTime indicates the allowance is valid for a single transaction only.
	OneTime AllowanceReason = "one_time"
)

// DelegatedPaymentHeaders holds the HTTP headers required by the delegated payment endpoint.
type DelegatedPaymentHeaders struct {
	Signature  string `json:"Signature"`
	Timestamp  string `json:"Timestamp"`
	APIVersion string `json:"API-Version,omitempty"`
}

// DelegatedPaymentMethodCard represents a card payment method used in a delegated payment request.
type DelegatedPaymentMethodCard struct {
	Type                   string                 `json:"type"`
	CardNumberType         CardNumberType         `json:"card_number_type"`
	Number                 string                 `json:"number"`
	ExpMonth               string                 `json:"exp_month,omitempty"`
	ExpYear                string                 `json:"exp_year,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Cvc                    string                 `json:"cvc,omitempty"`
	Cryptogram             string                 `json:"cryptogram,omitempty"`
	EciValue               string                 `json:"eci_value,omitempty"`
	ChecksPerformed        []string               `json:"checks_performed,omitempty"`
	Iin                    string                 `json:"iin,omitempty"`
	DisplayCardFundingType DisplayCardFundingType `json:"display_card_funding_type,omitempty"`
	DisplayWalletType      string                 `json:"display_wallet_type,omitempty"`
	DisplayBrand           string                 `json:"display_brand,omitempty"`
	DisplayLast4           string                 `json:"display_last4,omitempty"`
	Metadata               map[string]string      `json:"metadata"`
}

// NewDelegatedPaymentMethodCard returns a DelegatedPaymentMethodCard with the type field
// pre-set to "card".
func NewDelegatedPaymentMethodCard() *DelegatedPaymentMethodCard {
	return &DelegatedPaymentMethodCard{Type: "card"}
}

// DelegatedPaymentAllowance describes the permission and constraints for a delegated payment,
// including the maximum amount, currency, and expiry.
type DelegatedPaymentAllowance struct {
	Reason            AllowanceReason `json:"reason"`
	MaxAmount         int             `json:"max_amount"`
	Currency          common.Currency `json:"currency"`
	MerchantId        string          `json:"merchant_id"`
	CheckoutSessionId string          `json:"checkout_session_id"`
	ExpiresAt         *time.Time      `json:"expires_at"`
}

// DelegatedPaymentBillingAddress holds the billing address associated with a delegated payment.
type DelegatedPaymentBillingAddress struct {
	Name       string         `json:"name"`
	LineOne    string         `json:"line_one"`
	LineTwo    string         `json:"line_two,omitempty"`
	City       string         `json:"city"`
	State      string         `json:"state,omitempty"`
	PostalCode string         `json:"postal_code"`
	Country    common.Country `json:"country"`
}

// DelegatedPaymentRiskSignal represents a risk assessment signal provided by the token requestor.
type DelegatedPaymentRiskSignal struct {
	Type   string `json:"type"`
	Score  int    `json:"score"`
	Action string `json:"action"`
}

// CreateDelegatedPaymentTokenRequest is the request body for the create delegated payment token endpoint.
type CreateDelegatedPaymentTokenRequest struct {
	PaymentMethod  DelegatedPaymentMethodCard      `json:"payment_method"`
	Allowance      DelegatedPaymentAllowance       `json:"allowance"`
	BillingAddress *DelegatedPaymentBillingAddress `json:"billing_address,omitempty"`
	RiskSignals    []DelegatedPaymentRiskSignal    `json:"risk_signals"`
	Metadata       map[string]string               `json:"metadata"`
	Headers        DelegatedPaymentHeaders         `json:"-"`
}

// CreateDelegatedPaymentTokenResponse is the response returned by the create delegated payment
// token endpoint.
type CreateDelegatedPaymentTokenResponse struct {
	HttpMetadata common.HttpMetadata
	Id           string            `json:"id,omitempty"`
	Created      *time.Time        `json:"created,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}
