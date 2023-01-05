package test

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go"
	"github.com/checkout/checkout-sdk-go/accounts"
	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/instruments"
	"github.com/checkout/checkout-sdk-go/nas"
)

func TestSubmitFileAccounts(t *testing.T) {
	cases := []struct {
		name        string
		fileRequest accounts.File
		checker     func(*common.IdResponse, error)
	}{
		{
			name: "when data is correct then return ID for uploaded file - IMAGE",
			fileRequest: accounts.File{
				File:    "./checkout.jpeg",
				Purpose: common.BankVerification,
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name: "when data is correct then return ID for uploaded file - PDF",
			fileRequest: accounts.File{
				File:    "./checkout.pdf",
				Purpose: common.BankVerification,
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name:        "when file path is missing then return error",
			fileRequest: accounts.File{},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				assert.Equal(t, "Invalid file name", err.Error())
			},
		},
		{
			name: "when purpose is missing then return error",
			fileRequest: accounts.File{
				File: "./checkout.pdf",
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				assert.Equal(t, "Invalid purpose", err.Error())
			},
		},
	}

	client := buildPayoutsScheduleClient().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.SubmitFile(tc.fileRequest))
		})
	}
}

func TestCreateEntity(t *testing.T) {
	cases := []struct {
		name    string
		request accounts.OnboardEntityRequest
		checker func(*accounts.OnboardEntityResponse, error)
	}{
		{
			name: "when request is correct then create entity",
			request: accounts.OnboardEntityRequest{
				Reference:      GenerateRandomReference(),
				ContactDetails: &accounts.ContactDetails{Phone: &accounts.Phone{Number: "2345678910"}},
				Profile: &accounts.Profile{
					Urls: []string{"https://www.superheroexample.com"},
					Mccs: []string{"0742"},
				},
				Individual: &accounts.Individual{
					FirstName:         "Bruce",
					LastName:          "Wayne",
					TradingName:       "Batman's Super Hero Masks",
					NationalTaxId:     "TAX123456",
					RegisteredAddress: Address(),
					DateOfBirth:       &accounts.DateOfBirth{Day: 5, Month: 6, Year: 1995},
					Identification:    &accounts.Identification{NationalIdNumber: "AB123456C"},
				},
			},
			checker: func(response *accounts.OnboardEntityResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:    "when request is not correct then return error",
			request: accounts.OnboardEntityRequest{},
			checker: func(response *accounts.OnboardEntityResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "company_or_individual_required")
			},
		},
	}

	client := OAuthApi().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreateEntity(tc.request))
		})
	}
}

func TestGetEntity(t *testing.T) {
	var (
		entityId = createEntity(t)
	)

	cases := []struct {
		name     string
		entityId string
		checker  func(*accounts.OnboardEntityDetails, error)
	}{
		{
			name:     "when entity exists then return entity details",
			entityId: entityId,
			checker: func(response *accounts.OnboardEntityDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, entityId, response.Id)
			},
		},
		{
			name:     "when entity does not exist then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			checker: func(response *accounts.OnboardEntityDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}
	client := OAuthApi().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetEntity(tc.entityId))
		})
	}
}

func TestUpdateEntity(t *testing.T) {
	var (
		entityId = createEntity(t)
	)

	cases := []struct {
		name     string
		entityId string
		request  accounts.OnboardEntityRequest
		checker  func(*accounts.OnboardEntityResponse, error)
	}{
		{
			name:     "when request is correct then update entity",
			entityId: entityId,
			request: accounts.OnboardEntityRequest{
				ContactDetails: &accounts.ContactDetails{Phone: &accounts.Phone{Number: "2345678910"}},
				Profile: &accounts.Profile{
					Urls: []string{"https://www.superheroexample.com"},
					Mccs: []string{"0742"},
				},
				Individual: &accounts.Individual{
					FirstName:         "New Name",
					LastName:          "New LastName",
					TradingName:       "New Trading Name",
					NationalTaxId:     "TAX8765432",
					RegisteredAddress: Address(),
					DateOfBirth:       &accounts.DateOfBirth{Day: 5, Month: 6, Year: 1995},
					Identification:    &accounts.Identification{NationalIdNumber: "AB123456C"},
				},
			},
			checker: func(response *accounts.OnboardEntityResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:     "when entity not_found then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			request: accounts.OnboardEntityRequest{
				Individual: &accounts.Individual{
					FirstName:         "New Name",
					LastName:          "New LastName",
					TradingName:       "New Trading Name",
					NationalTaxId:     "TAX8765432",
					RegisteredAddress: Address(),
					DateOfBirth:       &accounts.DateOfBirth{Day: 5, Month: 6, Year: 1995},
					Identification:    &accounts.Identification{NationalIdNumber: "AB123456C"},
				},
			},
			checker: func(response *accounts.OnboardEntityResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := OAuthApi().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdateEntity(tc.entityId, tc.request))
		})
	}
}

func TestQueryPaymentInstruments(t *testing.T) {
	var (
		entityId = createEntityCompany(t)
		query    = accounts.PaymentInstrumentsQuery{
			Status: accounts.Unverified,
		}
	)

	cases := []struct {
		name     string
		entityId string
		query    accounts.PaymentInstrumentsQuery
		checker  func(*accounts.PaymentInstrumentQueryResponse, error)
	}{
		{
			name:     "when query a payment instrument then return payment instrument data",
			entityId: entityId,
			query:    query,
			checker: func(response *accounts.PaymentInstrumentQueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Data)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:     "when entity is incorrect then return error",
			entityId: "",
			checker: func(response *accounts.PaymentInstrumentQueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
			},
		},
	}
	client := buildAccountClient().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.QueryPaymentInstruments(tc.entityId, tc.query))
		})
	}
}

func TestRetrievePaymentInstrumentsDetails(t *testing.T) {
	var (
		entityId = createEntityCompany(t)

		file = accounts.File{
			File:    "./checkout.pdf",
			Purpose: common.BankVerification,
		}

		requestFileId = submitFile(t, file)

		paymentInstrumentId = paymentInstrumentRequest(t, entityId, requestFileId)
	)

	cases := []struct {
		name                string
		entityId            string
		paymentInstrumentId string
		checker             func(*accounts.PaymentInstrumentDetailsResponse, error)
	}{
		{
			name:                "when fetching valid payment instrument details then return payment instrument details",
			entityId:            entityId,
			paymentInstrumentId: paymentInstrumentId,
			checker: func(response *accounts.PaymentInstrumentDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.Label)
				assert.NotNil(t, response.Type)
				assert.NotNil(t, response.Currency)
				assert.NotNil(t, response.Country)
				assert.NotNil(t, response.Document)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name:                "when entity does not exist then return error",
			entityId:            "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			paymentInstrumentId: "ppi_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			checker: func(response *accounts.PaymentInstrumentDetailsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}
	client := buildAccountClient().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrievePaymentInstrumentDetails(tc.entityId, tc.paymentInstrumentId))
		})
	}
}

func TestUpdatePayoutSchedule(t *testing.T) {
	var (
		entityId = "ent_t2jwrwxhxdas5755cnctu7iwmm"
	)

	cases := []struct {
		name           string
		entityId       string
		currency       common.Currency
		request        accounts.CurrencySchedule
		checkerRequest func(*common.IdResponse, error)
		checkerInfo    func(*accounts.PayoutSchedule, error)
	}{
		{
			name:     "when request for daily frequency schedule is correct then update entity",
			entityId: entityId,
			currency: common.USD,
			request: accounts.CurrencySchedule{
				Enabled:    true,
				Threshold:  500,
				Recurrence: accounts.NewScheduleFrequencyDailyRequest(),
			},
			checkerRequest: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
			checkerInfo: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Currency)
				assert.IsType(t, accounts.NewScheduleFrequencyDailyRequest(), response.Currency[common.USD].Recurrence)
			},
		},
		{
			name:     "when request for weekly frequency schedule is correct then update entity",
			entityId: entityId,
			currency: common.USD,
			request: accounts.CurrencySchedule{
				Enabled:    true,
				Threshold:  1000,
				Recurrence: accounts.NewScheduleFrequencyWeeklyRequest([]accounts.DaySchedule{accounts.Monday}),
			},
			checkerRequest: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
			checkerInfo: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Currency)
				assert.IsType(t,
					accounts.NewScheduleFrequencyWeeklyRequest([]accounts.DaySchedule{accounts.Monday}),
					response.Currency[common.USD].Recurrence)
			},
		},
		{
			name:     "when request for monthly frequency schedule is correct then update entity",
			entityId: entityId,
			currency: common.USD,
			request: accounts.CurrencySchedule{
				Enabled:    true,
				Threshold:  1500,
				Recurrence: accounts.NewScheduleFrequencyMonthlyRequest([]int{5}),
			},
			checkerRequest: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
			checkerInfo: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, response.Currency)
				assert.IsType(t,
					accounts.NewScheduleFrequencyMonthlyRequest([]int{5}),
					response.Currency[common.USD].Recurrence)
			},
		},
		{
			name:     "when entity not_found then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			currency: common.USD,
			request:  accounts.CurrencySchedule{},
			checkerRequest: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
			checkerInfo: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	client := buildPayoutsScheduleClient().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checkerRequest(client.UpdatePayoutSchedule(tc.entityId, tc.currency, tc.request))

			tc.checkerInfo(client.RetrievePayoutSchedule(tc.entityId))
		})
	}
}

func TestGetPayoutSchedule(t *testing.T) {
	var (
		dailyEntity   = "ent_sdioy6bajpzxyl3utftdp7legq"
		weeklyEntity  = "ent_yvt7y275h6iu4diq4s6gxxepfm"
		monthlyEntity = "ent_224gcrnxtugb2hlqo62w625i6m"
	)

	cases := []struct {
		name     string
		entityId string
		checker  func(*accounts.PayoutSchedule, error)
	}{
		{
			name:     "when entity with daily schedule exists then return entity's payout schedule",
			entityId: dailyEntity,
			checker: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Currency[common.USD])
				assert.True(t, response.Currency[common.USD].Enabled)
				assert.Equal(t, 1000, response.Currency[common.USD].Threshold)
				assert.Equal(t, accounts.NewScheduleFrequencyDailyRequest(), response.Currency[common.USD].Recurrence)
			},
		},
		{
			name:     "when entity with weekly schedule exists then return entity's payout schedule",
			entityId: weeklyEntity,
			checker: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Currency[common.USD])
				assert.True(t, response.Currency[common.USD].Enabled)
				assert.Equal(t, 1000, response.Currency[common.USD].Threshold)
				assert.Equal(t,
					accounts.NewScheduleFrequencyWeeklyRequest([]accounts.DaySchedule{accounts.Wednesday}),
					response.Currency[common.USD].Recurrence)
			},
		},
		{
			name:     "when entity with monthly schedule exists then return entity's payout schedule",
			entityId: monthlyEntity,
			checker: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Currency[common.USD])
				assert.True(t, response.Currency[common.USD].Enabled)
				assert.Equal(t, 1000, response.Currency[common.USD].Threshold)
				assert.Equal(t,
					accounts.NewScheduleFrequencyMonthlyRequest([]int{15}),
					response.Currency[common.USD].Recurrence)
			},
		},
		{
			name:     "when entity does not exist then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			checker: func(response *accounts.PayoutSchedule, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}
	client := buildPayoutsScheduleClient().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrievePayoutSchedule(tc.entityId))
		})
	}
}

func createEntity(t *testing.T) string {
	r := accounts.OnboardEntityRequest{
		Reference:      GenerateRandomReference(),
		ContactDetails: &accounts.ContactDetails{Phone: &accounts.Phone{Number: "2345678910"}},
		Profile: &accounts.Profile{
			Urls: []string{"https://www.superheroexample.com"},
			Mccs: []string{"0742"},
		},
		Individual: &accounts.Individual{
			FirstName:         "Bruce",
			LastName:          "Wayne",
			TradingName:       "Batman's Super Hero Masks",
			NationalTaxId:     "TAX123456",
			RegisteredAddress: Address(),
			DateOfBirth:       &accounts.DateOfBirth{Day: 5, Month: 6, Year: 1995},
			Identification:    &accounts.Identification{NationalIdNumber: "AB123456C"},
		},
	}

	entity, err := OAuthApi().Accounts.CreateEntity(r)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating entity - %s", err.Error()))
	}

	return entity.Id
}

func createEntityCompany(t *testing.T) string {
	r := accounts.OnboardEntityRequest{
		Reference: GenerateRandomReference(),
		ContactDetails: &accounts.ContactDetails{
			Phone: &accounts.Phone{
				Number: "2345678910",
			},
			EntityEmailAddresses: &accounts.EntityEmailAddresses{
				Primary: []string{GenerateRandomEmail()},
			},
		},
		Profile: &accounts.Profile{
			Urls: []string{"https://www.superheroexample.com"},
			Mccs: []string{"0742"},
		},

		Company: &accounts.Company{
			BusinessRegistrationNumber: "01234567",
			BusinessType:               "",
			LegalName:                  "Super Hero Masks Inc.",
			TradingName:                "Super Hero Masks",
			PrincipalAddress:           Address(),
			RegisteredAddress:          Address(),
			Representatives: []accounts.Representative{
				{
					FirstName: "John",
					LastName:  "Doe",
					Address:   Address(),
					Identification: &accounts.Identification{
						NationalIdNumber: "AB123456C",
					},
				},
			},
		},
	}

	entity, err := buildAccountClient().Accounts.CreateEntity(r)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating entity company - %s", err.Error()))
	}

	return entity.Id
}

func submitFile(t *testing.T, fileRequest accounts.File) string {
	file, err := OAuthApi().Accounts.SubmitFile(fileRequest)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error uploading file - %s", err.Error()))
	}

	return file.Id
}

func paymentInstrumentRequest(t *testing.T, entityId string, fileId string) string {
	instrumentDocument := accounts.InstrumentDocument{
		Type:   "bank_statement",
		FileId: fileId,
	}

	instrumentDetails := accounts.InstrumentDetailsFasterPayments{
		AccountNumber: "12334454",
		BankCode:      "050389",
	}

	paymentInstrument := accounts.PaymentInstrumentRequest{
		Label:              "Barclays",
		Type:               instruments.BankAccount,
		Currency:           common.GBP,
		Country:            common.GB,
		DefaultDestination: false,
		Document:           &instrumentDocument,
		InstrumentDetails:  &instrumentDetails,
	}

	instrumentResponse, err := buildAccountClient().Accounts.CreatePaymentInstrument(entityId, paymentInstrument)
	if err != nil {
		assert.Fail(t, fmt.Sprintf("error creating payment instrument - %s", err.Error()))
	}

	return instrumentResponse.Id
}

func buildPayoutsScheduleClient() *nas.Api {
	oauthPayoutsScheduleApi, _ := checkout.Builder().OAuth().
		WithClientCredentials(
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_PAYOUT_SCHEDULE_CLIENT_ID"),
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_PAYOUT_SCHEDULE_CLIENT_SECRET")).
		WithEnvironment(configuration.Sandbox()).
		WithScopes([]string{configuration.Marketplace, configuration.Files}).
		Build()

	return oauthPayoutsScheduleApi
}

func buildAccountClient() *nas.Api {
	oauthAccountsApi, _ := checkout.Builder().OAuth().
		WithClientCredentials(
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_ACCOUNTS_CLIENT_ID"),
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_ACCOUNTS_CLIENT_SECRET")).
		WithEnvironment(configuration.Sandbox()).
		WithScopes([]string{configuration.Accounts}).
		Build()

	return oauthAccountsApi
}
