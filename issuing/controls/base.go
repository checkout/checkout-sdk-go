package issuing

type ControlType string

const (
	VelocityLimitType ControlType = "velocity_limit"
	MccLimitType      ControlType = "mcc_limit"
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
	VelocityWindow struct {
		Type VelocityWindowType `json:"type,omitempty" :"type"`
	}

	VelocityLimit struct {
		AmountLimit    int            `json:"amount_limit,omitempty"`
		VelocityWindow VelocityWindow `json:"velocity_window"`
		MccList        []string       `json:"mcc_list,omitempty"`
	}

	MccLimit struct {
		Type    MccControlType `json:"type,omitempty"`
		MccList []string       `json:"mcc_list,omitempty"`
	}
)
