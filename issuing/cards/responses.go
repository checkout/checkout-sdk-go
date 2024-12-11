package issuing

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
)

type (
	CardResponse struct {
		HttpMetadata    common.HttpMetadata
		Id              string          `json:"id,omitempty"`
		DisplayName     string          `json:"display_name,omitempty"`
		LastFour        string          `json:"last_four,omitempty"`
		ExpiryMonth     int             `json:"expiry_month,omitempty"`
		ExpiryYear      int             `json:"expiry_year,omitempty"`
		BillingCurrency common.Currency `json:"billing_currency,omitempty"`
		IssuingCountry  common.Country  `json:"issuing_country,omitempty"`
		Reference       string          `json:"reference,omitempty"`
		CreatedDate     *time.Time      `json:"created_date,omitempty"`
	}

	CardDetailsData struct {
		HttpMetadata     common.HttpMetadata
		Type             CardType        `json:"type,omitempty"`
		Id               string          `json:"id,omitempty"`
		CardholderId     string          `json:"cardholder_id,omitempty"`
		CardProductId    string          `json:"card_product_id,omitempty"`
		ClientId         string          `json:"client_id,omitempty"`
		LastFour         string          `json:"last_four,omitempty"`
		ExpiryMonth      int             `json:"expiry_month,omitempty"`
		ExpiryYear       int             `json:"expiry_year,omitempty"`
		Status           CardStatus      `json:"status,omitempty"`
		DisplayName      string          `json:"display_name,omitempty"`
		BillingCurrency  common.Currency `json:"billing_currency,omitempty"`
		IssuingCountry   common.Country  `json:"issuing_country,omitempty"`
		Reference        string          `json:"reference,omitempty"`
		CreatedDate      *time.Time      `json:"created_date,omitempty"`
		LastModifiedDate *time.Time      `json:"last_modified_date,omitempty"`
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
