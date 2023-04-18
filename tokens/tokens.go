package tokens

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

const tokensPath = "tokens"

type TokenType string

const (
	Card      TokenType = "card"
	ApplePay  TokenType = "applepay"
	GooglePay TokenType = "googlepay"
)

type (
	CardTokenRequest struct {
		Type           TokenType       `json:"type" binding:"required"`
		Number         string          `json:"number" binding:"required"`
		ExpiryMonth    int             `json:"expiry_month" binding:"required"`
		ExpiryYear     int             `json:"expiry_year" binding:"required"`
		Name           string          `json:"name,omitempty"`
		CVV            string          `json:"cvv,omitempty"`
		BillingAddress *common.Address `json:"billing_address,omitempty"`
		Phone          *common.Phone   `json:"phone,omitempty"`
	}

	WalletTokenRequest struct {
		Type      TokenType `json:"type" binding:"required"`
		TokenData TokenData `json:"token_data" binding:"required"`
	}
)

type (
	TokenData interface {
		GetType() TokenType
	}

	ApplePayTokenData struct {
		Version   string            `json:"version,omitempty"`
		Data      string            `json:"data,omitempty"`
		Signature string            `json:"signature,omitempty"`
		Header    map[string]string `json:"header,omitempty"`
	}

	GooglePayTokenData struct {
		Signature       string `json:"signature,omitempty"`
		ProtocolVersion string `json:"protocolVersion,omitempty"`
		SignedMessage   string `json:"signedMessage,omitempty"`
	}
)

func (t *ApplePayTokenData) GetType() TokenType {
	return ApplePay
}

func (t *GooglePayTokenData) GetType() TokenType {
	return GooglePay
}

type (
	CardTokenResponse struct {
		HttpMetadata   common.HttpMetadata
		Type           TokenType           `json:"type,omitempty"`
		Token          string              `json:"token" binding:"required"`
		ExpiresOn      time.Time           `json:"expires_on,omitempty"`
		ExpiryMonth    int                 `json:"expiry_month,omitempty"`
		ExpiryYear     int                 `json:"expiry_year,omitempty"`
		Scheme         string              `json:"scheme,omitempty"`
		Last4          string              `json:"last4,omitempty"`
		Bin            string              `json:"bin,omitempty"`
		CardType       common.CardType     `json:"card_type,omitempty"`
		CardCategory   common.CardCategory `json:"card_category,omitempty"`
		Issuer         string              `json:"issuer,omitempty"`
		IssuerCountry  string              `json:"issuer_country,omitempty"`
		ProductID      string              `json:"product_id,omitempty"`
		ProductType    string              `json:"product_type,omitempty"`
		TokenFormat    string              `json:"token_format,omitempty"`
		Name           string              `json:"name,omitempty"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
		Phone          *common.Phone       `json:"phone,omitempty"`
	}
)
