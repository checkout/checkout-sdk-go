package accounts

import (
	"github.com/checkout/checkout-sdk-go/common"
)

const (
	accountsPath           = "accounts"
	entitiesPath           = "entities"
	membersPath            = "members"
	filesPath              = "files"
	instrumentsPath        = "instruments"
	payoutSchedulesPath    = "payout-schedules"
	paymentInstrumentsPath = "payment-instruments"
)

type AccountHolderType string

const (
	IndividualType AccountHolderType = "individual"
	Corporate      AccountHolderType = "corporate"
	Government     AccountHolderType = "government"
)

type InstrumentStatus string

const (
	Verified          InstrumentStatus = "verified"
	Unverified        InstrumentStatus = "unverified"
	InstrumentPending InstrumentStatus = "pending"
)

type (
	InstrumentDocument struct {
		Type   string `json:"type,omitempty"`
		FileId string `json:"file_id,omitempty"`
	}
)

type (
	AccountHolder struct {
		Type              AccountHolderType                   `json:"type,omitempty"`
		FirstName         string                              `json:"first_name,omitempty"`
		LastName          string                              `json:"last_name,omitempty"`
		CompanyName       string                              `json:"company_name,omitempty"`
		TaxId             string                              `json:"tax_id,omitempty"`
		DateOfBirth       *DateOfBirth                        `json:"date_of_birth,omitempty"`
		CountryOfBirth    common.Country                      `json:"country_of_birth,omitempty"`
		ResidentialStatus string                              `json:"residential_status,omitempty"`
		BillingAddress    *common.Address                     `json:"billing_address,omitempty"`
		Phone             *common.Phone                       `json:"phone,omitempty"`
		Identification    *common.AccountHolderIdentification `json:"identification,omitempty"`
		Email             string                              `json:"email,omitempty"`
	}
)
