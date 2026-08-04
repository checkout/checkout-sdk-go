package sessions

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/sessions/sources"
)

// Spec-conformance guards for the sessions enums. These types carry the raw wire values sent to and
// returned by the API, so a typo is invisible at compile time and only fails against the live API.

// validAPIValue matches snake_case, or a single uppercase letter for the Y/N/U style codes.
var validAPIValue = regexp.MustCompile(`^([a-z0-9_]+|[A-Z])$`)

func TestCategory_MatchesSpec(t *testing.T) {
	// Spec: ["payment", "non_payment"]. Guards against the camelCase "nonPayment", which the API
	// rejects.
	assert.Equal(t, "payment", string(Payment))
	assert.Equal(t, "non_payment", string(NonPayment))
}

func TestTransactionType_MatchesSpec(t *testing.T) {
	// Spec: the five values below. The correct spelling is "quasi_card_transaction".
	assert.Equal(t, "account_funding", string(AccountFunding))
	assert.Equal(t, "check_acceptance", string(CheckAcceptance))
	assert.Equal(t, "goods_service", string(GoodsService))
	assert.Equal(t, "prepaid_activation_and_load", string(PrepaidActivationAndLoad))
	assert.Equal(t, "quasi_card_transaction", string(QuasiCardTransaction))
}

func TestShippingIndicator_MatchesSpec(t *testing.T) {
	// Spec: the seven values below. This type previously held a single wrong value, "visa", which
	// left MerchantRiskInfo.ShippingIndicator unusable.
	expected := []string{
		"billing_address",
		"another_address_on_file",
		"not_on_file",
		"store_pick_up",
		"digital_goods",
		"travel_and_event_no_shipping",
		"other",
	}
	actual := []string{
		string(ShippingBillingAddress),
		string(ShippingAnotherAddressOnFile),
		string(ShippingNotOnFile),
		string(ShippingStorePickUp),
		string(ShippingDigitalGoods),
		string(ShippingTravelAndEventNoShipping),
		string(ShippingOther),
	}

	assert.Equal(t, expected, actual)
	assert.NotContains(t, actual, "visa")
}

func TestThreeDsMethodCompletion_IsUppercase(t *testing.T) {
	// Spec enum is Y/N/U with minLength and maxLength 1. Lowercase values are rejected by the API.
	assert.Equal(t, "Y", string(common.Y))
	assert.Equal(t, "N", string(common.N))
	assert.Equal(t, "U", string(common.U))
}

func TestSessionScheme_MatchesSpec(t *testing.T) {
	// Spec: eight values. "discover" and "upi" were previously missing.
	expected := []string{
		"amex", "cartes_bancaires", "diners", "discover",
		"jcb", "mastercard", "upi", "visa",
	}
	actual := []string{
		string(sources.Amex), string(sources.CartesBancaires), string(sources.Diners),
		string(sources.Discover), string(sources.Jcb), string(sources.Mastercard),
		string(sources.Upi), string(sources.Visa),
	}

	assert.Equal(t, expected, actual)
	assert.Len(t, actual, 8)
}

func TestExperienceStatus_MatchesSpec(t *testing.T) {
	// Spec: PreferredExperiences status enum is the four values below.
	expected := []string{"available", "unprocessed", "processed", "unavailable"}
	actual := []string{
		string(ExperienceAvailable), string(ExperienceUnprocessed),
		string(ExperienceProcessed), string(ExperienceUnavailable),
	}

	assert.Equal(t, expected, actual)
}

// The 3ds, preferred_experiences, experience and google_spa fields of GetSessionResponse were
// previously unmodelled, so the Google SPA and preferred-experience data the API returns was dropped.
func TestSessionDetails_DeserializesTheExperienceFields(t *testing.T) {
	payload := []byte(`{
		"3ds":{"challenge_request":"creq","interaction_counter":"03",
			"error_details":{"error_code":"101","error_component":"D","error_detail":"acctNumber",
			"error_description":"missing"}},
		"preferred_experiences":{"google_spa":{"status":"available"},
			"3ds":{"status":"processed","reason":["Invalid response"]}},
		"experience":"3ds",
		"google_spa":{"challenge_url":"https://google.example/challenge","initial_timeout":"5",
			"max_timeout":"10","iframe":{"height":"400","width":"250"},
			"token":{"number":"4242","expiry_month":12,"expiry_year":2030}}
	}`)

	var details SessionDetails
	assert.Nil(t, json.Unmarshal(payload, &details))

	assert.NotNil(t, details.ThreeDs)
	assert.Equal(t, "creq", details.ThreeDs.ChallengeRequest)
	assert.Equal(t, "03", details.ThreeDs.InteractionCounter)
	assert.NotNil(t, details.ThreeDs.ErrorDetails)
	assert.Equal(t, "101", details.ThreeDs.ErrorDetails.ErrorCode)
	assert.Equal(t, "D", details.ThreeDs.ErrorDetails.ErrorComponent)

	assert.NotNil(t, details.PreferredExperiences)
	assert.NotNil(t, details.PreferredExperiences.GoogleSpa)
	assert.Equal(t, ExperienceAvailable, details.PreferredExperiences.GoogleSpa.Status)
	assert.NotNil(t, details.PreferredExperiences.ThreeDs)
	assert.Equal(t, ExperienceProcessed, details.PreferredExperiences.ThreeDs.Status)
	assert.Equal(t, []string{"Invalid response"}, details.PreferredExperiences.ThreeDs.Reason)

	assert.Equal(t, ThreeDsExperience, details.Experience)

	assert.NotNil(t, details.GoogleSpa)
	assert.Equal(t, "https://google.example/challenge", details.GoogleSpa.ChallengeUrl)
	assert.Equal(t, "5", details.GoogleSpa.InitialTimeout)
	assert.NotNil(t, details.GoogleSpa.Iframe)
	assert.Equal(t, "400", details.GoogleSpa.Iframe.Height)
	assert.NotNil(t, details.GoogleSpa.Token)
	assert.Equal(t, "4242", details.GoogleSpa.Token.Number)
	assert.Equal(t, 12, details.GoogleSpa.Token.ExpiryMonth)
}

// Structural guard: every wire value used by the sessions surface must be snake_case or a single
// uppercase code. This catches camelCase leaking into a value, which is how "nonPayment" survived.
func TestEverySessionsWireValueIsWellFormed(t *testing.T) {
	values := []string{
		string(Payment), string(NonPayment),
		string(AccountFunding), string(CheckAcceptance), string(GoodsService),
		string(PrepaidActivationAndLoad), string(QuasiCardTransaction),
		string(ShippingBillingAddress), string(ShippingAnotherAddressOnFile),
		string(ShippingNotOnFile), string(ShippingStorePickUp), string(ShippingDigitalGoods),
		string(ShippingTravelAndEventNoShipping), string(ShippingOther),
		string(common.Y), string(common.N), string(common.U),
		string(sources.Amex), string(sources.CartesBancaires), string(sources.Diners),
		string(sources.Discover), string(sources.Jcb), string(sources.Mastercard),
		string(sources.Upi), string(sources.Visa),
		string(ThreeDsExperience), string(GoogleSpaExperience),
	}
	values = append(values, sessionChallengeIndicatorWireValues...)

	assert.Greater(t, len(values), 30)
	for _, value := range values {
		assert.Regexp(t, validAPIValue, value, "%q is not a valid API value", value)
	}
}
