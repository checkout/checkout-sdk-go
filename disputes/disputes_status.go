package disputes

type DisputeStatus string

const (
	WON                    DisputeStatus = "won"
	LOST                   DisputeStatus = "lost"
	EXPIRED                DisputeStatus = "expired"
	ACCEPTED               DisputeStatus = "accepted"
	CANCELED               DisputeStatus = "canceled"
	RESOLVED               DisputeStatus = "resolved"
	ArbitrationWon         DisputeStatus = "arbitration_won"
	ArbitrationLost        DisputeStatus = "arbitration_lost"
	EvidenceRequired       DisputeStatus = "evidence_required"
	EvidenceUnderReview    DisputeStatus = "evidence_under_review"
	ArbitrationUnderReview DisputeStatus = "arbitration_under_review"
)
