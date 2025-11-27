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
        AirlineData: &setups.AirlineData{
            Ticket: &setups.AirlineTicket{
                Number:                 "1234567890123",
                IssueDate:              "2024-01-15",
                IssuingCarrierCode:     "CK",
                TravelPackageIndicator: "A",
                TravelAgencyName:       "Checkout Travel Agents",
                TravelAgencyCode:       "91114362",
            },
            Passengers: []setups.AirlinePassenger{
                {
                    FirstName:   "John",
                    LastName:    "Doe",
                    DateOfBirth: "1990-10-31",
                    Address: &common.Address{
                        Country: common.GB,
                    },
                },
            },
            FlightLegDetails: []setups.FlightLegDetail{
                {
                    FlightNumber:      "BA1483",
                    CarrierCode:       "BA",
                    ClassOfTravelling: "W",
                    DepartureAirport:  "LHW",
                    DepartureDate:     "2025-10-13",
                    DepartureTime:     "18:30",
                    ArrivalAirport:    "JFK",
                    StopOverCode:      "X",
                    FareBasisCode:     "WUP14B",
                },
            },
        },
        AccommodationData: []setups.AccommodationData{
            {
                Name:             "Checkout Lodge",
                BookingReference: "REF9083748",
                CheckInDate:      "2025-04-11",
                CheckOutDate:     "2025-04-18",
                NumberOfRooms:    2,
                Guests: []setups.AccommodationGuest{
                    {
                        FirstName:   "Jia",
                        LastName:    "Tsang",
                        DateOfBirth: "1970-03-19",
                    },
                },
                Room: []setups.AccommodationRoom{
                    {
                        Rate:           42.3,
                        NumberOfNights: 5,
                    },
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