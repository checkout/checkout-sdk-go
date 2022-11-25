package abc

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/payments"
)

type (
	RequestCardDestination struct {
		Type           payments.PaymentDestinationType `json:"type,omitempty"`
		Number         string                          `json:"number,omitempty"`
		ExpiryMonth    int                             `json:"expiry_month,omitempty"`
		ExpiryYear     int                             `json:"expiry_year,omitempty"`
		FirstName      string                          `json:"first_name,omitempty"`
		LastName       string                          `json:"last_name,omitempty"`
		Name           string                          `json:"name,omitempty"`
		BillingAddress *common.Address                 `json:"billing_address,omitempty"`
		Phone          *common.Phone                   `json:"phone,omitempty"`
	}

	RequestIdDestination struct {
		Type      payments.PaymentDestinationType `json:"type,omitempty"`
		Id        string                          `json:"id,omitempty"`
		FirstName string                          `json:"first_name,omitempty"`
		LastName  string                          `json:"last_name,omitempty"`
	}

	RequestTokenDestination struct {
		Type           payments.PaymentDestinationType `json:"type,omitempty"`
		Token          string                          `json:"token,omitempty"`
		FirstName      string                          `json:"first_name,omitempty"`
		LastName       string                          `json:"last_name,omitempty"`
		BillingAddress *common.Address                 `json:"billing_address,omitempty"`
		Phone          *common.Phone                   `json:"phone,omitempty"`
	}
)

func NewRequestCardDestination() *RequestCardDestination {
	return &RequestCardDestination{Type: payments.CardDestination}
}

func NewRequestIdDestination() *RequestIdDestination {
	return &RequestIdDestination{Type: payments.IdDestination}
}

func NewRequestTokenDestination() *RequestTokenDestination {
	return &RequestTokenDestination{Type: payments.TokenDestination}
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
