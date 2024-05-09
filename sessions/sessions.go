package sessions

const (
	SessionsPath          = "sessions"
	CollectDataPath       = "collect-data"
	CompletePath          = "complete"
	IssuerFingerprintPath = "issuer-fingerprint"
)

type AuthenticationType string

const (
	RegularAuthType      AuthenticationType = "regular"
	RecurringAuthType    AuthenticationType = "recurring"
	InstallmentAuthType  AuthenticationType = "installment"
	MaintainCardAuthType AuthenticationType = "maintain_card"
	AddCardAuthType      AuthenticationType = "add_card"
)

type Category string

const (
	Payment    Category = "payment"
	NonPayment Category = "nonPayment"
)

type TransactionType string

const (
	GoodsService             TransactionType = "goods_service"
	CheckAcceptance          TransactionType = "check_acceptance"
	AccountFunding           TransactionType = "account_funding"
	QuashiCardTransaction    TransactionType = "quashi_card_transaction"
	PrepaidActivationAndLoad TransactionType = "prepaid_activation_and_load"
)

type SessionStatus string

const (
	Pending            SessionStatus = "pending"
	Processing         SessionStatus = "processing"
	Challenged         SessionStatus = "challenged"
	ChallengeAbandoned SessionStatus = "challenge_abandoned"
	Expired            SessionStatus = "expired"
	Approved           SessionStatus = "approved"
	Attempted          SessionStatus = "attempted"
	Unavailable        SessionStatus = "unavailable"
	Declined           SessionStatus = "declined"
	Rejected           SessionStatus = "rejected"
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
	CollectChannelData  NextAction = "collect_channel_data"
	IssueFingerprint    NextAction = "issuer_fingerprint"
	ChallengeCardHolder NextAction = "challenge_cardholder"
	RedirectCardholder  NextAction = "redirect_cardholder"
	Complete            NextAction = "complete"
	Authenticate        NextAction = "authenticate"
)

type Recurring struct {
	DaysBetweenPayments int    `json:"days_between_payments,omitempty"`
	Expiry              string `json:"expiry,omitempty"`
}

type Installment struct {
	NumberOfPayments    int    `json:"number_of_payments,omitempty"`
	DaysBetweenPayments int    `json:"days_between_payments,omitempty"`
	Expiry              string `json:"expiry,omitempty"`
}
