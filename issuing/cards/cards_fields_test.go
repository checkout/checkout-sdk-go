package issuing

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies the card scheduling fields aligned with the 2026-09-02 Checkout.com swagger delta:
//   - scheduled_activation_date on add-card-request, update-card-request and get-card-response
//     (renamed from activation_date; the IssuingActivationDate schema was removed)
//   - revocation_date on add-card-request and update-card-request, now format: date
//   - encrypted_cvv on update-card-response
//   - the return-encrypted-cvv and Encryption-Key request headers

func TestCardDetailsRequest_ScheduledActivationAndRevocationDate(t *testing.T) {
	request := CardDetailsRequest{
		Type:                    Virtual,
		CardholderId:            "crh_test_abcdefghijklmnopqr",
		ScheduledActivationDate: "2026-06-01T10:00Z",
		RevocationDate:          "2026-12-01",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"scheduled_activation_date":"2026-06-01T10:00Z"`)
	assert.Contains(t, body, `"revocation_date":"2026-12-01"`)
	assert.NotContains(t, body, `"activation_date"`)

	var decoded CardDetailsRequest
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, request.ScheduledActivationDate, decoded.ScheduledActivationDate)
	assert.Equal(t, request.RevocationDate, decoded.RevocationDate)
}

func TestCardUpdateRequest_ScheduledActivationAndRevocationDate(t *testing.T) {
	request := CardUpdateRequest{
		Reference:               "X-123456-N11",
		ScheduledActivationDate: "2026-06-01T10:00Z",
		RevocationDate:          "2026-12-01",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"scheduled_activation_date":"2026-06-01T10:00Z"`)
	assert.Contains(t, body, `"revocation_date":"2026-12-01"`)
	assert.NotContains(t, body, `"activation_date"`)

	var decoded CardUpdateRequest
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, request.ScheduledActivationDate, decoded.ScheduledActivationDate)
	assert.Equal(t, request.RevocationDate, decoded.RevocationDate)
}

// The old key must not deserialize onto the new field. Go silently ignores unknown JSON keys, so
// without this guard a caller still sending activation_date would look fine and change nothing.
func TestCardUpdateRequest_IgnoresTheRemovedActivationDateKey(t *testing.T) {
	var decoded CardUpdateRequest
	assert.NoError(t, json.Unmarshal([]byte(`{"activation_date":"2026-06-01T10:00Z"}`), &decoded))
	assert.Empty(t, decoded.ScheduledActivationDate)
}

func TestCardUpdateRequest_SerializesEveryField(t *testing.T) {
	request := CardUpdateRequest{
		Reference:               "X-123456-N11",
		Metadata:                &CardMetadata{Udf1: "metadata1"},
		ExpiryMonth:             6,
		ExpiryYear:              2030,
		ScheduledActivationDate: "2026-06-01T10:00Z",
		RevocationDate:          "2027-03-12",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, map[string]interface{}{
		"reference":                 "X-123456-N11",
		"metadata":                  map[string]interface{}{"udf1": "metadata1"},
		"expiry_month":              float64(6),
		"expiry_year":               float64(2030),
		"scheduled_activation_date": "2026-06-01T10:00Z",
		"revocation_date":           "2027-03-12",
	}, decoded)
}

// The swagger example for update-card-request, verbatim.
func TestCardUpdateRequest_FromSwaggerExample(t *testing.T) {
	payload := `{
		"reference": "X-123456-N11",
		"expiry_month": 6,
		"expiry_year": 2030,
		"revocation_date": "2027-03-12",
		"scheduled_activation_date": "2026-06-01T10:00Z"
	}`

	var decoded CardUpdateRequest
	assert.NoError(t, json.Unmarshal([]byte(payload), &decoded))
	assert.Equal(t, "2026-06-01T10:00Z", decoded.ScheduledActivationDate)
	assert.Equal(t, "2027-03-12", decoded.RevocationDate)
	assert.Equal(t, 6, decoded.ExpiryMonth)
	assert.Equal(t, 2030, decoded.ExpiryYear)
}

func TestCardUpdateRequest_OmitsDatesWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(CardUpdateRequest{})
	assert.NoError(t, err)
	body := string(marshalled)
	assert.NotContains(t, body, "scheduled_activation_date")
	assert.NotContains(t, body, "revocation_date")
}

func TestCardDetailsData_DeserializeScheduleFields(t *testing.T) {
	payload := `{
		"id": "crd_test_abcdefghijklmnopqr",
		"scheduled_activation_date": "2026-06-01T10:00Z",
		"revocation_date": "2026-12-01",
		"user_id": "usr_test_abcdefghijklmnopqr",
		"root_card_id": "crd_root_abcdefghijklmnop",
		"parent_card_id": "crd_parent_abcdefghijkl"
	}`

	var data CardDetailsData
	assert.NoError(t, json.Unmarshal([]byte(payload), &data))
	assert.Equal(t, "2026-06-01T10:00Z", data.ScheduledActivationDate)
	assert.Equal(t, "2026-12-01", data.RevocationDate)
	assert.Equal(t, "usr_test_abcdefghijklmnopqr", data.UserId)
	assert.Equal(t, "crd_root_abcdefghijklmnop", data.RootCardId)
	assert.Equal(t, "crd_parent_abcdefghijkl", data.ParentCardId)
}

// Part D: the update response gains encrypted_cvv, returned only when the headers ask for it.
func TestCardUpdateResponse_DeserializesEncryptedCvv(t *testing.T) {
	payload := `{
		"last_modified_date": "2026-06-01T10:00:00Z",
		"encrypted_cvv": "oJMoNMEEUiQKYOsQ4Zd"
	}`

	var response CardUpdateResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.Equal(t, "oJMoNMEEUiQKYOsQ4Zd", response.EncryptedCvv)
	assert.NotNil(t, response.LastModifiedDate)
}

func TestCardUpdateResponse_EncryptedCvvIsEmptyWhenAbsent(t *testing.T) {
	var response CardUpdateResponse
	assert.NoError(t, json.Unmarshal([]byte(`{"last_modified_date":"2026-06-01T10:00:00Z"}`), &response))
	assert.Empty(t, response.EncryptedCvv)
}

// The headers travel on a field named Headers with json:"-", so they must never appear in the body.
func TestCardUpdateHeaders_AreNotPartOfTheRequestBody(t *testing.T) {
	body := struct {
		CardUpdateRequest
		Headers *CardUpdateHeaders `json:"-"`
	}{
		CardUpdateRequest{Reference: "X-123456-N11"},
		&CardUpdateHeaders{ReturnEncryptedCvv: "true", EncryptionKey: "MIIBIjAN"},
	}

	marshalled, err := json.Marshal(body)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"reference":"X-123456-N11"`)
	assert.NotContains(t, string(marshalled), "return-encrypted-cvv")
	assert.NotContains(t, string(marshalled), "Encryption-Key")
	assert.NotContains(t, string(marshalled), "Headers")
}

// The json tags double as header names, so they must match the swagger parameter names exactly.
func TestCardUpdateHeaders_TagsMatchTheSwaggerHeaderNames(t *testing.T) {
	marshalled, err := json.Marshal(CardUpdateHeaders{
		ReturnEncryptedCvv: "true",
		EncryptionKey:      "MIIBIjAN",
	})
	assert.NoError(t, err)

	var decoded map[string]interface{}
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, map[string]interface{}{
		"return-encrypted-cvv": "true",
		"Encryption-Key":       "MIIBIjAN",
	}, decoded)
}
