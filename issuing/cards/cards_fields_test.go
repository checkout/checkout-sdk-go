package issuing

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies the card scheduling fields aligned with the 2026-06-29 Checkout.com swagger delta:
//   - activation_date on add-card-request and update-card-request
//   - revocation_date on add-card-request and update-card-request
//   - activation_date / revocation_date on get-card-response

func TestCardDetailsRequest_ActivationAndRevocationDate(t *testing.T) {
	request := CardDetailsRequest{
		Type:           Virtual,
		CardholderId:   "crh_test_abcdefghijklmnopqr",
		ActivationDate: "2026-06-01T10:00Z",
		RevocationDate: "2026-12-01",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"activation_date":"2026-06-01T10:00Z"`)
	assert.Contains(t, body, `"revocation_date":"2026-12-01"`)

	var decoded CardDetailsRequest
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, request.ActivationDate, decoded.ActivationDate)
	assert.Equal(t, request.RevocationDate, decoded.RevocationDate)
}

func TestCardUpdateRequest_ActivationAndRevocationDate(t *testing.T) {
	request := CardUpdateRequest{
		Reference:      "X-123456-N11",
		ActivationDate: "2026-06-01T10:00Z",
		RevocationDate: "2026-12-01",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"activation_date":"2026-06-01T10:00Z"`)
	assert.Contains(t, body, `"revocation_date":"2026-12-01"`)

	var decoded CardUpdateRequest
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, request.ActivationDate, decoded.ActivationDate)
	assert.Equal(t, request.RevocationDate, decoded.RevocationDate)
}

func TestCardUpdateRequest_OmitsDatesWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(CardUpdateRequest{})
	assert.NoError(t, err)
	body := string(marshalled)
	assert.NotContains(t, body, "activation_date")
	assert.NotContains(t, body, "revocation_date")
}

func TestCardDetailsData_DeserializeScheduleFields(t *testing.T) {
	payload := `{
		"id": "crd_test_abcdefghijklmnopqr",
		"activation_date": "2026-06-01T10:00Z",
		"revocation_date": "2026-12-01",
		"user_id": "usr_test_abcdefghijklmnopqr",
		"root_card_id": "crd_root_abcdefghijklmnop",
		"parent_card_id": "crd_parent_abcdefghijkl"
	}`

	var data CardDetailsData
	assert.NoError(t, json.Unmarshal([]byte(payload), &data))
	assert.Equal(t, "2026-06-01T10:00Z", data.ActivationDate)
	assert.Equal(t, "2026-12-01", data.RevocationDate)
	assert.Equal(t, "usr_test_abcdefghijklmnopqr", data.UserId)
	assert.Equal(t, "crd_root_abcdefghijklmnop", data.RootCardId)
	assert.Equal(t, "crd_parent_abcdefghijkl", data.ParentCardId)
}
