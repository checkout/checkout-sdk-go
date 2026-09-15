package configuration

// OAuth 2.0 client credentials scopes.
//
// Mirrors components.securitySchemes.OAuth.flows.clientCredentials.scopes in the Checkout.com API
// specification, plus the scopes that appear only in per-operation security requirements and are
// never declared in that map: compliance-requests, compliance-requests:read,
// compliance-requests:respond, vault:gpayme-enrollment and vault:tokens-metadata.
//
// Five further constants -- issuing:card-mgmt, issuing:client, marketplace, middleware:gateway
// and middleware:payment-context -- appear nowhere in the specification at all, but the
// authorization server still grants them and callers still request them, so they are kept for
// backward compatibility. Each is marked inline. Do not assume a scope is dead because the
// specification omits it: the sandbox payouts client is provisioned for marketplace and answers
// a request for accounts with invalid_scope.
//
// Constants are ordered alphabetically. Note that PaymentContext and GatewayPaymentContexts are
// different scopes: the specification requires the former for GET /payment-contexts/{id} and the
// latter for POST /payment-contexts. "Payment Context" is the only scope whose wire value contains
// a space and a capital letter, which looks like a specification authoring defect; it is mirrored
// verbatim regardless, because that is the value the authorization server is documented to accept.
const (
	Accounts                    = "accounts"
	AgenticInventory            = "agentic:inventory"
	Balances                    = "balances"
	BalancesTopUpInstructions   = "balances:top-up-instructions"
	BalancesView                = "balances:view"
	CardManagement              = "card-management"
	ComplianceRequests          = "compliance-requests"
	ComplianceRequestsRead      = "compliance-requests:read"
	ComplianceRequestsRespond   = "compliance-requests:respond"
	Disputes                    = "disputes"
	DisputesAccept              = "disputes:accept"
	DisputesProvideEvidence     = "disputes:provide-evidence"
	DisputesSchemeFiles         = "disputes:scheme-files"
	DisputesView                = "disputes:view"
	Files                       = "files"
	FilesDownload               = "files:download"
	FilesRetrieve               = "files:retrieve"
	FilesUpload                 = "files:upload"
	FinancialActions            = "financial-actions"
	FinancialActionsView        = "financial-actions:view"
	Flow                        = "flow"
	FlowEvents                  = "flow:events"
	FlowReflow                  = "flow:reflow"
	FlowWorkflows               = "flow:workflows"
	Forward                     = "forward"
	ForwardSecrets              = "forward:secrets"
	Fx                          = "fx"
	Gateway                     = "gateway"
	GatewayPayment              = "gateway:payment"
	GatewayPaymentAuthorization = "gateway:payment-authorizations"
	GatewayPaymentCancellations = "gateway:payment-cancellations"
	GatewayPaymentCaptures      = "gateway:payment-captures"
	GatewayPaymentContexts      = "gateway:payment-contexts"
	GatewayPaymentDetails       = "gateway:payment-details"
	GatewayPaymentRefunds       = "gateway:payment-refunds"
	GatewayPaymentVoids         = "gateway:payment-voids"
	IdentityVerification        = "identity-verification"
	IssuingCardManagementRead   = "issuing:card-management-read"
	IssuingCardManagementWrite  = "issuing:card-management-write"
	IssuingCardMgmt             = "issuing:card-mgmt" // not in spec; kept for backward compat
	IssuingClient               = "issuing:client"    // not in spec; kept for backward compat
	IssuingControlsRead         = "issuing:controls-read"
	IssuingControlsWrite        = "issuing:controls-write"
	IssuingDisputes             = "issuing-disputes"
	IssuingDisputesRead         = "issuing:disputes-read"
	IssuingDisputesWrite        = "issuing:disputes-write"
	IssuingTransactionsRead     = "issuing:transactions-read"
	IssuingTransactionsWrite    = "issuing:transactions-write"
	Marketplace                 = "marketplace" // not in spec; kept for backward compat
	Middleware                  = "middleware"
	MiddlewareGateway           = "middleware:gateway" // not in spec; kept for backward compat
	MiddlewareMerchantsPublic   = "middleware:merchants-public"
	MiddlewareMerchantsSecret   = "middleware:merchants-secret"
	MiddlewarePaymentContext    = "middleware:payment-context" // not in spec; kept for backward compat
	PaymentContext              = "Payment Context"
	PaymentSessions             = "payment-sessions"
	PaymentsSearch              = "payments:search"
	PayoutsBankDetails          = "payouts:bank-details"
	Reports                     = "reports"
	ReportsView                 = "reports:view"
	SessionsApp                 = "sessions:app"
	SessionsBrowser             = "sessions:browser"
	Transactions                = "transactions"
	Transfers                   = "transfers"
	TransfersCreate             = "transfers:create"
	TransfersView               = "transfers:view"
	Vault                       = "vault"
	VaultApmeEnrollment         = "vault:apme-enrollment"
	VaultCardMetadata           = "vault:card-metadata"
	VaultCustomers              = "vault:customers"
	VaultGpaymeEnrollment       = "vault:gpayme-enrollment"
	VaultInstruments            = "vault:instruments"
	VaultNetworkTokens          = "vault:network-tokens"
	VaultRealTimeAccountUpdater = "vault:real-time-account-updater"
	VaultTokenization           = "vault:tokenization"
	VaultTokensMetadata         = "vault:tokens-metadata"
)
