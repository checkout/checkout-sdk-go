package accounts

import "github.com/checkout/checkout-sdk-go/common"

type BusinessType string

const (
	GeneralPartnership        BusinessType = "general_partnership"
	LimitedPartnership        BusinessType = "limited_partnership"
	PublicLimitedCompany      BusinessType = "public_limited_company"
	LimitedCompany            BusinessType = "limited_company"
	ProfessionalAssociation   BusinessType = "professional_association"
	UnincorporatedAssociation BusinessType = "unincorporated_association"
	AutoEntrepreneur          BusinessType = "auto_entrepreneur"
)

type (
	ContactDetails struct {
		Phone                *Phone                `json:"phone,omitempty"`
		EntityEmailAddresses *EntityEmailAddresses `json:"email_addresses,omitempty"`
	}

	EntityEmailAddresses struct {
		Primary string `json:"primary,omitempty"`
	}

	Profile struct {
		Urls                   []string        `json:"urls,omitempty"`
		Mccs                   []string        `json:"mccs,omitempty"`
		DefaultHoldingCurrency common.Currency `json:"default_holding_currency,omitempty"`
	}

	Company struct {
		BusinessRegistrationNumber string                  `json:"business_registration_number,omitempty"`
		BusinessType               BusinessType            `json:"business_type,omitempty"`
		LegalName                  string                  `json:"legal_name,omitempty"`
		TradingName                string                  `json:"trading_name,omitempty"`
		PrincipalAddress           *common.Address         `json:"principal_address,omitempty"`
		RegisteredAddress          *common.Address         `json:"registered_address,omitempty"`
		Document                   *EntityDocument         `json:"document,omitempty"`
		Representatives            []Representative        `json:"representatives,omitempty"`
		FinancialDetails           *EntityFinancialDetails `json:"financial_details,omitempty"`
	}

	EntityDocument struct {
		Type   string `json:"type,omitempty"`
		FileId string `json:"file_id,omitempty"`
	}

	Representative struct {
		FirstName      string          `json:"first_name,omitempty"`
		MiddleName     string          `json:"middle_name,omitempty"`
		LastName       string          `json:"last_name,omitempty"`
		Address        *common.Address `json:"address,omitempty"`
		Identification *Identification `json:"identification,omitempty"`
		Phone          *Phone          `json:"phone,omitempty"`
		DateOfBirth    *DateOfBirth    `json:"date_of_birth,omitempty"`
		PlaceOfBirth   *PlaceOfBirth   `json:"place_of_birth,omitempty"`
		Roles          []string        `json:"roles,omitempty"`
	}

	Identification struct {
		NationalIdNumber string    `json:"national_id_number,omitempty"`
		Document         *Document `json:"document,omitempty"`
	}

	Document struct {
		Type  common.DocumentType `json:"type,omitempty"`
		Front string              `json:"front,omitempty"`
		Back  string              `json:"back,omitempty"`
	}

	Phone struct {
		Number string `json:"number,omitempty"`
	}

	DateOfBirth struct {
		Day   int `json:"day,omitempty"`
		Month int `json:"month,omitempty"`
		Year  int `json:"year,omitempty"`
	}

	PlaceOfBirth struct {
		Country common.Country `json:"country,omitempty"`
	}

	Individual struct {
		FirstName         string          `json:"first_name,omitempty"`
		MiddleName        string          `json:"middle_name,omitempty"`
		LastName          string          `json:"last_name,omitempty"`
		TradingName       string          `json:"trading_name,omitempty"`
		NationalTaxId     string          `json:"national_tax_id,omitempty"`
		RegisteredAddress *common.Address `json:"registered_address,omitempty"`
		DateOfBirth       *DateOfBirth    `json:"date_of_birth,omitempty"`
		PlaceOfBirth      *PlaceOfBirth   `json:"place_of_birth,omitempty"`
		Identification    *Identification `json:"identification,omitempty"`
	}

	Capabilities struct {
		Payments *Payments `json:"payments,omitempty"`
		Payouts  *Payouts  `json:"payouts,omitempty"`
	}

	Payments struct {
		Available bool `json:"available,omitempty"`
		Enabled   bool `json:"enabled,omitempty"`
	}

	Payouts struct {
		Available bool `json:"available,omitempty"`
		Enabled   bool `json:"enabled,omitempty"`
	}

	RequirementsDue struct {
		Field  string `json:"field,omitempty"`
		Reason string `json:"reason,omitempty"`
	}

	Instrument struct {
		Id       string              `json:"id,omitempty"`
		Label    string              `json:"label,omitempty"`
		Status   InstrumentStatus    `json:"status,omitempty"`
		Document *InstrumentDocument `json:"document,omitempty"`
	}

	InstrumentDocument struct {
		Type   string `json:"type,omitempty"`
		FileId string `json:"file_id,omitempty"`
	}

	EntityFinancialDetails struct {
		AnnualProcessingVolume  int64                     `json:"annual_processing_volume,omitempty"`
		AverageTransactionValue int64                     `json:"average_transaction_value,omitempty"`
		HighestTransactionValue int64                     `json:"highest_transaction_value,omitempty"`
		Documents               *EntityFinancialDocuments `json:"documents,omitempty"`
	}

	EntityFinancialDocuments struct {
		BankStatement      *EntityDocument `json:"bank_statement,omitempty"`
		FinancialStatement *EntityDocument `json:"financial_statement,omitempty"`
	}
)
