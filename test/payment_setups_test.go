package test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/payments"
	"github.com/checkout/checkout-sdk-go/payments/setups"
)

var (
	paymentSetupRequest = setups.PaymentSetupRequest{
		Amount:      1000,
		Currency:    common.GBP,
		PaymentType: payments.Regular,
		Reference:   "PS-TEST-REF-001",
		Description: "Integration test payment setup",
		Customer: &setups.PaymentSetupCustomer{
			Email: &setups.PaymentSetupCustomerEmail{
				Address:  "integration-test@example.com",
				Verified: Bool(true),
			},
			Name:  "Integration Test Customer",
			Phone: Phone(),
		},
		ProcessingChannelId: os.Getenv("CHECKOUT_PROCESSING_CHANNEL_ID"),
		PaymentMethods: &setups.PaymentMethods{
			Klarna: &setups.KlarnaPaymentMethod{
				PaymentMethodBase: setups.PaymentMethodBase{
					Status:         "available",
					Flags:          []string{"supported"},
					Initialization: setups.PaymentMethodInitializationEnabled,
				},
				AccountHolder: &setups.KlarnaAccountHolder{
					BillingAddress: Address(),
				},
				PaymentMethodOptions: &setups.PaymentMethodOptions{
					Sdk: &setups.PaymentMethodOption{
						Id:     "sdk_option_test",
						Status: "active",
						Flags:  []string{"supported"},
						Action: &setups.PaymentMethodAction{
							Type:        "sdk",
							ClientToken: "test_client_token",
							SessionId:   "test_session_id",
						},
					},
				},
			},
		},
		Settings: &setups.PaymentSetupSettings{
			SuccessUrl: "https://example.com/success",
			FailureUrl: "https://example.com/failure",
		},
	}

	updatePaymentSetupRequest = setups.PaymentSetupRequest{
		Amount:      1500,
		Currency:    common.GBP,
		PaymentType: payments.Regular,
		Reference:   "PS-TEST-REF-001-UPDATED",
		Description: "Updated integration test payment setup",
		Customer: &setups.PaymentSetupCustomer{
			Email: &setups.PaymentSetupCustomerEmail{
				Address:  "integration-test-updated@example.com",
				Verified: Bool(true),
			},
			Name:  "Updated Integration Test Customer",
			Phone: Phone(),
		},
		ProcessingChannelId: os.Getenv("CHECKOUT_PROCESSING_CHANNEL_ID"),
		Settings: &setups.PaymentSetupSettings{
			SuccessUrl: "https://example.com/success-updated",
			FailureUrl: "https://example.com/failure-updated",
		},
	}
)

func TestCreatePaymentSetup(t *testing.T) {
	cases := []struct {
		name    string
		request setups.PaymentSetupRequest
		checker func(response *setups.PaymentSetupResponse, err error)
	}{
		{
			name:    "when payment setup request is valid then create payment setup",
			request: paymentSetupRequest,
			checker: func(response *setups.PaymentSetupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.NotEmpty(t, response.Id)
				assert.Equal(t, paymentSetupRequest.Amount, response.Amount)
				assert.Equal(t, paymentSetupRequest.Currency, response.Currency)
				assert.Equal(t, paymentSetupRequest.Reference, response.Reference)
				assert.Equal(t, paymentSetupRequest.Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.NotNil(t, response.PaymentMethods)
				assert.NotNil(t, response.Settings)
			},
		},
	}

	client := DefaultApi().PaymentSetups

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.CreatePaymentSetup(tc.request))
		})
	}
}

func TestUpdatePaymentSetup(t *testing.T) {
	// First create a payment setup
	createResponse, err := DefaultApi().PaymentSetups.CreatePaymentSetup(paymentSetupRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createResponse)
	setupId := createResponse.Id

	cases := []struct {
		name    string
		setupId string
		request setups.PaymentSetupRequest
		checker func(response *setups.PaymentSetupResponse, err error)
	}{
		{
			name:    "when payment setup update request is valid then update payment setup",
			setupId: setupId,
			request: updatePaymentSetupRequest,
			checker: func(response *setups.PaymentSetupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, setupId, response.Id)
				assert.Equal(t, updatePaymentSetupRequest.Amount, response.Amount)
				assert.Equal(t, updatePaymentSetupRequest.Currency, response.Currency)
				assert.Equal(t, updatePaymentSetupRequest.Reference, response.Reference)
				assert.Equal(t, updatePaymentSetupRequest.Description, response.Description)
			},
		},
	}

	client := DefaultApi().PaymentSetups

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.UpdatePaymentSetup(tc.setupId, tc.request))
		})
	}
}

func TestGetPaymentSetup(t *testing.T) {
	// First create a payment setup
	createResponse, err := DefaultApi().PaymentSetups.CreatePaymentSetup(paymentSetupRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createResponse)
	setupId := createResponse.Id

	cases := []struct {
		name    string
		setupId string
		checker func(response *setups.PaymentSetupResponse, err error)
	}{
		{
			name:    "when setup id is valid then get payment setup",
			setupId: setupId,
			checker: func(response *setups.PaymentSetupResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, setupId, response.Id)
				assert.Equal(t, paymentSetupRequest.Amount, response.Amount)
				assert.Equal(t, paymentSetupRequest.Currency, response.Currency)
				assert.Equal(t, paymentSetupRequest.Reference, response.Reference)
				assert.Equal(t, paymentSetupRequest.Description, response.Description)
				assert.NotNil(t, response.Customer)
				assert.NotNil(t, response.PaymentMethods)
				assert.NotNil(t, response.Settings)
			},
		},
		{
			name:    "when setup id is invalid then return error",
			setupId: "ps_invalid_setup_id",
			checker: func(response *setups.PaymentSetupResponse, err error) {
				assert.NotNil(t, err)
				assert.Nil(t, response)
			},
		},
	}

	client := DefaultApi().PaymentSetups

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.GetPaymentSetup(tc.setupId))
		})
	}
}

func TestConfirmPaymentSetup(t *testing.T) {
	t.Skip("Confirmation requires a valid payment method option ID from a real setup")

	// First create a payment setup
	createResponse, err := DefaultApi().PaymentSetups.CreatePaymentSetup(paymentSetupRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createResponse)
	setupId := createResponse.Id

	// Get a valid payment method option ID from the setup (would need real API response)
	paymentMethodOptionId := "pmo_real_option_id"

	cases := []struct {
		name                  string
		setupId               string
		paymentMethodOptionId string
		checker               func(response *setups.PaymentSetupConfirmResponse, err error)
	}{
		{
			name:                  "when setup id and option id are valid then confirm payment setup",
			setupId:               setupId,
			paymentMethodOptionId: paymentMethodOptionId,
			checker: func(response *setups.PaymentSetupConfirmResponse, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, 200, response.HttpMetadata.StatusCode)
				assert.Equal(t, setupId, response.Id)
				assert.NotEmpty(t, response.ActionId)
				assert.NotNil(t, response.Status)
				assert.NotNil(t, response.Source)
			},
		},
	}

	client := DefaultApi().PaymentSetups

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.checker(client.ConfirmPaymentSetup(tc.setupId, tc.paymentMethodOptionId))
		})
	}
}

func TestPaymentSetupFullWorkflow(t *testing.T) {
	t.Skip("Full workflow test - requires real payment method processing")

	client := DefaultApi().PaymentSetups

	// 1. Create payment setup
	createResponse, err := client.CreatePaymentSetup(paymentSetupRequest)
	assert.Nil(t, err)
	assert.NotNil(t, createResponse)
	setupId := createResponse.Id

	// 2. Get payment setup to verify creation
	getResponse, err := client.GetPaymentSetup(setupId)
	assert.Nil(t, err)
	assert.NotNil(t, getResponse)
	assert.Equal(t, setupId, getResponse.Id)

	// 3. Update payment setup
	updateResponse, err := client.UpdatePaymentSetup(setupId, updatePaymentSetupRequest)
	assert.Nil(t, err)
	assert.NotNil(t, updateResponse)
	assert.Equal(t, updatePaymentSetupRequest.Amount, updateResponse.Amount)

	// 4. Confirm payment setup (requires real payment method option)
	paymentMethodOptionId := "pmo_real_option_from_previous_steps"
	confirmResponse, err := client.ConfirmPaymentSetup(setupId, paymentMethodOptionId)
	assert.Nil(t, err)
	assert.NotNil(t, confirmResponse)
}
