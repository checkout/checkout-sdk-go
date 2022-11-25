package disputes

type DisputeCategory string

const (
	General                   DisputeCategory = "general"
	Duplicate                 DisputeCategory = "duplicate"
	Fraudulent                DisputeCategory = "fraudulent"
	Unrecognized              DisputeCategory = "unrecognized"
	IncorrectAmount           DisputeCategory = "incorrect_amount"
	NotAsDescribed            DisputeCategory = "not_as_described"
	CreditNotIssued           DisputeCategory = "credit_not_issued"
	CanceledRecurring         DisputeCategory = "canceled_recurring"
	ProductServiceNotReceived DisputeCategory = "product_service_not_received"
)
