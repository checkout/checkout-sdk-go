package sessions

import "time"

const (
	SessionsPath          = "sessions"
	CollectDataPath       = "collect-data"
	CompletePath          = "complete"
	IssuerFingerprintPath = "issuer-fingerprint"
)

type AuthenticationType string

const (
	AddCardAuthType      AuthenticationType = "add_card"
	InstallmentAuthType  AuthenticationType = "installment"
	MaintainCardAuthType AuthenticationType = "maintain_card"
	RecurringAuthType    AuthenticationType = "recurring"
	RegularAuthType      AuthenticationType = "regular"
)

type Category string

const (
	Payment    Category = "payment"
	NonPayment Category = "nonPayment"
)

type TransactionType string

const (
	AccountFunding           TransactionType = "account_funding"
	CheckAcceptance          TransactionType = "check_acceptance"
	GoodsService             TransactionType = "goods_service"
	PrepaidActivationAndLoad TransactionType = "prepaid_activation_and_load"
	QuashiCardTransaction    TransactionType = "quashi_card_transaction"
)

type SessionStatus string

const (
	Approved           SessionStatus = "approved"
	Attempted          SessionStatus = "attempted"
	Challenged         SessionStatus = "challenged"
	ChallengeAbandoned SessionStatus = "challenge_abandoned"
	Declined           SessionStatus = "declined"
	Expired            SessionStatus = "expired"
	Pending            SessionStatus = "pending"
	Processing         SessionStatus = "processing"
	Rejected           SessionStatus = "rejected"
	Unavailable        SessionStatus = "unavailable"
)

type StatusReason string

const (
	AresError    StatusReason = "ares_error"
	AresStatus   StatusReason = "ares_status"
	VeresError   StatusReason = "veres_error"
	VeresStatus  StatusReason = "veres_status"
	ParesError   StatusReason = "pares_error"
	ParesStatus  StatusReason = "pares_status"
	RreqError    StatusReason = "rreq_error"
	RreqStatus   StatusReason = "rreq_status"
	RiskDeclined StatusReason = "risk_declined"
)

type NextAction string

const (
	Authenticate        NextAction = "authenticate"
	ChallengeCardHolder NextAction = "challenge_cardholder"
	CollectChannelData  NextAction = "collect_channel_data"
	Complete            NextAction = "complete"
	IssueFingerprint    NextAction = "issuer_fingerprint"
	RedirectCardholder  NextAction = "redirect_cardholder"
)

type DeliveryTimeframe string

const (
	ElectronicDelivery DeliveryTimeframe = "electronic_delivery"
	SameDay            DeliveryTimeframe = "same_day"
	Overnight          DeliveryTimeframe = "overnight"
	TwoDayOrMore       DeliveryTimeframe = "two_day_or_more"
)

type PreOrderPurchaseIndicatorType string

const (
	FutureAvailability   PreOrderPurchaseIndicatorType = "future_availability"
	MerchandiseAvailable PreOrderPurchaseIndicatorType = "merchandise_available"
)

type ReorderItemsIndicatorType string

const (
	FirstTimeOrdered ReorderItemsIndicatorType = "first_time_ordered"
	Reordered        ReorderItemsIndicatorType = "reordered"
)

type Recurring struct {
	DaysBetweenPayments int    `json:"days_between_payments,omitempty" default:"1"`
	Expiry              string `json:"expiry,omitempty" default:"99991231"`
}

type Installment struct {
	NumberOfPayments    int    `json:"number_of_payments,omitempty"`
	DaysBetweenPayments int    `json:"days_between_payments,omitempty" default:"1"`
	Expiry              string `json:"expiry,omitempty" default:"99991231"`
}

type MerchantRiskInfo struct {
	DeliveryEmail             string                        `json:"delivery_email,omitempty"`
	DeliveryTimeframe         DeliveryTimeframe             `json:"delivery_timeframe,omitempty"`
	IsPreorder                bool                          `json:"is_preorder,omitempty"`
	IsReorder                 bool                          `json:"is_reorder,omitempty"`
	ShippingIndicator         ShippingIndicator             `json:"shipping_indicator,omitempty"`
	ReorderItemsIndicator     ReorderItemsIndicatorType     `json:"reorder_items_indicator,omitempty"`
	PreOrderPurchaseIndicator PreOrderPurchaseIndicatorType `json:"pre_order_purchase_indicator,omitempty"`
	PreOrderDate              *time.Time                    `json:"pre_order_date,omitempty"`
	GiftCardAmount            string                        `json:"gift_card_amount,omitempty"`
	GiftCardCurrency          string                        `json:"gift_card_currency,omitempty"`
	GiftCardCount             string                        `json:"gift_card_count,omitempty"`
}
