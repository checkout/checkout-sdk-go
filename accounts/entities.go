package accounts

import "github.com/checkout/checkout-sdk-go-beta/common"

type (
	OnboardEntityRequest struct {
		Reference      string          `json:"reference,omitempty"`
		ContactDetails *ContactDetails `json:"contact_details,omitempty"`
		Profile        *Profile        `json:"profile,omitempty"`
		Company        *Company        `json:"company,omitempty"`
		Individual     *Individual     `json:"individual,omitempty"`
	}

	OnboardEntityResponse struct {
		HttpMetadata    common.HttpMetadata `json:"http_metadata,omitempty"`
		Id              string              `json:"id,omitempty"`
		Reference       string              `json:"reference,omitempty"`
		Capabilities    *Capabilities       `json:"capabilities,omitempty"`
		Status          OnboardingStatus    `json:"status,omitempty"`
		RequirementsDue []RequirementsDue   `json:"requirements_due,omitempty"`
	}

	OnboardEntityDetails struct {
		HttpMetadata    common.HttpMetadata `json:"http_metadata,omitempty"`
		Id              string              `json:"id,omitempty"`
		Reference       string              `json:"reference,omitempty"`
		Capabilities    *Capabilities       `json:"capabilities,omitempty"`
		Status          OnboardingStatus    `json:"status,omitempty"`
		RequirementsDue []RequirementsDue   `json:"requirements_due,omitempty"`
		ContactDetails  *ContactDetails     `json:"contact_details,omitempty"`
		Profile         *Profile            `json:"profile,omitempty"`
		Company         *Company            `json:"company,omitempty"`
		Individual      *Individual         `json:"individual,omitempty"`
		Instruments     []Instrument        `json:"instruments,omitempty"`
	}
)
