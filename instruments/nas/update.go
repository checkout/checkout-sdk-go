package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
)

type (
	UpdateInstrumentRequest interface{}

	updateBankAccountInstrumentRequest struct {
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

	updateCardInstrumentRequest struct {
		Type          common.InstrumentType         `json:"type" binding:"required"`
		ExpiryMonth   int                           `json:"expiry_month" binding:"required"`
		ExpiryYear    int                           `json:"expiry_year" binding:"required"`
		Name          string                        `json:"name" binding:"required"`
		Customer      *common.UpdateCustomerRequest `json:"customer" binding:"required"`
		AccountHolder *common.AccountHolder         `json:"account_holder" binding:"required"`
	}

	updateTokenInstrumentRequest struct {
		Type  common.InstrumentType `json:"type" binding:"required"`
		Token string                `json:"account_number,omitempty"`
	}
)

func NewUpdateBankAccountInstrumentRequest() *updateBankAccountInstrumentRequest {
	return &updateBankAccountInstrumentRequest{
		Type: common.BankAccount,
	}
}

func NewUpdateCardInstrumentRequest() *updateCardInstrumentRequest {
	return &updateCardInstrumentRequest{
		Type: common.Card,
	}
}

func NewUpdateTokenInstrumentRequest() *updateTokenInstrumentRequest {
	return &updateTokenInstrumentRequest{
		Type: common.Token,
	}
}

type (
	UpdateInstrumentResponse struct {
		HttpMetadata                        common.HttpMetadata
		UpdateCardInstrumentResponse        *UpdateCardInstrumentResponse
		UpdateBankAccountInstrumentResponse *UpdateBankAccountInstrumentResponse
		AlternativeResponse                 *common.AlternativeResponse
	}

	// UpdateCardInstrumentResponse TODO review this response struct to check if we need both
	UpdateCardInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}

	// UpdateBankAccountInstrumentResponse TODO review this response struct to check if we need both
	UpdateBankAccountInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}
)

func (s *UpdateInstrumentResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(common.BankAccount):
		var response UpdateBankAccountInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.UpdateBankAccountInstrumentResponse = &response
	case string(common.Card):
		var response UpdateCardInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.UpdateCardInstrumentResponse = &response
	default:
		var response common.AlternativeResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.AlternativeResponse = &response
	}

	return nil
}
