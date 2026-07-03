package accounts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies PlatformsInstrumentDetailsAch aligned with the 2026-06-29 Checkout.com swagger delta:
// account_number, routing_number, account_type (enum savings/checking).

func TestInstrumentDetailsAch_Roundtrip(t *testing.T) {
	details := InstrumentDetailsAch{
		AccountNumber: "12345678",
		RoutingNumber: "021000021",
		AccountType:   InstrumentAccountChecking,
	}

	marshalled, err := json.Marshal(details)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"account_number":"12345678"`)
	assert.Contains(t, body, `"routing_number":"021000021"`)
	assert.Contains(t, body, `"account_type":"checking"`)

	var decoded InstrumentDetailsAch
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, details.AccountNumber, decoded.AccountNumber)
	assert.Equal(t, details.RoutingNumber, decoded.RoutingNumber)
	assert.Equal(t, details.AccountType, decoded.AccountType)
}

func TestInstrumentAccountType_Values(t *testing.T) {
	assert.Equal(t, InstrumentAccountType("savings"), InstrumentAccountSavings)
	assert.Equal(t, InstrumentAccountType("checking"), InstrumentAccountChecking)
}

func TestInstrumentDetailsAch_GetType(t *testing.T) {
	details := &InstrumentDetailsAch{}
	assert.Equal(t, "Ach", details.GetType())
}
