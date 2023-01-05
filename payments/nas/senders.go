package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type PaymentSenderType string

const (
	IndividualSender PaymentSenderType = "individual"
	CorporateSender  PaymentSenderType = "corporate"
	InstrumentSender PaymentSenderType = "instrument"
	Government       PaymentSenderType = "government"
)

type (
	PaymentCorporateSender struct {
		Type           PaymentSenderType                   `json:"type,omitempty"`
		CompanyName    string                              `json:"company_name,omitempty"`
		Address        *common.Address                     `json:"address,omitempty"`
		ReferenceType  string                              `json:"reference_type,omitempty"`
		SourceOfFunds  string                              `json:"source_of_funds,omitempty"`
		Identification *common.AccountHolderIdentification `json:"identification,omitempty"`
	}

	PaymentIndividualSender struct {
		Type           PaymentSenderType `json:"type,omitempty"`
		FirstName      string            `json:"first_name,omitempty"`
		LastName       string            `json:"last_name,omitempty"`
		Address        *common.Address   `json:"address,omitempty"`
		Identification *Identification   `json:"identification,omitempty"`
	}

	PaymentInstrumentSender struct {
		Type PaymentSenderType `json:"type,omitempty"`
	}
)

func NewPaymentCorporateSender() *PaymentCorporateSender {
	return &PaymentCorporateSender{Type: CorporateSender}
}

func NewPaymentIndividualSender() *PaymentIndividualSender {
	return &PaymentIndividualSender{Type: IndividualSender}
}

func NewPaymentInstrumentSender() *PaymentInstrumentSender {
	return &PaymentInstrumentSender{Type: InstrumentSender}
}

type (
	SenderResponse struct {
		HttpMetadata            common.HttpMetadata
		PaymentCorporateSender  *PaymentCorporateSender
		PaymentGovernmentSender *PaymentCorporateSender
		PaymentIndividualSender *PaymentIndividualSender
		PaymentInstrumentSender *PaymentInstrumentSender
		AlternativeResponse     *common.AlternativeResponse
	}
)

func (s *SenderResponse) UnmarshalJSON(data []byte) error {
	var typeMapping payments.SenderTypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Sender {
	case string(IndividualSender):
		var typeMapping PaymentIndividualSender
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentIndividualSender = &typeMapping
	case string(CorporateSender):
		var typeMapping PaymentCorporateSender
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentCorporateSender = &typeMapping
	case string(Government):
		var typeMapping PaymentCorporateSender
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentGovernmentSender = &typeMapping
	case string(InstrumentSender):
		var typeMapping PaymentInstrumentSender
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentInstrumentSender = &typeMapping
	default:
		var typeMapping common.AlternativeResponse
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.AlternativeResponse = &typeMapping
	}

	return nil
}
