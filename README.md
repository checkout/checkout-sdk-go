# Checkout.com Go SDK

![Status](https://img.shields.io/badge/status-beta-red.svg)

The official [Checkout][checkout] Go client library.

## Getting started

Make sure your project is using Go Modules (it will have a `go.mod` file in its root if it already is):

```sh
go mod init
```

```go
import (
    "github.com/checkout/checkout-sdk-go"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go toolchain will resolve and fetch the
checkout-sdk-go module automatically.

## API Keys

This SDK can be used with two different pair of API keys provided by Checkout. However, using different API keys imply
using specific API features. Please find in the table below the types of keys that can be used within this SDK.

| Account System | Public Key (example)                    | Secret Key (example)                    |
| -------------- | --------------------------------------- | --------------------------------------- |
| default        | pk_g650ff27-7c42-4ce1-ae90-5691a188ee7b | sk_gk3517a8-3z01-45fq-b4bd-4282384b0a64 |
| Previous       | pk_pkhpdtvabcf7hdgpwnbhw7r2uic          | sk_m73dzypy7cf3gf5d2xr4k7sxo4e          |

Note: sandbox keys have a `test_` or `sbox_` identifier, for current and previous api respectively.

If you don't have your own API keys, you can sign up for a test
account [here](https://www.checkout.com/get-test-account).

### OAuth

The SDK doesn't support any OAuth authentication flow natively, however it supports OAuth authorization tokens that can be used as API keys. For more information about OAuth please refer
to the official documentation.

## How to use the SDK

The SDK is structured by different modules and each module gives you access to different business features. All these modules can
be instantiated at once, or you can choose to create single modules separately.

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/client"
)

api := client.CheckoutApi(&secretKey, &publicKey, Sandbox) // or Production
var tokensClient = api.Tokens
```

```go
import (
"github.com/checkout/checkout-sdk-go"
"github.com/checkout/checkout-sdk-go/client"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var tokensClient = tokens.NewClient(*config)
```

### Tokens

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/tokens"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var client = tokens.NewClient(*config) // or api.Tokens
var card = &client.Card{
    Type:        common.Card,
    Number:      "4242424242424242",
    ExpiryMonth: 2,
    ExpiryYear:  2022,
    Name:        "Customer Name",
    CVV:         "100",
}
var request = &tokens.Request{
    Card: card,
}
response, err := client.Request(request)
```

### Payments

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var client = payments.NewClient(*config) // or api.Payments

var source = payments.TokenSource{
    Type:  common.Token.String(),
    Token: "tok_",
}
var request = &payments.Request{
    Source:   source,
    Amount:   "100",
    Currency: "USD",
    Reference: "Payment Reference",
    Customer: &payments.Customer{
        Email: "example@email.com",
        Name:  "First Name Last Name",
    },
    Metadata: map[string]string{
        "udf1": "User Define",
    },
}

idempotencyKey := checkout.NewIdempotencyKey()
params := checkout.Params{
    IdempotencyKey: &idempotencyKey,
}

response, err := client.Request(request, &params)
```

### Payment Detail

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var client = payments.NewClient(*config) // or api.Payments

response, err := client.Get("pay_")
```

### Actions

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var client = payments.NewClient(*config) // or api.Payments

response, err := client.Actions("pay_")
```

### Captures

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var client = payments.NewClient(*config) // or api.Payments

idempotencyKey := checkout.NewIdempotencyKey()
params := checkout.Params{
    IdempotencyKey: &idempotencyKey,
}

request := &client.CapturesRequest{
    Amount:    100,
    Reference: "Reference",
    Metadata: map[string]string{
        "udf1": "User Define",
    },
}
response, err := client.Captures("pay_", request, &params)
```

### Voids

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var client = payments.NewClient(*config) // or api.Payments

idempotencyKey := checkout.NewIdempotencyKey()
params := checkout.Params{
    IdempotencyKey: &idempotencyKey,
}

request := &client.VoidsRequest{
    Reference: "Reference",
    Metadata: map[string]string{
        "udf1": "User Define",
    },
}
response, err := client.Voids("pay_", request, &params)
```

### Refunds

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

config, err := checkout.SdkConfig(&secretKey, &publicKey, Sandbox) // or Production
var client = payments.NewClient(*config) // or api.Payments

request := &payments.RefundsRequest{
    Amount:    100,
    Reference: "Reference",
    Metadata: map[string]string{
        "udf1": "User Define",
    },
}

idempotencyKey := checkout.NewIdempotencyKey()
params := checkout.Params{
    IdempotencyKey: &idempotencyKey,
}

response, err := client.Refunds("pay_", request, &params)
```

More documentation related to Checkout API and the SDK is available at:

- [API Reference](https://api-reference.checkout.com/)
- [API Reference for our previous api reference](https://api-reference.checkout.com/previous/)
- [Docs](https://docs.checkout.com/)
- [Our previous docs](https://www.checkout.com/docs/previous)

The execution of integration tests require the following environment variables set in your system:

- Variables: `CHECKOUT_FOUR_PUBLIC_KEY` & `CHECKOUT_FOUR_SECRET_KEY`
- If you use our previous api set the variables: `CHECKOUT_PUBLIC_KEY` & `CHECKOUT_SECRET_KEY`

## Code of Conduct

Please refer to [Code of Conduct](CODE_OF_CONDUCT.md)

## Licensing

[MIT](LICENSE.md)
