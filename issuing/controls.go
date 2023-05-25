package issuing

import (
	"time"
)

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
		Type VelocityWindowType
	}

	VelocityLimit struct {
		AmountLimit    int
		VelocityWindow VelocityWindow
		MccList        []string
	}

	MccLimit struct {
		Type    MccControlType
		MccList []string
	}

	CardControlResponse interface {
		GetType() ControlType
	}

	CardControlTypeResponse struct {
		CardControlResponse
	}

	CardControlDetails struct {
		ControlType      ControlType `json:"control_type,omitempty"`
		Id               string      `json:"id,omitempty"`
		Description      string      `json:"description,omitempty"`
		TargetId         string      `json:"target_id,omitempty"`
		CreatedDate      *time.Time  `json:"created_date,omitempty"`
		LastModifiedDate *time.Time  `json:"last_modified_date,omitempty"`
	}

	velocityLimitCardControlTypeResponse struct {
		CardControlDetails
		VelocityLimit VelocityLimit `json:"velocity_limit,omitempty"`
	}

	mccLimitCardControlTypeResponse struct {
		CardControlDetails
		MccLimit MccLimit `json:"mcc_limit,omitempty"`
	}
)

func NewVelocityLimitCardControlTypeResponse() *velocityLimitCardControlTypeResponse {
	return &velocityLimitCardControlTypeResponse{
		CardControlDetails: CardControlDetails{ControlType: VelocityLimitType},
	}
}

func NewMccLimitCardControlTypeResponse() *mccLimitCardControlTypeResponse {
	return &mccLimitCardControlTypeResponse{
		CardControlDetails: CardControlDetails{ControlType: MccLimitType},
	}
}

func (c *velocityLimitCardControlTypeResponse) GetType() ControlType {
	return c.ControlType
}

func (c *mccLimitCardControlTypeResponse) GetType() ControlType {
	return c.ControlType
}
