package iddocumentverification

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

// iddvAttemptAssetsSwaggerExample is components.examples.iddv_attempt_assets_response_body from
// shared/swagger-latest.json, verbatim apart from shortened hrefs.
const iddvAttemptAssetsSwaggerExample = `{
	"total_count": 2,
	"skip": 0,
	"limit": 10,
	"data": [
		{
			"type": "document_front_image",
			"_links": {"asset_url": {"href": "https://storage-b.env.ubble.ai/ubble-ai/NDY/document_front.png"}}
		},
		{
			"type": "document_back_image",
			"_links": {"asset_url": {"href": "https://storage-b.env.ubble.ai/ubble-ai/NDY/document_back.png"}}
		}
	],
	"_links": {
		"self": {"href": "https://idv.checkout.com/id-document-verifications/iddv_1/attempts/datp_1/assets"}
	}
}`

// Part A: the new GET .../attempts/{attemptId}/assets response.
func TestIdDocumentVerificationAttemptAssetsResponse_FromSwaggerExample(t *testing.T) {
	var response IdDocumentVerificationAttemptAssetsResponse
	assert.NoError(t, json.Unmarshal([]byte(iddvAttemptAssetsSwaggerExample), &response))

	assert.Equal(t, 2, response.TotalCount)
	assert.Equal(t, 10, response.Limit)
	assert.Len(t, response.Data, 2)
	assert.Equal(t, identities.DocumentFrontImageIddvAsset, response.Data[0].Type)
	assert.Equal(t, identities.DocumentBackImageIddvAsset, response.Data[1].Type)
	assert.NotNil(t, response.Data[0].Links.AssetUrl.HRef)
	assert.Contains(t, *response.Data[0].Links.AssetUrl.HRef, "document_front.png")
	assert.Contains(t, *response.Data[1].Links.AssetUrl.HRef, "document_back.png")
	assert.Contains(t, response.Links, "self")
}

// T7: data declares minItems 0.
func TestIdDocumentVerificationAttemptAssetsResponse_EmptyPage(t *testing.T) {
	payload := `{"total_count":0,"skip":0,"limit":10,"data":[],"_links":{"self":{"href":"https://idv.checkout.com/a"}}}`

	var response IdDocumentVerificationAttemptAssetsResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))

	assert.Equal(t, 0, response.TotalCount)
	assert.Empty(t, response.Data)
}

// The IDDV attempt asset enum has exactly two values, and neither is the ADV one.
func TestIdDocumentVerificationAttemptAssetType_Values(t *testing.T) {
	assert.Equal(t, "document_front_image", string(identities.DocumentFrontImageIddvAsset))
	assert.Equal(t, "document_back_image", string(identities.DocumentBackImageIddvAsset))
}

// Part C: IdvPdf replaced signed_url with pdf_report.
func TestIdDocumentVerificationReportResponse_PdfReport(t *testing.T) {
	var response IdDocumentVerificationReportResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"pdf_report":"https://www.example.com/reports/iddv_1.pdf"}`), &response))

	assert.Equal(t, "https://www.example.com/reports/iddv_1.pdf", response.PdfReport)
}

func TestIdDocumentVerificationReportResponse_IgnoresSignedUrl(t *testing.T) {
	var response IdDocumentVerificationReportResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"signed_url":"https://www.example.com/reports/iddv_1.pdf"}`), &response))

	assert.Empty(t, response.PdfReport)
}

// M5 and D3 on the document the IDDV response returns: five new fields, and the issuing country
// is now typed.
func TestIdDocumentVerificationResponse_DocumentPermitFieldsAndTypedCountry(t *testing.T) {
	payload := `{
		"id": "iddv_1",
		"status": "approved",
		"document": {
			"document_type": "Residence Permit",
			"document_issuing_country": "FR",
			"nationality": "GB",
			"address": "123 Main Street, London, SW1A 1AA",
			"permit_obtaining_date": "2020-01-15",
			"permit_expiry_date": "2030-01-15",
			"permit_type_detailed": "Long term resident",
			"permit_type_remarks": "Work permitted"
		}
	}`

	var response IdDocumentVerificationResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))

	assert.NotNil(t, response.Document)
	assert.Equal(t, identities.ResidencePermit, response.Document.DocumentType)
	assert.Equal(t, common.FR, response.Document.DocumentIssuingCountry)
	assert.Equal(t, common.GB, response.Document.Nationality)
	assert.Equal(t, "123 Main Street, London, SW1A 1AA", response.Document.Address)
	assert.Equal(t, "2020-01-15", response.Document.PermitObtainingDate)
	assert.Equal(t, "2030-01-15", response.Document.PermitExpiryDate)
	assert.Equal(t, "Long term resident", response.Document.PermitTypeDetailed)
	assert.Equal(t, "Work permitted", response.Document.PermitTypeRemarks)
}

// Part E: the IDDV request takes the narrower IdvDeclaredData shape.
func TestCreateIdDocumentVerificationRequest_DeclaredDataBirthDate(t *testing.T) {
	request := CreateIdDocumentVerificationRequest{
		ApplicantId:   "aplt_tkoi5db4hryu5cei5vwoabr7we",
		UserJourneyId: "usj_tkoi5db4hryu5cei5vwoabr7we",
		DeclaredData: &identities.DeclaredData{
			Name:      "Hannah Bret",
			BirthDate: "1994-10-15",
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"birth_date":"1994-10-15"`)
	assert.NotContains(t, string(marshalled), "phone_number")
}

// The three IDDV responses all declare _links, which was absent entirely until the review pass.
func TestIdDocumentVerificationResponses_CarryTheirLinks(t *testing.T) {
	var verification IdDocumentVerificationResponse
	assert.NoError(t, json.Unmarshal([]byte(
		`{"id":"iddv_1","_links":{"self":{"href":"https://idv.checkout.com/iddv_1"},"applicant":{"href":"https://idv.checkout.com/aplt_1"}}}`),
		&verification))
	assert.NotNil(t, verification.Links.Self)
	assert.NotNil(t, verification.Links.Applicant)

	var attempt IdDocumentVerificationAttemptResponse
	assert.NoError(t, json.Unmarshal([]byte(
		`{"id":"datp_1","_links":{"self":{"href":"https://idv.checkout.com/datp_1"}}}`), &attempt))
	assert.NotNil(t, attempt.Links.Self)

	var list IdDocumentVerificationAttemptsResponse
	assert.NoError(t, json.Unmarshal([]byte(
		`{"total_count":25,"skip":0,"limit":10,"data":[],"_links":{"next":{"href":"https://idv.checkout.com/a?skip=10"}}}`),
		&list))
	assert.NotNil(t, list.Links.Next)
	assert.Contains(t, *list.Links.Next.HRef, "skip=10")
}
