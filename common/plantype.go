package common

type PlanType string

const (
	MerchantInitiatedBilling                PlanType = "MERCHANT_INITIATED_BILLING"
	MerchantInitiatedBillingSingleAgreement PlanType = "MERCHANT_INITIATED_BILLING_SINGLE_AGREEMENT"
	ChannelInitiatedBilling                 PlanType = "CHANNEL_INITIATED_BILLING"
	ChannelInitiatedBillingSingleAgreement  PlanType = "CHANNEL_INITIATED_BILLING_SINGLE_AGREEMENT"
	RecurringPayments                       PlanType = "RECURRING_PAYMENTS"
	PreApprovedPayments                     PlanType = "PRE_APPROVED_PAYMENTS"
)
