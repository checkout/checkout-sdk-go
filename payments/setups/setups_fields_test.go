package setups

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
)

// Verifies the Payment Setups schemas aligned with the 2026-06-29 Checkout.com swagger delta:
//   - PaymentSetup: billing_descriptor, presentment_details, terminal, latest_payment
//   - payment_methods: bacs, card_present, pay_by_bank, stablecoin
//   - order.amount_allocations[] (+ commission)
//   - KlarnaAccountHolder.name

func TestPaymentSetupRequest_NewTopLevelFields(t *testing.T) {
	request := PaymentSetupRequest{
		ProcessingChannelId: "pc_test_abcdefghijklmnopqrstuvw",
		Amount:              1000,
		Currency:            common.GBP,
		BillingDescriptor: &PaymentSetupBillingDescriptor{
			Name:      "SDK Test",
			City:      "London",
			Reference: "order-123",
		},
		PresentmentDetails: &PaymentSetupPresentmentDetails{
			Amount:   1200,
			Currency: common.USD,
		},
		Terminal: &PaymentSetupTerminal{
			Id:            "trm12345",
			LocalDateTime: timePtr(t, "2026-06-01T10:00:00Z"),
		},
	}

	marshalled, err := json.Marshal(request)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"billing_descriptor":`)
	assert.Contains(t, body, `"name":"SDK Test"`)
	assert.Contains(t, body, `"city":"London"`)
	assert.Contains(t, body, `"reference":"order-123"`)
	assert.Contains(t, body, `"presentment_details":`)
	assert.Contains(t, body, `"amount":1200`)
	assert.Contains(t, body, `"currency":"USD"`)
	assert.Contains(t, body, `"terminal":`)
	assert.Contains(t, body, `"id":"trm12345"`)
	assert.Contains(t, body, `"local_date_time":"2026-06-01T10:00:00Z"`)
}

func TestPaymentSetupResponse_NewFieldsAndLatestPayment(t *testing.T) {
	payload := `{
		"id": "psp_test_abcdefghijklmnopqr",
		"billing_descriptor": {"name": "SDK Test", "city": "London", "reference": "order-123"},
		"presentment_details": {"amount": 1200, "currency": "USD"},
		"terminal": {"id": "trm12345", "local_date_time": "2026-06-01T10:00:00Z"},
		"latest_payment": {"id": "pay_test_abcdefghijklmnopqr", "status": "Authorized"}
	}`

	var response PaymentSetupResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.Equal(t, "SDK Test", response.BillingDescriptor.Name)
	assert.Equal(t, "London", response.BillingDescriptor.City)
	assert.Equal(t, int64(1200), response.PresentmentDetails.Amount)
	assert.Equal(t, common.USD, response.PresentmentDetails.Currency)
	assert.Equal(t, "trm12345", response.Terminal.Id)
	assert.NotNil(t, response.Terminal.LocalDateTime)
	assert.Equal(t, "pay_test_abcdefghijklmnopqr", response.LatestPayment["id"])
	assert.Equal(t, "Authorized", response.LatestPayment["status"])
}

func TestPaymentMethods_NewConfigs(t *testing.T) {
	methods := PaymentMethods{
		Bacs: &BacsPaymentMethod{
			InstrumentId:      "src_test_abcdefghijklmnopqr",
			AccountNumber:     "12345678",
			BankCode:          "200000",
			Country:           common.GB,
			Currency:          "GBP",
			AllowPartialMatch: true,
			AccountHolder: &BacsAccountHolder{
				Type:      BacsAccountHolderIndividual,
				FirstName: "John",
				LastName:  "Smith",
				Email:     "john.smith@example.com",
			},
		},
		CardPresent: &CardPresentPaymentMethod{
			Track2:    "track2data",
			Emv:       "emvdata",
			EntryMode: "chip",
			Name:      "John Smith",
			Pin: &CardPresentPin{
				KeySetId:    "kset_123",
				Block:       "block_data",
				BlockFormat: "ISO-0",
			},
		},
		PayByBank: &PayByBankPaymentMethod{
			BankId: "bank_123",
			Action: &PayByBankAction{
				Type: "select_bank",
				Banks: []PayByBankBank{
					{BankId: "bank_123", DisplayName: "Test Bank", LogoUrl: "https://example.com/logo.png", Available: true},
				},
			},
		},
		Stablecoin: &StablecoinPaymentMethod{},
	}

	marshalled, err := json.Marshal(methods)
	assert.NoError(t, err)
	body := string(marshalled)
	// wiring keys
	assert.Contains(t, body, `"bacs":`)
	assert.Contains(t, body, `"card_present":`)
	assert.Contains(t, body, `"pay_by_bank":`)
	assert.Contains(t, body, `"stablecoin":`)
	// bacs
	assert.Contains(t, body, `"instrument_id":"src_test_abcdefghijklmnopqr"`)
	assert.Contains(t, body, `"account_number":"12345678"`)
	assert.Contains(t, body, `"bank_code":"200000"`)
	assert.Contains(t, body, `"allow_partial_match":true`)
	assert.Contains(t, body, `"type":"individual"`)
	// card_present
	assert.Contains(t, body, `"track2":"track2data"`)
	assert.Contains(t, body, `"entry_mode":"chip"`)
	assert.Contains(t, body, `"key_set_id":"kset_123"`)
	assert.Contains(t, body, `"block_format":"ISO-0"`)
	// pay_by_bank
	assert.Contains(t, body, `"bank_id":"bank_123"`)
	assert.Contains(t, body, `"type":"select_bank"`)
	assert.Contains(t, body, `"display_name":"Test Bank"`)
	assert.Contains(t, body, `"logo_url":"https://example.com/logo.png"`)
}

func TestBacsAccountHolderType_Values(t *testing.T) {
	assert.Equal(t, BacsAccountHolderType("individual"), BacsAccountHolderIndividual)
	assert.Equal(t, BacsAccountHolderType("corporate"), BacsAccountHolderCorporate)
}

func TestPaymentSetupOrder_AmountAllocations(t *testing.T) {
	order := PaymentSetupOrder{
		AmountAllocations: []PaymentSetupAmountAllocation{
			{
				Id:        "ent_test_abcdefghijklmnopqr",
				Amount:    750,
				Reference: "split-1",
				Commission: &AmountAllocationCommission{
					Amount:     50,
					Percentage: 2.5,
				},
			},
		},
	}

	marshalled, err := json.Marshal(order)
	assert.NoError(t, err)
	body := string(marshalled)
	assert.Contains(t, body, `"amount_allocations":`)
	assert.Contains(t, body, `"id":"ent_test_abcdefghijklmnopqr"`)
	assert.Contains(t, body, `"amount":750`)
	assert.Contains(t, body, `"reference":"split-1"`)
	assert.Contains(t, body, `"commission":`)
	assert.Contains(t, body, `"percentage":2.5`)

	var decoded PaymentSetupOrder
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Len(t, decoded.AmountAllocations, 1)
	assert.Equal(t, int64(750), decoded.AmountAllocations[0].Amount)
	assert.Equal(t, int64(50), decoded.AmountAllocations[0].Commission.Amount)
	assert.Equal(t, 2.5, decoded.AmountAllocations[0].Commission.Percentage)
}

func TestKlarnaAccountHolder_Name(t *testing.T) {
	marshalled, err := json.Marshal(KlarnaAccountHolder{Name: "John Smith"})
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"name":"John Smith"`)

	var decoded KlarnaAccountHolder
	assert.NoError(t, json.Unmarshal([]byte(`{"name":"Jane Doe"}`), &decoded))
	assert.Equal(t, "Jane Doe", decoded.Name)
}

func timePtr(t *testing.T, value string) *time.Time {
	parsed, err := time.Parse(time.RFC3339, value)
	assert.NoError(t, err)
	return &parsed
}
