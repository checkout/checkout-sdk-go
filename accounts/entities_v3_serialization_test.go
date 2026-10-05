package accounts

import (
	"encoding/json"
	"testing"
	"time"

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
	isRegistered := false
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
	assert.Contains(t, s, `"is_registered_company":false`)
	assert.Contains(t, s, `"business_type":"limited_company"`)
	assert.Contains(t, s, `"date_of_incorporation":{"day":1,"month":6,"year":2010}`)
}

func TestRepresentativeV3Fields_Roundtrip(t *testing.T) {
	representative := Representative{
		Id:                  "rep_2zhgixxlnq7e3xe433tlrevop5",
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
			Front: "file_qpoca3gpgfwdi6isotonx3kgwx",
		},
	}

	body, err := json.Marshal(documents)
	assert.NoError(t, err)
	s := string(body)
	assert.Contains(t, s, `"financial_statements":{`)
	assert.Contains(t, s, `"type":"financial_statements"`)
	assert.Contains(t, s, `"front":"file_qpoca3gpgfwdi6isotonx3kgwx"`)
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

// Every BusinessType value, marshalled inside Company, against its wire value.
func TestBusinessType_WireValues(t *testing.T) {
	expected := map[BusinessType]string{
		GeneralPartnership:             "general_partnership",
		LimitedPartnership:             "limited_partnership",
		PublicLimitedCompany:           "public_limited_company",
		LimitedCompany:                 "limited_company",
		ProfessionalAssociation:        "professional_association",
		UnincorporatedAssociation:      "unincorporated_association",
		AutoEntrepreneur:               "auto_entrepreneur",
		ScottishLimitedPartnership:     "scottish_limited_partnership",
		PrivateCorporation:             "private_corporation",
		LimitedLiabilityCorporation:    "limited_liability_corporation",
		PubliclyTradedCorporation:      "publicly_traded_corporation",
		RegulatedFinancialInstitution:  "regulated_financial_institution",
		SecRegisteredEntity:            "sec_registered_entity",
		CftcRegisteredEntity:           "cftc_registered_entity",
		IndividualOrSoleProprietorship: "individual_or_sole_proprietorship",
		GovernmentAgency:               "government_agency",
		NonProfitEntity:                "non_profit_entity",
		Trust:                          "trust",
		ClubOrSociety:                  "club_or_society",
	}
	assert.Len(t, expected, 19)
	for value, wire := range expected {
		body, err := json.Marshal(Company{BusinessType: value})
		assert.NoError(t, err)
		assert.JSONEq(t, `{"business_type":"`+wire+`"}`, string(body))

		var roundtrip Company
		assert.NoError(t, json.Unmarshal(body, &roundtrip))
		assert.Equal(t, value, roundtrip.BusinessType)
	}
}

// Every CompanyPositionType value, marshalled inside Representative, against its wire value.
func TestCompanyPosition_WireValues(t *testing.T) {
	expected := map[CompanyPositionType]string{
		CEOCPStringType:                        "ceo",
		CFOCPStringType:                        "cfo",
		COOCPStringType:                        "coo",
		ManagingMemberCPStringType:             "managing_member",
		GeneralPartnerCPStringType:             "general_partner",
		PresidentCPStringType:                  "president",
		VicePresidentCPStringType:              "vice_president",
		TreasurerCPStringType:                  "treasurer",
		OtherSeniorManagementCPStringType:      "other_senior_management",
		OtherExecutiveOfficerCPStringType:      "other_executive_officer",
		OtherNonExecutiveNonSeniorCPStringType: "other_non_executive_non_senior",
	}
	assert.Len(t, expected, 11)
	for value, wire := range expected {
		body, err := json.Marshal(Representative{CompanyPosition: companyPositionPtr(value)})
		assert.NoError(t, err)
		assert.JSONEq(t, `{"company_position":"`+wire+`"}`, string(body))

		var roundtrip Representative
		assert.NoError(t, json.Unmarshal(body, &roundtrip))
		assert.Equal(t, value, *roundtrip.CompanyPosition)
	}
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
	assert.Equal(t, "GB", details.ProcessingDetails.SettlementCountry)
	assert.Equal(t, []string{"GB"}, details.ProcessingDetails.TargetCountries)
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
		{name: "7 fractional digits and offset", value: `"2020-12-01T15:01:01.0000000+00:00"`, expected: "2020-12-01T15:01:01Z"},
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

// The US ISV Seller variants (3.0) require pci_compliance_contact next to primary in
// contact_details.email_addresses.
func TestEntityEmailAddressesPciComplianceContact_Roundtrip(t *testing.T) {
	emailAddresses := EntityEmailAddresses{
		Primary:              "admin@superhero1234.com",
		PciComplianceContact: "pci@superhero1234.com",
	}

	body, err := json.Marshal(emailAddresses)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"primary":"admin@superhero1234.com","pci_compliance_contact":"pci@superhero1234.com"}`, string(body))

	var roundtrip EntityEmailAddresses
	assert.NoError(t, json.Unmarshal(body, &roundtrip))
	assert.Equal(t, emailAddresses, roundtrip)

	contactDetails := ContactDetails{EntityEmailAddresses: &emailAddresses}
	body, err = json.Marshal(contactDetails)
	assert.NoError(t, err)
	assert.JSONEq(t,
		`{"email_addresses":{"primary":"admin@superhero1234.com","pci_compliance_contact":"pci@superhero1234.com"}}`,
		string(body))
}

// The 5 IdentityVerificationType values not covered by the every-field documents test, marshalled
// inside IdentityVerification.
func TestIdentityVerificationType_WireValues(t *testing.T) {
	const file = "file_identityverificationaaaaaa"
	expected := map[IdentityVerificationType]string{
		NationalIdentityCardIVStringType: "national_identity_card",
		DrivingLicenseIVStringType:       "driving_license",
		CitizenCardIVStringType:          "citizen_card",
		ResidencePermitIVStringType:      "residence_permit",
		ElectoralIdIVStringType:          "electoral_id",
	}
	for value, wire := range expected {
		document := IdentityVerification{Type: value, Front: file}
		body, err := json.Marshal(document)
		assert.NoError(t, err)
		assert.JSONEq(t, `{"type":"`+wire+`","front":"`+file+`"}`, string(body))

		var roundtrip IdentityVerification
		assert.NoError(t, json.Unmarshal(body, &roundtrip))
		assert.Equal(t, document, roundtrip)
	}
}

// articles_of_association as a company verification type (US Company (2.0)) and
// memorandum_of_association as an articles of association type.
func TestCompanyVerificationAndArticlesOfAssociationTypes_WireValues(t *testing.T) {
	const file = "file_jo7pnpns6kiuldcufo6fhnhgmm"
	documents := OnboardSubEntityDocuments{
		CompanyVerification:   &CompanyVerification{Type: ArticlesOfAssociationCVStringType, Front: file},
		ArticlesOfAssociation: &ArticlesOfAssociation{Type: MemorandumOfAssociationAOSStringType, Front: file},
	}

	body, err := json.Marshal(documents)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"company_verification":    {"type": "articles_of_association", "front": "`+file+`"},
		"articles_of_association": {"type": "memorandum_of_association", "front": "`+file+`"}
	}`, string(body))

	var roundtrip OnboardSubEntityDocuments
	assert.NoError(t, json.Unmarshal(body, &roundtrip))
	assert.Equal(t, documents, roundtrip)
}

// Every key of the request processing_details, then a full round trip.
func TestProcessingDetailsEveryField_Roundtrip(t *testing.T) {
	details := ProcessingDetails{
		SettlementCountry:           "GB",
		TargetCountries:             []string{"GB", "FR"},
		AnnualProcessingVolume:      1000000,
		AverageTransactionValue:     5000,
		AverageOrderFulfillmentTime: 3,
		HighestTransactionValue:     25000,
		Currency:                    common.GBP,
		Payments: &ProcessingDetailsPayments{
			Ach: &ProcessingDetailsAch{
				AnnualAchVolume:              1000000,
				AverageAchTransactionSize:    5000,
				EstimatedMonthlyCreditVolume: 100000,
				AverageCreditAmount:          2500,
			},
		},
	}

	body, err := json.Marshal(details)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"settlement_country": "GB",
		"target_countries": ["GB", "FR"],
		"annual_processing_volume": 1000000,
		"average_transaction_value": 5000,
		"average_order_fulfillment_time": 3,
		"highest_transaction_value": 25000,
		"currency": "GBP",
		"payments": {"ach": {
			"annual_ach_volume": 1000000,
			"average_ach_transaction_size": 5000,
			"estimated_monthly_credit_volume": 100000,
			"average_credit_amount": 2500
		}}
	}`, string(body))

	var roundtrip ProcessingDetails
	assert.NoError(t, json.Unmarshal(body, &roundtrip))
	assert.Equal(t, details, roundtrip)
}

// contact_details with invitee and a v3.0 phone (country_code and number).
func TestContactDetailsInviteeAndPhone_Roundtrip(t *testing.T) {
	contactDetails := ContactDetails{
		Invitee: &Invitee{Email: "invitee@example.com"},
		Phone:   &Phone{CountryCode: common.GB, Number: "2071234567"},
	}

	body, err := json.Marshal(contactDetails)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"invitee": {"email": "invitee@example.com"},
		"phone": {"country_code": "GB", "number": "2071234567"}
	}`, string(body))

	var roundtrip ContactDetails
	assert.NoError(t, json.Unmarshal(body, &roundtrip))
	assert.Equal(t, contactDetails, roundtrip)
}

// The v2.0 company representative: flat person fields, phone with number only, identification
// (US variants) and roles, plus id.
func TestRepresentativeV2Fields_Roundtrip(t *testing.T) {
	representative := Representative{
		Id:           "rep_2zhgixxlnq7e3xe433tlrevop5",
		FirstName:    "John",
		MiddleName:   "Paul",
		LastName:     "Doe",
		DateOfBirth:  &DateOfBirth{Day: 5, Month: 6, Year: 1995},
		PlaceOfBirth: &PlaceOfBirth{Country: common.FR},
		Phone:        &Phone{Number: "2345678910"},
		Address: &common.Address{
			AddressLine1: "90 Tottenham Court Road", City: "London", Zip: "W1T 4TJ", Country: common.GB},
		Identification: &Identification{NationalIdNumber: "123456789"},
		Roles:          []EntityRoles{UboERStringType, DirectorERStringType},
	}

	body, err := json.Marshal(representative)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"id": "rep_2zhgixxlnq7e3xe433tlrevop5",
		"first_name": "John",
		"middle_name": "Paul",
		"last_name": "Doe",
		"date_of_birth": {"day": 5, "month": 6, "year": 1995},
		"place_of_birth": {"country": "FR"},
		"phone": {"number": "2345678910"},
		"address": {"address_line1": "90 Tottenham Court Road", "city": "London", "zip": "W1T 4TJ", "country": "GB"},
		"identification": {"national_id_number": "123456789"},
		"roles": ["ubo", "director"]
	}`, string(body))

	var roundtrip Representative
	assert.NoError(t, json.Unmarshal(body, &roundtrip))
	assert.Equal(t, representative, roundtrip)
}

// The v2.0 sole trader top-level individual, every field, including financial_details (all 4 keys).
func TestIndividualV2TopLevel_Roundtrip(t *testing.T) {
	individual := Individual{
		FirstName:    "Jane",
		MiddleName:   "Ann",
		LastName:     "Doe",
		TradingName:  "Jane's Goods",
		DateOfBirth:  &DateOfBirth{Day: 1, Month: 2, Year: 1980},
		PlaceOfBirth: &PlaceOfBirth{Country: common.US},
		RegisteredAddress: &common.Address{
			AddressLine1: "123 Main Street", City: "San Francisco", State: "CA", Zip: "94105", Country: common.US},
		Identification: &Identification{NationalIdNumber: "123456789"},
		FinancialDetails: &EntityFinancialDetails{
			AnnualProcessingVolume:  120000,
			AverageTransactionValue: 10000,
			HighestTransactionValue: 50000,
			Currency:                common.USD,
		},
	}

	body, err := json.Marshal(individual)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"first_name": "Jane",
		"middle_name": "Ann",
		"last_name": "Doe",
		"trading_name": "Jane's Goods",
		"date_of_birth": {"day": 1, "month": 2, "year": 1980},
		"place_of_birth": {"country": "US"},
		"registered_address": {"address_line1": "123 Main Street", "city": "San Francisco", "state": "CA", "zip": "94105", "country": "US"},
		"identification": {"national_id_number": "123456789"},
		"financial_details": {
			"annual_processing_volume": 120000,
			"average_transaction_value": 10000,
			"highest_transaction_value": 50000,
			"currency": "USD"
		}
	}`, string(body))

	var roundtrip Individual
	assert.NoError(t, json.Unmarshal(body, &roundtrip))
	assert.Equal(t, individual, roundtrip)
}

// Every non-deprecated Company key on the wire. No single variant accepts all of them together; this
// checks the tags only. is_registered_company uses false, the only value the API accepts.
func TestCompanyEveryField_Roundtrip(t *testing.T) {
	isRegistered := false
	address := &common.Address{
		AddressLine1: "1 Rue de Rivoli", City: "Paris", Zip: "75001", Country: common.FR}
	company := Company{
		BusinessRegistrationNumber: "45228589000013",
		BusinessType:               LimitedCompany,
		LegalName:                  "Super Hero Masks SAS",
		TradingName:                "Super Hero Masks",
		AdditionalTradingNames:     []string{"SHM"},
		IsRegisteredCompany:        &isRegistered,
		PrincipalAddress:           address,
		RegisteredAddress:          address,
		Representatives: []Representative{{
			Roles:      []EntityRoles{UboERStringType},
			Individual: &Individual{FirstName: "Jane", LastName: "Doe"},
		}},
		FinancialDetails: &EntityFinancialDetails{
			AnnualProcessingVolume:  120000,
			AverageTransactionValue: 10000,
			HighestTransactionValue: 50000,
			Currency:                common.EUR,
		},
		DateOfIncorporation:     &DateOfIncorporation{Day: 1, Month: 6, Year: 2010},
		RegulatoryLicenceNumber: "LIC-1234",
	}

	body, err := json.Marshal(company)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"business_registration_number": "45228589000013",
		"business_type": "limited_company",
		"legal_name": "Super Hero Masks SAS",
		"trading_name": "Super Hero Masks",
		"additional_trading_names": ["SHM"],
		"is_registered_company": false,
		"principal_address": {"address_line1": "1 Rue de Rivoli", "city": "Paris", "zip": "75001", "country": "FR"},
		"registered_address": {"address_line1": "1 Rue de Rivoli", "city": "Paris", "zip": "75001", "country": "FR"},
		"representatives": [{"roles": ["ubo"], "individual": {"first_name": "Jane", "last_name": "Doe"}}],
		"financial_details": {
			"annual_processing_volume": 120000,
			"average_transaction_value": 10000,
			"highest_transaction_value": 50000,
			"currency": "EUR"
		},
		"date_of_incorporation": {"day": 1, "month": 6, "year": 2010},
		"regulatory_licence_number": "LIC-1234"
	}`, string(body))

	var roundtrip Company
	assert.NoError(t, json.Unmarshal(body, &roundtrip))
	assert.Equal(t, company, roundtrip)
}

// components.schemas["USISVSellerCompany3-0"].example, unmarshalled and marshalled back unchanged.
func TestUsIsvSellerCompanyExample_Roundtrip(t *testing.T) {
	example := `{
		"reference": "isv-seller-example001",
		"agreed_terms": {
			"date": "2026-07-02T10:30:00.0000000+00:00",
			"ip_address": "8.8.8.8",
			"name": "Toby Arden",
			"email": "toby.arden@example.com",
			"version": "cko-platform-terms-1.0.0"
		},
		"seller_category": "cat_retail_001",
		"processing_details": {
			"annual_processing_volume": 1000,
			"average_transaction_value": 2000,
			"average_order_fulfillment_time": 3,
			"target_countries": ["US"],
			"currency": "USD",
			"payments": {
				"ach": {
					"annual_ach_volume": 100000,
					"average_ach_transaction_size": 5000,
					"estimated_monthly_credit_volume": 50000,
					"average_credit_amount": 2500
				}
			}
		},
		"contact_details": {
			"phone": {"number": "4155678900", "country_code": "US"},
			"email_addresses": {"primary": "toby.arden@example.com", "pci_compliance_contact": "pci.contact@example.com"}
		},
		"profile": {
			"urls": ["https://www.isv-seller-example.com"],
			"mccs": ["5551"],
			"holding_currencies": ["USD"],
			"default_holding_currency": "USD"
		},
		"company": {
			"business_registration_number": "12-3456789",
			"business_type": "private_corporation",
			"legal_name": "ISV Seller Example Inc",
			"trading_name": "ISV Seller Example",
			"registered_address": {"address_line1": "123 Main Street", "city": "San Francisco", "state": "CA", "zip": "94105", "country": "US"},
			"principal_address": {"address_line1": "123 Main Street", "city": "San Francisco", "state": "CA", "zip": "94105", "country": "US"},
			"date_of_incorporation": {"year": 2025, "month": 10, "day": 1},
			"representatives": [
				{
					"roles": ["ubo", "control_person"],
					"ownership_percentage": 25,
					"company_position": "ceo",
					"individual": {
						"first_name": "Toby",
						"last_name": "Arden",
						"email_address": "toby.arden@example.com",
						"national_id_type": "ssn",
						"national_id_number": "123456789",
						"date_of_birth": {"day": 15, "month": 1, "year": 1990},
						"place_of_birth": {"country": "US"},
						"citizenships": [{"country": "US"}],
						"phone": {"country_code": "US", "number": "4155678901"},
						"address": {"address_line1": "123 Main Street", "city": "San Francisco", "state": "CA", "zip": "94105", "country": "US"}
					}
				},
				{
					"roles": ["authorised_signatory"],
					"individual": {
						"first_name": "Alex",
						"last_name": "Morgan",
						"email_address": "alex.morgan@example.com",
						"national_id_type": "ssn",
						"national_id_number": "987654321",
						"date_of_birth": {"day": 22, "month": 6, "year": 1985},
						"place_of_birth": {"country": "US"},
						"citizenships": [{"country": "US"}],
						"phone": {"country_code": "US", "number": "4155678902"},
						"address": {"address_line1": "123 Main Street", "city": "San Francisco", "state": "CA", "zip": "94105", "country": "US"}
					}
				}
			]
		}
	}`

	var request OnboardEntityRequest
	assert.NoError(t, json.Unmarshal([]byte(example), &request))
	body, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.JSONEq(t, example, string(body))
}

// components.schemas["USISVSellerSoleTrader3-0"].example, unmarshalled and marshalled back unchanged.
// is_registered_company false must be emitted, not dropped as a zero value.
func TestUsIsvSellerSoleTraderExample_Roundtrip(t *testing.T) {
	example := `{
		"reference": "isv-sole-trader-example001",
		"agreed_terms": {
			"date": "2026-07-02T10:30:00.0000000+00:00",
			"ip_address": "8.8.8.8",
			"name": "Hannah Bret",
			"email": "hannah.bret@example.com",
			"version": "cko-platform-terms-1.0.0"
		},
		"seller_category": "cat_retail_001",
		"processing_details": {
			"annual_processing_volume": 1000,
			"average_transaction_value": 2000,
			"average_order_fulfillment_time": 3,
			"target_countries": ["US"],
			"currency": "USD",
			"payments": {
				"ach": {
					"annual_ach_volume": 100000,
					"average_ach_transaction_size": 5000,
					"estimated_monthly_credit_volume": 50000,
					"average_credit_amount": 2500
				}
			}
		},
		"contact_details": {
			"phone": {"number": "4155678900", "country_code": "US"},
			"email_addresses": {"primary": "hannah.bret@example.com", "pci_compliance_contact": "pci.contact@example.com"}
		},
		"profile": {
			"urls": ["https://www.isv-sole-trader-example.com"],
			"mccs": ["5551"],
			"holding_currencies": ["USD"],
			"default_holding_currency": "USD"
		},
		"company": {
			"business_type": "individual_or_sole_proprietorship",
			"is_registered_company": false,
			"trading_name": "Hannah's Goods",
			"date_of_incorporation": {"year": 2025, "month": 10, "day": 1},
			"principal_address": {"address_line1": "123 Main Street", "city": "San Francisco", "state": "CA", "zip": "94105", "country": "US"},
			"representatives": [
				{
					"roles": ["ubo"],
					"ownership_percentage": 100,
					"individual": {
						"first_name": "Hannah",
						"last_name": "Bret",
						"email_address": "hannah.bret@example.com",
						"national_id_type": "ssn",
						"national_id_number": "123456789",
						"date_of_birth": {"day": 15, "month": 1, "year": 1990},
						"place_of_birth": {"country": "US"},
						"citizenships": [{"country": "US"}],
						"phone": {"country_code": "US", "number": "4155678901"},
						"address": {"address_line1": "123 Main Street", "city": "San Francisco", "state": "CA", "zip": "94105", "country": "US"}
					}
				}
			]
		}
	}`

	var request OnboardEntityRequest
	assert.NoError(t, json.Unmarshal([]byte(example), &request))
	assert.NotNil(t, request.Company.IsRegisteredCompany)
	assert.False(t, *request.Company.IsRegisteredCompany)

	body, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.Contains(t, string(body), `"is_registered_company":false`)
	assert.JSONEq(t, example, string(body))
}

// The PlatformsFileUploadResponse property examples. http_metadata is not on the wire, so fields are
// asserted rather than the whole JSON.
func TestUploadFileResponseSpecExample_Deserialization(t *testing.T) {
	const upload = "https://s3.eu-west-1.amazonaws.com/mp-files-api-staging-prod/ent_ociwguf5a5fe3ndmpnvpnwsi3e/file_6lbss42ezvoufcb2beo76rvwly?AWSAccessKeyId=ASIX4BFJOBCQFLAMPKU3&Expires=1661355993&x-amz-security-token=some_token"
	const self = "https://files.checkout.com/files/file_6lbss42ezvoufcb2beo76rvwly"

	var response UploadFileResponse
	assert.NoError(t, json.Unmarshal([]byte(`{
		"id": "file_6lbss42ezvoufcb2beo76rvwly",
		"maximum_size_in_bytes": 4194304,
		"document_types_for_purpose": ["image/jpeg", "image/png", "image/jpg"],
		"_links": {
			"upload": {"href": "`+upload+`"},
			"self": {"href": "`+self+`"}
		}
	}`), &response))

	assert.Equal(t, "file_6lbss42ezvoufcb2beo76rvwly", response.Id)
	assert.Equal(t, int64(4194304), response.MaximumSizeInBytes)
	assert.Equal(t, []string{"image/jpeg", "image/png", "image/jpg"}, response.DocumentTypesForPurpose)
	if assert.NotNil(t, response.Links["upload"].HRef) {
		assert.Equal(t, upload, *response.Links["upload"].HRef)
	}
	if assert.NotNil(t, response.Links["self"].HRef) {
		assert.Equal(t, self, *response.Links["self"].HRef)
	}
}

// The PlatformsFileRetrieveResponse property examples. uploaded_on is compared as an instant, since
// it is re-serialized in normalised form.
func TestFileDetailsResponseSpecExample_Deserialization(t *testing.T) {
	const download = "https://s3.eu-west-1.amazonaws.com/mp-files-api-clean-prod/ent_ociwguf5a5fe3ndmpnvpnwsi3e/file_6lbss42ezvoufcb2beo76rvwly?X-Amz-Expires=3600&x-amz-security-token=some_token"
	const self = "https://files.checkout.com/files/file_6lbss42ezvoufcb2beo76rvwly"

	var response FileDetailsResponse
	assert.NoError(t, json.Unmarshal([]byte(`{
		"id": "file_6lbss42ezvoufcb2beo76rvwly",
		"status": "invalid",
		"status_reasons": ["InvalidMimeType"],
		"size": 1024,
		"mime_type": "application/pdf",
		"uploaded_on": "2020-12-01T15:01:01.0000000+00:00",
		"purpose": "identity_verification",
		"_links": {
			"download": {"href": "`+download+`"},
			"self": {"href": "`+self+`"}
		}
	}`), &response))

	assert.Equal(t, "file_6lbss42ezvoufcb2beo76rvwly", response.Id)
	assert.Equal(t, "invalid", response.Status)
	assert.Equal(t, []string{"InvalidMimeType"}, response.StatusReasons)
	assert.Equal(t, int64(1024), response.Size)
	assert.Equal(t, "application/pdf", response.MimeType)
	assert.Equal(t, string(common.IdentityVerification), response.Purpose)
	if assert.NotNil(t, response.UploadedOn) {
		assert.True(t, response.UploadedOn.Equal(time.Date(2020, 12, 1, 15, 1, 1, 0, time.UTC)))
	}
	if assert.NotNil(t, response.Links["download"].HRef) {
		assert.Equal(t, download, *response.Links["download"].HRef)
	}
	if assert.NotNil(t, response.Links["self"].HRef) {
		assert.Equal(t, self, *response.Links["self"].HRef)
	}
}
