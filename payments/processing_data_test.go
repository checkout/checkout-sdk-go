package payments

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies the fields recently aligned with the Checkout.com swagger spec
// deserialize correctly from the API wire format.
//   - scheme
//   - partner_fraud_status
//   - partner_merchant_advice_code
//   - accommodation_data
//   - airline_data
//   - failure_code
//   - partner_code
//   - partner_response_code
//   - fallback_source_used
//   - scheme_transaction_link_id (Mastercard Transaction Link Identifier)
func TestProcessingData_UnmarshalAllNewFields(t *testing.T) {
	payload := `{
		"scheme":"ACCEL",
		"partner_fraud_status":"Accepted",
		"partner_merchant_advice_code":"24",
		"failure_code":"partner_error",
		"partner_code":"902111",
		"partner_response_code":"DECLINED",
		"fallback_source_used":true,
		"scheme_transaction_link_id":"MTL-XYZ-789",
		"accommodation_data":[{"name":"Grand Hotel"}],
		"airline_data":[{"ticket":{"number":"045-21351455613"}}]
	}`

	var data ProcessingData
	err := json.Unmarshal([]byte(payload), &data)

	assert.NoError(t, err)
	assert.Equal(t, "ACCEL", data.Scheme)
	assert.Equal(t, "Accepted", data.PartnerFraudStatus)
	assert.Equal(t, "24", data.PartnerMerchantAdviceCode)
	assert.Equal(t, "partner_error", data.FailureCode)
	assert.Equal(t, "902111", data.PartnerCode)
	assert.Equal(t, "DECLINED", data.PartnerResponseCode)
	assert.True(t, data.FallbackSourceUsed)
	assert.Equal(t, "MTL-XYZ-789", data.SchemeTransactionLinkId)

	assert.Len(t, data.AccommodationData, 1)
	assert.Equal(t, "Grand Hotel", data.AccommodationData[0].Name)

	assert.Len(t, data.AirlineData, 1)
	assert.NotNil(t, data.AirlineData[0].Ticket)
	assert.Equal(t, "045-21351455613", data.AirlineData[0].Ticket.Number)
}

// Verifies the new fields stay zero-valued (no spurious defaults) when absent
// from the API response.
func TestProcessingData_LeavesNewFieldsZeroWhenAbsent(t *testing.T) {
	payload := `{"locale":"en-GB"}`

	var data ProcessingData
	err := json.Unmarshal([]byte(payload), &data)

	assert.NoError(t, err)
	assert.Equal(t, "en-GB", data.Locale)
	assert.Empty(t, data.Scheme)
	assert.Empty(t, data.PartnerFraudStatus)
	assert.Empty(t, data.PartnerMerchantAdviceCode)
	assert.Empty(t, data.FailureCode)
	assert.Empty(t, data.PartnerCode)
	assert.Empty(t, data.PartnerResponseCode)
	assert.False(t, data.FallbackSourceUsed)
	assert.Empty(t, data.SchemeTransactionLinkId)
	assert.Nil(t, data.AccommodationData)
	assert.Nil(t, data.AirlineData)
}

// Verifies the Mastercard Transaction Link Identifier deserializes from the
// processing object returned in a payment response (PaymentProcessing), and
// serializes back to its snake_case wire name.
func TestPaymentProcessing_SchemeTransactionLinkId(t *testing.T) {
	payload := `{
		"retrieval_reference_number":"RRN001",
		"acquirer_transaction_id":"ACQ001",
		"scheme":"Mastercard",
		"scheme_transaction_link_id":"MTL-XYZ-789"
	}`

	var processing PaymentProcessing
	err := json.Unmarshal([]byte(payload), &processing)

	assert.NoError(t, err)
	assert.Equal(t, "RRN001", processing.RetrievalReferenceNumber)
	assert.Equal(t, "ACQ001", processing.AcquirerTransactionId)
	assert.Equal(t, "Mastercard", processing.Scheme)
	assert.Equal(t, "MTL-XYZ-789", processing.SchemeTransactionLinkId)

	marshalled, err := json.Marshal(PaymentProcessing{SchemeTransactionLinkId: "MTL-001"})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"scheme_transaction_link_id":"MTL-001"`)
}

// Verifies the field stays zero-valued when absent from the payment response
// processing object (it is optional and only populated for Mastercard transactions).
func TestPaymentProcessing_SchemeTransactionLinkIdAbsent(t *testing.T) {
	payload := `{"retrieval_reference_number":"RRN001"}`

	var processing PaymentProcessing
	err := json.Unmarshal([]byte(payload), &processing)

	assert.NoError(t, err)
	assert.Equal(t, "RRN001", processing.RetrievalReferenceNumber)
	assert.Empty(t, processing.SchemeTransactionLinkId)
}
