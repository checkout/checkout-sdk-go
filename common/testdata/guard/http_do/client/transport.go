package client

import "net/http"

func Send(c *http.Client, r *http.Request) (*http.Response, error) {
	return c.Do(r)
}
