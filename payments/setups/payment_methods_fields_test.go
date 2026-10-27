package setups

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
)

// Verifies the Payment Setup payment_type enum, the payment methods with their own spec fields
// (qpay, ideal, knet, bancontact, the Alipay family and similar wallets, tamara), the ach and sepa
// fields, order tipping and surcharge amounts, and the members kept only as deprecated.

func TestPaymentSetupPaymentType_Values(t *testing.T) {
	assert.Equal(t, PaymentSetupPaymentType("regular"), PaymentSetupPaymentTypeRegular)
	assert.Equal(t, PaymentSetupPaymentType("recurring"), PaymentSetupPaymentTypeRecurring)
	assert.Equal(t, PaymentSetupPaymentType("moto"), PaymentSetupPaymentTypeMoto)
	assert.Equal(t, PaymentSetupPaymentType("installment"), PaymentSetupPaymentTypeInstallment)
	assert.Equal(t, PaymentSetupPaymentType("pay_later"), PaymentSetupPaymentTypePayLater)
	assert.Equal(t, PaymentSetupPaymentType("unscheduled"), PaymentSetupPaymentTypeUnscheduled)
}

func TestPaymentSetupRequest_PaymentTypeIsLowercaseOnTheWire(t *testing.T) {
	values := map[PaymentSetupPaymentType]string{
		PaymentSetupPaymentTypeRegular:     `"payment_type":"regular"`,
		PaymentSetupPaymentTypeRecurring:   `"payment_type":"recurring"`,
		PaymentSetupPaymentTypeMoto:        `"payment_type":"moto"`,
		PaymentSetupPaymentTypeInstallment: `"payment_type":"installment"`,
		PaymentSetupPaymentTypePayLater:    `"payment_type":"pay_later"`,
		PaymentSetupPaymentTypeUnscheduled: `"payment_type":"unscheduled"`,
	}
	for value, wire := range values {
		request := PaymentSetupRequest{
			ProcessingChannelId: "pc_q4dbxom5jbgudnjzjpz7j2z6uq",
			Amount:              1000,
			Currency:            common.GBP,
			PaymentType:         value,
		}
		marshalled, err := json.Marshal(request)
		assert.NoError(t, err)
		assert.Contains(t, string(marshalled), wire)

		var decoded PaymentSetupRequest
		assert.NoError(t, json.Unmarshal(marshalled, &decoded))
		assert.Equal(t, request, decoded)
	}

	unset, err := json.Marshal(PaymentSetupRequest{ProcessingChannelId: "pc_q4dbxom5jbgudnjzjpz7j2z6uq"})
	assert.NoError(t, err)
	assert.NotContains(t, string(unset), "payment_type")
}

func TestPaymentSetupResponse_ReadsLowercasePaymentType(t *testing.T) {
	payload := `{"id":"set_6jy4zeybu3fylcm35yaoinjsgi","processing_channel_id":"pc_q4dbxom5jbgudnjzjpz7j2z6uq",
		"amount":1000,"currency":"GBP","payment_type":"pay_later"}`
	var response PaymentSetupResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	assert.Equal(t, PaymentSetupPaymentTypePayLater, response.PaymentType)

	marshalled, err := json.Marshal(response)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"payment_type":"pay_later"`)
}

func TestPaymentSetupTerminalOsKnetAndMandateEnums_Values(t *testing.T) {
	assert.Equal(t, PaymentSetupTerminalType("web"), PaymentSetupTerminalTypeWeb)
	assert.Equal(t, PaymentSetupTerminalType("wap"), PaymentSetupTerminalTypeWap)
	assert.Equal(t, PaymentSetupTerminalType("app"), PaymentSetupTerminalTypeApp)
	assert.Equal(t, PaymentSetupOsType("android"), PaymentSetupOsTypeAndroid)
	assert.Equal(t, PaymentSetupOsType("ios"), PaymentSetupOsTypeIos)
	assert.Equal(t, KnetLanguage("en"), KnetLanguageEn)
	assert.Equal(t, KnetLanguage("ar"), KnetLanguageAr)
	assert.Equal(t, SepaMandateType("core"), SepaMandateTypeCore)
	assert.Equal(t, SepaMandateType("b2b"), SepaMandateTypeB2b)
}

func methodSpecificPaymentMethods() PaymentMethods {
	status := PaymentSetupPaymentMethod{Status: "action_required", Flags: []string{"example_flag"}}
	return PaymentMethods{
		Qpay:       &QpayPaymentMethod{PaymentSetupPaymentMethod: status, NationalId: "21234567890", Description: "Order 123"},
		Ideal:      &IdealPaymentMethod{PaymentSetupPaymentMethod: status, Description: "2 t-shirts"},
		Knet:       &KnetPaymentMethod{PaymentSetupPaymentMethod: status, Language: KnetLanguageAr},
		Bancontact: &BancontactPaymentMethod{PaymentSetupPaymentMethod: status, AccountHolderName: "Jo Doe"},
		AlipayCn: &AlipayCnPaymentMethod{PaymentSetupPaymentMethod: status,
			TerminalType: PaymentSetupTerminalTypeWeb},
		AlipayHk: &AlipayHkPaymentMethod{PaymentSetupPaymentMethod: status,
			TerminalType: PaymentSetupTerminalTypeApp, OsType: PaymentSetupOsTypeIos},
		Gcash: &GCashPaymentMethod{PaymentSetupPaymentMethod: status,
			TerminalType: PaymentSetupTerminalTypeWap, OsType: PaymentSetupOsTypeAndroid},
		Tng: &TngPaymentMethod{PaymentSetupPaymentMethod: status,
			TerminalType: PaymentSetupTerminalTypeApp, OsType: PaymentSetupOsTypeAndroid},
		Dana: &DanaPaymentMethod{PaymentSetupPaymentMethod: status,
			TerminalType: PaymentSetupTerminalTypeWap, OsType: PaymentSetupOsTypeIos},
		Kakaopay: &KakaoPayPaymentMethod{PaymentSetupPaymentMethod: status,
			TerminalType: PaymentSetupTerminalTypeApp, OsType: PaymentSetupOsTypeIos},
		Truemoney: &TrueMoneyPaymentMethod{PaymentSetupPaymentMethod: status,
			TerminalType: PaymentSetupTerminalTypeWap, OsType: PaymentSetupOsTypeAndroid},
		Tamara: &TamaraPaymentMethod{
			PaymentMethodBase: PaymentMethodBase{Status: "action_required", Flags: []string{"example_flag"},
				Initialization: PaymentMethodInitializationEnabled},
			RiskAssessment: &TamaraRiskAssessment{
				PremiumCustomer:   boolPtr(false),
				RiskScore:         intPtr(42),
				OrderCountLast24h: intPtr(0),
				LinkedDeviceCount: intPtr(2),
			},
		},
	}
}

func TestMethodSpecificPaymentMethods_WireKeysAndRoundTrip(t *testing.T) {
	original := methodSpecificPaymentMethods()

	marshalled, err := json.Marshal(original)
	assert.NoError(t, err)
	assert.JSONEq(t, `{
		"qpay":{"status":"action_required","flags":["example_flag"],"national_id":"21234567890","description":"Order 123"},
		"ideal":{"status":"action_required","flags":["example_flag"],"description":"2 t-shirts"},
		"knet":{"status":"action_required","flags":["example_flag"],"language":"ar"},
		"bancontact":{"status":"action_required","flags":["example_flag"],"account_holder_name":"Jo Doe"},
		"alipay_cn":{"status":"action_required","flags":["example_flag"],"terminal_type":"web"},
		"alipay_hk":{"status":"action_required","flags":["example_flag"],"terminal_type":"app","os_type":"ios"},
		"gcash":{"status":"action_required","flags":["example_flag"],"terminal_type":"wap","os_type":"android"},
		"tng":{"status":"action_required","flags":["example_flag"],"terminal_type":"app","os_type":"android"},
		"dana":{"status":"action_required","flags":["example_flag"],"terminal_type":"wap","os_type":"ios"},
		"kakaopay":{"status":"action_required","flags":["example_flag"],"terminal_type":"app","os_type":"ios"},
		"truemoney":{"status":"action_required","flags":["example_flag"],"terminal_type":"wap","os_type":"android"},
		"tamara":{"status":"action_required","flags":["example_flag"],"initialization":"enabled",
			"risk_assessment":{"premium_customer":false,"risk_score":42,"order_count_last24h":0,"linked_device_count":2}}
	}`, string(marshalled))

	var decoded PaymentMethods
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, original, decoded)
}

func TestMethodSpecificPaymentMethods_FromSwaggerExamples(t *testing.T) {
	// Property examples from the spec: qpay.national_id, ideal.description, TerminalType, OsType and
	// tamara.risk_assessment. The spec gives no example for the other fields.
	payload := `{"payment_methods":{
		"qpay":{"national_id":"21234567890"},
		"ideal":{"description":"2 t-shirts"},
		"alipay_cn":{"terminal_type":"web","os_type":"android"},
		"tamara":{"status":"ready","initialization":"disabled",
			"risk_assessment":{"premium_customer":false,"risk_score":42,"order_count_last24h":0,"linked_device_count":2}}
	}}`
	var response PaymentSetupResponse
	assert.NoError(t, json.Unmarshal([]byte(payload), &response))
	methods := response.PaymentMethods
	assert.Equal(t, "21234567890", methods.Qpay.NationalId)
	assert.Equal(t, "2 t-shirts", methods.Ideal.Description)
	assert.Equal(t, PaymentSetupTerminalTypeWeb, methods.AlipayCn.TerminalType)
	assert.Equal(t, PaymentSetupOsTypeAndroid, methods.AlipayCn.OsType)
	assert.Equal(t, "ready", methods.Tamara.Status)
	assert.Equal(t, PaymentMethodInitializationDisabled, methods.Tamara.Initialization)
	risk := methods.Tamara.RiskAssessment
	if assert.NotNil(t, risk) {
		assert.False(t, *risk.PremiumCustomer)
		assert.Equal(t, 42, *risk.RiskScore)
		assert.Equal(t, 0, *risk.OrderCountLast24h)
		assert.Equal(t, 2, *risk.LinkedDeviceCount)
	}

	// An explicit false and zero are sent back, not dropped.
	marshalled, err := json.Marshal(methods.Tamara)
	assert.NoError(t, err)
	assert.Contains(t, string(marshalled), `"premium_customer":false`)
	assert.Contains(t, string(marshalled), `"order_count_last24h":0`)
}

func TestAchPaymentMethod_AllFieldsWireKeysAndRoundTrip(t *testing.T) {
	original := AchPaymentMethod{
		PaymentMethodBase: PaymentMethodBase{Status: "ready", Flags: []string{"example_flag"}},
		AccountType:       "current",
		AccountHolder: &AchAccountHolder{
			Type:        "corporate",
			FirstName:   "John",
			LastName:    "Smith",
			CompanyName: "Acme Corporation",
			DateOfBirth: "1985-05-15",
			Identification: &AchAccountHolderIdentification{
				Type:           "ssn",
				IssuingCountry: "US",
				Number:         "123456789",
			},
		},
		AccountNumber: "4099999992",
		BankCode:      "011075150",
		Country:       common.US,
	}

	marshalled, err := json.Marshal(original)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"status":"ready","flags":["example_flag"],"account_type":"current",
		"account_holder":{"type":"corporate","first_name":"John","last_name":"Smith",
			"company_name":"Acme Corporation","date_of_birth":"1985-05-15",
			"identification":{"type":"ssn","issuing_country":"US","number":"123456789"}},
		"account_number":"4099999992","bank_code":"011075150","country":"US"}`, string(marshalled))

	var decoded AchPaymentMethod
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, original, decoded)
}

func TestSepaPaymentMethod_AllFieldsWireKeysAndRoundTrip(t *testing.T) {
	original := SepaPaymentMethod{
		PaymentMethodBase: PaymentMethodBase{Status: "ready", Flags: []string{"example_flag"}},
		AccountHolder: &SepaAccountHolder{
			Type:        "individual",
			FirstName:   "John",
			LastName:    "Smith",
			CompanyName: "Acme Corporation",
		},
		AccountNumber: "DE89370400440532013000",
		Country:       common.DE,
		Currency:      "EUR",
		Mandate: &SepaMandate{
			Id:              "mandate_123",
			Type:            SepaMandateTypeB2b,
			DateOfSignature: shortDate(t, 2026, time.January, 15, 0, 0),
		},
	}

	marshalled, err := json.Marshal(original)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"status":"ready","flags":["example_flag"],
		"account_holder":{"type":"individual","first_name":"John","last_name":"Smith","company_name":"Acme Corporation"},
		"account_number":"DE89370400440532013000","country":"DE","currency":"EUR",
		"mandate":{"id":"mandate_123","type":"b2b","date_of_signature":"2026-01-15"}}`, string(marshalled))

	var decoded SepaPaymentMethod
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, original, decoded)

	var core SepaMandate
	assert.NoError(t, json.Unmarshal([]byte(`{"type":"core"}`), &core))
	assert.Equal(t, SepaMandateTypeCore, core.Type)
}

func TestPaymentSetupOrder_TippingAndSurchargeAmounts(t *testing.T) {
	original := PaymentSetupOrder{TippingAmount: 2, SurchargeAmount: 3}

	marshalled, err := json.Marshal(original)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"tipping_amount":2,"surcharge_amount":3}`, string(marshalled))

	var decoded PaymentSetupOrder
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, original, decoded)
}

func TestCustomerMerchantAccount_LastPaymentAmountKeepsDecimals(t *testing.T) {
	var account CustomerMerchantAccount
	assert.NoError(t, json.Unmarshal([]byte(`{"last_payment_amount":55.99}`), &account))
	assert.Equal(t, 55.99, account.LastPaymentAmount)

	marshalled, err := json.Marshal(account)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"last_payment_amount":55.99}`, string(marshalled))
}

func TestPaymentMethods_EmptyMethodsSendNothing(t *testing.T) {
	// Status and flags are read only and initialization has no SDK default, so an empty method
	// sends an empty object.
	methods := PaymentMethods{
		Bizum:      &BizumPaymentMethod{},
		Blik:       &BlikPaymentMethod{},
		PayByBank:  &PayByBankPaymentMethod{},
		Stablecoin: &StablecoinPaymentMethod{},
		Instrument: &InstrumentPaymentMethod{},
		Eps:        &SimplePaymentMethod{},
		Multibanco: &MultibancoPaymentMethod{},
		P24:        &P24PaymentMethod{},
		Swish:      &SwishPaymentMethod{},
		Ach:        &AchPaymentMethod{},
		Sepa:       &SepaPaymentMethod{},
		Qpay:       &QpayPaymentMethod{},
		Tamara:     &TamaraPaymentMethod{},
	}
	marshalled, err := json.Marshal(methods)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"bizum":{},"blik":{},"paybybank":{},"stablecoin":{},"instrument":{},"eps":{},
		"multibanco":{},"p24":{},"swish":{},"ach":{},"sepa":{},"qpay":{},"tamara":{}}`, string(marshalled))

	empty, err := json.Marshal(PaymentMethods{})
	assert.NoError(t, err)
	assert.Equal(t, `{}`, string(empty))
}

// The deprecated members stay on the wire until the next major version, with the same key and
// value as before, and the types keep embedding PaymentMethodBase, so literals that name it and
// the promoted Initialization compile as before.
func TestDeprecatedInitialization_StillSerializes(t *testing.T) {
	options := &PaymentMethodOptions{Sdk: &PaymentMethodOption{Id: "opt_123"}}
	base := PaymentMethodBase{Initialization: PaymentMethodInitializationEnabled}
	methods := PaymentMethods{
		Bizum:      &BizumPaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		Blik:       &BlikPaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		PayByBank:  &PayByBankPaymentMethod{PaymentMethodBase: base},
		Stablecoin: &StablecoinPaymentMethod{PaymentMethodBase: base},
		Instrument: &InstrumentPaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		Paynow:     &SimplePaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		Multibanco: &MultibancoPaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		P24:        &P24PaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		Swish:      &SwishPaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		Ach:        &AchPaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		Sepa:       &SepaPaymentMethod{PaymentMethodBase: base, PaymentMethodOptions: options},
		Bacs:       &BacsPaymentMethod{PaymentMethodOptions: options},
	}

	marshalled, err := json.Marshal(methods)
	assert.NoError(t, err)
	withBoth := `{"initialization":"enabled","payment_method_options":{"sdk":{"id":"opt_123"}}}`
	initOnly := `{"initialization":"enabled"}`
	assert.JSONEq(t, `{
		"bizum":`+withBoth+`,"blik":`+withBoth+`,"paybybank":`+initOnly+`,"stablecoin":`+initOnly+`,
		"instrument":`+withBoth+`,"paynow":`+withBoth+`,"multibanco":`+withBoth+`,"p24":`+withBoth+`,
		"swish":`+withBoth+`,"ach":`+withBoth+`,"sepa":`+withBoth+`,
		"bacs":{"payment_method_options":{"sdk":{"id":"opt_123"}}}
	}`, string(marshalled))

	var decoded PaymentMethods
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, methods, decoded)
	for _, initialization := range []PaymentMethodInitialization{
		decoded.Bizum.Initialization, decoded.Blik.Initialization, decoded.PayByBank.Initialization,
		decoded.Stablecoin.Initialization, decoded.Instrument.Initialization, decoded.Paynow.Initialization,
		decoded.Multibanco.Initialization, decoded.P24.Initialization, decoded.Swish.Initialization,
		decoded.Ach.Initialization, decoded.Sepa.Initialization,
	} {
		assert.Equal(t, PaymentMethodInitializationEnabled, initialization)
	}
}

func TestDeprecatedCardPresent_StillSerializes(t *testing.T) {
	methods := PaymentMethods{
		CardPresent: &CardPresentPaymentMethod{
			PaymentMethodBase: PaymentMethodBase{Initialization: PaymentMethodInitializationDisabled},
			Track2:            "track2data",
			Pin:               &CardPresentPin{KeySetId: "kset_123", Block: "block_data", BlockFormat: "ISO-0"},
		},
	}
	marshalled, err := json.Marshal(methods)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"card_present":{"initialization":"disabled","track2":"track2data",
		"pin":{"key_set_id":"kset_123","block":"block_data","block_format":"ISO-0"}}}`, string(marshalled))

	var decoded PaymentMethods
	assert.NoError(t, json.Unmarshal(marshalled, &decoded))
	assert.Equal(t, methods, decoded)
}
