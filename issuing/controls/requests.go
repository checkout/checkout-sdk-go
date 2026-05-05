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
		MccLimit MccLimit `json:"mcc_limit,omitempty"`
	}

	midCardControlRequest struct {
		CardControlCommon
		MidLimit MidLimit `json:"mid_limit,omitempty"`
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

func NewMidCardControlRequest() *midCardControlRequest {
	return &midCardControlRequest{
		CardControlCommon: CardControlCommon{ControlType: MidLimitType},
	}
}

func (c *velocityCardControlRequest) GetControlType() ControlType {
	return c.ControlType
}

func (c *mccCardControlRequest) GetControlType() ControlType {
	return c.ControlType
}

func (c *midCardControlRequest) GetControlType() ControlType {
	return c.ControlType
}

type (
	CardControlsQuery struct {
		TargetId string `url:"target_id,omitempty"`
	}
)

type (
	UpdateCardControlRequest struct {
		Description   string         `json:"description,omitempty"`
		VelocityLimit *VelocityLimit `json:"velocity_limit,omitempty"`
		MccLimit      *MccLimit      `json:"mcc_limit,omitempty"`
		MidLimit      *MidLimit      `json:"mid_limit,omitempty"`
	}
)
