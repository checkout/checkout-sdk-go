package nas

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
	"github.com/checkout/checkout-sdk-go/payments"
)

const (
	instrumentId = "src_wmlfc3zyhqzehihu7giusaaawu"
)

var (
	accountHolder = common.AccountHolder{
		FirstName: "Bruce",
		LastName:  "Wayne",
	}

	customerRequest = CreateCustomerInstrumentRequest{
		Id:    "cus_y3oqhf46pyzuxjbcn2giaqnb44",
		Email: "bruce@wayne-enterprises.com",
		Name:  "Bruce Wayne",
	}

	customerResponse = common.CustomerResponse{
		Id:    "cus_y3oqhf46pyzuxjbcn2giaqnb44",
		Email: "bruce@wayne-enterprises.com",
		Name:  "Bruce Wayne",
	}

	bank = common.BankDetails{
		Name:   "Lloyds TSB",
		Branch: "Bournemouth",
	}
)

func TestCreate(t *testing.T) {
	var (
		token = CreateTokenInstrumentResponse{
			Type:             common.Card,
			Id:               "src_wmlfc3zyhqzehihu7giusaaawu",
			CustomerResponse: &customerResponse,
			ExpiryMonth:      6,
			ExpiryYear:       2025,
			Last4:            "1234",
		}

		createTokenResponse = CreateInstrumentResponse{
			HttpMetadata:                  mocks.HttpMetadataStatusCreated,
			CreateTokenInstrumentResponse: &token,
		}

		bankAccount = CreateBankAccountInstrumentResponse{
			Type:             common.BankAccount,
			Id:               "src_wmlfc3zyhqzehihu7giusaaawu",
			CustomerResponse: &customerResponse,
			BankDetails:      &bank,
			SwiftBic:         "37040044",
			AccountNumber:    "12345",
			Iban:             "HU93116000060000000012345676",
		}

		createBankAccountResponse = CreateInstrumentResponse{
			HttpMetadata:                        mocks.HttpMetadataStatusCreated,
			CreateBankAccountInstrumentResponse: &bankAccount,
		}

		sepa = CreateSepaInstrumentResponse{
			Type:        common.Sepa,
			Id:          "src_wmlfc3zyhqzehihu7giusaaawu",
			Fingerprint: "vnsdrvikkvre3dtrjjvlm5du4q",
		}

		createSepaResponse = CreateInstrumentResponse{
			HttpMetadata:                 mocks.HttpMetadataStatusCreated,
			CreateSepaInstrumentResponse: &sepa,
		}
	)

	cases := []struct {
		name             string
		request          CreateInstrumentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CreateInstrumentResponse, error)
	}{
		{
			name:    "when request is for token instrument then create token instrument",
			request: getCreateTokenInstrumentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CreateInstrumentResponse)
						*respMapping = createTokenResponse
					})
			},
			checker: func(response *CreateInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.CreateTokenInstrumentResponse)
				assert.Equal(t, token.Id, response.CreateTokenInstrumentResponse.Id)
			},
		},
		{
			name:    "when request is for bank account instrument then create bank account instrument",
			request: getCreateBankAccountInstrumentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CreateInstrumentResponse)
						*respMapping = createBankAccountResponse
					})
			},
			checker: func(response *CreateInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.CreateBankAccountInstrumentResponse)
				assert.Equal(t, bankAccount.Id, response.CreateBankAccountInstrumentResponse.Id)
			},
		},
		{
			name:    "when request is for sepa instrument then create sepa instrument",
			request: getCreateSepaInstrumentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*CreateInstrumentResponse)
						*respMapping = createSepaResponse
					})
			},
			checker: func(response *CreateInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.CreateSepaInstrumentResponse)
				assert.Equal(t, sepa.Id, response.CreateSepaInstrumentResponse.Id)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *CreateInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: NewCreateTokenInstrumentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Invalid Request",
							Data: &errors.ErrorDetails{
								ErrorType: "request_invalid",
								ErrorCodes: []string{
									"email_required",
								},
							},
						})
			},
			checker: func(response *CreateInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
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
			tc.apiPost(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.Create(tc.request))
		})
	}
}

func getCreateTokenInstrumentRequest() *createTokenInstrumentRequest {
	r := NewCreateTokenInstrumentRequest()
	r.Token = "tok_asoto22g2fsu7prwomy12sgfsa"
	r.AccountHolder = &accountHolder
	r.Customer = &customerRequest
	return r
}

func getCreateSepaInstrumentRequest() *createSepaInstrumentRequest {
	// Create APIShortDate with current date
	apiDate := (*common.APIShortDate)(&time.Time{})
	*apiDate = common.APIShortDate(time.Now())

	r := NewCreateSepaInstrumentRequest()
	r.InstrumentData = &InstrumentData{
		AccountNumber:   "FR2810096000509685512959O86",
		Country:         common.GB,
		Currency:        common.GBP,
		PaymentType:     payments.Recurring,
		MandateId:       "1234567890",
		DateOfSignature: apiDate,
	}
	return r
}

func getCreateBankAccountInstrumentRequest() *createBankAccountInstrumentRequest {
	r := NewCreateBankAccountInstrumentRequest()
	r.AccountType = common.Savings
	r.AccountNumber = "12345"
	r.Iban = "HU93116000060000000012345676"
	r.SwiftBic = "37040044"
	r.Currency = common.GBP
	r.Country = common.GB
	r.AccountHolder = &accountHolder
	r.BankDetails = &bank
	return r
}

func TestGet(t *testing.T) {
	var (
		cardInstrument = GetCardInstrumentResponse{
			Type:        common.Card,
			Id:          "src_wmlfc3zyhqzehihu7giusaaawu",
			ExpiryMonth: 6,
			ExpiryYear:  2025,
			Last4:       "1234",
		}

		response = GetInstrumentResponse{
			HttpMetadata:              mocks.HttpMetadataStatusOk,
			GetCardInstrumentResponse: &cardInstrument,
		}
	)

	cases := []struct {
		name             string
		instrumentId     string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*GetInstrumentResponse, error)
	}{
		{
			name:         "when instrument exists then return instrument info",
			instrumentId: instrumentId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*GetInstrumentResponse)
						*respMapping = response
					})
			},
			checker: func(response *GetInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.GetCardInstrumentResponse)
				assert.Equal(t, cardInstrument.Id, response.GetCardInstrumentResponse.Id)
				assert.Equal(t, cardInstrument.Type, response.GetCardInstrumentResponse.Type)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *GetInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:         "when instrument not found then return error",
			instrumentId: "not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *GetInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
				assert.Equal(t, "404 Not Found", chkErr.Status)
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

			tc.checker(client.Get(tc.instrumentId))
		})
	}
}

func TestClientGetBankAccountFieldFormatting(t *testing.T) {
	var (
		allowOption = InstrumentSectionFieldAllowedOption{
			Id:      "1234",
			Display: "1234",
		}

		dependencies = InstrumentSectionFieldDependencies{
			FieldId: "1234",
			Value:   "1234",
		}

		sectionField = InstrumentSectionField{
			Id:              "1234",
			Section:         "1234",
			Display:         "1234",
			HelpText:        "1234",
			Type:            "1234",
			Required:        true,
			ValidationRegex: "1234",
			MinLength:       0,
			MaxLength:       1000,
			AllowedOptions:  []InstrumentSectionFieldAllowedOption{allowOption},
			Dependencies:    []InstrumentSectionFieldDependencies{dependencies},
		}

		section = InstrumentSection{
			Name:   "name",
			Fields: []InstrumentSectionField{sectionField},
		}

		query = QueryBankAccountFormatting{
			AccountHolderType: common.Individual,
			PaymentNetwork:    Ach,
		}

		response = GetBankAccountFieldFormattingResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Sections:     []InstrumentSection{section},
		}
	)

	cases := []struct {
		name             string
		country          string
		currency         string
		query            QueryBankAccountFormatting
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*GetBankAccountFieldFormattingResponse, error)
	}{
		{
			name:     "When fetching valid bank account validations then it returns validations",
			country:  string(common.GB),
			currency: string(common.GBP),
			query:    query,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*GetBankAccountFieldFormattingResponse)
						*respMapping = response
					})
			},
			checker: func(response *GetBankAccountFieldFormattingResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Sections)
				assert.NotNil(t, response.Sections[0].Name)
				assert.NotNil(t, response.Sections[0].Fields)
				assert.NotNil(t, response.Sections[0].Fields[0].Id)
				assert.NotNil(t, response.Sections[0].Fields[0].Section)
				assert.NotNil(t, response.Sections[0].Fields[0].Display)
				assert.NotNil(t, response.Sections[0].Fields[0].HelpText)
				assert.NotNil(t, response.Sections[0].Fields[0].Type)
				assert.NotNil(t, response.Sections[0].Fields[0].Required)
				assert.NotNil(t, response.Sections[0].Fields[0].ValidationRegex)
				assert.NotNil(t, response.Sections[0].Fields[0].MinLength)
				assert.NotNil(t, response.Sections[0].Fields[0].MaxLength)
			},
		},
		{
			name:     "when credentials invalid then return error",
			country:  string(common.GB),
			currency: string(common.GBP),
			query:    query,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *GetBankAccountFieldFormattingResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:     "when bank account validation not found then return error",
			country:  string(common.GB),
			currency: string(common.GBP),
			query:    query,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *GetBankAccountFieldFormattingResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
				assert.Equal(t, "404 Not Found", chkErr.Status)
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

			tc.checker(client.GetBankAccountFieldFormatting(tc.country, tc.currency, tc.query))
		})
	}
}

func TestUpdate(t *testing.T) {
	var (
		updateCardResponse = UpdateCardInstrumentResponse{
			Type:        common.Card,
			Fingerprint: "smoua2sbuqhupeofwbe77n5nsm",
		}

		response = UpdateInstrumentResponse{
			HttpMetadata:                 mocks.HttpMetadataStatusNoContent,
			UpdateCardInstrumentResponse: &updateCardResponse,
		}
	)

	cases := []struct {
		name             string
		instrumentId     string
		request          UpdateInstrumentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPatch         func(*mock.Mock) mock.Call
		checker          func(*UpdateInstrumentResponse, error)
	}{
		{
			name:         "when request is correct then update instrument",
			instrumentId: instrumentId,
			request:      NewUpdateCardInstrumentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*UpdateInstrumentResponse)
						*respMapping = response
					})
			},
			checker: func(response *UpdateInstrumentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.UpdateCardInstrumentResponse)
			},
		},
		{
			name:         "when credentials invalid then return error",
			instrumentId: instrumentId,
			request:      NewUpdateCardInstrumentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *UpdateInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:         "when instrument not found then return error",
			instrumentId: "not_found",
			request:      NewUpdateCardInstrumentRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *UpdateInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:         "when request invalid then return error",
			instrumentId: instrumentId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPatch: func(m *mock.Mock) mock.Call {
				return *m.On("PatchWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Invalid Request",
							Data: &errors.ErrorDetails{
								ErrorType: "request_invalid",
								ErrorCodes: []string{
									"email_invalid",
								},
							},
						})
			},
			checker: func(response *UpdateInstrumentResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
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
			tc.apiPatch(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.Update(tc.instrumentId, tc.request))
		})
	}
}

func TestDelete(t *testing.T) {
	var (
		response = common.MetadataResponse{HttpMetadata: mocks.HttpMetadataStatusNoContent}
	)

	cases := []struct {
		name             string
		instrumentId     string
		getAuthorization func(*mock.Mock) mock.Call
		apiDelete        func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:         "when request is correct then delete instrument",
			instrumentId: instrumentId,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.MetadataResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusNoContent, response.HttpMetadata.StatusCode)
			},
		},
		{
			name: "when credentials invalid then return error",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:         "when instrument not found then return error",
			instrumentId: "not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *common.MetadataResponse, err error) {
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
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemetry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiDelete(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.Delete(tc.instrumentId))
		})
	}
}
