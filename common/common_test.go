package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCardTypeNormalized(t *testing.T) {
	cases := []struct {
		name     string
		value    CardType
		expected CardType
	}{
		{
			name:     "when value uses the previous platform casing should map to the current platform constant",
			value:    Charge,
			expected: CardTypeCharge,
		},
		{
			name:     "when value uses the previous platform casing for credit should map to the current platform constant",
			value:    Credit,
			expected: CardTypeCredit,
		},
		{
			name:     "when value uses the previous platform casing for debit should map to the current platform constant",
			value:    Debit,
			expected: CardTypeDebit,
		},
		{
			name:     "when value is two words using the previous platform casing should map to the current platform constant",
			value:    DeferredDebit,
			expected: CardTypeDeferredDebit,
		},
		{
			name:     "when value uses the previous platform casing for prepaid should map to the current platform constant",
			value:    Prepaid,
			expected: CardTypePrepaid,
		},
		{
			name:     "when value is already a current platform constant should be unchanged",
			value:    CardTypeCredit,
			expected: CardTypeCredit,
		},
		{
			name:     "when value is a two word current platform constant should be unchanged",
			value:    CardTypeDeferredDebit,
			expected: CardTypeDeferredDebit,
		},
		{
			name:     "when value is the network token constant should be unchanged",
			value:    CardTypeNetworkToken,
			expected: CardTypeNetworkToken,
		},
		{
			name:     "when value is the unknown constant should be unchanged",
			value:    CardTypeUnknown,
			expected: CardTypeUnknown,
		},
		{
			name:     "when value is lowercase should map to the current platform constant",
			value:    CardType("deferred debit"),
			expected: CardTypeDeferredDebit,
		},
		{
			name:     "when value is empty should stay empty",
			value:    CardType(""),
			expected: CardType(""),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.value.Normalized())
		})
	}
}

func TestCardCategoryNormalized(t *testing.T) {
	cases := []struct {
		name     string
		value    CardCategory
		expected CardCategory
	}{
		{
			name:     "when value uses the previous platform casing should map to the current platform constant",
			value:    Commercial,
			expected: CardCategoryCommercial,
		},
		{
			name:     "when value uses the previous platform casing for consumer should map to the current platform constant",
			value:    Consumer,
			expected: CardCategoryConsumer,
		},
		{
			name:     "when value is already a current platform constant should be unchanged",
			value:    CardCategoryConsumer,
			expected: CardCategoryConsumer,
		},
		{
			name:     "when value is the commercial current platform constant should be unchanged",
			value:    CardCategoryCommercial,
			expected: CardCategoryCommercial,
		},
		{
			name:     "when value is the unknown constant should be unchanged",
			value:    CardCategoryUnknown,
			expected: CardCategoryUnknown,
		},
		{
			name:     "when value is lowercase should map to the current platform constant",
			value:    CardCategory("consumer"),
			expected: CardCategoryConsumer,
		},
		{
			name:     "when value is empty should stay empty",
			value:    CardCategory(""),
			expected: CardCategory(""),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.value.Normalized())
		})
	}
}
