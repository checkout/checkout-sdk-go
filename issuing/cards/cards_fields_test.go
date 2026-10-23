package issuing

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// Verifies the card scheduling fields:
//   - scheduled_activation_date on add-card-request, update-card-request and get-card-response
//     (renamed from activation_date; the IssuingActivationDate schema was removed)
//   - revocation_date on add-card-request and update-card-request, now format: date
//   - the deprecated return-encrypted-cvv and Encryption-Key request headers (still sent when set)
//
// update-card-response never includes encrypted_cvv (see the tests further down).

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

// Verifies these fields:
//   - scheduled_revocation_date on add-card-request, update-card-request, get-card-response
//   - status on update-card-request
//   - last_activated_on on activate-card-response, add-card-response/get-card-response
//   - update-card-response drops encrypted_cvv and keeps last_modified_date/_links

func TestCardDetailsRequest_ScheduledRevocationDate(t *testing.T) {
	request := CardDetailsRequest{
		Type:                    Virtual,
		CardholderId:            "crh_test_abcdefghijklmnopqr",
		ScheduledRevocationDate: "2027-03-12",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"scheduled_revocation_date":"2027-03-12"`)

	var decoded CardDetailsRequest
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, request.ScheduledRevocationDate, decoded.ScheduledRevocationDate)
}

func TestCardUpdateRequest_ScheduledRevocationDateAndStatus(t *testing.T) {
	request := CardUpdateRequest{
		Reference:               "X-123456-N11",
		ScheduledRevocationDate: "2027-03-12",
		Status:                  ActiveCardStatusUpdate,
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"scheduled_revocation_date":"2027-03-12"`)
	assert.Contains(t, body, `"status":"active"`)

	var decoded CardUpdateRequest
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, request.ScheduledRevocationDate, decoded.ScheduledRevocationDate)
	assert.Equal(t, request.Status, decoded.Status)
}

func TestCardUpdateRequest_OmitsScheduledRevocationDateAndStatusWhenUnset(t *testing.T) {
	marshalled, err := json.Marshal(CardUpdateRequest{})
	assert.NoError(t, err)
	body := string(marshalled)
	assert.NotContains(t, body, "scheduled_revocation_date")
	assert.NotContains(t, body, "status")
}

func TestCardDetailsData_DeserializeScheduledRevocationAndLastActivatedOn(t *testing.T) {
	// Swagger example payload shape for get-card-response.
	payload := `{
		"id": "crd_test_abcdefghijklmnopqr",
		"scheduled_revocation_date": "2027-03-12",
		"last_activated_on": "2019-09-10T10:11:12Z"
	}`

	var data CardDetailsData
	assert.NoError(t, json.Unmarshal([]byte(payload), &data))
	assert.Equal(t, "2027-03-12", data.ScheduledRevocationDate)
	assert.NotNil(t, data.LastActivatedOn)
	assert.Equal(t, "2019-09-10T10:11:12Z", data.LastActivatedOn.UTC().Format(time.RFC3339))
}

func TestCardDetailsData_LastActivatedOnNilWhenAbsent(t *testing.T) {
	payload := `{"id": "crd_test_abcdefghijklmnopqr"}`

	var data CardDetailsData
	assert.NoError(t, json.Unmarshal([]byte(payload), &data))
	assert.Nil(t, data.LastActivatedOn)
}

func TestActivateCardResponse_DeserializeLastActivatedOn(t *testing.T) {
	// Swagger example payload for activate-card-response.
	payload := `{
		"last_activated_on": "2019-09-10T10:11:12Z",
		"_links": {
			"self": {"href": "https://api.checkout.com/issuing/cards/crd_test"}
		}
	}`

	var response ActivateCardResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.NotNil(t, response.LastActivatedOn)
	assert.Equal(t, "2019-09-10T10:11:12Z", response.LastActivatedOn.UTC().Format(time.RFC3339))
	assert.Contains(t, response.Links, "self")
}

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

// The json tags double as header names. The headers are deprecated and no longer in the API
// specification, but they must keep the names they were sent with until they are removed.
func TestCardUpdateHeaders_TagsKeepTheHeaderNames(t *testing.T) {
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

func TestCardUpdateResponse_DeserializeSwaggerExample(t *testing.T) {
	payload := `{
		"scheduled_revocation_date": "2027-03-12",
		"last_activated_on": "2019-09-10T10:11:12Z",
		"last_modified_date": "2019-09-10T10:11:12Z",
		"encrypted_cvv": "should-be-ignored",
		"_links": {
			"self": {"href": "https://api.checkout.com/issuing/cards/crd_test"},
			"credentials": {"href": "https://api.checkout.com/issuing/cards/crd_test/credentials"},
			"revoke": {"href": "https://api.checkout.com/issuing/cards/crd_test/revoke"},
			"controls": {"href": "https://api.checkout.com/issuing/controls?target_id=crd_test"}
		}
	}`

	var response CardUpdateResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.Equal(t, "2027-03-12", response.ScheduledRevocationDate)
	assert.NotNil(t, response.LastActivatedOn)
	assert.NotNil(t, response.LastModifiedDate)
	assert.Contains(t, response.Links, "self")
	assert.Contains(t, response.Links, "credentials")
	assert.Contains(t, response.Links, "revoke")
	assert.Contains(t, response.Links, "controls")

	marshalled, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "encrypted_cvv")
}

// update-card-response has a virtual/physical discriminator; the virtual variant adds
// is_single_use.
func TestCardUpdateResponse_DeserializesIsSingleUse(t *testing.T) {
	payload := `{"type": "virtual", "last_modified_date": "2019-09-10T10:11:12Z", "is_single_use": true}`

	var response CardUpdateResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.True(t, response.IsSingleUse)
}

func TestCardUpdateResponse_IsSingleUseOmittedWhenAbsent(t *testing.T) {
	marshalled, err := json.Marshal(CardUpdateResponse{})
	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "is_single_use")
}

func TestCardUpdateResponse_FullRoundtrip(t *testing.T) {
	payload := `{
		"id": "crd_fa6psq242dcd6fdn5gifcq1491",
		"client_id": "cli_vkuhvk4vjn2edkps7dfsq6emqm",
		"entity_id": "ent_fa6psq242dcd6fdn5gifcq1491",
		"cardholder_id": "crh_d3ozhf43pcq2xbldn2g45qnb44",
		"card_product_id": "pro_7syjig3jq3mezlc3vjrdpfitl4",
		"user_id": "usr_fa6psq242dcd6fdn5gifcq1491",
		"last_four": "1234",
		"expiry_month": 5,
		"expiry_year": 2025,
		"status": "active",
		"display_name": "JOHN KENNEDY",
		"type": "virtual",
		"billing_currency": "USD",
		"issuing_country": "US",
		"reference": "X-123456-N11",
		"metadata": {"udf1": "metadata1"},
		"scheduled_revocation_date": "2027-03-12",
		"revocation_date": "2027-03-12",
		"scheduled_activation_date": "2026-06-01T10:00Z",
		"last_activated_on": "2019-09-10T10:11:12Z",
		"root_card_id": "crd_root",
		"parent_card_id": "crd_parent",
		"scheme": "visa",
		"created_date": "2021-09-09T19:41:39Z",
		"last_modified_date": "2021-09-09T19:41:39Z",
		"is_single_use": true,
		"_links": {"self": {"href": "https://api.checkout.com/issuing/cards/crd_test"}}
	}`

	var response CardUpdateResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assertFullCardUpdateResponse(t, response)

	marshalled, err := json.Marshal(response)
	assert.NoError(t, err)
	var roundtrip CardUpdateResponse
	assert.NoError(t, json.Unmarshal(marshalled, &roundtrip))
	assertFullCardUpdateResponse(t, roundtrip)
	assert.Equal(t, response, roundtrip)
}

func assertFullCardUpdateResponse(t *testing.T, response CardUpdateResponse) {
	assert.Equal(t, "crd_fa6psq242dcd6fdn5gifcq1491", response.Id)
	assert.Equal(t, "cli_vkuhvk4vjn2edkps7dfsq6emqm", response.ClientId)
	assert.Equal(t, "ent_fa6psq242dcd6fdn5gifcq1491", response.EntityId)
	assert.Equal(t, "crh_d3ozhf43pcq2xbldn2g45qnb44", response.CardholderId)
	assert.Equal(t, "pro_7syjig3jq3mezlc3vjrdpfitl4", response.CardProductId)
	assert.Equal(t, "usr_fa6psq242dcd6fdn5gifcq1491", response.UserId)
	assert.Equal(t, "1234", response.LastFour)
	assert.Equal(t, 5, response.ExpiryMonth)
	assert.Equal(t, 2025, response.ExpiryYear)
	assert.Equal(t, CardStatus("active"), response.Status)
	assert.Equal(t, "JOHN KENNEDY", response.DisplayName)
	assert.Equal(t, Virtual, response.Type)
	assert.Equal(t, "USD", string(response.BillingCurrency))
	assert.Equal(t, "US", string(response.IssuingCountry))
	assert.Equal(t, "X-123456-N11", response.Reference)
	if assert.NotNil(t, response.Metadata) {
		assert.Equal(t, "metadata1", response.Metadata.Udf1)
	}
	assert.Equal(t, "2027-03-12", response.ScheduledRevocationDate)
	assert.Equal(t, "2027-03-12", response.RevocationDate)
	assert.Equal(t, "2026-06-01T10:00Z", response.ScheduledActivationDate)
	if assert.NotNil(t, response.LastActivatedOn) {
		assert.Equal(t, "2019-09-10T10:11:12Z", response.LastActivatedOn.UTC().Format(time.RFC3339))
	}
	assert.Equal(t, "crd_root", response.RootCardId)
	assert.Equal(t, "crd_parent", response.ParentCardId)
	assert.Equal(t, CardScheme("visa"), response.Scheme)
	if assert.NotNil(t, response.CreatedDate) {
		assert.Equal(t, "2021-09-09T19:41:39Z", response.CreatedDate.UTC().Format(time.RFC3339))
	}
	if assert.NotNil(t, response.LastModifiedDate) {
		assert.Equal(t, "2021-09-09T19:41:39Z", response.LastModifiedDate.UTC().Format(time.RFC3339))
	}
	assert.True(t, response.IsSingleUse)
	if assert.Contains(t, response.Links, "self") {
		if assert.NotNil(t, response.Links["self"].HRef) {
			assert.Equal(t, "https://api.checkout.com/issuing/cards/crd_test", *response.Links["self"].HRef)
		}
	}
}

func TestRevokeAndSuspendCardRequest_EmptyBodySendsEmptyObject(t *testing.T) {
	revoke, err := common.Marshal(RevokeCardRequest{})
	assert.NoError(t, err)
	assert.JSONEq(t, `{}`, revoke.String())

	suspend, err := common.Marshal(SuspendCardRequest{})
	assert.NoError(t, err)
	assert.JSONEq(t, `{}`, suspend.String())

	withReason, err := common.Marshal(RevokeCardRequest{Reason: ReportedLost})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"reason":"reported_lost"}`, withReason.String())
}
