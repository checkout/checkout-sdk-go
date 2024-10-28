# Checkout.com Golang SDK

[![build-status](https://github.com/checkout/checkout-sdk-go/workflows/build-master/badge.svg)](https://github.com/checkout/checkout-sdk-go/actions/workflows/build-master.yml)
![CodeQL](https://github.com/checkout/checkout-sdk-go/workflows/CodeQL/badge.svg)

[![build-status](https://github.com/checkout/checkout-sdk-go/workflows/build-release/badge.svg)](https://github.com/checkout/checkout-sdk-go/actions/workflows/build-release.yml)
[![GitHub release](https://img.shields.io/github/release/checkout/checkout-sdk-go.svg)](https://github.com/checkout/checkout-sdk-go/releases/)
[![Go Reference](https://pkg.go.dev/badge/github.com/checkout/checkout-sdk-go.svg)](https://pkg.go.dev/github.com/checkout/checkout-sdk-go)

[![GitHub license](https://img.shields.io/github/license/checkout/checkout-sdk-go.svg)](https://github.com/checkout/checkout-sdk-go/blob/master/LICENSE.md)

## Getting started

> **Version 1.0.0 is here!**
> <br/><br/>
> We improved the initialization of SDK making it easier to understand the available options. <br/>
> Now `NAS` accounts are the default instance for the SDK and `ABC` structure was moved to a `previous` prefixes. <br/>

### Module installer
Make sure your project is using Go Modules:
```sh
go get github.com/checkout/checkout-sdk-go@{version}
```
Then import the library into your code:
```sh
import "github.com/checkout/checkout-sdk-go"
```

### :rocket: Please check in [GitHub releases](https://github.com/checkout/checkout-sdk-go/releases) for all the versions available.

### :book: Checkout our official documentation.

* [Official Docs (Default)](https://docs.checkout.com/)
* [Official Docs (Previous)](https://docs.checkout.com/previous)

### :books: Check out our official API documentation guide, where you can also find more usage examples.

* [API Reference (Default)](https://api-reference.checkout.com/)
* [API Reference (Previous)](https://api-reference.checkout.com/previous)

## How to use the SDK

This SDK can be used with two different pair of API keys provided by Checkout. However, using different API keys imply
using specific API features. </br>
Please find in the table below the types of keys that can be used within this SDK.

| Account System | Public Key (example)                    | Secret Key (example)                    |
|----------------|-----------------------------------------|-----------------------------------------|
| Default        | pk_pkhpdtvabcf7hdgpwnbhw7r2uic          | sk_m73dzypy7cf3gf5d2xr4k7sxo4e          |
| Previous       | pk_g650ff27-7c42-4ce1-ae90-5691a188ee7b | sk_gk3517a8-3z01-45fq-b4bd-4282384b0a64 |

Note: sandbox keys have a `sbox_` or `test_` identifier, for Default and Previous accounts respectively.

If you don't have your own API keys, you can sign up for a test
account [here](https://www.checkout.com/get-test-account).

**PLEASE NEVER SHARE OR PUBLISH YOUR CHECKOUT CREDENTIALS.**

### Default

Default keys client instantiation can be done as follows:

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/configuration"
)

api, err := checkout.Builder().
                     StaticKeys().
                     WithEnvironment(configuration.Sandbox()).
                     WithSecretKey("secret_key").
                     WithPublicKey("public_key"). // optional, only required for operations related with tokens
                     Build()
```

### Default OAuth

The SDK supports client credentials OAuth, when initialized as follows:

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/configuration"
)

api, err := checkout.Builder().
                     OAuth().
                     WithAuthorizationUri("https://access.sandbox.checkout.com/connect/token"). // optional, custom authorization URI
                     WithClientCredentials("client_id", "client_secret").
                     WithEnvironment(configuration.Sandbox()).
                     WithScopes(getOAuthScopes()).
                     Build()
```

### Previous

If your pair of keys matches the previous system type, this is how the SDK should be used:

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/configuration"
)

api, err := checkout.Builder().
                     Previous().
                     WithEnvironment(configuration.Sandbox()).
                     WithSecretKey("secret_key").
                     WithPublicKey("public_key"). // optional, only required for operations related with tokens
                     Build()
```

Then just get any client, and start making requests:

```go
import (
    "github.com/checkout/checkout-sdk-go/payments"
    "github.com/checkout/checkout-sdk-go/payments/nas"
)

request := nas.PaymentRequest{}
response, err := api.Payments.RequestPayment(request)
```

## Error Handling

All the API responses that do not fall in the 2** status codes will return a `errors.CheckoutApiError`. The
error encapsulates the `StatusCode`, `Status` and a the `ErrorDetails`, if available.

## Custom Http Client
Go SDK supports your own configuration for `http client` using `http.Client` from the standard library. You can pass it through when instantiating the SDK as follows:

```go
import (
    "net/http"
    
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/configuration"
)

httpClient := http.Client{
    Timeout: time.Duration(20) * time.Millisecond,
}

api, err := checkout.Builder().
                     StaticKeys().
                     WithEnvironment(configuration.Sandbox()).
                     WithHttpClient(&httpClient).
                     WithSecretKey("secret_key")).
                     WithPublicKey("public_key")). // optional, only required for operations related with tokens
                     Build()
```

## Logging

The SDK supports custom Log provider. You can provide your log configuration via SDK initialization. By default, the SDK uses the `log` package from the standard library.

```go
import (
    "log"	
	
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/configuration"
)

logger := log.New(os.Stderr, "checkout-sdk-go - ", log.LstdFlags)

api, err := checkout.Builder().
                     StaticKeys().
                     WithEnvironment(configuration.Sandbox()).
                     WithSecretKey("secret_key")).
                     WithPublicKey("public_key")). // optional, only required for operations related with tokens
                     WithLogger(logger) // your own custom configuration
                     Build()
```

## Custom Environment
In case that you want to use an integrator or mock server, you can specify your own URI configuration as follows:

```go
import (
    "github.com/checkout/checkout-sdk-go"
    "github.com/checkout/checkout-sdk-go/configuration"
)

environment := configuration.NewEnvironment(
	"https://the.base.uri/", // the uri for all CKO operations 
	"https://the.oauth.uri/connect/token", // the uri used for OAUTH authorization, only required for OAuth operations 
	"https://the.files.uri/", // the uri used for Files operations, only required for Accounts module 
	"https://the.transfers.uri/", // the uri used for Transfer operations, only required for Transfers module 
	"https://the.balances.uri/", // the uri used for Balances operations, only required for Balances module false 
)

api, err := checkout.Builder().
                     StaticKeys().
                     WithEnvironment(environment).
                     WithSecretKey("secret_key")).
                     WithPublicKey("public_key")). // optional, only required for operations related with tokens
                     Build()
```

## Building from source

Once you check out the code from GitHub, the project can be built using:

```sh
go mod tidy

go build
```

The execution of integration tests require the following environment variables set in your system:

* For default account systems (NAS): `CHECKOUT_DEFAULT_PUBLIC_KEY` & `CHECKOUT_DEFAULT_SECRET_KEY`
* For default account systems (OAuth): `CHECKOUT_DEFAULT_OAUTH_CLIENT_ID` & `CHECKOUT_DEFAULT_OAUTH_CLIENT_SECRET`
* For Previous account systems (ABC): `CHECKOUT_PREVIOUS_PUBLIC_KEY` & `CHECKOUT_PREVIOUS_SECRET_KEY`

## Telemetry
Request telementry is enabled by default in the Go SDK. Request latency is included in the telemetry data. Recording the request latency allows Checkout.com to continuously monitor and imporove the merchant experience.

Request telemetry can be disabled by opting out during checkout_sdk_builder builder step:

```
api := checkout.Builder().
        Previous().
		WithSecretKey("CHECKOUT_PREVIOUS_SECRET_KEY").
		WithEnvironment(configuration.Sandbox()).
        WithEnableTelemetry(false).
		Build()
```

## Code of Conduct

Please refer to [Code of Conduct](CODE_OF_CONDUCT.md)

## Licensing

[MIT](LICENSE.md)
