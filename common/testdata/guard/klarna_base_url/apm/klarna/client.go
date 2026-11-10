package klarna

type Client struct {
	path string
}

func (c *Client) getBaseUrl() string {
	return c.path
}
