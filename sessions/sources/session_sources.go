package sources

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type SessionSourceType string

const (
	Card         SessionSourceType = "card"
	Id           SessionSourceType = "id"
	Token        SessionSourceType = "token"
	NetworkToken SessionSourceType = "network_token"
)

type SessionScheme string

const (
	Amex            SessionScheme = "amex"
	CartesBancaires SessionScheme = "cartes_bancaires"
	Diners          SessionScheme = "diners"
	Jcb             SessionScheme = "jcb"
	Mastercard      SessionScheme = "mastercard"
	Visa            SessionScheme = "visa"
)

type (
	SessionSource interface {
		GetType() SessionSourceType
	}

	SessionSourceInfo struct {
		Type           SessionSourceType `json:"type,omitempty"`
		Scheme         SessionScheme     `json:"scheme,omitempty"`
		BillingAddress *SessionAddress   `json:"billing_address,omitempty"`
		HomePhone      *common.Phone     `json:"home_phone,omitempty"`
		MobilePhone    *common.Phone     `json:"mobile_phone,omitempty"`
		WorkPhone      *common.Phone     `json:"work_phone,omitempty"`
		Email          string            `json:"email,omitempty"`
	}

	sessionCardSource struct {
		SessionSourceInfo
		Number            string `json:"number,omitempty"`
		ExpiryMonth       int    `json:"expiry_month,omitempty"`
		ExpiryYear        int    `json:"expiry_year,omitempty"`
		Name              string `json:"name,omitempty"`
		Stored            bool   `json:"stored,omitempty" default:"false"`
		StoreForFutureUse bool   `json:"store_for_future_use,omitempty"`
	}

	sessionIdSource struct {
		SessionSourceInfo
		Id string `json:"id,omitempty"`
	}

	sessionTokenSource struct {
		SessionSourceInfo
		Token             string `json:"token,omitempty"`
		StoreForFutureUse bool   `json:"store_for_future_use,omitempty"`
	}

	sessionNetworkTokenSource struct {
		SessionSourceInfo
		Token       string `json:"token,omitempty"`
		ExpiryMonth int    `json:"expiry_month,omitempty"`
		ExpiryYear  int    `json:"expiry_year,omitempty"`
		Name        string `json:"name,omitempty"`
		Stored      bool   `json:"stored,omitempty"`
	}
)

func NewSessionCardSource() *sessionCardSource {
	return &sessionCardSource{SessionSourceInfo: SessionSourceInfo{Type: Card}}
}

func NewSessionIdSource() *sessionIdSource {
	return &sessionIdSource{SessionSourceInfo: SessionSourceInfo{Type: Id}}
}

func NewSessionTokenSource() *sessionTokenSource {
	return &sessionTokenSource{SessionSourceInfo: SessionSourceInfo{Type: Token}}
}

func NewSessionNetworkTokenSource() *sessionNetworkTokenSource {
	return &sessionNetworkTokenSource{SessionSourceInfo: SessionSourceInfo{Type: NetworkToken}}
}

func (s *sessionCardSource) GetType() SessionSourceType {
	return s.Type
}

func (s *sessionIdSource) GetType() SessionSourceType {
	return s.Type
}

func (s *sessionTokenSource) GetType() SessionSourceType {
	return s.Type
}

func (s *sessionNetworkTokenSource) GetType() SessionSourceType {
	return s.Type
}

type SessionAddress struct {
	common.Address
	AddressLine3 string `json:"address_line3,omitempty"`
}
