package common

// Plan  - The plan details for the recurring payment agreement.
type Plan struct {
	Type                     *PlanType `json:"type" binding:"required"`
	SkipShippingAddress      *bool     `json:"skip_shipping_address,omitempty"`
	ImmutableShippingAddress *bool     `json:"immutable_shipping_address,omitempty"`
}
