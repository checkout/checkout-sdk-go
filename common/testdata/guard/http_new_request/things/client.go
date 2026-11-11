package things

import "net/http"

func Request(id string) (*http.Request, error) {
	return http.NewRequest("GET", "https://example.com/things/"+id, nil)
}
