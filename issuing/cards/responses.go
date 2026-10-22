package issuing

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

type (
	CardResponse struct {
		HttpMetadata    common.HttpMetadata
		Id              string                 `json:"id,omitempty"`
		DisplayName     string                 `json:"display_name,omitempty"`
		LastFour        string                 `json:"last_four,omitempty"`
		ExpiryMonth     int                    `json:"expiry_month,omitempty"`
		ExpiryYear      int                    `json:"expiry_year,omitempty"`
		BillingCurrency common.Currency        `json:"billing_currency,omitempty"`
		IssuingCountry  common.Country         `json:"issuing_country,omitempty"`
		Reference       string                 `json:"reference,omitempty"`
		CreatedDate     *time.Time             `json:"created_date,omitempty"`
		Links           map[string]common.Link `json:"_links,omitempty"`
	}

	CardDetailsData struct {
		HttpMetadata    common.HttpMetadata
		Type            CardType        `json:"type,omitempty"`
		Id              string          `json:"id,omitempty"`
		CardholderId    string          `json:"cardholder_id,omitempty"`
		CardProductId   string          `json:"card_product_id,omitempty"`
		ClientId        string          `json:"client_id,omitempty"`
		EntityId        string          `json:"entity_id,omitempty"`
		UserId          string          `json:"user_id,omitempty"`
		LastFour        string          `json:"last_four,omitempty"`
		ExpiryMonth     int             `json:"expiry_month,omitempty"`
		ExpiryYear      int             `json:"expiry_year,omitempty"`
		Status          CardStatus      `json:"status,omitempty"`
		DisplayName     string          `json:"display_name,omitempty"`
		BillingCurrency common.Currency `json:"billing_currency,omitempty"`
		IssuingCountry  common.Country  `json:"issuing_country,omitempty"`
		Scheme          CardScheme      `json:"scheme,omitempty"`
		Reference       string          `json:"reference,omitempty"`
		Metadata        *CardMetadata   `json:"metadata,omitempty"`

		// RevocationDate is deprecated. Use ScheduledRevocationDate instead.
		//
		// Deprecated: use ScheduledRevocationDate.
		RevocationDate string `json:"revocation_date,omitempty"`

		// ScheduledRevocationDate is the date (format YYYY-MM-DD) on which the
		// card is revoked at midnight UTC. Replaces the deprecated RevocationDate.
		ScheduledRevocationDate string `json:"scheduled_revocation_date,omitempty"`

		// ScheduledActivationDate is the scheduled date of the card's first activation.
		// Replaces the former activation_date field.
		// [Optional]
		// Example: 2026-06-01T10:00Z
		ScheduledActivationDate string     `json:"scheduled_activation_date,omitempty"`
		RootCardId              string     `json:"root_card_id,omitempty"`
		ParentCardId            string     `json:"parent_card_id,omitempty"`
		CreatedDate             *time.Time `json:"created_date,omitempty"`

		// LastActivatedOn is the date and time the card was last activated.
		// Nil if the card has never been activated.
		LastActivatedOn  *time.Time             `json:"last_activated_on,omitempty"`
		LastModifiedDate *time.Time             `json:"last_modified_date,omitempty"`
		Links            map[string]common.Link `json:"_links,omitempty"`
	}

	ActivateCardResponse struct {
		HttpMetadata common.HttpMetadata

		// LastActivatedOn is the time the card was activated. Required on this
		// response (unlike the nullable LastActivatedOn on CardDetailsData).
		LastActivatedOn *time.Time             `json:"last_activated_on,omitempty"`
		Links           map[string]common.Link `json:"_links,omitempty"`
	}

	RevokeCardResponse struct {
		HttpMetadata common.HttpMetadata
		Links        map[string]common.Link `json:"_links,omitempty"`
	}

	SuspendCardResponse struct {
		HttpMetadata common.HttpMetadata
		Links        map[string]common.Link `json:"_links,omitempty"`
	}

	VirtualExtraData struct {
		IsSingleUse bool `json:"is_single_use,omitempty"`
	}

	CardDetailsResponse struct {
		CardDetailsData
		ExtraData ExtraData `json:"limit,omitempty"`
	}

	ExtraData interface {
		GetResponseType() CardType
	}
)

func (l VirtualExtraData) GetResponseType() CardType {
	return Virtual
}

func (s *CardDetailsResponse) UnmarshalJSON(data []byte) error {
	var cardDetailsData CardDetailsData
	if err := json.Unmarshal(data, &cardDetailsData); err != nil {
		return err
	}
	s.CardDetailsData = cardDetailsData

	switch cardDetailsData.Type {
	case Physical:
		s.ExtraData = nil
	case Virtual:
		var extraData = struct {
			VirtualExtraData
		}{}
		if err := json.Unmarshal(data, &extraData); err != nil {
			return nil
		}
		s.ExtraData = extraData.VirtualExtraData
	default:
		return errors.UnsupportedTypeError(fmt.Sprintf("%s unsupported", cardDetailsData.Type))
	}
	return nil
}

type (
	CardCredentialsResponse struct {
		Number string `json:"number,omitempty"`
		Cvc2   string `json:"cvc2,omitempty"`
	}
)

type (
	RenewCardResponse struct {
		HttpMetadata    common.HttpMetadata
		Id              string                 `json:"id,omitempty"`
		ParentCardId    string                 `json:"parent_card_id,omitempty"`
		CardholderId    string                 `json:"cardholder_id,omitempty"`
		Status          CardStatus             `json:"status,omitempty"`
		Type            CardType               `json:"type,omitempty"`
		ClientId        string                 `json:"client_id,omitempty"`
		EntityId        string                 `json:"entity_id,omitempty"`
		LastFour        string                 `json:"last_four,omitempty"`
		ExpiryYear      int                    `json:"expiry_year,omitempty"`
		ExpiryMonth     int                    `json:"expiry_month,omitempty"`
		DisplayName     string                 `json:"display_name,omitempty"`
		Reference       string                 `json:"reference,omitempty"`
		CreatedDate     *time.Time             `json:"created_date,omitempty"`
		BillingCurrency common.Currency        `json:"billing_currency,omitempty"`
		IssuingCountry  common.Country         `json:"issuing_country,omitempty"`
		Links           map[string]common.Link `json:"_links,omitempty"`
	}

	// CardUpdateResponse represents the response body for PATCH /issuing/cards/{cardId}
	// (update-card-response). The response never includes encrypted_cvv. IsSingleUse comes from
	// the virtual variant (update-virtual-card-response) and is absent for physical cards.
	CardUpdateResponse struct {
		HttpMetadata common.HttpMetadata

		// Id is the card's unique identifier.
		// [Required]
		// Pattern: ^crd_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		// Example: crd_fa6psq242dcd6fdn5gifcq1491
		Id string `json:"id,omitempty"`

		// ClientId is the client's unique identifier.
		// [Required]
		// Pattern: ^cli_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		ClientId string `json:"client_id,omitempty"`

		// EntityId is the entity's unique identifier.
		// [Required]
		// Pattern: ^ent_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		EntityId string `json:"entity_id,omitempty"`

		// CardholderId is the cardholder's unique identifier.
		// [Required]
		// Pattern: ^crh_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		CardholderId string `json:"cardholder_id,omitempty"`

		// CardProductId is the card product's unique identifier.
		// [Required]
		// Pattern: ^pro_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		CardProductId string `json:"card_product_id,omitempty"`

		// UserId is the Dashboard user's unique identifier.
		// [Optional]
		// Pattern: ^usr_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		UserId string `json:"user_id,omitempty"`

		// LastFour is the last four digits of the card number, also known as the PAN.
		// [Required]
		// Pattern: ^[0-9]{4}$
		// min 4, max 4
		// Example: 1234
		LastFour string `json:"last_four,omitempty"`

		// ExpiryMonth is the card's expiration month.
		// [Required]
		// Format: int32
		// min 1, max 12
		// Example: 5
		ExpiryMonth int `json:"expiry_month,omitempty"`

		// ExpiryYear is the card's expiration year.
		// [Required]
		// Format: int32
		// min 4 characters, max 4 characters
		// Example: 2025
		ExpiryYear int `json:"expiry_year,omitempty"`

		// Status is the card's status, which determines whether it can approve incoming authorizations:
		// inactive (the default; authorizations are declined until the card is activated), active
		// (authorizations can be approved), suspended (authorizations are declined; the card can be
		// reactivated) or revoked (authorizations are permanently declined; the card cannot be reactivated).
		// [Required]
		// Enum: "active" "inactive" "revoked" "suspended"
		Status CardStatus `json:"status,omitempty"`

		// DisplayName is the name to display on the card.
		// [Optional]
		// Pattern: ^[0-9a-zA-Z.\- ]{2,26}$
		// min 2 characters, max 26 characters
		// Example: JOHN KENNEDY
		DisplayName string `json:"display_name,omitempty"`

		// Type is the card type: virtual or physical. Discriminates update-virtual-card-response from
		// update-physical-card-response.
		// [Required]
		Type CardType `json:"type,omitempty"`

		// BillingCurrency is the issuing currency, as a three-letter ISO currency code.
		// [Required]
		// Format: ISO4217
		// Pattern: ^[a-zA-Z]{3}$
		// min 3 characters, max 3 characters
		// Example: USD
		BillingCurrency common.Currency `json:"billing_currency,omitempty"`

		// IssuingCountry is the issuing country, as a two-letter ISO country code.
		// [Required]
		// min 2 characters, max 2 characters
		// Example: US
		IssuingCountry common.Country `json:"issuing_country,omitempty"`

		// Reference is your reference.
		// [Optional]
		// max 256 characters
		// Example: X-123456-N11
		Reference string `json:"reference,omitempty"`

		// Metadata is user's metadata.
		// [Optional]
		Metadata *CardMetadata `json:"metadata,omitempty"`

		// ScheduledRevocationDate is the date on which the card is revoked, at midnight UTC.
		// [Optional]
		// Format: date (YYYY-MM-DD)
		ScheduledRevocationDate string `json:"scheduled_revocation_date,omitempty"`

		// RevocationDate is the date scheduling the card's automatic revocation (time is midnight UTC).
		// [Optional]
		// Format: date (YYYY-MM-DD)
		//
		// Deprecated: use ScheduledRevocationDate.
		RevocationDate string `json:"revocation_date,omitempty"`

		// ScheduledActivationDate is date scheduling the card's first activation. Only applies to the initial activation of a card.
		// Two formats are supported: date only (YYYY-MM-DD, treated as midnight UTC), or date with round
		// hour (YYYY-MM-DDTHH:mmZ in UTC, or YYYY-MM-DDTHH:mm+HH:mm with offset). Only round hours are
		// allowed when a time is provided (HH:00).
		// [Optional]
		// Example: 2026-06-01T10:00Z
		ScheduledActivationDate string `json:"scheduled_activation_date,omitempty"`

		// LastActivatedOn is the date and time the card was last activated. Nil if the card has never been activated.
		// [Optional]
		// Format: date-time
		// Nullable, read only.
		LastActivatedOn *time.Time `json:"last_activated_on,omitempty"`

		// RootCardId is, when a card is renewed, the unique identifier of the first card in the
		// renewal history.
		// [Optional]
		// Pattern: ^crd_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		RootCardId string `json:"root_card_id,omitempty"`

		// ParentCardId is, when a card is renewed, the unique identifier of the previous card in the
		// renewal history.
		// [Optional]
		// Pattern: ^crd_[a-z0-9]{26}$
		// min 30 characters, max 30 characters
		ParentCardId string `json:"parent_card_id,omitempty"`

		// Scheme is the card scheme.
		// [Required]
		// Enum: "mastercard" "visa"
		Scheme CardScheme `json:"scheme,omitempty"`

		// CreatedDate is the date and time the card was created.
		// [Optional]
		// Format: UTC date-time
		// Example: 2021-09-09T19:41:39Z
		CreatedDate *time.Time `json:"created_date,omitempty"`

		// LastModifiedDate is the date and time the card was last modified.
		// [Required]
		// Format: UTC date-time
		// Example: 2021-09-09T19:41:39Z
		LastModifiedDate *time.Time `json:"last_modified_date,omitempty"`

		// IsSingleUse specifies whether the virtual card is set to expire after a single use. Only
		// sent for virtual cards (update-virtual-card-response); physical cards never send it.
		// [Optional]
		// Default: false
		IsSingleUse bool `json:"is_single_use,omitempty"`

		// Links holds the links related to the card: self, credentials, revoke and controls.
		// [Optional]
		// min 1 item
		Links map[string]common.Link `json:"_links,omitempty"`
	}
)
