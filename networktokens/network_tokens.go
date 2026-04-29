package networktokens

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/payments"
)

const (
	networkTokensPath = "network-tokens"
	cryptogramsPath   = "cryptograms"
	deletePath        = "delete"
)

type TokenState string

const (
	Active    TokenState = "active"
	Suspended TokenState = "suspended"
	Inactive  TokenState = "inactive"
	Declined  TokenState = "declined"
	Requested TokenState = "requested"
)

type TransactionType string

const (
	Ecom      TransactionType = "ecom"
	Recurring TransactionType = "recurring"
	Pos       TransactionType = "pos"
	Aft       TransactionType = "aft"
)

type InitiatedBy string

const (
	Cardholder     InitiatedBy = "cardholder"
	TokenRequestor InitiatedBy = "token_requestor"
)

type Reason string

const (
	Fraud Reason = "fraud"
	Other Reason = "other"
)

// Source types — unexported so callers must use the constructors, ensuring
// the discriminator type field is always set correctly.

type cardSource struct {
	Type        payments.SourceType `json:"type"`
	Number      string              `json:"number"`
	ExpiryMonth string              `json:"expiry_month"`
	ExpiryYear  string              `json:"expiry_year"`
	Cvv         string              `json:"cvv,omitempty"`
}

type idSource struct {
	Type payments.SourceType `json:"type"`
	Id   string              `json:"id"`
}

func NewCardSource() *cardSource {
	return &cardSource{Type: payments.CardSource}
}

func NewIdSource() *idSource {
	return &idSource{Type: payments.IdSource}
}

// Requests

type ProvisionNetworkTokenRequest struct {
	Source interface{} `json:"source"`
}

type RequestCryptogramRequest struct {
	TransactionType TransactionType `json:"transaction_type"`
}

type DeleteNetworkTokenRequest struct {
	InitiatedBy InitiatedBy `json:"initiated_by"`
	Reason      Reason      `json:"reason"`
}

// Response sub-types

type NetworkTokenCard struct {
	Last4       string `json:"last4"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
}

type NetworkTokenDetails struct {
	Id                      string                    `json:"id"`
	State                   TokenState                `json:"state"`
	Number                  string                    `json:"number,omitempty"`
	ExpiryMonth             string                    `json:"expiry_month,omitempty"`
	ExpiryYear              string                    `json:"expiry_year,omitempty"`
	Type                    payments.NetworkTokenType `json:"type"`
	PaymentAccountReference string                    `json:"payment_account_reference,omitempty"`
	CreatedOn               *time.Time                `json:"created_on"`
	ModifiedOn              *time.Time                `json:"modified_on"`
}

// Responses
// NetworkTokenResponse is used for both ProvisionNetworkToken (POST) and GetNetworkToken (GET).
// Both endpoints share the same structure; the GET response allows additional state values
// (declined, requested) which are captured by the TokenState type.

type NetworkTokenResponse struct {
	HttpMetadata     common.HttpMetadata
	Card             *NetworkTokenCard      `json:"card"`
	NetworkToken     *NetworkTokenDetails   `json:"network_token"`
	TokenRequestorId string                 `json:"token_requestor_id,omitempty"`
	TokenSchemeId    string                 `json:"token_scheme_id,omitempty"`
	Links            map[string]common.Link `json:"_links,omitempty"`
}

type RequestCryptogramResponse struct {
	HttpMetadata common.HttpMetadata
	Cryptogram   string                 `json:"cryptogram"`
	Eci          string                 `json:"eci,omitempty"`
	Links        map[string]common.Link `json:"_links,omitempty"`
}
