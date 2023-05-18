package issuing

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
)

var (
	createdDate      = time.Now()
	lastModifiedDate = time.Now().Add(-5 * time.Hour)
	phone            = &common.Phone{
		CountryCode: "+1",
		Number:      "415 555 2671",
	}
	address = &common.Address{
		AddressLine1: "Checkout.com",
		AddressLine2: "90 Tottenham Court Road",
		City:         "London",
		State:        "London",
		Zip:          "W1T 4TJ",
		Country:      "GB",
	}
	document = &CardholderDocument{
		Type:            "national_identity_card",
		FrontDocumentId: "file_6lbss42ezvoufcb2beo76rvwly",
		BackDocumentId:  "file_aaz5pemp6326zbuvevp6qroqu4",
	}
	links = map[string]common.Link{
		"self": {
			HRef: &[]string{"https://api.checkout.com/issuing/cardholders/crh_d3ozhf43pcq2xbldn2g45qnb44"}[0],
		},
		"cards": {
			HRef: &[]string{"https://api.checkout.com/issuing/cards"}[0],
		},
	}
)

func TestCreateCardholder(t *testing.T) {
	var (
		response = CardholderResponse{
			HttpMetadata:     mocks.HttpMetadataStatusCreated,
			Id:               "crh_d3ozhf43pcq2xbldn2g45qnb44",
			Type:             "individual",
			Status:           "active",
			Reference:        "X-123456-N11",
			CreatedDate:      &createdDate,
			LastModifiedDate: &lastModifiedDate,
			Links:            links,
		}
	)

	cases := []struct {
		name             string
		request          CardholderRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*CardholderResponse, error)
	}{
		{
			name: "when request is correct then should return 201",
			request: CardholderRequest{
				Type:             Individual,
				Reference:        "X-123456-N11",
				EntityId:         "ent_fa6psq242dcd6fdn5gifcq1491",
				FirstName:        "John",
				MiddleName:       "Fitzgerald",
				LastName:         "Kennedy",
				Email:            "john.kennedy@myemaildomain.com",
				PhoneNumber:      phone,
				DateOfBirth:      "1985-05-15",
				BillingAddress:   address,
				ResidencyAddress: address,
				Document:         document,
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*CardholderResponse)
						*respMapping = response
					})
			},
			checker: func(response *CardholderResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.Equal(t, "crh_d3ozhf43pcq2xbldn2g45qnb44", response.Id)
				assert.Equal(t, Individual, response.Type)
				assert.Equal(t, CardholderActive, response.Status)
				assert.Equal(t, "X-123456-N11", response.Reference)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name:    "when request has invalid authorization then return error",
			request: CardholderRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnauthorized,
							Status:     "401 Unauthorized",
						})
			},
			checker: func(response *CardholderResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

			},
		},
		{
			name:    "when request is not correct then return error",
			request: CardholderRequest{},
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
								ErrorType:  "request_invalid",
								ErrorCodes: []string{"request_body_malformed"},
							},
						})
			},
			checker: func(response *CardholderResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "request_body_malformed")
			},
		},
		{
			name:    "when server has problems then return error",
			request: CardholderRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusInternalServerError,
							Status:     "500 Internal Server Error",
						})
			},
			checker: func(response *CardholderResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusInternalServerError, chkErr.StatusCode)

			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.CreateCardholder(tc.request))
		})
	}
}

func TestGetCardholder(t *testing.T) {
	var (
		response = CardholderDetailsResponse{
			HttpMetadata:      mocks.HttpMetadataStatusOk,
			Id:                "crh_d3ozhf43pcq2xbldn2g45qnb44",
			Type:              "individual",
			FirstName:         "John",
			MiddleName:        "Fitzgerald",
			LastName:          "Kennedy",
			Email:             "john.kennedy@myemaildomain.com",
			PhoneNumber:       phone,
			DateOfBirth:       "1985-05-15",
			BillingAddress:    address,
			ResidencyAddress:  address,
			Reference:         "X-123456-N11",
			AccountEntityId:   "ent_fa6psq242dcd6fdn5gifcq1491",
			ParentSubEntityId: "ent_fa6psq242dcd6fdn5gifcq1491",
			EntityId:          "ent_fa6psq242dcd6fdn5gifcq1491",
			CreatedDate:       &createdDate,
			LastModifiedDate:  &lastModifiedDate,
			Links:             links,
		}
	)
	cases := []struct {
		name             string
		cardholderId     string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*CardholderDetailsResponse, error)
	}{
		{
			name:         "when request is correct then should return a cardholder",
			cardholderId: "crh_d3ozhf43pcq2xbldn2g45qnb44",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*CardholderDetailsResponse)
						*respMapping = response
					})
			},
			checker: func(response *CardholderDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, "crh_d3ozhf43pcq2xbldn2g45qnb44", response.Id)
				assert.Equal(t, Individual, response.Type)
				assert.Equal(t, "John", response.FirstName)
				assert.Equal(t, "Fitzgerald", response.MiddleName)
				assert.Equal(t, "Kennedy", response.LastName)
				assert.Equal(t, "john.kennedy@myemaildomain.com", response.Email)
				assert.NotNil(t, response.PhoneNumber)
				assert.Equal(t, "1985-05-15", response.DateOfBirth)
				assert.NotNil(t, response.BillingAddress)
				assert.NotNil(t, response.ResidencyAddress)
				assert.Equal(t, "X-123456-N11", response.Reference)
				assert.Equal(t, "ent_fa6psq242dcd6fdn5gifcq1491", response.AccountEntityId)
				assert.Equal(t, "ent_fa6psq242dcd6fdn5gifcq1491", response.ParentSubEntityId)
				assert.Equal(t, "ent_fa6psq242dcd6fdn5gifcq1491", response.EntityId)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name:         "when request is not correct then return error",
			cardholderId: "crh_not_found",
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
			checker: func(response *CardholderDetailsResponse, err error) {
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardholder(tc.cardholderId))
		})
	}
}

func TestGetCardholderCards(t *testing.T) {
	CardTypeResponse := NewVirtualCardTypeResponse()
	CardTypeResponse.CardDetailsCardholder = CardDetailsCardholder{
		Type:             Virtual,
		Id:               "crd_fa6psq242dcd6fdn5gifcq1491",
		CardholderId:     "crh_d3ozhf43pcq2xbldn2g45qnb44",
		CardProductId:    "pro_7syjig3jq3mezlc3vjrdpfitl4",
		ClientId:         "cli_vkuhvk4vjn2edkps7dfsq6emqm",
		LastFour:         "1234",
		ExpiryMonth:      5,
		ExpiryYear:       2025,
		Status:           CardActive,
		DisplayName:      "JOHN KENNEDY",
		BillingCurrency:  common.USD,
		IssuingCountry:   common.US,
		Reference:        "X-123456-N11",
		CreatedDate:      &createdDate,
		LastModifiedDate: &lastModifiedDate,
	}
	CardTypeResponse.IsSingleUse = true

	var (
		cardDetailsResponse = CardDetailsResponse{
			CardTypeResponse: CardTypeResponse,
		}

		response = CardholderCardsResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Cards:        []CardDetailsResponse{cardDetailsResponse},
		}
	)

	cases := []struct {
		name             string
		cardholderId     string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*CardholderCardsResponse, error)
	}{
		{
			name:         "when request is correct then should return a cardholder",
			cardholderId: "crh_d3ozhf43pcq2xbldn2g45qnb44",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*CardholderCardsResponse)
						*respMapping = response
					})
			},
			checker: func(response *CardholderCardsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, Virtual, response.Cards[0].GetDetails().(virtualCardTypeResponse).Type)
				assert.Equal(t, "crd_fa6psq242dcd6fdn5gifcq1491", response.Cards[0].GetDetails().(virtualCardTypeResponse).Id)
				assert.Equal(t, "crh_d3ozhf43pcq2xbldn2g45qnb44", response.Cards[0].GetDetails().(virtualCardTypeResponse).CardholderId)
				assert.Equal(t, "pro_7syjig3jq3mezlc3vjrdpfitl4", response.Cards[0].GetDetails().(virtualCardTypeResponse).CardProductId)
				assert.Equal(t, "cli_vkuhvk4vjn2edkps7dfsq6emqm", response.Cards[0].GetDetails().(virtualCardTypeResponse).ClientId)
				assert.Equal(t, "1234", response.Cards[0].GetDetails().(virtualCardTypeResponse).LastFour)
				assert.Equal(t, 5, response.Cards[0].GetDetails().(virtualCardTypeResponse).ExpiryMonth)
				assert.Equal(t, 2025, response.Cards[0].GetDetails().(virtualCardTypeResponse).ExpiryYear)
				assert.Equal(t, CardActive, response.Cards[0].GetDetails().(virtualCardTypeResponse).Status)
				assert.Equal(t, "JOHN KENNEDY", response.Cards[0].GetDetails().(virtualCardTypeResponse).DisplayName)
				assert.Equal(t, common.USD, response.Cards[0].GetDetails().(virtualCardTypeResponse).BillingCurrency)
				assert.Equal(t, common.US, response.Cards[0].GetDetails().(virtualCardTypeResponse).IssuingCountry)
				assert.Equal(t, "X-123456-N11", response.Cards[0].GetDetails().(virtualCardTypeResponse).Reference)
				assert.NotNil(t, response.Cards[0].GetDetails().(virtualCardTypeResponse).CreatedDate)
				assert.NotNil(t, response.Cards[0].GetDetails().(virtualCardTypeResponse).LastModifiedDate)
				assert.Equal(t, true, response.Cards[0].GetDetails().(virtualCardTypeResponse).IsSingleUse)
			},
		},
		{
			name:         "when request is not correct then return error",
			cardholderId: "crh_not_found",
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
			checker: func(response *CardholderCardsResponse, err error) {
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

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardholderCards(tc.cardholderId))
		})
	}
}
