package sessions

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// sessionChallengeIndicatorValues are the nine values accepted by
// SessionRequest.challenge_indicator, per the API Reference ChallengeIndicator schema, in spec order.
var sessionChallengeIndicatorValues = []SessionChallengeIndicator{
	SessionChallengeNoPreference,
	SessionChallengeNoChallengeRequested,
	SessionChallengeRequested,
	SessionChallengeRequestedMandate,
	SessionChallengeLowValue,
	SessionChallengeTrustedListing,
	SessionChallengeTrustedListingPrompt,
	SessionChallengeTransactionRiskAssessment,
	SessionChallengeDataShare,
}

var sessionChallengeIndicatorWireValues = []string{
	"no_preference",
	"no_challenge_requested",
	"challenge_requested",
	"challenge_requested_mandate",
	"low_value",
	"trusted_listing",
	"trusted_listing_prompt",
	"transaction_risk_assessment",
	"data_share",
}

// paymentChallengeIndicatorWireValues are the four values accepted by the 3ds.challenge_indicator
// field on payments, hosted payments, payment links and payment sessions.
var paymentChallengeIndicatorWireValues = []string{
	"no_preference",
	"no_challenge_requested",
	"challenge_requested",
	"challenge_requested_mandate",
}

func TestSessionChallengeIndicator_MatchesSpecValueSet(t *testing.T) {
	assert.Len(t, sessionChallengeIndicatorValues, 9)

	for i, value := range sessionChallengeIndicatorValues {
		assert.Equal(t, sessionChallengeIndicatorWireValues[i], string(value))
	}
}

func TestSessionChallengeIndicator_SerializesOnSessionRequest(t *testing.T) {
	for i, value := range sessionChallengeIndicatorValues {
		request := SessionRequest{ChallengeIndicator: value}

		data, err := json.Marshal(request)
		assert.Nil(t, err)

		var decoded map[string]interface{}
		assert.Nil(t, json.Unmarshal(data, &decoded))
		assert.Equal(t, sessionChallengeIndicatorWireValues[i], decoded["challenge_indicator"])
	}
}

// The API Reference specifies only the four base values for the session response field, but the
// request accepts nine. An exemption value echoed back must still deserialize.
func TestSessionChallengeIndicator_DeserializesOnGetSessionResponse(t *testing.T) {
	for i, wire := range sessionChallengeIndicatorWireValues {
		payload := []byte(`{"challenge_indicator":"` + wire + `"}`)

		var response GetSessionResponse
		assert.Nil(t, json.Unmarshal(payload, &response))
		assert.NotNil(t, response.ChallengeIndicator)
		assert.Equal(t, sessionChallengeIndicatorValues[i], *response.ChallengeIndicator)
	}
}

func TestNewSessionRequest_DefaultsToNoPreference(t *testing.T) {
	request := NewSessionRequest()

	assert.Equal(t, SessionChallengeNoPreference, request.ChallengeIndicator)

	data, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.Contains(t, string(data), `"challenge_indicator":"no_preference"`)
}

// The shared payments enum must stay narrow: the exemption values are rejected by the
// 3ds.challenge_indicator fields that use common.ChallengeIndicator. This is the guard the split
// exists to provide.
func TestCommonChallengeIndicator_StaysNarrow(t *testing.T) {
	shared := []common.ChallengeIndicator{
		common.NoPreference,
		common.NoChallengeRequested,
		common.ChallengeRequested,
		common.ChallengeRequestedMandate,
	}

	assert.Len(t, shared, len(paymentChallengeIndicatorWireValues))
	for i, value := range shared {
		assert.Equal(t, paymentChallengeIndicatorWireValues[i], string(value))
	}

	exemptions := []string{
		"low_value",
		"trusted_listing",
		"trusted_listing_prompt",
		"transaction_risk_assessment",
		"data_share",
	}
	for _, exemption := range exemptions {
		assert.NotContains(t, paymentChallengeIndicatorWireValues, exemption)
		assert.Contains(t, sessionChallengeIndicatorWireValues, exemption)
	}
}
