package accounts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

func TestProcessingDetailsWithPayments_Roundtrip(t *testing.T) {
	details := ProcessingDetails{
		AnnualProcessingVolume:      1000000,
		AverageTransactionValue:     5000,
		AverageOrderFulfillmentTime: 3,
		HighestTransactionValue:     25000,
		Currency:                    common.Currency("GBP"),
		SettlementCountry:           "GB",
		TargetCountries:             []string{"GB"},
		Payments: &ProcessingDetailsPayments{
			Ach: &ProcessingDetailsAch{
				AnnualAchVolume:              1000000,
				AverageAchTransactionSize:    5000,
				EstimatedMonthlyCreditVolume: 100000,
				AverageCreditAmount:          5000,
			},
		},
	}

	body, err := json.Marshal(details)
	assert.NoError(t, err)
	s := string(body)
	assert.Contains(t, s, `"average_order_fulfillment_time":3`)
	assert.Contains(t, s, `"payments":{`)
	assert.Contains(t, s, `"ach":{`)
	assert.Contains(t, s, `"annual_ach_volume":1000000`)
	assert.Contains(t, s, `"average_ach_transaction_size":5000`)
	assert.Contains(t, s, `"estimated_monthly_credit_volume":100000`)
	assert.Contains(t, s, `"average_credit_amount":5000`)
}

func TestAgreedTerms_Roundtrip(t *testing.T) {
	agreedTerms := AgreedTerms{
		Date:      "2026-07-20T10:00:00Z",
		IpAddress: "203.0.113.42",
		Name:      "John Representative",
		Email:     "john@example.com",
		Version:   "1.0",
	}

	body, err := json.Marshal(agreedTerms)
	assert.NoError(t, err)
	s := string(body)
	assert.Contains(t, s, `"date":"2026-07-20T10:00:00Z"`)
	assert.Contains(t, s, `"ip_address":"203.0.113.42"`)
	assert.Contains(t, s, `"name":"John Representative"`)
	assert.Contains(t, s, `"email":"john@example.com"`)
	assert.Contains(t, s, `"version":"1.0"`)
}

func TestCompanyV3Fields_Roundtrip(t *testing.T) {
	isRegistered := true
	company := Company{
		LegalName:                  "Super Hero Masks Inc.",
		TradingName:                "Super Hero Masks",
		BusinessRegistrationNumber: "01234567",
		BusinessType:               LimitedCompany,
		AdditionalTradingNames:     []string{"SHM"},
		IsRegisteredCompany:        &isRegistered,
		DateOfIncorporation:        &DateOfIncorporation{Day: 1, Month: 6, Year: 2010},
	}

	body, err := json.Marshal(company)
	assert.NoError(t, err)
	s := string(body)
	assert.Contains(t, s, `"additional_trading_names":["SHM"]`)
	assert.Contains(t, s, `"is_registered_company":true`)
	assert.Contains(t, s, `"business_type":"limited_company"`)
	assert.Contains(t, s, `"date_of_incorporation":{"day":1,"month":6,"year":2010}`)
}

func TestRepresentativeV3Fields_Roundtrip(t *testing.T) {
	representative := Representative{
		Id:                  "rep_00000000000000000000000000",
		OwnershipPercentage: 100,
		CompanyPosition:     companyPositionPtr(CEOCPStringType),
		Roles:               []EntityRoles{UboERStringType, AuthorisedSignatoryERStringType, DirectorERStringType, ControlPersonERStringType},
		Individual: &Individual{
			FirstName:        "John",
			LastName:         "Doe",
			NationalIdType:   Ssn,
			NationalIdNumber: "AB123456C",
			EmailAddress:     "john@example.com",
			Citizenships:     []Citizenship{{Type: "citizenship", Country: common.Country("US")}},
		},
	}

	body, err := json.Marshal(representative)
	assert.NoError(t, err)
	s := string(body)
	assert.Contains(t, s, `"individual":{`)
	assert.Contains(t, s, `"national_id_type":"ssn"`)
	assert.Contains(t, s, `"citizenships":[{"type":"citizenship","country":"US"}]`)
	assert.Contains(t, s, `"company_position":"ceo"`)
	assert.Contains(t, s, `"ownership_percentage":100`)
	assert.Contains(t, s, `"roles":["ubo","authorised_signatory","director","control_person"]`)
}

func TestFinancialStatementsDocument_Roundtrip(t *testing.T) {
	documents := OnboardSubEntityDocuments{
		FinancialStatements: &FinancialStatements{
			Type:  FinancialStatementsFSStringType,
			Front: "file_00000000000000000000000000",
		},
	}

	body, err := json.Marshal(documents)
	assert.NoError(t, err)
	s := string(body)
	assert.Contains(t, s, `"financial_statements":{`)
	assert.Contains(t, s, `"type":"financial_statements"`)
	assert.Contains(t, s, `"front":"file_00000000000000000000000000"`)
}

func TestOnboardEntityRequestV3_Roundtrip(t *testing.T) {
	request := OnboardEntityRequest{
		Reference:      "ref_1",
		SellerCategory: "saas",
		AgreedTerms: &AgreedTerms{
			Date:    "2026-07-20T10:00:00Z",
			Version: "1.0",
		},
	}

	body, err := json.Marshal(request)
	assert.NoError(t, err)
	s := string(body)
	assert.Contains(t, s, `"seller_category":"saas"`)
	assert.Contains(t, s, `"agreed_terms":{`)
}

func TestEntityRoles_Values(t *testing.T) {
	assert.Equal(t, EntityRoles("ubo"), UboERStringType)
	assert.Equal(t, EntityRoles("legal_representative"), LegalRepresentativeERStringType)
	assert.Equal(t, EntityRoles("authorised_signatory"), AuthorisedSignatoryERStringType)
	assert.Equal(t, EntityRoles("director"), DirectorERStringType)
	assert.Equal(t, EntityRoles("control_person"), ControlPersonERStringType)
}

func TestNationalIdType_Values(t *testing.T) {
	assert.Equal(t, NationalIdType("ssn"), Ssn)
	assert.Equal(t, NationalIdType("itin"), Itin)
	assert.Equal(t, NationalIdType("passport"), Passport)
	assert.Equal(t, NationalIdType("driving_license"), DrivingLicense)
	assert.Equal(t, NationalIdType("national_id_card"), NationalIdCard)
	assert.Equal(t, NationalIdType("residence_permit"), ResidencePermit)
	assert.Equal(t, NationalIdType("other"), Other)
}

func TestBusinessType_HasAllNineteenValues(t *testing.T) {
	values := []BusinessType{
		IndividualOrSoleProprietorship, GeneralPartnership, LimitedPartnership,
		ScottishLimitedPartnership, PublicLimitedCompany, LimitedCompany,
		LimitedLiabilityCorporation, PrivateCorporation, PubliclyTradedCorporation,
		ProfessionalAssociation, UnincorporatedAssociation, AutoEntrepreneur,
		GovernmentAgency, NonProfitEntity, Trust, ClubOrSociety,
		RegulatedFinancialInstitution, CftcRegisteredEntity, SecRegisteredEntity,
	}
	assert.Len(t, values, 19)
}

func TestCompanyPosition_HasAllElevenValues(t *testing.T) {
	values := []CompanyPositionType{
		CEOCPStringType, CFOCPStringType, COOCPStringType, ManagingMemberCPStringType,
		GeneralPartnerCPStringType, PresidentCPStringType, VicePresidentCPStringType,
		TreasurerCPStringType, OtherSeniorManagementCPStringType,
		OtherExecutiveOfficerCPStringType, OtherNonExecutiveNonSeniorCPStringType,
	}
	assert.Len(t, values, 11)
}

func companyPositionPtr(v CompanyPositionType) *CompanyPositionType {
	return &v
}
