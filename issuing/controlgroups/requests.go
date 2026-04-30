package issuing

type (
	CreateControlGroupRequest struct {
		TargetId    string                `json:"target_id,omitempty"`
		FailIf      FailIfType            `json:"fail_if,omitempty"`
		Controls    []ControlGroupControl `json:"controls,omitempty"`
		Description string                `json:"description,omitempty"`
	}

	ControlGroupsQuery struct {
		TargetId string `url:"target_id,omitempty"`
	}
)
