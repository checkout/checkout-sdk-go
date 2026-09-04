package nas

import (
	"encoding/json"
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/instruments"
)

type QueryBankAccountFormatting struct {
	AccountHolderType common.AccountHolderType `json:"account-holder-type,omitempty"`
	PaymentNetwork    PaymentNetwork           `json:"payment-network,omitempty"`
}

type (
	GetInstrumentResponse struct {
		HttpMetadata                     common.HttpMetadata
		GetCardInstrumentResponse        *GetCardInstrumentResponse
		GetSepaInstrumentResponse        *GetSepaInstrumentResponse
		GetBankAccountInstrumentResponse *GetBankAccountInstrumentResponse
		GetAchInstrumentResponse         *GetAchInstrumentResponse
		GetBacsInstrumentResponse        *GetBacsInstrumentResponse
		AlternativeResponse              *common.AlternativeResponse
	}

	GetCardInstrumentResponse struct {
		Type                common.InstrumentType                   `json:"type" binding:"required"`
		Id                  string                                  `json:"id,omitempty"`
		Fingerprint         string                                  `json:"fingerprint,omitempty"`
		Customer            *instruments.InstrumentCustomerResponse `json:"customer,omitempty"`
		AccountHolder       *common.AccountHolder                   `json:"account_holder,omitempty"`
		EncryptedCardNumber string                                  `json:"encrypted_card_number,omitempty"`
		ExpiryMonth         int                                     `json:"expiry_month,omitempty"`
		ExpiryYear          int                                     `json:"expiry_year,omitempty"`
		Name                string                                  `json:"name,omitempty"`
		Scheme              string                                  `json:"scheme,omitempty"`
		SchemeLocal         string                                  `json:"scheme_local,omitempty"`
		Last4               string                                  `json:"last4,omitempty"`
		Bin                 string                                  `json:"bin,omitempty"`
		CardType            common.CardType                         `json:"card_type,omitempty"`
		CardCategory        common.CardCategory                     `json:"card_category,omitempty"`
		Issuer              string                                  `json:"issuer,omitempty"`
		IssuerCountry       common.Country                          `json:"issuer_country,omitempty"`
		ProductId           string                                  `json:"product_id,omitempty"`
		ProductType         string                                  `json:"product_type,omitempty"`
		NetworkToken        *NetworkTokenResponse                   `json:"network_token,omitempty"`
		CardWalletType      string                                  `json:"card_wallet_type,omitempty"`
		RegulatedIndicator  bool                                    `json:"regulated_indicator,omitempty"`
	}

	GetSepaInstrumentResponse struct {
		Type           common.InstrumentType `json:"type" binding:"required"`
		Id             string                `json:"id,omitempty"`
		Fingerprint    string                `json:"fingerprint,omitempty"`
		CreatedOn      *time.Time            `json:"created_on,omitempty"`
		ModifiedOn     *time.Time            `json:"modified_on,omitempty"`
		VaultId        string                `json:"vault_id,omitempty"`
		InstrumentData *InstrumentData       `json:"instrument_data,omitempty"`
		AccountHolder  *common.AccountHolder `json:"account_holder,omitempty"`
	}

	GetBankAccountInstrumentResponse struct {
		Type          common.InstrumentType                   `json:"type" binding:"required"`
		Id            string                                  `json:"id,omitempty"`
		Fingerprint   string                                  `json:"fingerprint,omitempty"`
		Customer      *instruments.InstrumentCustomerResponse `json:"customer,omitempty"`
		AccountHolder *common.AccountHolder                   `json:"account_holder,omitempty"`
		AccountType   common.AccountType                      `json:"account_type,omitempty"`
		AccountNumber string                                  `json:"account_number,omitempty"`
		BankCode      string                                  `json:"bank_code,omitempty"`
		Iban          string                                  `json:"iban,omitempty"`
		Bban          string                                  `json:"bban,omitempty"`
		SwiftBic      string                                  `json:"swift_bic,omitempty"`
		Currency      common.Currency                         `json:"currency,omitempty"`
		Country       common.Country                          `json:"country,omitempty"`
		BankDetails   *common.BankDetails                     `json:"bank,omitempty"`
	}

	GetAchInstrumentResponse struct {
		Type           common.InstrumentType                   `json:"type" binding:"required"`
		Id             string                                  `json:"id,omitempty"`
		Fingerprint    string                                  `json:"fingerprint,omitempty"`
		CreatedOn      *time.Time                              `json:"created_on,omitempty"`
		ModifiedOn     *time.Time                              `json:"modified_on,omitempty"`
		VaultId        string                                  `json:"vault_id,omitempty"`
		InstrumentData *AchInstrumentData                      `json:"instrument_data,omitempty"`
		AccountHolder  *AchAccountHolder                       `json:"account_holder,omitempty"`
		Customer       *instruments.InstrumentCustomerResponse `json:"customer,omitempty"`
	}

	// GetBacsAccountHolder is the account holder details of a stored Bacs Direct Debit instrument.
	//
	// Structurally identical to UpdateBacsAccountHolder, but kept separate so a response field is not
	// typed with a request-named type. RetrieveBacsInstrumentResponse requires FirstName, LastName and
	// BillingAddress.
	GetBacsAccountHolder struct {
		FirstName      string                      `json:"first_name,omitempty"`
		LastName       string                      `json:"last_name,omitempty"`
		CompanyName    string                      `json:"company_name,omitempty"`
		BillingAddress *BacsBillingAddress         `json:"billing_address,omitempty"`
		Type           InstrumentAccountHolderType `json:"type,omitempty"`
	}

	// GetBacsInstrumentAccount is the account configuration returned on a stored Bacs instrument.
	GetBacsInstrumentAccount struct {
		ClientId            string `json:"client_id,omitempty"`
		ProcessingChannelId string `json:"processing_channel_id,omitempty"`
	}

	// GetBacsInstrumentData is the details of a stored Bacs Direct Debit account.
	//
	// Status, MatchStatus, Description and MandateId are read-back fields the store and update
	// shapes do not declare. Status and MatchStatus are free-form strings in the specification.
	GetBacsInstrumentData struct {
		AccountNumber     string          `json:"account_number,omitempty"`
		BankCode          string          `json:"bank_code,omitempty"`
		Country           common.Country  `json:"country,omitempty"`
		Currency          common.Currency `json:"currency,omitempty"`
		PaymentType       BacsPaymentType `json:"payment_type,omitempty"`
		AllowPartialMatch *bool           `json:"allow_partial_match,omitempty"`
		Status            string          `json:"status,omitempty"`
		MatchStatus       string          `json:"match_status,omitempty"`
		Description       string          `json:"description,omitempty"`
		MandateId         string          `json:"mandate_id,omitempty"`
	}

	// GetBacsInstrumentResponse is the details of a stored Bacs Direct Debit instrument.
	GetBacsInstrumentResponse struct {
		Type        common.InstrumentType     `json:"type" binding:"required"`
		Id          string                    `json:"id,omitempty"`
		Fingerprint string                    `json:"fingerprint,omitempty"`
		CreatedOn   *time.Time                `json:"created_on,omitempty"`
		ModifiedOn  *time.Time                `json:"modified_on,omitempty"`
		VaultId     string                    `json:"vault_id,omitempty"`
		Account     *GetBacsInstrumentAccount `json:"account,omitempty"`
		// Validations is untyped because the specification declares the array items as a bare
		// object with no item schema. Retype it once the item schema is published.
		Validations    []map[string]interface{}                `json:"validations,omitempty"`
		InstrumentData *GetBacsInstrumentData                  `json:"instrument_data,omitempty"`
		AccountHolder  *GetBacsAccountHolder                   `json:"account_holder,omitempty"`
		Customer       *instruments.InstrumentCustomerResponse `json:"customer,omitempty"`
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
	case string(common.Sepa):
		var response GetSepaInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.GetSepaInstrumentResponse = &response
	case string(common.Ach):
		var response GetAchInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.GetAchInstrumentResponse = &response
	case string(common.Bacs):
		var response GetBacsInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.GetBacsInstrumentResponse = &response
	default:
		var response common.AlternativeResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.AlternativeResponse = &response
	}

	return nil
}
