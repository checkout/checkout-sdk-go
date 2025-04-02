package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type AccountType string

const (
	Savings AccountType = "savings"
	Current AccountType = "current"
	Cash    AccountType = "cash"
)

type (
	requestBankAccountDestination struct {
		Country       common.Country           `json:"country,omitempty"`
		AccountType   AccountType              `json:"account_type,omitempty"`
		Type          payments.DestinationType `json:"type,omitempty"`
		Iban          string                   `json:"iban,omitempty"`
		AccountNumber string                   `json:"account_number,omitempty"`
		BankCode      string                   `json:"bank_code,omitempty"`
		BranchCode    string                   `json:"branch_code,omitempty"`
		Bban          string                   `json:"bban,omitempty"`
		SwiftBic      string                   `json:"swift_bic,omitempty"`
		AccountHolder *common.AccountHolder    `json:"account_holder,omitempty"`
		Bank          *common.BankDetails      `json:"bank,omitempty"`
	}

	requestCardDestination struct {
		Type          payments.DestinationType `json:"type,omitempty"`
		Number        string                   `json:"number,omitempty"`
		ExpiryMonth   int                      `json:"expiry_month,omitempty"`
		ExpiryYear    int                      `json:"expiry_year,omitempty"`
		AccountHolder *common.AccountHolder    `json:"account_holder,omitempty"`
	}

	requestIdDestination struct {
		Type          payments.DestinationType `json:"type,omitempty"`
		Id            string                   `json:"id,omitempty"`
		AccountHolder *common.AccountHolder    `json:"account_holder,omitempty"`
	}

	requestTokenDestination struct {
		Type          payments.DestinationType `json:"type,omitempty"`
		Token         string                   `json:"token,omitempty"`
		AccountHolder *common.AccountHolder    `json:"account_holder,omitempty"`
	}

	requestNetworkTokenDestination struct {
		Type          payments.DestinationType  `json:"type,omitempty"`
		Token         string                    `json:"token,omitempty"`
		ExpiryMonth   int                       `json:"expiry_month,omitempty"`
		ExpiryYear    int                       `json:"expiry_year,omitempty"`
		TokenType     payments.NetworkTokenType `json:"token_type,omitempty"`
		Cryptogram    string                    `json:"cryptogram,omitempty"`
		Eci           string                    `json:"eci,omitempty"`
		AccountHolder *common.AccountHolder     `json:"account_holder,omitempty"`
	}
)

func NewRequestBankAccountDestination() *requestBankAccountDestination {
	return &requestBankAccountDestination{Type: payments.BankAccountDestination}
}

func NewRequestCardDestination() *requestCardDestination {
	return &requestCardDestination{Type: payments.CardDestination}
}

func NewRequestIdDestination() *requestIdDestination {
	return &requestIdDestination{Type: payments.IdDestination}
}

func NewRequestTokenDestination() *requestTokenDestination {
	return &requestTokenDestination{Type: payments.TokenDestination}
}

func NewRequestNetworkTokenDestination() *requestNetworkTokenDestination {
	return &requestNetworkTokenDestination{Type: payments.NetworkTokenDestination}
}

func (d *requestBankAccountDestination) GetType() payments.DestinationType {
	return d.Type
}

func (d *requestIdDestination) GetType() payments.DestinationType {
	return d.Type
}

func (d *requestTokenDestination) GetType() payments.DestinationType {
	return d.Type
}

func (d *requestCardDestination) GetType() payments.DestinationType {
	return d.Type
}

func (d *requestNetworkTokenDestination) GetType() payments.DestinationType {
	return d.Type
}

type (
	DestinationResponse struct {
		HttpMetadata                   common.HttpMetadata
		ResponseBankAccountDestination *ResponseBankAccountDestination
		AlternativeResponse            *common.AlternativeResponse
	}

	ResponseBankAccountDestination struct {
		Type          payments.DestinationType `json:"type,omitempty"`
		ExpiryMonth   int                      `json:"expiry_month,omitempty"`
		ExpiryYear    int                      `json:"expiry_year,omitempty"`
		Name          string                   `json:"name,omitempty"`
		Last4         string                   `json:"last4,omitempty"`
		Fingerprint   string                   `json:"fingerprint,omitempty"`
		Bin           string                   `json:"bin,omitempty"`
		CardType      common.CardType          `json:"card_type,omitempty"`
		CardCategory  common.CardCategory      `json:"card_category,omitempty"`
		Issuer        string                   `json:"issuer,omitempty"`
		IssuerCountry common.Country           `json:"issuer_country,omitempty"`
		ProductId     string                   `json:"product_id,omitempty"`
		ProductType   string                   `json:"product_type,omitempty"`
		AccountHolder *common.AccountHolder    `json:"account_holder,omitempty"`
	}
)

func (s *DestinationResponse) UnmarshalJSON(data []byte) error {
	var typeMapping payments.DestinationTypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Destination {
	case string(payments.BankAccountDestination):
		var typeMapping ResponseBankAccountDestination
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.ResponseBankAccountDestination = &typeMapping
	default:
		var typeMapping common.AlternativeResponse
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.AlternativeResponse = &typeMapping
	}

	return nil
}
