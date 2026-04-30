package accounts

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

type ReserveRuleType string

const (
	RollingType ReserveRuleType = "rolling"
)

type (
	HoldingDuration struct {
		Weeks *int `json:"weeks,omitempty"`
	}

	RollingReserveRule struct {
		Percentage      *float64         `json:"percentage,omitempty"`
		HoldingDuration *HoldingDuration `json:"holding_duration,omitempty"`
	}

	ReserveRuleRequest struct {
		Type      string              `json:"type,omitempty"`
		Rolling   *RollingReserveRule `json:"rolling,omitempty"`
		ValidFrom *time.Time          `json:"valid_from,omitempty"`
	}

	ReserveRuleResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Type         string                 `json:"type,omitempty"`
		Rolling      *RollingReserveRule    `json:"rolling,omitempty"`
		ValidFrom    *time.Time             `json:"valid_from,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	ReserveRulesResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []ReserveRuleResponse  `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}
)
