package sources

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	requestCardSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Number            string              `json:"number,omitempty"`
		ExpiryMonth       int                 `json:"expiry_month,omitempty"`
		ExpiryYear        int                 `json:"expiry_year,omitempty"`
		Name              string              `json:"name,omitempty"`
		Cvv               string              `json:"cvv,omitempty"`
		Stored            bool                `json:"stored,omitempty"`
		StoreForFutureUse bool                `json:"store_for_future_use,omitempty"`
		BillingAddress    *common.Address     `json:"billing_address,omitempty"`
		Phone             *common.Phone       `json:"phone,omitempty"`
	}

	requestIdSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Id                string              `json:"id,omitempty"`
		Cvv               string              `json:"cvv,omitempty"`
		Stored            *bool               `json:"stored,omitempty"`
		StoreForFutureUse *bool               `json:"storeForFutureUse,omitempty"`
	}

	requestCustomerSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"number,omitempty"`
	}

	requestTokenSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Token             string              `json:"token,omitempty"`
		BillingAddress    *common.Address     `json:"billing_address,omitempty"`
		Phone             *common.Phone       `json:"phone,omitempty"`
		Stored            *bool               `json:"stored,omitempty"`
		StoreForFutureUse bool                `json:"store_for_future_use,omitempty"`
	}

	requestNetworkTokenSource struct {
		Type           payments.SourceType       `json:"type,omitempty"`
		Token          string                    `json:"token,omitempty"`
		ExpiryMonth    int                       `json:"expiry_month,omitempty"`
		ExpiryYear     int                       `json:"expiry_year,omitempty"`
		TokenType      payments.NetworkTokenType `json:"token_type,omitempty"`
		Cryptogram     string                    `json:"cryptogram,omitempty"`
		Eci            string                    `json:"eci,omitempty"`
		Stored         bool                      `json:"stored"`
		Name           string                    `json:"name,omitempty"`
		Cvv            string                    `json:"cvv,omitempty"`
		BillingAddress *common.Address           `json:"billing_address,omitempty"`
		Phone          *common.Phone             `json:"phone,omitempty"`
	}

	requestDLocalSource struct {
		Type           payments.SourceType `json:"type,omitempty"`
		Number         string              `json:"number,omitempty"`
		ExpiryMonth    int                 `json:"expiry_month,omitempty"`
		ExpiryYear     int                 `json:"expiry_year,omitempty"`
		Name           string              `json:"name,omitempty"`
		Cvv            string              `json:"cvv,omitempty"`
		Stored         bool                `json:"stored"`
		BillingAddress *common.Address     `json:"billing_address,omitempty"`
		Phone          *common.Phone       `json:"phone,omitempty"`
	}
)

func NewRequestCardSource() *requestCardSource {
	return &requestCardSource{Type: payments.CardSource}
}

func NewRequestIdSource() *requestIdSource {
	return &requestIdSource{Type: payments.IdSource}
}

func NewRequestCustomerSource() *requestCustomerSource {
	return &requestCustomerSource{Type: payments.CustomerSource}
}

func NewRequestTokenSource() *requestTokenSource {
	return &requestTokenSource{Type: payments.TokenSource}
}

func NewRequestNetworkTokenSource() *requestNetworkTokenSource {
	return &requestNetworkTokenSource{Type: payments.NetworkTokenSource}
}

func NewRequestDLocalSource() *requestDLocalSource {
	return &requestDLocalSource{Type: payments.DLocalSource}
}

func (s *requestCardSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestIdSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestCustomerSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestTokenSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestNetworkTokenSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestDLocalSource) GetType() payments.SourceType {
	return s.Type
}
