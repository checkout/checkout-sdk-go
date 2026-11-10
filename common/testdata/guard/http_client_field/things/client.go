package things

import "github.com/checkout/checkout-sdk-go/v3/client"

func Transport(a *client.ApiClient) interface{} {
	return a.HttpClient
}
