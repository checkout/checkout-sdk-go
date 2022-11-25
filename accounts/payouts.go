package accounts

import (
	"encoding/json"
	"fmt"

	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/errors"
)

type Frequency string

const (
	Weekly  Frequency = "Weekly"
	Daily   Frequency = "Daily"
	Monthly Frequency = "Monthly"
)

type DaySchedule string

const (
	Monday    DaySchedule = "Monday"
	Tuesday   DaySchedule = "Tuesday"
	Wednesday DaySchedule = "Wednesday"
	Thursday  DaySchedule = "Thursday"
	Friday    DaySchedule = "Friday"
	Saturday  DaySchedule = "Saturday"
	Sunday    DaySchedule = "Sunday"
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
		Enabled    bool       `json:"enabled,omitempty"`
		Threshold  int        `json:"threshold,omitempty"`
		Recurrence Recurrence `json:"recurrence,omitempty"`
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
		Enabled    bool `json:"enabled,omitempty"`
		Threshold  int  `json:"threshold,omitempty"`
		Recurrence struct {
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
