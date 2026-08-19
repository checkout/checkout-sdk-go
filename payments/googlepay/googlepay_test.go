package googlepay

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// The body a real POST /googlepay/enrollments returns in sandbox. Used instead of the swagger
// example because the swagger example is the thing that was wrong: it omits merchant_id and
// names the timestamp tosAcceptedTime.
const realEnrollmentResponse = `{
	"merchant_id": "12345678901234567890",
	"tos_accepted_time": "2026-08-13T09:12:41Z",
	"state": "ACTIVE"
}`

func TestCreateEnrollmentResponseDeserialization(t *testing.T) {
	t.Run("should deserialize the real enrollment response", func(t *testing.T) {
		var response CreateEnrollmentResponse

		err := json.Unmarshal([]byte(realEnrollmentResponse), &response)

		assert.Nil(t, err)
		assert.Equal(t, "12345678901234567890", response.MerchantId)
		assert.NotNil(t, response.TosAcceptedTime)
		assert.Equal(t, time.Date(2026, 8, 13, 9, 12, 41, 0, time.UTC), response.TosAcceptedTime.UTC())
		assert.Equal(t, Active, response.State)
	})

	// This is the regression that matters most in Go. The field was tagged tosAcceptedTime, which
	// never matches the real payload, so it came back nil for every caller while looking correct.
	t.Run("should populate tos_accepted_time, which the camelCase tag never matched", func(t *testing.T) {
		var response CreateEnrollmentResponse

		err := json.Unmarshal([]byte(realEnrollmentResponse), &response)

		assert.Nil(t, err)
		assert.NotNil(t, response.TosAcceptedTime, "tos_accepted_time must deserialize, it was silently nil before")
	})

	// merchant_id is what the caller needs to initialise Google Pay on the client, so losing it
	// is the whole defect.
	t.Run("should not silently drop merchant_id", func(t *testing.T) {
		var response CreateEnrollmentResponse

		err := json.Unmarshal([]byte(realEnrollmentResponse), &response)

		assert.Nil(t, err)
		assert.NotEmpty(t, response.MerchantId)
	})

	t.Run("should leave the fields at their zero value when absent", func(t *testing.T) {
		var response CreateEnrollmentResponse

		err := json.Unmarshal([]byte(`{"state":"INACTIVE"}`), &response)

		assert.Nil(t, err)
		assert.Empty(t, response.MerchantId)
		assert.Nil(t, response.TosAcceptedTime)
		assert.Equal(t, Inactive, response.State)
	})

	t.Run("should marshal the field names the API uses", func(t *testing.T) {
		tosAcceptedTime := time.Date(2026, 8, 13, 9, 12, 41, 0, time.UTC)
		response := CreateEnrollmentResponse{
			MerchantId:      "12345678901234567890",
			TosAcceptedTime: &tosAcceptedTime,
			State:           Active,
		}

		body, err := json.Marshal(response)

		assert.Nil(t, err)
		assert.Contains(t, string(body), `"merchant_id":"12345678901234567890"`)
		assert.Contains(t, string(body), `"tos_accepted_time":"2026-08-13T09:12:41Z"`)
		assert.NotContains(t, string(body), "tosAcceptedTime")
	})
}
