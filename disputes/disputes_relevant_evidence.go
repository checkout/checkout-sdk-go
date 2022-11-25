package disputes

type RelevantEvidence string

const (
	ProofOfDeliveryOrService           RelevantEvidence = "proof_of_delivery_or_service"
	InvoiceOrReceipt                   RelevantEvidence = "invoice_or_receipt"
	InvoiceShowingDistinctTransactions RelevantEvidence = "invoice_showing_distinct_transactions"
	CustomerCommunication              RelevantEvidence = "customer_communication"
	RefundOrCancellationPolicy         RelevantEvidence = "refund_or_cancellation_policy"
	RecurringTransactionAgreement      RelevantEvidence = "recurring_transaction_agreement"
	AdditionalEvidence                 RelevantEvidence = "additional_evidence"
)
