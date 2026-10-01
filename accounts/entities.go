package accounts

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// BusinessType is the legal type of the company. Must be individual_or_sole_proprietorship for the
// sole trader variants; which other values a variant accepts depends on the variant.
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

// CompanyPositionType is the position of a representative within the company (required for the
// control_person role).
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

// EntityRoles is a role a representative holds within the company. For sole traders, the only
// accepted role is ubo.
type EntityRoles string

const (
	UboERStringType                 EntityRoles = "ubo"
	AuthorisedSignatoryERStringType EntityRoles = "authorised_signatory"
	DirectorERStringType            EntityRoles = "director"
	ControlPersonERStringType       EntityRoles = "control_person"
	LegalRepresentativeERStringType EntityRoles = "legal_representative"
)

// NationalIdType is the classification of a representative's national identification number (US
// ISV Seller variants).
type NationalIdType string

const (
	Ssn             NationalIdType = "ssn"
	Itin            NationalIdType = "itin"
	Passport        NationalIdType = "passport"
	DrivingLicense  NationalIdType = "driving_license"
	NationalIdCard  NationalIdType = "national_id_card"
	ResidencePermit NationalIdType = "residence_permit"
	Other           NationalIdType = "other"
)

// IdentityVerificationType is the document type accepted to confirm an individual's identity.
type IdentityVerificationType string

const (
	PassportIVStringType             IdentityVerificationType = "passport"
	NationalIdentityCardIVStringType IdentityVerificationType = "national_identity_card"
	DrivingLicenseIVStringType       IdentityVerificationType = "driving_license"
	CitizenCardIVStringType          IdentityVerificationType = "citizen_card"
	ResidencePermitIVStringType      IdentityVerificationType = "residence_permit"
	ElectoralIdIVStringType          IdentityVerificationType = "electoral_id"
)

// CompanyVerificationType is the document type accepted as company verification.
// articles_of_association is accepted on the US Company (2.0) variants only.
type CompanyVerificationType string

const (
	IncorporationDocumentCVStringType CompanyVerificationType = "incorporation_document"
	ArticlesOfAssociationCVStringType CompanyVerificationType = "articles_of_association"
)

// ArticlesOfAssociationType is the document type accepted as memorandum or articles of association.
type ArticlesOfAssociationType string

const (
	MemorandumOfAssociationAOSStringType ArticlesOfAssociationType = "memorandum_of_association"
	ArticlesOfAssociationAOSStringType   ArticlesOfAssociationType = "articles_of_association"
)

// BankVerificationType is the document type accepted as bank verification.
type BankVerificationType string

const (
	BankStatementBVStringType BankVerificationType = "bank_statement"
)

// TaxVerificationType is the document type accepted as tax verification: an IRS-issued Employer
// Identification Number letter.
type TaxVerificationType string

const (
	EinLetterTVStringType TaxVerificationType = "ein_letter"
)

// FinancialVerificationType is the document type accepted as financial verification. Note the
// singular financial_statement; FinancialStatementsType is a different type.
type FinancialVerificationType string

const (
	FinancialStatementFVStringType FinancialVerificationType = "financial_statement"
)

// FinancialStatementsType is the document type accepted as financial statements (US ISV Seller
// variants). Note the plural financial_statements; FinancialVerificationType is a different type.
type FinancialStatementsType string

const (
	FinancialStatementsFSStringType FinancialStatementsType = "financial_statements"
)

// ProofOfLegalityType is the document type accepted as proof of legality.
type ProofOfLegalityType string

const (
	ProofOfLegalityPOLStringType ProofOfLegalityType = "proof_of_legality"
)

// ProofOfPrincipalAddressType is the document type accepted as proof of the company's principal
// place of business. Same proof_of_address value as ProofOfResidentialAddressType, but the API
// defines the two as separate enums on separate documents.
type ProofOfPrincipalAddressType string

const (
	ProofOfAddressPOPAStringType ProofOfPrincipalAddressType = "proof_of_address"
)

// ProofOfResidentialAddressType is the document type accepted as a representative's proof of
// residential address (EEA Sole Trader Full (3.0)). Same proof_of_address value as
// ProofOfPrincipalAddressType, but the API defines the two as separate enums on separate documents.
type ProofOfResidentialAddressType string

const (
	ProofOfAddressPORAStringType ProofOfResidentialAddressType = "proof_of_address"
)

// ShareholderStructureType is the document type accepted as a certified shareholder structure.
type ShareholderStructureType string

const (
	CertifiedShareholderStructureSHSStringType ShareholderStructureType = "certified_shareholder_structure"
)

// CertifiedAuthorisedSignatoryType is the document type accepted as a representative's certified
// authorised signatory document.
type CertifiedAuthorisedSignatoryType string

const (
	PowerOfAttorneyCASStringType CertifiedAuthorisedSignatoryType = "power_of_attorney"
)

// ProofOfRegistrationType is the document type accepted as a sole trader's proof of registration
// (EEA Sole Trader Full (3.0)).
type ProofOfRegistrationType string

const (
	ExtractFromTradeRegisterPORStringType ProofOfRegistrationType = "extract_from_trade_register"
	OtherPORStringType                    ProofOfRegistrationType = "other"
)

// OnboardingStatus is the onboarding status of a sub-entity. POST always returns draft.
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
	// Phone is a phone number on the Accounts API: the sub-entity's contact phone, or a
	// representative's phone. See ContactDetails.Phone for the per-variant number format.
	Phone struct {
		// The ISO 3166-1 alpha-2 country where the number is registered, not the dialling code.
		// [Required] on Accounts API v3.0; not part of the v2.0 schemas.
		CountryCode common.Country `json:"country_code,omitempty"`
		// The phone number, without the country calling code.
		// [Required]
		Number string `json:"number,omitempty"`
	}
)

type (
	// Citizenship is a citizenship or legal-status record (US ISV Seller variants).
	Citizenship struct {
		// The type of citizenship or legal status (for example citizenship or residency).
		// [Optional]
		Type string `json:"type,omitempty"`
		// The two-letter ISO 3166-1 alpha-2 country code.
		// [Required]
		// Format: iso-3166-1-alpha-2
		Country common.Country `json:"country,omitempty"`
	}
)

type (
	// Profile is information about the profile of the sub-entity, primarily regarding the products and
	// services offered.
	Profile struct {
		// A collection of website URLs the sub-entity accepts payments on.
		// [Required]
		// max 100 items; each ^(http|https):\/\/\S{2,293}$, Format: uri
		Urls []string `json:"urls,omitempty"`
		// The merchant category codes (4-digit ISO 18245) that most closely describe the business.
		// [Required]
		// min 1 item, max 5 items; each ^[0-9]{4}$
		Mccs []string `json:"mccs,omitempty"`
		// The default holding currency (ISO 4217).
		// [Required] on every v3.0 variant; [Optional] on the v2.0 variants.
		// Format: iso-4217
		DefaultHoldingCurrency common.Currency `json:"default_holding_currency,omitempty"`
		// The currencies incoming funds are held in.
		// [Required] on every v3.0 variant; [Optional] on the v2.0 variants.
		// min 1 item on v3.0; USD only on the US variants
		HoldingCurrencies []common.Currency `json:"holding_currencies,omitempty"`
	}

	// AdditionalInfo is not defined by any Accounts API schema.
	//
	// Deprecated: not part of any Accounts API schema; the API does not read it.
	AdditionalInfo struct {
		// Deprecated: not defined by any Accounts API schema.
		Field1 string `json:"field1,omitempty"`
		// Deprecated: not defined by any Accounts API schema.
		Field2 string `json:"field2,omitempty"`
		// Deprecated: not defined by any Accounts API schema.
		Field3 string `json:"field3,omitempty"`
	}
)

type (
	// RequirementsDue is a field that needs attention before the sub-entity can be onboarded.
	RequirementsDue struct {
		// The field that needs to be addressed.
		Field string `json:"field,omitempty"`
		// The reason the field needs attention.
		Reason string `json:"reason,omitempty"`
		// A more descriptive message.
		Message string `json:"message,omitempty"`
	}
)

type (
	SubEntityMemberData struct {
		UserId string `json:"user_id,omitempty"`
	}
)

type (
	// ProcessingDetails is the sub-entity's expected processing, sent on onboarding requests (Accounts API
	// v3.0). For the GET response see EntityProcessingDetails, whose amounts are int64.
	ProcessingDetails struct {
		// The country code (iso-3166-1 alpha-2) where the settlement bank account is located.
		// [Required] on the EEA, GB and US Company and Sole Trader Full (3.0) variants; not part of the US
		// ISV Seller variants.
		// Format: iso-3166-1-alpha-2
		// 2 characters
		SettlementCountry string `json:"settlement_country,omitempty"`
		// Target country codes (iso-3166-1 alpha-2) with more than 10% expected volume processing with
		// Checkout.com.
		// [Required]
		// min 1 item, max 10 items
		TargetCountries []string `json:"target_countries,omitempty"`
		// The estimated annual processing volume. In minor units without decimals.
		// [Required]
		// min 0
		AnnualProcessingVolume int `json:"annual_processing_volume,omitempty"`
		// The expected average transaction value. In minor units without decimals.
		// [Required]
		// min 0
		AverageTransactionValue int `json:"average_transaction_value,omitempty"`
		// The average time in days between accepting payment and fulfilling the order.
		// [Required] on the US ISV Seller variants only.
		// min 0
		AverageOrderFulfillmentTime int `json:"average_order_fulfillment_time,omitempty"`
		// The expected highest transaction value. In minor units without decimals.
		// [Required] on the EEA, GB and US Company and Sole Trader Full (3.0) variants; not part of the US
		// ISV Seller variants.
		// min 0
		HighestTransactionValue int `json:"highest_transaction_value,omitempty"`
		// The currency used for the processing details provided.
		// [Required]
		Currency common.Currency `json:"currency,omitempty"`
		// Payment method-specific processing details.
		// [Required] on the US ISV Seller variants only.
		Payments *ProcessingDetailsPayments `json:"payments,omitempty"`
	}

	// ProcessingDetailsPayments holds payment method-specific processing details (US ISV Seller
	// variants).
	ProcessingDetailsPayments struct {
		// The ACH processing details.
		// [Required]
		Ach *ProcessingDetailsAch `json:"ach,omitempty"`
	}

	// ProcessingDetailsAch holds the expected ACH processing (US ISV Seller variants). All amounts are in
	// minor units without decimals.
	ProcessingDetailsAch struct {
		// The estimated annual ACH processing volume.
		// [Required]
		// min 0
		AnnualAchVolume int `json:"annual_ach_volume,omitempty"`
		// The expected average ACH transaction size.
		// [Required]
		// min 0
		AverageAchTransactionSize int `json:"average_ach_transaction_size,omitempty"`
		// The estimated monthly volume of ACH credit transactions (for example, refunds issued to
		// customers).
		// [Required]
		// min 0
		EstimatedMonthlyCreditVolume int `json:"estimated_monthly_credit_volume,omitempty"`
		// The average value of an ACH credit transaction (for example, a refund).
		// [Required]
		// min 0
		AverageCreditAmount int `json:"average_credit_amount,omitempty"`
	}
)

type (
	// ContactDetails holds the contact details of the sub-entity.
	ContactDetails struct {
		// The details of the user responsible for onboarding the sub-entity.
		// [Optional] (not part of the US ISV Seller variants)
		Invitee *Invitee `json:"invitee,omitempty"`
		// The phone number of the sub-entity.
		// [Required] for every Accounts API v2.0 variant and the US ISV Seller variants; [Optional] for the
		// other v3.0 variants.
		// On v3.0 CountryCode is required and is the ISO 3166-1 alpha-2 country where the number is
		// registered (for example FR), not the dialling code; v2.0 takes Number only. Number is the number
		// without the country calling code, and its format depends on the variant:
		//   v3.0 EEA: ^[0-9]{6,13}$, min 6 characters, max 13 characters
		//   v3.0 GB: ^[0-9]{7,11}$, min 7 characters, max 11 characters
		//   v3.0 US and US ISV Seller: ^[1-9][0-9]{9,16}$, min 10 characters, max 16 characters
		//   v2.0: ^[1-9][0-9]{7,15}$, min 8 characters, max 16 characters; on the US v2.0 variants
		//   ^[2-9]{1}[0-9]{9,15}$, min 10 characters
		Phone *Phone `json:"phone,omitempty"`
		// Email addresses for this sub-entity.
		// [Required] for every Accounts API v2.0 variant and the US ISV Seller variants; [Optional] for the
		// other v3.0 variants.
		EntityEmailAddresses *EntityEmailAddresses `json:"email_addresses,omitempty"`
	}

	// Invitee holds the details of the user responsible for onboarding the sub-entity.
	Invitee struct {
		// The main email address for this sub-entity. Despite the spec's wording, this is the address of
		// the invitee, the user responsible for onboarding the sub-entity.
		// [Optional]
		// Format: email
		Email string `json:"email,omitempty"`
	}

	// EntityEmailAddresses holds the email addresses for this sub-entity.
	EntityEmailAddresses struct {
		// The main email address for this sub-entity.
		// [Required]
		// Format: email
		Primary string `json:"primary,omitempty"`
	}
)

type (
	// Company holds information about the company represented by the sub-entity: on every company and
	// v3.0 sole trader variant, and as the controlling company of a Representative (where only
	// LegalName, TradingName and RegisteredAddress apply).
	Company struct {
		// The sub-entity's business registration number: a Commercial Registration or Ministry of Commerce
		// certificate number, or an equivalent registration number.
		// [Required] for the Full variants and US ISV Seller Company (3.0); [Optional] for the Lite (2.0)
		// variants. Not part of the sole trader variants.
		// The format depends on the variant:
		//   EEA: min 2 characters, max 39 characters; a SIRET number for sub-entities based in France.
		//   GB (3.0): a Companies House number, 8 characters, matching one of the three alternatives of the
		//   spec's pattern, ^(A|B|C)$:
		//     A: ((AC|CE|CS|FC|FE|GE|GS|IC|LP|NC|NF|NI|NL|NO|NP|OC|OE|PC|R0|RC|SA|SC|SE|SF|SG|SI|SL|SO|SR|SZ|ZC|\d{2})\d{6})
		//     B: ((IP|SP|RS)[A-Z\d]{6})
		//     C: (SL\d{5}[\dA])
		//   GB (2.0) accepts the same pattern case-insensitively.
		//   US: an Employer Identification Number (EIN), ^[0-9]{9}$, 9 characters; US ISV Seller Company
		//   (3.0) also accepts the hyphenated form, ^[0-9]{2}-?[0-9]{7}$, min 9 characters, max 11.
		BusinessRegistrationNumber string `json:"business_registration_number,omitempty"`
		// The legal type of the company. Must be individual_or_sole_proprietorship for the sole trader
		// variants.
		// [Required], except on EEA and US Company Lite (2.0) where it is [Optional]. Not part of GB
		// Company Full and Lite (2.0).
		BusinessType BusinessType `json:"business_type,omitempty"`
		// The legal name of the sub-entity.
		// [Required] for every company variant and the controlling company; not part of the sole trader
		// variants.
		// min 2 characters, max 300 characters
		LegalName string `json:"legal_name,omitempty"`
		// The trading name of the sub-entity, also referred to as 'doing business as'.
		// [Required]
		// min 2 characters, max 300 characters
		TradingName string `json:"trading_name,omitempty"`
		// The collection of additional trading names for the sub-entity.
		// [Optional] (US ISV Seller variants only)
		AdditionalTradingNames []string `json:"additional_trading_names,omitempty"`
		// Indicates whether the sub-entity is a registered legal entity. Must be false for US ISV Seller
		// Sole Trader (3.0).
		// [Required] for US ISV Seller Sole Trader (3.0); not part of the other variants.
		IsRegisteredCompany *bool `json:"is_registered_company,omitempty"`
		// The primary location where business is performed.
		// [Required] for every company and v3.0 sole trader variant.
		PrincipalAddress *common.Address `json:"principal_address,omitempty"`
		// The registered address of the company.
		// [Required] for every company variant and the controlling company; not part of the sole trader
		// variants.
		RegisteredAddress *common.Address `json:"registered_address,omitempty"`
		// Deprecated: not defined by any Accounts API company schema; the API does not read it. Company
		// documents go on the top-level request documents instead.
		Document *EntityDocument `json:"document,omitempty"`
		// Information about the representatives of this company.
		// [Required]
		// min 1 item; max 1 item for the sole trader variants (the individual themselves, with roles
		// [ubo]), max 5 on v2.0, max 25 on EEA, GB and US Company Full (3.0), no maximum on US ISV Seller
		// Company (3.0)
		Representatives []Representative `json:"representatives,omitempty"`
		// Seller financial questions.
		// [Required] for EEA and US Company Full (2.0); [Optional] for EEA and US Company Lite (2.0). Not
		// part of the other variants.
		FinancialDetails *EntityFinancialDetails `json:"financial_details,omitempty"`
		// The date the company was incorporated, or the date the sole trader started trading.
		// [Required] for every v3.0 variant; [Optional] for EEA, GB and US Company Full (2.0).
		DateOfIncorporation *DateOfIncorporation `json:"date_of_incorporation,omitempty"`

		// Deprecated: not part of the Accounts API schema (US spelling serializes to the non-existent
		// "regulatory_license_number" key). Use RegulatoryLicenceNumber ("regulatory_licence_number").
		RegulatoryLicenseNumber string `json:"regulatory_license_number,omitempty"`
		// The regulatory licence number of the company.
		// [Optional] (EEA Company Full (3.0) only)
		// ^[a-zA-Z0-9\-]+$
		// min 4 characters, max 32 characters
		RegulatoryLicenceNumber string `json:"regulatory_licence_number,omitempty"`
	}

	// EntityDocument is not defined by any Accounts API onboarding schema. It is referenced only by the
	// deprecated Company.Document and EntityFinancialDocuments.
	//
	// Deprecated: not part of any Accounts API onboarding schema.
	EntityDocument struct {
		// Deprecated: not defined by any Accounts API onboarding schema.
		Type string `json:"type,omitempty"`
		// Deprecated: not defined by any Accounts API onboarding schema.
		FileId string `json:"file_id,omitempty"`
	}

	// Representative is a representative of the sub-entity. One struct covers every shape the Accounts
	// API defines: the v3.0 person of interest (Individual, Roles, CompanyPosition,
	// OwnershipPercentage, Documents), the v3.0 controlling company of EEA and GB Company Full (Company,
	// OwnershipPercentage), and the v2.0 company representative (the flat person fields, Roles,
	// Documents and, on the US variants, Identification).
	Representative struct {
		// The representative's id.
		// [Optional]
		// ^rep_[a-z0-9]{26}$
		// 30 characters
		Id string `json:"id,omitempty"`
		// The percentage ownership of the UBO or controlling company (required when over 25%).
		// [Optional]
		// min 25, max 100 on the EEA, GB and US Company Full (3.0) variants; min 0, max 100 on the US ISV
		// Seller variants
		OwnershipPercentage int `json:"ownership_percentage,omitempty"`
		// The representative's first name. Accounts API v2.0 only; on v3.0 use Individual.
		// [Required] (v2.0)
		// min 2 characters, max 50 characters
		FirstName string `json:"first_name,omitempty"`
		// The representative's middle name. Required if it appears in official documents. Accounts API
		// v2.0 only; on v3.0 use Individual.
		// [Optional]
		// min 2 characters, max 50 characters
		MiddleName string `json:"middle_name,omitempty"`
		// The representative's last name. Accounts API v2.0 only; on v3.0 use Individual.
		// [Required] (v2.0)
		// min 2 characters, max 50 characters
		LastName string `json:"last_name,omitempty"`
		// The representative's address. Accounts API v2.0 only; on v3.0 use Individual.
		// [Required] (v2.0)
		Address *common.Address `json:"address,omitempty"`
		// The representative's phone number. Accounts API v2.0 only; on v3.0 use Individual.
		// [Optional]
		Phone *Phone `json:"phone,omitempty"`
		// The date of birth of the person according to the Gregorian calendar. Accounts API v2.0 only; on
		// v3.0 use Individual.
		// [Required] for the v2.0 Full variants; [Optional] for the v2.0 Lite variants.
		DateOfBirth *DateOfBirth `json:"date_of_birth,omitempty"`
		// The place of birth of the person. Accounts API v2.0 only; on v3.0 use Individual.
		// [Required] for EEA Company Full (2.0); [Optional] for EEA Company Lite (2.0). Not part of the
		// other v2.0 variants.
		PlaceOfBirth *PlaceOfBirth `json:"place_of_birth,omitempty"`
		// The representative's identification. Accounts API v2.0 US Company variants only.
		// [Required] for US Company Full (2.0); [Optional] for US Company Lite (2.0).
		Identification *Identification `json:"identification,omitempty"`
		// The individual's roles within the company. For sole traders, must be ubo only.
		// [Required] for every variant except EEA and US Company Lite (2.0), where it is [Optional].
		Roles []EntityRoles `json:"roles,omitempty"`
		// Verification documents for the individual representative. The API validates this object strictly
		// on v3.0: it accepts only identity_verification, certified_authorised_signatory,
		// proof_of_residential_address and proof_of_registration, and rejects any other key. See
		// OnboardSubEntityDocuments for which apply to each variant.
		// [Required] for the EEA, GB and US Sole Trader Full (3.0) variants and EEA Company Full (2.0);
		// [Optional] otherwise.
		Documents *OnboardSubEntityDocuments `json:"documents,omitempty"`
		// The position of the representative within the company (required for the control_person role).
		// [Optional] (EEA, GB and US Company Full (3.0) and US ISV Seller Company (3.0))
		CompanyPosition *CompanyPositionType `json:"company_position,omitempty"`
		// Information about the individual representing the sub-entity.
		// [Required] for every v3.0 person of interest.
		Individual *Individual `json:"individual,omitempty"`
		// The controlling company, when the representative is a company rather than an individual.
		// [Required] for a controlling company representative (EEA and GB Company Full (3.0) only).
		// The API reads only three fields here, all [Required]: LegalName, TradingName and
		// RegisteredAddress. Leave the other Company fields unset.
		Company *Company `json:"company,omitempty"`
	}

	// DateOfIncorporation is the date the company was incorporated, or the date the sole trader started
	// trading.
	DateOfIncorporation struct {
		// The day of the month the company was incorporated.
		// [Optional]
		// min 1, max 31
		Day int `json:"day,omitempty"`
		// The month the company was incorporated.
		// [Required]
		// min 1, max 12
		Month int `json:"month,omitempty"`
		// The year the company was incorporated.
		// [Required]
		// min 1500, max 2999
		Year int `json:"year,omitempty"`
	}

	// OnboardSubEntityDocuments holds verification documents for a sub-entity. This one struct serves two
	// different objects on the Accounts API, which accept different keys:
	//   - The top-level request documents (OnboardEntityRequest.Documents). The API ignores keys it does
	//     not recognise here rather than rejecting them, so a misplaced document is dropped silently.
	//   - A representative's documents (Representative.Documents). This object is strict: it accepts
	//     only IdentityVerification, CertifiedAuthorisedSignatory, ProofOfResidentialAddress and
	//     ProofOfRegistration, and rejects any other key.
	// Each field below says which of the two it belongs to.
	OnboardSubEntityDocuments struct {
		// The document to use to confirm the individual's identity. Valid in both objects:
		//   Representative: [Required] for the EEA, GB and US Sole Trader Full (3.0) variants; [Optional]
		//   for the company variants.
		//   Top level: [Required] for the six sole trader variants of Accounts API v2.0, the only variants
		//   that take it there.
		IdentityVerification *IdentityVerification `json:"identity_verification,omitempty"`
		// The document to use to confirm the company's identity (certified by a power of attorney within
		// the last 3 months). Top level only.
		// [Required] for EEA Company Full (2.0 and 3.0) and GB Company Full (2.0); [Optional] for the other
		// company variants and the US ISV Seller variants.
		CompanyVerification *CompanyVerification `json:"company_verification,omitempty"`
		// IRS-issued Employer Identification Number document used to verify the entity's tax
		// identification. Top level only.
		// [Optional] (US Company variants and the US ISV Seller variants only)
		TaxVerification *TaxVerification `json:"tax_verification,omitempty"`
		// Memorandum or Articles of Association document. Top level only.
		// [Required] for EEA and GB Company Full (3.0); [Optional] for US Company Full (3.0) and the US ISV
		// Seller variants.
		ArticlesOfAssociation *ArticlesOfAssociation `json:"articles_of_association,omitempty"`
		// Shareholder structure chart (including % of shares) certified by a competent authority individual
		// and dated within the last 3 months. Top level only.
		// [Required] for EEA and GB Company Full (3.0); [Optional] for US Company Full (3.0) and US ISV
		// Seller Company (3.0).
		ShareholderStructure *ShareholderStructure `json:"shareholder_structure,omitempty"`
		// A document showing transactions from the last 3 months. Top level only.
		// [Required] for EEA Company Full (3.0) and the EEA, GB and US Sole Trader Full (3.0) variants;
		// [Optional] for GB and US Company Full (3.0) and EEA Company Full and Lite (2.0).
		BankVerification *BankVerification `json:"bank_verification,omitempty"`
		// A regulatory licence document required for the company to operate (when applicable). Top level
		// only.
		// [Optional] (EEA, GB and US Company Full (3.0) and the US ISV Seller variants)
		ProofOfLegality *ProofOfLegality `json:"proof_of_legality,omitempty"`
		// Proof of the company's principal place of business. Top level only.
		// [Optional] (EEA, GB and US Company Full (3.0) and the US ISV Seller variants)
		ProofOfPrincipalAddress *ProofOfPrincipalAddress `json:"proof_of_principal_address,omitempty"`
		// Additional space for documents to be provided when requested. Top level only.
		// [Optional] (EEA, GB and US Company and Sole Trader Full (3.0); not the US ISV Seller variants)
		AdditionalDocument1 *AdditionalDocument `json:"additional_document1,omitempty"`
		// Additional space for documents to be provided when requested. Top level only.
		// [Optional] (EEA, GB and US Company and Sole Trader Full (3.0); not the US ISV Seller variants)
		AdditionalDocument2 *AdditionalDocument `json:"additional_document2,omitempty"`
		// Additional space for documents to be provided when requested. Top level only.
		// [Optional] (EEA, GB and US Company and Sole Trader Full (3.0); not the US ISV Seller variants)
		AdditionalDocument3 *AdditionalDocument `json:"additional_document3,omitempty"`
		// Certified authorised signatory document. Required when the legal representative or other role
		// owner is not registered on the certificate of incorporation. Representative only; not accepted at
		// the top level.
		// [Optional] (EEA, GB and US Company Full (3.0) and US ISV Seller Company (3.0))
		CertifiedAuthorisedSignatory *CertifiedAuthorisedSignatory `json:"certified_authorised_signatory,omitempty"`
		// Proof of residential address of the representative. Representative only; not accepted at the top
		// level.
		// [Required] for EEA Sole Trader Full (3.0), and only valid there.
		ProofOfResidentialAddress *ProofOfResidentialAddress `json:"proof_of_residential_address,omitempty"`
		// Proof of the sole trader's registration, for example an extract from a trade register.
		// Representative only; not accepted at the top level.
		// [Required] for EEA Sole Trader Full (3.0), and only valid there.
		ProofOfRegistration *ProofOfRegistration `json:"proof_of_registration,omitempty"`
		// Financial statement document. Becomes mandatory depending on the answer provided for
		// annual_processing_volume; the sub-entity's status changes to requirements_due when it is needed.
		// Top level only.
		// [Optional] (EEA Company Full and Lite (2.0) only)
		FinancialVerification *FinancialVerification `json:"financial_verification,omitempty"`
		// Audited or management-prepared financial statements (when applicable). Top level only.
		// [Optional] (US ISV Seller variants only)
		FinancialStatements *FinancialStatements `json:"financial_statements,omitempty"`
	}

	// IdentityVerification is the document to use to confirm an individual's identity.
	IdentityVerification struct {
		// The type of document used for identity verification.
		// [Required]
		Type IdentityVerificationType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
		// The ID of the back side of the document as represented within Checkout.com systems.
		// [Optional]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Back string `json:"back,omitempty"`
	}

	// CompanyVerification is the document to use to confirm the company's identity (certified by a power
	// of attorney within the last 3 months).
	CompanyVerification struct {
		// The type of document used for company verification. articles_of_association is accepted on
		// the US Company (2.0) variants only.
		// [Required]
		Type CompanyVerificationType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// TaxVerification is the IRS-issued Employer Identification Number document used to verify the
	// entity's tax identification (US variants).
	TaxVerification struct {
		// The type of IRS-issued document used for tax verification.
		// [Required]
		Type TaxVerificationType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// ArticlesOfAssociation is the memorandum or articles of association document.
	ArticlesOfAssociation struct {
		// The type of document used.
		// [Required]
		Type ArticlesOfAssociationType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// ShareholderStructure is the shareholder structure chart (including % of shares) certified by a
	// competent authority individual and dated within the last 3 months.
	ShareholderStructure struct {
		// The type of document.
		// [Required]
		Type ShareholderStructureType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// BankVerification is a document showing transactions from the last 3 months.
	BankVerification struct {
		// The type of document being used as bank verification.
		// [Required]
		Type BankVerificationType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// ProofOfLegality is a regulatory licence document required for the company to operate (when
	// applicable).
	ProofOfLegality struct {
		// The type of document used for proof of legality.
		// [Required]
		Type ProofOfLegalityType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// ProofOfPrincipalAddress is proof of the company's principal place of business.
	ProofOfPrincipalAddress struct {
		// The type of document being used as address verification.
		// [Required]
		Type ProofOfPrincipalAddressType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// AdditionalDocument is additional space for documents to be provided when requested. It carries a
	// file ID only; the API defines no document type for it.
	AdditionalDocument struct {
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// CertifiedAuthorisedSignatory is the certified authorised signatory document. Required when the legal
	// representative or other role owner is not registered on the certificate of incorporation.
	// Representative documents only (EEA, GB and US Company Full (3.0), US ISV Seller Company (3.0)).
	CertifiedAuthorisedSignatory struct {
		// The type of document.
		// [Required]
		Type CertifiedAuthorisedSignatoryType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
		// Deprecated: not defined by the Accounts API, certified_authorised_signatory has type and front
		// only. Retained so existing code keeps compiling.
		Back string `json:"back,omitempty"`
	}

	// ProofOfResidentialAddress is proof of residential address of the representative. Representative
	// documents only, EEA Sole Trader Full (3.0).
	ProofOfResidentialAddress struct {
		// The type of document being used as address verification.
		// [Required]
		Type ProofOfResidentialAddressType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// ProofOfRegistration is proof of the sole trader's registration, for example an extract from a trade
	// register. Representative documents only, EEA Sole Trader Full (3.0).
	ProofOfRegistration struct {
		// The type of document being used as proof of registration.
		// [Required]
		Type ProofOfRegistrationType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// FinancialVerification is the financial statement document. It becomes mandatory depending on the
	// answer provided for annual_processing_volume; the sub-entity's status changes to requirements_due
	// when it is needed.
	FinancialVerification struct {
		// The type of the file.
		// [Required]
		Type FinancialVerificationType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}

	// FinancialStatements is audited or management-prepared financial statements (when applicable). US
	// ISV Seller variants only.
	FinancialStatements struct {
		// The type of document.
		// [Required]
		Type FinancialStatementsType `json:"type,omitempty"`
		// The ID of the front side of the document as represented within Checkout.com systems.
		// [Required]
		// ^file_[a-z2-7]{26}$
		// 31 characters
		Front string `json:"front,omitempty"`
	}
)

type (
	// Individual is an individual on the Accounts API. One struct covers two objects: the v3.0
	// representative individual (Representative.Individual) and the v2.0 top-level individual of the sole
	// trader variants. Each field below says which of the two it belongs to.
	Individual struct {
		// The individual's first name.
		// [Required]
		// min 2 characters, max 50 characters
		FirstName string `json:"first_name,omitempty"`
		// The individual's middle name. Required if it appears in official documents.
		// [Optional]
		// min 2 characters, max 50 characters
		MiddleName string `json:"middle_name,omitempty"`
		// The individual's last name.
		// [Required]
		// min 2 characters, max 50 characters
		LastName string `json:"last_name,omitempty"`
		// The trading name of the sub-entity, also referred to as 'doing business as'. v2.0 sole traders
		// only.
		// [Required] (v2.0 sole traders)
		// min 2 characters, max 300 characters
		TradingName string `json:"trading_name,omitempty"`
		// Deprecated: not defined by the Accounts API for an individual; legal_name exists on the company
		// only.
		LegalName string `json:"legal_name,omitempty"`
		// Deprecated: not defined by any Accounts API schema; the API does not read it.
		NationalTaxId string `json:"national_tax_id,omitempty"`
		// The individual's national identification number. v3.0 only.
		// [Required] for the US ISV Seller variants; [Optional] for the other v3.0 variants.
		// The format depends on the variant:
		//   US ISV Seller: the number for the NationalIdType given. ^[a-zA-Z0-9\-]+$, min 5 characters,
		//   max 16 characters.
		//   Other v3.0 variants: a Social Security Number (SSN) or Individual Taxpayer Identification Number
		//   (ITIN), US residents only. ^\d{9}$, 9 characters.
		NationalIdNumber string `json:"national_id_number,omitempty"`
		// The classification of the national identification number provided. v3.0 only.
		// [Required] for the US ISV Seller variants only; not part of the other v3.0 schemas, leave unset
		// for them.
		NationalIdType NationalIdType `json:"national_id_type,omitempty"`
		// The list of citizenships or legal statuses for the representative. v3.0 only.
		// [Required] for the US ISV Seller variants only; not part of the other v3.0 schemas, leave unset
		// for them.
		Citizenships []Citizenship `json:"citizenships,omitempty"`
		// The representative's personal email address. v3.0 only.
		// [Required] for the US ISV Seller variants; [Optional] for the other v3.0 variants.
		// Format: email
		EmailAddress string `json:"email_address,omitempty"`
		// The representative's phone number. v3.0 only.
		// [Required] for the US ISV Seller variants; [Optional] for the other v3.0 variants.
		Phone *Phone `json:"phone,omitempty"`
		// The representative's address. v3.0 only.
		// [Required] (v3.0)
		Address *common.Address `json:"address,omitempty"`
		// The registered address of the sole trader's business. v2.0 sole traders only.
		// [Required] (v2.0 sole traders)
		RegisteredAddress *common.Address `json:"registered_address,omitempty"`
		// The date of birth of the person according to the Gregorian calendar.
		// [Required], except on GB Sole Trader Lite (2.0) where it is [Optional].
		DateOfBirth *DateOfBirth `json:"date_of_birth,omitempty"`
		// The place of birth of the person.
		// [Required] on every v3.0 variant and on EEA Sole Trader Full and Lite (2.0). Not part of the
		// other v2.0 variants.
		PlaceOfBirth *PlaceOfBirth `json:"place_of_birth,omitempty"`
		// The individual's identification. v2.0 US Sole Trader only.
		// [Required] for US Sole Trader Full (2.0); [Optional] for US Sole Trader Lite (2.0).
		Identification *Identification `json:"identification,omitempty"`
		// Seller financial questions. v2.0 US Sole Trader only.
		// [Required] for US Sole Trader Full (2.0); [Optional] for US Sole Trader Lite (2.0).
		FinancialDetails *EntityFinancialDetails `json:"financial_details,omitempty"`
	}

	// Identification is the identification of a representative or individual on the Accounts API v2.0
	// US variants.
	Identification struct {
		// Social Security Number (SSN), or Individual Taxpayer Identification Number (ITIN) for non-US
		// citizens.
		// [Required]
		// ^\d{9}$
		// 9 characters
		NationalIdNumber string `json:"national_id_number,omitempty"`
		// Deprecated: not defined by the Accounts API, the identification object carries
		// national_id_number only. Retained so existing code keeps compiling.
		Document *OnboardSubEntityDocuments `json:"document,omitempty"`
	}

	// DateOfBirth is the date of birth of the person according to the Gregorian calendar.
	DateOfBirth struct {
		// The calendar day of the month they were born.
		// [Required]
		// min 1, max 31
		Day int `json:"day,omitempty"`
		// The month of the year they were born.
		// [Required]
		// min 1, max 12
		Month int `json:"month,omitempty"`
		// The year they were born.
		// [Required]
		// min 1900, max 2999
		Year int `json:"year,omitempty"`
	}

	// PlaceOfBirth is the place of birth of the person.
	PlaceOfBirth struct {
		// The country code (iso-3166-1 alpha-2).
		// [Required]
		// Format: iso-3166-1-alpha-2
		Country common.Country `json:"country,omitempty"`
	}
)

type (
	// Capabilities is the capabilities of the entity, as returned by the API. The spec declares the
	// object without detailing it.
	Capabilities struct {
		// Whether the entity can process payments.
		Payments *Payments `json:"payments,omitempty"`
		// Whether the entity can receive payouts.
		Payouts *Payouts `json:"payouts,omitempty"`
		// Whether the entity can issue cards.
		Issuing *Issuing `json:"issuing,omitempty"`
	}

	// Payments is the payments capability of the entity.
	Payments struct {
		// Whether the capability is available.
		Available bool `json:"available,omitempty"`
		// Whether the capability is enabled.
		Enabled bool `json:"enabled,omitempty"`
	}

	// Payouts is the payouts capability of the entity.
	Payouts struct {
		// Whether the capability is available.
		Available bool `json:"available,omitempty"`
		// Whether the capability is enabled.
		Enabled bool `json:"enabled,omitempty"`
	}

	// Issuing is the card issuing capability of the entity.
	Issuing struct {
		// Whether the capability is available.
		Available bool `json:"available,omitempty"`
		// Whether the capability is enabled.
		Enabled bool `json:"enabled,omitempty"`
	}
)

type (
	// Instrument is a payment instrument of the sub-entity, as returned by GET /accounts/entities/{id}.
	// The spec declares the list without detailing its items.
	Instrument struct {
		// The ID of the payment instrument.
		Id string `json:"id,omitempty"`
		// The label of the payment instrument.
		Label string `json:"label,omitempty"`
		// The status of the payment instrument.
		Status InstrumentStatus `json:"status,omitempty"`
		// The document supplied with the payment instrument.
		Document *InstrumentDocument `json:"document,omitempty"`
	}
)

type (
	// EntityFinancialDetails holds the seller financial questions (financial_details): on the company of
	// EEA and US Company Full and Lite (2.0), and on the individual of US Sole Trader Full and Lite (2.0).
	EntityFinancialDetails struct {
		// The estimated annual processing volume. In minor units without decimals.
		// [Required] on the Full (2.0) variants; [Optional] on the Lite (2.0) variants.
		// min 0
		AnnualProcessingVolume int64 `json:"annual_processing_volume,omitempty"`
		// The expected average transaction value. In minor units without decimals.
		// [Required] on the Full (2.0) variants; [Optional] on the Lite (2.0) variants.
		// min 0
		AverageTransactionValue int64 `json:"average_transaction_value,omitempty"`
		// The expected highest transaction value. In minor units without decimals.
		// [Required] on the Full (2.0) variants; [Optional] on the Lite (2.0) variants.
		// min 0
		HighestTransactionValue int64 `json:"highest_transaction_value,omitempty"`
		// Deprecated: not defined by any Accounts API schema; the API does not read it. Supporting documents
		// go on the top-level request documents (OnboardSubEntityDocuments) instead.
		Documents *EntityFinancialDocuments `json:"documents,omitempty"`
		// The currency used for the financial details provided.
		// [Required] on US Company Full and US Sole Trader Full (2.0); [Optional] on the other variants.
		Currency common.Currency `json:"currency,omitempty"`
	}

	// EntityFinancialDocuments is not defined by any Accounts API schema: financial_details carries the
	// three amounts and the currency only.
	//
	// Deprecated: not part of any Accounts API schema.
	EntityFinancialDocuments struct {
		// Deprecated: not defined by any Accounts API schema.
		BankStatement *EntityDocument `json:"bank_statement,omitempty"`
		// Deprecated: not defined by any Accounts API schema.
		FinancialStatement *EntityDocument `json:"financial_statement,omitempty"`
	}
)

type (
	// OnboardEntityRequest is the request body of POST /accounts/entities and PUT
	// /accounts/entities/{id}. Which fields are required depends on the onboarding variant.
	OnboardEntityRequest struct {
		// A unique reference you can later use to identify the sub-entity.
		// [Required]
		// min 1 character, max 50 characters
		Reference string `json:"reference,omitempty"`
		// Contact details of this sub-entity.
		// [Required], except on EEA Company Full (3.0) where it is [Optional].
		ContactDetails *ContactDetails `json:"contact_details,omitempty"`
		// Information about the profile of the sub-entity.
		// [Required]
		Profile *Profile `json:"profile,omitempty"`
		// Information about the company represented by the sub-entity.
		// [Required] for every company and v3.0 sole trader variant.
		Company *Company `json:"company,omitempty"`
		// Information about the individual represented by the sub-entity. Accounts API v2.0 sole traders
		// only.
		// [Required] for the v2.0 sole trader variants.
		//
		// Deprecated: not used by the Accounts API v3.0 schema, where a sole trader is onboarded as Company
		// with a representative.
		Individual *Individual `json:"individual,omitempty"`
		// The top-level documents used to support the verification of the sub-entity's details.
		// [Required] on the EEA, GB and US Company and Sole Trader Full (3.0) variants, EEA Company Full
		// (2.0) and EEA Sole Trader Full (2.0); [Optional] otherwise.
		Documents *OnboardSubEntityDocuments `json:"documents,omitempty"`
		// Information about the sub-entity's expected processing.
		// [Required] for every v3.0 variant.
		ProcessingDetails *ProcessingDetails `json:"processing_details,omitempty"`
		// Whether the sub-entity should remain in draft on PUT, skipping due diligence checks. POST always
		// creates the entity in draft.
		// [Optional]
		IsDraft bool `json:"is_draft,omitempty"`
		// Deprecated: not defined by any Accounts API schema; the API does not read it.
		AdditionalInfo *AdditionalInfo `json:"additional_info,omitempty"`
		// The identifier of a seller category set up for your platform. Seller categories define the
		// pricing, capabilities and risk profile applied to sub-entities.
		// [Required] for the US ISV Seller variants only.
		SellerCategory string `json:"seller_category,omitempty"`
		// Details of the person who agreed to the terms and conditions on behalf of the sub-entity.
		// [Required] for the US ISV Seller variants only.
		AgreedTerms *AgreedTerms `json:"agreed_terms,omitempty"`
		// Deprecated: not defined by any Accounts API schema; the API does not read it.
		Submitter *Submitter `json:"submitter,omitempty"`
	}

	// AgreedTerms is the evidence of consent to Checkout.com onboarding: the person who agreed to the
	// terms and conditions (US ISV Seller variants).
	AgreedTerms struct {
		// Date and time the terms were agreed, in RFC 3339 or ISO 8601 format.
		// [Required]
		// Format: date-time
		Date string `json:"date,omitempty"`
		// IP address (IPv4 or IPv6) of the person at the time they agreed the terms.
		// [Required]
		IpAddress string `json:"ip_address,omitempty"`
		// First and last name of the person who agreed to the terms.
		// [Required]
		Name string `json:"name,omitempty"`
		// Email address of the person who agreed to the terms.
		// [Required]
		// Format: email
		Email string `json:"email,omitempty"`
		// Identifier of the terms version that was agreed.
		// [Required]
		Version string `json:"version,omitempty"`
	}

	OnboardSubEntityRequest struct {
		Request map[string]interface{} `json:"-"`
	}
)

type (
	// OnboardEntityResponse is the response of POST /accounts/entities and PUT /accounts/entities/{id}.
	OnboardEntityResponse struct {
		// The HTTP metadata of the response.
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		// The ID of the sub-entity.
		Id string `json:"id,omitempty"`
		// The reference supplied in the request.
		Reference string `json:"reference,omitempty"`
		// The onboarding status of the sub-entity; draft after POST.
		Status OnboardingStatus `json:"status,omitempty"`
		// The capabilities of the entity.
		Capabilities *Capabilities `json:"capabilities,omitempty"`
		// List of requirements due in order to be onboarded.
		RequirementsDue []RequirementsDue `json:"requirements_due,omitempty"`
	}

	// OnboardEntityDetails is the details of a sub-entity, as returned by GET /accounts/entities/{id}.
	OnboardEntityDetails struct {
		// The HTTP metadata of the response.
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		// The ID of the sub-entity.
		Id string `json:"id,omitempty"`
		// A unique reference you can later use to identify the sub-entity.
		Reference string `json:"reference,omitempty"`
		// The capabilities of the entity.
		Capabilities *Capabilities `json:"capabilities,omitempty"`
		// The onboarding status of the sub-entity.
		Status OnboardingStatus `json:"status,omitempty"`
		// List of requirements due in order to be onboarded.
		RequirementsDue []RequirementsDue `json:"requirements_due,omitempty"`
		// Contact details of this sub-entity.
		ContactDetails *ContactDetails `json:"contact_details,omitempty"`
		// Information about the profile of the sub-entity, primarily regarding the products and services
		// offered.
		Profile *Profile `json:"profile,omitempty"`
		// Information about the company represented by the sub-entity (company and v3.0 sole trader
		// variants).
		Company *Company `json:"company,omitempty"`
		// Information about the individual represented by the sub-entity (v2.0 sole trader variants).
		Individual *Individual `json:"individual,omitempty"`
		// The sub-entity's expected processing (Accounts API v3.0). Amounts are int64; see
		// EntityProcessingDetails.
		ProcessingDetails *EntityProcessingDetails `json:"processing_details,omitempty"`
		// The top-level documents used to support the verification of the sub-entity's details.
		// Representative documents are on Representative.Documents, under Company.
		Documents *OnboardSubEntityDocuments `json:"documents,omitempty"`
		// The sub-entity's payment instruments.
		Instruments []Instrument `json:"instruments,omitempty"`
	}

	// EntityProcessingDetails is the sub-entity's expected processing, as returned by
	// GET /accounts/entities/{id} (processing_details, Accounts API v3.0). A response-only type,
	// separate from the request's ProcessingDetails: the amounts are int64 here because the API
	// declares them as integers in minor units with no maximum.
	EntityProcessingDetails struct {
		// The country code (iso-3166-1 alpha-2) where the settlement bank account is located.
		// Format: iso-3166-1-alpha-2
		// 2 characters
		SettlementCountry string `json:"settlement_country,omitempty"`
		// Target country codes (iso-3166-1 alpha-2) with more than 10% expected volume processing with
		// Checkout.com.
		// min 1 item, max 10 items
		TargetCountries []string `json:"target_countries,omitempty"`
		// The estimated annual processing volume. In minor units without decimals.
		// min 0
		AnnualProcessingVolume int64 `json:"annual_processing_volume,omitempty"`
		// The expected average transaction value. In minor units without decimals.
		// min 0
		AverageTransactionValue int64 `json:"average_transaction_value,omitempty"`
		// The expected highest transaction value. In minor units without decimals.
		// min 0
		HighestTransactionValue int64 `json:"highest_transaction_value,omitempty"`
		// The currency used for the processing details provided.
		Currency common.Currency `json:"currency,omitempty"`
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

	// FileDetailsResponse is the details of a sub-entity's file, as returned by
	// GET /entities/{entityId}/files/{fileId}.
	FileDetailsResponse struct {
		// The HTTP metadata of the response.
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		// The ID of the file.
		Id string `json:"id,omitempty"`
		// The current status of the file.
		Status string `json:"status,omitempty"`
		// If Status is invalid, the reasons why the file was invalid; otherwise empty.
		StatusReasons []string `json:"status_reasons,omitempty"`
		// The size of the file, in KB.
		Size int64 `json:"size,omitempty"`
		// The MIME type of the file.
		MimeType string `json:"mime_type,omitempty"`
		// The date and time the file was uploaded, in ISO 8601 UTC format.
		// Format: date-time (RFC 3339)
		UploadedOn *time.Time `json:"uploaded_on,omitempty"`
		// The purpose of the file, as provided in the initial request.
		Purpose string `json:"purpose,omitempty"`
		// The links related to the file.
		Links map[string]common.Link `json:"_links,omitempty"`
	}

	// UploadFileResponse is the response of POST /entities/{entityId}/files: the file ID and the upload
	// link. The file content itself is sent to that link, not in the request.
	UploadFileResponse struct {
		// The HTTP metadata of the response.
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		// The file identifier.
		Id string `json:"id,omitempty"`
		// The maximum file size allowed, in bytes.
		MaximumSizeInBytes int64 `json:"maximum_size_in_bytes,omitempty"`
		// The MIME file types allowed for the document purpose provided on the initial request.
		DocumentTypesForPurpose []string `json:"document_types_for_purpose,omitempty"`
		// The links related to the file, including the upload link.
		Links map[string]common.Link `json:"_links,omitempty"`
	}
)

type EntityRequirementReason string

const (
	PeriodicReview EntityRequirementReason = "periodic_review"
	Attestation    EntityRequirementReason = "attestation"
)

type EntityRequirementPriority string

const (
	HighPriority     EntityRequirementPriority = "high"
	CriticalPriority EntityRequirementPriority = "critical"
)

type EntityRequirementUpdateStatus string

const (
	ProcessingStatus EntityRequirementUpdateStatus = "processing"
)

type (
	// Submitter is not defined by any Accounts API onboarding schema.
	//
	// Deprecated: not part of any Accounts API onboarding schema; the API does not read it.
	Submitter struct {
		// Deprecated: not defined by any Accounts API onboarding schema.
		IpAddress string `json:"ip_address,omitempty"`
	}

	EntityRequirementListItem struct {
		Id           string                    `json:"id,omitempty"`
		Resource     string                    `json:"resource,omitempty"`
		ResourceType string                    `json:"resource_type,omitempty"`
		Reason       EntityRequirementReason   `json:"reason,omitempty"`
		Priority     EntityRequirementPriority `json:"priority,omitempty"`
		Deadline     *time.Time                `json:"deadline,omitempty"`
		Urn          string                    `json:"urn,omitempty"`
		FieldPath    string                    `json:"field_path,omitempty"`
		FieldUrn     string                    `json:"field_urn,omitempty"`
		Metadata     map[string]interface{}    `json:"metadata,omitempty"`
		Links        map[string]common.Link    `json:"_links,omitempty"`
	}

	EntityRequirementDetails struct {
		EntityRequirementListItem
		Message string                 `json:"message,omitempty"`
		Schema  map[string]interface{} `json:"_schema,omitempty"`
	}

	EntityRequirementListResponse struct {
		HttpMetadata common.HttpMetadata         `json:"http_metadata,omitempty"`
		Data         []EntityRequirementListItem `json:"data,omitempty"`
	}

	EntityRequirementDetailsResponse struct {
		HttpMetadata common.HttpMetadata
		EntityRequirementDetails
	}

	EntityRequirementUpdateRequest struct {
		Value interface{} `json:"value"`
	}

	EntityRequirementUpdateResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                        `json:"id,omitempty"`
		Status       EntityRequirementUpdateStatus `json:"status,omitempty"`
		SubmittedAt  *time.Time                    `json:"submitted_at,omitempty"`
		Links        map[string]common.Link        `json:"_links,omitempty"`
	}
)
