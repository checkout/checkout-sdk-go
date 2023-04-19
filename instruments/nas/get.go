package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/instruments"
)

type QueryBankAccountFormatting struct {
	AccountHolderType common.AccountHolderType `json:"account-holder-type,omitempty"`
	PaymentNetwork    PaymentNetwork           `json:"payment-network,omitempty"`
}

type (
	GetInstrumentResponse struct {
		HttpMetadata                     common.HttpMetadata
		GetCardInstrumentResponse        *GetCardInstrumentResponse
		GetBankAccountInstrumentResponse *GetBankAccountInstrumentResponse
		AlternativeResponse              *common.AlternativeResponse
	}

	GetCardInstrumentResponse struct {
		Type          common.InstrumentType                   `json:"type" binding:"required"`
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
		Type          common.InstrumentType                   `json:"type" binding:"required"`
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

	InstrumentSectionFieldAllowedOption struct {
		Id      string `json:"id,omitempty"`
		Display string `json:"display,omitempty"`
	}

	InstrumentSectionFieldDependencies struct {
		FieldId string `json:"field_id,omitempty"`
		Value   string `json:"value,omitempty"`
	}

	InstrumentSectionField struct {
		Id              string                                `json:"id" binding:"required"`
		Section         string                                `json:"section,omitempty"`
		Display         string                                `json:"display" binding:"required"`
		HelpText        string                                `json:"help_text,omitempty"`
		Type            string                                `json:"type" binding:"required"`
		Required        bool                                  `json:"required" binding:"required"`
		ValidationRegex string                                `json:"validation_regex,omitempty"`
		MinLength       int                                   `json:"min_length,omitempty"`
		MaxLength       int                                   `json:"max_length,omitempty"`
		AllowedOptions  []InstrumentSectionFieldAllowedOption `json:"allowed_options,omitempty"`
		Dependencies    []InstrumentSectionFieldDependencies  `json:"dependencies,omitempty"`
	}

	InstrumentSection struct {
		Name   string                   `json:"name,omitempty" binding:"required"`
		Fields []InstrumentSectionField `json:"fields,omitempty"`
	}

	GetBankAccountFieldFormattingResponse struct {
		HttpMetadata common.HttpMetadata
		Sections     []InstrumentSection `json:"sections,omitempty"`
	}
)

func (s *GetInstrumentResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(common.BankAccount):
		var response GetBankAccountInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.GetBankAccountInstrumentResponse = &response
	case string(common.Card):
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
