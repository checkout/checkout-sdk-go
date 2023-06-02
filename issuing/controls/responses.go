package issuing

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/errors"
)

type (
	CardControlData struct {
		HttpMetadata     common.HttpMetadata
		ControlType      ControlType `json:"control_type,omitempty"`
		Id               string      `json:"id,omitempty"`
		Description      string      `json:"description,omitempty"`
		TargetId         string      `json:"target_id,omitempty"`
		CreatedDate      *time.Time  `json:"created_date,omitempty"`
		LastModifiedDate *time.Time  `json:"last_modified_date,omitempty"`
	}

	CardControlResponse struct {
		CardControlData
		Limit CardLimit `json:"limit,omitempty"`
	}

	CardLimit interface {
		GetType() ControlType
	}
)

func (l VelocityLimit) GetType() ControlType {
	return VelocityLimitType
}
func (l MccLimit) GetType() ControlType {
	return MccLimitType
}

func (s *CardControlResponse) UnmarshalJSON(data []byte) error {
	var controlData CardControlData
	if err := json.Unmarshal(data, &controlData); err != nil {
		return err
	}
	s.CardControlData = controlData

	switch controlData.ControlType {
	case VelocityLimitType:
		var limit = struct {
			VelocityLimit VelocityLimit `json:"velocity_limit,omitempty"`
		}{}
		if err := json.Unmarshal(data, &limit); err != nil {
			return nil
		}
		s.Limit = limit.VelocityLimit
	case MccLimitType:
		var limit = struct {
			MccLimit MccLimit `json:"mcc_limit,omitempty"`
		}{}
		if err := json.Unmarshal(data, &limit); err != nil {
			return nil
		}
		s.Limit = limit.MccLimit
	default:
		return errors.UnsupportedTypeError(fmt.Sprintf("%s unsupported", controlData.ControlType))
	}
	return nil
}

type (
	CardControlsQueryResponse struct {
		HttpMetadata common.HttpMetadata
		Controls     []CardControlResponse `json:"controls,omitempty"`
	}
)
