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
		{"slash", "x/y", "x%2Fy"},
		{"question mark", "x?y=1", "x%3Fy=1"},
		{"hash", "x#y", "x%23y"},
		{"slash, question mark and hash together", "x/y?z#w", "x%2Fy%3Fz%23w"},
		{"percent", "x%y", "x%25y"},
		{"percent-encoded value", "%2e%2e", "%252e%252e"},
		{"dots and slashes", "../x/y", "..%2Fx%2Fy"},
		{"backslash", `..\x`, "..%5Cx"},
		{"space", "x y", "x%20y"},
		{"blank", "   ", "%20%20%20"},
		{"semicolon", "x;y", "x%3By"},
		{"control characters", "x\r\ny\x00", "x%0D%0Ay%00"},
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
