package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	payment_sessions "github.com/checkout/checkout-sdk-go/payments/sessions"
)

var (
	paymentSessionsRequest = payment_sessions.PaymentSessionsRequest{
		Amount:    int64(2000),
		Currency:  common.GBP,
		Reference: "ORD-123A",
		Shipping: &payments.ShippingDetailsFlowHostedLinks{
			Address: Address(),
			Phone:   Phone(),
		},
		ThreeDsRequest: &payments.ThreeDsRequestFlowHostedLinks{
			Enabled:            false,
			AttemptN3D:         false,
			ChallengeIndicator: common.NoPreference,
			Exemption:          payments.LowValue,
			AllowUpgrade:       true,
		},
		Billing:    &payments.BillingInformation{Address: Address()},
		SuccessUrl: "https://example.com/payments/success",
		FailureUrl: "https://example.com/payments/fail",
	}
)

func TestRequestPaymentSessions(t *testing.T) {
	cases := []struct {
		name    string
		request payment_sessions.PaymentSessionsRequest
		checker func(response *payment_sessions.PaymentSessionsResponse, err error)
	}{
		{
			name:    "when payment context is valid the return a response",
			request: paymentSessionsRequest,
			checker: func(response *payment_sessions.PaymentSessionsResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 201, response.HttpMetadata.StatusCode)
				assert.NotNil(t, response.Id)
				if response.Links != nil {
					assert.NotNil(t, response.Links)
				}
			},
		},
	}

	client := DefaultApi().PaymentSessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestPaymentSessions(tc.request))
		})
	}
}

func TestRequestPaymentSessionsWithPayment(t *testing.T) {
	t.Skip("use on demand")

	paymentSessionsWithPaymentRequest := payment_sessions.PaymentSessionsWithPaymentRequest{
		SessionData: "test_session_data",
		Amount:      int64(2000),
		Currency:    common.GBP,
		Reference:   "ORD-123A",
		Shipping: &payments.ShippingDetailsFlowHostedLinks{
			Address: Address(),
			Phone:   Phone(),
		},
		ThreeDsRequest: &payments.ThreeDsRequestFlowHostedLinks{
			Enabled:            false,
			AttemptN3D:         false,
			ChallengeIndicator: common.NoPreference,
			Exemption:          payments.LowValue,
			AllowUpgrade:       true,
		},
		Billing:    &payments.BillingInformation{Address: Address()},
		SuccessUrl: "https://example.com/payments/success",
		FailureUrl: "https://example.com/payments/fail",
	}

	cases := []struct {
		name    string
		request payment_sessions.PaymentSessionsWithPaymentRequest
		checker func(response *payment_sessions.PaymentSessionPaymentResponse, err error)
	}{
		{
			name:    "when payment sessions with payment request is valid then return a response",
			request: paymentSessionsWithPaymentRequest,
			checker: func(response *payment_sessions.PaymentSessionPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.True(t, response.HttpMetadata.StatusCode == 201 || response.HttpMetadata.StatusCode == 202)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotEmpty(t, response.Type)
				assert.NotNil(t, response.PaymentSessionId)
				assert.NotNil(t, response.PaymentSessionSecret)
			},
		},
	}

	client := DefaultApi().PaymentSessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.RequestPaymentSessionsWithPayment(tc.request))
		})
	}
}

func TestSubmitPaymentSession(t *testing.T) {
	t.Skip("use on demand")

	submitPaymentSessionRequest := payment_sessions.SubmitPaymentSessionRequest{
		SessionData: "test_session_data",
		Amount:      int64(2000),
		Reference:   "ORD-123A",
		Items: []payments.Product{
			{
				Name:        "Test Product",
				Quantity:    1,
				UnitPrice:   2000,
				TotalAmount: 2000,
			},
		},
		ThreeDsRequest: &payments.ThreeDsRequestFlowHostedLinks{
			Enabled:            false,
			AttemptN3D:         false,
			ChallengeIndicator: common.NoPreference,
			Exemption:          payments.LowValue,
			AllowUpgrade:       true,
		},
		IpAddress:   "192.168.0.1",
		PaymentType: payments.Regular,
	}

	cases := []struct {
		name      string
		sessionId string
		request   payment_sessions.SubmitPaymentSessionRequest
		checker   func(response *payment_sessions.PaymentSessionPaymentResponse, err error)
	}{
		{
			name:      "when submit payment session request is valid then return a response",
			sessionId: "ps_test_session_id",
			request:   submitPaymentSessionRequest,
			checker: func(response *payment_sessions.PaymentSessionPaymentResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.True(t, response.HttpMetadata.StatusCode == 201 || response.HttpMetadata.StatusCode == 202)
				assert.NotNil(t, response.Id)
				assert.NotNil(t, response.Status)
				assert.NotEmpty(t, response.Type)
			},
		},
	}

	client := DefaultApi().PaymentSessions

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.SubmitPaymentSession(tc.sessionId, tc.request))
		})
	}
}
