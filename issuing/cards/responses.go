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
		// Example: crd_fa6psq242dcd6fdn5gifcq1491
		Id string `json:"id,omitempty"`

		// ClientId is the client's unique identifier.
		// [Required]
		ClientId string `json:"client_id,omitempty"`

		// EntityId is the entity's unique identifier.
		// [Required]
		EntityId string `json:"entity_id,omitempty"`

		// CardholderId is the cardholder's unique identifier.
		// [Required]
		CardholderId string `json:"cardholder_id,omitempty"`

		// CardProductId is the card product's unique identifier.
		// [Required]
		CardProductId string `json:"card_product_id,omitempty"`

		// UserId is the identifier of the user who issued the card.
		// [Optional]
		UserId string `json:"user_id,omitempty"`

		// LastFour is the last four digits of the card number.
		// [Required]
		// Example: 1234
		LastFour string `json:"last_four,omitempty"`

		// ExpiryMonth is the card's expiration month.
		// [Required]
		// Example: 5
		ExpiryMonth int `json:"expiry_month,omitempty"`

		// ExpiryYear is the card's expiration year.
		// [Required]
		// Example: 2025
		ExpiryYear int `json:"expiry_year,omitempty"`

		// Status is the card's status: active, inactive, revoked or suspended.
		// [Required]
		Status CardStatus `json:"status,omitempty"`

		// DisplayName is the cardholder's name as displayed on the card.
		// [Optional]
		// Example: JOHN KENNEDY
		DisplayName string `json:"display_name,omitempty"`

		// Type is the card type: virtual or physical.
		// [Required]
		Type CardType `json:"type,omitempty"`

		// BillingCurrency is the card's billing currency (ISO 4217).
		// [Required]
		BillingCurrency common.Currency `json:"billing_currency,omitempty"`

		// IssuingCountry is the card's issuing country (ISO 3166 alpha-2).
		// [Required]
		IssuingCountry common.Country `json:"issuing_country,omitempty"`

		// Reference is your reference for the card.
		// [Optional]
		// Example: X-123456-N11
		Reference string `json:"reference,omitempty"`

		// Metadata holds the user's metadata for the card.
		// [Optional]
		Metadata *CardMetadata `json:"metadata,omitempty"`

		// ScheduledRevocationDate is the date (format YYYY-MM-DD) on which the
		// card is revoked at midnight UTC.
		// [Optional]
		ScheduledRevocationDate string `json:"scheduled_revocation_date,omitempty"`

		// RevocationDate is the date on which the card is revoked.
		// [Optional]
		//
		// Deprecated: use ScheduledRevocationDate.
		RevocationDate string `json:"revocation_date,omitempty"`

		// ScheduledActivationDate is the scheduled date of the card's first activation.
		// [Optional]
		// Example: 2026-06-01T10:00Z
		ScheduledActivationDate string `json:"scheduled_activation_date,omitempty"`

		// LastActivatedOn is the date and time the card was last activated.
		// Nil if the card has never been activated.
		// [Optional]
		LastActivatedOn *time.Time `json:"last_activated_on,omitempty"`

		// RootCardId is the identifier of the first card in the renewal chain.
		// [Optional]
		RootCardId string `json:"root_card_id,omitempty"`

		// ParentCardId is the identifier of the card this card was renewed from.
		// [Optional]
		ParentCardId string `json:"parent_card_id,omitempty"`

		// Scheme is the card scheme: mastercard or visa.
		// [Required]
		Scheme CardScheme `json:"scheme,omitempty"`

		// CreatedDate is when the card was created.
		// [Optional]
		// Format: date-time
		CreatedDate *time.Time `json:"created_date,omitempty"`

		// LastModifiedDate is when the card was last modified.
		// [Required]
		// Format: date-time
		LastModifiedDate *time.Time `json:"last_modified_date,omitempty"`

		// IsSingleUse specifies whether the virtual card is set to expire after a single use.
		// Only present when the underlying card is virtual; physical cards never send it.
		// [Optional]
		IsSingleUse bool `json:"is_single_use,omitempty"`

		// Links holds the HAL links related to the card.
		// [Optional]
		Links map[string]common.Link `json:"_links,omitempty"`
	}
)
