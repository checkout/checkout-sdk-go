package test

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go-beta"
	"github.com/checkout/checkout-sdk-go-beta/accounts"
	"github.com/checkout/checkout-sdk-go-beta/common"
	"github.com/checkout/checkout-sdk-go-beta/configuration"
	"github.com/checkout/checkout-sdk-go-beta/errors"
	"github.com/checkout/checkout-sdk-go-beta/nas"
)

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
		entityId = createEntity()
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
		entityId = createEntity()
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

func createEntity() string {
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

	entity, _ := OAuthApi().Accounts.CreateEntity(r)

	return entity.Id
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

	client := buildAccountsClient().Accounts

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
	client := buildAccountsClient().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RetrievePayoutSchedule(tc.entityId))
		})
	}
}

func TestUploadFileAccounts(t *testing.T) {
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

	client := buildAccountsClient().Accounts

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UploadFile(tc.fileRequest))
		})
	}
}

func buildAccountsClient() *nas.Api {
	oauthAccountsApi, _ := checkout.Builder().OAuth().
		WithClientCredentials(
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_PAYOUT_SCHEDULE_CLIENT_ID"),
			os.Getenv("CHECKOUT_DEFAULT_OAUTH_PAYOUT_SCHEDULE_CLIENT_SECRET")).
		WithEnvironment(configuration.Sandbox()).
		WithScopes([]string{configuration.Marketplace, configuration.Files}).
		Build()

	return oauthAccountsApi
}
