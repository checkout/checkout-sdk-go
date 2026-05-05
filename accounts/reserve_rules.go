package accounts

import (
	"time"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

// ReserveRuleType identifies the reserve calculation model applied to an entity.
type ReserveRuleType string

const (
	// RollingType indicates a rolling reserve rule where a percentage of each transaction
	// is held for a defined period.
	RollingType ReserveRuleType = "rolling"
)

type (
	// HoldingDuration specifies the period for which funds are held under a rolling reserve rule.
	HoldingDuration struct {
		Weeks *int `json:"weeks,omitempty"`
	}

	// RollingReserveRule defines the parameters of a rolling reserve: the percentage withheld
	// and how long the funds are held before release.
	RollingReserveRule struct {
		Percentage      *float64         `json:"percentage,omitempty"`
		HoldingDuration *HoldingDuration `json:"holding_duration,omitempty"`
	}

	// ReserveRuleRequest is the request body for creating or updating a reserve rule on an entity.
	ReserveRuleRequest struct {
		Type      ReserveRuleType     `json:"type,omitempty"`
		Rolling   *RollingReserveRule `json:"rolling,omitempty"`
		ValidFrom *time.Time          `json:"valid_from,omitempty"`
	}

	// ReserveRuleResponse represents a single reserve rule as returned by the API.
	ReserveRuleResponse struct {
		HttpMetadata common.HttpMetadata
		Id           string                 `json:"id,omitempty"`
		Type         ReserveRuleType        `json:"type,omitempty"`
		Rolling      *RollingReserveRule    `json:"rolling,omitempty"`
		ValidFrom    *time.Time             `json:"valid_from,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}

	// ReserveRulesResponse is the paginated list response returned when querying all reserve
	// rules for an entity.
	ReserveRulesResponse struct {
		HttpMetadata common.HttpMetadata
		Data         []ReserveRuleResponse  `json:"data,omitempty"`
		Links        map[string]common.Link `json:"_links"`
	}
)
