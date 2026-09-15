package configuration

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The balances scopes are the only place the wire value of an OAuth scope is written down, and
// nothing else in the suite exercises BalancesTopUpInstructions: the sandbox fixture's
// getOAuthScopes() does not request it, so a typo would ship silently and every OAuth-configured
// caller of RetrieveTopUpInstructions would be rejected at the token endpoint.
//
// Values are taken from components.securitySchemes.OAuth.flows.clientCredentials.scopes in
// shared/swagger-latest.json.
func TestBalancesOAuthScopeValues(t *testing.T) {
	cases := []struct {
		name     string
		scope    string
		expected string
	}{
		{"Balances", Balances, "balances"},
		{"BalancesView", BalancesView, "balances:view"},
		{"BalancesTopUpInstructions", BalancesTopUpInstructions, "balances:top-up-instructions"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.scope)
		})
	}
}

// The scopes added when these constants were synced against the spec.
//
// The last five are not declared in clientCredentials.scopes at all: they appear only in the
// per-operation security requirements of GET/POST /compliance-requests/{payment_id}, the
// /googlepay/enrollments operations and GET /tokens/{tokenId}/metadata. Constants derived from the
// declared map alone would be missing them.
func TestOAuthScopeValuesAddedInSpecSync(t *testing.T) {
	cases := []struct {
		name     string
		scope    string
		expected string
	}{
		{"AgenticInventory", AgenticInventory, "agentic:inventory"},
		{"CardManagement", CardManagement, "card-management"},
		{"FlowReflow", FlowReflow, "flow:reflow"},
		{"GatewayPaymentContexts", GatewayPaymentContexts, "gateway:payment-contexts"},
		{"IssuingCardManagementRead", IssuingCardManagementRead, "issuing:card-management-read"},
		{"IssuingCardManagementWrite", IssuingCardManagementWrite, "issuing:card-management-write"},
		{"IssuingDisputes", IssuingDisputes, "issuing-disputes"},
		{"IssuingTransactionsWrite", IssuingTransactionsWrite, "issuing:transactions-write"},
		{"PaymentSessions", PaymentSessions, "payment-sessions"},
		{"Transactions", Transactions, "transactions"},
		{"ComplianceRequests", ComplianceRequests, "compliance-requests"},
		{"ComplianceRequestsRead", ComplianceRequestsRead, "compliance-requests:read"},
		{"ComplianceRequestsRespond", ComplianceRequestsRespond, "compliance-requests:respond"},
		{"VaultGpaymeEnrollment", VaultGpaymeEnrollment, "vault:gpayme-enrollment"},
		{"VaultTokensMetadata", VaultTokensMetadata, "vault:tokens-metadata"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.scope)
		})
	}
}

// These five scopes appear nowhere in the specification -- not in the clientCredentials scope map
// and not in any operation's security requirement -- so a sweep driven by the spec alone would
// delete them. They are kept deliberately: the authorization server still grants them and callers
// still request them. marketplace is the proof: the sandbox client behind
// CHECKOUT_DEFAULT_OAUTH_PAYOUT_SCHEDULE_CLIENT_ID is provisioned for it and answers a request for
// accounts with invalid_scope, so dropping it broke TestSubmitFileAccounts (PR #254).
//
// This test exists to stop the next specification-driven tidy-up from removing them again.
func TestLegacyOAuthScopeValuesAreRetained(t *testing.T) {
	cases := []struct {
		name     string
		scope    string
		expected string
	}{
		{"IssuingCardMgmt", IssuingCardMgmt, "issuing:card-mgmt"},
		{"IssuingClient", IssuingClient, "issuing:client"},
		{"Marketplace", Marketplace, "marketplace"},
		{"MiddlewareGateway", MiddlewareGateway, "middleware:gateway"},
		{"MiddlewarePaymentContext", MiddlewarePaymentContext, "middleware:payment-context"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.scope)
		})
	}
}

// PaymentContext and GatewayPaymentContexts are unrelated scopes despite reading alike, so this
// pins which is which. It also pins the singular: this constant held "Payment Contexts" (plural)
// until the spec sync, a value the authorization server defines under no reading of the spec, so
// every OAuth caller that requested it was rejected -- and rejected for the whole token request,
// losing every other scope asked for alongside it.
//
// The space and capital letter in "Payment Context" are almost certainly a spec authoring defect,
// asserted verbatim because that is the value GET /payment-contexts/{id} documents.
func TestPaymentContextOAuthScopeValues(t *testing.T) {
	assert.Equal(t, "Payment Context", PaymentContext)
	assert.Equal(t, "gateway:payment-contexts", GatewayPaymentContexts)
}

// Go gives these constants the least protection of any SDK in this family: they are untyped string
// constants at package scope, so nothing stops a caller passing an arbitrary string to WithScopes,
// and nothing here can be enumerated at runtime the way an enum can. A blank or duplicated value
// therefore has to be caught by reading the file, which is what allScopes below exists for -- it is
// the only inventory of these constants that the compiler will check.
//
// Keep allScopes in step with the const block: a constant missing from it is simply not covered.
func allScopes() map[string]string {
	return map[string]string{
		"Accounts":                    Accounts,
		"AgenticInventory":            AgenticInventory,
		"Balances":                    Balances,
		"BalancesTopUpInstructions":   BalancesTopUpInstructions,
		"BalancesView":                BalancesView,
		"CardManagement":              CardManagement,
		"ComplianceRequests":          ComplianceRequests,
		"ComplianceRequestsRead":      ComplianceRequestsRead,
		"ComplianceRequestsRespond":   ComplianceRequestsRespond,
		"Disputes":                    Disputes,
		"DisputesAccept":              DisputesAccept,
		"DisputesProvideEvidence":     DisputesProvideEvidence,
		"DisputesSchemeFiles":         DisputesSchemeFiles,
		"DisputesView":                DisputesView,
		"Files":                       Files,
		"FilesDownload":               FilesDownload,
		"FilesRetrieve":               FilesRetrieve,
		"FilesUpload":                 FilesUpload,
		"FinancialActions":            FinancialActions,
		"FinancialActionsView":        FinancialActionsView,
		"Flow":                        Flow,
		"FlowEvents":                  FlowEvents,
		"FlowReflow":                  FlowReflow,
		"FlowWorkflows":               FlowWorkflows,
		"Forward":                     Forward,
		"ForwardSecrets":              ForwardSecrets,
		"Fx":                          Fx,
		"Gateway":                     Gateway,
		"GatewayPayment":              GatewayPayment,
		"GatewayPaymentAuthorization": GatewayPaymentAuthorization,
		"GatewayPaymentCancellations": GatewayPaymentCancellations,
		"GatewayPaymentCaptures":      GatewayPaymentCaptures,
		"GatewayPaymentContexts":      GatewayPaymentContexts,
		"GatewayPaymentDetails":       GatewayPaymentDetails,
		"GatewayPaymentRefunds":       GatewayPaymentRefunds,
		"GatewayPaymentVoids":         GatewayPaymentVoids,
		"IdentityVerification":        IdentityVerification,
		"IssuingCardManagementRead":   IssuingCardManagementRead,
		"IssuingCardManagementWrite":  IssuingCardManagementWrite,
		"IssuingCardMgmt":             IssuingCardMgmt,
		"IssuingClient":               IssuingClient,
		"IssuingControlsRead":         IssuingControlsRead,
		"IssuingControlsWrite":        IssuingControlsWrite,
		"IssuingDisputes":             IssuingDisputes,
		"IssuingDisputesRead":         IssuingDisputesRead,
		"IssuingDisputesWrite":        IssuingDisputesWrite,
		"IssuingTransactionsRead":     IssuingTransactionsRead,
		"IssuingTransactionsWrite":    IssuingTransactionsWrite,
		"Marketplace":                 Marketplace,
		"Middleware":                  Middleware,
		"MiddlewareGateway":           MiddlewareGateway,
		"MiddlewareMerchantsPublic":   MiddlewareMerchantsPublic,
		"MiddlewareMerchantsSecret":   MiddlewareMerchantsSecret,
		"MiddlewarePaymentContext":    MiddlewarePaymentContext,
		"PaymentContext":              PaymentContext,
		"PaymentSessions":             PaymentSessions,
		"PaymentsSearch":              PaymentsSearch,
		"PayoutsBankDetails":          PayoutsBankDetails,
		"Reports":                     Reports,
		"ReportsView":                 ReportsView,
		"SessionsApp":                 SessionsApp,
		"SessionsBrowser":             SessionsBrowser,
		"Transactions":                Transactions,
		"Transfers":                   Transfers,
		"TransfersCreate":             TransfersCreate,
		"TransfersView":               TransfersView,
		"Vault":                       Vault,
		"VaultApmeEnrollment":         VaultApmeEnrollment,
		"VaultCardMetadata":           VaultCardMetadata,
		"VaultCustomers":              VaultCustomers,
		"VaultGpaymeEnrollment":       VaultGpaymeEnrollment,
		"VaultInstruments":            VaultInstruments,
		"VaultNetworkTokens":          VaultNetworkTokens,
		"VaultRealTimeAccountUpdater": VaultRealTimeAccountUpdater,
		"VaultTokenization":           VaultTokenization,
		"VaultTokensMetadata":         VaultTokensMetadata,
	}
}

// allScopes is hand-maintained, so on its own it would silently stop covering a constant the moment
// someone added one and forgot to extend the map -- the exact mistake the sweeps exist to catch. Go
// cannot enumerate a package's constants at runtime, so this parses the declaration instead and
// holds the two in step. If it fails, add the reported constant to allScopes rather than editing
// this test.
func TestAllScopesCoversEveryDeclaredConstant(t *testing.T) {
	fileSet := token.NewFileSet()
	parsed, err := parser.ParseFile(fileSet, "oauth_scopes.go", nil, 0)
	assert.NoError(t, err)

	declared := make(map[string]string)
	for _, declaration := range parsed.Decls {
		genDecl, ok := declaration.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || len(valueSpec.Names) != 1 || len(valueSpec.Values) != 1 {
				continue
			}
			literal, ok := valueSpec.Values[0].(*ast.BasicLit)
			if !ok || literal.Kind != token.STRING {
				continue
			}
			value, err := strconv.Unquote(literal.Value)
			assert.NoError(t, err)
			declared[valueSpec.Names[0].Name] = value
		}
	}

	assert.NotEmpty(t, declared, "parsed no constants out of oauth_scopes.go")
	assert.Equal(t, declared, allScopes())
}

// A blank value is not caught by the assertions above, which only read the constants they name.
// oauth_keys_credentials.go:89 joins the requested scopes with a space, so a blank constant would be
// sent as an empty entry and the token endpoint would reject the whole request, costing the caller
// every other scope it asked for.
func TestEveryOAuthScopeHasANonBlankValue(t *testing.T) {
	for name, scope := range allScopes() {
		assert.NotEmpty(t, strings.TrimSpace(scope), "%s has a blank wire value", name)
	}
}

// Two constants sharing a wire value means one of them is a copy-paste error, and it cannot be
// caught by the per-scope assertions above, which only ever read the constant they name. The
// consequence is silent in both directions: a caller selecting the mistyped constant requests a
// scope it did not ask for, and the scope that constant was supposed to carry is left with no
// constant at all, so it becomes unreachable through this package.
func TestNoOAuthScopeValueIsReused(t *testing.T) {
	owners := make(map[string][]string)
	for name, scope := range allScopes() {
		owners[scope] = append(owners[scope], name)
	}

	for scope, names := range owners {
		assert.Len(t, names, 1, "wire value %q is used by more than one constant: %v", scope, names)
	}
}
