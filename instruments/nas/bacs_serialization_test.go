package nas

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/payments"
)

// Schema validation tests for the Bacs Direct Debit, SEPA and ACH instrument families against
// StoreBacsInstrumentRequest, UpdateBacsInstrumentRequest, StoreSepaInstrumentRequest,
// StoreAchInstrumentRequest and RetrieveBacsInstrumentResponse.

func decode(t *testing.T, v interface{}) map[string]interface{} {
	t.Helper()
	raw, err := json.Marshal(v)
	assert.Nil(t, err)
	var out map[string]interface{}
	assert.Nil(t, json.Unmarshal(raw, &out))
	return out
}

func TestBacsStoreRequestSerializesEveryProperty(t *testing.T) {
	allowPartial := false
	r := NewCreateBacsInstrumentRequest()
	r.Account = &BacsInstrumentAccount{ProcessingChannelId: "pc_q4dbxom5jbgudnjzjpz7j2z6uq"}
	r.InstrumentData = &BacsInstrumentData{
		AccountNumber:     "86753246",
		BankCode:          "040004",
		Country:           common.GB,
		Currency:          common.GBP,
		PaymentType:       BacsRecurring,
		AllowPartialMatch: &allowPartial,
	}
	r.AccountHolder = &CreateBacsAccountHolder{
		FirstName: "John",
		LastName:  "Smith",
		BillingAddress: &BacsBillingAddress{
			AddressLine1: "Cloverfield St.",
			AddressLine2: "23A",
			City:         "London",
			Zip:          "SW1A 1AA",
			Country:      common.GB,
		},
	}

	body := decode(t, r)

	assert.Equal(t, "bacs", body["type"])
	assert.Equal(t, "pc_q4dbxom5jbgudnjzjpz7j2z6uq",
		body["account"].(map[string]interface{})["processing_channel_id"])

	data := body["instrument_data"].(map[string]interface{})
	assert.Equal(t, "86753246", data["account_number"])
	assert.Equal(t, "040004", data["bank_code"])
	assert.Equal(t, "GB", data["country"])
	assert.Equal(t, "GBP", data["currency"])
	assert.Equal(t, "Recurring", data["payment_type"])
	assert.Equal(t, false, data["allow_partial_match"])

	holder := body["account_holder"].(map[string]interface{})
	assert.Equal(t, "John", holder["first_name"])
	assert.Equal(t, "Smith", holder["last_name"])
	address := holder["billing_address"].(map[string]interface{})
	assert.Equal(t, "Cloverfield St.", address["address_line1"])
	assert.Equal(t, "23A", address["address_line2"])
	assert.Equal(t, "London", address["city"])
	assert.Equal(t, "SW1A 1AA", address["zip"])
	assert.Equal(t, "GB", address["country"])

	// StoreBacsInstrumentRequest.account_holder declares no company_name or type.
	_, hasCompany := holder["company_name"]
	_, hasType := holder["type"]
	assert.False(t, hasCompany)
	assert.False(t, hasType)
}

func TestBacsUpdateRequestSerializesTheFivePropertyAccountHolder(t *testing.T) {
	allowPartial := true
	r := NewUpdateBacsInstrumentRequest()
	r.InstrumentData = &BacsInstrumentData{
		PaymentType:       BacsRegular,
		AllowPartialMatch: &allowPartial,
	}
	r.AccountHolder = &UpdateBacsAccountHolder{
		FirstName:   "John",
		LastName:    "Smith",
		CompanyName: "Wayne Enterprises",
		Type:        InstrumentAccountHolderCorporate,
	}

	body := decode(t, r)

	assert.Equal(t, "bacs", body["type"])
	data := body["instrument_data"].(map[string]interface{})
	assert.Equal(t, "Regular", data["payment_type"])
	assert.Equal(t, true, data["allow_partial_match"])
	holder := body["account_holder"].(map[string]interface{})
	assert.Equal(t, "Wayne Enterprises", holder["company_name"])
	assert.Equal(t, "corporate", holder["type"])
}

func TestSepaPaymentTypeStaysLowercaseAndBacsStaysCapitalized(t *testing.T) {
	// The specification declares the SEPA payment_type lowercase and the Bacs Direct Debit
	// payment_type capitalized. This is the regression test that stops the two being unified.
	assert.Equal(t, "recurring", string(SepaRecurring))
	assert.Equal(t, "regular", string(SepaRegular))
	assert.Equal(t, "Recurring", string(BacsRecurring))
	assert.Equal(t, "Regular", string(BacsRegular))

	// payments.PaymentType serializes capitalized values and carries values neither instrument
	// schema allows, so it must not be used for either field.
	assert.NotEqual(t, string(SepaRecurring), string(payments.Recurring))
}

func TestSepaStoreRequestCarriesTheMandateType(t *testing.T) {
	r := NewCreateSepaInstrumentRequest()
	r.InstrumentData = &InstrumentData{
		Type:          SepaMandateB2B,
		AccountNumber: "FR2810096000509685512959O86",
		Country:       common.FR,
		Currency:      common.EUR,
		PaymentType:   SepaRecurring,
		MandateId:     "1234567890",
	}
	r.AccountHolder = &SepaAccountHolder{
		FirstName:   "John",
		LastName:    "Wick",
		CompanyName: "Checkout.com",
		Type:        InstrumentAccountHolderIndividual,
	}

	body := decode(t, r)
	data := body["instrument_data"].(map[string]interface{})

	assert.Equal(t, "sepa", body["type"])
	assert.Equal(t, "B2B", data["type"])
	assert.Equal(t, "recurring", data["payment_type"])
	assert.Equal(t, "individual", body["account_holder"].(map[string]interface{})["type"])
}

func TestAchAccountTypeIsNotTheBankAccountSet(t *testing.T) {
	// The ACH field declares savings and checking; common.AccountType declares savings, current
	// and cash.
	assert.Equal(t, "savings", string(AchSavings))
	assert.Equal(t, "checking", string(AchChecking))
	assert.Equal(t, "current", string(common.Current))
}

func TestAchAccountHolderDeclaresNoBillingAddress(t *testing.T) {
	r := NewCreateAchInstrumentRequest()
	r.InstrumentData = &AchInstrumentData{
		AccountType:   AchChecking,
		AccountNumber: "4099999992",
		BankCode:      "211370545",
		Currency:      common.USD,
		Country:       common.US,
	}
	r.AccountHolder = &AchAccountHolder{
		FirstName:   "John",
		LastName:    "Smith",
		CompanyName: "Smith Enterprises",
		Type:        InstrumentAccountHolderCorporate,
	}

	body := decode(t, r)
	holder := body["account_holder"].(map[string]interface{})

	assert.Equal(t, "checking", body["instrument_data"].(map[string]interface{})["account_type"])
	assert.Len(t, holder, 4)
	_, hasAddress := holder["billing_address"]
	assert.False(t, hasAddress)
}

func TestInstrumentResponseDispatchResolvesBacs(t *testing.T) {
	cases := []struct {
		name string
		body string
		want func(t *testing.T)
	}{}
	_ = cases

	createBody := `{"type":"bacs","id":"src_wmlfc3zyhqzehihu7giusaaawu","fingerprint":"vnsdrvikkvre3dtrjjvlm5du4q"}`

	var create CreateInstrumentResponse
	assert.Nil(t, json.Unmarshal([]byte(createBody), &create))
	assert.NotNil(t, create.CreateBacsInstrumentResponse)
	assert.Equal(t, common.Bacs, create.CreateBacsInstrumentResponse.Type)
	assert.Equal(t, "src_wmlfc3zyhqzehihu7giusaaawu", create.CreateBacsInstrumentResponse.Id)

	var update UpdateInstrumentResponse
	assert.Nil(t, json.Unmarshal([]byte(createBody), &update))
	assert.NotNil(t, update.UpdateBacsInstrumentResponse)
	assert.Equal(t, common.Bacs, update.UpdateBacsInstrumentResponse.Type)

	getBody := `{
		"type":"bacs",
		"id":"src_wmlfc3zyhqzehihu7giusaaawu",
		"fingerprint":"vnsdrvikkvre3dtrjjvlm5du4q",
		"created_on":"2021-01-01T00:00:00Z",
		"modified_on":"2021-02-02T10:30:00Z",
		"vault_id":"vid_wmlfc3zyhqzehihu7giusaaawu",
		"account":{"client_id":"cli_x","processing_channel_id":"pc_x"},
		"validations":[{"type":"name_match"}],
		"instrument_data":{
			"account_number":"86753246","bank_code":"040004","country":"GB","currency":"GBP",
			"payment_type":"Recurring","allow_partial_match":true,
			"status":"INVALID","match_status":"no match","description":"The name did not match.",
			"mandate_id":"6PZ6KFI3KW3UFHAM3J"
		},
		"account_holder":{"first_name":"Hannah","last_name":"Bret","type":"corporate"}
	}`

	var get GetInstrumentResponse
	assert.Nil(t, json.Unmarshal([]byte(getBody), &get))
	assert.NotNil(t, get.GetBacsInstrumentResponse)
	r := get.GetBacsInstrumentResponse
	assert.Equal(t, common.Bacs, r.Type)
	assert.Equal(t, "vid_wmlfc3zyhqzehihu7giusaaawu", r.VaultId)
	assert.Equal(t, "pc_x", r.Account.ProcessingChannelId)
	assert.Len(t, r.Validations, 1)
	assert.Equal(t, BacsRecurring, r.InstrumentData.PaymentType)
	assert.Equal(t, "INVALID", r.InstrumentData.Status)
	assert.Equal(t, "no match", r.InstrumentData.MatchStatus)
	assert.Equal(t, "6PZ6KFI3KW3UFHAM3J", r.InstrumentData.MandateId)
	assert.Equal(t, InstrumentAccountHolderCorporate, r.AccountHolder.Type)
}

func TestInstrumentTypeCarriesBacs(t *testing.T) {
	assert.Equal(t, "bacs", string(common.Bacs))
	assert.Equal(t, "ach", string(common.Ach))
	assert.Equal(t, common.Bacs, NewCreateBacsInstrumentRequest().Type)
	assert.Equal(t, common.Bacs, NewUpdateBacsInstrumentRequest().Type)
}

// The retrieve response uses a Get-specific account holder rather than the request-named
// UpdateBacsAccountHolder. The two are structurally identical, which is why the reuse compiled.
func TestRetrieveResponseUsesAGetSpecificAccountHolder(t *testing.T) {
	field, found := reflect.TypeOf(GetBacsInstrumentResponse{}).FieldByName("AccountHolder")
	assert.True(t, found)
	assert.Equal(t, "*nas.GetBacsAccountHolder", field.Type.String())

	// Same five properties as the update shape, per the specification.
	got := map[string]bool{}
	ht := reflect.TypeOf(GetBacsAccountHolder{})
	for i := 0; i < ht.NumField(); i++ {
		got[ht.Field(i).Tag.Get("json")] = true
	}
	for _, tag := range []string{
		"first_name,omitempty", "last_name,omitempty", "company_name,omitempty",
		"billing_address,omitempty", "type,omitempty",
	} {
		assert.True(t, got[tag], "expected json tag %s", tag)
	}
	assert.Len(t, got, 5)
}

// Every exported Bacs response type carries GoDoc, per the Go SDK conventions.
func TestBacsResponseTypesAreDocumented(t *testing.T) {
	for _, file := range []string{"create.go", "get.go", "update.go"} {
		src, err := os.ReadFile(file)
		assert.Nil(t, err)
		text := string(src)
		for _, name := range []string{
			"CreateBacsInstrumentResponse", "GetBacsInstrumentResponse", "UpdateBacsInstrumentResponse",
		} {
			idx := strings.Index(text, "\t"+name+" struct {")
			if idx < 0 {
				continue
			}
			preceding := text[:idx]
			lastLine := preceding[strings.LastIndex(preceding[:len(preceding)-1], "\n")+1:]
			assert.Contains(t, lastLine, "//", "%s in %s needs GoDoc", name, file)
		}
	}
}

// The validations slice is untyped because the specification publishes no item schema. The comment
// recording that must survive, so the next reader does not assume it was an oversight.
func TestValidationsIsDocumentedAsUntyped(t *testing.T) {
	src, err := os.ReadFile("get.go")
	assert.Nil(t, err)
	assert.Contains(t, string(src), "no item schema")
}
