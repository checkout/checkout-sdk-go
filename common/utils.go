package common

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func BuildPath(params ...string) string {
	var path string
	for _, s := range params {
		path += "/" + s
	}

	return path
}

// BuildQueryPath appends the encoded query values to path.
//
// An empty filter returns the path unchanged rather than appending a bare "?". That matters for
// the endpoints whose query filter is optional: before the 2026-09-02 pass the list-attempts
// endpoints used BuildPath and produced ".../attempts", and a caller passing a zero filter would
// otherwise have started getting ".../attempts?" instead. A trailing question mark with no query
// carries no information, so no caller loses anything.
func BuildQueryPath(path string, queryValues interface{}) (string, error) {
	values, err := query.Values(queryValues)
	if err != nil {
		return "", err
	}

	encoded := values.Encode()
	if encoded == "" {
		return path, nil
	}

	return fmt.Sprintf("%s?%s", path, encoded), nil
}

func EscapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func BuildDefaultClient() *http.Client {
	return &http.Client{Timeout: time.Duration(10) * time.Second}
}
