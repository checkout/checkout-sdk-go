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
	err := c.apiClient.Get(common.BuildPath("things", id), nil, &response)
	id, err = common.EncodePathSegment(id)
	if err != nil {
		return nil, err
	}
	return &response, err
}
