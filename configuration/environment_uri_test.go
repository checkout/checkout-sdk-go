package configuration

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/errors"
)

// Every SDK client builds its path with common.BuildPath, which prefixes each segment with "/".
// A base URI that also ends in "/" therefore produces a double slash in the request URL. Sandbox
// never had trailing slashes, so the sandbox-based test suite could not catch it; production
// carried them on the files, transfers and balances hosts.
func TestEnvironmentUrisHaveNoTrailingSlash(t *testing.T) {
	cases := []struct {
		name string
		env  *CheckoutEnv
	}{
		{"Sandbox", Sandbox()},
		{"Production", Production()},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			uris := map[string]string{
				"BaseUri":      tc.env.BaseUri(),
				"FilesUri":     tc.env.FilesUri(),
				"TransfersUri": tc.env.TransfersUri(),
				"BalancesUri":  tc.env.BalancesUri(),
				"ForwardUri":   tc.env.ForwardUri(),
				"IdentityUri":  tc.env.IdentityUri(),
			}
			for name, uri := range uris {
				assert.False(t, strings.HasSuffix(uri, "/"),
					"%s.%s must not end in \"/\" (got %q): paths already start with \"/\", so this yields a double slash",
					tc.name, name, uri)
			}
		})
	}
}

// Pins the exact URL the balances client produces in both environments.
func TestBalancesUriJoinsCleanlyWithABuiltPath(t *testing.T) {
	// common.BuildPath("balances", "ent_x") -> "/balances/ent_x". Inlined to avoid an import
	// cycle between configuration and common.
	const path = "/balances/ent_x"

	assert.Equal(t, "https://balances.sandbox.checkout.com/balances/ent_x", Sandbox().BalancesUri()+path)
	assert.Equal(t, "https://balances.checkout.com/balances/ent_x", Production().BalancesUri()+path)
}

// An environment URI that cannot be parsed returns a CheckoutArgumentError that does not repeat the
// URI.
func TestEnvironmentSubdomainWithAnInvalidUriDoesNotRepeatIt(t *testing.T) {
	cases := map[string]*CheckoutEnv{
		"authorization URI": NewEnvironment("https://api.example.com", "https://user:userinfo-value@access.example.com/%zz", "", "", "", "", "", false),
		"base URI":          NewEnvironment("https://user:userinfo-value@api.example.com/%zz", "https://access.example.com/connect/token", "", "", "", "", "", false),
	}
	for name, env := range cases {
		env := env
		t.Run(name, func(t *testing.T) {
			subdomain, err := NewEnvironmentSubdomain(env, "abc123")

			assert.Nil(t, subdomain)
			if assert.IsType(t, errors.CheckoutArgumentError(""), err) {
				assert.NotContains(t, err.Error(), "userinfo-value")
			}
		})
	}
}
