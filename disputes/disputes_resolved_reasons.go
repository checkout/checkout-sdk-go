package disputes

type DisputeResolvedReason string

const (
	RapidDisputeResolution DisputeResolvedReason = "rapid_dispute_resolution"
	NegativeAmount         DisputeResolvedReason = "negative_amount"
	AlreadyRefunded        DisputeResolvedReason = "already_refunded"
)
