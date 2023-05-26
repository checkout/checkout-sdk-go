package issuing

import (
	"encoding/json"
	"fmt"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
	"time"
)

type (
	CardControlResponse struct {
		HttpMetadata                     common.HttpMetadata
		VelocityLimitCardControlResponse *VelocityLimitCardControlResponse
		MccLimitCardControlResponse      *MccLimitCardControlResponse
	}

	VelocityLimitCardControlResponse struct {
		ControlType      ControlType   `json:"control_type,omitempty"`
		Id               string        `json:"id,omitempty"`
		Description      string        `json:"description,omitempty"`
		TargetId         string        `json:"target_id,omitempty"`
		CreatedDate      *time.Time    `json:"created_date,omitempty"`
		LastModifiedDate *time.Time    `json:"last_modified_date,omitempty"`
		VelocityLimit    VelocityLimit `json:"velocity_limit,omitempty"`
	}

	MccLimitCardControlResponse struct {
		ControlType      ControlType `json:"control_type,omitempty"`
		Id               string      `json:"id,omitempty"`
		Description      string      `json:"description,omitempty"`
		TargetId         string      `json:"target_id,omitempty"`
		CreatedDate      *time.Time  `json:"created_date,omitempty"`
		LastModifiedDate *time.Time  `json:"last_modified_date,omitempty"`
		MccLimit         MccLimit    `json:"mcc_limit,omitempty"`
	}
)

func (s *CardControlResponse) UnmarshalJSON(data []byte) error {
	var typeMapping common.TypeMapping
	if err := json.Unmarshal(data, &typeMapping); err != nil {
		return err
	}

	switch typeMapping.ControlType {
	case string(VelocityLimitType):
		var response VelocityLimitCardControlResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.VelocityLimitCardControlResponse = &response
	case string(MccLimitType):
		var response MccLimitCardControlResponse
		if err := json.Unmarshal(data, &response); err != nil {
			return nil
		}
		s.MccLimitCardControlResponse = &response
	default:
		return errors.UnsupportedTypeError(fmt.Sprintf("%s unsupported", typeMapping.ControlType))
	}
	return nil
}

type (
	CardControlsQueryResponse struct {
		HttpMetadata common.HttpMetadata
		Controls     []CardControlResponse `json:"controls,omitempty"`
	}
)
