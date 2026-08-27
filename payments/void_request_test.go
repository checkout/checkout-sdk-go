package payments

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies VoidRequest (POST /payments/{id}/voids) marshalling: amount serializes
// under the "amount" key, and an unset amount is absent from the JSON body so
// existing full-void callers keep sending the same request.
func TestVoidRequest_MarshalAmount(t *testing.T) {
	amount := int64(6540)
	request := VoidRequest{
		Amount:    &amount,
		Reference: "ORD-5023-4E89",
	}

	data, err := json.Marshal(request)
	assert.NoError(t, err)

	var body map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &body))
	assert.Equal(t, float64(6540), body["amount"])
	assert.Equal(t, "ORD-5023-4E89", body["reference"])
}

func TestVoidRequest_MarshalWithoutAmount(t *testing.T) {
	request := VoidRequest{
		Reference: "ORD-5023-4E89",
	}

	data, err := json.Marshal(request)
	assert.NoError(t, err)

	var body map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &body))
	_, present := body["amount"]
	assert.False(t, present)
	assert.Equal(t, map[string]interface{}{"reference": "ORD-5023-4E89"}, body)
}

// The spec allows an explicit zero amount (minimum 0) and an entirely empty
// request body (every field optional): both must be expressible, which is why
// Amount is a pointer rather than a value with omitempty.
func TestVoidRequest_MarshalExplicitZeroAmount(t *testing.T) {
	amount := int64(0)
	request := VoidRequest{Amount: &amount}

	data, err := json.Marshal(request)
	assert.NoError(t, err)

	var body map[string]interface{}
	assert.NoError(t, json.Unmarshal(data, &body))
	assert.Equal(t, float64(0), body["amount"])
}

func TestVoidRequest_MarshalEmptyRequest(t *testing.T) {
	data, err := json.Marshal(VoidRequest{})
	assert.NoError(t, err)
	assert.Equal(t, "{}", string(data))
}
