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
	RevokeCardRequest struct {
		Reason RevokeReason `json:"reason,omitempty"`
	}
)

type SuspendReason string

const (
	SuspectedLost   SuspendReason = "suspected_lost"
	SuspectedStolen SuspendReason = "suspected_stolen"
)

type (
	SuspendCardRequest struct {
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
		Type          CardType      `json:"type,omitempty"`
		CardholderId  string        `json:"cardholder_id,omitempty"`
		Lifetime      CardLifetime  `json:"lifetime"`
		Reference     string        `json:"reference,omitempty"`
		CardProductId string        `json:"card_product_id,omitempty"`
		DisplayName   string        `json:"display_name,omitempty"`
		ActivateCard  bool          `json:"activate_card,omitempty"`
		Metadata      *CardMetadata `json:"metadata,omitempty"`

		// RevocationDate schedules the card's automatic revocation.
		// [Optional]
		// Format: yyyy-MM-dd (time is midnight UTC)
		// Example: 2027-03-12
		RevocationDate string `json:"revocation_date,omitempty"`

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

type (
	// CardUpdateRequest represents the request body for PATCH /issuing/cards/{cardId}.
	CardUpdateRequest struct {
		// Reference is your reference.
		// [Optional]
		// max 256 characters
		// Example: X-123456-N11
		Reference string `json:"reference,omitempty"`

		// Metadata is the user's metadata. Declared by update-card-request but absent from this
		// struct until the 2026-09-02 pass, so it could not be sent at all.
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

		// ScheduledActivationDate schedules the card's first activation. Only applies to the
		// initial activation of a card. Two formats are supported: date only (yyyy-MM-dd,
		// treated as midnight UTC), or date with round hour (yyyy-MM-ddTHH:mmZ in UTC, or
		// yyyy-MM-ddTHH:mm+HH:mm with offset). Only round hours are allowed when a time is
		// provided (HH:00). The value must be at least the next round hour after the request
		// time.
		// [Optional]
		// Example: 2026-06-01T10:00Z
		ScheduledActivationDate string `json:"scheduled_activation_date,omitempty"`

		// RevocationDate schedules the card's automatic revocation.
		// [Optional]
		// Format: yyyy-MM-dd (time is midnight UTC)
		// Example: 2027-03-12
		RevocationDate string `json:"revocation_date,omitempty"`
	}

	// CardUpdateHeaders are the optional HTTP headers accepted when updating a card's details.
	//
	// Go canonicalises outgoing header names, so return-encrypted-cvv reaches the wire as
	// Return-Encrypted-Cvv no matter how the tag is spelled. That is correct and unavoidable:
	// net/http applies textproto.CanonicalMIMEHeaderKey on write, and HTTP header names are
	// case insensitive per RFC 7230, so the API reads it either way. Do not try to "fix" the
	// casing; it cannot be done through net/http.
	CardUpdateHeaders struct {
		// ReturnEncryptedCvv set to "true" retrieves the card's encrypted credentials in the
		// response. Requires an RSA public key in the Encryption-Key header, otherwise the API
		// answers 422 with error code encryption_key_required.
		// [Optional]
		// Maps to HTTP header return-encrypted-cvv.
		// Example: "true"
		ReturnEncryptedCvv string `json:"return-encrypted-cvv,omitempty"`

		// EncryptionKey is the RSA public key used to encrypt returned credentials. Required when
		// ReturnEncryptedCvv is set. Provide the public key with the BEGIN PUBLIC KEY and END
		// PUBLIC KEY headers and any newline characters removed, encoded as Base64.
		// [Optional]
		// Maps to HTTP header Encryption-Key.
		EncryptionKey string `json:"Encryption-Key,omitempty"`
	}
)
