package issuing

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/errors"
)

type (
	// CardControlData holds the fields shared by every control response
	// (create, get, list and update control).
	CardControlData struct {
		// HttpMetadata contains the HTTP status code and headers of the response.
		HttpMetadata common.HttpMetadata
		// ControlType is the control's type: velocity_limit, mcc_limit or mid_limit.
		// It selects which limit is populated in CardControlResponse.Limit.
		ControlType ControlType `json:"control_type,omitempty"`
		// Id is the control's unique identifier (prefix ctr_).
		Id string `json:"id,omitempty"`
		// Description is the description of the control. Maximum 256 characters.
		Description string `json:"description,omitempty"`
		// TargetId is the ID of the card or control profile the control applies to.
		TargetId string `json:"target_id,omitempty"`
		// IsEditable indicates whether you can change this control. false means an immutable
		// control applied by Checkout.com; true means you applied this control and can change it.
		IsEditable bool `json:"is_editable,omitempty"`
		// CreatedDate is the date and time, in UTC, when the control was created.
		CreatedDate *time.Time `json:"created_date,omitempty"`
		// LastModifiedDate is the date and time, in UTC, when the control was last modified.
		LastModifiedDate *time.Time `json:"last_modified_date,omitempty"`
		// Links contains the links related to the control, when returned.
		Links map[string]common.Link `json:"_links,omitempty"`
	}

	// CardControlResponse is a control returned by the API, discriminated on ControlType.
	CardControlResponse struct {
		CardControlData
		// Limit holds the control's limit: a VelocityLimit value for velocity_limit,
		// a MccLimit value for mcc_limit or a MidLimit value for mid_limit.
		Limit CardLimit `json:"limit,omitempty"`
	}

	// CardLimit is implemented by every control limit type.
	CardLimit interface {
		// GetType returns the ControlType that corresponds to the limit.
		GetType() ControlType
	}
)

func (l VelocityLimit) GetType() ControlType {
	return VelocityLimitType
}
func (l MccLimit) GetType() ControlType {
	return MccLimitType
}
func (l MidLimit) GetType() ControlType {
	return MidLimitType
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
	case MidLimitType:
		var limit = struct {
			MidLimit MidLimit `json:"mid_limit,omitempty"`
		}{}
		if err := json.Unmarshal(data, &limit); err != nil {
			return nil
		}
		s.Limit = limit.MidLimit
	default:
		return errors.UnsupportedTypeError(fmt.Sprintf("%s unsupported", controlData.ControlType))
	}
	return nil
}

type (
	// CardControlsQueryResponse is the list of controls applied to a target.
	CardControlsQueryResponse struct {
		// HttpMetadata contains the HTTP status code and headers of the response.
		HttpMetadata common.HttpMetadata
		// Controls is the list of controls applied to the specified target.
		Controls []CardControlResponse `json:"controls,omitempty"`
	}
)
