package accounts

import (
	"encoding/json"
	"fmt"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/errors"
)

// Frequency indicates how often funds should be paid out to a sub-entity.
type Frequency string

const (
	// Weekly pays out on the configured days of the week. SaaS seller (ISV)
	// sub-entities accept working days only, Monday to Friday: a schedule set
	// to a Saturday or Sunday is rejected. Standard sub-entities accept any day.
	Weekly Frequency = "weekly"
	// Daily pays out every day for standard sub-entities. For SaaS seller (ISV)
	// sub-entities it runs on working days only, Monday to Friday, with no
	// payout at weekends, based on the available balance as of 00:00 in the
	// sub-entity's time zone.
	Daily Frequency = "daily"
	// Monthly pays out on the configured days of the month. Standard
	// sub-entities accept any day from 1 to 28. SaaS seller (ISV) sub-entities
	// accept only these combinations, in any order: [1], [15], [1, 15] or [1, 16].
	Monthly Frequency = "monthly"
)

type DaySchedule string

const (
	Monday    DaySchedule = "monday"
	Tuesday   DaySchedule = "tuesday"
	Wednesday DaySchedule = "wednesday"
	Thursday  DaySchedule = "thursday"
	Friday    DaySchedule = "friday"
	Saturday  DaySchedule = "saturday"
	Sunday    DaySchedule = "sunday"
)

type (
	Recurrence interface {
		GetSchedule() Frequency
	}

	scheduleFrequencyDaily struct {
		Frequency
	}

	scheduleFrequencyWeekly struct {
		Frequency
		ByDay []DaySchedule `json:"by_day,omitempty"`
	}

	scheduleFrequencyMonthly struct {
		Frequency
		ByMonthDay []int `json:"by_month_day,omitempty"`
	}
)

func NewScheduleFrequencyDailyRequest() scheduleFrequencyDaily {
	return scheduleFrequencyDaily{
		Frequency: Daily,
	}
}

func NewScheduleFrequencyWeeklyRequest(days []DaySchedule) scheduleFrequencyWeekly {
	return scheduleFrequencyWeekly{
		Frequency: Weekly,
		ByDay:     days,
	}
}

func NewScheduleFrequencyMonthlyRequest(days []int) scheduleFrequencyMonthly {
	return scheduleFrequencyMonthly{
		Frequency:  Monthly,
		ByMonthDay: days,
	}
}

func (s scheduleFrequencyDaily) GetSchedule() Frequency {
	return s.Frequency
}

func (s scheduleFrequencyWeekly) GetSchedule() Frequency {
	return s.Frequency
}

func (s scheduleFrequencyMonthly) GetSchedule() Frequency {
	return s.Frequency
}

type (
	CurrencySchedule struct {
		Enabled   bool `json:"enabled,omitempty"`
		Threshold int  `json:"threshold,omitempty"`
		// PaymentInstrumentId is the ID of the platforms payment instrument used
		// as the payout destination for this schedule. Optional for SaaS seller
		// (ISV) schedules; if included, it must reference a verified payment
		// instrument, otherwise the request fails.
		PaymentInstrumentId string `json:"payment_instrument_id,omitempty"`
		// BalanceMinimum is the amount, in the minor units of the schedule's
		// currency, to retain in the sub-entity's available balance. Only the
		// funds above this are paid out, and no payout is generated if there are
		// none. Defaults to 0 if not set. SaaS seller (ISV) schedules only.
		BalanceMinimum *int64 `json:"balance_minimum,omitempty"`
		// CarryForwardEnabled indicates whether to carry forward any balance
		// below the configured minimum to the next payout. Always returned for
		// SaaS sellers, and defaults to false if not set. SaaS seller (ISV)
		// schedules only.
		CarryForwardEnabled *bool      `json:"carry_forward_enabled,omitempty"`
		Recurrence          Recurrence `json:"recurrence,omitempty"`
	}

	PayoutSchedule struct {
		HttpMetadata common.HttpMetadata `json:"http_metadata,omitempty"`
		Currency     map[common.Currency]CurrencySchedule
		Links        map[string]common.Link `json:"_links"`
	}
)

func (p *PayoutSchedule) UnmarshalJSON(data []byte) error {
	p.Currency = make(map[common.Currency]CurrencySchedule)

	var currencyMap map[common.Currency]currencyUnmarshaler
	if err := json.Unmarshal(data, &currencyMap); err != nil {
		return err
	}

	var currency CurrencySchedule
	for k := range currencyMap {
		if k != "_links" {
			switch currencyMap[k].Recurrence.Frequency {
			case Daily:
				var schedule map[common.Currency]dailyScheduleUnmarshaler
				if err := json.Unmarshal(data, &schedule); err != nil {
					return err
				}
				currency.Recurrence = schedule[k].Recurrence
			case Weekly:
				var schedule map[common.Currency]weeklyScheduleUnmarshaler
				if err := json.Unmarshal(data, &schedule); err != nil {
					return err
				}
				currency.Recurrence = schedule[k].Recurrence
			case Monthly:
				var schedule map[common.Currency]monthlyScheduleUnmarshaler
				if err := json.Unmarshal(data, &schedule); err != nil {
					return err
				}
				currency.Recurrence = schedule[k].Recurrence
			default:
				return errors.UnsupportedTypeError(fmt.Sprintf("%s currency frequency is unsupported", k))
			}

			currency.Enabled = currencyMap[k].Enabled
			currency.Threshold = currencyMap[k].Threshold
			currency.PaymentInstrumentId = currencyMap[k].PaymentInstrumentId
			currency.BalanceMinimum = currencyMap[k].BalanceMinimum
			currency.CarryForwardEnabled = currencyMap[k].CarryForwardEnabled
			p.Currency[k] = currency
		}
	}

	var links linksUnmarshaler
	if err := json.Unmarshal(data, &links); err != nil {
		return err
	}
	p.Links = links.Links

	return nil
}

type (
	currencyUnmarshaler struct {
		Enabled             bool   `json:"enabled,omitempty"`
		Threshold           int    `json:"threshold,omitempty"`
		PaymentInstrumentId string `json:"payment_instrument_id,omitempty"`
		BalanceMinimum      *int64 `json:"balance_minimum,omitempty"`
		CarryForwardEnabled *bool  `json:"carry_forward_enabled,omitempty"`
		Recurrence          struct {
			Frequency
		}
	}

	linksUnmarshaler struct {
		Links map[string]common.Link `json:"_links"`
	}

	dailyScheduleUnmarshaler struct {
		Recurrence scheduleFrequencyDaily `json:"recurrence,omitempty"`
	}

	weeklyScheduleUnmarshaler struct {
		Recurrence scheduleFrequencyWeekly `json:"recurrence,omitempty"`
	}

	monthlyScheduleUnmarshaler struct {
		Recurrence scheduleFrequencyMonthly `json:"recurrence,omitempty"`
	}
)
