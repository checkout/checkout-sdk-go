package common_test

import (
	"encoding/json"
	stderrors "errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v3/accounts"
	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/errors"
	eventsabc "github.com/checkout/checkout-sdk-go/v3/events/abc"
	instruments "github.com/checkout/checkout-sdk-go/v3/instruments/nas"
	issuingcards "github.com/checkout/checkout-sdk-go/v3/issuing/cards"
	issuingcontrols "github.com/checkout/checkout-sdk-go/v3/issuing/controls"
	"github.com/checkout/checkout-sdk-go/v3/onboardingsimulator"
	"github.com/checkout/checkout-sdk-go/v3/payments"
	paymentsabc "github.com/checkout/checkout-sdk-go/v3/payments/abc"
	"github.com/checkout/checkout-sdk-go/v3/payments/contexts"
	paymentsnas "github.com/checkout/checkout-sdk-go/v3/payments/nas"
	"github.com/checkout/checkout-sdk-go/v3/payments/setups"
	webhooksabc "github.com/checkout/checkout-sdk-go/v3/webhooks/abc"
	workflowevents "github.com/checkout/checkout-sdk-go/v3/workflows/events"
)

const decoderMarker = "body_marker_9c1e"

// Every type of the SDK with its own UnmarshalJSON (git grep -n 'func (.*) UnmarshalJSON').
func customDecoders() map[string]func() interface{} {
	return map[string]func() interface{}{
		"accounts.PayoutSchedule":                      func() interface{} { return new(accounts.PayoutSchedule) },
		"common.APIShortDate":                          func() interface{} { return new(common.APIShortDate) },
		"events/abc.EventTypesResponse":                func() interface{} { return new(eventsabc.EventTypesResponse) },
		"instruments/nas.CreateInstrumentResponse":     func() interface{} { return new(instruments.CreateInstrumentResponse) },
		"instruments/nas.GetInstrumentResponse":        func() interface{} { return new(instruments.GetInstrumentResponse) },
		"instruments/nas.UpdateInstrumentResponse":     func() interface{} { return new(instruments.UpdateInstrumentResponse) },
		"issuing/cards.CardDetailsResponse":            func() interface{} { return new(issuingcards.CardDetailsResponse) },
		"issuing/controls.CardControlResponse":         func() interface{} { return new(issuingcontrols.CardControlResponse) },
		"onboardingsimulator.AvailableRequirements":    func() interface{} { return new(onboardingsimulator.SimulatorAvailableRequirementsResponse) },
		"onboardingsimulator.ScenariosResponse":        func() interface{} { return new(onboardingsimulator.SimulatorScenariosResponse) },
		"payments/abc.DestinationResponse":             func() interface{} { return new(paymentsabc.DestinationResponse) },
		"payments/abc.GetPaymentActionsResponse":       func() interface{} { return new(paymentsabc.GetPaymentActionsResponse) },
		"payments/abc.SourceResponse":                  func() interface{} { return new(paymentsabc.SourceResponse) },
		"payments.AirlineData":                         func() interface{} { return new(payments.AirlineData) },
		"payments/contexts.PaymentContextsAirlineData": func() interface{} { return new(contexts.PaymentContextsAirlineData) },
		"payments/nas.DestinationResponse":             func() interface{} { return new(paymentsnas.DestinationResponse) },
		"payments/nas.GetPaymentActionsResponse":       func() interface{} { return new(paymentsnas.GetPaymentActionsResponse) },
		"payments/nas.SenderResponse":                  func() interface{} { return new(paymentsnas.SenderResponse) },
		"payments/nas.ResponseCardSource":              func() interface{} { return new(paymentsnas.ResponseCardSource) },
		"payments/nas.SourceResponse":                  func() interface{} { return new(paymentsnas.SourceResponse) },
		"payments/setups.BizumPaymentMethod":           func() interface{} { return new(setups.BizumPaymentMethod) },
		"payments/setups.InstrumentPaymentMethod":      func() interface{} { return new(setups.InstrumentPaymentMethod) },
		"payments/setups.MultibancoPaymentMethod":      func() interface{} { return new(setups.MultibancoPaymentMethod) },
		"payments/setups.P24PaymentMethod":             func() interface{} { return new(setups.P24PaymentMethod) },
		"payments/setups.SwishPaymentMethod":           func() interface{} { return new(setups.SwishPaymentMethod) },
		"payments/setups.AchPaymentMethod":             func() interface{} { return new(setups.AchPaymentMethod) },
		"payments/setups.SepaPaymentMethod":            func() interface{} { return new(setups.SepaPaymentMethod) },
		"payments/setups.PayByBankPaymentMethod":       func() interface{} { return new(setups.PayByBankPaymentMethod) },
		"payments/setups.StablecoinPaymentMethod":      func() interface{} { return new(setups.StablecoinPaymentMethod) },
		"payments/setups.BlikPaymentMethod":            func() interface{} { return new(setups.BlikPaymentMethod) },
		"payments/setups.IdealPaymentMethod":           func() interface{} { return new(setups.IdealPaymentMethod) },
		"payments/setups.KnetPaymentMethod":            func() interface{} { return new(setups.KnetPaymentMethod) },
		"payments/setups.BancontactPaymentMethod":      func() interface{} { return new(setups.BancontactPaymentMethod) },
		"webhooks/abc.WebhooksResponse":                func() interface{} { return new(webhooksabc.WebhooksResponse) },
		"workflows/events.EventTypesResponse":          func() interface{} { return new(workflowevents.EventTypesResponse) },
	}
}

var decoderBodies = []string{
	`null`, `{}`, `[]`, `[null]`, `[{}]`, `[[]]`, `""`, `"` + decoderMarker + `"`, `0`, `-1`, `true`,
	`{"type":null}`, `{"type":1}`, `{"type":"` + decoderMarker + `"}`, `{"type":""}`,
	`{"type":"card"}`, `{"type":"card","source":null}`, `{"control_type":null}`,
	`{"control_type":"velocity_limit"}`, `{"control_type":"mcc_limit"}`,
	`{"data":null}`, `{"data":[null]}`, `{"data":{}}`, `{"items":null}`, `{"items":[null]}`,
	`{"passenger":null}`, `{"passenger":[null]}`, `{"passenger":"` + decoderMarker + `"}`, `{"passenger":{}}`,
	`{"initialization":null}`, `{"initialization":"` + decoderMarker + `"}`,
	`{"frequency":null}`, `{"frequency":"` + decoderMarker + `"}`,
	`{"actions":null}`, `{"_links":null}`, `{"requirements":null}`, `{"scenarios":null}`,
	`{"type":"`, `{`, `[`, `nul`,
}

func TestEveryCustomDecoderAcceptsAnyBody(t *testing.T) {
	for name, newValue := range customDecoders() {
		for _, body := range decoderBodies {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("%s: json.Unmarshal(%s): unexpected %v", name, body, r)
					}
				}()
				_ = json.Unmarshal([]byte(body), newValue())
			}()
		}
	}
}

func TestEveryCustomDecoderThroughUnmarshalGivesAResultOrAnSdkError(t *testing.T) {
	for name, newValue := range customDecoders() {
		for _, body := range decoderBodies {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("%s: common.Unmarshal(%s): unexpected %v", name, body, r)
					}
				}()
				header := http.Header{}
				header.Set("Content-Type", "application/json")
				metadata := &common.HttpMetadata{StatusCode: http.StatusOK, ResponseBody: []byte(body), Headers: &common.Headers{Header: header}}
				err := common.Unmarshal(metadata, newValue())
				if err == nil {
					return
				}
				var decodingErr errors.CheckoutDecodingError
				assert.True(t, stderrors.As(err, &decodingErr), "%s %s: got %T", name, body, err)
				assert.NotContains(t, err.Error(), decoderMarker, "%s %s", name, body)
				assert.False(t, strings.Contains(fmt.Sprintf("%v", err), decoderMarker), "%s %s", name, body)
			}()
		}
	}
}

func TestEveryCustomDecoderAcceptsNull(t *testing.T) {
	for name, newValue := range customDecoders() {
		assert.NoError(t, json.Unmarshal([]byte(`null`), newValue()), name)
		assert.NoError(t, json.Unmarshal([]byte(`[null]`), &[]interface{}{newValue()}), name)
	}
}

func TestNestedCardDetailsAcceptNull(t *testing.T) {
	var list []issuingcards.CardDetailsResponse
	assert.NoError(t, json.Unmarshal([]byte(`[null]`), &list))
	assert.Len(t, list, 1)
}

func TestTypedDecodersReturnTheErrorOfTheirVariant(t *testing.T) {
	cases := map[string]struct {
		newValue func() interface{}
		body     string
	}{
		"get instrument card":         {func() interface{} { return new(instruments.GetInstrumentResponse) }, `{"type":"card","id":987654321}`},
		"get instrument bank account": {func() interface{} { return new(instruments.GetInstrumentResponse) }, `{"type":"bank_account","id":1}`},
		"get instrument sepa":         {func() interface{} { return new(instruments.GetInstrumentResponse) }, `{"type":"sepa","id":1}`},
		"get instrument ach":          {func() interface{} { return new(instruments.GetInstrumentResponse) }, `{"type":"ach","id":1}`},
		"get instrument bacs":         {func() interface{} { return new(instruments.GetInstrumentResponse) }, `{"type":"bacs","id":1}`},
		"get instrument other":        {func() interface{} { return new(instruments.GetInstrumentResponse) }, `{"type":"other","x":[}`},
		"create instrument card":      {func() interface{} { return new(instruments.CreateInstrumentResponse) }, `{"type":"card","id":1}`},
		"create instrument bank":      {func() interface{} { return new(instruments.CreateInstrumentResponse) }, `{"type":"bank_account","id":1}`},
		"create instrument sepa":      {func() interface{} { return new(instruments.CreateInstrumentResponse) }, `{"type":"sepa","id":1}`},
		"create instrument ach":       {func() interface{} { return new(instruments.CreateInstrumentResponse) }, `{"type":"ach","id":1}`},
		"create instrument bacs":      {func() interface{} { return new(instruments.CreateInstrumentResponse) }, `{"type":"bacs","id":1}`},
		"create instrument token":     {func() interface{} { return new(instruments.CreateInstrumentResponse) }, `{"type":"token","id":1}`},
		"update instrument card":      {func() interface{} { return new(instruments.UpdateInstrumentResponse) }, `{"type":"card","id":1}`},
		"update instrument bank":      {func() interface{} { return new(instruments.UpdateInstrumentResponse) }, `{"type":"bank_account","id":1}`},
		"update instrument sepa":      {func() interface{} { return new(instruments.UpdateInstrumentResponse) }, `{"type":"sepa","id":1}`},
		"update instrument ach":       {func() interface{} { return new(instruments.UpdateInstrumentResponse) }, `{"type":"ach","id":1}`},
		"update instrument bacs":      {func() interface{} { return new(instruments.UpdateInstrumentResponse) }, `{"type":"bacs","id":1}`},
		"card details virtual":        {func() interface{} { return new(issuingcards.CardDetailsResponse) }, `{"type":"virtual","is_single_use":"x"}`},
		"card control velocity limit": {func() interface{} { return new(issuingcontrols.CardControlResponse) }, `{"control_type":"velocity_limit","velocity_limit":"x"}`},
		"card control mcc limit":      {func() interface{} { return new(issuingcontrols.CardControlResponse) }, `{"control_type":"mcc_limit","mcc_limit":"x"}`},
		"card control mid limit":      {func() interface{} { return new(issuingcontrols.CardControlResponse) }, `{"control_type":"mid_limit","mid_limit":"x"}`},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Error(t, json.Unmarshal([]byte(tc.body), tc.newValue()))

			header := http.Header{}
			header.Set("Content-Type", "application/json")
			metadata := &common.HttpMetadata{StatusCode: http.StatusOK, ResponseBody: []byte(tc.body), Headers: &common.Headers{Header: header}}
			assert.IsType(t, errors.CheckoutDecodingError{}, common.Unmarshal(metadata, tc.newValue()))
		})
	}
}
