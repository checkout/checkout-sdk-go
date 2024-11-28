package common

import (
	"net/http"
	"time"
)

type CardWalletType string

const (
	Applepay  CardWalletType = "applepay"
	Googlepay CardWalletType = "googlepay"
)

type AccountType string

const (
	Savings AccountType = "savings"
	Current AccountType = "current"
	Cash    AccountType = "cash"
)

type CardType string

const (
	Charge        CardType = "Charge"
	Credit        CardType = "Credit"
	Debit         CardType = "Debit"
	DeferredDebit CardType = "Deferred Debit"
	Prepaid       CardType = "Prepaid"
)

type CardCategory string

const (
	All               CardCategory = "All"
	Commercial        CardCategory = "Commercial"
	Consumer          CardCategory = "Consumer"
	NotSet            CardCategory = "NotSet"
	OtherCardCategory CardCategory = "Other"
)

type AccountHolderType string

const (
	Individual AccountHolderType = "individual"
	Corporate  AccountHolderType = "corporate"
	Government AccountHolderType = "government"
)

type ChallengeIndicator string

const (
	ChallengeRequested        ChallengeIndicator = "challenge_requested"
	ChallengeRequestedMandate ChallengeIndicator = "challenge_requested_mandate"
	NoChallengeRequested      ChallengeIndicator = "no_challenge_requested"
	NoPreference              ChallengeIndicator = "no_preference"
)

type DocumentType string

const (
	PassportDocumentType DocumentType = "passport"
	NationalIdentityCard DocumentType = "national_identity_card"
	DrivingLicense       DocumentType = "driving_license"
	CitizenCard          DocumentType = "citizen_card"
	ResidencePermit      DocumentType = "residence_permit"
	ElectoralId          DocumentType = "electoral_id"
)

type AccountHolderIdentificationType string

const (
	Passport            AccountHolderIdentificationType = "passport"
	DrivingLicence      AccountHolderIdentificationType = "driving_licence"
	NationalId          AccountHolderIdentificationType = "national_id"
	CompanyRegistration AccountHolderIdentificationType = "company_registration"
	TaxId               AccountHolderIdentificationType = "tax_id"
)

type ThreeDsFlowType string

const (
	Challenged            ThreeDsFlowType = "challenged"
	Frictionless          ThreeDsFlowType = "frictionless"
	FrictionlessDelegated ThreeDsFlowType = "frictionless_delegated"
)

type Exemption string

const (
	LowRiskProgram            Exemption = "low_risk_program"
	LowValue                  Exemption = "low_value"
	None                      Exemption = "none"
	Other                     Exemption = "other"
	OutOfScaScope             Exemption = "out_of_sca_scope"
	RecurringOperation        Exemption = "recurring_operation"
	ScaDelegation             Exemption = "sca_delegation"
	SecureCorporatePayment    Exemption = "secure_corporate_payment"
	ThreeDsOutage             Exemption = "3ds_outage"
	TransactionRiskAssessment Exemption = "transaction_risk_assessment"
	TrustedListing            Exemption = "trusted_listing"
	TrustedListingPrompt      Exemption = "trusted_listing_prompt"
)

type ThreeDsMethodCompletion string

const (
	Y ThreeDsMethodCompletion = "y"
	N ThreeDsMethodCompletion = "n"
	U ThreeDsMethodCompletion = "u"
)

type (
	Address struct {
		AddressLine1 string  `json:"address_line1,omitempty"`
		AddressLine2 string  `json:"address_line2,omitempty"`
		City         string  `json:"city,omitempty"`
		State        string  `json:"state,omitempty"`
		Zip          string  `json:"zip,omitempty"`
		Country      Country `json:"country,omitempty"`
	}

	Phone struct {
		CountryCode string `json:"country_code,omitempty"`
		Number      string `json:"number,omitempty"`
	}

	BankDetails struct {
		Name    string   `json:"name,omitempty"`
		Branch  string   `json:"branch,omitempty"`
		Address *Address `json:"address,omitempty"`
	}
)

type (
	IdResponse struct {
		HttpMetadata HttpMetadata
		Id           string          `json:"id,omitempty"`
		Links        map[string]Link `json:"_links"`
	}

	MetadataResponse struct {
		HttpMetadata HttpMetadata
	}

	Data interface{}

	ContentResponse struct {
		HttpMetadata HttpMetadata
		Content      Data `json:"content,omitempty"`
	}

	HttpMetadata struct {
		Status       string     `json:"status,omitempty"`
		StatusCode   int        `json:"status_code,omitempty"`
		ResponseBody []byte     `json:"response_body,omitempty"`
		ResponseCSV  [][]string `json:"response_csv,omitempty"`
		Headers      *Headers   `json:"headers,omitempty"`
	}

	AlternativeResponse map[string]interface{}

	Headers struct {
		Header       http.Header
		CKORequestID *string `json:"cko-request-id,omitempty"`
		CKOVersion   *string `json:"cko-version,omitempty"`
	}

	Link struct {
		HRef  *string `json:"href,omitempty"`
		Title *string `json:"title,omitempty"`
	}
)

type (
	AccountHolderIdentification struct {
		Type           AccountHolderIdentificationType `json:"type,omitempty"`
		Number         string                          `json:"number,omitempty"`
		IssuingCountry Country                         `json:"issuing_country,omitempty"`
		DateOfExpiry   string                          `json:"date_of_expiry,omitempty"`
	}

	AccountHolder struct {
		Type              AccountHolderType            `json:"type,omitempty"`
		Title             string                       `json:"title,omitempty"`
		FullName          string                       `json:"full_name,omitempty"`
		FirstName         string                       `json:"first_name,omitempty"`
		MiddleName        string                       `json:"middle_name,omitempty"`
		LastName          string                       `json:"last_name,omitempty"`
		Email             string                       `json:"email,omitempty"`
		Gender            string                       `json:"gender,omitempty"`
		CompanyName       string                       `json:"company_name,omitempty"`
		TaxId             string                       `json:"tax_id,omitempty"`
		DateOfBirth       string                       `json:"date_of_birth,omitempty"`
		CountryOfBirth    string                       `json:"country_of_birth,omitempty"`
		ResidentialStatus string                       `json:"residential_status,omitempty"`
		BillingAddress    *Address                     `json:"billing_address,omitempty"`
		Phone             *Phone                       `json:"phone,omitempty"`
		Identification    *AccountHolderIdentification `json:"identification,omitempty"`
	}
)

type InstrumentType string

const (
	Card        InstrumentType = "card"
	BankAccount InstrumentType = "bank_account"
	Token       InstrumentType = "token"
	Sepa        InstrumentType = "sepa"
	CardToken   InstrumentType = "card_token"
)

type (
	InstrumentDetails struct {
		Id                         string                      `json:"id,omitempty"`
		Fingerprint                string                      `json:"fingerprint,omitempty"`
		InstrumentCustomerResponse *InstrumentCustomerResponse `json:"customer,omitempty"`
		AccountHolder              *AccountHolder              `json:"account_holder,omitempty"`
	}

	InstrumentCustomerResponse struct {
		Id      string `json:"id,omitempty"`
		Email   string `json:"email,omitempty"`
		Name    string `json:"name,omitempty"`
		Phone   *Phone `json:"phone,omitempty"`
		Default bool   `json:"default,omitempty"`
	}
)

type (
	CustomerRequest struct {
		Id        string `json:"id,omitempty"`
		Email     string `json:"email,omitempty"`
		Name      string `json:"name,omitempty"`
		TaxNumber string `json:"tax_number,omitempty"`
		Phone     *Phone `json:"phone,omitempty"`
		Default   bool   `json:"default,omitempty"`
	}

	CustomerResponse struct {
		Id    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
		Name  string `json:"name,omitempty"`
		Phone *Phone `json:"phone,omitempty"`
	}

	UpdateCustomerRequest struct {
		Id      string `json:"id,omitempty"`
		Default bool   `json:"default,omitempty"`
	}
)

type (
	// Deprecated: should use AmountAllocations instead
	MarketplaceData struct {
		SubEntityId string              `json:"sub_entity_id,omitempty"`
		SubEntities []AmountAllocations `json:"sub_entities,omitempty"`
	}

	AmountAllocations struct {
		Id         string      `json:"id,omitempty"`
		Amount     int64       `json:"amount"`
		Reference  string      `json:"reference,omitempty"`
		Commission *Commission `json:"commission,omitempty"`
	}

	Commission struct {
		Amount     int64   `json:"amount,omitempty"`
		Percentage float32 `json:"percentage,omitempty"`
	}
)

type (
	DateRangeQuery struct {
		From *time.Time `url:"from,omitempty" layout:"2006-01-02T15:04:05Z"`
		To   *time.Time `url:"to,omitempty" layout:"2006-01-02T15:04:05Z"`
	}
)

type (
	Destination struct {
		AccountType   AccountType    `json:"account_type,omitempty"`
		AccountNumber string         `json:"account_number,omitempty"`
		BankCode      string         `json:"bank_code,omitempty"`
		BranchCode    string         `json:"branch_code,omitempty"`
		Iban          string         `json:"iban,omitempty"`
		Bban          string         `json:"bban,omitempty"`
		SwiftBic      string         `json:"swift_bic,omitempty"`
		Country       Country        `json:"country,omitempty"`
		AccountHolder *AccountHolder `json:"account_holder,omitempty"`
		Bank          *BankDetails   `json:"bank,omitempty"`
	}
)

type CardholderAccountAgeIndicatorType string

const (
	CardholderLessThanThirtyDays CardholderAccountAgeIndicatorType = "less_than_thirty_days"
	CardholderMoreThanSixtyDays  CardholderAccountAgeIndicatorType = "more_than_sixty_days"
	CardholderNoAccount          CardholderAccountAgeIndicatorType = "no_account"
	CardholderThirtyToSixtyDays  CardholderAccountAgeIndicatorType = "thirty_to_sixty_days"
	CardholderThisTransaction    CardholderAccountAgeIndicatorType = "this_transaction"
)

type AccountChangeIndicatorType string

const (
	AccountChangeLessThanThirtyDays AccountChangeIndicatorType = "less_than_thirty_days"
	AccountChangeMoreThanSixtyDays  AccountChangeIndicatorType = "more_than_sixty_days"
	AccountChangeThirtyToSixtyDays  AccountChangeIndicatorType = "thirty_to_sixty_days"
	AccountChangeThisTransaction    AccountChangeIndicatorType = "this_transaction"
)

type AccountPasswordChangeIndicatorType string

const (
	PasswordChangeLessThanThirtyDays AccountPasswordChangeIndicatorType = "less_than_thirty_days"
	PasswordChangeMoreThanSixtyDays  AccountPasswordChangeIndicatorType = "more_than_sixty_days"
	PasswordChangeNoChange           AccountPasswordChangeIndicatorType = "no_change"
	PasswordChangeThirtyToSixtyDays  AccountPasswordChangeIndicatorType = "thirty_to_sixty_days"
	PasswordChangeThisTransaction    AccountPasswordChangeIndicatorType = "this_transaction"
)
