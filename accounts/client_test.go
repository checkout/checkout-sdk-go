package accounts

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/instruments"
	"github.com/checkout/checkout-sdk-go/mocks"
)

func TestCreateEntity(t *testing.T) {
	var (
		onboardEntity = OnboardEntityResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "ent_1234",
			Reference:    "reference",
			Status:       Active,
			Capabilities: &Capabilities{
				Payments: &Payments{Available: true},
			},
		}
	)

	cases := []struct {
		name             string
		request          OnboardEntityRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*OnboardEntityResponse, error)
	}{
		{
			name: "when request is correct then create entity",
			request: OnboardEntityRequest{
				Reference:      "reference",
				ContactDetails: &ContactDetails{Phone: &Phone{Number: "2345678910"}},
				Profile: &Profile{
					Urls: []string{"https://www.superheroexample.com"},
					Mccs: []string{"0742"},
				},
				Individual: &Individual{
					FirstName:      "Bruce",
					LastName:       "Wayne",
					TradingName:    "Batman's Super Hero Masks",
					NationalTaxId:  "TAX123456",
					DateOfBirth:    &DateOfBirth{Day: 5, Month: 6, Year: 1995},
					Identification: &Identification{NationalIdNumber: "AB123456C"},
				},
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*OnboardEntityResponse)
						*respMapping = onboardEntity
					})
			},
			checker: func(response *OnboardEntityResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, onboardEntity.Id, response.Id)
				assert.Equal(t, onboardEntity.Reference, response.Reference)
				assert.Equal(t, onboardEntity.Status, response.Status)
				assert.Equal(t, onboardEntity.Capabilities, response.Capabilities)
			},
		},
		{
			name:    "when request is not correct then return error",
			request: OnboardEntityRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "invalid_request",
								ErrorCodes: []string{"company_or_individual_required"},
							},
						})
			},
			checker: func(response *OnboardEntityResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "company_or_individual_required")
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *OnboardEntityResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.CreateEntity(tc.request))
		})
	}
}

func TestGetEntity(t *testing.T) {
	var (
		entityId = "ent_1234"

		entityDetails = OnboardEntityDetails{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           entityId,
			Reference:    "reference",
			Capabilities: &Capabilities{
				Payments: &Payments{Available: true},
			},
			Status:         Active,
			ContactDetails: &ContactDetails{Phone: &Phone{Number: "2345678910"}},
			Profile: &Profile{
				Urls: []string{"https://www.superheroexample.com"},
				Mccs: []string{"0742"},
			},
			Individual: &Individual{
				FirstName:      "Bruce",
				LastName:       "Wayne",
				TradingName:    "Batman's Super Hero Masks",
				NationalTaxId:  "TAX123456",
				DateOfBirth:    &DateOfBirth{Day: 5, Month: 6, Year: 1995},
				Identification: &Identification{NationalIdNumber: "AB123456C"},
			},
		}
	)

	cases := []struct {
		name             string
		entityId         string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*OnboardEntityDetails, error)
	}{
		{
			name:     "when entity exists then return entity details",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*OnboardEntityDetails)
						*respMapping = entityDetails
					})
			},
			checker: func(response *OnboardEntityDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, entityId, response.Id)
				assert.Equal(t, entityDetails.Reference, response.Reference)
				assert.Equal(t, entityDetails.Reference, response.Reference)
				assert.Equal(t, entityDetails.Status, response.Status)
				assert.Equal(t, entityDetails.Capabilities, response.Capabilities)
				assert.Equal(t, entityDetails.ContactDetails, response.ContactDetails)
				assert.Equal(t, entityDetails.Profile, response.Profile)
				assert.Equal(t, entityDetails.Individual, response.Individual)
			},
		},
		{
			name:     "when entity does not exist then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *OnboardEntityDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *OnboardEntityDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.GetEntity(tc.entityId))
		})
	}
}

func TestUpdateEntity(t *testing.T) {
	var (
		onboardEntity = OnboardEntityResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "ent_1234",
			Reference:    "reference",
			Status:       Active,
			Capabilities: &Capabilities{
				Payments: &Payments{Available: true},
			},
		}
	)

	cases := []struct {
		name             string
		entityId         string
		request          OnboardEntityRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPut           func(*mock.Mock) mock.Call
		checker          func(*OnboardEntityResponse, error)
	}{
		{
			name:     "when request is correct then update entity",
			entityId: "ent_1234",
			request: OnboardEntityRequest{
				Reference:      "reference",
				ContactDetails: &ContactDetails{Phone: &Phone{Number: "2345678910"}},
				Profile: &Profile{
					Urls: []string{"https://www.superheroexample.com"},
					Mccs: []string{"0742"},
				},
				Individual: &Individual{
					FirstName:      "Bruce",
					LastName:       "Wayne",
					TradingName:    "Batman's Super Hero Masks",
					NationalTaxId:  "TAX123456",
					DateOfBirth:    &DateOfBirth{Day: 5, Month: 6, Year: 1995},
					Identification: &Identification{NationalIdNumber: "AB123456C"},
				},
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPut: func(m *mock.Mock) mock.Call {
				return *m.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*OnboardEntityResponse)
						*respMapping = onboardEntity
					})
			},
			checker: func(response *OnboardEntityResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, onboardEntity.Id, response.Id)
				assert.Equal(t, onboardEntity.Reference, response.Reference)
				assert.Equal(t, onboardEntity.Status, response.Status)
				assert.Equal(t, onboardEntity.Capabilities, response.Capabilities)
			},
		},
		{
			name:     "when entity not_found then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			request: OnboardEntityRequest{
				Reference:      "reference",
				ContactDetails: &ContactDetails{Phone: &Phone{Number: "2345678910"}},
				Profile: &Profile{
					Urls: []string{"https://www.superheroexample.com"},
					Mccs: []string{"0742"},
				},
				Individual: &Individual{
					FirstName:      "Bruce",
					LastName:       "Wayne",
					TradingName:    "Batman's Super Hero Masks",
					NationalTaxId:  "TAX123456",
					DateOfBirth:    &DateOfBirth{Day: 5, Month: 6, Year: 1995},
					Identification: &Identification{NationalIdNumber: "AB123456C"},
				},
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPut: func(m *mock.Mock) mock.Call {
				return *m.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *OnboardEntityResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiPut(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.UpdateEntity(tc.entityId, tc.request))
		})
	}
}

func TestCreatePaymentInstruments(t *testing.T) {
	var (
		entityId = "ent_1234"

		metadataResponse = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusAccepted,
		}

		instrumentDocument = InstrumentDocument{
			Type:   "bank_statement",
			FileId: "file_wxglze3wwywujg4nna5fb7ldli",
		}

		address = common.Address{
			AddressLine1: "90 Tottenham Court Road",
			AddressLine2: "",
			City:         "London",
			State:        "London",
			Zip:          "W1T 4TJ",
			Country:      common.GB,
		}

		accountHolder = AccountHolder{
			FirstName:      "Peter",
			LastName:       "Parker",
			BillingAddress: &address,
		}

		paymentInstrument = PaymentInstrument{
			Type:          instruments.Card,
			Label:         "Peter's Personal Account",
			AccountType:   common.Cash,
			AccountNumber: "12345678",
			BankCode:      "050389",
			Currency:      common.GBP,
			Country:       common.GB,
			Document:      &instrumentDocument,
			AccountHolder: &accountHolder,
		}
	)

	cases := []struct {
		name              string
		entityId          string
		paymentInstrument PaymentInstrument
		getAuthorization  func(*mock.Mock) mock.Call
		apiPost           func(*mock.Mock) mock.Call
		checker           func(*common.MetadataResponse, error)
	}{
		{
			name:              "when create a payment instrument then return 202 status",
			entityId:          entityId,
			paymentInstrument: paymentInstrument,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.MetadataResponse)
						*respMapping = metadataResponse
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusAccepted, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:     "when send a bad request then return error",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusBadRequest,
							Status:     "400 Bad Request",
						})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusBadRequest, chkErr.StatusCode)
			},
		},
		{
			name:              "when request is not correct then return error",
			paymentInstrument: PaymentInstrument{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "invalid_request",
								ErrorCodes: []string{"company_or_individual_required"},
							},
						})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "company_or_individual_required")
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.CreatePaymentInstruments(tc.entityId, tc.paymentInstrument))
		})
	}
}

func TestCreatePaymentInstrument(t *testing.T) {
	var (
		entityId = "ent_1234"

		idResponse = common.IdResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "ppi_qn4nis4k3ykpzzu7cvtuvhqqga",
		}

		instrumentDocument = InstrumentDocument{
			Type:   "bank_statement",
			FileId: "file_wxglze3wwywujg4nna5fb7ldli",
		}

		instrumentDetails = InstrumentDetailsFasterPayments{
			AccountNumber: "12345678",
			BankCode:      "050389",
		}

		paymentInstrumentRequest = PaymentInstrumentRequest{
			Label:              "Bob's Bank Account",
			Type:               instruments.BankAccount,
			Currency:           common.GBP,
			Country:            common.GB,
			DefaultDestination: true,
			Document:           &instrumentDocument,
			InstrumentDetails:  &instrumentDetails,
		}
	)

	cases := []struct {
		name                     string
		entityId                 string
		paymentInstrumentRequest PaymentInstrumentRequest
		getAuthorization         func(*mock.Mock) mock.Call
		apiPost                  func(*mock.Mock) mock.Call
		checker                  func(*common.IdResponse, error)
	}{
		{
			name:                     "when create a payment instrument then return 201 status",
			entityId:                 entityId,
			paymentInstrumentRequest: paymentInstrumentRequest,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = idResponse
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, idResponse.Id, response.Id)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:     "when send a bad request then return error",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusBadRequest,
							Status:     "400 Bad Request",
						})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusBadRequest, chkErr.StatusCode)
			},
		},
		{
			name:                     "when request is not correct then return error",
			paymentInstrumentRequest: PaymentInstrumentRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "invalid_request",
								ErrorCodes: []string{"company_or_individual_required"},
							},
						})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "company_or_individual_required")
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.CreatePaymentInstrument(tc.entityId, tc.paymentInstrumentRequest))
		})
	}
}

func TestQueryPaymentInstruments(t *testing.T) {
	var (
		entityId = "ent_1234"

		query = PaymentInstrumentsQuery{
			Status: InstrumentPending,
		}

		instrumentDocument = InstrumentDocument{
			Type:   "bank_statement",
			FileId: "file_wxglze3wwywujg4nna5fb7ldli",
		}

		paymentInstrumentDetailsResponse = PaymentInstrumentDetailsResponse{
			HttpMetadata:       mocks.HttpMetadataStatusOk,
			Id:                 "ppi_qn4nis4k3ykpzzu7cvtuvhqqga",
			Status:             "verified",
			InstrumentId:       "src_pdasnoaxrtoevpyh3opgaxcrti",
			Label:              "Bob's Bank Account",
			Type:               "bank_account",
			Currency:           common.GBP,
			Country:            common.GB,
			DefaultDestination: true,
			Document:           &instrumentDocument,
		}

		paymentInstrumentQueryResponse = PaymentInstrumentQueryResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Data:         []PaymentInstrumentDetailsResponse{paymentInstrumentDetailsResponse},
		}
	)

	cases := []struct {
		name             string
		entityId         string
		query            PaymentInstrumentsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*PaymentInstrumentQueryResponse, error)
	}{
		{
			name:     "when query status of payment instruments then return a list of payment instruments",
			entityId: entityId,
			query:    query,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*PaymentInstrumentQueryResponse)
						*respMapping = paymentInstrumentQueryResponse
					})
			},
			checker: func(response *PaymentInstrumentQueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:     "when send a bad request then return error",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusBadRequest,
							Status:     "400 Bad Request",
						})
			},
			checker: func(response *PaymentInstrumentQueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusBadRequest, chkErr.StatusCode)
			},
		},
		{
			name:  "when request is not correct then return error",
			query: PaymentInstrumentsQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "invalid_request",
								ErrorCodes: []string{"company_or_individual_required"},
							},
						})
			},
			checker: func(response *PaymentInstrumentQueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "company_or_individual_required")
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentInstrumentQueryResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.QueryPaymentInstruments(tc.entityId, tc.query))
		})
	}
}

func TestRetrievePaymentInstrumentDetails(t *testing.T) {
	var (
		entityId            = "ent_1234"
		paymentInstrumentId = "1234"

		instrumentDocument = InstrumentDocument{
			Type:   "bank_statement",
			FileId: "file_wxglze3wwywujg4nna5fb7ldli",
		}

		paymentInstrumentDetailsResponse = PaymentInstrumentDetailsResponse{
			HttpMetadata:       mocks.HttpMetadataStatusOk,
			Id:                 "ppi_qn4nis4k3ykpzzu7cvtuvhqqga",
			Status:             "verified",
			InstrumentId:       "src_pdasnoaxrtoevpyh3opgaxcrti",
			Label:              "Bob's Bank Account",
			Type:               "bank_account",
			Currency:           common.GBP,
			Country:            common.GB,
			DefaultDestination: true,
			Document:           &instrumentDocument,
		}
	)

	cases := []struct {
		name                string
		entityId            string
		paymentInstrumentId string
		getAuthorization    func(*mock.Mock) mock.Call
		apiGet              func(*mock.Mock) mock.Call
		checker             func(*PaymentInstrumentDetailsResponse, error)
	}{
		{
			name:                "when retrieve a payment instrument details then return payment instrument details",
			entityId:            entityId,
			paymentInstrumentId: paymentInstrumentId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*PaymentInstrumentDetailsResponse)
						*respMapping = paymentInstrumentDetailsResponse
					})
			},
			checker: func(response *PaymentInstrumentDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotNil(t, paymentInstrumentDetailsResponse.Id, response.Id)
				assert.NotNil(t, paymentInstrumentDetailsResponse.Label, response.Label)
				assert.NotNil(t, paymentInstrumentDetailsResponse.Type, response.Type)
				assert.NotNil(t, paymentInstrumentDetailsResponse.Currency, response.Currency)
				assert.NotNil(t, paymentInstrumentDetailsResponse.Country, response.Country)
				assert.NotNil(t, paymentInstrumentDetailsResponse.Document, response.Document)
				assert.NotNil(t, paymentInstrumentDetailsResponse.Status, response.Status)
				assert.NotNil(t, paymentInstrumentDetailsResponse.DefaultDestination, response.DefaultDestination)
				assert.NotNil(t, paymentInstrumentDetailsResponse.InstrumentId, response.InstrumentId)
			},
		},
		{
			name:     "when send a bad request then return error",
			entityId: entityId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusBadRequest,
							Status:     "400 Bad Request",
						})
			},
			checker: func(response *PaymentInstrumentDetailsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusBadRequest, chkErr.StatusCode)
			},
		},
		{
			name:                "when request is not correct then return error",
			paymentInstrumentId: "",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Unprocessable",
							Data: &errors.ErrorDetails{
								ErrorType:  "invalid_request",
								ErrorCodes: []string{"company_or_individual_required"},
							},
						})
			},
			checker: func(response *PaymentInstrumentDetailsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "company_or_individual_required")
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *PaymentInstrumentDetailsResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.RetrievePaymentInstrumentDetails(tc.entityId, tc.paymentInstrumentId))
		})
	}
}

func TestGetPayoutSchedule(t *testing.T) {
	var (
		schedule = PayoutSchedule{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Currency: map[common.Currency]CurrencySchedule{
				common.USD: {
					Enabled:    true,
					Threshold:  500,
					Recurrence: NewScheduleFrequencyDailyRequest(),
				},
			},
			Links: map[string]common.Link{
				"self": {
					HRef: &[]string{"https://www.test-link.com"}[0],
				},
			},
		}
	)

	cases := []struct {
		name             string
		entityId         string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*PayoutSchedule, error)
	}{
		{
			name:     "when entity schedule exists then return entity's payout schedule",
			entityId: "ent_1234",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*PayoutSchedule)
						*respMapping = schedule
					})
			},
			checker: func(response *PayoutSchedule, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Currency[common.USD])
				assert.True(t, response.Currency[common.USD].Enabled)
				assert.Equal(t, 500, response.Currency[common.USD].Threshold)
				assert.Equal(t, NewScheduleFrequencyDailyRequest(), response.Currency[common.USD].Recurrence)
			},
		},
		{
			name:     "when entity does not exist then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *PayoutSchedule, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.RetrievePayoutSchedule(tc.entityId))
		})
	}
}

func TestUpdatePayoutSchedule(t *testing.T) {
	var (
		idResponse = common.IdResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Links: map[string]common.Link{
				"self": {
					HRef: &[]string{"https://www.test-link.com"}[0],
				},
			},
		}
	)

	cases := []struct {
		name             string
		entityId         string
		currency         common.Currency
		request          CurrencySchedule
		getAuthorization func(*mock.Mock) mock.Call
		apiPut           func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:     "when request is correct then update entity",
			entityId: "ent_1234",
			currency: common.USD,
			request: CurrencySchedule{
				Enabled:    true,
				Threshold:  500,
				Recurrence: NewScheduleFrequencyDailyRequest(),
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPut: func(m *mock.Mock) mock.Call {
				return *m.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = idResponse
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Links)
				assert.Equal(t, idResponse.Links, response.Links)
			},
		},
		{
			name:     "when entity not_found then return error",
			entityId: "ent_zzzzzzzzzzzzzzzzzzzzzzzzzz",
			currency: common.USD,
			request:  CurrencySchedule{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPut: func(m *mock.Mock) mock.Call {
				return *m.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			filesClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiPut(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient, filesClient)

			tc.checker(client.UpdatePayoutSchedule(tc.entityId, tc.currency, tc.request))
		})
	}
}
