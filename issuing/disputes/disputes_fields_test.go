package issuing

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies the dispute schemas aligned with the 2026-06-29 Checkout.com swagger delta serialize to
// their exact snake_case wire names:
//   - amend-dispute-request (new): reason, amount, evidence, fraud_details,
//     reason_change_justification, action_response
//   - fraud_details on create-dispute-request and escalate-dispute-request
//   - IssuingDisputeFraudType (11-value enum)
//   - action_details on the dispute response

func TestAmendDisputeRequest_RoundtripAllFields(t *testing.T) {
	amount := int64(1500)
	request := AmendDisputeRequest{
		Reason: "4807",
		Amount: &amount,
		Evidence: []DisputeEvidence{
			{Name: "amended_evidence.pdf", Content: "QW1lbmRlZA==", Description: "Amended evidence"},
		},
		FraudDetails: &IssuingDisputeFraudDetails{
			FraudType:   FraudTypeCardNotPresentFraud,
			Description: "No online purchases on this date.",
		},
		ReasonChangeJustification: "New evidence confirms an unauthorized transaction.",
		ActionResponse:            "Updated the reason code as requested.",
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"reason":"4807"`)
	assert.Contains(t, body, `"amount":1500`)
	assert.Contains(t, body, `"evidence":`)
	assert.Contains(t, body, `"fraud_details":`)
	assert.Contains(t, body, `"fraud_type":"card_not_present_fraud"`)
	assert.Contains(t, body, `"reason_change_justification":"New evidence confirms an unauthorized transaction."`)
	assert.Contains(t, body, `"action_response":"Updated the reason code as requested."`)

	var decoded AmendDisputeRequest
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, request.Reason, decoded.Reason)
	assert.Equal(t, *request.Amount, *decoded.Amount)
	assert.Equal(t, request.Evidence, decoded.Evidence)
	assert.Equal(t, request.FraudDetails.FraudType, decoded.FraudDetails.FraudType)
	assert.Equal(t, request.FraudDetails.Description, decoded.FraudDetails.Description)
	assert.Equal(t, request.ReasonChangeJustification, decoded.ReasonChangeJustification)
	assert.Equal(t, request.ActionResponse, decoded.ActionResponse)
}

func TestAmendDisputeRequest_FromSwaggerExample(t *testing.T) {
	payload := `{
		"reason": "4808",
		"amount": 2000,
		"evidence": [{"name": "file.pdf", "content": "Zm9v", "description": "doc"}],
		"fraud_details": {"fraud_type": "counterfeit_card", "description": "cloned card"},
		"reason_change_justification": "reason updated",
		"action_response": "question about the requested change"
	}`

	var request AmendDisputeRequest
	assert.NoError(t, json.Unmarshal([]byte(payload), &request))
	assert.Equal(t, "4808", request.Reason)
	assert.Equal(t, int64(2000), *request.Amount)
	assert.Len(t, request.Evidence, 1)
	assert.Equal(t, "file.pdf", request.Evidence[0].Name)
	assert.Equal(t, FraudTypeCounterfeitCard, request.FraudDetails.FraudType)
	assert.Equal(t, "cloned card", request.FraudDetails.Description)
	assert.Equal(t, "reason updated", request.ReasonChangeJustification)
	assert.Equal(t, "question about the requested change", request.ActionResponse)
}

func TestAmendDisputeRequest_OmitsUnsetFields(t *testing.T) {
	marshalled, err := json.Marshal(AmendDisputeRequest{})
	assert.NoError(t, err)
	body := string(marshalled)
	assert.NotContains(t, body, "reason")
	assert.NotContains(t, body, "amount")
	assert.NotContains(t, body, "fraud_details")
	assert.NotContains(t, body, "reason_change_justification")
	assert.NotContains(t, body, "action_response")
}

// Every swagger fraud_type value must serialize to its exact wire string (review-integrity §8).
func TestIssuingDisputeFraudType_AllValuesSerialize(t *testing.T) {
	cases := map[IssuingDisputeFraudType]string{
		FraudTypeCardLost:                  "card_lost",
		FraudTypeCardStolen:                "card_stolen",
		FraudTypeCardNeverReceived:         "card_never_received",
		FraudTypeFraudulentAccount:         "fraudulent_account",
		FraudTypeCounterfeitCard:           "counterfeit_card",
		FraudTypeAccountTakeover:           "account_takeover",
		FraudTypeCardNotPresentFraud:       "card_not_present_fraud",
		FraudTypeMerchantMisrepresentation: "merchant_misrepresentation",
		FraudTypeCardholderManipulation:    "cardholder_manipulation",
		FraudTypeIncorrectProcessing:       "incorrect_processing",
		FraudTypeOther:                     "other",
	}
	assert.Len(t, cases, 11)

	for value, wire := range cases {
		marshalled, err := json.Marshal(IssuingDisputeFraudDetails{FraudType: value})
		assert.NoError(t, err)
		assert.Contains(t, string(marshalled), `"fraud_type":"`+wire+`"`)

		var decoded IssuingDisputeFraudDetails
		assert.NoError(t, json.Unmarshal([]byte(`{"fraud_type":"`+wire+`"}`), &decoded))
		assert.Equal(t, value, decoded.FraudType)
	}
}

func TestCreateDisputeRequest_FraudDetails(t *testing.T) {
	request := CreateDisputeRequest{
		TransactionId: "trx_test_abcdefghijklmnopqr",
		Reason:        "4837",
		FraudDetails: &IssuingDisputeFraudDetails{
			FraudType: FraudTypeCardLost,
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"fraud_details":`)
	assert.Contains(t, string(marshalled), `"fraud_type":"card_lost"`)
}

func TestEscalateDisputeRequest_FraudDetails(t *testing.T) {
	request := EscalateDisputeRequest{
		Justification: "Escalating",
		FraudDetails: &IssuingDisputeFraudDetails{
			FraudType: FraudTypeAccountTakeover,
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"fraud_details":`)
	assert.Contains(t, string(marshalled), `"fraud_type":"account_takeover"`)
}

func TestIssuingDisputeResponse_ActionDetails(t *testing.T) {
	payload := `{
		"id": "idsp_test_12345abcdefghijklmnop",
		"status": "action_required",
		"action_details": {
			"instructions": "Provide the missing receipt.",
			"last_action_response": "Awaiting cardholder response."
		}
	}`

	var response IssuingDisputeResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.NotNil(t, response.ActionDetails)
	assert.Equal(t, "Provide the missing receipt.", response.ActionDetails.Instructions)
	assert.Equal(t, "Awaiting cardholder response.", response.ActionDetails.LastActionResponse)
}
