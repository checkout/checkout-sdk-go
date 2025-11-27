# Payment Setups

The Payment Setups module provides functionality to manage payment setups in the Checkout.com ecosystem. This module allows you to create, update, retrieve, and confirm payment setups.

## Features

- **Create Payment Setup**: Create a new payment setup with various payment methods and configurations
- **Update Payment Setup**: Update an existing payment setup
- **Get Payment Setup**: Retrieve details of a specific payment setup
- **Confirm Payment Setup**: Confirm a payment setup with payment source information

## Usage

### Initialize the Client

```go
import (
    "github.com/checkout/checkout-sdk-go/nas"
    "github.com/checkout/checkout-sdk-go/configuration"
    "github.com/checkout/checkout-sdk-go/payments/setups"
)

credentials := configuration.NewDefaultKeysSdkCredentials(
    "sk_sbox_...", // Replace with your secret key
    "pk_sbox_...", // Replace with your public key
)

enableTelemetry := true
config := configuration.NewConfiguration(
    credentials, 
    &enableTelemetry, 
    configuration.Sandbox(),
    &http.Client{}, 
    nil,
)

api := nas.CheckoutApi(config)
```

### Create Payment Setup

```go
createRequest := setups.PaymentSetupRequest{
    Amount:      1000, // £10.00
    Currency:    common.GBP,
    Reference:   "ORDER-123",
    Description: "Test payment setup",
    Customer: &setups.PaymentSetupCustomer{
        Email: "customer@example.com",
        Name:  "John Doe",
    },
    ReturnUrl: "https://example.com/return",
    PaymentMethods: &setups.PaymentMethods{
        Klarna: &setups.KlarnaPaymentMethod{
            Options: &setups.PaymentMethodOptions{
                Initialization: &setups.PaymentMethodInitialization{
                    PaymentMethodActions: []setups.PaymentMethodAction{
                        {
                            Type: "initialize",
                        },
                    },
                },
            },
        },
    },
}

response, err := api.PaymentSetups.CreatePaymentSetup(createRequest)
```

### Get Payment Setup

```go
response, err := api.PaymentSetups.GetPaymentSetup("ps_123456789")
```

### Update Payment Setup

```go
updateRequest := setups.PaymentSetupRequest{
    Amount:      1500, // £15.00
    Currency:    common.GBP,
    Description: "Updated payment setup",
}

response, err := api.PaymentSetups.UpdatePaymentSetup("ps_123456789", updateRequest)
```

### Confirm Payment Setup

```go
confirmRequest := setups.PaymentSetupConfirmRequest{
    Source: &setups.PaymentSetupSource{
        Type: "card",
    },
    ThreeDs: &setups.PaymentSetupThreeDs{
        Enabled: true,
    },
}

paymentMethodOptionId := "pmo_123456789" // Payment method option ID from setup creation
response, err := api.PaymentSetups.ConfirmPaymentSetup("ps_123456789", paymentMethodOptionId, confirmRequest)
```

## Supported Payment Methods

The Payment Setups module supports various payment methods:

- **Bizum**: Spanish mobile payment system
- **Klarna**: Buy now, pay later service
- **Stcpay**: Saudi Telecom Company payment service
- **Tabby**: Buy now, pay later service for MENA region

## Payment Setup Statuses

- `Pending`: Payment setup has been created but not yet confirmed
- `Confirmed`: Payment setup has been confirmed with payment source
- `Completed`: Payment setup has been successfully processed
- `Expired`: Payment setup has expired
- `Cancelled`: Payment setup has been cancelled

## Error Handling

All methods return appropriate errors. Handle them as follows:

```go
response, err := api.PaymentSetups.CreatePaymentSetup(request)
if err != nil {
    switch e := err.(type) {
    case errors.CheckoutAuthorizationError:
        // Handle authorization error
    case errors.CheckoutAPIError:
        // Handle API error
        fmt.Printf("API Error: %s\n", e.Data.ErrorType)
    default:
        // Handle other errors
    }
}
```

## Industry-Specific Data

The module supports industry-specific data such as airline data for travel industry:

```go
request := setups.PaymentSetupRequest{
    Industry: &setups.PaymentSetupIndustry{
        Type: "airline",
        AirlineData: &setups.AirlineData{
            Ticket: &setups.AirlineTicket{
                Number:            "1234567890123",
                IssueDate:         "2024-01-15",
                IssuingCarrierCode: "CK",
            },
            Passenger: &setups.AirlinePassenger{
                Name: &setups.AirlinePassengerName{
                    First: "John",
                    Last:  "Doe",
                },
            },
        },
    },
}
```

## Testing

Run tests with:

```bash
go test ./payments/setups
```

## Links

- [Checkout.com API Documentation](https://api-reference.checkout.com/)
- [Go SDK Documentation](https://github.com/checkout/checkout-sdk-go)