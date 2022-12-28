package sources

import (
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type (
	RequestCardSource struct {
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

	RequestIdSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Id                string              `json:"id,omitempty"`
		Cvv               string              `json:"cvv,omitempty"`
		Stored            *bool               `json:"stored,omitempty"`
		StoreForFutureUse *bool               `json:"storeForFutureUse,omitempty"`
	}

	RequestCustomerSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"number,omitempty"`
	}

	RequestTokenSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Token             string              `json:"token,omitempty"`
		BillingAddress    *common.Address     `json:"billing_address,omitempty"`
		Phone             *common.Phone       `json:"phone,omitempty"`
		Stored            *bool               `json:"stored,omitempty"`
		StoreForFutureUse bool                `json:"store_for_future_use,omitempty"`
	}

	RequestNetworkTokenSource struct {
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

	RequestDLocalSource struct {
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

func NewRequestCardSource() *RequestCardSource {
	return &RequestCardSource{Type: payments.CardSource}
}

func NewRequestIdSource() *RequestIdSource {
	return &RequestIdSource{Type: payments.IdSource}
}

func NewRequestCustomerSource() *RequestCustomerSource {
	return &RequestCustomerSource{Type: payments.CustomerSource}
}

func NewRequestTokenSource() *RequestTokenSource {
	return &RequestTokenSource{Type: payments.TokenSource}
}

func NewRequestNetworkTokenSource() *RequestNetworkTokenSource {
	return &RequestNetworkTokenSource{Type: payments.NetworkTokenSource}
}

func NewRequestDLocalSource() *RequestDLocalSource {
	return &RequestDLocalSource{Type: payments.DLocalSource}
}
