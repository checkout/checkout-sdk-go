package issuing

import (
	"github.com/checkout/checkout-sdk-go/v2/common"
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

	CardDetailsRequest struct {
		Type          CardType     `json:"type,omitempty"`
		CardholderId  string       `json:"cardholder_id,omitempty"`
		Lifetime      CardLifetime `json:"lifetime"`
		Reference     string       `json:"reference,omitempty"`
		CardProductId string       `json:"card_product_id,omitempty"`
		DisplayName   string       `json:"display_name,omitempty"`
		ActivateCard  bool         `json:"activate_card,omitempty"`
	}

	physicalCardRequest struct {
		CardDetailsRequest
		ShippingInstructions ShippingInstruction `json:"shipping_instructions,omitempty"`
	}

	virtualCardRequest struct {
		CardDetailsRequest
		IsSingleUse bool `json:"is_single_use,omitempty"`
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
	RenewCardMetadata struct {
		Udf1 string `json:"udf1,omitempty"`
		Udf2 string `json:"udf2,omitempty"`
		Udf3 string `json:"udf3,omitempty"`
		Udf4 string `json:"udf4,omitempty"`
		Udf5 string `json:"udf5,omitempty"`
	}

	renewCardRequestBase struct {
		DisplayName string             `json:"display_name,omitempty"`
		Reference   string             `json:"reference,omitempty"`
		Metadata    *RenewCardMetadata `json:"metadata,omitempty"`
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
	ScheduleRevocationRequest struct {
		RevocationDate string `json:"revocation_date,omitempty"`
	}

	CardUpdateRequest struct {
		Reference   string `json:"reference,omitempty"`
		ExpiryMonth int    `json:"expiry_month,omitempty"`
		ExpiryYear  int    `json:"expiry_year,omitempty"`
	}
)
