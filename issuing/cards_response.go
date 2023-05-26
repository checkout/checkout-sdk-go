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

	CardDetailsResponse struct {
		HttpMetadata         common.HttpMetadata
		PhysicalCardResponse *PhysicalCardResponse
		VirtualCardResponse  *VirtualCardResponse
	}

	PhysicalCardResponse struct {
		Type             CardType        `json:"type" binding:"required"`
		Id               string          `json:"id,omitempty"`
		CardholderId     string          `json:"cardholder_id,omitempty"`
		CardProductId    string          `json:"card_product_id,omitempty"`
		ClientId         string          `json:"client_id,omitempty"`
		LastFour         string          `json:"last_four,omitempty"`
		ExpiryMonth      int             `json:"expiry_month,omitempty"`
		Status           CardStatus      `json:"status,omitempty"`
		DisplayName      string          `json:"display_name,omitempty"`
		BillingCurrency  common.Currency `json:"billing_currency,omitempty"`
		IssuingCountry   common.Country  `json:"issuing_country,omitempty"`
		Reference        string          `json:"reference,omitempty"`
		CreatedDate      *time.Time      `json:"created_date,omitempty"`
		LastModifiedDate *time.Time      `json:"last_modified_date,omitempty"`
	}

	VirtualCardResponse struct {
		Type             CardType        `json:"type" binding:"required"`
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
		IsSingleUse      bool            `json:"is_single_use,omitempty"`
	}
)

func (s *CardDetailsResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(Physical):
		var response PhysicalCardResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.PhysicalCardResponse = &response
	case string(Virtual):
		var response VirtualCardResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.VirtualCardResponse = &response
	default:
		return errors.UnsupportedTypeError(fmt.Sprintf("%s unsupported", typeMapping.Type))
	}
	return nil
}

type (
	CardCredentialsResponse struct {
		Number string `json:"number,omitempty"`
		Cvc2   string `json:"cvc_2,omitempty"`
	}
)
