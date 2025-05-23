package transfers

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

func TestInitiateTransferOfFounds(t *testing.T) {
	var (
		transfer = TransferResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "tra_y3oqhf46pyzuxjbcn2giaqnb4",
			Status:       "pending",
		}
	)

	cases := []struct {
		name             string
		request          TransferRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*TransferResponse, error)
	}{
		{
			name: "when request is correct then create transfer of founds",
			request: TransferRequest{
				Reference:    "reference",
				TransferType: Commission,
				Source:       &TransferSourceRequest{},
				Destination:  &TransferDestinationRequest{},
			},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*TransferResponse)
						*respMapping = transfer
					})
			},
			checker: func(response *TransferResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
				assert.Equal(t, transfer.Id, response.Id)
				assert.Equal(t, transfer.Status, response.Status)
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
			checker: func(response *TransferResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request invalid then return error",
			request: TransferRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(
						errors.CheckoutAPIError{
							StatusCode: http.StatusUnprocessableEntity,
							Status:     "422 Invalid Request",
							Data: &errors.ErrorDetails{
								ErrorType: "request_invalid",
							},
						})
			},
			checker: func(response *TransferResponse, err error) {
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
			enableTelemertry := true

			tc.getAuthorization(&credentials.Mock)
			tc.apiPost(&apiClient.Mock)

			conf := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(conf, apiClient)

			tc.checker(client.InitiateTransferOfFounds(tc.request, nil))
		})
	}
}

func TestRetrieveTransfer(t *testing.T) {
	var (
		now = time.Now()

		transferDetails = TransferDetails{
			HttpMetadata: mocks.HttpMetadataStatusOk,
			Id:           "tra_y3oqhf46pyzuxjbcn2giaqnb4",
			Reference:    "superhero1234",
			Status:       "rejected",
			TransferType: Commission,
			RequestedOn:  &now,
			ReasonCodes: []string{"destination_transfers_capability_disabled",
				"source_and_destination_currency_accounts_must_be_different",
			},
			Source: &TransferSourceResponse{
				EntityId: "ent_azsiyswl7bwe2ynjzujy7lcjca",
				Amount:   int64(100),
				Currency: common.GBP,
			},
			Destination: &TransferDestinationResponse{EntityId: "ent_bqik7gxoavwhmy3ot6kvmbx6py"},
		}
	)

	cases := []struct {
		name             string
		transferId       string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*TransferDetails, error)
	}{
		{
			name:       "when transfer exists then return transfer details",
			transferId: "tra_y3oqhf46pyzuxjbcn2giaqnb4",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*TransferDetails)
						*respMapping = transferDetails
					})
			},
			checker: func(response *TransferDetails, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, transferDetails.Id, response.Id)
				assert.Equal(t, transferDetails.Reference, response.Reference)
				assert.Equal(t, transferDetails.Status, response.Status)
				assert.Equal(t, transferDetails.TransferType, response.TransferType)
				assert.Equal(t, transferDetails.Source, response.Source)
				assert.Equal(t, int64(100), response.Source.Amount)
				assert.Equal(t, common.GBP, response.Source.Currency)
				assert.Equal(t, transferDetails.Destination, response.Destination)
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
			checker: func(response *TransferDetails, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:       "when transfer not found then return error",
			transferId: "not_found",
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
			checker: func(response *TransferDetails, err error) {
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

			conf := configuration.NewConfiguration(credentials, &enableTelemertry, environment, &http.Client{}, nil)
			client := NewClient(conf, apiClient)

			tc.checker(client.RetrieveTransfer(tc.transferId))
		})
	}
}
