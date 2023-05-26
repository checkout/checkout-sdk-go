package issuing

type (
	CardControlRequest interface {
		GetControlType() ControlType
	}

	CardControlTypeRequest struct {
		CardControlRequest
	}

	CardControlCommon struct {
		ControlType ControlType `json:"control_type,omitempty"`
		Description string      `json:"description,omitempty"`
		TargetId    string      `json:"target_id,omitempty"`
	}

	velocityCardControlRequest struct {
		CardControlCommon
		VelocityLimit VelocityLimit `json:"velocity_limit,omitempty"`
	}

	mccCardControlRequest struct {
		CardControlCommon
		MccLimit MccLimit
	}
)

func NewVelocityCardControlRequest() *velocityCardControlRequest {
	return &velocityCardControlRequest{
		CardControlCommon: CardControlCommon{ControlType: VelocityLimitType},
	}
}

func NewMccCardControlRequest() *mccCardControlRequest {
	return &mccCardControlRequest{
		CardControlCommon: CardControlCommon{ControlType: MccLimitType},
	}
}

func (c *velocityCardControlRequest) GetControlType() ControlType {
	return c.ControlType
}

func (c *mccCardControlRequest) GetControlType() ControlType {
	return c.ControlType
}

type (
	CardControlsQuery struct {
		TargetId string `json:"target_id,omitempty"`
	}
)

type (
	UpdateCardControlRequest struct {
		Description   string        `json:"description,omitempty"`
		VelocityLimit VelocityLimit `json:"velocity_limit"`
		MccLimit      MccLimit      `json:"mcc_limit"`
	}
)
