package issuing

import (
	"github.com/checkout/checkout-sdk-go/v3/common"
)

type LifetimeUnit string

const (
	Months LifetimeUnit = "Months"
	Years  LifetimeUnit = "Years"
)

type (
	CardCredentialsQuery struct {
		Credentials string `json:"credentials,omitempty"`
	}
)

type RevokeReason string

const (
	Expired        RevokeReason = "expired"
	ReportedLost   RevokeReason = "reported_lost"
	ReportedStolen RevokeReason = "reported_stolen"
)

type (
	// RevokeCardRequest is the optional body for POST /issuing/cards/{cardId}/revoke. The body is
	// not required: pass RevokeCardRequest{} to revoke without a reason, which sends {}.
	RevokeCardRequest struct {
		// Reason is the reason for revoking the card.
		// [Optional]
		// Enum: expired, reported_lost, reported_stolen
		Reason RevokeReason `json:"reason,omitempty"`
	}
)

type SuspendReason string

const (
	SuspectedLost   SuspendReason = "suspected_lost"
	SuspectedStolen SuspendReason = "suspected_stolen"
)

type (
	// SuspendCardRequest is the optional body for POST /issuing/cards/{cardId}/suspend. The body is
	// not required: pass SuspendCardRequest{} to suspend without a reason, which sends {}.
	SuspendCardRequest struct {
		// Reason is the reason for suspending the card.
		// [Optional]
		// Enum: suspected_lost, suspected_stolen
		Reason SuspendReason `json:"reason,omitempty"`
	}
)

type (
	CardLifetime struct {
		Unit  LifetimeUnit `json:"unit,omitempty"`
		Value int          `json:"value,omitempty"`
	}

	ShippingInstruction struct {
		ShippingRecipient string          `json:"shipping_recipient,omitempty"`
		ShippingAddress   *common.Address `json:"shipping_address,omitempty"`
		AdditionalComment string          `json:"additional_comment,omitempty"`
	}

	CardRequest interface {
		GetRequestType() CardType
	}

	CardTypeRequest struct {
		CardRequest
	}

	CardMetadata struct {
		Udf1 string `json:"udf1,omitempty"`
		Udf2 string `json:"udf2,omitempty"`
		Udf3 string `json:"udf3,omitempty"`
		Udf4 string `json:"udf4,omitempty"`
		Udf5 string `json:"udf5,omitempty"`
	}

	CardDetailsRequest struct {
		// Type is the card type.
		// [Required]
		// Enum: "virtual" "physical"
		Type CardType `json:"type,omitempty"`
		// CardholderId is the cardholder's unique identifier.
		// [Required]
		// ^crh_[a-z0-9]{26}$
		CardholderId string `json:"cardholder_id,omitempty"`
		// Lifetime is how long the card accepts incoming authorizations.
		// [Optional]
		Lifetime CardLifetime `json:"lifetime"`
		// Reference is your reference.
		// [Optional]
		// max 256 characters
		// Example: X-123456-N11
		Reference string `json:"reference,omitempty"`
		// CardProductId is the card product's unique identifier.
		// [Required]
		CardProductId string `json:"card_product_id,omitempty"`
		// DisplayName is the name to display on the card.
		// [Optional]
		// min 2 characters, max 26 characters
		// Example: JOHN KENNEDY
		DisplayName string `json:"display_name,omitempty"`
		// ActivateCard sets whether to activate the card on creation.
		// [Optional]
		// Default: true
		ActivateCard bool `json:"activate_card,omitempty"`
		// Metadata is the user's metadata.
		// [Optional]
		Metadata *CardMetadata `json:"metadata,omitempty"`

		// RevocationDate is deprecated. Use ScheduledRevocationDate instead.
		//
		// Deprecated: use ScheduledRevocationDate. If both fields are provided,
		// ScheduledRevocationDate overrides this value.
		RevocationDate string `json:"revocation_date,omitempty"`

		// ScheduledRevocationDate is the date (format YYYY-MM-DD) on which the
		// card is revoked at midnight UTC. Replaces the deprecated RevocationDate.
		ScheduledRevocationDate string `json:"scheduled_revocation_date,omitempty"`

		// ScheduledActivationDate schedules the card's first activation. Only applies to the
		// initial activation of a card. Two formats are supported: date only (yyyy-MM-dd,
		// treated as midnight UTC), or date with round hour (yyyy-MM-ddTHH:mmZ in UTC, or
		// yyyy-MM-ddTHH:mm+HH:mm with offset). Only round hours are allowed when a time is
		// provided (HH:00). The value must be at least the next round hour after the request
		// time.
		// [Optional]
		// Example: 2026-06-01T10:00Z
		ScheduledActivationDate string `json:"scheduled_activation_date,omitempty"`
	}

	physicalCardRequest struct {
		CardDetailsRequest
		ShippingInstructions ShippingInstruction `json:"shipping_instructions,omitempty"`
	}

	virtualCardRequest struct {
		CardDetailsRequest
		IsSingleUse       bool     `json:"is_single_use,omitempty"`
		ReturnCredentials []string `json:"return_credentials,omitempty"`
		ControlProfiles   []string `json:"control_profiles,omitempty"`
	}
)

func NewPhysicalCardRequest() *physicalCardRequest {
	return &physicalCardRequest{
		CardDetailsRequest: CardDetailsRequest{Type: Physical},
	}
}

func NewVirtualCardRequest() *virtualCardRequest {
	return &virtualCardRequest{
		CardDetailsRequest: CardDetailsRequest{Type: Virtual},
	}
}

func (c *physicalCardRequest) GetRequestType() CardType {
	return c.Type
}

func (c *virtualCardRequest) GetRequestType() CardType {
	return c.Type
}

type (
	RenewCardMetadata = CardMetadata

	renewCardRequestBase struct {
		DisplayName string        `json:"display_name,omitempty"`
		Reference   string        `json:"reference,omitempty"`
		Metadata    *CardMetadata `json:"metadata,omitempty"`
	}

	physicalCardRenewRequest struct {
		renewCardRequestBase
		ShippingInstructions *ShippingInstruction `json:"shipping_instructions,omitempty"`
	}

	virtualCardRenewRequest struct {
		renewCardRequestBase
	}

	RenewCardRequest interface {
		GetRenewType() CardType
	}
)

func NewPhysicalCardRenewRequest() *physicalCardRenewRequest {
	return &physicalCardRenewRequest{}
}

func NewVirtualCardRenewRequest() *virtualCardRenewRequest {
	return &virtualCardRenewRequest{}
}

func (c *physicalCardRenewRequest) GetRenewType() CardType {
	return Physical
}

func (c *virtualCardRenewRequest) GetRenewType() CardType {
	return Virtual
}

// CardStatusUpdate is the only status value accepted by CardUpdateRequest.Status,
// used to reactivate an inactive or suspended card.
type CardStatusUpdate string

const (
	ActiveCardStatusUpdate CardStatusUpdate = "active"
)

type (
	// CardUpdateRequest represents the request body for PATCH /issuing/cards/{cardId}.
	CardUpdateRequest struct {
		// Reference is your reference.
		// [Optional]
		// max 256 characters
		// Example: X-123456-N11
		Reference string `json:"reference,omitempty"`

		// Metadata is the user's metadata.
		// [Optional]
		Metadata *CardMetadata `json:"metadata,omitempty"`

		// ExpiryMonth is the card's expiration month.
		// [Optional]
		// min 1, max 12
		// Example: 5
		ExpiryMonth int `json:"expiry_month,omitempty"`

		// ExpiryYear is the card's expiration year.
		// [Optional]
		// min 4 characters, max 4 characters
		// Example: 2025
		ExpiryYear int `json:"expiry_year,omitempty"`

		// Status can be set to "active" to reactivate an inactive or suspended
		// card. Mutually exclusive with the scheduled_activation_date field:
		// submitting both results in a scheduled_activation_date_conflicts_with_activation error.
		// [Optional]
		// Enum: "active"
		// Example: active
		Status CardStatusUpdate `json:"status,omitempty"`

		// ScheduledActivationDate schedules the card's first activation. Only applies to the
		// initial activation of a card. Two formats are supported: date only (yyyy-MM-dd,
		// treated as midnight UTC), or date with round hour (yyyy-MM-ddTHH:mmZ in UTC, or
		// yyyy-MM-ddTHH:mm+HH:mm with offset). Only round hours are allowed when a time is
		// provided (HH:00). The value must be at least the next round hour after the request
		// time.
		// [Optional]
		// Example: 2026-06-01T10:00Z
		ScheduledActivationDate string `json:"scheduled_activation_date,omitempty"`

		// RevocationDate is the date scheduling the card's automatic revocation (time is
		// midnight UTC). If both fields are provided, ScheduledRevocationDate overrides this value.
		// [Optional]
		// Format: date (YYYY-MM-DD)
		//
		// Deprecated: use ScheduledRevocationDate. If both fields are provided,
		// ScheduledRevocationDate overrides this value.
		RevocationDate string `json:"revocation_date,omitempty"`

		// ScheduledRevocationDate is the date on which the card is revoked, at midnight UTC.
		// Replaces the deprecated RevocationDate.
		// [Optional]
		// Format: date (YYYY-MM-DD)
		ScheduledRevocationDate string `json:"scheduled_revocation_date,omitempty"`
	}

	// CardUpdateHeaders are the HTTP headers the SDK can send when updating a card's details.
	// The API no longer documents them and ignores them (verified against the sandbox). Headers
	// that are set are still sent unchanged.
	//
	// Deprecated: The API no longer supports these headers; use UpdateCard without headers.
	// Will be removed in the next major version.
	//
	// Go canonicalises outgoing header names, so return-encrypted-cvv reaches the wire as
	// Return-Encrypted-Cvv no matter how the tag is spelled. That is correct and unavoidable:
	// net/http applies textproto.CanonicalMIMEHeaderKey on write, and HTTP header names are
	// case insensitive per RFC 7230, so the API reads it either way. Do not try to "fix" the
	// casing; it cannot be done through net/http.
	CardUpdateHeaders struct {
		// ReturnEncryptedCvv is sent as the return-encrypted-cvv header when set. The API no
		// longer documents this header and ignores it; the response never includes an
		// encrypted CVV.
		// [Optional]
		//
		// Deprecated: The API no longer supports this header. Will be removed in the next
		// major version.
		// Maps to HTTP header return-encrypted-cvv.
		// Example: "true"
		ReturnEncryptedCvv string `json:"return-encrypted-cvv,omitempty"`

		// EncryptionKey is sent as the Encryption-Key header when set. The API no longer
		// documents this header and ignores it.
		// [Optional]
		//
		// Deprecated: The API no longer supports this header. Will be removed in the next
		// major version.
		// Maps to HTTP header Encryption-Key.
		EncryptionKey string `json:"Encryption-Key,omitempty"`
	}
)
