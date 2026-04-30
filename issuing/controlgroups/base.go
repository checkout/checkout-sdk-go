package issuing

type ControlGroupControlType string

const (
	VelocityLimitControlType ControlGroupControlType = "velocity_limit"
	MccLimitControlType      ControlGroupControlType = "mcc_limit"
	MidLimitControlType      ControlGroupControlType = "mid_limit"
)

type FailIfType string

const (
	AllFail FailIfType = "all_fail"
	AnyFail FailIfType = "any_fail"
)

type LimitControlType string

const (
	Allow LimitControlType = "allow"
	Block LimitControlType = "block"
)

type VelocityWindowType string

const (
	Daily   VelocityWindowType = "daily"
	Weekly  VelocityWindowType = "weekly"
	Monthly VelocityWindowType = "monthly"
	AllTime VelocityWindowType = "all_time"
)

type (
	VelocityWindow struct {
		Type VelocityWindowType `json:"type,omitempty"`
	}

	VelocityGroupLimit struct {
		AmountLimit    int64          `json:"amount_limit,omitempty"`
		VelocityWindow VelocityWindow `json:"velocity_window"`
		MccList        []string       `json:"mcc_list,omitempty"`
		MidList        []string       `json:"mid_list,omitempty"`
	}

	MccGroupLimit struct {
		Type    LimitControlType `json:"type,omitempty"`
		MccList []string         `json:"mcc_list,omitempty"`
	}

	MidGroupLimit struct {
		Type    LimitControlType `json:"type,omitempty"`
		MidList []string         `json:"mid_list,omitempty"`
	}

	ControlGroupControl struct {
		ControlType   ControlGroupControlType `json:"control_type,omitempty"`
		Description   string                  `json:"description,omitempty"`
		MccLimit      *MccGroupLimit          `json:"mcc_limit,omitempty"`
		MidLimit      *MidGroupLimit          `json:"mid_limit,omitempty"`
		VelocityLimit *VelocityGroupLimit     `json:"velocity_limit,omitempty"`
	}
)
