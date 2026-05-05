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

// TokenState represents the lifecycle state of a network token.
type TokenState string

const (
	// Active indicates the network token is active and can be used for transactions.
	Active TokenState = "active"
	// Suspended indicates the network token has been temporarily suspended.
	Suspended TokenState = "suspended"
	// Inactive indicates the network token is no longer usable.
	Inactive TokenState = "inactive"
	// Declined indicates the network token provisioning request was declined.
	Declined TokenState = "declined"
	// Requested indicates the network token provisioning request has been submitted but
	// is awaiting a final state.
	Requested TokenState = "requested"
)

// TransactionType describes the channel through which a network token transaction is initiated.
type TransactionType string

const (
	// Ecom represents an e-commerce (card-not-present) transaction.
	Ecom TransactionType = "ecom"
	// Recurring represents a recurring or subscription transaction.
	Recurring TransactionType = "recurring"
	// Pos represents an in-store point-of-sale transaction.
	Pos TransactionType = "pos"
	// Aft represents an account funding transaction.
	Aft TransactionType = "aft"
)

// InitiatedBy identifies who initiated a network token lifecycle event such as deletion.
type InitiatedBy string

const (
	// Cardholder indicates the action was initiated by the cardholder.
	Cardholder InitiatedBy = "cardholder"
	// TokenRequestor indicates the action was initiated by the token requestor (the merchant
	// or wallet provider).
	TokenRequestor InitiatedBy = "token_requestor"
)

// Reason describes why a network token deletion was requested.
type Reason string

const (
	// Fraud indicates the token is being deleted due to suspected or confirmed fraud.
	Fraud Reason = "fraud"
	// Other indicates the token is being deleted for a reason not covered by another constant.
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

// NewCardSource returns a card source with the type field pre-set to the card source type.
func NewCardSource() *cardSource {
	return &cardSource{Type: payments.CardSource}
}

// NewIdSource returns an ID source with the type field pre-set to the ID source type.
func NewIdSource() *idSource {
	return &idSource{Type: payments.IdSource}
}

// Requests

// ProvisionNetworkTokenRequest is the request body for provisioning a new network token.
type ProvisionNetworkTokenRequest struct {
	Source interface{} `json:"source"`
}

// RequestCryptogramRequest is the request body for generating a cryptogram for a network token.
type RequestCryptogramRequest struct {
	TransactionType TransactionType `json:"transaction_type"`
}

// DeleteNetworkTokenRequest is the request body for deleting a network token.
type DeleteNetworkTokenRequest struct {
	InitiatedBy InitiatedBy `json:"initiated_by"`
	Reason      Reason      `json:"reason"`
}

// Response sub-types

// NetworkTokenCard holds the masked card details associated with a network token.
type NetworkTokenCard struct {
	Last4       string `json:"last4"`
	ExpiryMonth string `json:"expiry_month"`
	ExpiryYear  string `json:"expiry_year"`
}

// NetworkTokenDetails holds the full details of a provisioned network token.
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

// NetworkTokenResponse is returned by both ProvisionNetworkToken (POST) and GetNetworkToken
// (GET). Both endpoints share the same response structure; the GET response allows additional
// state values (Declined, Requested) captured by TokenState.
type NetworkTokenResponse struct {
	HttpMetadata     common.HttpMetadata
	Card             *NetworkTokenCard      `json:"card"`
	NetworkToken     *NetworkTokenDetails   `json:"network_token"`
	TokenRequestorId string                 `json:"token_requestor_id,omitempty"`
	TokenSchemeId    string                 `json:"token_scheme_id,omitempty"`
	Links            map[string]common.Link `json:"_links,omitempty"`
}

// RequestCryptogramResponse is the response returned when a cryptogram is generated for a
// network token.
type RequestCryptogramResponse struct {
	HttpMetadata common.HttpMetadata
	Cryptogram   string                 `json:"cryptogram"`
	Eci          string                 `json:"eci,omitempty"`
	Links        map[string]common.Link `json:"_links,omitempty"`
}
