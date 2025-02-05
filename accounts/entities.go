package accounts

import (
	"time"

	"github.com/checkout/checkout-sdk-go/common"
)

type BusinessType string

const (
	GeneralPartnership             BusinessType = "general_partnership"
	LimitedPartnership             BusinessType = "limited_partnership"
	PublicLimitedCompany           BusinessType = "public_limited_company"
	LimitedCompany                 BusinessType = "limited_company"
	ProfessionalAssociation        BusinessType = "professional_association"
	UnincorporatedAssociation      BusinessType = "unincorporated_association"
	AutoEntrepreneur               BusinessType = "auto_entrepreneur"
	ScottishLimitedPartnership     BusinessType = "scottish_limited_partnership"
	PrivateCorporation             BusinessType = "private_corporation"
	LimitedLiabilityCorporation    BusinessType = "limited_liability_corporation"
	PubliclyTradedCorporation      BusinessType = "publicly_traded_corporation"
	RegulatedFinancialInstitution  BusinessType = "regulated_financial_institution"
	SecRegisteredEntity            BusinessType = "sec_registered_entity"
	CftcRegisteredEntity           BusinessType = "cftc_registered_entity"
	IndividualOrSoleProprietorship BusinessType = "individual_or_sole_proprietorship"
	GovernmentAgency               BusinessType = "government_agency"
	NonProfitEntity                BusinessType = "non_profit_entity"
	Trust                          BusinessType = "trust"
	ClubOrSociety                  BusinessType = "club_or_society"
)

type CompanyPositionType string

const (
	CEOCPStringType                        CompanyPositionType = "ceo"
	CFOCPStringType                        CompanyPositionType = "cfo"
	COOCPStringType                        CompanyPositionType = "coo"
	ManagingMemberCPStringType             CompanyPositionType = "managing_member"
	GeneralPartnerCPStringType             CompanyPositionType = "general_partner"
	PresidentCPStringType                  CompanyPositionType = "president"
	VicePresidentCPStringType              CompanyPositionType = "vice_president"
	TreasurerCPStringType                  CompanyPositionType = "treasurer"
	OtherSeniorManagementCPStringType      CompanyPositionType = "other_senior_management"
	OtherExecutiveOfficerCPStringType      CompanyPositionType = "other_executive_officer"
	OtherNonExecutiveNonSeniorCPStringType CompanyPositionType = "other_non_executive_non_senior"
)

type EntityRoles string

const (
	UboERStringType                 EntityRoles = "ubo"
	AuthorisedSignatoryERStringType EntityRoles = "authorised_signatory"
	DirectorERStringType            EntityRoles = "director"
	ControlPersonERStringType       EntityRoles = "control_person"
	LegalRepresentativeERStringType EntityRoles = "legal_representative"
)

type IdentityVerificationType string

const (
	PassportIVStringType             IdentityVerificationType = "passport"
	NationalIdentityCardIVStringType IdentityVerificationType = "national_identity_card"
	DrivingLicenseIVStringType       IdentityVerificationType = "driving_license"
	CitizenCardIVStringType          IdentityVerificationType = "citizen_card"
	ResidencePermitIVStringType      IdentityVerificationType = "residence_permit"
	ElectoralIdIVStringType          IdentityVerificationType = "electoral_id"
)

type CompanyVerificationType string

const (
	IncorporationDocumentCVStringType CompanyVerificationType = "incorporation_document"
	ArticlesOfAssociationCVStringType CompanyVerificationType = "articles_of_association"
)

type ArticlesOfAssociationType string

const (
	MemorandumOfAssociationAOSStringType ArticlesOfAssociationType = "memorandum_of_association"
	ArticlesOfAssociationAOSStringType   ArticlesOfAssociationType = "articles_of_association"
)

type BankVerificationType string

const (
	BankStatementBVStringType BankVerificationType = "bank_statement"
)

type TaxVerificationType string

const (
	EinLetterTVStringType TaxVerificationType = "ein_letter"
)

type FinancialVerificationType string

const (
	FinancialStatementFVStringType FinancialVerificationType = "financial_statement"
)

type ProofOfLegalityType string

const (
	ProofOfLegalityPOLStringType ProofOfLegalityType = "proof_of_legality"
)

type ProofOfPrincipalAddressType string

const (
	ProofOfAddressPOPAStringType ProofOfPrincipalAddressType = "proof_of_address"
)

type ProofOfResidentialAddressType string

const (
	ProofOfAddressPORAStringType ProofOfResidentialAddressType = "proof_of_address"
)

type ShareholderStructureType string

const (
	CertifiedShareholderStructureSHSStringType ShareholderStructureType = "certified_shareholder_structure"
)

type CertifiedAuthorisedSignatoryType string

const (
	PowerOfAttorneyCASStringType CertifiedAuthorisedSignatoryType = "power_of_attorney"
)

type ProofOfRegistrationType string

const (
	ExtractFromTradeRegisterPORStringType ProofOfRegistrationType = "extract_from_trade_register"
	OtherPORStringType                    ProofOfRegistrationType = "other"
)

type OnboardingStatus string

const (
	Draft          OnboardingStatus = "draft"
	Active         OnboardingStatus = "active"
	Pending        OnboardingStatus = "pending"
	Restricted     OnboardingStatus = "restricted"
	RequirementDue OnboardingStatus = "requirements_due"
	Inactive       OnboardingStatus = "inactive"
	Rejected       OnboardingStatus = "rejected"
)

type (
	Phone struct {
		CountryCode common.Country `json:"country_code,omitempty"`
		Number      string         `json:"number,omitempty"`
	}
)

type (
	Profile struct {
		Urls                   []string          `json:"urls,omitempty"`
		Mccs                   []string          `json:"mccs,omitempty"`
		DefaultHoldingCurrency common.Currency   `json:"default_holding_currency,omitempty"`
		HoldingCurrencies      []common.Currency `json:"holding_currencies,omitempty"`
	}

	AdditionalInfo struct {
		Field1 string `json:"field1,omitempty"`
		Field2 string `json:"field2,omitempty"`
		Field3 string `json:"field3,omitempty"`
	}
)

type (
	RequirementsDue struct {
		Field   string `json:"field,omitempty"`
		Reason  string `json:"reason,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

type (
	SubEntityMemberData struct {
		UserId string `json:"user_id,omitempty"`
	}
)

type (
	ProcessingDetails struct {
		SettlementCountry       string          `json:"settlement_country,omitempty"`
		TargetCountries         []string        `json:"target_countries,omitempty"`
		AnnualProcessingVolume  int             `json:"annual_processing_volume,omitempty"`
		AverageTransactionValue int             `json:"average_transaction_value,omitempty"`
		HighestTransactionValue int             `json:"highest_transaction_value,omitempty"`
		Currency                common.Currency `json:"currency,omitempty"`
	}
)

type (
	ContactDetails struct {
		Invitee              *Invitee              `json:"invitee,omitempty"`
		Phone                *Phone                `json:"phone,omitempty"`
		EntityEmailAddresses *EntityEmailAddresses `json:"email_addresses,omitempty"`
	}

	Invitee struct {
		Email string `json:"email,omitempty"`
	}

	EntityEmailAddresses struct {
		Primary string `json:"primary,omitempty"`
	}
)

type (
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
		DateOfIncorporation        *DateOfIncorporation    `json:"date_of_incorporation,omitempty"`
		RegulatoryLicenseNumber    string                  `json:"regulatory_license_number,omitempty"`
		RegulatoryLicenceNumber    string                  `json:"regulatory_licence_number,omitempty"`
	}

	EntityDocument struct {
		Type   string `json:"type,omitempty"`
		FileId string `json:"file_id,omitempty"`
	}

	Representative struct {
		Id                  string                     `json:"id,omitempty"`
		OwnershipPercentage int                        `json:"ownership_percentage,omitempty"`
		FirstName           string                     `json:"first_name,omitempty"`
		MiddleName          string                     `json:"middle_name,omitempty"`
		LastName            string                     `json:"last_name,omitempty"`
		Address             *common.Address            `json:"address,omitempty"`
		Phone               *Phone                     `json:"phone,omitempty"`
		DateOfBirth         *DateOfBirth               `json:"date_of_birth,omitempty"`
		PlaceOfBirth        *PlaceOfBirth              `json:"place_of_birth,omitempty"`
		Roles               []EntityRoles              `json:"roles,omitempty"`
		Documents           *OnboardSubEntityDocuments `json:"documents,omitempty"`
		CompanyPosition     *CompanyPositionType       `json:"company_position,omitempty"`
		Individual          *Individual                `json:"individual,omitempty"`
		Company             *Company                   `json:"company,omitempty"`
	}

	DateOfIncorporation struct {
		Day   int `json:"day,omitempty"`
		Month int `json:"month,omitempty"`
		Year  int `json:"year,omitempty"`
	}

	OnboardSubEntityDocuments struct {
		IdentityVerification         *IdentityVerification         `json:"identity_verification,omitempty"`
		CompanyVerification          *CompanyVerification          `json:"company_verification,omitempty"`
		TaxVerification              *TaxVerification              `json:"tax_verification,omitempty"`
		ArticlesOfAssociation        *ArticlesOfAssociation        `json:"articles_of_association,omitempty"`
		ShareholderStructure         *ShareholderStructure         `json:"shareholder_structure,omitempty"`
		BankVerification             *BankVerification             `json:"bank_verification,omitempty"`
		ProofOfLegality              *ProofOfLegality              `json:"proof_of_legality,omitempty"`
		ProofOfPrincipalAddress      *ProofOfPrincipalAddress      `json:"proof_of_principal_address,omitempty"`
		AdditionalDocument1          *AdditionalDocument           `json:"additional_document1,omitempty"`
		AdditionalDocument2          *AdditionalDocument           `json:"additional_document2,omitempty"`
		AdditionalDocument3          *AdditionalDocument           `json:"additional_document3,omitempty"`
		CertifiedAuthorisedSignatory *CertifiedAuthorisedSignatory `json:"certified_authorised_signatory,omitempty"`
		ProofOfResidentialAddress    *ProofOfResidentialAddress    `json:"proof_of_residential_address,omitempty"`
		ProofOfRegistration          *ProofOfRegistration          `json:"proof_of_registration,omitempty"`
		FinancialVerification        *FinancialVerification        `json:"financial_verification,omitempty"`
	}

	IdentityVerification struct {
		Type  IdentityVerificationType `json:"type,omitempty"`
		Front string                   `json:"front,omitempty"`
		Back  string                   `json:"back,omitempty"`
	}

	CompanyVerification struct {
		Type  CompanyVerificationType `json:"type,omitempty"`
		Front string                  `json:"front,omitempty"`
	}

	TaxVerification struct {
		Type  TaxVerificationType `json:"type,omitempty"`
		Front string              `json:"front,omitempty"`
	}

	ArticlesOfAssociation struct {
		Type  ArticlesOfAssociationType `json:"type,omitempty"`
		Front string                    `json:"front,omitempty"`
	}

	ShareholderStructure struct {
		Type  ShareholderStructureType `json:"type,omitempty"`
		Front string                   `json:"front,omitempty"`
	}

	BankVerification struct {
		Type  BankVerificationType `json:"type,omitempty"`
		Front string               `json:"front,omitempty"`
	}

	ProofOfLegality struct {
		Type  ProofOfLegalityType `json:"type,omitempty"`
		Front string              `json:"front,omitempty"`
	}

	ProofOfPrincipalAddress struct {
		Type  ProofOfPrincipalAddressType `json:"type,omitempty"`
		Front string                      `json:"front,omitempty"`
	}

	AdditionalDocument struct {
		Front string `json:"front,omitempty"`
	}

	CertifiedAuthorisedSignatory struct {
		Type  CertifiedAuthorisedSignatoryType `json:"type,omitempty"`
		Front string                           `json:"front,omitempty"`
		Back  string                           `json:"back,omitempty"`
	}

	ProofOfResidentialAddress struct {
		Type  ProofOfResidentialAddressType `json:"type,omitempty"`
		Front string                        `json:"front,omitempty"`
	}

	ProofOfRegistration struct {
		Type  ProofOfRegistrationType `json:"type,omitempty"`
		Front string                  `json:"front,omitempty"`
	}

	FinancialVerification struct {
		Type  FinancialVerificationType `json:"type,omitempty"`
		Front string                    `json:"front,omitempty"`
	}
)

type (
	Individual struct {
		FirstName         string                  `json:"first_name,omitempty"`
		MiddleName        string                  `json:"middle_name,omitempty"`
		LastName          string                  `json:"last_name,omitempty"`
		TradingName       string                  `json:"trading_name,omitempty"`
		LegalName         string                  `json:"legal_name,omitempty"`
		NationalTaxId     string                  `json:"national_tax_id,omitempty"`
		NationalIdNumber  string                  `json:"national_id_number,omitempty"`
		EmailAddress      string                  `json:"email_address,omitempty"`
		Phone             *Phone                  `json:"phone,omitempty"`
		Address           *common.Address         `json:"address,omitempty"`
		RegisteredAddress *common.Address         `json:"registered_address,omitempty"`
		DateOfBirth       *DateOfBirth            `json:"date_of_birth,omitempty"`
		PlaceOfBirth      *PlaceOfBirth           `json:"place_of_birth,omitempty"`
		Identification    *Identification         `json:"identification,omitempty"`
		FinancialDetails  *EntityFinancialDetails `json:"financial_details,omitempty"`
	}

	Identification struct {
		NationalIdNumber string                     `json:"national_id_number,omitempty"`
		Document         *OnboardSubEntityDocuments `json:"document,omitempty"`
	}

	DateOfBirth struct {
		Day   int `json:"day,omitempty"`
		Month int `json:"month,omitempty"`
		Year  int `json:"year,omitempty"`
	}

	PlaceOfBirth struct {
		Country common.Country `json:"country,omitempty"`
	}
)

type (
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
)

type (
	Instrument struct {
		Id       string              `json:"id,omitempty"`
		Label    string              `json:"label,omitempty"`
		Status   InstrumentStatus    `json:"status,omitempty"`
		Document *InstrumentDocument `json:"document,omitempty"`
	}
)

type (
	EntityFinancialDetails struct {
		AnnualProcessingVolume  int64                     `json:"annual_processing_volume,omitempty"`
		AverageTransactionValue int64                     `json:"average_transaction_value,omitempty"`
		HighestTransactionValue int64                     `json:"highest_transaction_value,omitempty"`
		Documents               *EntityFinancialDocuments `json:"documents,omitempty"`
		Currency                common.Currency           `json:"currency,omitempty"`
	}

	EntityFinancialDocuments struct {
		BankStatement      *EntityDocument `json:"bank_statement,omitempty"`
		FinancialStatement *EntityDocument `json:"financial_statement,omitempty"`
	}
)

type (
	OnboardEntityRequest struct {
		Reference         string                     `json:"reference,omitempty"`
		ContactDetails    *ContactDetails            `json:"contact_details,omitempty"`
		Profile           *Profile                   `json:"profile,omitempty"`
		Company           *Company                   `json:"company,omitempty"`
		Individual        *Individual                `json:"individual,omitempty"`
		Documents         *OnboardSubEntityDocuments `json:"documents,omitempty"`
		ProcessingDetails *ProcessingDetails         `json:"processing_details,omitempty"`
		Draft             bool                       `json:"draft,omitempty"`
		IsDraft           bool                       `json:"is_draft,omitempty"`
		AdditionalInfo    *AdditionalInfo            `json:"additional_info,omitempty"`
	}

	OnboardSubEntityRequest struct {
		Request map[string]interface{} `json:"-"`
	}
)

type (
	OnboardEntityResponse struct {
		HttpMetadata    common.HttpMetadata `json:"http_metadata,omitempty"`
		Id              string              `json:"id,omitempty"`
		Reference       string              `json:"reference,omitempty"`
		Status          OnboardingStatus    `json:"status,omitempty"`
		Capabilities    *Capabilities       `json:"capabilities,omitempty"`
		RequirementsDue []RequirementsDue   `json:"requirements_due,omitempty"`
	}

	OnboardEntityDetails struct {
		HttpMetadata    common.HttpMetadata `json:"http_metadata,omitempty"`
		Id              string              `json:"id,omitempty"`
		Reference       string              `json:"reference,omitempty"`
		Capabilities    *Capabilities       `json:"capabilities,omitempty"`
		Status          OnboardingStatus    `json:"status,omitempty"`
		RequirementsDue []RequirementsDue   `json:"requirements_due,omitempty"`
		ContactDetails  *ContactDetails     `json:"contact_details,omitempty"`
		Profile         *Profile            `json:"profile,omitempty"`
		Company         *Company            `json:"company,omitempty"`
		Individual      *Individual         `json:"individual,omitempty"`
		Instruments     []Instrument        `json:"instruments,omitempty"`
	}

	OnboardSubEntityResponse struct {
		HttpMetadata common.HttpMetadata    `json:"http_metadata,omitempty"`
		Response     map[string]interface{} `json:"-"`
	}

	OnboardSubEntityDetailsResponse struct {
		HttpMetadata common.HttpMetadata    `json:"http_metadata,omitempty"`
		Data         []SubEntityMemberData  `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links,omitempty"`
	}

	FileDetailsResponse struct {
		HttpMetadata  common.HttpMetadata    `json:"http_metadata,omitempty"`
		Id            string                 `json:"id,omitempty"`
		Status        string                 `json:"status,omitempty"`
		StatusReasons []string               `json:"status_reasons,omitempty"`
		Size          string                 `json:"size,omitempty"`
		MimeType      string                 `json:"mime_type,omitempty"`
		UploadedOn    *time.Time             `json:"uploaded_on,omitempty"`
		Filename      string                 `json:"filename,omitempty"`
		Purpose       string                 `json:"purpose,omitempty"`
		Links         map[string]common.Link `json:"_links,omitempty"`
	}

	UploadFileResponse struct {
		HttpMetadata            common.HttpMetadata    `json:"http_metadata,omitempty"`
		Id                      string                 `json:"id,omitempty"`
		MaximumSizeInBytes      int64                  `json:"maximum_size_in_bytes,omitempty"`
		DocumentTypesForPurpose []string               `json:"document_types_for_purpose,omitempty"`
		Links                   map[string]common.Link `json:"_links,omitempty"`
	}
)
