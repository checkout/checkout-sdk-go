package accounts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies CurrencySchedule aligned with the 2026-08-05 Checkout.com swagger delta:
// balance_minimum and carry_forward_enabled (SaaS seller ISV variant) plus
// payment_instrument_id on both the update request and the retrieve response.

func TestCurrencySchedule_MarshalNewFieldsPresentWhenSet(t *testing.T) {
	balanceMinimum := int64(500)
	carryForward := false

	schedule := CurrencySchedule{
		Enabled:             true,
		Threshold:           100,
		PaymentInstrumentId: "ppi_w4jelhppmfiufdnatam37wrfc4",
		BalanceMinimum:      &balanceMinimum,
		CarryForwardEnabled: &carryForward,
		Recurrence:          NewScheduleFrequencyWeeklyRequest([]DaySchedule{Monday}),
	}

	marshalled, err := json.Marshal(schedule)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"payment_instrument_id":"ppi_w4jelhppmfiufdnatam37wrfc4"`)
	assert.Contains(t, body, `"balance_minimum":500`)
	assert.Contains(t, body, `"carry_forward_enabled":false`)
}

func TestCurrencySchedule_MarshalNewFieldsAbsentWhenUnset(t *testing.T) {
	schedule := CurrencySchedule{
		Enabled:    true,
		Threshold:  100,
		Recurrence: NewScheduleFrequencyDailyRequest(),
	}

	marshalled, err := json.Marshal(schedule)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.NotContains(t, body, "payment_instrument_id")
	assert.NotContains(t, body, "balance_minimum")
	assert.NotContains(t, body, "carry_forward_enabled")
}

func TestPayoutSchedule_UnmarshalNewFieldsPresent(t *testing.T) {
	payload := `{
		"GBP": {
			"enabled": true,
			"threshold": 100,
			"balance_minimum": 500,
			"carry_forward_enabled": true,
			"payment_instrument_id": "ppi_w4jelhppmfiufdnatam37wrfc4",
			"recurrence": {
				"frequency": "weekly",
				"by_day": ["monday"]
			}
		},
		"_links": {
			"self": {
				"href": "https://api.checkout.com/accounts/entities/ent_wxglze3wwywujg4nna5fb7ldli"
			}
		}
	}`

	var schedule PayoutSchedule
	assert.NoError(t, json.Unmarshal([]byte(payload), &schedule))

	gbp := schedule.Currency["GBP"]
	assert.True(t, gbp.Enabled)
	assert.Equal(t, 100, gbp.Threshold)
	assert.Equal(t, "ppi_w4jelhppmfiufdnatam37wrfc4", gbp.PaymentInstrumentId)
	if assert.NotNil(t, gbp.BalanceMinimum) {
		assert.Equal(t, int64(500), *gbp.BalanceMinimum)
	}
	if assert.NotNil(t, gbp.CarryForwardEnabled) {
		assert.True(t, *gbp.CarryForwardEnabled)
	}
}

func TestPayoutSchedule_UnmarshalNewFieldsAbsent(t *testing.T) {
	payload := `{
		"USD": {
			"enabled": true,
			"threshold": 100,
			"recurrence": {
				"frequency": "monthly",
				"by_month_day": [1, 15]
			}
		},
		"_links": {}
	}`

	var schedule PayoutSchedule
	assert.NoError(t, json.Unmarshal([]byte(payload), &schedule))

	usd := schedule.Currency["USD"]
	assert.Empty(t, usd.PaymentInstrumentId)
	assert.Nil(t, usd.BalanceMinimum)
	assert.Nil(t, usd.CarryForwardEnabled)
}
