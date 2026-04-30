package issuing

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type (
	ControlProfileRequest struct {
		Name string `json:"name,omitempty"`
	}

	ControlProfilesQuery struct {
		TargetId string `url:"target_id,omitempty"`
	}
)

type (
	ControlProfileResponse struct {
		HttpMetadata     common.HttpMetadata
		Id               string                 `json:"id,omitempty"`
		Name             string                 `json:"name,omitempty"`
		CreatedDate      *time.Time             `json:"created_date,omitempty"`
		LastModifiedDate *time.Time             `json:"last_modified_date,omitempty"`
		Links            map[string]common.Link `json:"_links,omitempty"`
	}

	ControlProfilesResponse struct {
		HttpMetadata    common.HttpMetadata
		ControlProfiles []ControlProfileResponse `json:"control_profiles,omitempty"`
	}
)
