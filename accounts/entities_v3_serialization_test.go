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

// Regression: EEA Sole Trader (3.0) needs proof_of_residential_address and proof_of_registration on
// the representative, with bank_verification alone at the top level.
func TestEeaSoleTraderRepresentativeDocuments_Serialization(t *testing.T) {
	request := OnboardEntityRequest{
		Reference: "ref_sole_trader",
		Company: &Company{
			BusinessType: IndividualOrSoleProprietorship,
			Representatives: []Representative{{
				Individual: &Individual{FirstName: "Jane", LastName: "Doe"},
				Roles:      []EntityRoles{UboERStringType},
				Documents: &OnboardSubEntityDocuments{
					IdentityVerification: &IdentityVerification{
						Type: PassportIVStringType, Front: "file_identityverificationaaaaaa"},
					ProofOfResidentialAddress: &ProofOfResidentialAddress{
						Type: ProofOfAddressPORAStringType, Front: "file_proofofresidentialaddressa"},
					ProofOfRegistration: &ProofOfRegistration{
						Type: ExtractFromTradeRegisterPORStringType, Front: "file_proofofregistrationaaaaaaa"},
				},
			}},
		},
		Documents: &OnboardSubEntityDocuments{
			BankVerification: &BankVerification{Type: BankStatementBVStringType, Front: "file_bankverificationaaaaaaaaaa"},
		},
	}

	body, err := json.Marshal(request)
	assert.Nil(t, err)

	var decoded struct {
		Company struct {
			Representatives []struct {
				Documents json.RawMessage `json:"documents"`
			} `json:"representatives"`
		} `json:"company"`
		Documents json.RawMessage `json:"documents"`
	}
	assert.Nil(t, json.Unmarshal(body, &decoded))
	assert.JSONEq(t, `{
		"identity_verification":        {"type": "passport",                    "front": "file_identityverificationaaaaaa"},
		"proof_of_residential_address": {"type": "proof_of_address",            "front": "file_proofofresidentialaddressa"},
		"proof_of_registration":        {"type": "extract_from_trade_register", "front": "file_proofofregistrationaaaaaaa"}
	}`, string(decoded.Company.Representatives[0].Documents))
	assert.JSONEq(t, `{"bank_verification": {"type": "bank_statement", "front": "file_bankverificationaaaaaaaaaa"}}`,
		string(decoded.Documents))
	// Key-level check on the raw body, so a tag change cannot pass silently.
	assert.Contains(t, string(body), `"proof_of_residential_address":{`)
	assert.Contains(t, string(body), `"proof_of_registration":{`)
}

// Every field of OnboardSubEntityDocuments, so a tag change on any key cannot pass silently, then a
// full round trip.
func TestOnboardSubEntityDocumentsEveryField_Roundtrip(t *testing.T) {
	const file = "file_aaaaaaaaaaaaaaaaaaaaaaaaaa"
	documents := OnboardSubEntityDocuments{
		IdentityVerification:         &IdentityVerification{Type: PassportIVStringType, Front: file, Back: file},
		CompanyVerification:          &CompanyVerification{Type: IncorporationDocumentCVStringType, Front: file},
		TaxVerification:              &TaxVerification{Type: EinLetterTVStringType, Front: file},
		ArticlesOfAssociation:        &ArticlesOfAssociation{Type: ArticlesOfAssociationAOSStringType, Front: file},
		ShareholderStructure:         &ShareholderStructure{Type: CertifiedShareholderStructureSHSStringType, Front: file},
		BankVerification:             &BankVerification{Type: BankStatementBVStringType, Front: file},
		ProofOfLegality:              &ProofOfLegality{Type: ProofOfLegalityPOLStringType, Front: file},
		ProofOfPrincipalAddress:      &ProofOfPrincipalAddress{Type: ProofOfAddressPOPAStringType, Front: file},
		AdditionalDocument1:          &AdditionalDocument{Front: file},
		AdditionalDocument2:          &AdditionalDocument{Front: file},
		AdditionalDocument3:          &AdditionalDocument{Front: file},
		CertifiedAuthorisedSignatory: &CertifiedAuthorisedSignatory{Type: PowerOfAttorneyCASStringType, Front: file},
		ProofOfResidentialAddress:    &ProofOfResidentialAddress{Type: ProofOfAddressPORAStringType, Front: file},
		ProofOfRegistration:          &ProofOfRegistration{Type: ExtractFromTradeRegisterPORStringType, Front: file},
		FinancialVerification:        &FinancialVerification{Type: FinancialStatementFVStringType, Front: file},
		FinancialStatements:          &FinancialStatements{Type: FinancialStatementsFSStringType, Front: file},
	}

	body, err := json.Marshal(documents)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
		"identity_verification":          {"type": "passport", "front": "`+file+`", "back": "`+file+`"},
		"company_verification":           {"type": "incorporation_document", "front": "`+file+`"},
		"tax_verification":               {"type": "ein_letter", "front": "`+file+`"},
		"articles_of_association":        {"type": "articles_of_association", "front": "`+file+`"},
		"shareholder_structure":          {"type": "certified_shareholder_structure", "front": "`+file+`"},
		"bank_verification":              {"type": "bank_statement", "front": "`+file+`"},
		"proof_of_legality":              {"type": "proof_of_legality", "front": "`+file+`"},
		"proof_of_principal_address":     {"type": "proof_of_address", "front": "`+file+`"},
		"additional_document1":           {"front": "`+file+`"},
		"additional_document2":           {"front": "`+file+`"},
		"additional_document3":           {"front": "`+file+`"},
		"certified_authorised_signatory": {"type": "power_of_attorney", "front": "`+file+`"},
		"proof_of_residential_address":   {"type": "proof_of_address", "front": "`+file+`"},
		"proof_of_registration":          {"type": "extract_from_trade_register", "front": "`+file+`"},
		"financial_verification":         {"type": "financial_statement", "front": "`+file+`"},
		"financial_statements":           {"type": "financial_statements", "front": "`+file+`"}
	}`, string(body))

	var roundTripped OnboardSubEntityDocuments
	assert.Nil(t, json.Unmarshal(body, &roundTripped))
	assert.Equal(t, documents, roundTripped)
}

func TestProofOfRegistrationOtherType_Deserialization(t *testing.T) {
	var documents OnboardSubEntityDocuments
	assert.Nil(t, json.Unmarshal(
		[]byte(`{"proof_of_registration": {"type": "other", "front": "file_proofofregistrationaaaaaaa"}}`), &documents))
	assert.Equal(t, OtherPORStringType, documents.ProofOfRegistration.Type)
}

// EEA and GB Company Full (3.0) allow a representative that is a company:
// { company: { legal_name, trading_name, registered_address }, ownership_percentage }.
func TestControllingCompanyRepresentative_Serialization(t *testing.T) {
	representative := Representative{
		OwnershipPercentage: 60,
		Company: &Company{
			LegalName:   "Parent Holdings Ltd",
			TradingName: "Parent Holdings",
			RegisteredAddress: &common.Address{
				AddressLine1: "1 Main Street", City: "London", Zip: "W1T 4TJ", Country: common.GB},
		},
	}

	body, err := json.Marshal(representative)
	assert.Nil(t, err)
	assert.JSONEq(t, `{
		"ownership_percentage": 60,
		"company": {
			"legal_name": "Parent Holdings Ltd",
			"trading_name": "Parent Holdings",
			"registered_address": {"address_line1": "1 Main Street", "city": "London", "zip": "W1T 4TJ", "country": "GB"}
		}
	}`, string(body))
}

// The v2.0 US Company representative identification was not modelled.
func TestV2UsRepresentativeIdentification_Serialization(t *testing.T) {
	representative := Representative{
		FirstName:      "John",
		LastName:       "Doe",
		Identification: &Identification{NationalIdNumber: "123456789"},
	}

	body, err := json.Marshal(representative)
	assert.Nil(t, err)
	assert.JSONEq(t, `{"first_name": "John", "last_name": "Doe", "identification": {"national_id_number": "123456789"}}`,
		string(body))
}

// GET /accounts/entities/{id} returns documents and processing_details; neither was modelled. The
// processing amounts are integers in minor units with no maximum, so they are int64.
func TestOnboardEntityDetailsDocumentsAndProcessingDetails_Deserialization(t *testing.T) {
	var details OnboardEntityDetails
	assert.Nil(t, json.Unmarshal([]byte(`{
		"id": "ent_aaaaaaaaaaaaaaaaaaaaaaaaaa",
		"processing_details": {
			"settlement_country": "GB", "target_countries": ["GB"], "currency": "USD",
			"annual_processing_volume": 3000000000,
			"average_transaction_value": 2500000000,
			"highest_transaction_value": 9000000000
		},
		"documents": {"bank_verification": {"type": "bank_statement", "front": "file_bankverificationaaaaaaaaaa"}},
		"company": {"representatives": [{"documents": {"proof_of_registration":
			{"type": "extract_from_trade_register", "front": "file_proofofregistrationaaaaaaa"}}}]}
	}`), &details))

	assert.Equal(t, int64(3000000000), details.ProcessingDetails.AnnualProcessingVolume)
	assert.Equal(t, int64(2500000000), details.ProcessingDetails.AverageTransactionValue)
	assert.Equal(t, int64(9000000000), details.ProcessingDetails.HighestTransactionValue)
	assert.Equal(t, common.USD, details.ProcessingDetails.Currency)
	assert.Equal(t, BankStatementBVStringType, details.Documents.BankVerification.Type)
	assert.Equal(t, ExtractFromTradeRegisterPORStringType,
		details.Company.Representatives[0].Documents.ProofOfRegistration.Type)
}

// uploaded_on is read as an RFC 3339 date-time, and left nil when the API omits it.
func TestFileDetailsResponseUploadedOn_Deserialization(t *testing.T) {
	cases := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "RFC 3339", value: `"2026-10-01T12:30:00Z"`, expected: "2026-10-01T12:30:00Z"},
		{name: "absent"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			body := `{"id": "file_aaaaaaaaaaaaaaaaaaaaaaaaaa", "status": "created", "purpose": "proof_of_registration"}`
			if tc.value != "" {
				body = `{"id": "file_aaaaaaaaaaaaaaaaaaaaaaaaaa", "status": "created", "purpose": "proof_of_registration", "uploaded_on": ` + tc.value + `}`
			}
			var response FileDetailsResponse
			assert.Nil(t, json.Unmarshal([]byte(body), &response))
			assert.Equal(t, "file_aaaaaaaaaaaaaaaaaaaaaaaaaa", response.Id)
			assert.Equal(t, "created", response.Status)
			assert.Equal(t, "proof_of_registration", response.Purpose)
			if tc.expected == "" {
				assert.Nil(t, response.UploadedOn)
			} else {
				assert.Equal(t, tc.expected, response.UploadedOn.Format("2006-01-02T15:04:05.999999999Z07:00"))
			}
		})
	}
}

// The upload endpoints send the purpose's value on the wire; all values are asserted as strings.
func TestPurpose_WireValues(t *testing.T) {
	expected := map[common.Purpose]string{
		common.DisputesEvidence:             "dispute_evidence",
		common.AdditionalDocument:           "additional_document",
		common.ArticlesOfAssociation:        "articles_of_association",
		common.BankVerification:             "bank_verification",
		common.CertifiedAuthorisedSignatory: "certified_authorised_signatory",
		common.CompanyOwnership:             "company_ownership",
		common.CompanyVerification:          "company_verification",
		common.FinancialVerification:        "financial_verification",
		common.Identification:               "identification",
		common.IdentityVerification:         "identity_verification",
		common.TaxVerification:              "tax_verification",
		common.ProofOfLegality:              "proof_of_legality",
		common.ProofOfPrincipalAddress:      "proof_of_principal_address",
		common.ShareholderStructure:         "shareholder_structure",
		common.ProofOfResidentialAddress:    "proof_of_residential_address",
		common.ProofOfRegistration:          "proof_of_registration",
	}
	for purpose, wire := range expected {
		assert.Equal(t, wire, string(purpose))
	}
}
