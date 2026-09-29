package payments

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// The swagger types tax_amount, discount_amount, shipping_amount, shipping_tax_amount,
// duty_amount and original_order_amount as `number`, not `integer`, and the live API honours
// that: POST /payments with "tax_amount": 10.5 returns 201 and GET /payments/{id} echoes 10.5
// back. While these fields were int64, that response failed to unmarshal with
// "cannot unmarshal number 10.5 into Go struct field ... of type int64", and because
// encoding/json aborts on the first error, the whole payment response was lost, not just the
// field. Keep them float64.
func TestProcessingDataAcceptsFractionalAmounts(t *testing.T) {
	var data ProcessingData
	err := json.Unmarshal([]byte(`{"tax_amount":10.5}`), &data)

	assert.Nil(t, err)
	assert.Equal(t, 10.5, data.TaxAmount)
}

func TestProcessingSettingsAcceptsFractionalAmounts(t *testing.T) {
	body := `{"tax_amount":10.5,"discount_amount":0.25,"shipping_amount":3.75,` +
		`"shipping_tax_amount":1.5,"duty_amount":2.05,"original_order_amount":99.99}`

	var settings ProcessingSettings
	err := json.Unmarshal([]byte(body), &settings)

	assert.Nil(t, err)
	assert.Equal(t, 10.5, settings.TaxAmount)
	assert.Equal(t, 0.25, settings.DiscountAmount)
	assert.Equal(t, 3.75, settings.ShippingAmount)
	assert.Equal(t, 1.5, settings.ShippingTaxAmount)
	assert.Equal(t, 2.05, settings.DutyAmount)
	assert.Equal(t, 99.99, settings.OriginalOrderAmount)
}

func TestProcessingSettingsStillAcceptsWholeAmounts(t *testing.T) {
	var settings ProcessingSettings
	err := json.Unmarshal([]byte(`{"tax_amount":3000}`), &settings)

	assert.Nil(t, err)
	assert.Equal(t, float64(3000), settings.TaxAmount)
}
