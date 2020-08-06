# Checkout.com Go

![Status](https://img.shields.io/badge/status-BETA-red.svg)

The official [Checkout][checkout] Go client library.

## :rocket: Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference checkout-sdk-go in a Go program with `import`:

``` go
import (
    "github.com/checkout/checkout-sdk-go"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the checkout-sdk-go module automatically.


## :book: Documentation

You can see the [SDK documentation here][api-docs].

For details on all the functionality in this library, see the [GoDoc][godoc]
documentation.

Below are a few simple examples:

### API

If you're dealing with multiple keys, it is recommended you use `client.API`.
This allows you to create as many clients as needed, each with their own
individual key.

```go
import (
    "github.com/shiuh-yaw-cko/checkout"
    "github.com/shiuh-yaw-cko/checkout/client"
)

api := &client.API{}
api.Init(secretKey, &publicKey)
```

### Tokens

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/tokens"
)

idempotencyKey := checkout.NewIdempotencyKey()
con:q!fig, err := checkout.Create(secretKe, &publicKey, &idempotencyKey)
if err != nil {
    return
}
var client = tokens.NewClient(*config)
var card = &tokens.Card{
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

idempotencyKey := checkout.NewIdempotencyKey()
config, err := checkout.Create(secretKey, &publicKey, &idempotencyKey)
if err != nil {
    return
}
var client = payments.NewClient(*config)
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
response, err := client.Request(request)
```

### Payment Detail

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

idempotencyKey := checkout.NewIdempotencyKey()
config, err := checkout.Create(secretKey, &publicKey, &idempotencyKey)
if err != nil {
    return
}
var client = payments.NewClient(*config)
response, err := client.Get("pay_")
```

### Actions

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

idempotencyKey := checkout.NewIdempotencyKey()
config, err := checkout.Create(secretKey, &publicKey, &idempotencyKey)
if err != nil {
    return
}
var client = payments.NewClient(*config)
response, err := client.Actions("pay_")
```

### Captures

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

idempotencyKey := checkout.NewIdempotencyKey()
config, err := checkout.Create(secretKey, &publicKey, &idempotencyKey)
if err != nil {
    return
}
var client = payments.NewClient(*config)
request := &payments.CapturesRequest{
    Amount:    100,
    Reference: "Reference",
    Metadata: map[string]string{
        "udf1": "User Define",
    },
}
response, err := client.Captures("pay_", request)
```

### Voids

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

idempotencyKey := checkout.NewIdempotencyKey()
config, err := checkout.Create(secretKey, &publicKey, &idempotencyKey)
if err != nil {
    return
}
var client = payments.NewClient(*config)
request := &payments.VoidsRequest{
    Reference: "Reference",
    Metadata: map[string]string{
        "udf1": "User Define",
    },
}
response, err := client.Voids("pay_", request)
```

### Refunds

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/payments"
)

idempotencyKey := checkout.NewIdempotencyKey()
config, err := checkout.Create(secretKey, &publicKey, &idempotencyKey)
if err != nil {
    return
}
var client = payments.NewClient(*config)
request := &payments.RefundsRequest{
    Amount:    100,
    Reference: "Reference",
    Metadata: map[string]string{
        "udf1": "User Define",
    },
}
response, err := client.Refunds("pay_", request)
```

For any requests, bug or comments, please [open an issue][issues] or [submit a
pull request][pulls].

[issues]: https://github.com/checkout/checkout-sdk-go/issues/new
[pulls]: https://github.com/checkout/checkout-sdk-go/pulls
[api-docs]: https://api-reference.checkout.com/
[checkout]: https://checkout.com
[godoc]: http://godoc.org/github.com/checkout/checkout-sdk-go
