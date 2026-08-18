package common

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Reported internally against the PHP SDK: bank_statement was missing everywhere. It is the
// document type for the bank account behind a payment instrument, not an identity document.
func TestInstrumentDocumentType(t *testing.T) {
	t.Run("should expose bank_statement for instrument documents", func(t *testing.T) {
		assert.Equal(t, InstrumentDocumentType("bank_statement"), BankStatement)
	})

	t.Run("should serialize to the value the API accepts", func(t *testing.T) {
		body, err := json.Marshal(map[string]interface{}{"type": BankStatement})

		assert.Nil(t, err)
		assert.Equal(t, `{"type":"bank_statement"}`, string(body))
	})

	// bank_statement belongs to the instrument document enum, not the identity one. This is what
	// stops the two being merged the next time someone reports it missing from DocumentType.
	t.Run("should keep bank_statement out of the identity document type", func(t *testing.T) {
		identity := []DocumentType{
			PassportDocumentType,
			NationalIdentityCard,
			DrivingLicense,
			CitizenCard,
			ResidencePermit,
			ElectoralId,
		}

		assert.NotContains(t, identity, DocumentType(BankStatement))
	})
}
