package accounts

import (
	"github.com/checkout/checkout-sdk-go/common"
)

const (
	accountsPath           = "accounts"
	entitiesPath           = "entities"
	instrumentsPath        = "instruments"
	payoutSchedulesPath    = "payout-schedules"
	filesPath              = "files"
	paymentInstrumentsPath = "payment-instruments"
)

type AccountHolderIdentificationType string

const (
	Passport            AccountHolderIdentificationType = "passport"
	DrivingLicence      AccountHolderIdentificationType = "driving_licence"
	NationalId          AccountHolderIdentificationType = "national_id"
	CompanyRegistration AccountHolderIdentificationType = "company_registration"
	TaxId               AccountHolderIdentificationType = "tax_id"
)

type AccountHolderType string

const (
	IndividualType AccountHolderType = "individual"
	Corporate      AccountHolderType = "corporate"
	Government     AccountHolderType = "government"
)

type (
	AccountHolder struct {
		Type              AccountHolderType           `json:"type,omitempty"`
		FirstName         string                      `json:"first_name,omitempty"`
		LastName          string                      `json:"last_name,omitempty"`
		CompanyName       string                      `json:"company_name,omitempty"`
		TaxId             string                      `json:"tax_id,omitempty"`
		DateOfBirth       *DateOfBirth                `json:"date_of_birth,omitempty"`
		CountryOfBirth    common.Country              `json:"country_of_birth,omitempty"`
		ResidentialStatus string                      `json:"residential_status,omitempty"`
		BillingAddress    *common.Address             `json:"billing_address,omitempty"`
		Phone             *common.Phone               `json:"phone,omitempty"`
		Identification    AccountHolderIdentification `json:"identification,omitempty"`
		Email             string                      `json:"email,omitempty"`
	}

	AccountHolderIdentification struct {
		Type           AccountHolderIdentificationType `json:"type,omitempty"`
		Number         string                          `json:"number,omitempty"`
		IssuingCountry common.Country                  `json:"issuing_country,omitempty"`
	}
)
