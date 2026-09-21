package faceauthentication

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/identities"
)

// Verifies the face authentication attempt assets response (FavAttemptAssets)
// deserializes from the API wire format, including the asset type enum and the
// nested _links.asset_url HAL link.
func TestFaceAuthenticationAttemptAssetsResponse_Unmarshal(t *testing.T) {
	payload := `{
		"total_count":2,
		"skip":0,
		"limit":10,
		"data":[
			{"type":"face_image","_links":{"asset_url":{"href":"https://example.com/face-image.jpg"}}},
			{"type":"face_video","_links":{"asset_url":{"href":"https://example.com/face-video.mp4"}}}
		],
		"_links":{"self":{"href":"https://example.com/assets"},"next":{"href":"https://example.com/assets?skip=10"}}
	}`

	var response FaceAuthenticationAttemptAssetsResponse
	err := json.Unmarshal([]byte(payload), &response)

	assert.NoError(t, err)
	assert.Equal(t, 2, response.TotalCount)
	assert.Equal(t, 0, response.Skip)
	assert.Equal(t, 10, response.Limit)
	assert.Len(t, response.Data, 2)
	assert.Equal(t, identities.FaceImageFavAsset, response.Data[0].Type)
	assert.NotNil(t, response.Data[0].Links.AssetUrl.HRef)
	assert.Equal(t, "https://example.com/face-image.jpg", *response.Data[0].Links.AssetUrl.HRef)
	assert.Equal(t, identities.FaceVideoFavAsset, response.Data[1].Type)
	assert.Contains(t, response.Links, "self")
}

// D1: the face authentication attempt takes the narrow FavClientInformation shape, so the two
// IDV-only document fields cannot be sent here even by mistake. The type is the guard.
func TestCreateFaceAuthenticationAttemptRequest_PhoneNumberAndNarrowClientInformation(t *testing.T) {
	request := CreateFaceAuthenticationAttemptRequest{
		RedirectUrl: "https://example.com/success",
		PhoneNumber: &identities.PhoneNumber{CountryCode: "+1", Number: "5555550102"},
		ClientInformation: &identities.ClientInformation{
			PreSelectedResidenceCountry: common.US,
			PreSelectedLanguage:         "en-US",
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"phone_number":{"country_code":"+1","number":"5555550102"}`)
	assert.NotContains(t, body, "pre_selected_document_type")
	assert.NotContains(t, body, "pre_selected_document_issuing_country")
}

// Part E: risk labels are typed on the FAV response too.
func TestFaceAuthenticationResponse_TypedRiskLabels(t *testing.T) {
	payload := `{"id":"fav_1","status":"approved","risk_labels":["risky_document_format"],"face":{"image_signed_url":"https://example.com/f.png"}}`

	var response FaceAuthenticationResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))

	assert.Equal(t, []identities.RiskLabel{identities.RiskyDocumentFormat}, response.RiskLabels)
	assert.NotNil(t, response.Face)
}

// The FAV response has always used the `face` tag correctly; P11 was IDV only. Guard so a future
// "consistency" edit does not break the one that works.
func TestFaceAuthenticationResponse_FaceUsesTheSpecTag(t *testing.T) {
	var response FaceAuthenticationResponse
	assert.NoError(t, json.Unmarshal(
		[]byte(`{"face":{"image_signed_url":"https://example.com/f.png"}}`), &response))
	assert.NotNil(t, response.Face)
}

// Part E on the attempt response.
func TestFaceAuthenticationAttemptResponse_PhoneNumber(t *testing.T) {
	payload := `{"id":"att_1","status":"completed","phone_number":{"country_code":"+1","number":"5555550102"}}`

	var response FaceAuthenticationAttemptResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))

	assert.NotNil(t, response.PhoneNumber)
	assert.Equal(t, "5555550102", response.PhoneNumber.Number)
}
