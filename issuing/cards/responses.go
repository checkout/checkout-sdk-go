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
		// Replaced activation_date in the 2026-09-02 spec; IssuingActivationDate was removed.
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

	// CardUpdateResponse represents the response body for PATCH /issuing/cards/{cardId}.
	//
	// The 2026-09-02 spec (INT-1695) added encrypted_cvv here; the 2026-09-17 spec (INT-1700)
	// removed it again, so the current spec never includes it. LastModifiedDate stays required
	// (per swagger's allOf against get-card-response).
	CardUpdateResponse struct {
		HttpMetadata common.HttpMetadata

		// ScheduledRevocationDate is the date (format YYYY-MM-DD) on which the
		// card is revoked at midnight UTC.
		ScheduledRevocationDate string `json:"scheduled_revocation_date,omitempty"`

		// LastActivatedOn is the date and time the card was last activated.
		// Nil if the card has never been activated.
		LastActivatedOn *time.Time `json:"last_activated_on,omitempty"`

		// LastModifiedDate is when the card was last modified.
		// [Optional]
		// Format: date-time
		LastModifiedDate *time.Time `json:"last_modified_date,omitempty"`

		// IsSingleUse specifies whether the virtual card is set to expire after a single use.
		// Only present when the underlying card is virtual; physical cards never send it.
		IsSingleUse bool `json:"is_single_use,omitempty"`

		// Links holds the HAL links related to the card.
		// [Optional]
		Links map[string]common.Link `json:"_links,omitempty"`
	}
)
