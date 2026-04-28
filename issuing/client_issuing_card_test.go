package issuing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"

	cards "github.com/checkout/checkout-sdk-go/v2/issuing/cards"
)

// # tests

func TestRenewCard(t *testing.T) {
	var (
		response = cards.RenewCardResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
			Id:           "crd_renewed_fa6psq242dcd6fdn5gifcq1491",
			ParentCardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			Status:       cards.CardActive,
			Type:         cards.Virtual,
		}
	)

	cases := []struct {
		name             string
		cardId           string
		request          cards.RenewCardRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.RenewCardResponse, error)
	}{
		{
			name:    "when request is correct then should return 201",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: buildVirtualRenewCardRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*cards.RenewCardResponse)
						*respMapping = response
					})
			},
			checker: assertRenewCardSuccess(t, &response),
		},
		{
			name:   "when credentials are invalid then return authorization error",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			request: buildVirtualRenewCardRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *cards.RenewCardResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when card not found then return 404",
			cardId:  "crd_not_found",
			request: buildVirtualRenewCardRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *cards.RenewCardResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:    "when request is invalid then return 422",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: buildVirtualRenewCardRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable",
						Data: &errors.ErrorDetails{
							ErrorType:  "request_invalid",
							ErrorCodes: []string{"request_body_malformed"},
						},
					})
			},
			checker: func(response *cards.RenewCardResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "request_body_malformed")
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

			tc.checker(client.RenewCard(tc.cardId, tc.request))
		})
	}
}

func TestScheduleCardRevocation(t *testing.T) {
	var (
		response = cards.ScheduleRevocationResponse{
			HttpMetadata: mocks.HttpMetadataStatusCreated,
		}
	)

	cases := []struct {
		name             string
		cardId           string
		request          cards.ScheduleRevocationRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*cards.ScheduleRevocationResponse, error)
	}{
		{
			name:    "when request is correct then should return 201",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: buildScheduleRevocationRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*cards.ScheduleRevocationResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.ScheduleRevocationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:    "when card not found then return 404",
			cardId:  "crd_not_found",
			request: buildScheduleRevocationRequest(),
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *cards.ScheduleRevocationResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:    "when request is invalid then return 422",
			cardId:  "crd_fa6psq242dcd6fdn5gifcq1491",
			request: cards.ScheduleRevocationRequest{},
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusUnprocessableEntity,
						Status:     "422 Unprocessable",
						Data: &errors.ErrorDetails{
							ErrorType:  "request_invalid",
							ErrorCodes: []string{"revocation_date_required"},
						},
					})
			},
			checker: func(response *cards.ScheduleRevocationResponse, err error) {
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

			tc.checker(client.ScheduleCardRevocation(tc.cardId, tc.request))
		})
	}
}

func TestDeleteScheduledRevocation(t *testing.T) {
	var (
		response = cards.ScheduleRevocationResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
		}
	)

	cases := []struct {
		name             string
		cardId           string
		getAuthorization func(*mock.Mock) mock.Call
		apiDelete        func(*mock.Mock) mock.Call
		checker          func(*cards.ScheduleRevocationResponse, error)
	}{
		{
			name:   "when request is correct then should return 200",
			cardId: "crd_fa6psq242dcd6fdn5gifcq1491",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*cards.ScheduleRevocationResponse)
						*respMapping = response
					})
			},
			checker: func(response *cards.ScheduleRevocationResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:   "when card not found then return 404",
			cardId: "crd_not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiDelete: func(m *mock.Mock) mock.Call {
				return *m.On("DeleteWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *cards.ScheduleRevocationResponse, err error) {
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

			tc.checker(client.DeleteScheduledRevocation(tc.cardId))
		})
	}
}

// # common methods

func buildVirtualRenewCardRequest() cards.RenewCardRequest {
	req := cards.NewVirtualCardRenewRequest()
	req.DisplayName = "John Kennedy"
	req.Reference = "X-123456-N11"
	return req
}

func buildScheduleRevocationRequest() cards.ScheduleRevocationRequest {
	return cards.ScheduleRevocationRequest{
		RevocationDate: "2026-12-31",
	}
}

func assertRenewCardSuccess(t *testing.T, expected *cards.RenewCardResponse) func(*cards.RenewCardResponse, error) {
	return func(response *cards.RenewCardResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, http.StatusCreated, response.HttpMetadata.StatusCode)
		assert.Equal(t, expected.Id, response.Id)
		assert.Equal(t, expected.ParentCardId, response.ParentCardId)
		assert.Equal(t, expected.Status, response.Status)
		assert.Equal(t, expected.Type, response.Type)
	}
}
