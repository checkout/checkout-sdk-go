package nas

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies PaymentInstructionResponse aligned with the 2026-08-05 Checkout.com
// swagger delta: funds_transfer_type on the card payout response instruction.

func TestPaymentInstructionResponse_FundsTransferTypePresentWhenSet(t *testing.T) {
	payload := `{"value_date":"2026-08-05T00:00:00Z","funds_transfer_type":"AA"}`

	var instruction PaymentInstructionResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &instruction))
	assert.Equal(t, "AA", instruction.FundsTransferType)

	marshalled, err := json.Marshal(instruction)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"funds_transfer_type":"AA"`)
}

func TestPaymentInstructionResponse_FundsTransferTypeAbsentWhenUnset(t *testing.T) {
	instruction := PaymentInstructionResponse{}

	marshalled, err := json.Marshal(instruction)
	assert.NoError(t, err)
	assert.NotContains(t, string(marshalled), "funds_transfer_type")
}
