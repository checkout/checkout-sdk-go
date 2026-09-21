package identityverification

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

// Verifies the identity verification attempt assets response (IdvAttemptAssets)
// deserializes from the API wire format, including the asset type enum and the
// nested _links.asset_url HAL link.
func TestIdentityVerificationAttemptAssetsResponse_Unmarshal(t *testing.T) {
	payload := `{
		"total_count":2,
		"skip":0,
		"limit":10,
		"data":[
			{"type":"document_front_image","_links":{"asset_url":{"href":"https://example.com/document-front.jpg"}}},
			{"type":"face_image","_links":{"asset_url":{"href":"https://example.com/face-image.jpg"}}}
		],
		"_links":{"self":{"href":"https://example.com/assets"},"previous":{"href":"https://example.com/assets?skip=0"}}
	}`

	var response IdentityVerificationAttemptAssetsResponse
	err := json.Unmarshal([]byte(payload), &response)

	assert.NoError(t, err)
	assert.Equal(t, 2, response.TotalCount)
	assert.Equal(t, 0, response.Skip)
	assert.Equal(t, 10, response.Limit)
	assert.Len(t, response.Data, 2)
	assert.Equal(t, identities.DocumentFrontImageIdvAsset, response.Data[0].Type)
	assert.NotNil(t, response.Data[0].Links.AssetUrl.HRef)
	assert.Equal(t, "https://example.com/document-front.jpg", *response.Data[0].Links.AssetUrl.HRef)
	assert.Equal(t, identities.FaceImageIdvAsset, response.Data[1].Type)
	assert.Contains(t, response.Links, "self")
}

// The IDV response's face field is tagged `face`, which is what the spec declares. It read
// `face_image` until the 2026-09-02 pass, a key the API never sends, so the field silently stayed
// nil on every response. This is a live data fix, not a rename.
func TestIdentityVerificationResponse_FaceUsesTheSpecTag(t *testing.T) {
	var response IdentityVerificationResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"id":"idv_1","face":{"image_signed_url":"https://example.com/face.png"}}`), &response))

	assert.NotNil(t, response.Face, "face must populate from the `face` key the API actually sends")
	assert.Equal(t, "https://example.com/face.png", response.Face.ImageSignedUrl)
}

// The old tag must no longer populate anything, otherwise the bug could silently return.
func TestIdentityVerificationResponse_IgnoresTheOldFaceImageKey(t *testing.T) {
	var response IdentityVerificationResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"id":"idv_1","face_image":{"image_signed_url":"https://example.com/face.png"}}`), &response))

	assert.Nil(t, response.Face)
}

// The response: risk labels are typed, certifications nest two deep, and the policy
// version is carried.
func TestIdentityVerificationResponse_RiskLabelsCertificationsAndPolicyVersion(t *testing.T) {
	payload := `{
		"id":"idv_1",
		"status":"created",
		"risk_labels":["multiple_faces_detected","mcc_not_confident"],
		"verification_policy_version":"2026-09-01",
		"certifications":[{"type":"diatf","data":{"gpg45_profile":"H1A","level_of_confidence":"high"}}]
	}`

	var response IdentityVerificationResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))

	assert.Equal(t, identities.IdvCreated, response.Status)
	assert.Equal(t, []identities.RiskLabel{identities.MultipleFacesDetected, identities.MccNotConfident},
		response.RiskLabels)
	assert.Equal(t, "2026-09-01", response.VerificationPolicyVersion)
	assert.Len(t, response.Certifications, 1)
	assert.Equal(t, identities.Diatf, response.Certifications[0].Type)
	assert.Equal(t, identities.Gpg45H1A, response.Certifications[0].Data.Gpg45Profile)
}

// The report response carries pdf_report, not the removed signed_url.
func TestIdentityVerificationReportResponse_PdfReport(t *testing.T) {
	var response IdentityVerificationReportResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"pdf_report":"https://www.example.com/reports/idv_1.pdf"}`), &response))
	assert.Equal(t, "https://www.example.com/reports/idv_1.pdf", response.PdfReport)

	var stale IdentityVerificationReportResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"signed_url":"https://www.example.com/reports/idv_1.pdf"}`), &stale))
	assert.Empty(t, stale.PdfReport)
}

// The attempt request: phone_number is sendable, and the client information is
// the wider IDV shape.
func TestCreateIdentityVerificationAttemptRequest_PhoneNumberAndIdvClientInformation(t *testing.T) {
	request := CreateIdentityVerificationAttemptRequest{
		RedirectUrl: "https://example.com/success",
		PhoneNumber: &identities.PhoneNumber{CountryCode: "+33", Number: "5555550102"},
		ClientInformation: &identities.IdentityVerificationClientInformation{
			ClientInformation: identities.ClientInformation{
				PreSelectedResidenceCountry: common.FR,
				PreSelectedLanguage:         "en-US",
			},
			PreSelectedDocumentIssuingCountry: common.GB,
			PreSelectedDocumentType:           identities.TravelDocument,
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"phone_number":{"country_code":"+33","number":"5555550102"}`)
	assert.Contains(t, body, `"pre_selected_document_type":"Travel Document"`)
	assert.Contains(t, body, `"pre_selected_document_issuing_country":"GB"`)
}

// The IDV create requests take the five field IdvIdentityDeclaredData, and all three
// extra fields are genuinely sendable.
func TestCreateIdentityVerificationRequest_CarriesTheIdentityDeclaredData(t *testing.T) {
	request := CreateIdentityVerificationRequest{
		ApplicantId: "aplt_tkoi5db4hryu5cei5vwoabr7we",
		DeclaredData: &identities.IdentityDeclaredData{
			DeclaredData: identities.DeclaredData{Name: "Hannah Bret", BirthDate: "1994-10-15"},
			Email:        "hannah.bret@example.com",
			PhoneNumber:  &identities.PhoneNumber{CountryCode: "+33", Number: "5555550102"},
			Address:      &identities.IdvAddress{City: "London", Country: common.GB},
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"birth_date":"1994-10-15"`)
	assert.Contains(t, body, `"email":"hannah.bret@example.com"`)
	assert.Contains(t, body, `"country":"GB"`)
}

// The attempt response: phone_number, and the session information gained three fields.
func TestIdentityVerificationAttemptResponse_PhoneNumberAndSessionFields(t *testing.T) {
	payload := `{
		"id":"att_1",
		"status":"terminated",
		"phone_number":{"country_code":"+44","number":"7700900000"},
		"applicant_session_information":{
			"ip_address":"1.2.3.4","number_of_sessions":2,"user_agent":"Mozilla/5.0","initial_device":"desktop"
		}
	}`

	var response IdentityVerificationAttemptResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))

	assert.Equal(t, identities.AttemptTerminated, response.Status)
	assert.NotNil(t, response.PhoneNumber)
	assert.Equal(t, "+44", response.PhoneNumber.CountryCode)
	assert.Equal(t, 2, response.ApplicantSessionInformation.NumberOfSessions)
	assert.Equal(t, identities.Desktop, response.ApplicantSessionInformation.InitialDevice)
}

// The IDV responses all declare _links. The attempt carries verification_url, which is how a
// caller sends the applicant to complete the attempt.
func TestIdentityVerificationResponses_CarryTheirLinks(t *testing.T) {
	var verification IdentityVerificationResponse
	assert.NoError(t, json.Unmarshal([]byte(
		`{"id":"idv_1","_links":{"self":{"href":"https://idv.checkout.com/idv_1"},"applicant":{"href":"https://idv.checkout.com/aplt_1"}}}`),
		&verification))
	assert.NotNil(t, verification.Links.Self)
	assert.NotNil(t, verification.Links.Applicant)

	var attempt IdentityVerificationAttemptResponse
	assert.NoError(t, json.Unmarshal([]byte(
		`{"id":"att_1","_links":{"self":{"href":"https://idv.checkout.com/att_1"},"verification_url":{"href":"https://verify.checkout.com/att_1"}}}`),
		&attempt))
	assert.NotNil(t, attempt.Links.Self)
	assert.NotNil(t, attempt.Links.VerificationUrl)
	assert.Contains(t, *attempt.Links.VerificationUrl.HRef, "verify.checkout.com")

	var list IdentityVerificationAttemptsResponse
	assert.NoError(t, json.Unmarshal([]byte(
		`{"total_count":25,"skip":10,"limit":10,"data":[],"_links":{"next":{"href":"https://idv.checkout.com/a?skip=20"},"previous":{"href":"https://idv.checkout.com/a?skip=0"}}}`),
		&list))
	assert.NotNil(t, list.Links.Next)
	assert.NotNil(t, list.Links.Previous)
}

// D4: the attempt inherits declared_data from the shared base, but the attempt schema does not
// declare it, so it never populates. Retained deliberately; this pins the documented behaviour so
// nobody "fixes" it by wiring it to something.
func TestIdentityVerificationAttemptResponse_DeclaredDataNeverPopulates(t *testing.T) {
	var attempt IdentityVerificationAttemptResponse
	assert.NoError(t, json.Unmarshal([]byte(`{"id":"att_1","status":"completed"}`), &attempt))
	assert.Nil(t, attempt.DeclaredData,
		"the attempt schema declares no declared_data, so the inherited field stays nil")
}
