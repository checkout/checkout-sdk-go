package identityverification

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/identities"
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
