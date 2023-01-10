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
		Stored            bool                `json:"stored"`
		StoreForFutureUse bool                `json:"store_for_future_use,omitempty"`
		BillingAddress    *common.Address     `json:"billing_address,omitempty"`
		Phone             *common.Phone       `json:"phone,omitempty"`
	}

	RequestIdSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Id                string              `json:"id,omitempty"`
		Cvv               string              `json:"cvv,omitempty"`
		PaymentMethod     string              `json:"payment_method,omitempty"`
		Stored            *bool               `json:"stored,omitempty"`
		StoreForFutureUse *bool               `json:"storeForFutureUse,omitempty"`
	}

	RequestTokenSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Token             string              `json:"token,omitempty"`
		BillingAddress    *common.Address     `json:"billing_address,omitempty"`
		Phone             *common.Phone       `json:"phone,omitempty"`
		Stored            *bool               `json:"stored,omitempty"`
		StoreForFutureUse bool                `json:"store_for_future_use,omitempty"`
	}

	RequestProviderTokenSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		PaymentMethod string                `json:"payment_method,omitempty"`
		Token         string                `json:"token,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	RequestNetworkTokenSource struct {
		Type           payments.SourceType       `json:"type,omitempty"`
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

	RequestBankAccountSource struct {
		Type          payments.SourceType  `json:"type,omitempty"`
		PaymentMethod string               `json:"payment_method,omitempty"`
		AccountType   string               `json:"account_type,omitempty"`
		Country       common.Country       `json:"country,omitempty"`
		AccountNumber string               `json:"account_number,omitempty"`
		BankCode      string               `json:"bank_code,omitempty"`
		AccountHolder common.AccountHolder `json:"account_holder,omitempty"`
	}

	RequestCustomerSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"number,omitempty"`
	}
)

func NewRequestCardSource() *RequestCardSource {
	return &RequestCardSource{Type: payments.CardSource}
}

func NewRequestIdSource() *RequestIdSource {
	return &RequestIdSource{Type: payments.IdSource}
}

func NewRequestTokenSource() *RequestTokenSource {
	return &RequestTokenSource{Type: payments.TokenSource}
}

func NewRequestProviderTokenSource() *RequestProviderTokenSource {
	return &RequestProviderTokenSource{Type: payments.ProviderTokenSource}
}

func NewRequestNetworkTokenSource() *RequestNetworkTokenSource {
	return &RequestNetworkTokenSource{Type: payments.NetworkTokenSource}
}

func NewRequestBankAccountSource() *RequestBankAccountSource {
	return &RequestBankAccountSource{Type: payments.BankAccountSource}
}

func NewRequestCustomerSource() *RequestCustomerSource {
	return &RequestCustomerSource{Type: payments.CustomerSource}
}

type (
	PayoutSource interface {
		GetType() payments.SourceType
	}

	payoutRequestSource struct {
		Type   payments.SourceType `json:"type,omitempty"`
		Id     string              `json:"id,omitempty"`
		Amount int64               `json:"amount,omitempty"`
	}
)

func NewPayoutRequestSource() *payoutRequestSource {
	return &payoutRequestSource{Type: payments.CurrencyAccountSource}
}

func (s *payoutRequestSource) GetType() payments.SourceType {
	return s.Type
}
