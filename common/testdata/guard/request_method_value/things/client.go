package things

import "github.com/checkout/checkout-sdk-go/v3/client"

type Client struct {
	apiClient client.HttpClient
}

func (c *Client) Getter() interface{} {
	return c.apiClient.Get
}
