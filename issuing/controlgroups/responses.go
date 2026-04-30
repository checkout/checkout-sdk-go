package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type (
	ControlGroupResponse struct {
		HttpMetadata     common.HttpMetadata
		Id               string                 `json:"id,omitempty"`
		TargetId         string                 `json:"target_id,omitempty"`
		FailIf           FailIfType             `json:"fail_if,omitempty"`
		Controls         []ControlGroupControl  `json:"controls,omitempty"`
		IsEditable       *bool                  `json:"is_editable,omitempty"`
		Description      string                 `json:"description,omitempty"`
		CreatedDate      *time.Time             `json:"created_date,omitempty"`
		LastModifiedDate *time.Time             `json:"last_modified_date,omitempty"`
		Links            map[string]common.Link `json:"_links,omitempty"`
	}

	ControlGroupsResponse struct {
		HttpMetadata  common.HttpMetadata
		ControlGroups []ControlGroupResponse `json:"control_groups,omitempty"`
	}
)
