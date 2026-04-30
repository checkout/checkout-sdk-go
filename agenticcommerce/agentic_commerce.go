package agenticcommerce

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

const delegatePaymentPath = "agentic_commerce/delegate_payment"

type CardNumberType string

const (
	Fpan         CardNumberType = "fpan"
	NetworkToken CardNumberType = "network_token"
)

type DisplayCardFundingType string

const (
	Credit  DisplayCardFundingType = "credit"
	Debit   DisplayCardFundingType = "debit"
	Prepaid DisplayCardFundingType = "prepaid"
)

type AllowanceReason string

const (
	OneTime AllowanceReason = "one_time"
)

type DelegatedPaymentHeaders struct {
	Signature  string `json:"Signature"`
	Timestamp  string `json:"Timestamp"`
	APIVersion string `json:"API-Version,omitempty"`
}

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

func NewDelegatedPaymentMethodCard() *DelegatedPaymentMethodCard {
	return &DelegatedPaymentMethodCard{Type: "card"}
}

type DelegatedPaymentAllowance struct {
	Reason            AllowanceReason `json:"reason"`
	MaxAmount         int             `json:"max_amount"`
	Currency          common.Currency `json:"currency"`
	MerchantId        string          `json:"merchant_id"`
	CheckoutSessionId string          `json:"checkout_session_id"`
	ExpiresAt         *time.Time      `json:"expires_at"`
}

type DelegatedPaymentBillingAddress struct {
	Name       string         `json:"name"`
	LineOne    string         `json:"line_one"`
	LineTwo    string         `json:"line_two,omitempty"`
	City       string         `json:"city"`
	State      string         `json:"state,omitempty"`
	PostalCode string         `json:"postal_code"`
	Country    common.Country `json:"country"`
}

type DelegatedPaymentRiskSignal struct {
	Type   string `json:"type"`
	Score  int    `json:"score"`
	Action string `json:"action"`
}

type CreateDelegatedPaymentTokenRequest struct {
	PaymentMethod  DelegatedPaymentMethodCard      `json:"payment_method"`
	Allowance      DelegatedPaymentAllowance       `json:"allowance"`
	BillingAddress *DelegatedPaymentBillingAddress `json:"billing_address,omitempty"`
	RiskSignals    []DelegatedPaymentRiskSignal    `json:"risk_signals"`
	Metadata       map[string]string               `json:"metadata"`
	Headers        DelegatedPaymentHeaders         `json:"-"`
}

type CreateDelegatedPaymentTokenResponse struct {
	HttpMetadata common.HttpMetadata
	Id           string            `json:"id,omitempty"`
	Created      *time.Time        `json:"created,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}
