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
		Stored            bool                `json:"stored"`
		StoreForFutureUse bool                `json:"store_for_future_use,omitempty"`
		BillingAddress    *common.Address     `json:"billing_address,omitempty"`
		Phone             *common.Phone       `json:"phone,omitempty"`
	}

	requestIdSource struct {
		Type              payments.SourceType   `json:"type,omitempty"`
		Id                string                `json:"id,omitempty"`
		Cvv               string                `json:"cvv,omitempty"`
		PaymentMethod     string                `json:"payment_method,omitempty"`
		Stored            *bool                 `json:"stored,omitempty"`
		StoreForFutureUse *bool                 `json:"storeForFutureUse,omitempty"`
		BillingAddress    *common.Address       `json:"billing_address,omitempty"`
		Phone             *common.Phone         `json:"phone,omitempty"`
		AccountHolder     *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestTokenSource struct {
		Type              payments.SourceType `json:"type,omitempty"`
		Token             string              `json:"token,omitempty"`
		BillingAddress    *common.Address     `json:"billing_address,omitempty"`
		Phone             *common.Phone       `json:"phone,omitempty"`
		Stored            *bool               `json:"stored,omitempty"`
		StoreForFutureUse bool                `json:"store_for_future_use,omitempty"`
	}

	requestProviderTokenSource struct {
		Type          payments.SourceType   `json:"type,omitempty"`
		PaymentMethod string                `json:"payment_method,omitempty"`
		Token         string                `json:"token,omitempty"`
		AccountHolder *common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestNetworkTokenSource struct {
		Type              payments.SourceType       `json:"type,omitempty"`
		Token             string                    `json:"token,omitempty"`
		ExpiryMonth       int                       `json:"expiry_month,omitempty"`
		ExpiryYear        int                       `json:"expiry_year,omitempty"`
		TokenType         payments.NetworkTokenType `json:"token_type,omitempty"`
		Cryptogram        string                    `json:"cryptogram,omitempty"`
		Eci               string                    `json:"eci,omitempty"`
		Stored            bool                      `json:"stored"`
		StoreForFutureUse bool                      `json:"store_for_future_use,omitempty"`
		Name              string                    `json:"name,omitempty"`
		Cvv               string                    `json:"cvv,omitempty"`
		BillingAddress    *common.Address           `json:"billing_address,omitempty"`
		Phone             *common.Phone             `json:"phone,omitempty"`
	}

	requestBankAccountSource struct {
		Type          payments.SourceType  `json:"type,omitempty"`
		PaymentMethod string               `json:"payment_method,omitempty"`
		AccountType   string               `json:"account_type,omitempty"`
		Country       common.Country       `json:"country,omitempty"`
		AccountNumber string               `json:"account_number,omitempty"`
		BankCode      string               `json:"bank_code,omitempty"`
		AccountHolder common.AccountHolder `json:"account_holder,omitempty"`
	}

	requestCustomerSource struct {
		Type payments.SourceType `json:"type,omitempty"`
		Id   string              `json:"number,omitempty"`
	}
)

func NewRequestCardSource() *requestCardSource {
	return &requestCardSource{Type: payments.CardSource}
}

func NewRequestIdSource() *requestIdSource {
	return &requestIdSource{Type: payments.IdSource}
}

func NewRequestTokenSource() *requestTokenSource {
	return &requestTokenSource{Type: payments.TokenSource}
}

func NewRequestProviderTokenSource() *requestProviderTokenSource {
	return &requestProviderTokenSource{Type: payments.ProviderTokenSource}
}

func NewRequestNetworkTokenSource() *requestNetworkTokenSource {
	return &requestNetworkTokenSource{Type: payments.NetworkTokenSource}
}

func NewRequestBankAccountSource() *requestBankAccountSource {
	return &requestBankAccountSource{Type: payments.BankAccountSource}
}

func NewRequestCustomerSource() *requestCustomerSource {
	return &requestCustomerSource{Type: payments.CustomerSource}
}

func (s *requestCardSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestIdSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestTokenSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestProviderTokenSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestNetworkTokenSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestBankAccountSource) GetType() payments.SourceType {
	return s.Type
}

func (s *requestCustomerSource) GetType() payments.SourceType {
	return s.Type
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

func NewPayoutCurrencyAccountSource() *payoutRequestSource {
	return &payoutRequestSource{Type: payments.CurrencyAccountSource}
}

func NewPayoutEntitySource() *payoutRequestSource {
	return &payoutRequestSource{Type: payments.EntitySource}
}

func (s *payoutRequestSource) GetType() payments.SourceType {
	return s.Type
}
