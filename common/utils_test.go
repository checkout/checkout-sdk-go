package common

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/errors"
)

func TestEncodePathSegment(t *testing.T) {
	cases := []struct {
		name     string
		segment  string
		expected string
	}{
		{"plain id is unchanged", "pay_mbabizu24mvu3mela5njyhpit4", "pay_mbabizu24mvu3mela5njyhpit4"},
		{"hyphen, underscore, dot and tilde are unchanged", "a-b_c.d~e", "a-b_c.d~e"},
		{"question mark", "x?y=1", "x%3Fy=1"},
		{"hash", "x#y", "x%23y"},
		{"question mark and hash together", "x?z#w", "x%3Fz%23w"},
		{"percent", "x%y", "x%25y"},
		{"percent-encoded value", "%2e%2e", "%252e%252e"},
		{"space", "x y", "x%20y"},
		{"blank", "   ", "%20%20%20"},
		{"semicolon", "x;y", "x%3By"},
		{"unicode", "pé", "p%C3%A9"},
		{"dots inside a value", "x..y", "x..y"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encoded, err := EncodePathSegment(tc.segment)

			assert.Nil(t, err)
			assert.Equal(t, tc.expected, encoded)
		})
	}
}

func TestEncodePathSegmentRejectsInvalidSegments(t *testing.T) {
	cases := []struct {
		name    string
		segment string
	}{
		{"empty", ""},
		{"dot", "."},
		{"dot dot", ".."},
		// The API decodes %2F and %5C back into separators, so these are refused, not encoded.
		{"slash", "x/y"},
		{"dots and slashes", "../x/y"},
		{"encoded dots and slash", "%2e%2e/x"},
		{"backslash", `..\x`},
		{"carriage return and line feed", "x\r\ny"},
		{"null", "x\x00y"},
		{"unit separator", "x\x1fy"},
		{"delete", "x\x7fy"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			encoded, err := EncodePathSegment(tc.segment)

			assert.Equal(t, "", encoded)
			assert.IsType(t, errors.CheckoutArgumentError(""), err)
			assert.Equal(t, "invalid path segment", err.Error())
		})
	}
}

func TestBuildPathKeepsConstantSegmentsAsGiven(t *testing.T) {
	assert.Equal(t, "/sepa/mandates/man_1", BuildPath("sepa/mandates", "man_1"))
	assert.Equal(t, "/payments", BuildPath("payments"))
}
