package nas

import (
	"encoding/json"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
)

type SenderType string

const (
	Individual SenderType = "individual"
	Corporate  SenderType = "corporate"
	Instrument SenderType = "instrument"
	Government SenderType = "government"
)

type (
	Sender interface {
		GetType() SenderType
	}

	CorporateSender struct {
		Type           SenderType                          `json:"type,omitempty"`
		CompanyName    string                              `json:"company_name,omitempty"`
		Address        *common.Address                     `json:"address,omitempty"`
		Reference      string                              `json:"reference,omitempty"`
		ReferenceType  string                              `json:"reference_type,omitempty"`
		SourceOfFunds  string                              `json:"source_of_funds,omitempty"`
		Identification *common.AccountHolderIdentification `json:"identification,omitempty"`
	}

	IndividualSender struct {
		Type           SenderType      `json:"type,omitempty"`
		FirstName      string          `json:"first_name,omitempty"`
		MiddleName     string          `json:"middle_name,omitempty"`
		LastName       string          `json:"last_name,omitempty"`
		Address        *common.Address `json:"address,omitempty"`
		Dob            string          `json:"dob,omitempty"`
		DateOfBirth    string          `json:"date_of_birth,omitempty"`
		Identification *Identification `json:"identification,omitempty"`
		Reference      string          `json:"reference,omitempty"`
		ReferenceType  string          `json:"reference_type,omitempty"`
		SourceOfFunds  string          `json:"source_of_funds,omitempty"`
		CountryOfBirth common.Country  `json:"country_of_birth,omitempty"`
		Nationality    common.Country  `json:"nationality,omitempty"`
	}

	InstrumentSender struct {
		Type      SenderType `json:"type,omitempty"`
		Reference string     `json:"reference,omitempty"`
	}
)

func NewRequestCorporateSender() *CorporateSender {
	return &CorporateSender{Type: Corporate}
}

func NewRequestGovernmentSender() *CorporateSender {
	return &CorporateSender{Type: Government}
}

func NewRequestIndividualSender() *IndividualSender {
	return &IndividualSender{Type: Individual}
}

func NewRequestInstrumentSender() *InstrumentSender {
	return &InstrumentSender{Type: Instrument}
}

func (s *CorporateSender) GetType() SenderType {
	return s.Type
}

func (s *IndividualSender) GetType() SenderType {
	return s.Type
}

func (s *InstrumentSender) GetType() SenderType {
	return s.Type
}

type (
	SenderResponse struct {
		HttpMetadata            common.HttpMetadata
		PaymentCorporateSender  *CorporateSender
		PaymentGovernmentSender *CorporateSender
		PaymentIndividualSender *IndividualSender
		PaymentInstrumentSender *InstrumentSender
		AlternativeResponse     *common.AlternativeResponse
	}
)

func (s *SenderResponse) UnmarshalJSON(data []byte) error {
	var typeMapping payments.SenderTypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.Sender {
	case string(Individual):
		var typeMapping IndividualSender
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentIndividualSender = &typeMapping
	case string(Corporate):
		var typeMapping CorporateSender
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentCorporateSender = &typeMapping
	case string(Government):
		var typeMapping CorporateSender
		if err := json.Unmarshal(data, &typeMapping); err != nil {
			return err
		}
		s.PaymentGovernmentSender = &typeMapping
	case string(Instrument):
		var typeMapping InstrumentSender
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
