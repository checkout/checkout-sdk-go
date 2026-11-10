package things

import "github.com/checkout/checkout-sdk-go/v3/client"

func Base(a *client.ApiClient) string {
	return a.BaseUri
}
