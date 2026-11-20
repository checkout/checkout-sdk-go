package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/v3/common"
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
		Token string                `json:"token,omitempty"`
	}

	updateSepaInstrumentRequest struct {
		Type           common.InstrumentType `json:"type" binding:"required"`
		InstrumentData *InstrumentData       `json:"instrument_data,omitempty"`
		AccountHolder  *SepaAccountHolder    `json:"account_holder,omitempty"`
	}

	updateAchInstrumentRequest struct {
		Type           common.InstrumentType `json:"type" binding:"required"`
		InstrumentData *AchInstrumentData    `json:"instrument_data,omitempty"`
		AccountHolder  *AchAccountHolder     `json:"account_holder,omitempty"`
	}

	updateBacsInstrumentRequest struct {
		Type           common.InstrumentType    `json:"type" binding:"required"`
		InstrumentData *BacsInstrumentData      `json:"instrument_data,omitempty"`
		AccountHolder  *UpdateBacsAccountHolder `json:"account_holder,omitempty"`
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

func NewUpdateSepaInstrumentRequest() *updateSepaInstrumentRequest {
	return &updateSepaInstrumentRequest{
		Type: common.Sepa,
	}
}

func NewUpdateAchInstrumentRequest() *updateAchInstrumentRequest {
	return &updateAchInstrumentRequest{
		Type: common.Ach,
	}
}

func NewUpdateBacsInstrumentRequest() *updateBacsInstrumentRequest {
	return &updateBacsInstrumentRequest{
		Type: common.Bacs,
	}
}

type (
	// UpdateInstrumentResponse is the response to updating a payment instrument. One variant field
	// is set, chosen by the instrument type of the response; a response of another type is kept in
	// AlternativeResponse.
	UpdateInstrumentResponse struct {
		// HttpMetadata holds the status, headers and body of the HTTP response.
		HttpMetadata common.HttpMetadata
		// UpdateCardInstrumentResponse is set for the card type.
		UpdateCardInstrumentResponse *UpdateCardInstrumentResponse
		// UpdateBankAccountInstrumentResponse is set for the bank_account type.
		UpdateBankAccountInstrumentResponse *UpdateBankAccountInstrumentResponse
		// UpdateSepaInstrumentResponse is set for the sepa type.
		UpdateSepaInstrumentResponse *UpdateSepaInstrumentResponse
		// UpdateAchInstrumentResponse is set for the ach type.
		UpdateAchInstrumentResponse *UpdateAchInstrumentResponse
		// UpdateBacsInstrumentResponse is set for the bacs type.
		UpdateBacsInstrumentResponse *UpdateBacsInstrumentResponse
		// AlternativeResponse holds the response as a map for any other type.
		AlternativeResponse *common.AlternativeResponse
	}

	UpdateCardInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}

	UpdateBankAccountInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}

	UpdateSepaInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}

	UpdateAchInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}

	// UpdateBacsInstrumentResponse is the response returned after updating a stored Bacs Direct
	// Debit instrument. UpdateBacsInstrumentResponse declares type, id and fingerprint.
	UpdateBacsInstrumentResponse struct {
		Type        common.InstrumentType `json:"type" binding:"required"`
		Id          string                `json:"id,omitempty"`
		Fingerprint string                `json:"fingerprint,omitempty"`
	}
)

// UnmarshalJSON decodes the response into the variant field of its type, or into
// AlternativeResponse for any other type, and returns the error of that decoding.
func (s *UpdateInstrumentResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Type {
	case string(common.BankAccount):
		var response UpdateBankAccountInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return err
		}
		s.UpdateBankAccountInstrumentResponse = &response
	case string(common.Card):
		var response UpdateCardInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return err
		}
		s.UpdateCardInstrumentResponse = &response
	case string(common.Sepa):
		var response UpdateSepaInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return err
		}
		s.UpdateSepaInstrumentResponse = &response
	case string(common.Ach):
		var response UpdateAchInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return err
		}
		s.UpdateAchInstrumentResponse = &response
	case string(common.Bacs):
		var response UpdateBacsInstrumentResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return err
		}
		s.UpdateBacsInstrumentResponse = &response
	default:
		var response common.AlternativeResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return err
		}
		s.AlternativeResponse = &response
	}

	return nil
}
