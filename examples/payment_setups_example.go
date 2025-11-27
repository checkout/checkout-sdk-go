package main

import (
	"fmt"
	"net/http"

	"github.com/checkout/checkout-sdk-go/common"
	"github.com/checkout/checkout-sdk-go/configuration"
	"github.com/checkout/checkout-sdk-go/nas"
	"github.com/checkout/checkout-sdk-go/payments/setups"
)

func main() {
	// Initialize the Checkout SDK with your credentials
	credentials := configuration.NewDefaultKeysSdkCredentials(
		"sk_sbox_...", // Replace with your secret key
		"pk_sbox_...", // Replace with your public key
	)

	enableTelemetry := true
	config := configuration.NewConfiguration(
		credentials,
		&enableTelemetry,
		configuration.Sandbox(), // Use Sandbox for testing
		&http.Client{},
		nil,
	)

	// Create the API client
	api := nas.CheckoutApi(config)

	// Example 1: Create a Payment Setup
	fmt.Println("=== Creating Payment Setup ===")

	createRequest := setups.PaymentSetupRequest{
		Amount:      1000, // £10.00
		Currency:    common.GBP,
		Reference:   "ORDER-123",
		Description: "Test payment setup",
		Customer: &setups.PaymentSetupCustomer{
			Email: &setups.PaymentSetupCustomerEmail{
				Address:  "customer@example.com",
				Verified: true,
			},
			Name: "John Doe",
		},
		PaymentMethods: &setups.PaymentMethods{
			Klarna: &setups.KlarnaPaymentMethod{
				Initialization: "disabled",
				PaymentMethodOptions: &setups.KlarnaPaymentMethodOptions{
					Sdk: &setups.KlarnaSDKOption{
						Id: "opt_123456789",
						Action: &setups.KlarnaSDKAction{
							Type:        "sdk",
							ClientToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.ewog",
							SessionId:   "0b1d9815-165e-42e2-8867-35bc03789e00",
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

	createResponse, err := api.PaymentSetups.CreatePaymentSetup(createRequest)
	if err != nil {
		fmt.Printf("Error creating payment setup: %v\n", err)
		return
	}

	fmt.Printf("Payment Setup created: ID=%s, Status=%s\n",
		createResponse.Id, createResponse.Status)

	// Example 2: Get Payment Setup Details
	fmt.Println("\n=== Getting Payment Setup Details ===")

	getResponse, err := api.PaymentSetups.GetPaymentSetup(createResponse.Id)
	if err != nil {
		fmt.Printf("Error getting payment setup: %v\n", err)
		return
	}

	fmt.Printf("Payment Setup details: ID=%s, Amount=%d, Currency=%s\n",
		getResponse.Id, getResponse.Amount, getResponse.Currency)

	// Example 3: Update Payment Setup
	fmt.Println("\n=== Updating Payment Setup ===")

	updateRequest := setups.PaymentSetupRequest{
		Amount:      1500, // Update to £15.00
		Currency:    common.GBP,
		Reference:   "ORDER-123-UPDATED",
		Description: "Updated test payment setup",
	}

	updateResponse, err := api.PaymentSetups.UpdatePaymentSetup(createResponse.Id, updateRequest)
	if err != nil {
		fmt.Printf("Error updating payment setup: %v\n", err)
		return
	}

	fmt.Printf("Payment Setup updated: ID=%s, Amount=%d\n",
		updateResponse.Id, updateResponse.Amount)

	// Example 4: Confirm Payment Setup
	fmt.Println("\n=== Confirming Payment Setup ===")

	paymentMethodOptionId := "pmo_123456789" // Payment method option ID from setup creation
	confirmResponse, err := api.PaymentSetups.ConfirmPaymentSetup(createResponse.Id, paymentMethodOptionId)
	if err != nil {
		fmt.Printf("Error confirming payment setup: %v\n", err)
		return
	}

	fmt.Printf("Payment Setup confirmed: ID=%s, Status=%s\n",
		confirmResponse.Id, confirmResponse.Status)

	fmt.Println("\n=== Payment Setups Example Complete ===")
}
