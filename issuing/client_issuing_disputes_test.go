package issuing

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/checkout/checkout-sdk-go/v2/common"
	"github.com/checkout/checkout-sdk-go/v2/configuration"
	"github.com/checkout/checkout-sdk-go/v2/errors"
	"github.com/checkout/checkout-sdk-go/v2/mocks"

	disputes "github.com/checkout/checkout-sdk-go/v2/issuing/disputes"
)

// # tests

func TestCreateDispute(t *testing.T) {
	idempotencyKey := "test-idempotency-key-12345"

	var (
		response = disputes.IssuingDisputeResponse{
			HttpMetadata:  mocks.HttpMetadataStatusCreated,
			Id:            "idsp_test_12345abcdefghijklmnop",
			Reason:        "4837",
			Status:        disputes.DisputeCreated,
			TransactionId: "trx_test_abcdefghijklmnopqr",
		}
	)

	cases := []struct {
		name             string
		request          disputes.CreateDisputeRequest
		idempotencyKey   *string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*disputes.IssuingDisputeResponse, error)
	}{
		{
			name:           "when request is correct then should return 201",
			request:        buildCreateDisputeRequest(),
			idempotencyKey: &idempotencyKey,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*disputes.IssuingDisputeResponse)
						*respMapping = response
					})
			},
			checker: assertIssuingDisputeSuccess(t, &response),
		},
		{
			name:           "when credentials are invalid then return authorization error",
			request:        disputes.CreateDisputeRequest{},
			idempotencyKey: nil,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(nil, errors.CheckoutAuthorizationError("Invalid authorization type"))
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil)
			},
			checker: func(response *disputes.IssuingDisputeResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:           "when request is invalid then return 422",
			request:        disputes.CreateDisputeRequest{},
			idempotencyKey: nil,
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
							ErrorCodes: []string{"transaction_id_required"},
						},
					})
			},
			checker: func(response *disputes.IssuingDisputeResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "request_invalid", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "transaction_id_required")
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

			tc.checker(client.CreateDispute(tc.request, tc.idempotencyKey))
		})
	}
}

func TestGetDispute(t *testing.T) {
	var (
		response = disputes.IssuingDisputeResponse{
			HttpMetadata:  mocks.HttpMetadataStatusOk,
			Id:            "idsp_test_12345abcdefghijklmnop",
			Reason:        "4837",
			Status:        disputes.DisputeCreated,
			TransactionId: "trx_test_abcdefghijklmnopqr",
		}
	)

	cases := []struct {
		name             string
		disputeId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*disputes.IssuingDisputeResponse, error)
	}{
		{
			name:      "when request is correct then should return 200",
			disputeId: "idsp_test_12345abcdefghijklmnop",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*disputes.IssuingDisputeResponse)
						*respMapping = response
					})
			},
			checker: assertIssuingDisputeSuccess(t, &response),
		},
		{
			name:      "when dispute not found then return 404",
			disputeId: "idsp_not_found",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("GetWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(errors.CheckoutAPIError{
						StatusCode: http.StatusNotFound,
						Status:     "404 Not Found",
					})
			},
			checker: func(response *disputes.IssuingDisputeResponse, err error) {
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
			tc.apiGet(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.GetDispute(tc.disputeId))
		})
	}
}

func TestCancelDispute(t *testing.T) {
	idempotencyKey := "test-idempotency-key-cancel"

	var (
		response = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
		}
	)

	cases := []struct {
		name             string
		disputeId        string
		idempotencyKey   *string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:           "when request is correct then should return 200",
			disputeId:      "idsp_test_12345abcdefghijklmnop",
			idempotencyKey: &idempotencyKey,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*common.MetadataResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:           "when dispute not found then return 404",
			disputeId:      "idsp_not_found",
			idempotencyKey: nil,
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
			tc.apiPost(&apiClient.Mock)

			config := configuration.NewConfiguration(credentials, &enableTelemetry, environment, &http.Client{}, nil)
			client := NewClient(config, apiClient)

			tc.checker(client.CancelDispute(tc.disputeId, tc.idempotencyKey))
		})
	}
}

func TestEscalateDispute(t *testing.T) {
	idempotencyKey := "test-idempotency-key-escalate"

	var (
		response = common.MetadataResponse{
			HttpMetadata: mocks.HttpMetadataStatusOk,
		}
	)

	cases := []struct {
		name             string
		disputeId        string
		request          disputes.EscalateDisputeRequest
		idempotencyKey   *string
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*common.MetadataResponse, error)
	}{
		{
			name:           "when request is correct then should return 200",
			disputeId:      "idsp_test_12345abcdefghijklmnop",
			request:        buildEscalateDisputeRequest(),
			idempotencyKey: &idempotencyKey,
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("PostWithContext", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(4).(*common.MetadataResponse)
						*respMapping = response
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
			},
		},
		{
			name:           "when dispute not found then return 404",
			disputeId:      "idsp_not_found",
			request:        buildEscalateDisputeRequest(),
			idempotencyKey: nil,
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
			checker: func(response *common.MetadataResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusNotFound, chkErr.StatusCode)
			},
		},
		{
			name:           "when request is invalid then return 422",
			disputeId:      "idsp_test_12345abcdefghijklmnop",
			request:        disputes.EscalateDisputeRequest{},
			idempotencyKey: nil,
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
							ErrorCodes: []string{"justification_required"},
						},
					})
			},
			checker: func(response *common.MetadataResponse, err error) {
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

			tc.checker(client.EscalateDispute(tc.disputeId, tc.request, tc.idempotencyKey))
		})
	}
}

// # common methods

func buildCreateDisputeRequest() disputes.CreateDisputeRequest {
	amount := int64(1000)
	return disputes.CreateDisputeRequest{
		TransactionId:        "trx_test_abcdefghijklmnopqr",
		Reason:               "4837",
		Amount:               &amount,
		PresentmentMessageId: "msg_test_abcdefghijklmnopqr",
		Justification:        "Customer dispute",
		Evidence: []disputes.DisputeEvidence{
			{
				Name:        "receipt.pdf",
				Content:     "SGVsbG8gV29ybGQ=",
				Description: "Transaction receipt",
			},
		},
	}
}

func buildEscalateDisputeRequest() disputes.EscalateDisputeRequest {
	amount := int64(500)
	return disputes.EscalateDisputeRequest{
		Justification: "Escalating due to additional evidence",
		Amount:        &amount,
		AdditionalEvidence: []disputes.DisputeEvidence{
			{
				Name:        "additional_evidence.pdf",
				Content:     "QWRkaXRpb25hbCBFdmlkZW5jZQ==",
				Description: "Additional supporting documentation",
			},
		},
	}
}

func assertIssuingDisputeSuccess(t *testing.T, expected *disputes.IssuingDisputeResponse) func(*disputes.IssuingDisputeResponse, error) {
	return func(response *disputes.IssuingDisputeResponse, err error) {
		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.Equal(t, expected.Id, response.Id)
		assert.Equal(t, expected.Reason, response.Reason)
		assert.Equal(t, expected.Status, response.Status)
		assert.Equal(t, expected.TransactionId, response.TransactionId)
	}
}
