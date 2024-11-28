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

	cardholders "github.com/checkout/checkout-sdk-go/issuing/cardholders"
	cards "github.com/checkout/checkout-sdk-go/issuing/cards"
	controls "github.com/checkout/checkout-sdk-go/issuing/controls"
	issuingTesting "github.com/checkout/checkout-sdk-go/issuing/testing"
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
		AddressLine2: "ABC build",
		City:         "London",
		State:        "London",
		Zip:          "W1T 4TJ",
		Country:      "GB",
	}
	document = &cardholders.CardholderDocument{
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
		response = cardholders.CardholderResponse{
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
		request          cardholders.CardholderRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cardholders.CardholderResponse, error)
	}{
		{
			name: "when request is correct then should return 201",
			request: cardholders.CardholderRequest{
				Type:             cardholders.Individual,
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
						respMapping := args.Get(3).(*cardholders.CardholderResponse)
						*respMapping = response
					})
			},
			checker: func(response *cardholders.CardholderResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.Equal(t, "crh_d3ozhf43pcq2xbldn2g45qnb44", response.Id)
				assert.Equal(t, cardholders.Individual, response.Type)
				assert.Equal(t, cardholders.CardholderActive, response.Status)
				assert.Equal(t, "X-123456-N11", response.Reference)
				assert.NotNil(t, response.CreatedDate)
				assert.NotNil(t, response.LastModifiedDate)
				assert.NotNil(t, response.Links)
			},
		},
		{
			name:    "when request has invalid authorization then return error",
			request: cardholders.CardholderRequest{},
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
			checker: func(response *cardholders.CardholderResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnauthorized, chkErr.StatusCode)

			},
		},
		{
			name:    "when request is not correct then return error",
			request: cardholders.CardholderRequest{},
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
			checker: func(response *cardholders.CardholderResponse, err error) {
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
			request: cardholders.CardholderRequest{},
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
			checker: func(response *cardholders.CardholderResponse, err error) {
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.CreateCardholder(tc.request))
		})
	}
}

func TestGetCardholder(t *testing.T) {
	var (
		response = cardholders.CardholderDetailsResponse{
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
		checker          func(*cardholders.CardholderDetailsResponse, error)
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
						respMapping := args.Get(2).(*cardholders.CardholderDetailsResponse)
						*respMapping = response
					})
			},
			checker: func(response *cardholders.CardholderDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, "crh_d3ozhf43pcq2xbldn2g45qnb44", response.Id)
				assert.Equal(t, cardholders.Individual, response.Type)
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
			checker: func(response *cardholders.CardholderDetailsResponse, err error) {
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardholder(tc.cardholderId))
		})
	}
}

func TestGetCardholderCards(t *testing.T) {
	var (
		cardResponse = cards.CardDetailsResponse{}

		response = cardholders.CardholderCardsResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Cards:        []cards.CardDetailsResponse{cardResponse},
		}
	)

	cases := []struct {
		name             string
		cardholderId     string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*cardholders.CardholderCardsResponse, error)
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
						respMapping := args.Get(2).(*cardholders.CardholderCardsResponse)
						*respMapping = response
					})
			},
			checker: func(response *cardholders.CardholderCardsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
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
			checker: func(response *cardholders.CardholderCardsResponse, err error) {
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardholderCards(tc.cardholderId))
		})
	}
}

func TestCreateCard(t *testing.T) {
	request := cards.NewVirtualCardRequest()
	response := cards.CardResponse{}

	cases := []struct {
		name             string
		request          cards.CardRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.CardResponse, error)
	}{
		{
			name:    "when create a card and this request is correct then should return a response",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*cards.CardResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.CardResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.CreateCard(tc.request))
		})
	}
}

func TestGetCardDetails(t *testing.T) {
	response := cards.CardDetailsResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.CardDetailsResponse, error)
	}{
		{
			name:   "when get a card and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*cards.CardDetailsResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.CardDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardDetails(tc.cardId))
		})
	}
}

func TestEnrollThreeDS(t *testing.T) {
	request := cards.ThreeDSEnrollmentRequest{}
	response := cards.ThreeDSEnrollmentResponse{}

	cases := []struct {
		name             string
		cardId           string
		request          cards.ThreeDSEnrollmentRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.ThreeDSEnrollmentResponse, error)
	}{
		{
			name:    "when enroll a card three DS and this request is correct then should return a response",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*cards.ThreeDSEnrollmentResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.ThreeDSEnrollmentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.EnrollThreeDS(tc.cardId, tc.request))
		})
	}
}

func TestUpdateThreeDS(t *testing.T) {
	request := cards.ThreeDSUpdateRequest{}
	response := cards.ThreeDSUpdateResponse{}

	cases := []struct {
		name             string
		cardId           string
		request          cards.ThreeDSUpdateRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.ThreeDSUpdateResponse, error)
	}{
		{
			name:    "when update a card enroll three DS and this request is correct then should return 201",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Patch", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*cards.ThreeDSUpdateResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.ThreeDSUpdateResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.UpdateThreeDS(tc.cardId, tc.request))
		})
	}
}

func TestGetCardThreeDSDetails(t *testing.T) {
	response := cards.ThreeDSEnrollmentDetailsResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.ThreeDSEnrollmentDetailsResponse, error)
	}{
		{
			name:   "when get a card enroll three DS details and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*cards.ThreeDSEnrollmentDetailsResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.ThreeDSEnrollmentDetailsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardThreeDSDetails(tc.cardId))
		})
	}
}

func TestActivateCard(t *testing.T) {
	response := common.IdResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:   "when activate a card and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.ActivateCard(tc.cardId))
		})
	}
}

func TestGetCardCredentials(t *testing.T) {
	response := cards.CardCredentialsResponse{}

	cases := []struct {
		name             string
		cardId           string
		query            cards.CardCredentialsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.CardCredentialsResponse, error)
	}{
		{
			name:   "when get card credentials and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			query:  cards.CardCredentialsQuery{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*cards.CardCredentialsResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.CardCredentialsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardCredentials(tc.cardId, tc.query))
		})
	}
}

func TestRevokeCard(t *testing.T) {
	request := cards.RevokeCardRequest{}
	response := common.IdResponse{}

	cases := []struct {
		name             string
		cardId           string
		request          cards.RevokeCardRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:    "when revoke a card and this request is correct then should return a response",
			cardId:  "",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.RevokeCard(tc.cardId, tc.request))
		})
	}
}

func TestSuspendCard(t *testing.T) {
	response := common.IdResponse{}

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:   "when suspend a card and this request is correct then should return a response",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.IdResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.SuspendCard(tc.cardId))
		})
	}
}

func TestCreateControl(t *testing.T) {
	request := controls.NewVelocityCardControlRequest()
	response := controls.CardControlResponse{}

	cases := []struct {
		name             string
		request          controls.CardControlRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*controls.CardControlResponse, error)
	}{
		{
			name:    "when create a card control and this request is correct then should return a response",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*controls.CardControlResponse)
						*respMapping = response
					})
			},
			checker: func(response *controls.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.CreateControl(tc.request))
		})
	}
}

func TestGetCardControls(t *testing.T) {
	query := controls.CardControlsQuery{}
	response := controls.CardControlsQueryResponse{}

	cases := []struct {
		name             string
		query            controls.CardControlsQuery
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*controls.CardControlsQueryResponse, error)
	}{
		{
			name:  "when get a card control and this request is correct then should return a response",
			query: query,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*controls.CardControlsQueryResponse)
						*respMapping = response
					})
			},
			checker: func(response *controls.CardControlsQueryResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardControls(tc.query))
		})
	}
}

func TestGetCardControlDetails(t *testing.T) {
	response := controls.CardControlResponse{}

	cases := []struct {
		name             string
		controlId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*controls.CardControlResponse, error)
	}{
		{
			name:      "when get a card control details and this request is correct then should return a response",
			controlId: "ctr_gp7vkmxayztufjz6top5bjcdra",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*controls.CardControlResponse)
						*respMapping = response
					})
			},
			checker: func(response *controls.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiGet(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.GetCardControlDetails(tc.controlId))
		})
	}
}

func TestUpdateCardControl(t *testing.T) {
	request := controls.UpdateCardControlRequest{}
	response := controls.CardControlResponse{}

	cases := []struct {
		name             string
		controlId        string
		request          controls.UpdateCardControlRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiUpdate        func(*mock.Mock) mock.Call
		checker          func(*controls.CardControlResponse, error)
	}{
		{
			name:      "when update a card control and this request is correct then should return a response",
			controlId: "ctr_gp7vkmxayztufjz6top5bjcdra",
			request:   request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiUpdate: func(m *mock.Mock) mock.Call {
				return *m.On("Put", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*controls.CardControlResponse)
						*respMapping = response
					})
			},
			checker: func(response *controls.CardControlResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiUpdate(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.UpdateCardControl(tc.controlId, tc.request))
		})
	}
}

func TestRemoveCardControl(t *testing.T) {
	response := common.IdResponse{}

	cases := []struct {
		name             string
		controlId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiDelete        func(*mock.Mock) mock.Call
		checker          func(*common.IdResponse, error)
	}{
		{
			name:      "when delete a card control and this request is correct then should return a response",
			controlId: "ctr_gp7vkmxayztufjz6top5bjcdra",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("Delete", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*common.IdResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.IdResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiDelete(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.RemoveCardControl(tc.controlId))
		})
	}
}

func TestSimulateAuthorization(t *testing.T) {
	request := issuingTesting.CardAuthorizationRequest{}
	response := issuingTesting.CardAuthorizationResponse{}

	cases := []struct {
		name             string
		request          issuingTesting.CardAuthorizationRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*issuingTesting.CardAuthorizationResponse, error)
	}{
		{
			name:    "when simulate an authorization and this request is correct then should return a response",
			request: request,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*issuingTesting.CardAuthorizationResponse)
						*respMapping = response
					})
			},
			checker: func(response *issuingTesting.CardAuthorizationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := new(mocks.ApiClientMock)
			credentials := new(mocks.CredentialsMock)
			environment := new(mocks.EnvironmentMock)
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.SimulateAuthorization(tc.request))
		})
	}
}

func TestSimulateIncrement(t *testing.T) {
	issuingResponse := issuingTesting.CardSimulationResponse{
		HttpMetadata: mocks.HttpMetadataStatusCreated,
		Status:       issuingTesting.Authorized,
	}

	cases := []struct {
		name             string
		transactionId    string
		request          issuingTesting.CardSimulationRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*issuingTesting.CardSimulationResponse, error)
	}{
		{
			name:          "when simulating an increment authorization with valid request then return response",
			transactionId: "transaction_id",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*issuingTesting.CardSimulationResponse)
						*respMapping = issuingResponse
					})
			},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, issuingResponse.HttpMetadata.StatusCode, response.HttpMetadata.StatusCode)
				assert.Equal(t, issuingResponse.Status, response.Status)
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
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:          "when simulating an increment authorization with invalid transactionId then return error",
			transactionId: "not_found",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.SimulateIncrement(tc.transactionId, tc.request))
		})
	}
}

func TestSimulateClearing(t *testing.T) {
	issuingResponse := common.MetadataResponse{
		HttpMetadata: mocks.HttpMetadataStatusAccepted,
	}

	cases := []struct {
		name             string
		transactionId    string
		request          issuingTesting.CardSimulationRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:          "when simulating a clearing authorization with valid request then return response",
			transactionId: "transaction_id",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*common.MetadataResponse)
						*respMapping = issuingResponse
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, issuingResponse.HttpMetadata.StatusCode, response.HttpMetadata.StatusCode)
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
		{
			name:          "when simulating a clearing authorization with invalid transactionId then return error",
			transactionId: "not_found",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.SimulateClearing(tc.transactionId, tc.request))
		})
	}
}

func TestSimulateReversal(t *testing.T) {
	issuingResponse := issuingTesting.CardSimulationResponse{
		HttpMetadata: mocks.HttpMetadataStatusCreated,
		Status:       issuingTesting.Reversed,
	}

	cases := []struct {
		name             string
		transactionId    string
		request          issuingTesting.CardSimulationRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*issuingTesting.CardSimulationResponse, error)
	}{
		{
			name:          "when simulating a reversal authorization with valid request then return response",
			transactionId: "transaction_id",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*issuingTesting.CardSimulationResponse)
						*respMapping = issuingResponse
					})
			},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, issuingResponse.HttpMetadata.StatusCode, response.HttpMetadata.StatusCode)
				assert.Equal(t, issuingResponse.Status, response.Status)
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
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:          "when simulating an reversal authorization with invalid transactionId then return error",
			transactionId: "not_found",
			request:       issuingTesting.CardSimulationRequest{Amount: 100},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusNotFound,
							Status:     "404 Not Found",
						})
			},
			checker: func(response *issuingTesting.CardSimulationResponse, err error) {
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			configuration := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(configuration, apiClient)

			tc.checker(client.SimulateReversal(tc.transactionId, tc.request))
		})
	}
}
