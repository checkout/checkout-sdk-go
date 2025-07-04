package configuration

const (
	Vault                       = "vault"
	VaultInstruments            = "vault:instruments"
	VaultTokenization           = "vault:tokenization"
	Gateway                     = "gateway"
	GatewayPayment              = "gateway:payment"
	GatewayPaymentDetails       = "gateway:payment-details"
	GatewayPaymentAuthorization = "gateway:payment-authorizations"
	GatewayPaymentVoids         = "gateway:payment-voids"
	GatewayPaymentCaptures      = "gateway:payment-captures"
	GatewayPaymentRefunds       = "gateway:payment-refunds"
	Fx                          = "fx"
	PayoutsBankDetails          = "payouts:bank-details"
	SessionsApp                 = "sessions:app"
	SessionsBrowser             = "sessions:browser"
	Disputes                    = "disputes"
	DisputesView                = "disputes:view"
	DisputesProvideEvidence     = "disputes:provide-evidence"
	DisputesAccept              = "disputes:accept"
	DisputesSchemeFiles         = "disputes:scheme-files"
	Marketplace                 = "marketplace"
	Accounts                    = "accounts"
	Flow                        = "flow"
	FlowWorkflows               = "flow:workflows"
	FlowEvents                  = "flow:events"
	Files                       = "files"
	FilesRetrieve               = "files:retrieve"
	FilesUpload                 = "files:upload"
	FilesDownload               = "files:download"
	Transfers                   = "transfers"
	TransfersCreate             = "transfers:create"
	TransfersView               = "transfers:view"
	Balances                    = "balances"
	BalancesView                = "balances:view"
	Middleware                  = "middleware"
	MiddlewareGateway           = "middleware:gateway"
	MiddlewarePaymentContext    = "middleware:payment-context"
	MiddlewareMerchantsSecret   = "middleware:merchants-secret"
	MiddlewareMerchantsPublic   = "middleware:merchants-public"
	Reports                     = "reports"
	ReportsView                 = "reports:view"
	VaultCardMetadata           = "vault:card-metadata"
	FinancialActions            = "financial-actions"
	FinancialActionsView        = "financial-actions:view"
	IssuingClient               = "issuing:client"
	IssuingCardMgmt             = "issuing:card-mgmt"
	IssuingControlsRead         = "issuing:controls-read"
	IssuingControlsWrite        = "issuing:controls-write"
	PaymentContexts             = "Payment Contexts"
	Forward                     = "forward"
)
