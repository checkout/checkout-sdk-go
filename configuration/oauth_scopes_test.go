package configuration

import (
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
