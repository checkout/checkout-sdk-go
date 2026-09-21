package addressdocumentverification

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

// advAttemptAssetsSwaggerExample is components.examples.adv_attempt_assets_response_body from
// shared/swagger-latest.json, verbatim apart from shortened hrefs. asset_url is the only link
// AdvAttemptAsset declares, and it is required, so using the spec's own example rather than a
// hand-written fixture is what keeps this honest.
const advAttemptAssetsSwaggerExample = `{
	"total_count": 1,
	"skip": 0,
	"limit": 10,
	"data": [
		{
			"type": "document",
			"_links": {
				"asset_url": {"href": "https://storage-b.env.ubble.ai/ubble-ai/NDY/address_document.png"}
			}
		}
	],
	"_links": {
		"self": {"href": "https://idv.checkout.com/address-document-verifications/adv_1/attempts/adva_1/assets"},
		"next": {"href": "https://idv.checkout.com/address-document-verifications/adv_1/attempts/adva_1/assets?skip=10"},
		"previous": {"href": "https://idv.checkout.com/address-document-verifications/adv_1/attempts/adva_1/assets?skip=0"}
	}
}`

// Part A: the new GET .../attempts/{attemptId}/assets response.
func TestAddressDocumentVerificationAttemptAssetsResponse_FromSwaggerExample(t *testing.T) {
	var response AddressDocumentVerificationAttemptAssetsResponse
	assert.NoError(t, json.Unmarshal([]byte(advAttemptAssetsSwaggerExample), &response))

	assert.Equal(t, 1, response.TotalCount)
	assert.Equal(t, 0, response.Skip)
	assert.Equal(t, 10, response.Limit)
	assert.Len(t, response.Data, 1)
	assert.Equal(t, identities.DocumentAdvAsset, response.Data[0].Type)
	assert.Equal(t, "document", string(response.Data[0].Type))
	assert.NotNil(t, response.Data[0].Links.AssetUrl.HRef)
	assert.Contains(t, *response.Data[0].Links.AssetUrl.HRef, "address_document.png")
	assert.Contains(t, response.Links, "self")
	assert.Contains(t, response.Links, "next")
	assert.Contains(t, response.Links, "previous")
}

// T7: data declares minItems 0, so an attempt with no assets yet is a legal page.
func TestAddressDocumentVerificationAttemptAssetsResponse_EmptyPage(t *testing.T) {
	payload := `{"total_count":0,"skip":0,"limit":10,"data":[],"_links":{"self":{"href":"https://idv.checkout.com/a"}}}`

	var response AddressDocumentVerificationAttemptAssetsResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))

	assert.Equal(t, 0, response.TotalCount)
	assert.Empty(t, response.Data)
	assert.Contains(t, response.Links, "self")
}

// The ADV attempt asset enum has exactly one value.
func TestAddressDocumentVerificationAttemptAssetType_Value(t *testing.T) {
	assert.Equal(t, identities.AddressDocumentVerificationAttemptAssetType("document"), identities.DocumentAdvAsset)
}

// Part C: IdvPdf replaced signed_url with pdf_report, and now requires it.
func TestAddressDocumentVerificationReportResponse_PdfReport(t *testing.T) {
	var response AddressDocumentVerificationReportResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"pdf_report":"https://www.example.com/reports/adv_1.pdf"}`), &response))

	assert.Equal(t, "https://www.example.com/reports/adv_1.pdf", response.PdfReport)
}

// The removed key must not populate the new field. Go ignores unknown JSON keys silently, so
// without this guard a stale server response would look like an empty report rather than a
// mismatch.
func TestAddressDocumentVerificationReportResponse_IgnoresSignedUrl(t *testing.T) {
	var response AddressDocumentVerificationReportResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"signed_url":"https://www.example.com/reports/adv_1.pdf"}`), &response))

	assert.Empty(t, response.PdfReport)
}

// D3: AdvAddress.country is an ISO 3166-1 alpha-2 code, so it is typed common.Country. The plan
// recorded only identities.go as the untyped Go site and missed this one.
func TestAddressDocumentVerificationAddress_CountryIsTyped(t *testing.T) {
	payload := `{
		"document_type":"utility_bill",
		"issuer":"EDF Energy",
		"full_names":["Hannah Bret"],
		"issue_date":"2024-01-15",
		"address":{"address_line1":"123 Main Street","city":"London","zip":"SW1A 1AA","country":"GB"}
	}`

	var result AddressDocumentResult
	assert.NoError(t, json.Unmarshal([]byte(payload), &result))

	assert.NotNil(t, result.Address)
	assert.Equal(t, common.GB, result.Address.Country)
	assert.Equal(t, "123 Main Street", result.Address.AddressLine1)
}

// Part E: the ADV request takes the narrower IdvDeclaredData shape, which gained birth_date.
func TestCreateAddressDocumentVerificationRequest_DeclaredDataBirthDate(t *testing.T) {
	request := CreateAddressDocumentVerificationRequest{
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
	// The three IdvIdentityDeclaredData-only fields must not appear here.
	assert.NotContains(t, string(marshalled), "phone_number")
	assert.NotContains(t, string(marshalled), "email")
	assert.NotContains(t, string(marshalled), `"address"`)
}
