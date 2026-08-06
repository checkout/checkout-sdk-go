package payments

import (
	"encoding/json"
	"testing"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/stretchr/testify/assert"
)

// Verifies PaymentProcessing (PaymentResponse.processing, POST /payments) deserializes
// every field declared by the inline swagger schema. This schema is distinct from
// ProcessingData, which is returned by GET /payments/{id}.
func TestPaymentProcessing_UnmarshalAllFields(t *testing.T) {
	payload := `{
		"retrieval_reference_number":"909913440644",
		"acquirer_transaction_id":"440644309099499894406",
		"recommendation_code":"02",
		"scheme":"Mastercard",
		"partner_merchant_advice_code":"24",
		"partner_response_code":"ER_WRONG_TICKET",
		"partner_order_id":"5GK24544NA744002L",
		"partner_payment_id":"440644309099499894406",
		"partner_status":"pending",
		"partner_transaction_id":"txn_abc",
		"partner_session_id":"session_abc",
		"partner_error_codes":["ERR_001","ERR_002"],
		"partner_error_message":"Payment declined",
		"partner_authorization_code":"auth_123",
		"partner_authorization_response_code":"00",
		"surcharge_amount":200,
		"pan_type_processed":"dpan",
		"fallback_source_used":true,
		"cko_network_token_available":true,
		"purchase_country":"GB",
		"foreign_retailer_amount":200,
		"scheme_merchant_id":"123456",
		"reconciliation_id":"4123495123",
		"aggregator":{
			"sub_merchant_id":"9cf70789ba90123",
			"aggregator_id_visa":"10012345",
			"aggregator_id_mc":"00000123456"
		},
		"partner_client_token":"token_abc",
		"continuation_payload":"payload_abc",
		"pun":"pun_abc",
		"merchant_category_code":"5311",
		"scheme_transaction_link_id":"MTL-XYZ-789"
	}`

	var processing PaymentProcessing
	err := json.Unmarshal([]byte(payload), &processing)

	assert.NoError(t, err)
	assert.Equal(t, "909913440644", processing.RetrievalReferenceNumber)
	assert.Equal(t, "440644309099499894406", processing.AcquirerTransactionId)
	assert.Equal(t, "02", processing.RecommendationCode)
	assert.Equal(t, "Mastercard", processing.Scheme)
	assert.Equal(t, "24", processing.PartnerMerchantAdviceCode)
	assert.Equal(t, "ER_WRONG_TICKET", processing.PartnerResponseCode)
	assert.Equal(t, "5GK24544NA744002L", processing.PartnerOrderId)
	assert.Equal(t, "440644309099499894406", processing.PartnerPaymentId)
	assert.Equal(t, "pending", processing.PartnerStatus)
	assert.Equal(t, "txn_abc", processing.PartnerTransactionId)
	assert.Equal(t, "session_abc", processing.PartnerSessionId)
	assert.Equal(t, []string{"ERR_001", "ERR_002"}, processing.PartnerErrorCodes)
	assert.Equal(t, "Payment declined", processing.PartnerErrorMessage)
	assert.Equal(t, "auth_123", processing.PartnerAuthorizationCode)
	assert.Equal(t, "00", processing.PartnerAuthorizationResponseCode)
	assert.Equal(t, int64(200), processing.SurchargeAmount)
	assert.Equal(t, DPAN, processing.PanTypeProcessed)
	assert.True(t, processing.FallbackSourceUsed)
	assert.True(t, processing.CkoNetworkTokenAvailable)
	assert.Equal(t, common.GB, processing.PurchaseCountry)
	assert.Equal(t, int64(200), processing.ForeignRetailerAmount)
	assert.Equal(t, "123456", processing.SchemeMerchantId)
	assert.Equal(t, "4123495123", processing.ReconciliationId)
	assert.NotNil(t, processing.Aggregator)
	assert.Equal(t, "9cf70789ba90123", processing.Aggregator.SubMerchantId)
	assert.Equal(t, "10012345", processing.Aggregator.AggregatorIdVisa)
	assert.Equal(t, "00000123456", processing.Aggregator.AggregatorIdMc)
	assert.Equal(t, "token_abc", processing.PartnerClientToken)
	assert.Equal(t, "payload_abc", processing.ContinuationPayload)
	assert.Equal(t, "pun_abc", processing.Pun)
	assert.Equal(t, "5311", processing.MerchantCategoryCode)
	assert.Equal(t, "MTL-XYZ-789", processing.SchemeTransactionLinkId)
}

// scheme_merchant_id is a string in the spec, so alphanumeric merchant identifiers
// must deserialize without error.
func TestPaymentProcessing_UnmarshalAlphanumericSchemeMerchantId(t *testing.T) {
	var processing PaymentProcessing
	err := json.Unmarshal([]byte(`{"scheme_merchant_id":"MID-0012AB"}`), &processing)

	assert.NoError(t, err)
	assert.Equal(t, "MID-0012AB", processing.SchemeMerchantId)
}

// Round-trips the fields added for the POST /payments processing alignment.
func TestPaymentProcessing_MarshalRoundTrip(t *testing.T) {
	original := PaymentProcessing{
		SchemeMerchantId:      "123456",
		PurchaseCountry:       common.GB,
		ForeignRetailerAmount: 200,
		ReconciliationId:      "4123495123",
		FallbackSourceUsed:    true,
		Aggregator: &Aggregator{
			SubMerchantId:    "9cf70789ba90123",
			AggregatorIdVisa: "10012345",
			AggregatorIdMc:   "00000123456",
		},
	}

	data, err := json.Marshal(original)
	assert.NoError(t, err)

	var roundTripped PaymentProcessing
	assert.NoError(t, json.Unmarshal(data, &roundTripped))
	assert.Equal(t, original.SchemeMerchantId, roundTripped.SchemeMerchantId)
	assert.Equal(t, original.PurchaseCountry, roundTripped.PurchaseCountry)
	assert.Equal(t, original.ForeignRetailerAmount, roundTripped.ForeignRetailerAmount)
	assert.Equal(t, original.ReconciliationId, roundTripped.ReconciliationId)
	assert.Equal(t, original.FallbackSourceUsed, roundTripped.FallbackSourceUsed)
	assert.Equal(t, original.Aggregator.SubMerchantId, roundTripped.Aggregator.SubMerchantId)
}

// Absent fields must stay zero-valued, with no spurious defaults.
func TestPaymentProcessing_UnmarshalOmittedFields(t *testing.T) {
	var processing PaymentProcessing
	err := json.Unmarshal([]byte(`{"scheme":"Visa"}`), &processing)

	assert.NoError(t, err)
	assert.Equal(t, "Visa", processing.Scheme)
	assert.Empty(t, processing.SchemeMerchantId)
	assert.Empty(t, processing.ReconciliationId)
	assert.Empty(t, processing.PurchaseCountry)
	assert.Zero(t, processing.ForeignRetailerAmount)
	assert.False(t, processing.FallbackSourceUsed)
	assert.Nil(t, processing.Aggregator)
}
