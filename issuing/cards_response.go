package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

type (
	CardTypeResponse interface {
		GetType() CardType
		GetDetails() interface{}
	}

	CardDetailsResponse struct {
		CardTypeResponse
	}

	CardDetailsCardholder struct {
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

	physicalCardTypeResponse struct {
		CardDetailsCardholder
	}

	virtualCardTypeResponse struct {
		CardDetailsCardholder
		IsSingleUse bool `json:"is_single_use,omitempty"`
	}
)

func NewPhysicalCardTypeResponse() *physicalCardTypeResponse {
	return &physicalCardTypeResponse{
		CardDetailsCardholder: CardDetailsCardholder{Type: Physical},
	}
}

func NewVirtualCardTypeResponse() *virtualCardTypeResponse {
	return &virtualCardTypeResponse{
		CardDetailsCardholder: CardDetailsCardholder{Type: Virtual},
	}
}

func (c *physicalCardTypeResponse) GetType() CardType {
	return c.Type
}

func (c *virtualCardTypeResponse) GetType() CardType {
	return c.Type
}

func (c *physicalCardTypeResponse) GetDetails() interface{} {
	return physicalCardTypeResponse{
		CardDetailsCardholder: c.CardDetailsCardholder,
	}
}

func (c *virtualCardTypeResponse) GetDetails() interface{} {
	return virtualCardTypeResponse{
		CardDetailsCardholder: c.CardDetailsCardholder,
		IsSingleUse:           c.IsSingleUse,
	}
}
