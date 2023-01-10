package abc

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	requestCardDestination struct {
		Type           payments.DestinationType `json:"type,omitempty"`
		Number         string                   `json:"number,omitempty"`
		ExpiryMonth    int                      `json:"expiry_month,omitempty"`
		ExpiryYear     int                      `json:"expiry_year,omitempty"`
		FirstName      string                   `json:"first_name,omitempty"`
		LastName       string                   `json:"last_name,omitempty"`
		Name           string                   `json:"name,omitempty"`
		BillingAddress *common.Address          `json:"billing_address,omitempty"`
		Phone          *common.Phone            `json:"phone,omitempty"`
	}

	requestIdDestination struct {
		Type      payments.DestinationType `json:"type,omitempty"`
		Id        string                   `json:"id,omitempty"`
		FirstName string                   `json:"first_name,omitempty"`
		LastName  string                   `json:"last_name,omitempty"`
	}

	requestTokenDestination struct {
		Type           payments.DestinationType `json:"type,omitempty"`
		Token          string                   `json:"token,omitempty"`
		FirstName      string                   `json:"first_name,omitempty"`
		LastName       string                   `json:"last_name,omitempty"`
		BillingAddress *common.Address          `json:"billing_address,omitempty"`
		Phone          *common.Phone            `json:"phone,omitempty"`
	}
)

func NewRequestCardDestination() *requestCardDestination {
	return &requestCardDestination{Type: payments.CardDestination}
}

func NewRequestIdDestination() *requestIdDestination {
	return &requestIdDestination{Type: payments.IdDestination}
}

func NewRequestTokenDestination() *requestTokenDestination {
	return &requestTokenDestination{Type: payments.TokenDestination}
}

func (d *requestCardDestination) GetType() payments.DestinationType {
	return d.Type
}

func (d *requestIdDestination) GetType() payments.DestinationType {
	return d.Type
}

func (d *requestTokenDestination) GetType() payments.DestinationType {
	return d.Type
}

type (
	DestinationResponse struct {
		*ResponseCardDestination
		*common.AlternativeResponse
	}

	ResponseCardDestination struct {
		ExpiryMonth   int                 `json:"expiry_month,omitempty"`
		ExpiryYear    int                 `json:"expiry_year,omitempty"`
		Name          string              `json:"name,omitempty"`
		Last4         string              `json:"last4,omitempty"`
		Fingerprint   string              `json:"fingerprint,omitempty"`
		Bin           string              `json:"bin,omitempty"`
		CardType      common.CardType     `json:"card_type,omitempty"`
		CardCategory  common.CardCategory `json:"card_category,omitempty"`
		Issuer        string              `json:"issuer,omitempty"`
		IssuerCountry common.Country      `json:"issuer_country,omitempty"`
		ProductId     string              `json:"product_id,omitempty"`
		ProductType   string              `json:"product_type,omitempty"`
	}
)

func (s *DestinationResponse) UnmarshalJSON(data []byte) error {
	var typeMapping payments.DestinationTypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Destination {
	case string(payments.CardDestination):
		var typeMapping ResponseCardDestination
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.ResponseCardDestination = &typeMapping
	default:
		var typeMapping common.AlternativeResponse
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.AlternativeResponse = &typeMapping
	}

	return nil
}
