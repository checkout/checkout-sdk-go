package apm

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/payments"
)

// Schema validation tests for requestBacsSource against PaymentRequestBacsSource, which declares a
// type and an id only, and for the SEPA source's mandate_type.

func TestRequestBacsSourceSerializesTypeAndId(t *testing.T) {
	s := NewRequestBacsSource()
	s.Id = "src_wmlfc3zyhqzehihu7giusaaawu"

	raw, err := json.Marshal(s)
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "bacs", body["type"])
	assert.Equal(t, "src_wmlfc3zyhqzehihu7giusaaawu", body["id"])
	assert.Len(t, body, 2)
	assert.Equal(t, payments.BacsSource, s.GetType())
}

func TestRequestSepaSourceCarriesTheMandateType(t *testing.T) {
	s := NewRequestSepaSource()
	s.MandateType = SepaMandateB2B

	raw, err := json.Marshal(s)
	assert.Nil(t, err)

	var body map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &body))

	assert.Equal(t, "sepa", body["type"])
	assert.Equal(t, "B2B", body["mandate_type"])
}

func TestSourceTypeCarriesTheFourAddedValues(t *testing.T) {
	// bacs, mobilepay, swish and vipps are declared by PaymentRequestSourceType and were missing.
	assert.Equal(t, "bacs", string(payments.BacsSource))
	assert.Equal(t, "mobilepay", string(payments.MobilepaySource))
	assert.Equal(t, "swish", string(payments.SwishSource))
	assert.Equal(t, "vipps", string(payments.VippsSource))
}
