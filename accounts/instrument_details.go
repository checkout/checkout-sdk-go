package accounts

import (
	"github.com/checkout/checkout-sdk-go/common"
)

type InstrumentDetail string

type (
	InstrumentDetails interface {
		GetType() string
	}

	InstrumentDetailsFasterPayments struct {
		AccountNumber string `json:"account_number,omitempty"`
		BankCode      string `json:"bank_code,omitempty"`
	}

	InstrumentDetailsSepa struct {
		Iban     string `json:"iban,omitempty"`
		SwiftBic string `json:"swift_bic,omitempty"`
	}

	InstrumentDetailsCardToken struct {
		Token string `json:"token,omitempty"`
	}
)

type (
	PaymentInstrumentDetailsResponse struct {
		HttpMetadata       common.HttpMetadata
		Id                 string                 `json:"id,omitempty"`
		Status             InstrumentStatus       `json:"status,omitempty"`
		InstrumentId       string                 `json:"instrument_id,omitempty"`
		Label              string                 `json:"label,omitempty"`
		Type               common.InstrumentType  `json:"type,omitempty"`
		Currency           common.Currency        `json:"currency,omitempty"`
		Country            common.Country         `json:"country,omitempty"`
		DefaultDestination bool                   `json:"default,omitempty"`
		Document           *InstrumentDocument    `json:"document,omitempty"`
		Links              map[string]common.Link `json:"_links"`
	}

	PaymentInstrumentQueryResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []PaymentInstrumentDetailsResponse `json:"data,omitempty"`
	}
)

type (
	PaymentInstrumentRequest struct {
		Label              string                `json:"label,omitempty"`
		Type               common.InstrumentType `json:"type,omitempty"`
		Currency           common.Currency       `json:"currency,omitempty"`
		Country            common.Country        `json:"country,omitempty"`
		DefaultDestination bool                  `json:"default,omitempty"`
		Document           *InstrumentDocument   `json:"document"`
		InstrumentDetails  InstrumentDetails     `json:"instrument_details,omitempty"`
	}

	PaymentInstrumentsQuery struct {
		Status InstrumentStatus `json:"status,omitempty"`
	}

	UpdatePaymentInstrumentRequest struct {
		Label   string  `json:"label,omitempty"`
		Default bool    `json:"default,omitempty"`
		Headers Headers `json:"headers,omitempty"`
	}
)

type Headers struct {
	IfMatch string `json:"if-match,omitempty"`
}

func (s *InstrumentDetailsFasterPayments) GetType() string {
	return "FasterPayment"
}

func (s *InstrumentDetailsSepa) GetType() string {
	return "Sepa"
}

func (s *InstrumentDetailsCardToken) GetType() string {
	return string(common.CardToken)
}
