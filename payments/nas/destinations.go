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
	PaymentRequestBankAccountDestination struct {
		Type          payments.PaymentDestinationType `json:"type,omitempty"`
		AccountType   AccountType                     `json:"account_type,omitempty"`
		AccountNumber string                          `json:"account_number,omitempty"`
		BankCode      string                          `json:"bank_code,omitempty"`
		BranchCode    string                          `json:"branch_code,omitempty"`
		Iban          string                          `json:"iban,omitempty"`
		Bban          string                          `json:"bban,omitempty"`
		SwiftBic      string                          `json:"swift_bic,omitempty"`
		Country       common.Country                  `json:"country,omitempty"`
		AccountHolder *common.AccountHolder           `json:"account_holder,omitempty"`
		Bank          *common.BankDetails             `json:"bank,omitempty"`
	}

	RequestIdDestination struct {
		Type payments.PaymentDestinationType `json:"type,omitempty"`
		Id   string                          `json:"id,omitempty"`
	}
)

type (
	DestinationResponse struct {
		HttpMetadata                   common.HttpMetadata
		ResponseBankAccountDestination *ResponseBankAccountDestination
		AlternativeResponse            *common.AlternativeResponse
	}

	ResponseBankAccountDestination struct {
		Type          payments.PaymentDestinationType `json:"type,omitempty"`
		ExpiryMonth   int                             `json:"expiry_month,omitempty"`
		ExpiryYear    int                             `json:"expiry_year,omitempty"`
		Name          string                          `json:"name,omitempty"`
		Last4         string                          `json:"last4,omitempty"`
		Fingerprint   string                          `json:"fingerprint,omitempty"`
		Bin           string                          `json:"bin,omitempty"`
		CardType      common.CardType                 `json:"card_type,omitempty"`
		CardCategory  common.CardCategory             `json:"card_category,omitempty"`
		Issuer        string                          `json:"issuer,omitempty"`
		IssuerCountry common.Country                  `json:"issuer_country,omitempty"`
		ProductId     string                          `json:"product_id,omitempty"`
		ProductType   string                          `json:"product_type,omitempty"`
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
