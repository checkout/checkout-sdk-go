package things

import (
	"github.com/checkout/checkout-sdk-go/v3/common"

	"github.com/checkout/checkout-sdk-go/v3/client"
)

type Client struct {
	apiClient client.HttpClient
}

func (c *Client) Get(id string) (*common.MetadataResponse, error) {
	var response common.MetadataResponse
	err := c.apiClient.Get("/things/"+id, nil, &response)
	return &response, err
}
