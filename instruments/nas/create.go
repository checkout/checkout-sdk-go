package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type (
	CreateInstrumentRequest interface{}

	createBankAccountInstrumentRequest struct {
		Type                common.InstrumentType            `json:"type" binding:"required"`
		AccountType         common.AccountType               `json:"account_type,omitempty"`
		AccountNumber       string                           `json:"account_number,omitempty"`
		BankCode            string                           `json:"bank_code,omitempty"`
		BranchCode          string                           `json:"branch_code,omitempty"`
		Iban                string                           `json:"iban,omitempty"`
		Bban                string                           `json:"bban,omitempty"`
		SwiftBic            string                           `json:"swift_bic,omitempty"`
		Currency            common.Currency                  `json:"currency,omitempty"`
		Country             common.Country                   `json:"country,omitempty"`
		ProcessingChannelId string                           `json:"processing_channel_id,omitempty"`
		AccountHolder       *common.AccountHolder            `json:"account_holder,omitempty"`
		BankDetails         *common.BankDetails              `json:"bank,omitempty"`
		Customer            *CreateCustomerInstrumentRequest `json:"customer,omitempty"`
	}

	createTokenInstrumentRequest struct {
		Type          common.InstrumentType            `json:"type" binding:"required"`
		Token         string                           `json:"token" binding:"required"`
		AccountHolder *common.AccountHolder            `json:"account_holder" binding:"required"`
		Customer      *CreateCustomerInstrumentRequest `json:"customer,omitempty"`
	}

	createCardInstrumentRequest struct {
		Type                common.InstrumentType            `json:"type" binding:"required"`
		Number              string                           `json:"number,omitempty"`
		ExpiryMonth         int                              `json:"expiry_month,omitempty"`
		ExpiryYear          int                              `json:"expiry_year,omitempty"`
		AccountHolder       *common.AccountHolder            `json:"account_holder,omitempty"`
		Customer            *CreateCustomerInstrumentRequest `json:"customer,omitempty"`
		EntityId            string                           `json:"entity_id,omitempty"`
		ProcessingChannelId string                           `json:"processing_channel_id,omitempty"`
		NetworkToken        *ProvisionNetworkToken           `json:"network_token,omitempty"`
	}

	createSepaInstrumentRequest struct {
		Type           common.InstrumentType            `json:"type" binding:"required"`
		InstrumentData *InstrumentData                  `json:"instrument_data,omitempty"`
		AccountHolder  *common.AccountHolder            `json:"account_holder" binding:"required"`
		Customer       *CreateCustomerInstrumentRequest `json:"customer,omitempty"`
	}

	createAchInstrumentRequest struct {
		Type           common.InstrumentType            `json:"type" binding:"required"`
		InstrumentData *AchInstrumentData               `json:"instrument_data,omitempty"`
		AccountHolder  *common.AccountHolder            `json:"account_holder" binding:"required"`
		Customer       *CreateCustomerInstrumentRequest `json:"customer,omitempty"`
	}
)

func NewCreateBankAccountInstrumentRequest() *createBankAccountInstrumentRequest {
	return &createBankAccountInstrumentRequest{
		Type: common.BankAccount,
	}
}

func NewCreateTokenInstrumentRequest() *createTokenInstrumentRequest {
	return &createTokenInstrumentRequest{
		Type: common.Token,
	}
}

func NewCreateCardInstrumentRequest() *createCardInstrumentRequest {
	return &createCardInstrumentRequest{
		Type: common.Card,
	}
}

func NewCreateSepaInstrumentRequest() *createSepaInstrumentRequest {
	return &createSepaInstrumentRequest{
		Type: common.Sepa,
	}
}

func NewCreateAchInstrumentRequest() *createAchInstrumentRequest {
	return &createAchInstrumentRequest{
		Type: common.Ach,
	}
}

type (
	CreateInstrumentResponse struct {
		HttpMetadata                        common.HttpMetadata
		CreateBankAccountInstrumentResponse *CreateBankAccountInstrumentResponse
		CreateCardInstrumentResponse        *CreateCardInstrumentResponse
		CreateTokenInstrumentResponse       *CreateTokenInstrumentResponse
		CreateSepaInstrumentResponse        *CreateSepaInstrumentResponse
		CreateAchInstrumentResponse         *CreateAchInstrumentResponse
		AlternativeResponse                 *common.AlternativeResponse
	}

	CreateBankAccountInstrumentResponse struct {
		Type common.InstrumentType `json:"type" binding:"required"`
		// common
		Id               string                   `json:"id,omitempty"`
		Fingerprint      string                   `json:"fingerprint,omitempty"`
		CustomerResponse *common.CustomerResponse `json:"customer,omitempty"`
		// specific
		BankDetails   *common.BankDetails `json:"bank,omitempty"`
		SwiftBic      string              `json:"swift_bic,omitempty"`
		AccountNumber string              `json:"account_number,omitempty"`
		BankCode      string              `json:"bank_code,omitempty"`
		Iban          string              `json:"iban,omitempty"`
	}

	CreateTokenInstrumentResponse struct {
		Type          common.InstrumentType    `json:"type" binding:"required"`
		Id            string                   `json:"id,omitempty"`
		Fingerprint   string                   `json:"fingerprint,omitempty"`
		Customer      *common.CustomerResponse `json:"customer,omitempty"`
		AccountHolder *common.AccountHolder    `json:"account_holder,omitempty"`
		ExpiryMonth   int                      `json:"expiry_month,omitempty"`
		ExpiryYear    int                      `json:"expiry_year,omitempty"`
		Scheme        string                   `json:"scheme,omitempty"`
		SchemeLocal   string                   `json:"scheme_local,omitempty"`
		Last4         string                   `json:"last4,omitempty"`
		Bin           string                   `json:"bin,omitempty"`
		CardType      common.CardType          `json:"card_type,omitempty"`
		CardCategory  common.CardCategory      `json:"card_category,omitempty"`
		Issuer        string                   `json:"issuer,omitempty"`
		IssuerCountry common.Country           `json:"issuer_country,omitempty"`
		ProductId     string                   `json:"product_id,omitempty"`
		ProductType   string                   `json:"product_type,omitempty"`
		NetworkToken  *NetworkTokenResponse    `json:"network_token,omitempty"`
	}

	CreateCardInstrumentResponse struct {
		Type          common.InstrumentType    `json:"type" binding:"required"`
		Id            string                   `json:"id,omitempty"`
		Fingerprint   string                   `json:"fingerprint,omitempty"`
		Customer      *common.CustomerResponse `json:"customer,omitempty"`
		AccountHolder *common.AccountHolder    `json:"account_holder,omitempty"`
		ExpiryMonth   int                      `json:"expiry_month,omitempty"`
		ExpiryYear    int                      `json:"expiry_year,omitempty"`
		Scheme        string                   `json:"scheme,omitempty"`
		SchemeLocal   string                   `json:"scheme_local,omitempty"`
		Last4         string                   `json:"last4,omitempty"`
		Bin           string                   `json:"bin,omitempty"`
		CardType      common.CardType          `json:"card_type,omitempty"`
		CardCategory  common.CardCategory      `json:"card_category,omitempty"`
		Issuer        string                   `json:"issuer,omitempty"`
		IssuerCountry common.Country           `json:"issuer_country,omitempty"`
		ProductId     string                   `json:"product_id,omitempty"`
		ProductType   string                   `json:"product_type,omitempty"`
		NetworkToken  *NetworkTokenResponse    `json:"network_token,omitempty"`
	}

	CreateSepaInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}

	CreateAchInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}
)

func (s *CreateInstrumentResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(common.BankAccount):
		var response CreateBankAccountInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateBankAccountInstrumentResponse = &response
	case string(common.Card):
		var response CreateCardInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateCardInstrumentResponse = &response
	case string(common.Token):
		var response CreateTokenInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateTokenInstrumentResponse = &response
	case string(common.Sepa):
		var response CreateSepaInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateSepaInstrumentResponse = &response
	case string(common.Ach):
		var response CreateAchInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.CreateAchInstrumentResponse = &response
	default:
		var response common.AlternativeResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.AlternativeResponse = &response
	}

	return nil
}
