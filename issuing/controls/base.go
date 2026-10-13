package issuing

type ControlType string

const (
	VelocityLimitType ControlType = "velocity_limit"
	MccLimitType      ControlType = "mcc_limit"
	MidLimitType      ControlType = "mid_limit"
)

type VelocityWindowType string

const (
	Daily   VelocityWindowType = "daily"
	Weekly  VelocityWindowType = "weekly"
	Monthly VelocityWindowType = "monthly"
	AllTime VelocityWindowType = "all_time"
)

type MccControlType string

const (
	Allow MccControlType = "allow"
	Block MccControlType = "block"
)

type (
	// VelocityWindow is the period of time over which the velocity limit amount can be spent.
	VelocityWindow struct {
		// Type is the velocity window's unit of time. One of daily, weekly, monthly or all_time.
		Type VelocityWindowType `json:"type,omitempty"`
	}

	// VelocityLimit determines how much a target card can spend over a given timeframe.
	// It is used both in control requests and in control responses.
	VelocityLimit struct {
		// AmountLimit is the amount that can be spent, in minor units. Minimum 0.
		AmountLimit int64 `json:"amount_limit,omitempty"`
		// AmountRemaining is the remaining amount that can be spent, in minor units. Minimum 0.
		// Response only: it is returned by the API and is never sent on requests while left nil.
		AmountRemaining *int64 `json:"amount_remaining,omitempty"`
		// VelocityWindow is the period of time over which the specified AmountLimit can be spent.
		VelocityWindow VelocityWindow `json:"velocity_window"`
		// MccList is the list of merchant category codes (MCCs) that the velocity limit applies to,
		// as four-digit ISO 18245 codes.
		MccList []string `json:"mcc_list,omitempty"`
		// MidList is the list of merchant identification (MID) codes to allow or block transactions from.
		// You can provide either MccList or MidList, but not both.
		MidList []string `json:"mid_list,omitempty"`
	}

	// MccLimit is the merchant category code (MCC) rule, which determines the types of businesses
	// transactions can be processed from.
	MccLimit struct {
		// Type sets whether to allow or block the list of MCCs supplied. One of allow or block.
		Type MccControlType `json:"type,omitempty"`
		// MccList is the list of MCCs to allow or block transactions from, as 4-digit ISO 18245 codes.
		MccList []string `json:"mcc_list,omitempty"`
	}

	// MidLimit is the merchant identification (MID) code rule, which determines the merchants
	// from whom transactions can be processed.
	MidLimit struct {
		// Type sets whether to allow or block the list of MIDs supplied. One of allow or block.
		Type MccControlType `json:"type,omitempty"`
		// MidList is the list of merchant identification (MID) codes to allow or block transactions from.
		// Each code is 1 to 15 characters long.
		MidList []string `json:"mid_list,omitempty"`
	}
)
