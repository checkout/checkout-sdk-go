package identities

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

func TestDocumentType_MatchesTheSwaggerEnumExactly(t *testing.T) {
	assert.Equal(t, DocumentType("Driving licence"), DrivingLicence)
	assert.Equal(t, DocumentType("ID"), IdCard)
	assert.Equal(t, DocumentType("Other"), Other)
	assert.Equal(t, DocumentType("Passport"), Passport)
	assert.Equal(t, DocumentType("Residence Permit"), ResidencePermit)
	assert.Equal(t, DocumentType("Travel Document"), TravelDocument)
	assert.Equal(t, DocumentType("Visa"), Visa)
}

// M3. Not to be confused with AdvAttemptTerminated and IddvAttemptTerminated, which are a
// different enum and already carried terminated before this change.
func TestAttemptVerificationStatus_GainedTerminated(t *testing.T) {
	assert.Equal(t, AttemptVerificationStatus("terminated"), AttemptTerminated)
}

// M4. Again distinct from AdvCreated and IddvCreated, which are different enums.
func TestIdentityVerificationStatus_GainedCreated(t *testing.T) {
	assert.Equal(t, IdentityVerificationStatus("created"), IdvCreated)
}

func TestRiskLabel_MatchesTheSwaggerEnum(t *testing.T) {
	assert.Equal(t, RiskLabel("multiple_faces_detected"), MultipleFacesDetected)
	assert.Equal(t, RiskLabel("mcc_not_confident"), MccNotConfident)
	assert.Equal(t, RiskLabel("risky_document_format"), RiskyDocumentFormat)
}

func TestInitialDevice_MatchesTheSwaggerEnum(t *testing.T) {
	assert.Equal(t, InitialDevice("desktop"), Desktop)
	assert.Equal(t, InitialDevice("mobile"), Mobile)
}

func TestCertificationEnums_MatchTheSwagger(t *testing.T) {
	assert.Equal(t, CertificationType("diatf"), Diatf)
	assert.Equal(t, Gpg45Profile("M1A"), Gpg45M1A)
	assert.Equal(t, Gpg45Profile("M1C"), Gpg45M1C)
	assert.Equal(t, Gpg45Profile("H1A"), Gpg45H1A)
	assert.Equal(t, LevelOfConfidence("medium"), LevelOfConfidenceMedium)
	assert.Equal(t, LevelOfConfidence("high"), LevelOfConfidenceHigh)
}

// ---------------------------------------------------------------------------------------------
// Part E: the new entities
// ---------------------------------------------------------------------------------------------

func TestPhoneNumber_SerializesBothProperties(t *testing.T) {
	marshalled, err := json.Marshal(PhoneNumber{CountryCode: "+33", Number: "5555550102"})
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, map[string]interface{}{
		"country_code": "+33",
		"number":       "5555550102",
	}, decoded)
}

// T17: the phone country code is a dialling prefix, not an ISO code, so it stays a string.
func TestPhoneNumber_CountryCodeIsADiallingPrefix(t *testing.T) {
	marshalled, err := json.Marshal(PhoneNumber{CountryCode: "+33", Number: "5555550102"})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"country_code":"+33"`)
}

func TestIdvAddress_SerializesEveryProperty(t *testing.T) {
	address := IdvAddress{
		AddressLine1: "123 Main Street",
		AddressLine2: "Apt 4B",
		City:         "London",
		State:        "Greater London",
		Zip:          "SW1A 1AA",
		Country:      common.GB,
	}

	marshalled, err := json.Marshal(address)
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, map[string]interface{}{
		"address_line1": "123 Main Street",
		"address_line2": "Apt 4B",
		"city":          "London",
		"state":         "Greater London",
		"zip":           "SW1A 1AA",
		"country":       "GB",
	}, decoded)
}

func TestDeclaredData_CarriesOnlyTheSharedShape(t *testing.T) {
	marshalled, err := json.Marshal(DeclaredData{Name: "Hannah Bret", BirthDate: "1994-10-15"})
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, map[string]interface{}{
		"name":       "Hannah Bret",
		"birth_date": "1994-10-15",
	}, decoded)
}

// D1: IdentityDeclaredData embeds DeclaredData, so the embedded fields flatten into the same
// JSON object rather than nesting.
func TestIdentityDeclaredData_FlattensTheEmbeddedShape(t *testing.T) {
	declaredData := IdentityDeclaredData{
		DeclaredData: DeclaredData{Name: "Hannah Bret", BirthDate: "1994-10-15"},
		Email:        "hannah.bret@example.com",
		PhoneNumber:  &PhoneNumber{CountryCode: "+33", Number: "5555550102"},
		Address:      &IdvAddress{AddressLine1: "123 Main Street", Country: common.GB},
	}

	marshalled, err := json.Marshal(declaredData)
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, "Hannah Bret", decoded["name"])
	assert.Equal(t, "1994-10-15", decoded["birth_date"])
	assert.Equal(t, "hannah.bret@example.com", decoded["email"])
	assert.NotContains(t, decoded, "DeclaredData")
	assert.Equal(t, map[string]interface{}{"country_code": "+33", "number": "5555550102"},
		decoded["phone_number"])
}

func TestIdentityDeclaredData_RoundTrip(t *testing.T) {
	original := IdentityDeclaredData{
		DeclaredData: DeclaredData{Name: "Hannah Bret", BirthDate: "1994-10-15"},
		Email:        "hannah.bret@example.com",
		PhoneNumber:  &PhoneNumber{CountryCode: "+33", Number: "5555550102"},
		Address:      &IdvAddress{City: "London", Country: common.GB},
	}

	marshalled, err := json.Marshal(original)
	assert.NoError(t, err)

	var decoded IdentityDeclaredData
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, original.Name, decoded.Name)
	assert.Equal(t, original.BirthDate, decoded.BirthDate)
	assert.Equal(t, original.Email, decoded.Email)
	assert.Equal(t, original.PhoneNumber.CountryCode, decoded.PhoneNumber.CountryCode)
	assert.Equal(t, common.GB, decoded.Address.Country)
}

// D1: the face authentication attempt schema declares neither document field, so sending them
// would be a request the API rejects. The narrow type is what prevents it.
func TestClientInformation_KeepsTheTwoFieldFavShape(t *testing.T) {
	marshalled, err := json.Marshal(ClientInformation{
		PreSelectedResidenceCountry: common.FR,
		PreSelectedLanguage:         "en-US",
	})
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, map[string]interface{}{
		"pre_selected_residence_country": "FR",
		"pre_selected_language":          "en-US",
	}, decoded)
	assert.NotContains(t, decoded, "pre_selected_document_type")
	assert.NotContains(t, decoded, "pre_selected_document_issuing_country")
}

func TestIdentityVerificationClientInformation_AddsTheTwoIdvOnlyFields(t *testing.T) {
	clientInformation := IdentityVerificationClientInformation{
		ClientInformation: ClientInformation{
			PreSelectedResidenceCountry: common.FR,
			PreSelectedLanguage:         "en-US",
		},
		PreSelectedDocumentIssuingCountry: common.GB,
		PreSelectedDocumentType:           TravelDocument,
	}

	marshalled, err := json.Marshal(clientInformation)
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, map[string]interface{}{
		"pre_selected_residence_country":        "FR",
		"pre_selected_language":                 "en-US",
		"pre_selected_document_issuing_country": "GB",
		"pre_selected_document_type":            "Travel Document",
	}, decoded)
}

// Part E: certifications nest two deep.
func TestCertification_DeserializesNestedDiatfData(t *testing.T) {
	payload := `[{"type":"diatf","data":{"gpg45_profile":"M1A","level_of_confidence":"high","right_to_work":"GRANTED"}}]`

	var certifications []Certification
	assert.NoError(t, json.Unmarshal([]byte(payload), &certifications))

	assert.Len(t, certifications, 1)
	assert.Equal(t, Diatf, certifications[0].Type)
	assert.NotNil(t, certifications[0].Data)
	assert.Equal(t, Gpg45M1A, certifications[0].Data.Gpg45Profile)
	assert.Equal(t, LevelOfConfidenceHigh, certifications[0].Data.LevelOfConfidence)
	assert.Equal(t, "GRANTED", certifications[0].Data.RightToWork)
}

// Part E: the three new session information fields.
func TestApplicantSessionInformation_DeserializesTheNewFields(t *testing.T) {
	payload := `{
		"ip_address":"1.2.3.4",
		"number_of_sessions":3,
		"user_agent":"Mozilla/5.0",
		"initial_device":"mobile",
		"selected_documents":[{"country":"GB","document_type":"Passport"}]
	}`

	var session ApplicantSessionInformation
	assert.NoError(t, json.Unmarshal([]byte(payload), &session))

	assert.Equal(t, "1.2.3.4", session.IpAddress)
	assert.Equal(t, 3, session.NumberOfSessions)
	assert.Equal(t, "Mozilla/5.0", session.UserAgent)
	assert.Equal(t, Mobile, session.InitialDevice)
	assert.Len(t, session.SelectedDocuments, 1)
	assert.Equal(t, common.GB, session.SelectedDocuments[0].Country)
}

// M5 and D3 on DocumentDetails: the five new fields, plus the now-typed issuing country.
func TestDocumentDetails_PermitFieldsAndTypedIssuingCountry(t *testing.T) {
	payload := `{
		"document_type":"Residence Permit",
		"document_issuing_country":"FR",
		"nationality":"GB",
		"address":"123 Main Street, London, SW1A 1AA",
		"permit_obtaining_date":"2020-01-15",
		"permit_expiry_date":"2030-01-15",
		"permit_type_detailed":"Long term resident",
		"permit_type_remarks":"Work permitted"
	}`

	var document DocumentDetails
	assert.NoError(t, json.Unmarshal([]byte(payload), &document))

	assert.Equal(t, common.FR, document.DocumentIssuingCountry)
	assert.Equal(t, common.GB, document.Nationality)
	assert.Equal(t, "123 Main Street, London, SW1A 1AA", document.Address)
	assert.Equal(t, "2020-01-15", document.PermitObtainingDate)
	assert.Equal(t, "2030-01-15", document.PermitExpiryDate)
	assert.Equal(t, "Long term resident", document.PermitTypeDetailed)
	assert.Equal(t, "Work permitted", document.PermitTypeRemarks)
}

// The document address is a flat string here, unlike the structured IdvAddress the declared data
// uses. Same word, two shapes, in the same spec.
func TestDocumentDetails_AddressIsAFlatString(t *testing.T) {
	var document DocumentDetails
	assert.NoError(t, json.Unmarshal([]byte(`{"address":"123 Main Street"}`), &document))
	assert.Equal(t, "123 Main Street", document.Address)
}

// ---------------------------------------------------------------------------------------------
// M1: the attempts pagination filter
// ---------------------------------------------------------------------------------------------

func TestAttemptsQueryFilter_IsSeparateFromTheAssetsFilter(t *testing.T) {
	// Same shape, different documented resource: one counts attempts, the other counts assets.
	attempts := AttemptsQueryFilter{Skip: 5, Limit: 25}
	assets := AttemptAssetsQueryFilter{Skip: 5, Limit: 25}

	assert.Equal(t, 5, attempts.Skip)
	assert.Equal(t, 25, attempts.Limit)
	assert.Equal(t, 5, assets.Skip)
	assert.Equal(t, 25, assets.Limit)
}

// Documents a real limitation rather than asserting desired behaviour: url omitempty drops a zero
// Skip, so skip=0 cannot be sent. Harmless because 0 is the API default, but it is not
// expressible. The PHP SDK has the same gap; python does not.
func TestAttemptsQueryFilter_ZeroSkipIsNotExpressible(t *testing.T) {
	query, err := common.BuildQueryPath("attempts", AttemptsQueryFilter{Skip: 0, Limit: 10})
	assert.NoError(t, err)
	assert.Equal(t, "attempts?limit=10", query)
	assert.NotContains(t, query, "skip")
}

func TestAttemptsQueryFilter_BuildsTheQueryString(t *testing.T) {
	query, err := common.BuildQueryPath("attempts", AttemptsQueryFilter{Skip: 5, Limit: 25})
	assert.NoError(t, err)
	assert.Contains(t, query, "skip=5")
	assert.Contains(t, query, "limit=25")
}

func TestAttemptsQueryFilter_EmptyFilterAddsNoQueryString(t *testing.T) {
	query, err := common.BuildQueryPath("attempts", AttemptsQueryFilter{})
	assert.NoError(t, err)
	assert.Equal(t, "attempts", query)
}
