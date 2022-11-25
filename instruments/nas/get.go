package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/instruments"
)

type (
	GetInstrumentResponse struct {
		HttpMetadata                     common.HttpMetadata
		GetCardInstrumentResponse        *GetCardInstrumentResponse
		GetBankAccountInstrumentResponse *GetBankAccountInstrumentResponse
		AlternativeResponse              *common.AlternativeResponse
	}

	GetCardInstrumentResponse struct {
		Type          instruments.InstrumentType              `json:"type" binding:"required"`
		Id            string                                  `json:"id,omitempty"`
		Fingerprint   string                                  `json:"fingerprint,omitempty"`
		Customer      *instruments.InstrumentCustomerResponse `json:"customer,omitempty"`
		AccountHolder *common.AccountHolder                   `json:"account_holder,omitempty"`

		ExpiryMonth   int                 `json:"expiry_month,omitempty"`
		ExpiryYear    int                 `json:"expiry_year,omitempty"`
		Name          string              `json:"name,omitempty"`
		Scheme        string              `json:"scheme,omitempty"`
		SchemeLocal   string              `json:"scheme_local,omitempty"`
		Last4         string              `json:"last4,omitempty"`
		Bin           string              `json:"bin,omitempty"`
		CardType      common.CardType     `json:"card_type,omitempty"`
		CardCategory  common.CardCategory `json:"card_category,omitempty"`
		Issuer        string              `json:"issuer,omitempty"`
		IssuerCountry common.Country      `json:"issuer_country,omitempty"`
		ProductId     string              `json:"product_id,omitempty"`
		ProductType   string              `json:"product_type,omitempty"`
	}

	GetBankAccountInstrumentResponse struct {
		Type          instruments.InstrumentType              `json:"type" binding:"required"`
		Id            string                                  `json:"id,omitempty"`
		Fingerprint   string                                  `json:"fingerprint,omitempty"`
		Customer      *instruments.InstrumentCustomerResponse `json:"customer,omitempty"`
		AccountHolder *common.AccountHolder                   `json:"account_holder,omitempty"`

		AccountType   common.AccountType  `json:"account_type,omitempty"`
		AccountNumber string              `json:"account_number,omitempty"`
		BankCode      string              `json:"bank_code,omitempty"`
		Iban          string              `json:"iban,omitempty"`
		Bban          string              `json:"bban,omitempty"`
		SwiftBic      string              `json:"swift_bic,omitempty"`
		Currency      common.Currency     `json:"currency,omitempty"`
		Country       common.Country      `json:"country,omitempty"`
		BankDetails   *common.BankDetails `json:"bank,omitempty"`
	}
)

func (s *GetInstrumentResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(instruments.BankAccount):
		var response GetBankAccountInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.GetBankAccountInstrumentResponse = &response
	case string(instruments.Card):
		var response GetCardInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.GetCardInstrumentResponse = &response
	default:
		var response common.AlternativeResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.AlternativeResponse = &response
	}

	return nil
}
