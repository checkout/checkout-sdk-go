package balances

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v3/common"
	"github.com/checkout/checkout-sdk-go/v3/configuration"
	"github.com/checkout/checkout-sdk-go/v3/errors"
	"github.com/checkout/checkout-sdk-go/v3/mocks"
)

const (
	testEntityId          = "ent_w4jelhppmfiufdnatam37wrfc4"
	testCurrencyAccountId = "ca_g5y7d6jo4e2urgforcbf2ey5jm"
	// The exact path the SDK must build. common.BuildPath prefixes each segment with "/".
	expectedTopUpPath = "/entities/" + testEntityId + "/currency-accounts/" + testCurrencyAccountId +
		"/top-up-instructions"
)

// fullRailJson and the payloads below use the field-level "example" values from
// shared/swagger-latest.json. The response schema carries no top-level example.
const fullRailJson = `{
	"beneficiary_account_name": "Acme Inc",
	"beneficiary_address": "1 Example Street, Exampleville, EX, 00000, US",
	"bank_name": "Example Bank",
	"bank_address": "1 Example Street, Exampleville, EX, 00000, US",
	"account_number": "1234567890",
	"sort_code": "000000",
	"routing_number": "000000000",
	"iban": "GB00EXAM00000000000000",
	"swift_code": "TESTUS00XXX"
}`

func fullRail() *TopUpFundingDetails {
	return &TopUpFundingDetails{
		BeneficiaryAccountName: "Acme Inc",
		BeneficiaryAddress:     "1 Example Street, Exampleville, EX, 00000, US",
		BankName:               "Example Bank",
		BankAddress:            "1 Example Street, Exampleville, EX, 00000, US",
		AccountNumber:          "1234567890",
		SortCode:               "000000",
		RoutingNumber:          "000000000",
		Iban:                   "GB00EXAM00000000000000",
		SwiftCode:              "TESTUS00XXX",
	}
}

func assertFullRail(t *testing.T, rail *TopUpFundingDetails) {
	assert.NotNil(t, rail)
	assert.Equal(t, "Acme Inc", rail.BeneficiaryAccountName)
	assert.Equal(t, "1 Example Street, Exampleville, EX, 00000, US", rail.BeneficiaryAddress)
	assert.Equal(t, "Example Bank", rail.BankName)
	assert.Equal(t, "1 Example Street, Exampleville, EX, 00000, US", rail.BankAddress)
	assert.Equal(t, "1234567890", rail.AccountNumber)
	assert.Equal(t, "000000", rail.SortCode)
	assert.Equal(t, "000000000", rail.RoutingNumber)
	assert.Equal(t, "GB00EXAM00000000000000", rail.Iban)
	assert.Equal(t, "TESTUS00XXX", rail.SwiftCode)
}

func TestRetrieveTopUpInstructions(t *testing.T) {
	var response = TopUpInstructionsResponse{
		HttpMetadata:      mocks.HttpMetadataStatusOk,
		CurrencyAccountId: testCurrencyAccountId,
		Currency:          common.USD,
		PaymentReference:  "TP-ABC123",
		BankDetails: &TopUpBankDetails{
			Domestic:      fullRail(),
			International: fullRail(),
		},
	}

	cases := []struct {
		name              string
		entityId          string
		currencyAccountId string
		getAuthorization  func(*mock.Mock) mock.Call
		apiGet            func(*mock.Mock) mock.Call
		checker           func(*TopUpInstructionsResponse, error)
	}{
		{
			name:              "when request is correct then return top-up instructions",
			entityId:          testEntityId,
			currencyAccountId: testCurrencyAccountId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				// The path is asserted here, not with mock.Anything: a wrong path is the
				// most likely way this endpoint breaks.
				return *m.On("GetWithContext", mock.Anything, expectedTopUpPath, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*TopUpInstructionsResponse)
						*respMapping = response
					})
			},
			checker: func(response *TopUpInstructionsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, testCurrencyAccountId, response.CurrencyAccountId)
				assert.Equal(t, common.USD, response.Currency)
				assert.Equal(t, "TP-ABC123", response.PaymentReference)
				assert.NotNil(t, response.BankDetails)
				assertFullRail(t, response.BankDetails.Domestic)
				assertFullRail(t, response.BankDetails.International)
			},
		},
		{
			name:              "when credentials invalid then return error",
			entityId:          testEntityId,
			currencyAccountId: testCurrencyAccountId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *TopUpInstructionsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				assert.Equal(t, "Invalid authorization type", err.Error())
			},
		},
		{
			name:              "when top-ups are not enabled then return 403 error",
			entityId:          testEntityId,
			currencyAccountId: testCurrencyAccountId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusForbidden,
						Status:     "403 Forbidden",
					})
			},
			checker: func(response *TopUpInstructionsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusForbidden, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.RetrieveTopUpInstructions(tc.entityId, tc.currencyAccountId))
		})
	}
}

// --- response shape: both rails, single rail, empty bank_details ---------------
// The spec declares no `required` array on TopUpBankDetails, so domestic-only,
// international-only and an empty bank_details are all legal 200 bodies.

func TestTopUpInstructionsResponse_UnmarshalsBothRails(t *testing.T) {
	body := `{
		"currency_account_id": "` + testCurrencyAccountId + `",
		"currency": "USD",
		"payment_reference": "TP-ABC123",
		"bank_details": {"domestic": ` + fullRailJson + `, "international": ` + fullRailJson + `}
	}`

	var response TopUpInstructionsResponse
	assert.Nil(t, json.Unmarshal([]byte(body), &response))

	assert.Equal(t, testCurrencyAccountId, response.CurrencyAccountId)
	assert.Equal(t, common.USD, response.Currency)
	assert.Equal(t, "TP-ABC123", response.PaymentReference)
	assert.NotNil(t, response.BankDetails)
	assertFullRail(t, response.BankDetails.Domestic)
	assertFullRail(t, response.BankDetails.International)
}

func TestTopUpInstructionsResponse_UnmarshalsDomesticOnly(t *testing.T) {
	// A United States domestic rail, per "Returned for United States domestic transfers".
	body := `{
		"currency_account_id": "` + testCurrencyAccountId + `",
		"currency": "USD",
		"payment_reference": "TP-ABC123",
		"bank_details": {"domestic": {
			"beneficiary_account_name": "Acme Inc",
			"bank_name": "Example Bank",
			"account_number": "1234567890",
			"routing_number": "000000000"
		}}
	}`

	var response TopUpInstructionsResponse
	assert.Nil(t, json.Unmarshal([]byte(body), &response))

	assert.NotNil(t, response.BankDetails)
	assert.Nil(t, response.BankDetails.International, "international rail must be nil when absent")
	assert.NotNil(t, response.BankDetails.Domestic)
	assert.Equal(t, "000000000", response.BankDetails.Domestic.RoutingNumber)
	assert.Empty(t, response.BankDetails.Domestic.SortCode)
	assert.Empty(t, response.BankDetails.Domestic.Iban)
	assert.Empty(t, response.BankDetails.Domestic.SwiftCode)
}

func TestTopUpInstructionsResponse_UnmarshalsInternationalOnly(t *testing.T) {
	// An international rail, per "Returned for international transfers".
	body := `{
		"currency_account_id": "` + testCurrencyAccountId + `",
		"currency": "USD",
		"payment_reference": "TP-ABC123",
		"bank_details": {"international": {
			"beneficiary_account_name": "Acme Inc",
			"bank_name": "Example Bank",
			"iban": "GB00EXAM00000000000000",
			"swift_code": "TESTUS00XXX"
		}}
	}`

	var response TopUpInstructionsResponse
	assert.Nil(t, json.Unmarshal([]byte(body), &response))

	assert.NotNil(t, response.BankDetails)
	assert.Nil(t, response.BankDetails.Domestic, "domestic rail must be nil when absent")
	assert.NotNil(t, response.BankDetails.International)
	assert.Equal(t, "TESTUS00XXX", response.BankDetails.International.SwiftCode)
	assert.Empty(t, response.BankDetails.International.AccountNumber)
	assert.Empty(t, response.BankDetails.International.RoutingNumber)
}

func TestTopUpInstructionsResponse_UnmarshalsEmptyBankDetails(t *testing.T) {
	body := `{
		"currency_account_id": "` + testCurrencyAccountId + `",
		"currency": "USD",
		"payment_reference": "TP-ABC123",
		"bank_details": {}
	}`

	var response TopUpInstructionsResponse
	assert.Nil(t, json.Unmarshal([]byte(body), &response))

	assert.NotNil(t, response.BankDetails)
	assert.Nil(t, response.BankDetails.Domestic)
	assert.Nil(t, response.BankDetails.International)
}

func TestTopUpInstructionsResponse_RoundTrips(t *testing.T) {
	original := TopUpInstructionsResponse{
		CurrencyAccountId: testCurrencyAccountId,
		Currency:          common.USD,
		PaymentReference:  "TP-ABC123",
		BankDetails: &TopUpBankDetails{
			Domestic:      fullRail(),
			International: fullRail(),
		},
	}

	marshalled, err := json.Marshal(original)
	assert.Nil(t, err)

	// Wire names must be snake_case.
	body := string(marshalled)
	for _, key := range []string{
		`"currency_account_id"`, `"currency"`, `"payment_reference"`, `"bank_details"`,
		`"domestic"`, `"international"`, `"beneficiary_account_name"`, `"beneficiary_address"`,
		`"bank_name"`, `"bank_address"`, `"account_number"`, `"sort_code"`,
		`"routing_number"`, `"iban"`, `"swift_code"`,
	} {
		assert.Contains(t, body, key)
	}

	var deserialized TopUpInstructionsResponse
	assert.Nil(t, json.Unmarshal(marshalled, &deserialized))
	assert.Equal(t, original.CurrencyAccountId, deserialized.CurrencyAccountId)
	assert.Equal(t, original.Currency, deserialized.Currency)
	assert.Equal(t, original.PaymentReference, deserialized.PaymentReference)
	assertFullRail(t, deserialized.BankDetails.Domestic)
	assertFullRail(t, deserialized.BankDetails.International)
}

func TestTopUpFundingDetails_OmitsUnsetOptionalFields(t *testing.T) {
	marshalled, err := json.Marshal(TopUpBankDetails{
		Domestic: &TopUpFundingDetails{
			BeneficiaryAccountName: "Acme Inc",
			BankName:               "Example Bank",
		},
	})
	assert.Nil(t, err)

	body := string(marshalled)
	assert.Contains(t, body, `"beneficiary_account_name":"Acme Inc"`)
	assert.Contains(t, body, `"bank_name":"Example Bank"`)
	assert.NotContains(t, body, "international")
	assert.NotContains(t, body, "sort_code")
	assert.NotContains(t, body, "routing_number")
	assert.NotContains(t, body, "iban")
	assert.NotContains(t, body, "swift_code")
	assert.NotContains(t, body, "beneficiary_address")
	assert.NotContains(t, body, "bank_address")
	assert.NotContains(t, body, "account_number")
}
