package nas

import (
	"github.com/checkout/checkout-sdk-go/v3/common"
)

type PaymentNetwork string

const (
	Local   PaymentNetwork = "local"
	Sepa    PaymentNetwork = "sepa"
	Fps     PaymentNetwork = "fps"
	Ach     PaymentNetwork = "ach"
	Fedwire PaymentNetwork = "fedwire"
	Swift   PaymentNetwork = "swift"
)

// SepaPaymentType is the type of payment for a SEPA instrument.
//
// The wire values are lowercase. The equivalent Bacs Direct Debit field is capitalized, so do not
// share one type between the two. Do not use payments.PaymentType either: it serializes capitalized
// values and also carries MOTO, Installment, PayLater and Unscheduled, which SEPA does not allow.
type SepaPaymentType string

const (
	SepaRecurring SepaPaymentType = "recurring"
	SepaRegular   SepaPaymentType = "regular"
)

// BacsPaymentType is the type of payment for a Bacs Direct Debit instrument.
//
// The wire values are capitalized. The equivalent SEPA field is lowercase, so do not share one type
// between the two.
type BacsPaymentType string

const (
	BacsRecurring BacsPaymentType = "Recurring"
	BacsRegular   BacsPaymentType = "Regular"
)

// SepaMandateType is the type of SEPA mandate.
type SepaMandateType string

const (
	SepaMandateCore SepaMandateType = "Core"
	SepaMandateB2B  SepaMandateType = "B2B"
)

// AchAccountType is the type of Direct Debit account of an ACH instrument.
//
// Distinct from common.AccountType, which declares savings, current and cash for the bank-account
// instrument. The ACH field declares savings and checking.
type AchAccountType string

const (
	AchSavings  AchAccountType = "savings"
	AchChecking AchAccountType = "checking"
)

// InstrumentAccountHolderType is the type of account holder on a stored instrument.
//
// The instrument schemas declare individual and corporate only, unlike common.AccountHolderType.
type InstrumentAccountHolderType string

const (
	InstrumentAccountHolderIndividual InstrumentAccountHolderType = "individual"
	InstrumentAccountHolderCorporate  InstrumentAccountHolderType = "corporate"
)

// SepaBillingAddress is the billing address of the account holder of a SEPA instrument.
//
// AddressLine1 max 200, AddressLine2 max 10 and Country min 2 max 2 characters. City and Zip are
// max 35 and max 16 when storing, and max 50 both when updating.
type SepaBillingAddress struct {
	AddressLine1 string         `json:"address_line1,omitempty"`
	AddressLine2 string         `json:"address_line2,omitempty"`
	City         string         `json:"city,omitempty"`
	Zip          string         `json:"zip,omitempty"`
	Country      common.Country `json:"country,omitempty"`
}

// SepaAccountHolder is the account holder details of a SEPA instrument.
//
// The schema declares these five properties only. Deliberately not common.AccountHolder, which is a
// superset carrying a phone, identification, a date of birth and a tax ID this schema does not
// declare.
type SepaAccountHolder struct {
	FirstName      string                      `json:"first_name,omitempty"`
	LastName       string                      `json:"last_name,omitempty"`
	CompanyName    string                      `json:"company_name,omitempty"`
	BillingAddress *SepaBillingAddress         `json:"billing_address,omitempty"`
	Type           InstrumentAccountHolderType `json:"type,omitempty"`
}

// InstrumentData is the details of a SEPA account.
//
// AccountNumber is the IBAN, min 15 max 34 characters. MandateId min 1 max 35 characters.
// DateOfSignature is required when MandateId is provided and defaults to the current date otherwise.
type InstrumentData struct {
	Type            SepaMandateType      `json:"type,omitempty"`
	AccountNumber   string               `json:"account_number,omitempty"`
	Country         common.Country       `json:"country,omitempty"`
	Currency        common.Currency      `json:"currency,omitempty"`
	PaymentType     SepaPaymentType      `json:"payment_type,omitempty"`
	MandateId       string               `json:"mandate_id,omitempty"`
	DateOfSignature *common.APIShortDate `json:"date_of_signature,omitempty"`
}

// AchAccountHolder is the account holder details of an ACH instrument.
//
// The schema marks all four properties required, but the descriptions qualify that: the names apply
// to an individual account holder and the company name to a corporate one. The ACH account holder
// declares no billing address.
type AchAccountHolder struct {
	FirstName   string                      `json:"first_name,omitempty"`
	LastName    string                      `json:"last_name,omitempty"`
	CompanyName string                      `json:"company_name,omitempty"`
	Type        InstrumentAccountHolderType `json:"type,omitempty"`
}

// AchInstrumentData is the details of an ACH bank account.
//
// AccountNumber min 4 max 17 characters. BankCode is the routing number, min 8 max 9 characters.
type AchInstrumentData struct {
	AccountType   AchAccountType  `json:"account_type,omitempty"`
	AccountNumber string          `json:"account_number,omitempty"`
	BankCode      string          `json:"bank_code,omitempty"`
	Currency      common.Currency `json:"currency,omitempty"`
	Country       common.Country  `json:"country,omitempty"`
}

// BacsBillingAddress is the billing address of the account holder of a Bacs Direct Debit instrument.
//
// AddressLine1 max 200 and AddressLine2 max 10 characters. City and Zip are max 35 and max 16 when
// storing, and max 50 both when updating. Country is min 2 max 2 characters and is the only required
// property when storing.
type BacsBillingAddress struct {
	AddressLine1 string         `json:"address_line1,omitempty"`
	AddressLine2 string         `json:"address_line2,omitempty"`
	City         string         `json:"city,omitempty"`
	Zip          string         `json:"zip,omitempty"`
	Country      common.Country `json:"country,omitempty"`
}

// CreateBacsAccountHolder is the account holder details of a Bacs instrument being stored.
//
// The store schema declares FirstName, LastName and BillingAddress only. It adds CompanyName and
// Type on update, which UpdateBacsAccountHolder carries.
type CreateBacsAccountHolder struct {
	FirstName      string              `json:"first_name,omitempty"`
	LastName       string              `json:"last_name,omitempty"`
	BillingAddress *BacsBillingAddress `json:"billing_address,omitempty"`
}

// UpdateBacsAccountHolder is the account holder details of a Bacs instrument being updated.
type UpdateBacsAccountHolder struct {
	FirstName      string                      `json:"first_name,omitempty"`
	LastName       string                      `json:"last_name,omitempty"`
	CompanyName    string                      `json:"company_name,omitempty"`
	BillingAddress *BacsBillingAddress         `json:"billing_address,omitempty"`
	Type           InstrumentAccountHolderType `json:"type,omitempty"`
}

// BacsInstrumentAccount is the account configuration for a Bacs Direct Debit instrument.
//
// ProcessingChannelId matches the pattern ^(pc)_(\w{26})$.
type BacsInstrumentAccount struct {
	ProcessingChannelId string `json:"processing_channel_id,omitempty"`
}

// BacsInstrumentData is the details of a Bacs Direct Debit account.
//
// AccountNumber is min 8 max 8 characters and BankCode is the six-digit sort code. PaymentType is
// capitalized, unlike the SEPA equivalent.
type BacsInstrumentData struct {
	AccountNumber     string          `json:"account_number,omitempty"`
	BankCode          string          `json:"bank_code,omitempty"`
	Country           common.Country  `json:"country,omitempty"`
	Currency          common.Currency `json:"currency,omitempty"`
	PaymentType       BacsPaymentType `json:"payment_type,omitempty"`
	AllowPartialMatch *bool           `json:"allow_partial_match,omitempty"`
}

type ProvisionNetworkToken struct {
	Provision bool `json:"provision,omitempty"`
}

type NetworkTokenResponse struct {
	Id    string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
}

type CreateCustomerInstrumentRequest struct {
	Id      string        `json:"id,omitempty"`
	Email   string        `json:"email,omitempty"`
	Name    string        `json:"name,omitempty"`
	Phone   *common.Phone `json:"phone,omitempty"`
	Default bool          `json:"default,omitempty"`
}
