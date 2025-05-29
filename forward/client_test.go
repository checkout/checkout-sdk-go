package forward

import (
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/errors"
	"github.com/checkout/checkout-sdk-go/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
)

func TestClient_ForwardAnApiRequest(t *testing.T) {
	var response = ForwardAnApiResponse{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		RequestId:    "fwd_01HK153X00VZ1K15Z3HYC0QGPN",
		DestinationResponse: &DestinationResponse{
			Status: 201,
			Headers: map[string][]string{
				"Cko-Request-Id": {"5fa7ee8c-f82d-4440-a6dc-e8c859b03235"},
				"Content-Type":   {"application/json"}},
			Body: `{\"id\": \"pay_mbabizu24mvu3mela5njyhpit4\", \"action_id\": \"act_mbabizu24mvu3mela5njyhpit4\", \"amount\": 6540, \"currency\": \"USD\", \"approved\": true, \"status\": \"Authorized\", \"auth_code\": \"770687\", \"response_code\": \"10000\", \"response_summary\": \"Approved\", \"_links\": {\"self\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4\"}, \"action\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/actions\"}, \"void\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/voids\"}, \"capture\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/captures\"}}}`,
		},
	}

	cases := []struct {
		name             string
		request          ForwardRequest
		getAuthorization func(*mock.Mock) mock.Call
		apiPost          func(*mock.Mock) mock.Call
		checker          func(*ForwardAnApiResponse, error)
	}{
		{
			name: "when request is valid then return a forward response",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiPost: func(m *mock.Mock) mock.Call {
				return *m.On("Post", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(3).(*ForwardAnApiResponse)
						*respMapping = response
					})
			},
			checker: func(response *ForwardAnApiResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, "fwd_01HK153X00VZ1K15Z3HYC0QGPN", response.RequestId)
				assert.NotNil(t, response.DestinationResponse)
				assert.Equal(t, 201, response.DestinationResponse.Status)
				assert.NotNil(t, response.DestinationResponse.Headers)
				assert.Equal(t, "5fa7ee8c-f82d-4440-a6dc-e8c859b03235", response.DestinationResponse.Headers["Cko-Request-Id"][0])
				assert.Equal(t, "application/json", response.DestinationResponse.Headers["Content-Type"][0])
				assert.Equal(t, `{\"id\": \"pay_mbabizu24mvu3mela5njyhpit4\", \"action_id\": \"act_mbabizu24mvu3mela5njyhpit4\", \"amount\": 6540, \"currency\": \"USD\", \"approved\": true, \"status\": \"Authorized\", \"auth_code\": \"770687\", \"response_code\": \"10000\", \"response_summary\": \"Approved\", \"_links\": {\"self\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4\"}, \"action\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/actions\"}, \"void\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/voids\"}, \"capture\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/captures\"}}}`, response.DestinationResponse.Body)
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
			checker: func(response *ForwardAnApiResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:    "when request is not correct then return error",
			request: ForwardRequest{},
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
			checker: func(response *ForwardAnApiResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAPIError)
				assert.Equal(t, http.StatusUnprocessableEntity, chkErr.StatusCode)
				assert.Equal(t, "invalid_request", chkErr.Data.ErrorType)
				assert.Contains(t, chkErr.Data.ErrorCodes, "company_or_individual_required")
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

			tc.checker(client.ForwardAnApiRequest(tc.request))
		})
	}

}

func TestClient_GetForwardRequest(t *testing.T) {
	var response = GetForwardResponse{
		HttpMetadata: mocks.HttpMetadataStatusOk,
		RequestId:    "fwd_01HK153X00VZ1K15Z3HYC0QGPN",
		Reference:    "ORD-5023-4E89",
		EntityId:     "ent_lp6h57qskk6ubewfk3pq4f2c2y",
		DestinationRequest: &DestinationRequestResponse{
			Url:    "https://example.com/payments",
			Method: "POST",
			Headers: map[string]string{
				"Authorization":   "***redacted***",
				"Idempotency-Key": "xe4fad12367dfgrds",
				"Content-Type":    "application/json",
			},
			Body: `{\"amount\": 1000, \"currency\": \"USD\", \"reference\": \"some_reference\", \"source\": {\"type\": \"card\", \"number\": \"{{card_number}}\", \"expiry_month\": \"{{card_expiry_month}}\", \"expiry_year\": \"{{card_expiry_year_yyyy}}\", \"name\": \"Ali Farid\"}, \"payment_type\": \"Regular\", \"authorization_type\": \"Final\", \"capture\": true, \"processing_channel_id\": \"pc_xxxxxxxxxxx\", \"risk\": {\"enabled\": false}, \"merchant_initiated\": true}`,
		},
		DestinationResponse: &DestinationResponse{
			Status: 201,
			Headers: map[string][]string{
				"Cko-Request-Id": {"5fa7ee8c-f82d-4440-a6dc-e8c859b03235"},
				"Content-Type":   {"application/json"}},
			Body: `{\"id\": \"pay_mbabizu24mvu3mela5njyhpit4\", \"action_id\": \"act_mbabizu24mvu3mela5njyhpit4\", \"amount\": 6540, \"currency\": \"USD\", \"approved\": true, \"status\": \"Authorized\", \"auth_code\": \"770687\", \"response_code\": \"10000\", \"response_summary\": \"Approved\", \"_links\": {\"self\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4\"}, \"action\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/actions\"}, \"void\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/voids\"}, \"capture\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/captures\"}}}`,
		},
		CreatedOn: time.Date(2024, 1, 2, 15, 4, 5, 0, time.UTC),
	}

	cases := []struct {
		name             string
		forwardId        string
		getAuthorization func(*mock.Mock) mock.Call
		apiGet           func(*mock.Mock) mock.Call
		checker          func(*GetForwardResponse, error)
	}{
		{
			name:      "when forward request exists then return response",
			forwardId: "fwd_01HK153X00VZ1K15Z3HYC0QGPN",
			getAuthorization: func(m *mock.Mock) mock.Call {
				return *m.On("GetAuthorization", mock.Anything).
					Return(&configuration.SdkAuthorization{}, nil)
			},
			apiGet: func(m *mock.Mock) mock.Call {
				return *m.On("Get", mock.Anything, mock.Anything, mock.Anything).
					Return(nil).
					Run(func(args mock.Arguments) {
						respMapping := args.Get(2).(*GetForwardResponse)
						*respMapping = response
					})
			},
			checker: func(response *GetForwardResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, http.StatusOK, response.HttpMetadata.StatusCode)
				assert.Equal(t, "fwd_01HK153X00VZ1K15Z3HYC0QGPN", response.RequestId)
				assert.Equal(t, "ORD-5023-4E89", response.Reference)
				assert.Equal(t, "ent_lp6h57qskk6ubewfk3pq4f2c2y", response.EntityId)
				assert.NotNil(t, response.DestinationRequest)
				assert.Equal(t, "https://example.com/payments", response.DestinationRequest.Url)
				assert.Equal(t, "POST", response.DestinationRequest.Method)
				assert.Equal(t, "***redacted***", response.DestinationRequest.Headers["Authorization"])
				assert.Equal(t, "xe4fad12367dfgrds", response.DestinationRequest.Headers["Idempotency-Key"])
				assert.Equal(t, "application/json", response.DestinationRequest.Headers["Content-Type"])
				assert.Equal(t, `{\"amount\": 1000, \"currency\": \"USD\", \"reference\": \"some_reference\", \"source\": {\"type\": \"card\", \"number\": \"{{card_number}}\", \"expiry_month\": \"{{card_expiry_month}}\", \"expiry_year\": \"{{card_expiry_year_yyyy}}\", \"name\": \"Ali Farid\"}, \"payment_type\": \"Regular\", \"authorization_type\": \"Final\", \"capture\": true, \"processing_channel_id\": \"pc_xxxxxxxxxxx\", \"risk\": {\"enabled\": false}, \"merchant_initiated\": true}`, response.DestinationRequest.Body)
				assert.NotNil(t, response.DestinationResponse)
				assert.Equal(t, 201, response.DestinationResponse.Status)
				assert.NotNil(t, response.DestinationResponse.Headers)
				assert.Equal(t, "5fa7ee8c-f82d-4440-a6dc-e8c859b03235", response.DestinationResponse.Headers["Cko-Request-Id"][0])
				assert.Equal(t, "application/json", response.DestinationResponse.Headers["Content-Type"][0])
				assert.Equal(t, `{\"id\": \"pay_mbabizu24mvu3mela5njyhpit4\", \"action_id\": \"act_mbabizu24mvu3mela5njyhpit4\", \"amount\": 6540, \"currency\": \"USD\", \"approved\": true, \"status\": \"Authorized\", \"auth_code\": \"770687\", \"response_code\": \"10000\", \"response_summary\": \"Approved\", \"_links\": {\"self\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4\"}, \"action\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/actions\"}, \"void\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/voids\"}, \"capture\": {\"href\": \"https://api.sandbox.checkout.com/payments/pay_mbabizu24mvu3mela5njyhpit4/captures\"}}}`, response.DestinationResponse.Body)
				assert.Equal(t, time.Date(2024, 1, 2, 15, 4, 5, 0, time.UTC), response.CreatedOn)
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
			checker: func(response *GetForwardResponse, err error) {
				assert.Nil(t, response)
				assert.NotNil(t, err)
				chkErr := err.(errors.CheckoutAuthorizationError)
				assert.Equal(t, "Invalid authorization type", chkErr.Error())
			},
		},
		{
			name:      "when forward request id not found then return error",
			forwardId: "not_found",
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
			checker: func(response *GetForwardResponse, err error) {
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

			tc.checker(client.GetForwardRequest(tc.forwardId))
		})
	}
}
