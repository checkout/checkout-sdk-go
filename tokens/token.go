package tokens

import (
	"encoding/json"
	"time"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/common"
)

type (
	// Request -
	Request struct {
		*Card
		*Wallet
	}

	// Card -
	Card struct {
		Type           common.TokenType `json:"type" binding:"required"`
		Number         string           `json:"number" binding:"required"`
		ExpiryMonth    uint64           `json:"expiry_month" binding:"required"`
		ExpiryYear     uint64           `json:"expiry_year" binding:"required"`
		Name           string           `json:"name,omitempty"`
		CVV            string           `json:"cvv,omitempty"`
		BillingAddress *common.Address  `json:"billing_address,omitempty"`
		Phone          *common.Phone    `json:"phone,omitempty"`
	}

	// Wallet -
	Wallet struct {
		Type      common.TokenType       `json:"type" binding:"required"`
		TokenData map[string]interface{} `json:"token_data" binding:"required"`
	}
)

// MarshalJSON ...
func (s *Request) MarshalJSON() ([]byte, error) {
	if s.Card != nil {
		return json.Marshal(s.Card)
	} else if s.Wallet != nil {
		return json.Marshal(s.Wallet)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (s *Request) UnmarshalJSON(data []byte) error {
	temp := &struct {
		Type string `json:"type"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	switch temp.Type {
	case common.Card.String():
		var source Card
		if err := json.Unmarshal(data, &source); err != nil {
			return err
		}
		s.Card = &source
		s.Wallet = nil
	default:
		var source Wallet
		if err := json.Unmarshal(data, &source); err != nil {
			return err
		}
		s.Wallet = &source
		s.Card = nil
	}
	return nil
}

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
		Created        *Created                 `json:"created,omitempty"`
	}

	// Created -
	Created struct {
		Type          string              `json:"type,omitempty"`
		Token         string              `json:"token" binding:"required"`
		ExpiresOn     time.Time           `json:"expires_on,omitempty"`
		ExpiryMonth   uint64              `json:"expiry_month,omitempty"`
		ExpiryYear    uint64              `json:"expiry_year,omitempty"`
		Scheme        string              `json:"scheme,omitempty"`
		Last4         string              `json:"last4,omitempty"`
		Bin           string              `json:"bin,omitempty"`
		CardType      common.CardType     `json:"card_type,omitempty"`
		CardCategory  common.CardCategory `json:"card_category,omitempty"`
		Issuer        string              `json:"issuer,omitempty"`
		IssuerCountry string              `json:"issuer_country,omitempty"`
		ProductID     string              `json:"product_id,omitempty"`
		ProductType   string              `json:"product_type,omitempty"`
	}
)
